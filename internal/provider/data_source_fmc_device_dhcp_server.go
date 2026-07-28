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
	_ datasource.DataSource              = &DeviceDHCPServerDataSource{}
	_ datasource.DataSourceWithConfigure = &DeviceDHCPServerDataSource{}
)

func NewDeviceDHCPServerDataSource() datasource.DataSource {
	return &DeviceDHCPServerDataSource{}
}

type DeviceDHCPServerDataSource struct {
	client *fmc.Client
}

func (d *DeviceDHCPServerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_dhcp_server"
}

func (d *DeviceDHCPServerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the Device DHCP Server.").AddMinimumVersionHeaderDescription().AddMinimumVersionDescription("7.4").String,

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
				MarkdownDescription: "Type of the object; this value is always 'DHCPServerSettings'.",
				Computed:            true,
			},
			"ping_timeout": schema.Int64Attribute{
				MarkdownDescription: "Ping timeout in milliseconds.",
				Computed:            true,
			},
			"lease_length": schema.Int64Attribute{
				MarkdownDescription: "Duration (in seconds) that a client can use an assigned IP address before it must renew the lease.",
				Computed:            true,
			},
			"auto_config_interface_id": schema.StringAttribute{
				MarkdownDescription: "Id of the interface to be used for autoconfiguration. Id, name, and type must all be specified together.",
				Computed:            true,
			},
			"auto_config_interface_type": schema.StringAttribute{
				MarkdownDescription: "Type of the interface to be used for autoconfiguration. Id, name, and type must all be specified together.",
				Computed:            true,
			},
			"auto_config_interface_name": schema.StringAttribute{
				MarkdownDescription: "Name of the interface to be used for autoconfiguration. Id, name, and type must all be specified together.",
				Computed:            true,
			},
			"domain_name": schema.StringAttribute{
				MarkdownDescription: "Domain name given to DHCP clients.",
				Computed:            true,
			},
			"primary_dns_server_id": schema.StringAttribute{
				MarkdownDescription: "Id of the Host object used as the primary DNS server offered to DHCP clients.",
				Computed:            true,
			},
			"secondary_dns_server_id": schema.StringAttribute{
				MarkdownDescription: "Id of the Host object used as the secondary DNS server offered to DHCP clients.",
				Computed:            true,
			},
			"primary_wins_server_id": schema.StringAttribute{
				MarkdownDescription: "Id of the Host object used as the primary WINS server offered to DHCP clients.",
				Computed:            true,
			},
			"secondary_wins_server_id": schema.StringAttribute{
				MarkdownDescription: "Id of the Host object used as the secondary WINS server offered to DHCP clients.",
				Computed:            true,
			},
			"servers": schema.ListNestedAttribute{
				MarkdownDescription: "List of per-interface DHCP server pools.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_id": schema.StringAttribute{
							MarkdownDescription: "Id of the interface on which the DHCP server is enabled.",
							Computed:            true,
						},
						"interface_name": schema.StringAttribute{
							MarkdownDescription: "Name of the interface on which the DHCP server is enabled.",
							Computed:            true,
						},
						"interface_type": schema.StringAttribute{
							MarkdownDescription: "Type of the interface on which the DHCP server is enabled.",
							Computed:            true,
						},
						"address_pool": schema.StringAttribute{
							MarkdownDescription: "Range of IP addresses (formatted as `start_ip-end_ip`) the DHCP server can hand out. The pool must be on the same subnet as the associated interface.",
							Computed:            true,
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: "Enable the DHCP server on the associated interface.",
							Computed:            true,
						},
					},
				},
			},
			"dhcp_options": schema.ListNestedAttribute{
				MarkdownDescription: "List of advanced DHCP option codes (parameters) sent to DHCP clients.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"code": schema.Int64Attribute{
							MarkdownDescription: "DHCP option code.",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Type of the value carried by the option. `IP` uses `first_ip_address_id` (and optionally `second_ip_address_id`), `ASCII` uses `ascii` and `HEX` uses `hex`.",
							Computed:            true,
						},
						"first_ip_address_id": schema.StringAttribute{
							MarkdownDescription: "Id of the Host object used as the first IP address, when `type` is `IP`.",
							Computed:            true,
						},
						"second_ip_address_id": schema.StringAttribute{
							MarkdownDescription: "Id of the Host object used as the second IP address, when `type` is `IP`.",
							Computed:            true,
						},
						"ascii": schema.StringAttribute{
							MarkdownDescription: "ASCII value, when `type` is `ASCII`.",
							Computed:            true,
						},
						"hex": schema.StringAttribute{
							MarkdownDescription: "Hexadecimal value, when `type` is `HEX`.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *DeviceDHCPServerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *DeviceDHCPServerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	// Check if FMC client is connected to supports this object
	if d.client.FMCVersionParsed.LessThan(minFMCVersionDeviceDHCPServer) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support Device DHCP Server, minimum required version is 7.4", d.client.FMCVersion))
		return
	}
	var config DeviceDHCPServer

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
