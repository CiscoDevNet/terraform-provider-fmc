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
	"regexp"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &NetworkGroupOverridesResource{}
	_ resource.ResourceWithImportState = &NetworkGroupOverridesResource{}
)

func NewNetworkGroupOverridesResource() resource.Resource {
	return &NetworkGroupOverridesResource{}
}

type NetworkGroupOverridesResource struct {
	client *fmc.Client
}

func (r *NetworkGroupOverridesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_group_overrides"
}

func (r *NetworkGroupOverridesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages a Network Group Overrides.").String,

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
			"parent_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the parent Network Group object.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"parent_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of the parent Network Group object.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"overrides": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Override entries for the object.").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"target_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of the override target Device or Domain. Note that each target can be defined once only.").String,
							Required:            true,
						},
						"target_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of the target.").AddStringEnumDescription("Device", "Domain").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("Device", "Domain"),
							},
						},
						"description": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Description of the object.").String,
							Optional:            true,
						},
						"objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of network objects (Host, Network, Range, FQDN or Network Group).").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("ID of the network object.").String,
										Optional:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Name of the network object.").String,
										Optional:            true,
									},
								},
							},
						},
						"literals": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of literal values (Host or Network).").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"value": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("IP address or network in CIDR format. Please do not use /32 mask for host.").String,
										Optional:            true,
									},
								},
							},
						},
					},
				},
				Validators: []validator.List{
					listvalidator.SizeAtMost(1000),
				},
			},
		},
	}
}

func (r *NetworkGroupOverridesResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *NetworkGroupOverridesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan NetworkGroupOverrides

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))
	body := plan.toBodyOverrides(ctx, NetworkGroupOverrides{})
	res, err := r.client.Post(plan.getPath()+"?bulk=true", body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = plan.ParentId
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *NetworkGroupOverridesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state NetworkGroupOverrides

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

	urlPath := state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()) + "/overrides"
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
	res = state.synthesizeOverrides(ctx, res)

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

// Section below is generated&owned by "gen/generator.go". //template:begin update

func (r *NetworkGroupOverridesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state NetworkGroupOverrides

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))
	var res fmc.Res
	var err error
	planOwnedTargetIds := make(map[string]bool)
	for _, override := range plan.Overrides {
		planOwnedTargetIds[override.TargetId.ValueString()] = true
	}

	stateOwnedTargetIds := make(map[string]bool)
	for _, override := range state.Overrides {
		stateOwnedTargetIds[override.TargetId.ValueString()] = true
	}

	// DELETE
	// Delete objects that are present in state, but missing in plan
	var toDelete []NetworkGroupOverridesOverrides
	for _, override := range state.Overrides {
		if _, ok := planOwnedTargetIds[override.TargetId.ValueString()]; !ok {
			toDelete = append(toDelete, override)
		}
	}

	for idx, override := range toDelete {
		urlPath := state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()) + "?overrideTargetId=" + url.QueryEscape(override.TargetId.ValueString())
		res, err := r.client.Delete(urlPath, reqMods...)
		if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
			continue
		} else if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete override entry (DELETE), got error: %s, %s", err, res.String()))
			// On error, save state with items still existing on the server
			var remaining []NetworkGroupOverridesOverrides
			for _, o := range state.Overrides {
				if _, ok := planOwnedTargetIds[o.TargetId.ValueString()]; ok {
					remaining = append(remaining, o)
				}
			}
			remaining = append(remaining, toDelete[idx:]...)
			state.Overrides = remaining
			diags = resp.State.Set(ctx, &state)
			resp.Diagnostics.Append(diags...)
			return
		}
	}

	// UPDATE
	// Check for updates to existing objects (that are present in both plan and state)
	stateTargetIdToIdx := make(map[string]int)
	for i, override := range state.Overrides {
		stateTargetIdToIdx[override.TargetId.ValueString()] = i
	}

	updatedTargetIds := make(map[string]bool)
	for pi, override := range plan.Overrides {
		si, ok := stateTargetIdToIdx[override.TargetId.ValueString()]
		if !ok {
			continue
		}

		// Compare plan and state at respective list indices
		var pv, sv attr.Value
		diags = req.Plan.GetAttribute(ctx, path.Root("overrides").AtListIndex(pi), &pv)
		if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
			return
		}
		diags = req.State.GetAttribute(ctx, path.Root("overrides").AtListIndex(si), &sv)
		if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
			return
		}

		if pv.Equal(sv) {
			continue
		}

		// Build body for single override update
		toUpdate := plan
		toUpdate.Overrides = []NetworkGroupOverridesOverrides{override}
		updateBody := toUpdate.toBodyOverrides(ctx, NetworkGroupOverrides{})
		singleBody := gjson.Get(updateBody, "0").Raw

		urlPath := plan.getPath() + "/" + url.QueryEscape(plan.Id.ValueString())
		res, err = r.client.Put(urlPath, singleBody, reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update override entry (PUT), got error: %s, %s", err, res.String()))
			// On error, save state with items still existing on the server:
			// only state overrides that were not scheduled for deletion (intersection of state and plan).
			// New items from the plan were not yet created, so they must not be in state.
			// For successfully updated items, use plan values (server has new values).
			// For not-yet-updated items, use state values (server still has old values).
			planTargetIdToOverride := make(map[string]NetworkGroupOverridesOverrides)
			for _, o := range plan.Overrides {
				planTargetIdToOverride[o.TargetId.ValueString()] = o
			}
			var remaining []NetworkGroupOverridesOverrides
			for _, o := range state.Overrides {
				tid := o.TargetId.ValueString()
				if _, ok := planOwnedTargetIds[tid]; !ok {
					continue
				}
				if updatedTargetIds[tid] {
					remaining = append(remaining, planTargetIdToOverride[tid])
				} else {
					remaining = append(remaining, o)
				}
			}
			state.Overrides = remaining
			diags = resp.State.Set(ctx, &state)
			resp.Diagnostics.Append(diags...)
			return
		}
		updatedTargetIds[override.TargetId.ValueString()] = true
	}

	// CREATE
	// Create objects (that are present in plan, but missing in state)
	toCreate := plan
	toCreate.Overrides = []NetworkGroupOverridesOverrides{}
	for _, override := range plan.Overrides {
		if _, ok := stateOwnedTargetIds[override.TargetId.ValueString()]; !ok {
			toCreate.Overrides = append(toCreate.Overrides, override)
		}
	}

	if len(toCreate.Overrides) > 0 {
		body := toCreate.toBodyOverrides(ctx, NetworkGroupOverrides{})
		res, err = r.client.Post(toCreate.getPath()+"?bulk=true", body, reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
			// On error, save state with items still existing on the server:
			// deletes and updates succeeded, so plan items that existed in state are on the server
			// with their updated values. New items from the plan were not created, so exclude them.
			var remaining []NetworkGroupOverridesOverrides
			for _, o := range plan.Overrides {
				if _, ok := stateOwnedTargetIds[o.TargetId.ValueString()]; ok {
					remaining = append(remaining, o)
				}
			}
			plan.Overrides = remaining
			diags = resp.State.Set(ctx, &plan)
			resp.Diagnostics.Append(diags...)
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *NetworkGroupOverridesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworkGroupOverrides

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	var err error
	var errElement int
	var res fmc.Res

	for idx, override := range state.Overrides {
		urlPath := state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()) + "?overrideTargetId=" + url.QueryEscape(override.TargetId.ValueString())
		res, err = r.client.Delete(urlPath, reqMods...)
		if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
			err = nil
			continue
		} else if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete override entry (DELETE), got error: %s, %s", err, res.String()))
			errElement = idx
			break
		}
	}

	if err != nil {
		state.Overrides = state.Overrides[errElement:]
		diags = resp.State.Set(ctx, &state)
		resp.Diagnostics.Append(diags...)
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *NetworkGroupOverridesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Parse import ID
	var inputPattern = regexp.MustCompile(`^(?:(?P<domain>[^\s,]+),)?(?P<id>[^\s,]+?)$`)
	match := inputPattern.FindStringSubmatch(req.ID)
	if match == nil {
		errMsg := "Failed to parse import parameters.\nPlease provide import string in the following format: <domain>,<id>\n<domain> is optional. If not provided, `Global` is used implicitly and resource's `domain` attribute is not set.\n" + fmt.Sprintf("Got: %q", req.ID)
		resp.Diagnostics.AddError("Import error", errMsg)
		return
	}

	// Set domain, if provided
	if tmpDomain := match[inputPattern.SubexpIndex("domain")]; tmpDomain != "" {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("domain"), tmpDomain)...)
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), match[inputPattern.SubexpIndex("id")])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin createSubresources

// End of section. //template:end createSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin deleteSubresources

// End of section. //template:end deleteSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin updateSubresources

// End of section. //template:end updateSubresources
