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

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
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
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &DeviceClusterHealthMonitorResource{}
	_ resource.ResourceWithImportState = &DeviceClusterHealthMonitorResource{}
)

func NewDeviceClusterHealthMonitorResource() resource.Resource {
	return &DeviceClusterHealthMonitorResource{}
}

type DeviceClusterHealthMonitorResource struct {
	client *fmc.Client
}

func (r *DeviceClusterHealthMonitorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_cluster_health_monitor"
}

func (r *DeviceClusterHealthMonitorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages a Device Cluster Health Monitor.").String,

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
			"cluster_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the parent cluster").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the resource; This is always `ClusterHealthMonitorSetting`.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"health_check": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable health check.").String,
				Optional:            true,
			},
			"hold_time": schema.Float64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Time (in seconds) to wait before declaring an unresponsive peer as down.").AddFloatRangeDescription(0.3, 45).String,
				Optional:            true,
				Validators: []validator.Float64{
					float64validator.Between(0.3, 45),
				},
			},
			"debounce_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The time (in milliseconds) before the interface is considered to have failed.").AddIntegerRangeDescription(300, 9000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(300, 9000),
				},
			},
			"data_interface_auto_rejoin_attempts": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Number of rejoin attempts. Use '-1' for unlimited attempts.").AddIntegerRangeDescription(-1, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(-1, 65535),
				},
			},
			"data_interface_auto_rejoin_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interval duration (in minutes) between rejoin attempts.").AddIntegerRangeDescription(2, 60).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(2, 60),
				},
			},
			"data_interface_auto_rejoin_interval_variation": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rejoin interval increase strategy. Possible values are 1 (no change); 2 (2 x the previous duration), or 3 (3 x the previous duration)").AddIntegerRangeDescription(1, 3).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 3),
				},
			},
			"cluster_interface_auto_rejoin_attempts": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Number of rejoin attempts. Use '-1' for unlimited attempts.").AddIntegerRangeDescription(-1, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(-1, 65535),
				},
			},
			"cluster_interface_auto_rejoin_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interval duration (in minutes) between rejoin attempts.").AddIntegerRangeDescription(2, 60).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(2, 60),
				},
			},
			"cluster_interface_auto_rejoin_interval_variation": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rejoin interval increase strategy. Possible values are 1 (no change); 2 (2 x the previous duration), or 3 (3 x the previous duration)").AddIntegerRangeDescription(1, 3).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 3),
				},
			},
			"system_auto_rejoin_attempts": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Number of rejoin attempts. Use '-1' for unlimited attempts.").AddIntegerRangeDescription(-1, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(-1, 65535),
				},
			},
			"system_auto_rejoin_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interval duration (in minutes) between rejoin attempts.").AddIntegerRangeDescription(2, 60).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(2, 60),
				},
			},
			"system_auto_rejoin_interval_variation": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rejoin interval increase strategy. Possible values are 1 (no change); 2 (2 x the previous duration), or 3 (3 x the previous duration)").AddIntegerRangeDescription(1, 3).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 3),
				},
			},
			"unmonitored_interfaces": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of interfaces excluded from monitoring.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"service_application_monitoring": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable service application monitoring (Snort and disk-full processes).").String,
				Optional:            true,
			},
		},
	}
}

func (r *DeviceClusterHealthMonitorResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

func (r *DeviceClusterHealthMonitorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DeviceClusterHealthMonitor

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

	// ID of the object is equal to ID of the cluster
	plan.Id = plan.ClusterId

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, DeviceClusterHealthMonitor{})
	res, err := r.client.Put(plan.getPath()+"/"+url.PathEscape(plan.Id.ValueString()), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("id").String())
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *DeviceClusterHealthMonitorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DeviceClusterHealthMonitor

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

func (r *DeviceClusterHealthMonitorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state DeviceClusterHealthMonitor

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

func (r *DeviceClusterHealthMonitorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DeviceClusterHealthMonitor

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import

func (r *DeviceClusterHealthMonitorResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <cluster_id>,<id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("cluster_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin createSubresources

// End of section. //template:end createSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin deleteSubresources

// End of section. //template:end deleteSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin updateSubresources

// End of section. //template:end updateSubresources
