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
				MarkdownDescription: "Type of the object; this value is always ''.",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the object.",
				Computed:            true,
			},
			"enable_ssl_protocol": schema.BoolAttribute{
				MarkdownDescription: "Whether the SSL protocol is enabled.",
				Computed:            true,
			},
			"enable_ipsec_ikev2_protocol": schema.BoolAttribute{
				MarkdownDescription: "Whether the IPsec IKEv2 protocol is enabled.",
				Computed:            true,
			},
			"ipv4_address_pools": schema.ListNestedAttribute{
				MarkdownDescription: "List of IPv4 address pools for the Group Policy.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Unique identifier for the IPv4 address pool.",
							Computed:            true,
						},
					},
				},
			},
			"banner": schema.StringAttribute{
				MarkdownDescription: "Banner text to be displayed to users.",
				Computed:            true,
			},
			"primary_dns_server": schema.StringAttribute{
				MarkdownDescription: "Primary DNS server for the Group Policy.",
				Computed:            true,
			},
			"secondary_dns_server": schema.StringAttribute{
				MarkdownDescription: "Secondary DNS server for the Group Policy.",
				Computed:            true,
			},
			"primary_wins_server": schema.StringAttribute{
				MarkdownDescription: "Primary WINS server for the Group Policy.",
				Computed:            true,
			},
			"secondary_wins_server": schema.StringAttribute{
				MarkdownDescription: "Secondary WINS server for the Group Policy.",
				Computed:            true,
			},
			"dhcp_scope_network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID for the DHCP scope.",
				Computed:            true,
			},
			"default_domain": schema.StringAttribute{
				MarkdownDescription: "Default domain name for the Group Policy.",
				Computed:            true,
			},
			"ipv4_split_tunnel_policy": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"ipv6_split_tunnel_policy": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"split_tunnel_acl_id": schema.StringAttribute{
				MarkdownDescription: "ACL ID for the split tunnel configuration.",
				Computed:            true,
			},
			"split_d_n_s_request_policy": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"split_d_n_s_domain_list": schema.StringAttribute{
				MarkdownDescription: "List of domains for split DNS requests.",
				Computed:            true,
			},
			"secure_client_client_profile_id": schema.StringAttribute{
				MarkdownDescription: "ID of the Secure Client profile.",
				Computed:            true,
			},
			"secure_client_management_profile_id": schema.StringAttribute{
				MarkdownDescription: "ID of the Secure Client management profile.",
				Computed:            true,
			},
			"ssl_compression": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"dtls_compression": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"mtu_size": schema.Int64Attribute{
				MarkdownDescription: "Maximum Transmission Unit size for SSL connections.",
				Computed:            true,
			},
			"ignore_df_bit": schema.BoolAttribute{
				MarkdownDescription: "Whether to ignore the Don't Fragment bit in packets.",
				Computed:            true,
			},
			"enable_keep_alive_messages": schema.BoolAttribute{
				MarkdownDescription: "Whether to enable keep-alive messages for the connection.",
				Computed:            true,
			},
			"keep_alive_message_interval": schema.Int64Attribute{
				MarkdownDescription: "Interval for keep-alive messages in seconds.",
				Computed:            true,
			},
			"enable_gateway_dpd": schema.BoolAttribute{
				MarkdownDescription: "Whether to enable Dead Peer Detection (DPD) for the connection.",
				Computed:            true,
			},
			"gateway_dpd_interval": schema.Int64Attribute{
				MarkdownDescription: "Interval for Dead Peer Detection (DPD) messages in seconds.",
				Computed:            true,
			},
			"enable_client_dpd": schema.BoolAttribute{
				MarkdownDescription: "Whether to enable Dead Peer Detection (DPD) for the connection.",
				Computed:            true,
			},
			"client_dpd_interval": schema.Int64Attribute{
				MarkdownDescription: "Interval for Dead Peer Detection (DPD) messages in seconds.",
				Computed:            true,
			},
			"client_bypass_protocol": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_ssl_rekey": schema.BoolAttribute{
				MarkdownDescription: "Whether to enable SSL rekeying.",
				Computed:            true,
			},
			"rekey_method": schema.StringAttribute{
				MarkdownDescription: "Method to use for SSL rekeying.",
				Computed:            true,
			},
			"rekey_interval": schema.Int64Attribute{
				MarkdownDescription: "Interval for SSL rekeying in minutes.",
				Computed:            true,
			},
			"client_firewall_private_network_rules_id": schema.StringAttribute{
				MarkdownDescription: "ID of the client firewall private network rules.",
				Computed:            true,
			},
			"client_firewall_public_network_rules_id": schema.StringAttribute{
				MarkdownDescription: "ID of the client firewall public network rules.",
				Computed:            true,
			},
			"custom_attributes": schema.ListNestedAttribute{
				MarkdownDescription: "List of custom attributes for the Group Policy.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "ID of the custom attribute.",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "AnyConnect specific custom attribute.",
							Computed:            true,
						},
					},
				},
			},
			"traffic_filter_acl_id": schema.StringAttribute{
				MarkdownDescription: "ACL ID for the traffic filter.",
				Computed:            true,
			},
			"restrict_vpn_to_vlan_id": schema.Int64Attribute{
				MarkdownDescription: "VLAN ID to restrict VPN access.",
				Computed:            true,
			},
			"access_hours_id": schema.StringAttribute{
				MarkdownDescription: "ID of the access hours settings.",
				Computed:            true,
			},
			"simultaneous_login_per_user": schema.Int64Attribute{
				MarkdownDescription: "Maximum number of simultaneous logins allowed per user.",
				Computed:            true,
			},
			"max_connection_time": schema.Int64Attribute{
				MarkdownDescription: "Maximum connection timeout in minutes.",
				Computed:            true,
			},
			"max_connection_time_alert_interval": schema.Int64Attribute{
				MarkdownDescription: "Alert interval for maximum connection time in minutes.",
				Computed:            true,
			},
			"vpn_idle_timeout": schema.Int64Attribute{
				MarkdownDescription: "VPN idle timeout in minutes.",
				Computed:            true,
			},
			"vpn_idle_timeout_alert_interval": schema.Int64Attribute{
				MarkdownDescription: "Alert interval for VPN idle timeout in minutes.",
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
