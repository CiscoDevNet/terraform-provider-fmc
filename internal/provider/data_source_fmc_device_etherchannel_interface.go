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

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &DeviceEtherChannelInterfaceDataSource{}
	_ datasource.DataSourceWithConfigure = &DeviceEtherChannelInterfaceDataSource{}
)

func NewDeviceEtherChannelInterfaceDataSource() datasource.DataSource {
	return &DeviceEtherChannelInterfaceDataSource{}
}

type DeviceEtherChannelInterfaceDataSource struct {
	client *fmc.Client
}

func (d *DeviceEtherChannelInterfaceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_etherchannel_interface"
}

func (d *DeviceEtherChannelInterfaceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the Device EtherChannel Interface.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the object",
				Optional:            true,
				Computed:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "Name of the FMC domain",
				Optional:            true,
			},
			"device_id": schema.StringAttribute{
				MarkdownDescription: "Id of the parent device.",
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the object.",
				Computed:            true,
			},
			"is_multi_instance": schema.BoolAttribute{
				MarkdownDescription: "Is parent device multi-instance.",
				Computed:            true,
			},
			"logical_name": schema.StringAttribute{
				MarkdownDescription: "Logical name of the interface, unique on the device. Should not contain whitespace or slash characters.",
				Computed:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable the interface.",
				Computed:            true,
			},
			"management_only": schema.BoolAttribute{
				MarkdownDescription: "Whether this interface limits traffic to management traffic; when true, through-the-box traffic is disallowed. Value true conflicts with mode INLINE, PASSIVE, TAP, ERSPAN, or with security_zone_id.",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional user-created description.",
				Computed:            true,
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "Mode of the interface. Use INLINE if, and only if, the interface is part of fmc_inline_set with tap_mode=false or tap_mode unset. Use TAP if, and only if, the interface is part of fmc_inline_set with tap_mode = true. Use ERSPAN only when both erspan_source_ip and erspan_flow_id are set.",
				Computed:            true,
			},
			"security_zone_id": schema.StringAttribute{
				MarkdownDescription: "Id of the assigned security zone.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the interface; it must already be present on the device.",
				Optional:            true,
				Computed:            true,
			},
			"mtu": schema.Int64Attribute{
				MarkdownDescription: "Maximum transmission unit. Can only be used when logical_name is set.",
				Computed:            true,
			},
			"priority": schema.Int64Attribute{
				MarkdownDescription: "Priority. Can only be set for routed interfaces.",
				Computed:            true,
			},
			"enable_sgt_propagate": schema.BoolAttribute{
				MarkdownDescription: "Enable SGT propagation.",
				Computed:            true,
			},
			"ether_channel_id": schema.StringAttribute{
				MarkdownDescription: "Value of Ether Channel ID, allowed range 1 to 48.",
				Computed:            true,
			},
			"selected_interfaces": schema.SetNestedAttribute{
				MarkdownDescription: "Set of objects representing physical interfaces.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Id of the object.",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Type of the selected interface",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Name of the selected interface",
							Computed:            true,
						},
					},
				},
			},
			"nve_only": schema.BoolAttribute{
				MarkdownDescription: "Used for VTEP's source interface to restrict it to NVE only. For routed mode (NONE mode) the `nve_only` restricts interface to VxLAN traffic and common management traffic. For transparent firewall modes, the `nve_only` is automatically enabled.",
				Computed:            true,
			},
			"ipv4_static_address": schema.StringAttribute{
				MarkdownDescription: "Static IPv4 address. Conflicts with mode INLINE, PASSIVE, TAP, ERSPAN.",
				Computed:            true,
			},
			"ipv4_static_netmask": schema.StringAttribute{
				MarkdownDescription: "Netmask (width) for ipv4_static_address.",
				Computed:            true,
			},
			"ipv4_address_pool_id": schema.StringAttribute{
				MarkdownDescription: "Id of the assigned IPv4 address pool.",
				Computed:            true,
			},
			"ipv4_dhcp_obtain_default_route": schema.BoolAttribute{
				MarkdownDescription: "Any non-null value here indicates to enable DHCPv4. Value `false` indicates to enable DHCPv4 without obtaining default IPv4 route but anyway requires also `ipv4_dhcp_route_metric` to be set to exactly 1. Value `true` indicates to enable DHCPv4 and obtain the route and also requires `ipv4_dhcp_route_metric` to be non-null. The `ipv4_dhcp_obtain_default_route` must be null when using `ipv4_static_address`.",
				Computed:            true,
			},
			"ipv4_dhcp_default_route_metric": schema.Int64Attribute{
				MarkdownDescription: "The metric for `ipv4_dhcp_obtain_default_route`. Any non-null value enables DHCP as a side effect. Must be null when using `ipv4_static_address`.",
				Computed:            true,
			},
			"ipv4_pppoe_vpdn_group_name": schema.StringAttribute{
				MarkdownDescription: "PPPoE Configuration - PPPoE Group Name.",
				Computed:            true,
			},
			"ipv4_pppoe_user": schema.StringAttribute{
				MarkdownDescription: "PPPoE Configuration - PPPoE User.",
				Computed:            true,
			},
			"ipv4_pppoe_password": schema.StringAttribute{
				MarkdownDescription: "PPPoE Configuration - PPPoE Password.",
				Computed:            true,
			},
			"ipv4_pppoe_authentication": schema.StringAttribute{
				MarkdownDescription: "PPPoE Configuration - PPPoE Authentication, can be one of PAP, CHAP, MSCHAP.",
				Computed:            true,
			},
			"ipv4_pppoe_route_metric": schema.Int64Attribute{
				MarkdownDescription: "PPPoE Configuration - PPPoE route metric, can be value between 1 - 255.",
				Computed:            true,
			},
			"ipv4_pppoe_route_settings": schema.BoolAttribute{
				MarkdownDescription: "PPPoE Configuration - PPPoE Enable Route Settings.",
				Computed:            true,
			},
			"ipv4_pppoe_store_credentials_in_flash": schema.BoolAttribute{
				MarkdownDescription: "PPPoE Configuration - PPPoE store username and password in Flash.",
				Computed:            true,
			},
			"ipv6": schema.BoolAttribute{
				MarkdownDescription: "Whether to enable IPv6.",
				Computed:            true,
			},
			"ipv6_enforce_eui": schema.BoolAttribute{
				MarkdownDescription: "Whether to enforce IPv6 Extended Unique Identifier (EUI64 from RFC2373).",
				Computed:            true,
			},
			"ipv6_link_local_address": schema.StringAttribute{
				MarkdownDescription: "IPv6 Configuration - Link-Local Address.",
				Computed:            true,
			},
			"ipv6_auto_config": schema.BoolAttribute{
				MarkdownDescription: "Whether to enable IPv6 autoconfiguration.",
				Computed:            true,
			},
			"ipv6_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: "IPv6 address without a slash and prefix.",
							Computed:            true,
						},
						"prefix": schema.StringAttribute{
							MarkdownDescription: "Prefix width for the IPv6 address.",
							Computed:            true,
						},
						"enforce_eui": schema.BoolAttribute{
							MarkdownDescription: "Whether to enforce IPv6 Extended Unique Identifier (EUI64 from RFC2373).",
							Computed:            true,
						},
					},
				},
			},
			"ipv6_address_pool_id": schema.StringAttribute{
				MarkdownDescription: "Id of the assigned IPv6 address pool.",
				Computed:            true,
			},
			"ipv6_prefixes": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: "IPv6 address with the prefix length.",
							Computed:            true,
						},
						"default": schema.BoolAttribute{
							MarkdownDescription: "Use default prefix.",
							Computed:            true,
						},
					},
				},
			},
			"ipv6_dad": schema.BoolAttribute{
				MarkdownDescription: "Whether to enable IPv6 DAD Loopback Detect (DAD).",
				Computed:            true,
			},
			"ipv6_dad_attempts": schema.Int64Attribute{
				MarkdownDescription: "Number of Duplicate Address Detection (DAD) attempts.",
				Computed:            true,
			},
			"ipv6_ns_interval": schema.Int64Attribute{
				MarkdownDescription: "Neighbor Solicitation (NS) interval.",
				Computed:            true,
			},
			"ipv6_reachable_time": schema.Int64Attribute{
				MarkdownDescription: "The amount of time that a remote IPv6 node is considered reachable after a reachability confirmation event has occurred",
				Computed:            true,
			},
			"ipv6_ra": schema.BoolAttribute{
				MarkdownDescription: "Whether to enable IPv6 router advertisement (RA).",
				Computed:            true,
			},
			"ipv6_ra_life_time": schema.Int64Attribute{
				MarkdownDescription: "Router Advertisement (RA) lifetime.",
				Computed:            true,
			},
			"ipv6_ra_interval": schema.Int64Attribute{
				MarkdownDescription: "Interval between Router Advertisements (RA) transmissions",
				Computed:            true,
			},
			"ipv6_dhcp": schema.BoolAttribute{
				MarkdownDescription: "Enable DHCPv6 client.",
				Computed:            true,
			},
			"ipv6_dhcp_obtain_default_route": schema.BoolAttribute{
				MarkdownDescription: "Whether to obtain default route from DHCPv6.",
				Computed:            true,
			},
			"ipv6_dhcp_pool_id": schema.StringAttribute{
				MarkdownDescription: "Id of the assigned DHCPv6 pool",
				Computed:            true,
			},
			"ipv6_dhcp_pool_type": schema.StringAttribute{
				MarkdownDescription: "Type of the object; this value is always 'IPv6AddressPool'.",
				Computed:            true,
			},
			"ipv6_dhcp_address_config": schema.BoolAttribute{
				MarkdownDescription: "Whether to enable DHCPv6 for address config.",
				Computed:            true,
			},
			"ipv6_dhcp_nonaddress_config": schema.BoolAttribute{
				MarkdownDescription: "Whether to enable DHCPv6 for non-address config.",
				Computed:            true,
			},
			"ipv6_dhcp_client_pd_prefix_name": schema.StringAttribute{
				MarkdownDescription: "Prefix Name for Prefix Delegation",
				Computed:            true,
			},
			"ipv6_dhcp_client_pd_hint_prefixes": schema.StringAttribute{
				MarkdownDescription: "Hint Prefixes for Prefix Delegation (PD)",
				Computed:            true,
			},
			"ip_based_monitoring": schema.BoolAttribute{
				MarkdownDescription: "Whether to enable IP based Monitoring.",
				Computed:            true,
			},
			"ip_based_monitoring_type": schema.StringAttribute{
				MarkdownDescription: "PPPoE Configuration - PPPoE route metric, [ AUTO, PEER_IPV4, PEER_IPV6, AUTO4, AUTO6 ]",
				Computed:            true,
			},
			"ip_based_monitoring_next_hop": schema.StringAttribute{
				MarkdownDescription: "IP address to monitor.",
				Computed:            true,
			},
			"auto_negotiation": schema.BoolAttribute{
				MarkdownDescription: "Enables auto negotiation of duplex and speed.",
				Computed:            true,
			},
			"duplex": schema.StringAttribute{
				MarkdownDescription: "Duplex configuraion, can be one of INLINE, PASSIVE, TAP, ERSPAN.",
				Computed:            true,
			},
			"speed": schema.StringAttribute{
				MarkdownDescription: "Speed configuraion, can be one of AUTO, TEN, HUNDRED, THOUSAND, TEN_THOUSAND, TWENTY_FIVE_THOUSAND, FORTY_THOUSAND, HUNDRED_THOUSAND, TWO_HUNDRED_THOUSAND, DETECT_SFP",
				Computed:            true,
			},
			"lldp_receive": schema.BoolAttribute{
				MarkdownDescription: "LLDP receive configuraion.",
				Computed:            true,
			},
			"lldp_transmit": schema.BoolAttribute{
				MarkdownDescription: "LLDP transmit configuraion.",
				Computed:            true,
			},
			"flow_control_send": schema.StringAttribute{
				MarkdownDescription: "Flow Control Send configuraion, can be one of ON, OFF.",
				Computed:            true,
			},
			"fec_mode": schema.StringAttribute{
				MarkdownDescription: "Path Monitoring - Monitoring Type, can be one of AUTO, CL108RS, CL74FC, CL91RS, DISABLE.",
				Computed:            true,
			},
			"management_access": schema.BoolAttribute{
				MarkdownDescription: "Whether to enable Management Access.",
				Computed:            true,
			},
			"management_access_network_objects": schema.SetNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "ID of the network object (Host, Network or Range).",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Type of the object",
							Computed:            true,
						},
					},
				},
			},
			"active_mac_address": schema.StringAttribute{
				MarkdownDescription: "MAC address for active interface in format 0123.4567.89ab.",
				Computed:            true,
			},
			"standby_mac_address": schema.StringAttribute{
				MarkdownDescription: "MAC address for standby interface in format 0123.4567.89ab.",
				Computed:            true,
			},
			"arp_table_entries": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"mac_address": schema.StringAttribute{
							MarkdownDescription: "MAC address for custom ARP entry in format 0123.4567.89ab.",
							Computed:            true,
						},
						"ip_address": schema.StringAttribute{
							MarkdownDescription: "IP address for custom ARP entry",
							Computed:            true,
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: "Enable Alias for custom ARP entry",
							Computed:            true,
						},
					},
				},
			},
			"enable_anti_spoofing": schema.BoolAttribute{
				MarkdownDescription: "Enable Anti Spoofing",
				Computed:            true,
			},
			"allow_full_fragment_reassembly": schema.BoolAttribute{
				MarkdownDescription: "Allow Full Fragment Reassembly",
				Computed:            true,
			},
			"override_default_fragment_setting_chain": schema.Int64Attribute{
				MarkdownDescription: "Override Default Fragment Setting - Chain value",
				Computed:            true,
			},
			"override_default_fragment_setting_size": schema.Int64Attribute{
				MarkdownDescription: "Override Default Fragment Setting - Fragment Size value",
				Computed:            true,
			},
			"override_default_fragment_setting_timeout": schema.Int64Attribute{
				MarkdownDescription: "Override Default Fragment Setting - Time Out value",
				Computed:            true,
			},
		},
	}
}
func (d *DeviceEtherChannelInterfaceDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *DeviceEtherChannelInterfaceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

func (d *DeviceEtherChannelInterfaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DeviceEtherChannelInterface

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !config.Domain.IsNull() && config.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(config.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))
	if config.Id.IsNull() && !config.Name.IsNull() {
		offset := 0
		limit := 1000
		for page := 1; ; page++ {
			queryString := fmt.Sprintf("?limit=%d&offset=%d&expanded=true", limit, offset)
			res, err := d.client.Get(config.getPath()+queryString, reqMods...)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects, got error: %s", err))
				return
			}
			if value := res.Get("items"); len(value.Array()) > 0 {
				value.ForEach(func(k, v gjson.Result) bool {
					if config.Name.ValueString() == v.Get("name").String() {
						config.Id = types.StringValue(v.Get("id").String())
						tflog.Debug(ctx, fmt.Sprintf("%s: Found object with name '%v', id: %v", config.Id.String(), config.Name.ValueString(), config.Id.String()))
						return false
					}
					return true
				})
			}
			if !config.Id.IsNull() || !res.Get("paging.next.0").Exists() {
				break
			}
			offset += limit
		}

		if config.Id.IsNull() {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with name: %s", config.Name.ValueString()))
			return
		}
	}
	urlPath := config.getPath() + "/" + url.QueryEscape(config.Id.ValueString())
	res, err := d.client.Get(urlPath, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	is_multi_instance, diags := FMCIsDeviceMultiInstance(ctx, d.client, config.DeviceId.ValueString(), reqMods)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	config.IsMultiInstance = types.BoolValue(is_multi_instance)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
