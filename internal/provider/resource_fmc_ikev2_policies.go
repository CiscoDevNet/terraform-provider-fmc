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
	"github.com/google/uuid"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
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

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &IKEv2PoliciesResource{}
	_ resource.ResourceWithImportState = &IKEv2PoliciesResource{}
)

func NewIKEv2PoliciesResource() resource.Resource {
	return &IKEv2PoliciesResource{}
}

type IKEv2PoliciesResource struct {
	client *fmc.Client
}

func (r *IKEv2PoliciesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ikev2_policies"
}

func (r *IKEv2PoliciesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages IKEv2 Policies through bulk operations.").AddMinimumVersionHeaderDescription().AddMinimumVersionBulkCreateDescription("999").AddMinimumVersionBulkDeleteDescription("999").AddMinimumVersionBulkUpdateDescription().String,

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
			"items": schema.MapNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Map of IKEv2 Policies. The key of the map is the name of the individual IKEv2 Policy.").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the IKEv2 Policy object.").String,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"description": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Description of the object.").String,
							Optional:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this value is always 'IKEv2Policy'.").String,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"priority": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Priority of the IKEv1 Policy.").AddIntegerRangeDescription(1, 65535).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65535),
							},
						},
						"lifetime": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Lifetime in seconds.").AddIntegerRangeDescription(120, 2147483647).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(120, 2147483647),
							},
						},
						"integrity_algorithms": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IKEv2 Integrity algorithms.").AddStringEnumDescription("SHA", "SHA-256", "SHA-384", "SHA-512", "MD5", "NULL").String,
							ElementType:         types.StringType,
							Required:            true,
							Validators: []validator.Set{
								setvalidator.ValueStringsAre(
									stringvalidator.OneOf("SHA", "SHA-256", "SHA-384", "SHA-512", "MD5", "NULL"),
								),
							},
						},
						"encryption_algorithms": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IKEv2 Encryption algorithms.").AddStringEnumDescription("DES", "3DES", "AES", "AES-192", "AES-256", "AES-GCM", "AES-GCM-192", "AES-GCM-256", "NULL").String,
							ElementType:         types.StringType,
							Required:            true,
							Validators: []validator.Set{
								setvalidator.ValueStringsAre(
									stringvalidator.OneOf("DES", "3DES", "AES", "AES-192", "AES-256", "AES-GCM", "AES-GCM-192", "AES-GCM-256", "NULL"),
								),
							},
						},
						"prf_algorithms": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IKEv2 Pseudo-Random Function (PRF) algorithms.").AddStringEnumDescription("SHA", "SHA-256", "SHA-384", "SHA-512", "MD5").String,
							ElementType:         types.StringType,
							Required:            true,
							Validators: []validator.Set{
								setvalidator.ValueStringsAre(
									stringvalidator.OneOf("SHA", "SHA-256", "SHA-384", "SHA-512", "MD5"),
								),
							},
						},
						"dh_groups": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IKEv2 Diffie-Hellman groups.").AddStringEnumDescription("1", "2", "5", "14", "15", "16", "19", "20", "21", "24", "31").String,
							ElementType:         types.StringType,
							Required:            true,
							Validators: []validator.Set{
								setvalidator.ValueStringsAre(
									stringvalidator.OneOf("1", "2", "5", "14", "15", "16", "19", "20", "21", "24", "31"),
								),
							},
						},
					},
				},
			},
		},
	}
}

func (r *IKEv2PoliciesResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *IKEv2PoliciesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IKEv2Policies

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

	//// Prepare state to track creation process. Create request is split to multiple requests, where just subset of them may be successful
	// Copy fields, as those may contain domain information or other references
	state := plan
	// Create random ID to track bulk resource. This does not relate to FMC in any way
	state.Id = types.StringValue(uuid.New().String())
	// Erase all Items, those will be filled in after creation
	state.Items = make(map[string]IKEv2PoliciesItems, len(plan.Items))
	// Creation process is put in a separate function, as that same proces will be needed with `Update`
	plan, diags = r.createSubresources(ctx, state, plan, reqMods...)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		// Save state for whatever was already created
		diags = resp.State.Set(ctx, &plan)
		tflog.Debug(ctx, fmt.Sprintf("%s: Create failed, some items might have been created", plan.Id.ValueString()))
		resp.Diagnostics.Append(diags...)
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *IKEv2PoliciesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IKEv2Policies

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

	// Get all objects from FMC
	urlPath := state.getPath() + "?expanded=true"
	res, err := r.client.Get(urlPath, reqMods...)
	if err != nil {
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

// Section below is generated&owned by "gen/generator.go". //template:begin update

func (r *IKEv2PoliciesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IKEv2Policies

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

	// DELETE
	// Delete objects (that are present in state, but missing in plan)
	var toDelete IKEv2Policies
	toDelete.Items = make(map[string]IKEv2PoliciesItems)
	planOwnedIDs := make(map[string]string, len(plan.Items))

	// Prepare list of ID that are in plan
	for k, v := range plan.Items {
		planOwnedIDs[v.Id.ValueString()] = k
	}

	// Check if ID from state list is in plan as well. If not, mark it for delete
	for k, v := range state.Items {
		if _, ok := planOwnedIDs[v.Id.ValueString()]; !ok {
			toDelete.Items[k] = v
		}
	}

	// If there are objects marked to be deleted
	if len(toDelete.Items) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to delete: %d", state.Id.ValueString(), len(toDelete.Items)))
		state, diags = r.deleteSubresources(ctx, state, toDelete, reqMods...)
		if diags != nil {
			resp.Diagnostics.Append(diags...)
			diags = resp.State.Set(ctx, &state)
			resp.Diagnostics.Append(diags...)
			return
		}
	}

	// CREATE
	// Create new objects (objects that have missing IDs in plan)
	var toCreate IKEv2Policies
	toCreate.Items = make(map[string]IKEv2PoliciesItems)
	// Scan plan for items with no ID
	for k, v := range plan.Items {
		if v.Id.IsUnknown() || v.Id.IsNull() {
			toCreate.Items[k] = v
		}
	}

	// If there are objects marked for create
	if len(toCreate.Items) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to create: %d", state.Id.ValueString(), len(toCreate.Items)))
		state, diags = r.createSubresources(ctx, state, toCreate, reqMods...)
		if diags != nil {
			resp.Diagnostics.Append(diags...)
			diags = resp.State.Set(ctx, &state)
			resp.Diagnostics.Append(diags...)
			return
		}
	}

	// UPDATE
	// Update objects (objects that have different definition in plan and state)
	var notEqual bool
	var toUpdate IKEv2Policies
	toUpdate.Items = make(map[string]IKEv2PoliciesItems)

	for _, valueState := range state.Items {

		// Check if the ID from plan exists on list of ID owned by state
		if keyState, ok := planOwnedIDs[valueState.Id.ValueString()]; ok {

			// Check if items in state and plan are qual
			notEqual, diags = helpers.IsConfigUpdatingAt(ctx, req.Plan, req.State, path.Root("items").AtMapKey(keyState))
			if diags != nil {
				resp.Diagnostics.Append(diags...)
				diags = resp.State.Set(ctx, &state)
				resp.Diagnostics.Append(diags...)
				return
			}

			// If definitions differ, add object to update list
			if notEqual {
				toUpdate.Items[keyState] = plan.Items[keyState]
			}
		}
	}

	// If there are objects marked for update
	if len(toUpdate.Items) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to update: %d", state.Id.ValueString(), len(toUpdate.Items)))
		state, diags = r.updateSubresources(ctx, state, toUpdate, reqMods...)
		if diags != nil {
			resp.Diagnostics.Append(diags...)
			diags = resp.State.Set(ctx, &state)
			resp.Diagnostics.Append(diags...)
			return
		}
	}
	plan = state

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *IKEv2PoliciesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IKEv2Policies

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

	// Execute delete
	state, diags = r.deleteSubresources(ctx, state, state, reqMods...)
	resp.Diagnostics.Append(diags...)

	// Check if every element was removed
	if len(state.Items) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Not all elements have been removed", state.Id.ValueString()))
		diags = resp.State.Set(ctx, &state)
		resp.Diagnostics.Append(diags...)
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *IKEv2PoliciesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import looks for string in the following format: <domain_name>,<ref_id>,[<object1_name>,<object2_name>,...]
	// <domain_name> is optional
	// <ref_id> for objects that have `reference` attributes
	// <object1_name>,<object2_name>,... is coma-separated list of object names
	var config IKEv2Policies

	// Compile pattern for import command parsing
	var inputPattern = regexp.MustCompile(`^(?:(?P<domain>[^\s,]+),)?\[(?P<names>.*?)\]$`)

	// Parse parameter
	match := inputPattern.FindStringSubmatch(req.ID)

	// Check if regex matched
	if match == nil {
		resp.Diagnostics.AddError("Import error", "Failed to parse import parameters. Please provide import string in the following format: <domain_name>,[<object1_name>,<object2_name>,...]")
		return
	}

	// Extract values
	if tmpDomain := match[inputPattern.SubexpIndex("domain")]; tmpDomain != "" {
		config.Domain = types.StringValue(tmpDomain)
	}
	names := strings.Split(match[inputPattern.SubexpIndex("names")], ",")

	// Fill state with names of objects to import
	config.Items = make(map[string]IKEv2PoliciesItems, len(names))
	for _, v := range names {
		config.Items[v] = IKEv2PoliciesItems{}
	}

	// Generate new ID
	config.Id = types.StringValue(uuid.New().String())

	// Set filled in structure
	diags := resp.State.Set(ctx, config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set import flag
	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin createSubresources
// createSubresources takes list of objects, splits them into bulks and creates them
// We want to save the state after each create event, to be able track already created resources
func (r *IKEv2PoliciesResource) createSubresources(ctx context.Context, state, plan IKEv2Policies, reqMods ...func(*fmc.Req)) (IKEv2Policies, diag.Diagnostics) {
	// Get FMC version from the clinet
	fmcVersion, _ := version.NewVersion(strings.Split(r.client.FMCVersion, " ")[0])

	// Check if FMC version supports bulk creates
	if fmcVersion.LessThan(minFMCVersionBulkCreateIKEv2Policies) {
		tflog.Debug(ctx, fmt.Sprintf("%s: One-by-one creation mode (IKEv2 Policies)", state.Id.ValueString()))
		var tmpObject IKEv2Policies
		tmpObject.Items = make(map[string]IKEv2PoliciesItems, 1)
		for k, v := range plan.Items {
			tmpObject.Items[k] = v

			body := tmpObject.toBodyNonBulk(ctx, state)
			res, err := r.client.Post(state.getPath(), body, reqMods...)
			if err != nil {
				return state, diag.Diagnostics{
					diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("%s: Failed to create object (POST) id %s, got error: %s, %s", state.Id.ValueString(), v.Id.ValueString(), err, res.String())),
				}
			}

			// fromBodyUnknowns expect result to be listed under "items" key
			body, _ = sjson.SetRaw("{items:[]}", "items.-1", res.String())
			res = gjson.Parse(body)

			// Read computed values
			tmpObject.fromBodyUnknowns(ctx, res)

			// Save object to plan
			state.Items[k] = tmpObject.Items[k]

			// Clear tmpObject.Items
			delete(tmpObject.Items, k)

		}
	} else {
		var idx = 0
		var bulk IKEv2Policies
		bulk.Items = make(map[string]IKEv2PoliciesItems, bulkSizeCreate)

		tflog.Debug(ctx, fmt.Sprintf("%s: Bulk creation mode (IKEv2 Policies)", state.Id.ValueString()))

		// iterate over all items
		for k, v := range plan.Items {
			// count loops
			idx++

			// add object to current bulk
			bulk.Items[k] = v

			// If bulk size was reached or all entries have been processed
			if idx%bulkSizeCreate == 0 || idx == len(plan.Items) {

				// Parse body of the request to string
				body := bulk.toBody(ctx, IKEv2Policies{})

				// Execute request
				urlPath := state.getPath() + "?bulk=true"
				res, err := r.client.Post(urlPath, body, reqMods...)
				if err != nil {
					return state, diag.Diagnostics{
						diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Failed to create a bulk (POST) id: %s, got error: %s, %s", state.Id.ValueString(), err, res.String())),
					}
				}

				// Read result and save it to the state
				bulk.fromBodyUnknowns(ctx, res)
				for k, v := range bulk.Items {
					state.Items[k] = v
				}

				// Clear bulk item for next run
				bulk.Items = make(map[string]IKEv2PoliciesItems, bulkSizeCreate)
			}
		}
	}

	return state, nil
}

// End of section. //template:end createSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin deleteSubresources
// deleteSubresources takes list of objects and deletes them either in bulk, or one-by-one, depending on FMC version
func (r *IKEv2PoliciesResource) deleteSubresources(ctx context.Context, state, plan IKEv2Policies, reqMods ...func(*fmc.Req)) (IKEv2Policies, diag.Diagnostics) {
	objectsToRemove := plan.Clone()

	// Get FMC version from the clinet
	fmcVersion, _ := version.NewVersion(strings.Split(r.client.FMCVersion, " ")[0])

	// Check if FMC version supports bulk deletes
	if fmcVersion.LessThan(minFMCVersionBulkDeleteIKEv2Policies) {
		tflog.Debug(ctx, fmt.Sprintf("%s: One-by-one deletion mode (IKEv2 Policies)", state.Id.ValueString()))
		for k, v := range objectsToRemove.Items {
			// Check if the object was not already deleted
			if v.Id.IsNull() {
				delete(state.Items, k)
				continue
			}

			urlPath := state.getPath() + "/" + url.QueryEscape(v.Id.ValueString())
			res, err := r.client.Delete(urlPath, reqMods...)
			if err != nil {
				return state, diag.Diagnostics{
					diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("%s: Failed to delete object (DELETE) id %s, got error: %s, %s", state.Id.ValueString(), v.Id.ValueString(), err, res.String())),
				}
			}

			// Remove deleted item from state
			delete(state.Items, k)
		}
	} else {
		tflog.Debug(ctx, fmt.Sprintf("%s: Bulk deletion mode (IKEv2 Policies)", state.Id.ValueString()))

		var idx = 0
		var idsToRemove strings.Builder
		var alreadyDeleted []string

		for k, v := range objectsToRemove.Items {
			// Counter
			idx++

			// Check if the object was not already deleted
			if v.Id.IsNull() {
				alreadyDeleted = append(alreadyDeleted, k)
				continue
			}

			// Create list of IDs of items to delete
			idsToRemove.WriteString(v.Id.ValueString() + ",")

			// If bulk size was reached or all entries have been processed
			if idx%bulkSizeDelete == 0 || idx == len(objectsToRemove.Items) {
				urlPath := state.getPath() + "?bulk=true&filter=ids:" + url.QueryEscape(idsToRemove.String())
				res, err := r.client.Delete(urlPath, reqMods...)
				if err != nil {
					return state, diag.Diagnostics{
						diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("%s: Failed to delete subobject(s) (DELETE), got error: %s, %s", state.Id.ValueString(), err, res.String())),
					}
				}

				// Read result and remove deleted items from state
				deletedItems := res.Get("items.#.name").Array()
				for _, name := range deletedItems {
					delete(state.Items, name.String())
				}

				// Reset ID string
				idsToRemove.Reset()
			}
		}

		for _, v := range alreadyDeleted {
			delete(state.Items, v)
		}
	}

	return state, nil
}

// End of section. //template:end deleteSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin updateSubresources

// updateSubresources take elements one-by-one and updates them, as bulks are not supported
func (r *IKEv2PoliciesResource) updateSubresources(ctx context.Context, state, plan IKEv2Policies, reqMods ...func(*fmc.Req)) (IKEv2Policies, diag.Diagnostics) {
	var tmpObject IKEv2Policies
	tmpObject.Items = make(map[string]IKEv2PoliciesItems, 1)

	tflog.Debug(ctx, fmt.Sprintf("%s: One-by-one update mode (IKEv2 Policies)", state.Id.ValueString()))

	for k, v := range plan.Items {
		tmpObject.Items[k] = v

		body := tmpObject.toBodyNonBulk(ctx, state)
		urlPath := state.getPath() + "/" + url.QueryEscape(v.Id.ValueString())
		res, err := r.client.Put(urlPath, body, reqMods...)
		if err != nil {
			return state, diag.Diagnostics{
				diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Failed to update object (PUT) id %s, got error: %s, %s", state.Id.ValueString(), err, res.String())),
			}
		}

		// Update state
		state.Items[k] = v

		// Clear tmpObject.Items
		delete(tmpObject.Items, k)
	}

	return state, nil
}

// End of section. //template:end updateSubresources
