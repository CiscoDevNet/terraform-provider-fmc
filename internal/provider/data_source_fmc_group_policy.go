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
	_ datasource.DataSource              = &GroupPolicyDataSource{}
	_ datasource.DataSourceWithConfigure = &GroupPolicyDataSource{}
)

func NewGroupPolicyDataSource() datasource.DataSource {
	return &GroupPolicyDataSource{}
}

type GroupPolicyDataSource struct {
	client *fmc.Client
}

func (d *GroupPolicyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_group_policy"
}

func (d *GroupPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the Group Policy.").String,

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
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the Group Policy object.",
				Optional:            true,
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the object; this value is always 'GroupPolicy'.",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the object.",
				Computed:            true,
			},
			"protocol_ssl": schema.BoolAttribute{
				MarkdownDescription: "Enable SSL protocol for VPN connections.",
				Computed:            true,
			},
			"protocol_ipsec_ikev2": schema.BoolAttribute{
				MarkdownDescription: "Enable IPsec IKEv2 protocol for VPN connections.",
				Computed:            true,
			},
			"ipv4_address_pools": schema.ListNestedAttribute{
				MarkdownDescription: "List of IPv4 Address Pools for address assignment.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Pool Id.",
							Computed:            true,
						},
					},
				},
			},
			"banner": schema.StringAttribute{
				MarkdownDescription: "Banner text to be displayed to users.",
				Computed:            true,
			},
			"primary_dns_server_host_id": schema.StringAttribute{
				MarkdownDescription: "Id of host object that represents primary DNS server.",
				Computed:            true,
			},
			"secondary_dns_server_host_id": schema.StringAttribute{
				MarkdownDescription: "Id of host object that represents secondary DNS server.",
				Computed:            true,
			},
			"primary_wins_server_host_id": schema.StringAttribute{
				MarkdownDescription: "Id of host object that represents primary WINS server.",
				Computed:            true,
			},
			"secondary_wins_server_host_id": schema.StringAttribute{
				MarkdownDescription: "Id of host object that represents secondary WINS server.",
				Computed:            true,
			},
			"dhcp_network_scope_network_object_id": schema.StringAttribute{
				MarkdownDescription: "Id of the Network Object used to determine the DHCP scope.",
				Computed:            true,
			},
			"default_domain": schema.StringAttribute{
				MarkdownDescription: "Name of the default domain.",
				Computed:            true,
			},
			"ipv4_split_tunnel_policy": schema.StringAttribute{
				MarkdownDescription: "IPv4 split tunnel policy.",
				Computed:            true,
			},
			"ipv6_split_tunnel_policy": schema.StringAttribute{
				MarkdownDescription: "IPv6 split tunnel policy.",
				Computed:            true,
			},
			"split_tunnel_acl_id": schema.StringAttribute{
				MarkdownDescription: "Id of standard or extended ACL used for split tunnel configuration.",
				Computed:            true,
			},
			"split_tunnel_acl_type": schema.StringAttribute{
				MarkdownDescription: "Type of ACL used for split tunnel configuration. Mandatory, when `split_tunnel_acl_id` is set.",
				Computed:            true,
			},
			"dns_request_split_tunnel_policy": schema.StringAttribute{
				MarkdownDescription: "Define if DNS requests should be send over the tunnel or not.",
				Computed:            true,
			},
			"split_dns_domain_list": schema.StringAttribute{
				MarkdownDescription: "Up to 10, comma separated domains for split DNS requests.",
				Computed:            true,
			},
			"secure_client_profile_id": schema.StringAttribute{
				MarkdownDescription: "ID of the Secure Client Profile.",
				Computed:            true,
			},
			"secure_client_management_profile_id": schema.StringAttribute{
				MarkdownDescription: "ID of the Secure Client Management Profile.",
				Computed:            true,
			},
			"secure_client_modules": schema.ListNestedAttribute{
				MarkdownDescription: "List of Secure Client Modules to be enabled.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"profile_id": schema.StringAttribute{
							MarkdownDescription: "ID of the module profile.",
							Computed:            true,
						},
						"download_module": schema.BoolAttribute{
							MarkdownDescription: "Enable module download.",
							Computed:            true,
						},
					},
				},
			},
			"ssl_compression": schema.StringAttribute{
				MarkdownDescription: "SSL compression method for the connection.",
				Computed:            true,
			},
			"dtls_compression": schema.StringAttribute{
				MarkdownDescription: "DTLS compression method for the connection.",
				Computed:            true,
			},
			"mtu_size": schema.Int64Attribute{
				MarkdownDescription: "Maximum Transmission Unit (MTU) size for SSL connections.",
				Computed:            true,
			},
			"ignore_df_bit": schema.BoolAttribute{
				MarkdownDescription: "Whether to ignore the Don't Fragment bit in packets.",
				Computed:            true,
			},
			"keep_alive_messages": schema.BoolAttribute{
				MarkdownDescription: "Enable Keepalive Messages between Secure Client and VPN gateway.",
				Computed:            true,
			},
			"keep_alive_messages_interval": schema.Int64Attribute{
				MarkdownDescription: "Keepalive message interval in seconds.",
				Computed:            true,
			},
			"gateway_dpd": schema.BoolAttribute{
				MarkdownDescription: "Enable VPN secure gateway Dead Peer Detection (DPD).",
				Computed:            true,
			},
			"gateway_dpd_interval": schema.Int64Attribute{
				MarkdownDescription: "VPN secure gateway Dead Peer Detection (DPD) messages interval in seconds.",
				Computed:            true,
			},
			"client_dpd": schema.BoolAttribute{
				MarkdownDescription: "Enable VPN client Dead Peer Detection (DPD).",
				Computed:            true,
			},
			"client_dpd_interval": schema.Int64Attribute{
				MarkdownDescription: "VPN client Dead Peer Detection (DPD) messages interval in seconds.",
				Computed:            true,
			},
			"client_bypass_protocol": schema.BoolAttribute{
				MarkdownDescription: "Drop network traffic for which the headend did not assign an IP address. Applicable if headend assigned only IPv4 or only IPv6 address.",
				Computed:            true,
			},
			"ssl_rekey": schema.BoolAttribute{
				MarkdownDescription: "Enables the client to rekey the connection.",
				Computed:            true,
			},
			"ssl_rekey_method": schema.StringAttribute{
				MarkdownDescription: "Method to use for SSL rekeying.",
				Computed:            true,
			},
			"ssl_rekey_interval": schema.Int64Attribute{
				MarkdownDescription: "Interval for SSL rekeying in minutes.",
				Computed:            true,
			},
			"client_firewall_private_network_rules_acl_id": schema.StringAttribute{
				MarkdownDescription: "Id of extended ACL to configure firewall settings for the VPN client's platform.",
				Computed:            true,
			},
			"client_firewall_public_network_rules_acl_id": schema.StringAttribute{
				MarkdownDescription: "Id of extended ACL to configure firewall settings for the VPN client's platform.",
				Computed:            true,
			},
			"secure_client_custom_attributes": schema.ListNestedAttribute{
				MarkdownDescription: "Secure Client Custom Attributes that are used by the Secure Client to configure features.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Id of the Custom Attribute.",
							Computed:            true,
						},
					},
				},
			},
			"traffic_filter_acl_id": schema.StringAttribute{
				MarkdownDescription: "Id of Extended ACL that determine whether to allow or block tunneled data packets coming through the VPN connection.",
				Computed:            true,
			},
			"restrict_vpn_to_vlan": schema.Int64Attribute{
				MarkdownDescription: "Specifies the egress VLAN ID for sessions to which this Group Policy applies.",
				Computed:            true,
			},
			"access_hours_time_range_id": schema.StringAttribute{
				MarkdownDescription: "ID of Time Range object that specifies when this group policy is available to be applied to a remote access user.",
				Computed:            true,
			},
			"simultaneous_logins_per_user": schema.Int64Attribute{
				MarkdownDescription: "Maximum number of simultaneous logins allowed for a user.",
				Computed:            true,
			},
			"maximum_connection_time": schema.Int64Attribute{
				MarkdownDescription: "Maximum user connection time in minutes.",
				Computed:            true,
			},
			"maximum_connection_time_alert_interval": schema.Int64Attribute{
				MarkdownDescription: "Specifies the interval of time before maximum connection time is reached to display a message to the user.",
				Computed:            true,
			},
			"idle_timeout": schema.Int64Attribute{
				MarkdownDescription: "VPN idle timeout in minutes.",
				Computed:            true,
			},
			"idle_timeout_alert_interval": schema.Int64Attribute{
				MarkdownDescription: "Interval of time before idle time is reached to display a message to the user.",
				Computed:            true,
			},
		},
	}
}
func (d *GroupPolicyDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *GroupPolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *GroupPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GroupPolicy

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
						tflog.Debug(ctx, fmt.Sprintf("%s: Found object with name '%v', id: %v", config.Id.ValueString(), config.Name.ValueString(), config.Id.ValueString()))
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
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with name: %v", config.Name.ValueString()))
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

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
