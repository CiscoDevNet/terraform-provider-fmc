// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// resolvePolicyDomain returns the name of the domain in which a policy exists. Policies may reside in a different
// domain than the device itself, hence a per-policy domain can be configured. If it is not set, device's domain is assumed.
func resolvePolicyDomain(policyDomain, deviceDomain types.String) string {
	if !policyDomain.IsNull() && policyDomain.ValueString() != "" {
		return policyDomain.ValueString()
	}
	return deviceDomain.ValueString()
}

// policyReqMods translates a domain name into request modifiers. Empty domain name means that no modifier is applied,
// hence the provider-level default domain is used.
func policyReqMods(domain string) []func(*fmc.Req) {
	reqMods := [](func(*fmc.Req)){}
	if domain != "" {
		reqMods = append(reqMods, fmc.DomainName(domain))
	}
	return reqMods
}

// displayDomain renders a domain name for diagnostics. Empty domain name means the provider-level default domain.
func displayDomain(domain string) string {
	if domain == "" {
		return "Global"
	}
	return domain
}

// policyDomainPath maps an attribute holding a policy id to the attribute holding the domain of that policy,
// eg. `nat_policy_id` to `nat_policy_domain`.
func policyDomainPath(policyPath path.Path) path.Path {
	return path.Root(strings.TrimSuffix(policyPath.String(), "_id") + "_domain")
}

// updatePolicyAssignment updates policy-to-device assignment of one specific device (UUID) and of one specific policy type
// (policyPath points to a different attribute for Access Policy, NAT Policy, Platform Settings Policy, etc.).
func updatePolicyAssignment(ctx context.Context, client *fmc.Client, deviceId string, deviceType string, policyPath path.Path, plan tfsdk.Plan, state tfsdk.State) diag.Diagnostics {
	var planPolicy, statePolicy types.String

	if diags := plan.GetAttribute(ctx, policyPath, &planPolicy); diags.HasError() {
		return diags
	}

	if diags := state.GetAttribute(ctx, policyPath, &statePolicy); diags.HasError() {
		return diags
	}

	tflog.Debug(ctx, fmt.Sprintf("policy assignment %s, id %s, policy planned %s, state %s", policyPath, deviceId, planPolicy, statePolicy))

	if statePolicy.Equal(planPolicy) {
		return nil
	}

	// A policy may exist in a different domain than the device itself. When the policy is being de-assigned, the
	// assignment lives in the domain recorded in the state, otherwise the planned domain applies.
	var policyDomain, deviceDomain types.String
	var diags diag.Diagnostics
	if planPolicy.IsNull() {
		diags.Append(state.GetAttribute(ctx, policyDomainPath(policyPath), &policyDomain)...)
		diags.Append(state.GetAttribute(ctx, path.Root("domain"), &deviceDomain)...)
	} else {
		diags.Append(plan.GetAttribute(ctx, policyDomainPath(policyPath), &policyDomain)...)
		diags.Append(plan.GetAttribute(ctx, path.Root("domain"), &deviceDomain)...)
	}
	if diags.HasError() {
		return diags
	}

	domain := resolvePolicyDomain(policyDomain, deviceDomain)
	reqMods := policyReqMods(domain)
	domainHint := fmt.Sprintf("Verify that `%s` is the domain in which the policy exists (set via `%s`, defaulting to the device's `domain`).", displayDomain(domain), policyDomainPath(policyPath))

	// Marginal risk of data race follows (GET/PUT).
	// Partial self-protection in case user specifies terraform -parallelism 2 or more:
	policyMu.Lock()
	defer policyMu.Unlock()

	if planPolicy.IsNull() {
		if policyPath.String() == "health_policy_id" {
			return nil
		}
		res, err := client.Get("/api/fmc_config/v1/domain/{DOMAIN_UUID}/assignment/policyassignments/"+url.QueryEscape(statePolicy.ValueString()), reqMods...)
		// If assignment does not exist, do noting
		if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
			return nil
		} else if err != nil {
			return diag.Diagnostics{diag.NewErrorDiagnostic(
				"Client Error",
				fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s\n%s", err, res.String(), domainHint),
			)}
		}
		targets := res.Get(fmt.Sprintf(`targets.#(id != "%s")#`, deviceId))
		body, err := sjson.SetRaw(res.String(), "targets", targets.Raw)
		tflog.Debug(ctx, fmt.Sprintf("policy assignment with device %s removed: %s", deviceId, res))
		if err != nil {
			return diag.Diagnostics{diag.NewAttributeErrorDiagnostic(
				path.Root("id"),
				"Internal Error",
				fmt.Sprintf("Failed to set JSON \"targets\", got error: %s, %s, %s", err, res, targets),
			)}
		}

		res, err = client.Put("/api/fmc_config/v1/domain/{DOMAIN_UUID}/assignment/policyassignments"+"/"+url.QueryEscape(statePolicy.ValueString()), body, reqMods...)
		if err != nil {
			return diag.Diagnostics{diag.NewErrorDiagnostic(
				"Client Error",
				fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s\n%s", err, res.String(), domainHint),
			)}
		}
		return nil
	}

	tflog.Debug(ctx, fmt.Sprintf("policy path: %s", policyPath.String()))
	if policyPath.String() == "health_policy_id" {
		stub, _ := sjson.Set("{}", "id", planPolicy.ValueString())
		stub, _ = sjson.Set(stub, "policy.id", planPolicy.ValueString())
		stub, _ = sjson.Set(stub, "policy.type", "HealthPolicy")
		stub, _ = sjson.Set(stub, "targets", []any{})
		stubdev, _ := sjson.Set("{}", "id", deviceId)
		stubdev, _ = sjson.Set(stubdev, "type", deviceType)
		stub, _ = sjson.SetRaw(stub, "targets.-1", stubdev)
		res, err := client.Post("/api/fmc_config/v1/domain/{DOMAIN_UUID}/assignment/policyassignments", stub, reqMods...)
		if err != nil {
			return diag.Diagnostics{diag.NewErrorDiagnostic(
				"Client Error",
				fmt.Sprintf("Failed to configure object (POST), got error: %s, %s\n%s", err, res.String(), domainHint),
			)}
		}
		return nil
	}

	res, err := client.Get("/api/fmc_config/v1/domain/{DOMAIN_UUID}/assignment/policyassignments"+"/"+url.QueryEscape(planPolicy.ValueString()), reqMods...)
	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
		stub, _ := sjson.Set("{}", "id", planPolicy.ValueString())
		stub, _ = sjson.Set(stub, "policy.id", planPolicy.ValueString())
		stub, _ = sjson.Set(stub, "targets", []any{})
		res = gjson.Parse(stub)
	} else if err != nil {
		return diag.Diagnostics{diag.NewErrorDiagnostic(
			"Client Error",
			fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()),
		)}
	}

	if res.Get(fmt.Sprintf(`targets.#(id=="%s")`, deviceId)).Exists() {
		tflog.Debug(ctx, fmt.Sprintf("target %s already assigned to policy %s", deviceId, planPolicy))
		return nil
	}

	polBody, err := sjson.Set(res.String(), `targets.-1`, map[string]any{
		"id":   deviceId,
		"type": deviceType,
	})
	if err != nil {
		return diag.Diagnostics{diag.NewAttributeErrorDiagnostic(
			path.Root("id"),
			"Internal Error",
			fmt.Sprintf("Failed to append to JSON list \"targets\", got error: %s, %s", err, res),
		)}
	}

	res, err = client.Put("/api/fmc_config/v1/domain/{DOMAIN_UUID}/assignment/policyassignments"+"/"+url.QueryEscape(planPolicy.ValueString()), polBody, reqMods...)
	if err != nil {
		return diag.Diagnostics{diag.NewErrorDiagnostic(
			"Client Error",
			fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s\n%s", err, res.String(), domainHint),
		)}
	}

	return nil
}
