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

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &FTDPlatformSettingsSyslogServersDataSource{}
	_ datasource.DataSourceWithConfigure = &FTDPlatformSettingsSyslogServersDataSource{}
)

func NewFTDPlatformSettingsSyslogServersDataSource() datasource.DataSource {
	return &FTDPlatformSettingsSyslogServersDataSource{}
}

type FTDPlatformSettingsSyslogServersDataSource struct {
	client *fmc.Client
}

func (d *FTDPlatformSettingsSyslogServersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ftd_platform_settings_syslog_servers"
}

func (d *FTDPlatformSettingsSyslogServersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the FTD Platform Settings Syslog Servers.").AddMinimumVersionHeaderDescription().AddMinimumVersionDescription("7.7").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the object",
				Required:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "Name of the FMC domain",
				Optional:            true,
			},
			"ftd_platform_settings_id": schema.StringAttribute{
				MarkdownDescription: "Id of the parent FTD Platform Settings.",
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the object; this value is always 'Server'.",
				Computed:            true,
			},
			"allow_user_traffic_when_tcp_syslog_server_is_down": schema.BoolAttribute{
				MarkdownDescription: "Allow user traffic when TCP syslog server is down.",
				Computed:            true,
			},
			"message_queue_size": schema.Int64Attribute{
				MarkdownDescription: "Size of the queue for storing syslog messages on the security appliance when syslog server is busy. Specify 0 to allow an unlimited number of messages to be queued (subject to available block memory).",
				Computed:            true,
			},
			"servers": schema.ListNestedAttribute{
				MarkdownDescription: "List of syslog servers.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_object_id": schema.StringAttribute{
							MarkdownDescription: "Id of host object representing the IP address of the syslog server.",
							Computed:            true,
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: "Protocol used to send syslog messages to the server.",
							Computed:            true,
						},
						"port": schema.Int64Attribute{
							MarkdownDescription: "Port number used to send syslog messages to the server.",
							Computed:            true,
						},
						"emblem_format": schema.BoolAttribute{
							MarkdownDescription: "(UDP only) Log messages in EMBLEM format.",
							Computed:            true,
						},
						"secure_syslog": schema.BoolAttribute{
							MarkdownDescription: "(TCP only) Use TLS to secure syslog messages sent to the server.",
							Computed:            true,
						},
						"use_management_interface": schema.BoolAttribute{
							MarkdownDescription: "Use management interface to reach syslog server (true) or use data interfaces (false).",
							Computed:            true,
						},
						"interface_literals": schema.SetAttribute{
							MarkdownDescription: "List of interface literals to reach syslog server.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"interface_objects": schema.SetNestedAttribute{
							MarkdownDescription: "List of interface objects (Security Zones or Interface Groups) to reach syslog server.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the interface object.",
										Computed:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of the interface object; either 'SecurityZone' or 'InterfaceGroup'.",
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: "Name of the interface object.",
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

func (d *FTDPlatformSettingsSyslogServersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *FTDPlatformSettingsSyslogServersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// Get FMC version
	fmcVersion, _ := version.NewVersion(strings.Split(d.client.FMCVersion, " ")[0])

	// Check if FMC client is connected to supports this object
	if fmcVersion.LessThan(minFMCVersionFTDPlatformSettingsSyslogServers) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support FTD Platform Settings Syslog Servers, minimum required version is 7.7", d.client.FMCVersion))
		return
	}
	var config FTDPlatformSettingsSyslogServers

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
