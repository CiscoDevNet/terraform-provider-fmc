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
	_ resource.Resource                = &DeviceBGPGeneralSettingsResource{}
	_ resource.ResourceWithImportState = &DeviceBGPGeneralSettingsResource{}
)

func NewDeviceBGPGeneralSettingsResource() resource.Resource {
	return &DeviceBGPGeneralSettingsResource{}
}

type DeviceBGPGeneralSettingsResource struct {
	client *fmc.Client
}

func (r *DeviceBGPGeneralSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_bgp_general_settings"
}

func (r *DeviceBGPGeneralSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages a Device BGP General Settings.").String,

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
			"device_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the parent device.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Object name; Always set to 'AsaBGPGeneralTable'").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"as_number": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Autonomous System (AS) number in 'asplain' or 'asdot' format").String,
				Required:            true,
			},
			"router_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("String value for the routerID. Possible values can be 'AUTOMATIC' or valid ipv4 address").String,
				Optional:            true,
			},
			"scanning_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Scanning interval of BGP routers for next hop validation in Seconds.").AddIntegerRangeDescription(5, 60).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(5, 60),
				},
			},
			"as_number_in_path_attribute": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Range to discard routes that have as-path segments that exceed a specified value.").AddIntegerRangeDescription(1, 254).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 254),
				},
			},
			"log_neighbor_changes": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Logging of BGP neighbor status changes.").String,
				Optional:            true,
			},
			"tcp_path_mtu_discovery": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use TCP path MTU discovery.").String,
				Optional:            true,
			},
			"reset_session_upon_failover": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Reset session upon failover").String,
				Optional:            true,
			},
			"enforce_first_peer_as": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Discard updates received from an external BGP (eBGP) peers that do not list their autonomous system (AS) number.").String,
				Optional:            true,
			},
			"use_dot_notation": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Change format of BGP 4-byte autonomous system numbers from asplain (decimal values) to dot notation.").String,
				Optional:            true,
			},
			"aggregate_timer": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interval at which BGP routes will be aggregated or to disable timer-based router aggregation (in seconds).").AddIntegerRangeDescription(6, 60).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(6, 60),
				},
			},
			"default_local_preference": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Default local preference").AddIntegerRangeDescription(0, 4294967295).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 4294967295),
				},
			},
			"compare_med_from_different_neighbors": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allow comparing MED from different neighbors").String,
				Optional:            true,
			},
			"compare_router_id_in_path": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Compare Router ID for identical EBGP paths").String,
				Optional:            true,
			},
			"pick_best_med": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Pick the best-MED path among paths advertised by neighbor AS").String,
				Optional:            true,
			},
			"missing_med_as_best": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Treat missing MED as the best preferred path").String,
				Optional:            true,
			},
			"keepalive_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Keepalive interval in seconds").AddIntegerRangeDescription(0, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
				},
			},
			"hold_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Hold time in seconds").AddIntegerRangeDescription(0, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
				},
			},
			"min_hold_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Minimum hold time (0 or 3-65535 seconds)").AddIntegerRangeDescription(0, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
				},
			},
			"next_hop_address_tracking": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable next hop address tracking").String,
				Optional:            true,
			},
			"next_hop_delay_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Next hop delay interval in seconds").AddIntegerRangeDescription(0, 100).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 100),
				},
			},
			"graceful_restart": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable graceful restart").String,
				Optional:            true,
			},
			"graceful_restart_restart_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Graceful Restart Time in seconds").AddIntegerRangeDescription(1, 3600).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 3600),
				},
			},
			"graceful_restart_stale_path_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Stalepath Time in seconds").AddIntegerRangeDescription(1, 3600).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 3600),
				},
			},
		},
	}
}

func (r *DeviceBGPGeneralSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *DeviceBGPGeneralSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DeviceBGPGeneralSettings

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
	body := plan.toBody(ctx, DeviceBGPGeneralSettings{})
	res, err := r.client.Post(plan.getPath(), body, reqMods...)
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

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *DeviceBGPGeneralSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DeviceBGPGeneralSettings

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

func (r *DeviceBGPGeneralSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state DeviceBGPGeneralSettings

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

func (r *DeviceBGPGeneralSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DeviceBGPGeneralSettings

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

func (r *DeviceBGPGeneralSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <device_id>,<id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("device_id"), idParts[0])...)
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
