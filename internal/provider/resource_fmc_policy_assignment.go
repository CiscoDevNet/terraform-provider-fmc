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
	"fmt"
	"net/url"
	"slices"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

/*
	Each device must have Access Control Policy and Health Policy attached.
	When policy assignment is updated, the device is automatically removed from previous attachment.
	Health policy is always assigned by POST request with only devices that need to be assigned. Adding already assigned devices would trigger re-assignment and re-deploy.
	Other policies with each POST/PUT request require all devices that should be attached.
*/

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &PolicyAssignmentResource{}
	_ resource.ResourceWithImportState = &PolicyAssignmentResource{}
)

func NewPolicyAssignmentResource() resource.Resource {
	return &PolicyAssignmentResource{}
}

type PolicyAssignmentResource struct {
	client *fmc.Client
}

func (r *PolicyAssignmentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policy_assignment"
}

func (r *PolicyAssignmentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages a Policy Assignment.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "Name of the FMC domain",
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this is always 'PolicyAssignment'").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"policy_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the policy to be assigned.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"policy_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the policy to be assigned.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"policy_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the policy to be assigned.").AddStringEnumDescription("FTDNatPolicy", "HealthPolicy", "AccessPolicy", "RAVpn", "FTDPlatformSettingsPolicy").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("FTDNatPolicy", "HealthPolicy", "AccessPolicy", "RAVpn", "FTDPlatformSettingsPolicy"),
				},
			},
			"after_destroy_policy_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the Policy to be assigned after this policy assignment is destroyed. Applicable for Health and Access Control Policies only.").String,
				Optional:            true,
			},
			"targets": schema.SetNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of devices to which the policy should be attached").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the device").String,
							Required:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of the device").AddStringEnumDescription("Device", "DeviceHAPair", "DeviceGroup").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("Device", "DeviceHAPair", "DeviceGroup"),
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the device").String,
							Required:            true,
						},
					},
				},
			},
		},
	}
}

func (r *PolicyAssignmentResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

func (r *PolicyAssignmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PolicyAssignment
	var res gjson.Result
	var err error

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !plan.Domain.IsNull() && plan.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(plan.Domain.ValueString()))
	}

	// make sure only one policy assignment is updated at a time
	policyMu.Lock()
	defer policyMu.Unlock()

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))
	// Check if this is Health Policy, as this is handled differently
	if plan.PolicyType.ValueString() == "HealthPolicy" {
		res, diags = r.createPolicyAssignment(ctx, plan, reqMods...)
		resp.Diagnostics.Append(diags...)
	} else {
		// Check if policy assignment already exists
		res, err = r.client.Get(plan.getPath()+"/"+url.QueryEscape(plan.PolicyId.ValueString()), reqMods...)
		if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
			// Policy assignment does not exist - need to create it
			tflog.Debug(ctx, fmt.Sprintf("%s: Policy assignment does not exist", plan.Id.ValueString()))
			res, diags = r.createPolicyAssignment(ctx, plan, reqMods...)
			if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
				return
			}
		} else if err != nil {
			// Failed to retrieve policy assignments
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
			return
		} else {
			// Policy assignment already exists - need to update it
			tflog.Debug(ctx, fmt.Sprintf("%s: Policy assignment already exists", plan.Id.ValueString()))
			res, diags = r.updatePolicyAssignment(ctx, res, plan, PolicyAssignment{}, plan, reqMods...)
			if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
				return
			}
		}
	}

	// Extract object ID (Equal to policy ID)
	plan.Id = types.StringValue(res.Get("id").String())
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *PolicyAssignmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PolicyAssignment

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !state.Domain.IsNull() && state.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(state.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	urlPath := state.getPath() + "/" + url.QueryEscape(state.Id.ValueString())
	res, err := r.client.Get(urlPath, reqMods...)

	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	imp, diags := helpers.IsFlagImporting(ctx, req)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// After `terraform import` we switch to a full read.
	if imp {
		state.fromBody(ctx, res)
	} else {
		state.fromBodyPartial(ctx, res)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end read

func (r *PolicyAssignmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state PolicyAssignment
	var toAdd, toRemove PolicyAssignment

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Read state
	diags = req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !plan.Domain.IsNull() && plan.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(plan.Domain.ValueString()))
	}

	// Compute list of targets that should be removed (exist in state, but not in plan)
	toRemove = plan
	toRemove.Targets = []PolicyAssignmentTargets{}
	for _, target := range state.Targets {
		if !plan.containsTarget(target) {
			toRemove.Targets = append(toRemove.Targets, target)
		}
	}

	// Compute list of targets that should be added (exist in plan, but not in state)
	toAdd = plan
	toAdd.Targets = []PolicyAssignmentTargets{}
	for _, target := range plan.Targets {
		if !state.containsTarget(target) {
			toAdd.Targets = append(toAdd.Targets, target)
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	// make sure only one policy assignment is updated at a time
	policyMu.Lock()
	defer policyMu.Unlock()

	// Check if this is Health Policy, as this is handled differently
	if state.PolicyType.ValueString() == "HealthPolicy" {
		// Check if there are any devices to be removed and policy `after destroy` is set
		if !state.AfterDestroyPolicyId.IsNull() && len(toRemove.Targets) > 0 {
			toRemove.PolicyId = state.AfterDestroyPolicyId
			_, diags = r.createPolicyAssignment(ctx, toRemove, reqMods...)
			resp.Diagnostics.Append(diags...)
		}
		if len(toAdd.Targets) > 0 {
			_, diags = r.createPolicyAssignment(ctx, toAdd, reqMods...)
			resp.Diagnostics.Append(diags...)
		}
		plan.PolicyName = types.StringNull()
	} else {
		// This is policy other than Health Policy
		// Check if Access Policy target needs to be re-assigned
		if state.PolicyType.ValueString() == "AccessPolicy" && len(toRemove.Targets) > 0 && !plan.AfterDestroyPolicyId.IsNull() {
			toRemove.PolicyId = plan.AfterDestroyPolicyId
			res, err := r.client.Get(plan.getPath()+"/"+url.QueryEscape(toRemove.PolicyId.ValueString()), reqMods...)
			if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
				// Policy assignment does not exist - need to create it
				tflog.Debug(ctx, fmt.Sprintf("%s: Policy assignment does not exist", plan.Id.ValueString()))
				_, diags = r.createPolicyAssignment(ctx, toRemove, reqMods...)
				if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
					return
				}
			} else if err != nil {
				// Failed to retrieve policy assignments
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
				return
			} else {
				// Policy assignment already exists - need to update it
				tflog.Debug(ctx, fmt.Sprintf("%s: Policy assignment already exists", plan.Id.ValueString()))
				_, diags = r.updatePolicyAssignment(ctx, res, toRemove, toRemove, PolicyAssignment{}, reqMods...)
				if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
					return
				}
			}
		}

		// Check if policy assignment exists
		// It may not exist if devices were reassigned to other policies outside of this resource
		urlPath := state.getPath() + "/" + url.QueryEscape(state.Id.ValueString())
		res, err := r.client.Get(urlPath, reqMods...)
		if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
			// Policy assignment does not exist - need to create it
			tflog.Debug(ctx, fmt.Sprintf("%s: Policy assignment does not exist", plan.Id.ValueString()))
			_, diags := r.createPolicyAssignment(ctx, plan, reqMods...)
			if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
				return
			}
		} else if err != nil {
			// Failed to retrieve policy assignments
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
			return
		} else {
			// Policy assignment already exists - need to update it (PUT)
			tflog.Debug(ctx, fmt.Sprintf("%s: Policy assignment already exists", plan.Id.ValueString()))
			_, diags := r.updatePolicyAssignment(ctx, res, plan, toRemove, toAdd, reqMods...)
			if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
				return
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *PolicyAssignmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// There is no DELETE endpoint for this resource. Every update happens through PUT request or re-assignment.
	var state PolicyAssignment

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !state.Domain.IsNull() && state.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(state.Domain.ValueString()))
	}

	// make sure only one policy assignment is updated at a time
	policyMu.Lock()
	defer policyMu.Unlock()

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	// Check if this is Health Policy, as this is handled differently
	if state.PolicyType.ValueString() == "HealthPolicy" {
		if state.AfterDestroyPolicyId.IsNull() {
			tflog.Debug(ctx, fmt.Sprintf("%s: No after destroy policy ID provided", state.Id.ValueString()))
		} else {
			state.PolicyId = state.AfterDestroyPolicyId
			_, diags := r.createPolicyAssignment(ctx, state, reqMods...)
			resp.Diagnostics.Append(diags...)
		}
	} else if state.PolicyType.ValueString() == "AccessPolicy" {
		if state.AfterDestroyPolicyId.IsNull() {
			// No 'after destroy' policy ID provided - just remove the policy assignment resource
			tflog.Debug(ctx, fmt.Sprintf("%s: No after destroy policy ID provided", state.Id.ValueString()))
		} else {
			// 'After destroy' policy ID provided - reassign devices to the new policy
			res, err := r.client.Get(state.getPath()+"/"+url.QueryEscape(state.AfterDestroyPolicyId.ValueString()), reqMods...)
			if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
				tflog.Debug(ctx, fmt.Sprintf("%s: After destroy policy assignment does not exist", state.Id.ValueString()))

				// Set desired policy to after destroy policy
				state.PolicyId = state.AfterDestroyPolicyId
				_, diags := r.createPolicyAssignment(ctx, state, reqMods...)
				if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
					return
				}
			} else if err != nil {
				// Failed to retrieve policy assignments
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
				return
			} else {
				// Policy assignment already exists - need to update it with destroy policy
				tflog.Debug(ctx, fmt.Sprintf("%s: Policy assignment already exists", state.Id.ValueString()))

				state.PolicyId = state.AfterDestroyPolicyId
				_, diags := r.updatePolicyAssignment(ctx, res, state, PolicyAssignment{}, state, reqMods...)
				resp.Diagnostics.Append(diags...)
			}
		}
	} else {
		// Any other policy assignment is optional, hence we can just remove that assignment
		res, err := r.client.Get(state.getPath()+"/"+url.QueryEscape(state.PolicyId.ValueString()), reqMods...)
		// If there is no error, policy exists and we can remove assignments from it
		if err == nil {
			// Policy assignment already exists - need to update it (remove configured devices, but keep non-managed ones)
			tflog.Debug(ctx, fmt.Sprintf("%s: Policy assignment exists", state.Id.ValueString()))

			_, diag := r.updatePolicyAssignment(ctx, res, state, state, PolicyAssignment{}, reqMods...)
			resp.Diagnostics.Append(diag...)
		}
	}

	if !resp.Diagnostics.HasError() {
		tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))
	}

	resp.State.RemoveResource(ctx)
}

// Section below is generated&owned by "gen/generator.go". //template:begin import

func (r *PolicyAssignmentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import

// createPolicyAssignment creates a new policy assignment using POST request
func (r *PolicyAssignmentResource) createPolicyAssignment(ctx context.Context, data PolicyAssignment, reqMods ...func(*fmc.Req)) (gjson.Result, diag.Diagnostics) {
	var diag diag.Diagnostics

	body := data.toBody(ctx, PolicyAssignment{})
	body, _ = sjson.Delete(body, "dummy_after_destroy_policy_id")

	res, err := r.client.Post(data.getPath(), body, reqMods...)
	if err != nil {
		diag.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return res, diag
	}
	return res, nil
}

// updatePolicyAssignment updates exsiting policy assigment using PUT request
func (r *PolicyAssignmentResource) updatePolicyAssignment(ctx context.Context, res gjson.Result, data, toRemove, toAdd PolicyAssignment, reqMods ...func(*fmc.Req)) (gjson.Result, diag.Diagnostics) {
	var diag diag.Diagnostics

	if len(toRemove.Targets) == 0 && len(toAdd.Targets) == 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: No changes to policy assignment requested", data.Id.ValueString()))
		return res, diag
	}

	updateBody := res.String()
	updateBody, _ = sjson.Delete(updateBody, "links")

	// Remove targets from updateBody based on `toRemove` list
	if len(toRemove.Targets) > 0 {
		// Get list of IDs from devices from state
		var stateTargets = make([]string, len(toRemove.Targets))
		for _, target := range toRemove.Targets {
			stateTargets = append(stateTargets, target.Id.ValueString())
		}

		existingTargets := gjson.Get(updateBody, "targets").Array()
		updateBody, _ = sjson.Set(updateBody, "targets", []any{})
		for _, target := range existingTargets {
			if !slices.Contains(stateTargets, target.Get("id").String()) {
				updateBody, _ = sjson.Set(updateBody, "targets.-1", map[string]string{
					"id":   target.Get("id").String(),
					"type": target.Get("type").String(),
					"name": target.Get("name").String(),
				})
			}
		}
	}

	// Add targets to updateBody based on `toAdd` list
	if len(toAdd.Targets) > 0 {
		for _, target := range toAdd.Targets {
			if !res.Get(fmt.Sprintf("targets.#(id==%s)", target.Id.ValueString())).Exists() {
				updateBody, _ = sjson.Set(updateBody, "targets.-1", map[string]string{
					"id":   target.Id.ValueString(),
					"type": target.Type.ValueString(),
					"name": target.Name.ValueString(),
				})
			}
		}
	}

	// Update object
	resPut, err := r.client.Put(data.getPath()+"/"+url.QueryEscape(data.PolicyId.ValueString()), updateBody, reqMods...)
	if err != nil {
		diag.AddError("Client Error", fmt.Sprintf("Failed to update policy assignment (PUT), got error: %s, %s", err, res.String()))
		return resPut, diag
	}
	return resPut, diag
}

// Checks if given target is on the target list
func (r *PolicyAssignment) containsTarget(target PolicyAssignmentTargets) bool {
	for _, t := range r.Targets {
		if t.Id.ValueString() == target.Id.ValueString() {
			return true
		}
	}
	return false
}
