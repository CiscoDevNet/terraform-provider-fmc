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

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/netascode/terraform-provider-fmc/internal/provider/helpers"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &DevicePhysicalInterfaceResource{}
	_ resource.ResourceWithImportState = &DevicePhysicalInterfaceResource{}
)

func NewDevicePhysicalInterfaceResource() resource.Resource {
	return &DevicePhysicalInterfaceResource{}
}

type DevicePhysicalInterfaceResource struct {
	client *fmc.Client
}

func (r *DevicePhysicalInterfaceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_physical_interface"
}

func (r *DevicePhysicalInterfaceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Device Physical Interface.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "The name of the FMC domain",
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"device_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("UUID of the parent device (fmc_device.example.id).").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether to enable the interface.").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Mode of the interface. Use INLINE if, and only if, the interface is part of fmc_inline_set with tap_mode=false or tap_mode unset. Use TAP if, and only if, the interface is part of fmc_inline_set with tap_mode = true. Use ERSPAN only when both erspan_source_ip and erspan_flow_id are set.").AddStringEnumDescription("INLINE", "PASSIVE", "TAP", "ERSPAN", "NONE", "SWITCHPORT").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("INLINE", "PASSIVE", "TAP", "ERSPAN", "NONE", "SWITCHPORT"),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the interface; it must already be present on the device.").String,
				Required:            true,
			},
			"logical_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Customizable logical name of the interface, unique on the device. Should not contain whitespace or slash characters. Must be non-empty in order to set security_zone_id, mtu, inline sets, etc.").String,
				Optional:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
			},
			"management_only": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether this interface limits traffic to management traffic; when true, through-the-box traffic is disallowed. Value true conflicts with mode INLINE, PASSIVE, TAP, ERSPAN, or with security_zone_id.").String,
				Optional:            true,
			},
			"mtu": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Maximum transmission unit. Can only be used when logical_name is set.").String,
				Optional:            true,
			},
			"priority": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Priority 0-65535. Can only be set for routed interfaces.").String,
				Optional:            true,
			},
			"security_zone_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("UUID of the assigned security zone (fmc_security_zone.example.id). Can only be used when logical_name is set.").String,
				Optional:            true,
			},
			"ipv4_static_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Static IPv4 address. Conflicts with mode INLINE, PASSIVE, TAP, ERSPAN.").String,
				Optional:            true,
			},
			"ipv4_static_netmask": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Netmask (width) for ipv4_static_address.").String,
				Optional:            true,
			},
			"ipv4_dhcp_obtain_route": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Any non-null value here indicates to enable DHCPv4. Value `false` indicates to enable DHCPv4 without obtaining from there the default IPv4 route but anyway requires also ipv4_dhcp_route_metric to be set to exactly 1. Value `true` indicates to enable DHCPv4 and obtain the route and also requires ipv4_dhcp_route_metric to be non-null. The ipv4_dhcp_obtain_route must be null when using ipv4_static_netmask.").String,
				Optional:            true,
			},
			"ipv4_dhcp_route_metric": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The metric for ipv4_dhcp_obtain_route. Any non-null value enables DHCP as a side effect. Must be null when using ipv4_static_netmask.").AddIntegerRangeDescription(1, 255).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
			},
			"ipv6_enable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether to enable IPv6.").String,
				Optional:            true,
			},
			"ipv6_enforce_eui": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether to enforce IPv6 Extended Unique Identifier (EUI64 from RFC2373).").String,
				Optional:            true,
			},
			"ipv6_enable_auto_config": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether to enable IPv6 autoconfiguration.").String,
				Optional:            true,
			},
			"ipv6_enable_dhcp_address": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether to enable DHCPv6 for address config.").String,
				Optional:            true,
			},
			"ipv6_enable_dhcp_nonaddress": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether to enable DHCPv6 for non-address config.").String,
				Optional:            true,
			},
			"ipv6_enable_ra": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether to enable IPv6 router advertisement (RA).").String,
				Optional:            true,
			},
			"ipv6_addresses": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv6 address without a slash and prefix.").String,
							Optional:            true,
						},
						"prefix": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Prefix width for the IPv6 address.").String,
							Optional:            true,
						},
						"enforce_eui": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Indicates whether to enforce IPv6 Extended Unique Identifier (EUI64 from RFC2373).").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *DevicePhysicalInterfaceResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *DevicePhysicalInterfaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DevicePhysicalInterface

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !plan.Domain.IsNull() && plan.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(plan.Domain.ValueString()))
	}
	tflog.Debug(ctx, fmt.Sprintf("%s: considering object name %s", plan.Id, plan.Name))

	if plan.Id.ValueString() == "" && plan.Name.ValueString() != "" {
		offset := 0
		limit := 1000
		for page := 1; ; page++ {
			queryString := fmt.Sprintf("?limit=%d&offset=%d", limit, offset)
			res, err := r.client.Get(plan.getPath()+queryString, reqMods...)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects, got error: %s", err))
				return
			}
			if value := res.Get("items"); len(value.Array()) > 0 {
				value.ForEach(func(k, v gjson.Result) bool {
					if plan.Name.ValueString() == v.Get("name").String() {
						plan.Id = types.StringValue(v.Get("id").String())
						tflog.Debug(ctx, fmt.Sprintf("%s: Found object with name '%s', id: %s", plan.Id, plan.Name.ValueString(), plan.Id))
						return false
					}
					return true
				})
			}
			if plan.Id.ValueString() != "" || !res.Get("paging.next.0").Exists() {
				break
			}
			offset += limit
		}

		if plan.Id.ValueString() == "" {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with name: %s", plan.Name.ValueString()))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, DevicePhysicalInterface{})
	res, err := r.client.Put(plan.getPath()+"/"+url.PathEscape(plan.Id.ValueString()), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("id").String())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *DevicePhysicalInterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DevicePhysicalInterface

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !state.Domain.IsNull() && state.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(state.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	res, err := r.client.Get(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString()), reqMods...)
	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	// If every attribute is set to null we are dealing with an import operation and therefore reading all attributes
	if state.isNull(ctx, res) {
		state.fromBody(ctx, res)
	} else {
		state.updateFromBody(ctx, res)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update

func (r *DevicePhysicalInterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state DevicePhysicalInterface

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
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

func (r *DevicePhysicalInterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DevicePhysicalInterface

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
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

func (r *DevicePhysicalInterfaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
}

// End of section. //template:end import
