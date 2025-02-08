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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
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
	_ resource.Resource                = &DeviceSubinterfaceResource{}
	_ resource.ResourceWithImportState = &DeviceSubinterfaceResource{}
)

func NewDeviceSubinterfaceResource() resource.Resource {
	return &DeviceSubinterfaceResource{}
}

type DeviceSubinterfaceResource struct {
	client *fmc.Client
}

func (r *DeviceSubinterfaceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_subinterface"
}

func (r *DeviceSubinterfaceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages a Device Subinterface.").String,

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
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the subinterface in format `interface_name.subinterface_id` (eg. GigabitEthernet0/1.7).").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"logical_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Logical name of the interface, unique on the device. Should not contain whitespace or slash characters.").String,
				Optional:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether to enable the interface.").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"management_only": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether this interface limits traffic to management traffic; when true, through-the-box traffic is disallowed. Value true conflicts with mode INLINE, PASSIVE, TAP, ERSPAN, or with security_zone_id.").String,
				Optional:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Optional user-created description.").String,
				Optional:            true,
			},
			"security_zone_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the assigned security zone. Can only be used when logical_name is set.").String,
				Optional:            true,
			},
			"mtu": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Maximum transmission unit. Can only be used when logical_name is set.").AddIntegerRangeDescription(64, 9000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(64, 9000),
				},
			},
			"priority": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Priority 0-65535. Can only be set for routed interfaces.").AddIntegerRangeDescription(0, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
				},
			},
			"enable_sgt_propagate": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether to propagate SGT.").String,
				Optional:            true,
			},
			"interface_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the parent interface (fmc_device_physical_interface.example.name).").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"sub_interface_id": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The numerical id of this subinterface, unique on the parent interface.").AddIntegerRangeDescription(0, 4294967295).String,
				Required:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 4294967295),
				},
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"vlan_id": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("VLAN identifier, unique per the parent interface.").AddIntegerRangeDescription(1, 4094).String,
				Required:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4094),
				},
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
				MarkdownDescription: helpers.NewAttributeDescription("Any non-null value here indicates to enable DHCPv4. Value `false` indicates to enable DHCPv4 without obtaining from there the default IPv4 route but anyway requires also ipv4_dhcp_route_metric to be set to exactly 1. Value `true` indicates to enable DHCPv4 and obtain the route and also requires ipv4_dhcp_route_metric to be non-null. The ipv4_dhcp_obtain_route must be null when using ipv4_static_address.").String,
				Optional:            true,
			},
			"ipv4_dhcp_route_metric": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The metric for ipv4_dhcp_obtain_route. Any non-null value enables DHCP as a side effect. Must be null when using ipv4_static_address.").AddIntegerRangeDescription(1, 255).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
			},
			"ipv4_pppoe_vpdn_group_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("PPPoE Configuration - PPPoE Group Name.").String,
				Optional:            true,
			},
			"ipv4_pppoe_user": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("PPPoE Configuration - PPPoE User.").String,
				Optional:            true,
			},
			"ipv4_pppoe_password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("PPPoE Configuration - PPPoE Password.").String,
				Optional:            true,
			},
			"ipv4_pppoe_authentication": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("PPPoE Configuration - PPPoE Authentication, can be one of PAP, CHAP, MSCHAP.").AddStringEnumDescription("PAP", "CHAP", "MSCHAP").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("PAP", "CHAP", "MSCHAP"),
				},
			},
			"ipv4_pppoe_route_metric": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("PPPoE Configuration - PPPoE route metric, can be value between 1 - 255.").AddIntegerRangeDescription(1, 255).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
			},
			"ipv4_pppoe_route_settings": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("PPPoE Configuration - PPPoE Enable Route Settings.").String,
				Optional:            true,
			},
			"ipv4_pppoe_store_credentials_in_flash": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("PPPoE Configuration - PPPoE store username and password in Flash.").String,
				Optional:            true,
			},
			"ipv6_enable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether to enable IPv6.").String,
				Optional:            true,
			},
			"ipv6_enforce_eui": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether to enforce IPv6 Extended Unique Identifier (EUI64 from RFC2373).").String,
				Optional:            true,
			},
			"ipv6_link_local_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv6 Configuration - Link-Local Address.").String,
				Optional:            true,
			},
			"ipv6_enable_auto_config": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether to enable IPv6 autoconfiguration.").String,
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
			"ipv6_prefixes": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv6 address without a slash and prefix.").String,
							Optional:            true,
						},
						"default": schema.StringAttribute{
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
				MarkdownDescription: helpers.NewAttributeDescription("Neighbor Solicitation (NS) interval.").AddIntegerRangeDescription(1000, 3600000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1000, 3600000),
				},
			},
			"ipv6_reachable_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The amount of time that a remote IPv6 node is considered reachable after a reachability confirmation event has occurred").AddIntegerRangeDescription(0, 3600000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 3600000),
				},
			},
			"ipv6_enable_ra": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether to enable IPv6 router advertisement (RA).").String,
				Optional:            true,
			},
			"ipv6_ra_life_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Router Advertisement (RA) lifetime.").AddIntegerRangeDescription(0, 9000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 9000),
				},
			},
			"ipv6_ra_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interval between Router Advertisements (RA) transmissions").AddIntegerRangeDescription(3, 1800).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(3, 1800),
				},
			},
			"ipv6_dhcp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable DHCPv6 client.").String,
				Optional:            true,
			},
			"ipv6_default_route_by_dhcp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether to obtain default route from DHCPv6.").String,
				Optional:            true,
			},
			"ipv6_dhcp_pool_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the assigned DHCPv6 pool").String,
				Optional:            true,
			},
			"ipv6_dhcp_pool_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this value is always 'IPv6AddressPool'.").String,
				Optional:            true,
			},
			"ipv6_enable_dhcp_address_config": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether to enable DHCPv6 for address config.").String,
				Optional:            true,
			},
			"ipv6_enable_dhcp_nonaddress_config": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether to enable DHCPv6 for non-address config.").String,
				Optional:            true,
			},
			"ipv6_dhcp_client_pd_prefix_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Prefix Name for Prefix Delegation (PD)").String,
				Optional:            true,
			},
			"ipv6_dhcp_client_pd_hint_prefixes": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Hint Prefixes for Prefix Delegation (PD)").String,
				Optional:            true,
			},
			"ip_based_monitoring": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether to enable IP based Monitoring.").String,
				Optional:            true,
			},
			"ip_based_monitoring_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("PPPoE Configuration - PPPoE route metric, [ AUTO, PEER_IPV4, PEER_IPV6, AUTO4, AUTO6 ]").AddStringEnumDescription("AUTO", "PEER_IPV4", "PEER_IPV6", "AUTO4", "AUTO6").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("AUTO", "PEER_IPV4", "PEER_IPV6", "AUTO4", "AUTO6"),
				},
			},
			"ip_based_monitoring_next_hop": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP address to monitor.").String,
				Optional:            true,
			},
			"active_mac_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("MAC address for active interface in format 0123.4567.89ab.").String,
				Optional:            true,
			},
			"standby_mac_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("MAC address for standby interface in format 0123.4567.89ab.").String,
				Optional:            true,
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
			"enable_anti_spoofing": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Anti Spoofing").String,
				Optional:            true,
			},
			"allow_full_fragment_reassembly": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allow Full Fragment Reassembly").String,
				Optional:            true,
			},
			"override_default_fragment_setting_chain": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Override Default Fragment Setting - Chain value").AddIntegerRangeDescription(1, 8200).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 8200),
				},
			},
			"override_default_fragment_setting_size": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Override Default Fragment Setting - Fragment Size value").AddIntegerRangeDescription(1, 30000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 30000),
				},
			},
			"override_default_fragment_setting_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Override Default Fragment Setting - Time Out value").AddIntegerRangeDescription(1, 30).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 30),
				},
			},
		},
	}
}

func (r *DeviceSubinterfaceResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

func (r *DeviceSubinterfaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DeviceSubinterface

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
	body := plan.toBody(ctx, DeviceSubinterface{})
	res, err := r.client.Post(plan.getPath(), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("id").String())
	plan.fromBodyUnknowns(ctx, res)

	// Fix 'name`
	plan.Name = types.StringValue(fmt.Sprintf("%s.%d", plan.Name.ValueString(), plan.SubInterfaceId.ValueInt64()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *DeviceSubinterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DeviceSubinterface

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

	// Fix 'name`
	state.Name = types.StringValue(fmt.Sprintf("%s.%d", state.Name.ValueString(), state.SubInterfaceId.ValueInt64()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *DeviceSubinterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state DeviceSubinterface

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

	// Fix 'name`
	plan.Name = types.StringValue(fmt.Sprintf("%s.%d", strings.Split(plan.Name.ValueString(), ".")[0], plan.SubInterfaceId.ValueInt64()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *DeviceSubinterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DeviceSubinterface

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

func (r *DeviceSubinterfaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
