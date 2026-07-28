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
	_ datasource.DataSource              = &DeviceDDNSDataSource{}
	_ datasource.DataSourceWithConfigure = &DeviceDDNSDataSource{}
)

func NewDeviceDDNSDataSource() datasource.DataSource {
	return &DeviceDDNSDataSource{}
}

type DeviceDDNSDataSource struct {
	client *fmc.Client
}

func (d *DeviceDDNSDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_ddns"
}

func (d *DeviceDDNSDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the Device DDNS.").AddMinimumVersionHeaderDescription().AddMinimumVersionDescription("7.4").String,

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
				MarkdownDescription: "Type of the object; this value is always 'DDNSSetting'.",
				Computed:            true,
			},
			"dhcp_client_request_type": schema.StringAttribute{
				MarkdownDescription: "Determine which records you want the DHCP server to update.",
				Computed:            true,
			},
			"dhcp_client_broadcast": schema.BoolAttribute{
				MarkdownDescription: "Request that the DHCP server broadcast the DHCP reply (DHCP option 1).",
				Computed:            true,
			},
			"dynamic_dns_update": schema.StringAttribute{
				MarkdownDescription: "Which DNS RRs you want the DHCP server to update.",
				Computed:            true,
			},
			"dhcp_client_request_override": schema.BoolAttribute{
				MarkdownDescription: "Override the update actions requested by the DHCP client.",
				Computed:            true,
			},
			"dhcp_client_id_interfaces": schema.ListNestedAttribute{
				MarkdownDescription: "List of interfaces on which the DHCP client sets the client identifier.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Id of the interface.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Name of the interface.",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Type of the interface.",
							Computed:            true,
						},
					},
				},
			},
			"ddns_update_methods": schema.ListNestedAttribute{
				MarkdownDescription: "List of DDNS update methods.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Name of the DDNS update method.",
							Computed:            true,
						},
						"method": schema.StringAttribute{
							MarkdownDescription: "DDNS update method type.",
							Computed:            true,
						},
						"update_interval_day": schema.Int64Attribute{
							MarkdownDescription: "Update interval day portion.",
							Computed:            true,
						},
						"update_interval_hour": schema.Int64Attribute{
							MarkdownDescription: "Update interval hour portion.",
							Computed:            true,
						},
						"update_interval_minute": schema.Int64Attribute{
							MarkdownDescription: "Update interval minute portion.",
							Computed:            true,
						},
						"update_interval_second": schema.Int64Attribute{
							MarkdownDescription: "Update interval second portion.",
							Computed:            true,
						},
						"web_url": schema.StringAttribute{
							MarkdownDescription: "URL of the web update server, used when `method` is `WEB`. Syntax https://username:password@provider-domain/xyz?hostname=<h>&myip=<a>.",
							Computed:            true,
						},
						"web_update_type": schema.StringAttribute{
							MarkdownDescription: "Addresses to update when `method` is `WEB`.",
							Computed:            true,
						},
						"update_records": schema.StringAttribute{
							MarkdownDescription: "DNS records to update.",
							Computed:            true,
						},
					},
				},
			},
			"ddns_interface_settings": schema.ListNestedAttribute{
				MarkdownDescription: "List of per-interface DDNS settings.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_id": schema.StringAttribute{
							MarkdownDescription: "Id of the interface. Id, name, and type must all be specified together.",
							Computed:            true,
						},
						"interface_name": schema.StringAttribute{
							MarkdownDescription: "Name of the interface. Id, name, and type must all be specified together.",
							Computed:            true,
						},
						"interface_type": schema.StringAttribute{
							MarkdownDescription: "Type of the interface. Id, name, and type must all be specified together.",
							Computed:            true,
						},
						"method_name": schema.StringAttribute{
							MarkdownDescription: "Name of the DDNS update method (from `update_methods`) applied to the interface.",
							Computed:            true,
						},
						"hostname": schema.StringAttribute{
							MarkdownDescription: "Set hostname for the interface.",
							Computed:            true,
						},
						"dhcp_client_request_type": schema.StringAttribute{
							MarkdownDescription: "How the DHCP client requests the DHCP server to perform DNS updates.",
							Computed:            true,
						},
						"dynamic_dns_update": schema.StringAttribute{
							MarkdownDescription: "DNS records that the DDNS method updates.",
							Computed:            true,
						},
						"dhcp_client_request_override": schema.BoolAttribute{
							MarkdownDescription: "Override the DDNS updates that the DHCP client requests the DHCP server to perform.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *DeviceDDNSDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *DeviceDDNSDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	// Check if FMC client is connected to supports this object
	if d.client.FMCVersionParsed.LessThan(minFMCVersionDeviceDDNS) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support Device DDNS, minimum required version is 7.4", d.client.FMCVersion))
		return
	}
	var config DeviceDDNS

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
