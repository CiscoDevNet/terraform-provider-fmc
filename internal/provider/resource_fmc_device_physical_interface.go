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
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
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
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages a Device Physical Interface.").String,

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
				MarkdownDescription: helpers.NewAttributeDescription("Type of the resource.").String,
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
				MarkdownDescription: helpers.NewAttributeDescription("Whether to enable the interface.").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"management_only": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether this interface limits traffic to management traffic; when true, through-the-box traffic is disallowed. Value true conflicts with mode INLINE, PASSIVE, TAP, ERSPAN, or with security_zone_id.").String,
				Optional:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Optional user-created description.").String,
				Optional:            true,
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Mode of the interface. Use INLINE if, and only if, the interface is part of fmc_inline_set with tap_mode=false or tap_mode unset. Use TAP if, and only if, the interface is part of fmc_inline_set with tap_mode = true. Use ERSPAN only when both erspan_source_ip and erspan_flow_id are set.").AddStringEnumDescription("INLINE", "PASSIVE", "TAP", "ERSPAN", "NONE", "SWITCHPORT").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("INLINE", "PASSIVE", "TAP", "ERSPAN", "NONE", "SWITCHPORT"),
				},
			},
			"security_zone_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the assigned security zone. Can only be used when `logical_name` is set.").String,
				Optional:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the interface; it must already be present on the device.").String,
				Required:            true,
			},
			"mtu": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Maximum transmission unit. Can only be used when logical_name is set.").AddIntegerRangeDescription(64, 9000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(64, 9000),
				},
			},
			"priority": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Priority. Can only be set for routed interfaces.").AddIntegerRangeDescription(0, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
				},
			},
			"enable_sgt_propagate": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether to propagate SGT.").String,
				Optional:            true,
			},
			"nve_only": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Used for VTEP's source interface to restrict it to NVE only. For routed mode (NONE mode) the `nve_only` restricts interface to VxLAN traffic and common management traffic. For transparent firewall modes, the `nve_only` is automatically enabled.").String,
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
			"ipv4_address_pool_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the assigned IPv4 address pool.").String,
				Optional:            true,
			},
			"ipv4_dhcp_obtain_default_route": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Any non-null value here indicates to enable DHCPv4. Value `false` indicates to enable DHCPv4 without obtaining default IPv4 route but anyway requires also `ipv4_dhcp_route_metric` to be set to exactly 1. Value `true` indicates to enable DHCPv4 and obtain the route and also requires `ipv4_dhcp_route_metric` to be non-null. The `ipv4_dhcp_obtain_default_route` must be null when using `ipv4_static_address`.").String,
				Optional:            true,
			},
			"ipv4_dhcp_default_route_metric": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The metric for `ipv4_dhcp_obtain_default_route`. Any non-null value enables DHCP as a side effect. Must be null when using `ipv4_static_address`.").AddIntegerRangeDescription(1, 255).String,
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
			"ipv6": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether to enable IPv6.").String,
				Optional:            true,
			},
			"ipv6_enforce_eui": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether to enforce IPv6 Extended Unique Identifier (EUI64 from RFC2373).").String,
				Optional:            true,
			},
			"ipv6_link_local_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv6 Configuration - Link-Local Address.").String,
				Optional:            true,
			},
			"ipv6_auto_config": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether to enable IPv6 autoconfiguration.").String,
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
							MarkdownDescription: helpers.NewAttributeDescription("Whether to enforce IPv6 Extended Unique Identifier (EUI64 from RFC2373).").String,
							Optional:            true,
						},
					},
				},
			},
			"ipv6_address_pool_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the assigned IPv6 address pool.").String,
				Optional:            true,
			},
			"ipv6_prefixes": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv6 address with the prefix length.").String,
							Optional:            true,
						},
						"default": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Use default prefix.").String,
							Optional:            true,
						},
					},
				},
			},
			"ipv6_dad": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether to enable IPv6 DAD Loopback Detect (DAD).").String,
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
			"ipv6_ra": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether to enable IPv6 router advertisement (RA).").String,
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
			"ipv6_dhcp_obtain_default_route": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether to obtain default route from DHCPv6.").String,
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
			"ipv6_dhcp_address_config": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether to enable DHCPv6 for address config.").String,
				Optional:            true,
			},
			"ipv6_dhcp_nonaddress_config": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether to enable DHCPv6 for non-address config.").String,
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
				MarkdownDescription: helpers.NewAttributeDescription("Whether to enable IP based Monitoring.").String,
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
			"auto_negotiation": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enables auto negotiation of duplex and speed.").String,
				Optional:            true,
			},
			"duplex": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Duplex configuraion.").AddStringEnumDescription("AUTO", "FULL", "HALF").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("AUTO", "FULL", "HALF"),
				},
			},
			"speed": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Speed configuraion, can be one of AUTO, TEN, HUNDRED, THOUSAND, TEN_THOUSAND, TWENTY_FIVE_THOUSAND, FORTY_THOUSAND, HUNDRED_THOUSAND, TWO_HUNDRED_THOUSAND, DETECT_SFP").AddStringEnumDescription("AUTO", "TEN", "HUNDRED", "THOUSAND", "TEN_THOUSAND", "TWENTY_FIVE_THOUSAND", "FORTY_THOUSAND", "HUNDRED_THOUSAND", "TWO_HUNDRED_THOUSAND", "DETECT_SFP").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("AUTO", "TEN", "HUNDRED", "THOUSAND", "TEN_THOUSAND", "TWENTY_FIVE_THOUSAND", "FORTY_THOUSAND", "HUNDRED_THOUSAND", "TWO_HUNDRED_THOUSAND", "DETECT_SFP"),
				},
			},
			"lldp_receive": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("LLDP receive configuraion.").String,
				Optional:            true,
			},
			"lldp_transmit": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("LLDP transmit configuraion.").String,
				Optional:            true,
			},
			"flow_control_send": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Flow Control Send configuraion, can be one of ON, OFF.").AddStringEnumDescription("ON", "OFF").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("ON", "OFF"),
				},
			},
			"fec_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Path Monitoring - Monitoring Type, can be one of AUTO, CL108RS, CL74FC, CL91RS, DISABLE.").AddStringEnumDescription("AUTO", "CL108RS", "CL74FC", "CL91RS", "DISABLE").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("AUTO", "CL108RS", "CL74FC", "CL91RS", "DISABLE"),
				},
			},
			"management_access": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether to enable Management Access.").String,
				Optional:            true,
			},
			"management_access_network_objects": schema.SetNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of the network object (Host, Network or Range).").String,
							Optional:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of the object").String,
							Optional:            true,
						},
					},
				},
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
							Required:            true,
						},
						"ip_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP address for custom ARP entry").String,
							Required:            true,
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable Alias for custom ARP entry").AddDefaultValueDescription("true").String,
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(true),
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
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
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
			queryString := fmt.Sprintf("?limit=%d&offset=%d&expanded=true", limit, offset)
			res, err := r.client.Get(plan.getPath()+queryString, reqMods...)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects, got error: %s", err))
				return
			}
			if value := res.Get("items"); len(value.Array()) > 0 {
				value.ForEach(func(k, v gjson.Result) bool {
					if plan.Name.ValueString() == v.Get("name").String() {
						plan.Id = types.StringValue(v.Get("id").String())
						tflog.Debug(ctx, fmt.Sprintf("%s: Found object with name '%v', id: %s", plan.Id.ValueString(), plan.Name.ValueString(), plan.Id.ValueString()))
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
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with name: %v", plan.Name.ValueString()))
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
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *DevicePhysicalInterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DevicePhysicalInterface

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

func (r *DevicePhysicalInterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state DevicePhysicalInterface

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

func (r *DevicePhysicalInterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DevicePhysicalInterface

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

	// This is physical interface, so delete is done by PUT with minimal body
	// This needs to be done in two steps due to some dependencies in the FMC
	// Step 1: Remove all attributes except 'ifname'
	body := state.toBodyPutDelete(ctx)
	res, err := r.client.Put(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString()), body, reqMods...)
	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
		tflog.Debug(ctx, fmt.Sprintf("%s: Interface not found", state.Id.ValueString()))
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to remove object configuration phase 1 (PUT), got error: %s, %s", err, res.String()))
		return
	}

	// Step 2: Remove 'ifname' (if still configured) from body and re-run request
	if !state.LogicalName.IsNull() {
		body, _ = sjson.Delete(body, "ifname")
		res, err = r.client.Put(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString()), body, reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to remove object configuration phase 2 (PUT), got error: %s, %s", err, res.String()))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *DevicePhysicalInterfaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Parse import ID
	var inputPattern = regexp.MustCompile(`^(?:(?P<domain>[^\s,]+),)?(?P<device_id>[^\s,]+),(?P<id>[^\s,]+?)$`)
	match := inputPattern.FindStringSubmatch(req.ID)
	if match == nil {
		errMsg := "Failed to parse import parameters.\nPlease provide import string in the following format: <domain>,<device_id>,<id>\n<domain> is optional. If not provided, `Global` is used implicitly and resource's `domain` attribute is not set.\n" + fmt.Sprintf("Got: %q", req.ID)
		resp.Diagnostics.AddError("Import error", errMsg)
		return
	}

	// Set domain, if provided
	if tmpDomain := match[inputPattern.SubexpIndex("domain")]; tmpDomain != "" {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("domain"), tmpDomain)...)
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), match[inputPattern.SubexpIndex("id")])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("device_id"), match[inputPattern.SubexpIndex("device_id")])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
