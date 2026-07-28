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
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &DeviceDHCPRelayDataSource{}
	_ datasource.DataSourceWithConfigure = &DeviceDHCPRelayDataSource{}
)

func NewDeviceDHCPRelayDataSource() datasource.DataSource {
	return &DeviceDHCPRelayDataSource{}
}

type DeviceDHCPRelayDataSource struct {
	client *fmc.Client
}

func (d *DeviceDHCPRelayDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_dhcp_relay"
}

func (d *DeviceDHCPRelayDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the Device DHCP Relay.").AddMinimumVersionHeaderDescription().AddMinimumVersionDescription("7.4").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the object",
				Required:            true,
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
				MarkdownDescription: "Type of the object; this value is always 'DHCPRelaySettings'.",
				Computed:            true,
			},
			"ipv4_relay_timeout": schema.Int64Attribute{
				MarkdownDescription: "Amount of time in seconds that the Firewall Threat Defense device waits to time out the DHCP relay agent.",
				Computed:            true,
			},
			"ipv6_relay_timeout": schema.Int64Attribute{
				MarkdownDescription: "Amount of time in seconds that the Firewall Threat Defense device waits to time out the DHCPv6 relay agent.",
				Computed:            true,
			},
			"trust_all_information": schema.BoolAttribute{
				MarkdownDescription: "Set all client interfaces as trusted.",
				Computed:            true,
			},
			"relay_agents": schema.ListNestedAttribute{
				MarkdownDescription: "List of per-interface DHCP relay agents.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_id": schema.StringAttribute{
							MarkdownDescription: "Id of the interface connected to DHCP clients.",
							Computed:            true,
						},
						"interface_name": schema.StringAttribute{
							MarkdownDescription: "Name of the interface connected to DHCP clients.",
							Computed:            true,
						},
						"interface_type": schema.StringAttribute{
							MarkdownDescription: "Type of the interface connected to DHCP clients.",
							Computed:            true,
						},
						"ipv4_relay": schema.BoolAttribute{
							MarkdownDescription: "Enable the IPv4 DHCP relay agent on the interface.",
							Computed:            true,
						},
						"ipv6_relay": schema.BoolAttribute{
							MarkdownDescription: "Enable the IPv6 DHCP relay agent on the interface.",
							Computed:            true,
						},
						"set_route": schema.BoolAttribute{
							MarkdownDescription: "(IPv4 only) Changes the default gateway address in the DHCP message from the server to that of the Firewall Threat Defense device interface that is closest to the DHCP client, which relayed the original DHCP request.",
							Computed:            true,
						},
					},
				},
			},
			"servers": schema.ListNestedAttribute{
				MarkdownDescription: "List of DHCP servers that client requests are relayed to.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"server_id": schema.StringAttribute{
							MarkdownDescription: "Id of the Host identifying the DHCP server.",
							Computed:            true,
						},
						"server_interface_id": schema.StringAttribute{
							MarkdownDescription: "Id of the server-side interface used to reach the DHCP server. Id, name, and type must all be specified together. Mutually exclusive with `client_interfaces`.",
							Computed:            true,
						},
						"server_interface_name": schema.StringAttribute{
							MarkdownDescription: "Name of the server-side interface used to reach the DHCP server. Id, name, and type must all be specified together. Mutually exclusive with `client_interfaces`.",
							Computed:            true,
						},
						"server_interface_type": schema.StringAttribute{
							MarkdownDescription: "Type of the server-side interface used to reach the DHCP server. Id, name, and type must all be specified together. Mutually exclusive with `client_interfaces`.",
							Computed:            true,
						},
						"client_interfaces": schema.ListNestedAttribute{
							MarkdownDescription: "List of client-side interfaces served by this DHCP relay server. Mutually exclusive with `server_interface_id`.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the client-side interface.",
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: "Name of the client-side interface.",
										Computed:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of the client-side interface.",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *DeviceDHCPRelayDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *DeviceDHCPRelayDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	// Check if FMC client is connected to supports this object
	if d.client.FMCVersionParsed.LessThan(minFMCVersionDeviceDHCPRelay) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support Device DHCP Relay, minimum required version is 7.4", d.client.FMCVersion))
		return
	}
	var config DeviceDHCPRelay

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
