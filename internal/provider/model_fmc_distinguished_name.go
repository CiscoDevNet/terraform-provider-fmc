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

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"sort"
	"strings"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DistinguishedName struct {
	Id                types.String `tfsdk:"id"`
	Domain            types.String `tfsdk:"domain"`
	Name              types.String `tfsdk:"name"`
	Type              types.String `tfsdk:"type"`
	DistinguishedName types.String `tfsdk:"distinguished_name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionDistinguishedName = version.Must(version.NewVersion("7.4"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DistinguishedName) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/distinguishednames"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DistinguishedName) toBody(ctx context.Context, state DistinguishedName) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	body, _ = sjson.Set(body, "type", "DistinguishedName")
	if !data.DistinguishedName.IsNull() {
		body, _ = sjson.Set(body, "dn", data.DistinguishedName.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DistinguishedName) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("dn"); value.Exists() {
		data.DistinguishedName = types.StringValue(value.String())
	} else {
		data.DistinguishedName = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *DistinguishedName) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("dn"); value.Exists() && !data.DistinguishedName.IsNull() {
		data.DistinguishedName = types.StringValue(value.String())
	} else {
		data.DistinguishedName = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DistinguishedName) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin Clone

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// End of section. //template:end toBodyNonBulk

// Section below is generated&owned by "gen/generator.go". //template:begin findObjectsToBeReplaced

// End of section. //template:end findObjectsToBeReplaced

// Section below is generated&owned by "gen/generator.go". //template:begin clearItemIds

// End of section. //template:end clearItemIds

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyPutDelete

// End of section. //template:end toBodyPutDelete

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBody

// End of section. //template:end adjustBody

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBodyBulk

// End of section. //template:end adjustBodyBulk

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyOverrides

// End of section. //template:end toBodyOverrides

// Section below is generated&owned by "gen/generator.go". //template:begin synthesizeOverrides

// End of section. //template:end synthesizeOverrides

// The helpers below are hand-owned (they live outside of any generator section, so they
// survive `go generate`). They are shared by both the fmc_distinguished_name and
// fmc_distinguished_names Read() functions.
//
// The FMC DistinguishedName API accepts a comma-separated distinguished name in the
// request (for example "CN=example.com,OU=IT,O=Example,C=US") but returns it as a
// slash-separated value in a canonicalized attribute order (for example
// "C=US/CN=example.com/O=Example/OU=IT"). These helpers reconcile the API response with
// the user-provided value so that the provider does not report a spurious diff.

// normalizeDistinguishedName converts the slash-separated distinguished name returned
// by the FMC API into the comma-separated form that users supply in their configuration.
func normalizeDistinguishedName(apiValue string) string {
	return strings.ReplaceAll(apiValue, "/", ",")
}

// distinguishedNameComponents splits a distinguished name (comma- or slash-separated)
// into its trimmed relative-distinguished-name components.
func distinguishedNameComponents(dn string) []string {
	fields := strings.FieldsFunc(dn, func(r rune) bool {
		return r == ',' || r == '/'
	})
	components := make([]string, 0, len(fields))
	for _, f := range fields {
		if trimmed := strings.TrimSpace(f); trimmed != "" {
			components = append(components, trimmed)
		}
	}
	return components
}

// distinguishedNamesEquivalent reports whether two distinguished names describe the same
// set of components, ignoring their order and separator style. The FMC API reorders the
// attributes it stores, so two values that differ only by ordering are semantically equal.
func distinguishedNamesEquivalent(a, b string) bool {
	ca := distinguishedNameComponents(a)
	cb := distinguishedNameComponents(b)
	if len(ca) != len(cb) {
		return false
	}
	sort.Strings(ca)
	sort.Strings(cb)
	for i := range ca {
		if ca[i] != cb[i] {
			return false
		}
	}
	return true
}

// reconcileDistinguishedName returns the value to store in state for the distinguished
// name attribute. If the configured value (stateValue) describes the same components as
// the API value (after re-ordering), the user's original value is preserved to avoid a
// diff; otherwise the normalized (comma-separated) API value is returned.
func reconcileDistinguishedName(stateValue, apiValue string) string {
	normalized := normalizeDistinguishedName(apiValue)
	if stateValue != "" && distinguishedNamesEquivalent(stateValue, normalized) {
		return stateValue
	}
	return normalized
}
