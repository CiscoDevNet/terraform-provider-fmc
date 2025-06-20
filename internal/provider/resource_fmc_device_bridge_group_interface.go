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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
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
	_ resource.Resource                = &DeviceBridgeGroupInterfaceResource{}
	_ resource.ResourceWithImportState = &DeviceBridgeGroupInterfaceResource{}
)

func NewDeviceBridgeGroupInterfaceResource() resource.Resource {
	return &DeviceBridgeGroupInterfaceResource{}
}

type DeviceBridgeGroupInterfaceResource struct {
	client *fmc.Client
}

func (r *DeviceBridgeGroupInterfaceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_bridge_group_interface"
}

func (r *DeviceBridgeGroupInterfaceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages Device Bridge Group Interface, known as BVI interfaces.").String,

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
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this value is always 'BridgeGroupInterface'.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the Bridge Group interface in format BVI<bridge_group_id>.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"logical_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Logical name of the Brige Group interface.").String,
				Optional:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description of the object.").String,
				Optional:            true,
			},
			"bridge_group_id": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Bridge Group Id.").AddIntegerRangeDescription(1, 250).String,
				Required:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 250),
				},
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"selected_interfaces": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of interfaces that are part of the bridge group.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the interface").String,
							Required:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the interface").String,
							Required:            true,
						},
					},
				},
			},
			"ipv4_static_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Static IPv4 address.").String,
				Optional:            true,
			},
			"ipv4_static_netmask": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Netmask for ipv4_static_address.").String,
				Optional:            true,
			},
			"ipv4_dhcp_obtain_route": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Value `false` indicates to enable DHCPv4 without obtaining default route. Value `true` indicates to enable DHCPv4 and obtain the default route. The ipv4_dhcp_obtain_route must be null when using ipv4_static_address.").String,
				Optional:            true,
			},
			"ipv6_addresses": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv6 address without a slash and prefix.").String,
							Required:            true,
						},
						"prefix": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Prefix width for the IPv6 address.").String,
							Required:            true,
						},
					},
				},
			},
			"ipv6_enable_dad": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether to enable IPv6 DAD Loopback Detect (DAD).").String,
				Optional:            true,
			},
			"ipv6_dad_attempts": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Number of Duplicate Address Detection (DAD) attempts.").AddIntegerRangeDescription(0, 600).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 600),
				},
			},
			"ipv6_ns_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Neighbor Solicitation (NS) interval in Milliseconds.").AddIntegerRangeDescription(1000, 3600000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1000, 3600000),
				},
			},
			"ipv6_reachable_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The amount of time (in Milliseconds) that a remote IPv6 node is considered reachable after a reachability confirmation event has occurred.").AddIntegerRangeDescription(0, 3600000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 3600000),
				},
			},
			"arp_table_entries": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"mac_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("MAC address for custom ARP entry in format 0123.4567.89ab.").String,
							Optional:            true,
						},
						"ip_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP address for custom ARP entry").String,
							Optional:            true,
						},
						"enable_alias": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable Alias for custom ARP entry").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *DeviceBridgeGroupInterfaceResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *DeviceBridgeGroupInterfaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DeviceBridgeGroupInterface

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
	body := plan.toBody(ctx, DeviceBridgeGroupInterface{})
	body = plan.adjustBody(ctx, body)
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

func (r *DeviceBridgeGroupInterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DeviceBridgeGroupInterface

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

func (r *DeviceBridgeGroupInterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state DeviceBridgeGroupInterface

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
	body = plan.adjustBody(ctx, body)
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

func (r *DeviceBridgeGroupInterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DeviceBridgeGroupInterface

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

func (r *DeviceBridgeGroupInterfaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
