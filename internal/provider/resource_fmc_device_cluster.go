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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/netascode/terraform-provider-fmc/internal/provider/helpers"
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
		MarkdownDescription: helpers.NewAttributeDescription("This device manages FTD Device Cluster configuration. This is resource may be re-designed in future releases.").String,

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
			"control_node_device_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Cluster Control Node device ID.").String,
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
				MarkdownDescription: helpers.NewAttributeDescription("List of data nodes where hardware needs to match the control node hardware.").String,
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

	// Create object
	body := plan.toBody(ctx, DeviceCluster{})
	res, err := r.client.Post(plan.getPath(), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}

	taskID := res.Get("metadata.task.id").String()
	tflog.Debug(ctx, fmt.Sprintf("%s: Async task initiated successfully", taskID))

	const atom time.Duration = 5 * time.Second
	// We need device's UUID, but it only shows after the task succeeds. Poll the task.
	for i := time.Duration(0); i < 5*time.Minute; i += atom {
		task, err := r.client.Get("/api/fmc_config/v1/domain/{DOMAIN_UUID}/job/taskstatuses/"+url.QueryEscape(taskID), reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to read object (GET), got error: %s, %s", err, task.String()))
			return
		}
		stat := strings.ToUpper(task.Get("status").String())
		if stat == "FAILED" {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("API task for the new device failed: %s, %s", task.Get("message"), task.Get("description")))
			return
		}
		if stat == "UNKNOWN" {
			//how to check if field exist?
			sub_stat := strings.ToUpper(task.Get("subTasks.0.status").String())
			if sub_stat != "PENDING" && stat != "RUNNING" {
				break
			}
		} else if stat != "PENDING" && stat != "RUNNING" {
			break
		}
		time.Sleep(atom)
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

// Section below is generated&owned by "gen/generator.go". //template:begin update

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

	body := plan.toBody(ctx, state)
	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete

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
	res, err := r.client.Delete(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString()), reqMods...)
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import

func (r *DeviceClusterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin createSubresources

// End of section. //template:end createSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin deleteSubresources

// End of section. //template:end deleteSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin updateSubresources

// End of section. //template:end updateSubresources
