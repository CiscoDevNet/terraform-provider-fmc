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
	"slices"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &DeviceClusterResource{}
	_ resource.ResourceWithImportState = &DeviceClusterResource{}
)

func NewDeviceClusterResource() resource.Resource {
	return &DeviceClusterResource{}
}

type DeviceClusterResource struct {
	client *fmc.Client
}

func (r *DeviceClusterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_cluster"
}

func (r *DeviceClusterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This device manages FTD Device Cluster configuration.\n Configuration of the Cluster is replicated from the Cluster Control Node. Nevertheless, please make sure that the configuration of the control and all the data nodes is consistent.\n The following actions are not supported:\n - Renaming the cluster,\n - Disabling/Enabling cluster node,\n - Changing cluster control node,\n").String,

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
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the FTD Cluster.").String,
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the resource; This is always `DeviceCluster`.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"cluster_key": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Secret key for the cluster, between 1 nd 63 characters.").String,
				Required:            true,
			},
			"control_node_vni_prefix": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Cluster Control VXLAN Network Identifier (VNI) Network").String,
				Optional:            true,
			},
			"control_node_ccl_prefix": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Cluster Control Link Network / Virtual Tunnel Endpoint (VTEP) Network").String,
				Required:            true,
			},
			"control_node_interface_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Cluster control link interface ID.").String,
				Required:            true,
			},
			"control_node_interface_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Cluster control link interface Name.").String,
				Required:            true,
			},
			"control_node_interface_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Cluster control link interface Type.").String,
				Required:            true,
			},
			"control_node_device_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Cluster Control Node device ID.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"control_node_ccl_ipv4_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Cluster control link IPv4 address / VTEP IPv4 address.").String,
				Required:            true,
			},
			"control_node_priority": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Priority of cluster controle node.").AddIntegerRangeDescription(1, 255).String,
				Required:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
			},
			"data_devices": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of cluster data nodes.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"data_node_device_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Cluster Data Node device ID.").String,
							Required:            true,
						},
						"data_node_ccl_ipv4_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Cluster Data Node link IPv4 address / VTEP IPv4 address.").String,
							Required:            true,
						},
						"data_node_priority": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Priority of cluster data node.").AddIntegerRangeDescription(1, 255).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 255),
							},
						},
					},
				},
			},
		},
	}
}

func (r *DeviceClusterResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

func (r *DeviceClusterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DeviceCluster

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

	// Deploy configuration to devices in Cluster (pre-requisite for cluster creation)
	var clusterDeviceIds = []string{plan.ControlNodeDeviceId.ValueString()}
	for _, v := range plan.DataDevices {
		clusterDeviceIds = append(clusterDeviceIds, v.DataNodeDeviceId.ValueString())
	}

	deploy := DeviceDeploy{
		Id:            types.StringValue(plan.Id.ValueString()),
		IgnoreWarning: types.BoolValue(true),
		DeviceIdList:  helpers.GetStringListFromStringSlice(clusterDeviceIds),
	}
	diags = FMCDeviceDeploy(ctx, r.client, deploy, reqMods)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Create object
	body := plan.toBody(ctx, DeviceCluster{})
	res, err := r.client.Post(plan.getPath(), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}

	// Wait for cluster to be created
	taskID := res.Get("metadata.task.id").String()
	tflog.Debug(ctx, fmt.Sprintf("%s: Async task initiated successfully (id: %s)", plan.Id.ValueString(), taskID))

	diags = FMCWaitForJobToFinish(ctx, r.client, taskID, reqMods)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	check, err := r.client.Get(plan.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to read object (GET), got error: %s, %s", err, check))
		return
	}
	name := "items.#(name==" + url.QueryEscape(plan.Name.ValueString()) + ").id"
	id := check.Get(name).String()
	plan.Id = types.StringValue(id)

	if plan.Id.ValueString() == "" {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("No Cluster named %q: %s", plan.Name.ValueString(), check))
		return
	}
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))
	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *DeviceClusterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DeviceCluster

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

func (r *DeviceClusterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state DeviceCluster

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

	// Check if name has changed
	if plan.Name.ValueString() != state.Name.ValueString() {
		tflog.Debug(ctx, fmt.Sprintf("%s: Name has changed", plan.Id.ValueString()))
		body := plan.toBody(ctx, DeviceCluster{})
		body, _ = sjson.Set(body, "action", "UPDATE_CLUSTER_NAME")
		res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body, reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
			return
		}
	}

	// Check if bootsrap configuration has changed
	if state.hasBootstrapChanged(ctx, plan) {
		tflog.Debug(ctx, fmt.Sprintf("%s: Bootstrap configuration has changed", plan.Id.ValueString()))
		body := plan.toBody(ctx, DeviceCluster{})
		body, _ = sjson.Set(body, "action", "UPDATE_BOOTSTRAP")
		res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body, reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
			return
		}
		// Wait for cluster to be updated
		if res.Get("metadata.task.id").Exists() {
			taskID := res.Get("metadata.task.id").String()
			tflog.Debug(ctx, fmt.Sprintf("%s: Async task initiated successfully (id: %s)", plan.Id.ValueString(), taskID))

			diags = FMCWaitForJobToFinish(ctx, r.client, taskID, reqMods)
			if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
				return
			}
		}
	}

	// Get list of state data devices
	stateDevices := make([]string, len(state.DataDevices))
	for i, v := range state.DataDevices {
		stateDevices[i] = v.DataNodeDeviceId.ValueString()
	}

	// Get list of plan data devices
	planDevices := make([]string, len(plan.DataDevices))
	for i, v := range plan.DataDevices {
		planDevices[i] = v.DataNodeDeviceId.ValueString()
	}

	// Check if any device needs to be removed from cluster
	toBeRemoved := plan
	toBeRemoved.DataDevices = []DeviceClusterDataDevices{}
	for _, v := range state.DataDevices {
		if !slices.Contains(planDevices, v.DataNodeDeviceId.ValueString()) {
			toBeRemoved.DataDevices = append(toBeRemoved.DataDevices, v)
		}
	}

	if len(toBeRemoved.DataDevices) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Data devices to be removed from cluster: %v", plan.Id.ValueString(), toBeRemoved.DataDevices))
		body := toBeRemoved.toBody(ctx, DeviceCluster{})
		body, _ = sjson.Set(body, "action", "REMOVE_NODES")
		res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body, reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
			return
		}

		// Wait for cluster to be updated
		taskID := res.Get("metadata.task.id").String()
		tflog.Debug(ctx, fmt.Sprintf("%s: Async task initiated successfully (id: %s)", plan.Id.ValueString(), taskID))

		diags = FMCWaitForJobToFinish(ctx, r.client, taskID, reqMods)
		if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
			return
		}

	}

	// Check if any device needs to be added to cluster
	toBeAdded := plan
	toBeAdded.DataDevices = []DeviceClusterDataDevices{}
	for _, v := range plan.DataDevices {
		if !slices.Contains(stateDevices, v.DataNodeDeviceId.ValueString()) {
			toBeAdded.DataDevices = append(toBeAdded.DataDevices, v)
		}
	}

	if len(toBeAdded.DataDevices) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Data devices to be added to cluster: %v", plan.Id.ValueString(), toBeAdded.DataDevices))
		body := toBeAdded.toBody(ctx, DeviceCluster{})
		body, _ = sjson.Set(body, "action", "ADD_NODES")
		res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body, reqMods...)
		if err != nil {
			// Error saying that device is already in the cluster
			// This will trigger even if a single device out of multiple is already in the cluster
			// Second terraform apply will fix the problem (1st run updates state, 2nd add missing ones)
			if !strings.Contains(res.String(), "is a duplicate device id") {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
				return
			}
		}

		// Wait for cluster to be updated
		if res.Get("metadata.task.id").Exists() {
			taskID := res.Get("metadata.task.id").String()
			tflog.Debug(ctx, fmt.Sprintf("%s: Async task initiated successfully (id: %s)", plan.Id.ValueString(), taskID))

			diags = FMCWaitForJobToFinish(ctx, r.client, taskID, reqMods)
			if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
				return
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *DeviceClusterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DeviceCluster

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
	// Delete should break the cluster, instead of deleting it
	body := state.toBodyPutDelete(ctx, DeviceCluster{})
	res, err := r.client.Put(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString()), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to remove object configuration (PUT), got error: %s, %s", err, res.String()))
		return
	}

	taskID := res.Get("metadata.task.id").String()
	tflog.Debug(ctx, fmt.Sprintf("%s: Async task initiated successfully (id: %s)", state.Id.ValueString(), taskID))

	diags = FMCWaitForJobToFinish(ctx, r.client, taskID, reqMods)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *DeviceClusterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var config DeviceCluster

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
		config.Domain = types.StringValue(tmpDomain)
	}
	config.Id = types.StringValue(match[inputPattern.SubexpIndex("id")])

	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
