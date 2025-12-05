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
	"maps"
	"net/url"
	"regexp"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/google/uuid"
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
	_ resource.Resource                = &ASPathsResource{}
	_ resource.ResourceWithImportState = &ASPathsResource{}
)

func NewASPathsResource() resource.Resource {
	return &ASPathsResource{}
}

type ASPathsResource struct {
	client *fmc.Client
}

func (r *ASPathsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_as_paths"
}

func (r *ASPathsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages AS Paths through bulk operations.").AddMinimumVersionHeaderDescription().AddMinimumVersionBulkCreateDescription("999").AddMinimumVersionBulkDeleteDescription("999").AddMinimumVersionBulkUpdateDescription().String,

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
				MarkdownDescription: helpers.NewAttributeDescription("Map of AS Paths. The key of the map is the name of the individual AS Path object.").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the managed AS Path object.").String,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseNonNullStateForUnknown(),
							},
						},
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this value is always 'ASPathList'.").String,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseNonNullStateForUnknown(),
							},
						},
						"overridable": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether the object values can be overridden.").String,
							Optional:            true,
						},
						"entries": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of AS Path entries.").String,
							Required:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"action": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Action to take.").AddStringEnumDescription("PERMIT", "DENY").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("PERMIT", "DENY"),
										},
									},
									"regular_expression": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Regular expression.").String,
										Required:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *ASPathsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *ASPathsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ASPaths

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
	state.Items = make(map[string]ASPathsItems, len(plan.Items))
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

func (r *ASPathsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ASPaths

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

func (r *ASPathsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ASPaths

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
	var toDelete ASPaths
	toDelete.Items = make(map[string]ASPathsItems, len(state.Items))
	planOwnedIDs := make(map[string]string, len(plan.Items))

	// Prepare list of ID that are in plan
	for k, v := range plan.Items {
		if !v.Id.IsUnknown() && v.Id.ValueString() != "" {
			planOwnedIDs[v.Id.ValueString()] = k
		}
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
	var toCreate ASPaths
	toCreate.Items = make(map[string]ASPathsItems, len(plan.Items))
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
	var toUpdate ASPaths
	toUpdate.Items = make(map[string]ASPathsItems, len(plan.Items))

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

func (r *ASPathsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ASPaths

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
func (r *ASPathsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Parse import ID
	var inputPattern = regexp.MustCompile(`^(?:(?P<domain>[^\s,]+),)?\[(?P<names>.*?)\]$`)
	match := inputPattern.FindStringSubmatch(req.ID)
	if match == nil {
		errMsg := "Failed to parse import parameters.\nPlease provide import string in the following format: <domain>,[<item1_name>,<item2_name>,...]\n<domain> is optional. If not provided, `Global` is used implicitly and resource's `domain` attribute is not set.\n" + fmt.Sprintf("Got: %q", req.ID)
		resp.Diagnostics.AddError("Import error", errMsg)
		return
	}

	// Set domain, if provided
	if tmpDomain := match[inputPattern.SubexpIndex("domain")]; tmpDomain != "" {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("domain"), tmpDomain)...)
	}
	// Generate new ID (random, does not relate to FMC in any way)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), uuid.New().String())...)

	// Fill state with names of objects to import
	names := strings.Split(match[inputPattern.SubexpIndex("names")], ",")
	itemsMap := make(map[string]ASPathsItems, len(names))
	for _, v := range names {
		itemsMap[v] = ASPathsItems{}
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("items"), itemsMap)...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin createSubresources
// createSubresources takes list of objects, splits them into bulks and creates them
// We want to save the state after each create event, to be able track already created resources
func (r *ASPathsResource) createSubresources(ctx context.Context, state, plan ASPaths, reqMods ...func(*fmc.Req)) (ASPaths, diag.Diagnostics) {
	// Check if FMC version supports bulk creates
	if r.client.FMCVersionParsed.LessThan(minFMCVersionBulkCreateASPaths) {
		tflog.Debug(ctx, fmt.Sprintf("%s: One-by-one creation mode (AS Paths)", state.Id.ValueString()))
		var tmpObject ASPaths
		tmpObject.Items = make(map[string]ASPathsItems, 1)
		for k, v := range plan.Items {
			tmpObject.Items[k] = v

			body := tmpObject.toBodyNonBulk(ctx, state)
			body = tmpObject.adjustBody(ctx, body)
			res, err := r.client.Post(state.getPath(), body, reqMods...)
			if err != nil {
				return state, diag.Diagnostics{
					diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("%s: Failed to create object (POST) id %s, got error: %s, %s", state.Id.ValueString(), v.Id.ValueString(), err, res.String())),
				}
			}

			// fromBodyUnknowns expect result to be listed under "items" key
			body, _ = sjson.SetRaw("{}", "items.-1", res.String())
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
		var bulk ASPaths
		bulk.Items = make(map[string]ASPathsItems, bulkSizeCreate)

		tflog.Debug(ctx, fmt.Sprintf("%s: Bulk creation mode (AS Paths)", state.Id.ValueString()))

		// iterate over all items
		for k, v := range plan.Items {
			// count loops
			idx++

			// add object to current bulk
			bulk.Items[k] = v

			// If bulk size was reached or all entries have been processed
			if idx%bulkSizeCreate == 0 || idx == len(plan.Items) {

				// Parse body of the request to string
				body := bulk.toBody(ctx, ASPaths{})
				body = bulk.adjustBodyBulk(ctx, body)

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
				maps.Copy(state.Items, bulk.Items)

				// Clear bulk item for next run
				bulk.Items = make(map[string]ASPathsItems, bulkSizeCreate)
			}
		}
	}

	return state, nil
}

// End of section. //template:end createSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin deleteSubresources
// deleteSubresources takes list of objects and deletes them either in bulk, or one-by-one, depending on FMC version
func (r *ASPathsResource) deleteSubresources(ctx context.Context, state, plan ASPaths, reqMods ...func(*fmc.Req)) (ASPaths, diag.Diagnostics) {
	objectsToRemove := plan.Clone()

	// Check if FMC version supports bulk deletes
	if r.client.FMCVersionParsed.LessThan(minFMCVersionBulkDeleteASPaths) {
		tflog.Debug(ctx, fmt.Sprintf("%s: One-by-one deletion mode (AS Paths)", state.Id.ValueString()))
		for k, v := range objectsToRemove.Items {
			// Check if the object was not already deleted
			if v.Id.IsNull() {
				delete(state.Items, k)
				continue
			}

			urlPath := state.getPath() + "/" + url.QueryEscape(v.Id.ValueString())
			res, err := r.client.Delete(urlPath, reqMods...)
			if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
				return state, diag.Diagnostics{
					diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("%s: Failed to delete object (DELETE) id %s, got error: %s, %s", state.Id.ValueString(), v.Id.ValueString(), err, res.String())),
				}
			}

			// Remove deleted item from state
			delete(state.Items, k)
		}
	} else {
		tflog.Debug(ctx, fmt.Sprintf("%s: Bulk deletion mode (AS Paths)", state.Id.ValueString()))

		var idx = 0

		estimatedIDLength := 37 // UUID length + comma
		estimatedCapacity := min(len(objectsToRemove.Items)*estimatedIDLength, maxUrlParamLength)
		var idsToRemove strings.Builder
		idsToRemove.Grow(estimatedCapacity)

		for k, v := range objectsToRemove.Items {
			// Counter
			idx++

			// Check if the object was not already deleted
			if v.Id.IsNull() {
				delete(state.Items, k)
				continue
			}

			// Create list of IDs of items to delete
			idsToRemove.WriteString(v.Id.ValueString())
			idsToRemove.WriteString(",")

			// If bulk size was reached or all entries have been processed
			if idsToRemove.Len() >= maxUrlParamLength || idx == len(objectsToRemove.Items) {
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
	}

	return state, nil
}

// End of section. //template:end deleteSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin updateSubresources

// updateSubresources take elements one-by-one and updates them, as bulks are not supported
func (r *ASPathsResource) updateSubresources(ctx context.Context, state, plan ASPaths, reqMods ...func(*fmc.Req)) (ASPaths, diag.Diagnostics) {
	var tmpObject ASPaths
	tmpObject.Items = make(map[string]ASPathsItems, 1)

	tflog.Debug(ctx, fmt.Sprintf("%s: One-by-one update mode (AS Paths)", state.Id.ValueString()))

	for k, v := range plan.Items {
		tmpObject.Items[k] = v

		body := tmpObject.toBodyNonBulk(ctx, state)
		body = tmpObject.adjustBody(ctx, body)
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
