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
	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/planmodifiers"
	"github.com/google/uuid"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
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
	_ resource.Resource                = &SLAMonitorsResource{}
	_ resource.ResourceWithImportState = &SLAMonitorsResource{}
)

func NewSLAMonitorsResource() resource.Resource {
	return &SLAMonitorsResource{}
}

type SLAMonitorsResource struct {
	client *fmc.Client
}

func (r *SLAMonitorsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sla_monitors"
}

func (r *SLAMonitorsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages SLA Monitors through bulk operations.").AddMinimumVersionHeaderDescription().AddMinimumVersionBulkDeleteDescription("999").AddMinimumVersionBulkUpdateDescription().String,

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
				MarkdownDescription: helpers.NewAttributeDescription("Map of SLA Monitors. The key of the map is the name of the individual SLA monitor.").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the managed SLA Monitor.").String,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								planmodifiers.UseStateForUnknownKeepNonNullStateString(),
							},
						},
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this value is always 'SLAMonitor'.").String,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								planmodifiers.UseStateForUnknownKeepNonNullStateString(),
							},
						},
						"description": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Description of the object.").String,
							Optional:            true,
						},
						"sla_monitor_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID number of the SLA operation.").AddIntegerRangeDescription(1, 2147483647).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 2147483647),
							},
						},
						"timeout": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Amount of time (in milliseconds) that the SLA operation waits for a response to the ICMP echo requests.").AddIntegerRangeDescription(0, 604800000).AddDefaultValueDescription("5000").String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 604800000),
							},
							Default: int64default.StaticInt64(5000),
						},
						"frequency": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Frequency (in seconds) of ICMP echo request transmissions.").AddIntegerRangeDescription(1, 604800).AddDefaultValueDescription("60").String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 604800),
							},
							Default: int64default.StaticInt64(60),
						},
						"threshold": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Amount of time (in milliseconds) that must pass after an ICMP echo request before a rising threshold is declared.").AddIntegerRangeDescription(0, 2147483647).AddDefaultValueDescription("5000").String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 2147483647),
							},
							Default: int64default.StaticInt64(5000),
						},
						"data_size": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Size (in bytes) of the ICMP request packet payload.").AddIntegerRangeDescription(0, 16384).AddDefaultValueDescription("28").String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 16384),
							},
							Default: int64default.StaticInt64(28),
						},
						"tos": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of Service (ToS) defined in the IP header of the ICMP request packet.").AddIntegerRangeDescription(0, 255).AddDefaultValueDescription("0").String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 255),
							},
							Default: int64default.StaticInt64(0),
						},
						"number_of_packets": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Number of packets that are sent.").AddIntegerRangeDescription(1, 100).AddDefaultValueDescription("1").String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 100),
							},
							Default: int64default.StaticInt64(1),
						},
						"monitor_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP address to monitor.").String,
							Required:            true,
						},
						"selected_interfaces": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Security zones or interface groups that contain interfaces through which the device communicates.").String,
							Required:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the security zone or interface group object.").String,
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

func (r *SLAMonitorsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *SLAMonitorsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SLAMonitors

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
	state.Items = make(map[string]SLAMonitorsItems, len(plan.Items))
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

func (r *SLAMonitorsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SLAMonitors

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

func (r *SLAMonitorsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SLAMonitors

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
	var toDelete SLAMonitors
	toDelete.Items = make(map[string]SLAMonitorsItems, len(state.Items))
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
	var toCreate SLAMonitors
	toCreate.Items = make(map[string]SLAMonitorsItems, len(plan.Items))
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
	var toUpdate SLAMonitors
	toUpdate.Items = make(map[string]SLAMonitorsItems, len(plan.Items))

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

func (r *SLAMonitorsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SLAMonitors

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
func (r *SLAMonitorsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import looks for string in the following format: <domain_name>,<ref_id>,[<object1_name>,<object2_name>,...]
	// <domain_name> is optional
	// <ref_id> for objects that have `reference` attributes
	// <object1_name>,<object2_name>,... is coma-separated list of object names
	var config SLAMonitors

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
	config.Items = make(map[string]SLAMonitorsItems, len(names))
	for _, v := range names {
		config.Items[v] = SLAMonitorsItems{}
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
func (r *SLAMonitorsResource) createSubresources(ctx context.Context, state, plan SLAMonitors, reqMods ...func(*fmc.Req)) (SLAMonitors, diag.Diagnostics) {
	var idx = 0
	var bulk SLAMonitors
	bulk.Items = make(map[string]SLAMonitorsItems, bulkSizeCreate)

	tflog.Debug(ctx, fmt.Sprintf("%s: Bulk creation mode (SLA Monitors)", state.Id.ValueString()))

	// iterate over all items
	for k, v := range plan.Items {
		// count loops
		idx++

		// add object to current bulk
		bulk.Items[k] = v

		// If bulk size was reached or all entries have been processed
		if idx%bulkSizeCreate == 0 || idx == len(plan.Items) {

			// Parse body of the request to string
			body := bulk.toBody(ctx, SLAMonitors{})

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
			bulk.Items = make(map[string]SLAMonitorsItems, bulkSizeCreate)
		}
	}

	return state, nil
}

// End of section. //template:end createSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin deleteSubresources
// deleteSubresources takes list of objects and deletes them either in bulk, or one-by-one, depending on FMC version
func (r *SLAMonitorsResource) deleteSubresources(ctx context.Context, state, plan SLAMonitors, reqMods ...func(*fmc.Req)) (SLAMonitors, diag.Diagnostics) {
	objectsToRemove := plan.Clone()

	// Get FMC version from the clinet
	fmcVersion, _ := version.NewVersion(strings.Split(r.client.FMCVersion, " ")[0])

	// Check if FMC version supports bulk deletes
	if fmcVersion.LessThan(minFMCVersionBulkDeleteSLAMonitors) {
		tflog.Debug(ctx, fmt.Sprintf("%s: One-by-one deletion mode (SLA Monitors)", state.Id.ValueString()))
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
		tflog.Debug(ctx, fmt.Sprintf("%s: Bulk deletion mode (SLA Monitors)", state.Id.ValueString()))

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
func (r *SLAMonitorsResource) updateSubresources(ctx context.Context, state, plan SLAMonitors, reqMods ...func(*fmc.Req)) (SLAMonitors, diag.Diagnostics) {
	var tmpObject SLAMonitors
	tmpObject.Items = make(map[string]SLAMonitorsItems, 1)

	tflog.Debug(ctx, fmt.Sprintf("%s: One-by-one update mode (SLA Monitors)", state.Id.ValueString()))

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
