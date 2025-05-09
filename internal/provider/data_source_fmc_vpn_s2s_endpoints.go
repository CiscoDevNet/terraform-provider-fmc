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

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &VPNS2SEndpointsDataSource{}
	_ datasource.DataSourceWithConfigure = &VPNS2SEndpointsDataSource{}
)

func NewVPNS2SEndpointsDataSource() datasource.DataSource {
	return &VPNS2SEndpointsDataSource{}
}

type VPNS2SEndpointsDataSource struct {
	client *fmc.Client
}

func (d *VPNS2SEndpointsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_s2s_endpoints"
}

func (d *VPNS2SEndpointsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the VPN S2S Endpoints.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the object",
				Computed:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "Name of the FMC domain",
				Optional:            true,
			},
			"vpn_s2s_id": schema.StringAttribute{
				MarkdownDescription: "Id of the parent VPN S2S Topology.",
				Required:            true,
			},
			"items": schema.MapNestedAttribute{
				MarkdownDescription: "Map of Endpoints. The key of the map is the name of the Endpoint.",
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Id of the Endpoint object.",
							Computed:            true,
						},
						"peer_type": schema.StringAttribute{
							MarkdownDescription: "Role of the device in the topology.",
							Computed:            true,
						},
						"extranet_device": schema.BoolAttribute{
							MarkdownDescription: "Is the device managed by local FMC.",
							Computed:            true,
						},
						"extranet_ip_address": schema.StringAttribute{
							MarkdownDescription: "IP address of extranet device (optionally coma separated Backup IP Addresses).",
							Computed:            true,
						},
						"extranet_dynamic_ip": schema.BoolAttribute{
							MarkdownDescription: "Is the IP address of the extranet device dynamic.",
							Computed:            true,
						},
						"device_id": schema.StringAttribute{
							MarkdownDescription: "Id of the device managed by local FMC.",
							Computed:            true,
						},
						"interface_id": schema.StringAttribute{
							MarkdownDescription: "Id of the primary interface to source the VPN connection.",
							Computed:            true,
						},
						"interface_ipv6_address": schema.StringAttribute{
							MarkdownDescription: "IPv6 address of the interface. If not set, IPv4 address will be used.",
							Computed:            true,
						},
						"interface_public_ip_address": schema.StringAttribute{
							MarkdownDescription: "Public address of the interface, in case the one configured on the interface is private.",
							Computed:            true,
						},
						"connection_type": schema.StringAttribute{
							MarkdownDescription: "Connection type.",
							Computed:            true,
						},
						"allow_incoming_ikev2_routes": schema.BoolAttribute{
							MarkdownDescription: "Allow incoming IKEv2 routes.",
							Computed:            true,
						},
						"send_vti_ip_to_peer": schema.BoolAttribute{
							MarkdownDescription: "Send Virtual Tunnel Interface IP to the peers",
							Computed:            true,
						},
						"protected_networks": schema.SetNestedAttribute{
							MarkdownDescription: "Set of protected networks.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the protected network.",
										Computed:            true,
									},
								},
							},
						},
						"protected_networks_acl_id": schema.StringAttribute{
							MarkdownDescription: "Id of the ACL that defines protected networks.",
							Computed:            true,
						},
						"nat_traversal": schema.BoolAttribute{
							MarkdownDescription: "Enable NAT traversal.",
							Computed:            true,
						},
						"nat_exemption": schema.BoolAttribute{
							MarkdownDescription: "Enable NAT exemption.",
							Computed:            true,
						},
						"nat_exemption_inside_interface_id": schema.StringAttribute{
							MarkdownDescription: "Id of the inside Security Zone for NAT Exemption identification.",
							Computed:            true,
						},
						"reverse_route_injection": schema.BoolAttribute{
							MarkdownDescription: "Enable Reverse Route Injection (RRI).",
							Computed:            true,
						},
						"local_identity_type": schema.StringAttribute{
							MarkdownDescription: "Type of the local identity.",
							Computed:            true,
						},
						"local_identity_string": schema.StringAttribute{
							MarkdownDescription: "String of the local identity (applicable only for types KEYID and EMAILID)",
							Computed:            true,
						},
						"vpn_filter_acl_id": schema.StringAttribute{
							MarkdownDescription: "Id of the VPN filter ACL.",
							Computed:            true,
						},
						"override_remote_vpn_filter_acl_id": schema.StringAttribute{
							MarkdownDescription: "Id of the ACL to override VPN filter on the Hub.",
							Computed:            true,
						},
						"backup_interface_id": schema.StringAttribute{
							MarkdownDescription: "Id of the backup interface to source the VPN connection.",
							Computed:            true,
						},
						"backup_interface_public_ip_address": schema.StringAttribute{
							MarkdownDescription: "Public address of the backup VIT interface, in case the one configured on the interface is private. (NAT Address)",
							Computed:            true,
						},
						"backup_local_identity_type": schema.StringAttribute{
							MarkdownDescription: "Type of the local identity for the backup tunnel.",
							Computed:            true,
						},
						"backup_local_identity_string": schema.StringAttribute{
							MarkdownDescription: "String of the local identity for the backup tunnel (applicable only for types KEYID and EMAILID)",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *VPNS2SEndpointsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *VPNS2SEndpointsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config VPNS2SEndpoints

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

	// Get all objects from FMC
	urlPath := config.getPath() + "?expanded=true"
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
