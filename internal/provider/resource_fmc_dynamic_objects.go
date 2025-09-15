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
	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/planmodifiers"
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
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &DynamicObjectsResource{}
	_ resource.ResourceWithImportState = &DynamicObjectsResource{}
)

func NewDynamicObjectsResource() resource.Resource {
	return &DynamicObjectsResource{}
}

type DynamicObjectsResource struct {
	client *fmc.Client
}

func (r *DynamicObjectsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dynamic_objects"
}

func (r *DynamicObjectsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages Dynamic Objects through bulk operations.").AddMinimumVersionHeaderDescription().AddMinimumVersionBulkUpdateDescription().String,

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
				MarkdownDescription: helpers.NewAttributeDescription("Map of Dynamic Objects. The key of the map is the name of the individual Dynamic Object.").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the managed Dynamic Object.").String,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								planmodifiers.UseStateForUnknownKeepNonNullStateString(),
							},
						},
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this value is always 'DynamicObject'.").String,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								planmodifiers.UseStateForUnknownKeepNonNullStateString(),
							},
						},
						"description": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Description of the object.").String,
							Optional:            true,
						},
						"object_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of dynamic object mappings.").AddStringEnumDescription("IP").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("IP"),
							},
						},
						"mappings": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of mappings.").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *DynamicObjectsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *DynamicObjectsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DynamicObjects

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
	state.Items = make(map[string]DynamicObjectsItems, len(plan.Items))
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

func (r *DynamicObjectsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DynamicObjects

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Save objects with mappings saved
	hasMappings := make(map[string]string, len(state.Items))
	for k, v := range state.Items {
		if !v.Mappings.IsNull() {
			hasMappings[k] = v.Id.ValueString()
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("State at the beginning: %v", state))

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

	// Read mappings
	for k, v := range hasMappings {
		tflog.Debug(ctx, fmt.Sprintf("State item entry this particualr: %v", v))

		// Get mappings for the object
		urlPath := state.getPath() + "/" + url.QueryEscape(v) + "/mappings"
		res, err := r.client.Get(urlPath, reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object mappings (GET), got error: %s, %s", err, res.String()))
			return
		}

		// Read mappings from response
		tmp := state.Items[k]
		tmp.Mappings = state.fromBodyMapping(ctx, res)
		state.Items[k] = tmp
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *DynamicObjectsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state DynamicObjects

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
	var toDelete DynamicObjects
	toDelete.Items = make(map[string]DynamicObjectsItems)
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
	var toCreate DynamicObjects
	toCreate.Items = make(map[string]DynamicObjectsItems)
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
	var toUpdate DynamicObjects
	var toUpdateMappings DynamicObjectsMappings
	toUpdate.Items = make(map[string]DynamicObjectsItems)
	toUpdateMappings.Items = make(map[string]DynamicObjectsMappingsItems)

	for _, valueState := range state.Items {

		// Check if the ID from plan exists on list of ID owned by state
		if keyState, ok := planOwnedIDs[valueState.Id.ValueString()]; ok {

			// Check if items in state and plan are equal
			notEqual, diags = helpers.IsConfigUpdatingAt(ctx, req.Plan, req.State, path.Root("items").AtMapKey(keyState).AtName("description"))
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

			// Check if items in state and plan are equal
			notEqual, diags = helpers.IsConfigUpdatingAt(ctx, req.Plan, req.State, path.Root("items").AtMapKey(keyState).AtName("mappings"))
			if diags != nil {
				resp.Diagnostics.Append(diags...)
				diags = resp.State.Set(ctx, &state)
				resp.Diagnostics.Append(diags...)
				return
			}

			// If mappings differ, add object to update list
			if notEqual {
				var tmpObject DynamicObjectsMappingsItems

				newMappings := plan.Items[keyState].Mappings
				oldMappings := state.Items[keyState].Mappings

				tmpObject.Id = plan.Items[keyState].Id
				tmpObject.Add = helpers.DifferenceStringSet(ctx, oldMappings, newMappings)
				tmpObject.Remove = helpers.DifferenceStringSet(ctx, newMappings, oldMappings)

				tflog.Debug(ctx, fmt.Sprintf("Mappings to update: %+v", tmpObject))

				toUpdateMappings.Items[keyState] = tmpObject
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

	// If there are mappings marked for update
	if len(toUpdateMappings.Items) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Number of mappings to update: %d", state.Id.ValueString(), len(toUpdateMappings.Items)))
		state, diags = r.updateMappings(ctx, state, plan, toUpdateMappings, reqMods...)
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

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *DynamicObjectsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DynamicObjects

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
func (r *DynamicObjectsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import looks for string in the following format: <domain_name>,<ref_id>,[<object1_name>,<object2_name>,...]
	// <domain_name> is optional
	// <ref_id> for objects that have `reference` attributes
	// <object1_name>,<object2_name>,... is coma-separated list of object names
	var config DynamicObjects

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
	config.Items = make(map[string]DynamicObjectsItems, len(names))
	for _, v := range names {
		config.Items[v] = DynamicObjectsItems{}
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

// createSubresources takes list of objects, splits them into bulks and creates them
// We want to save the state after each create event, to be able track already created resources
func (r *DynamicObjectsResource) createSubresources(ctx context.Context, state, plan DynamicObjects, reqMods ...func(*fmc.Req)) (DynamicObjects, diag.Diagnostics) {
	var idx = 0
	var bulk DynamicObjects
	var bulkMappings DynamicObjectsMappings
	var mapping DynamicObjectsMappingsItems
	bulk.Items = make(map[string]DynamicObjectsItems, bulkSizeCreate)
	bulkMappings.Items = make(map[string]DynamicObjectsMappingsItems, bulkSizeCreate)

	tflog.Debug(ctx, fmt.Sprintf("%s: Creating bulk of objects", state.Id.ValueString()))

	// iterate over all items
	for k, v := range plan.Items {
		// count loops
		idx++

		// remove mappings, as those are created separately
		tflog.Debug(ctx, fmt.Sprintf("Mappings found: %v", v.Mappings))
		if !v.Mappings.IsNull() {
			tflog.Debug(ctx, fmt.Sprintf("Non empty mappings found %v", v.Mappings))
			mapping.Add = v.Mappings
			bulkMappings.Items[k] = mapping
			v.Mappings = types.Set{}
		}
		tflog.Debug(ctx, fmt.Sprintf("Bulk mappings: %v", bulkMappings))

		// add object to current bulk
		bulk.Items[k] = v

		// If bulk size was reached or all entries have been processed
		if idx%bulkSizeCreate == 0 || idx == len(plan.Items) {

			// Parse body of the request to string
			body := bulk.toBody(ctx, DynamicObjects{})

			// Execute request
			urlPath := bulk.getPath() + "?bulk=true"
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
				// Add ID to mappings
				if _, ok := bulkMappings.Items[k]; ok {
					tmp := bulkMappings.Items[k]
					tmp.Id = v.Id
					bulkMappings.Items[k] = tmp
				}
			}

			// Clear bulk item for next run
			bulk.Items = make(map[string]DynamicObjectsItems, bulkSizeCreate)
		}
	}

	// Create mappings
	state, diags := r.updateMappings(ctx, state, plan, bulkMappings, reqMods...)

	return state, diags
}

// Section below is generated&owned by "gen/generator.go". //template:begin deleteSubresources
// deleteSubresources takes list of objects and deletes them either in bulk, or one-by-one, depending on FMC version
func (r *DynamicObjectsResource) deleteSubresources(ctx context.Context, state, plan DynamicObjects, reqMods ...func(*fmc.Req)) (DynamicObjects, diag.Diagnostics) {
	objectsToRemove := plan.Clone()
	tflog.Debug(ctx, fmt.Sprintf("%s: Bulk deletion mode (Dynamic Objects)", state.Id.ValueString()))

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

	return state, nil
}

// End of section. //template:end deleteSubresources

// updateSubresources take elements one-by-one and updates them, as bulks are not supported
func (r *DynamicObjectsResource) updateSubresources(ctx context.Context, state, plan DynamicObjects, reqMods ...func(*fmc.Req)) (DynamicObjects, diag.Diagnostics) {
	var tmpObject DynamicObjects
	tmpObject.Items = make(map[string]DynamicObjectsItems, 1)

	tflog.Debug(ctx, fmt.Sprintf("%s: Updating bulk of objects", state.Id.ValueString()))

	for k, v := range plan.Items {
		// Remove mappings, as those are handled separately
		v.Mappings = types.Set{}
		tmpObject.Items[k] = v

		body := tmpObject.toBodyNonBulk(ctx, state)
		urlPath := tmpObject.getPath() + "/" + url.QueryEscape(v.Id.ValueString())
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

// updateMappings updates (add & remove) mappings based on provided mappings structure
func (r *DynamicObjectsResource) updateMappings(ctx context.Context, state, plan DynamicObjects, mappings DynamicObjectsMappings, reqMods ...func(*fmc.Req)) (DynamicObjects, diag.Diagnostics) {
	var idx = 0
	var bulk DynamicObjectsMappings
	// each item may be converted to two (for 'add' and 'remove')
	mappingBulkSizeCreate := int(bulkSizeCreate / 2)
	bulk.Items = make(map[string]DynamicObjectsMappingsItems, mappingBulkSizeCreate)

	tflog.Debug(ctx, fmt.Sprintf("%s: Creating bulk of dynamic objects mappings", state.Id.ValueString()))

	// iterate over all items
	for k, v := range mappings.Items {
		// count loops
		idx++

		// add object to current bulk
		bulk.Items[k] = v

		// If bulk size was reached or all entries have been processed
		if idx%mappingBulkSizeCreate == 0 || idx == len(mappings.Items) {

			// Parse body of the request to string
			body := bulk.toBody(ctx, DynamicObjectsMappings{})

			// Execute request
			urlPath := bulk.getPath()
			res, err := r.client.Post(urlPath, body, reqMods...)
			if err != nil {
				return state, diag.Diagnostics{
					diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Failed to create a bulk (POST) id: %s, got error: %s, %s", state.Id.ValueString(), err, res.String())),
				}
			}

			// Read result and save it to the state
			for k := range bulk.Items {
				tmp := state.Items[k]
				tmp.Mappings = plan.Items[k].Mappings
				state.Items[k] = tmp
			}

			// Clear bulk item for next run
			bulk.Items = make(map[string]DynamicObjectsMappingsItems, mappingBulkSizeCreate)
		}
	}

	return state, nil
}
