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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &FTDPlatformSettingsSyslogLoggingSetupDataSource{}
	_ datasource.DataSourceWithConfigure = &FTDPlatformSettingsSyslogLoggingSetupDataSource{}
)

func NewFTDPlatformSettingsSyslogLoggingSetupDataSource() datasource.DataSource {
	return &FTDPlatformSettingsSyslogLoggingSetupDataSource{}
}

type FTDPlatformSettingsSyslogLoggingSetupDataSource struct {
	client *fmc.Client
}

func (d *FTDPlatformSettingsSyslogLoggingSetupDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ftd_platform_settings_syslog_logging_setup"
}

func (d *FTDPlatformSettingsSyslogLoggingSetupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the FTD Platform Settings Syslog Logging Setup.").AddMinimumVersionHeaderDescription().AddMinimumVersionDescription("7.7").String,

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
				MarkdownDescription: "Type of the object; this value is always 'SyslogSetting'.",
				Computed:            true,
			},
			"enable_logging": schema.BoolAttribute{
				MarkdownDescription: "Turns on the data plane system logging.",
				Computed:            true,
			},
			"enable_logging_on_failover_standby_unit": schema.BoolAttribute{
				MarkdownDescription: "Turns on logging for the failover standby unit, if available.",
				Computed:            true,
			},
			"emblem_format": schema.BoolAttribute{
				MarkdownDescription: "Enables EMBLEM format logging.",
				Computed:            true,
			},
			"send_debug_messages_as_syslog": schema.BoolAttribute{
				MarkdownDescription: "Redirects all the debug trace output to the syslog.",
				Computed:            true,
			},
			"internal_buffer_memory_size": schema.Int64Attribute{
				MarkdownDescription: "Size of the internal buffer to which syslog messages are saved if the logging buffer is enabled.",
				Computed:            true,
			},
			"fmc_logging_type": schema.StringAttribute{
				MarkdownDescription: "FMC logging mode.",
				Computed:            true,
			},
			"fmc_logging_level": schema.StringAttribute{
				MarkdownDescription: "FMC logging level. Required if `fmc_logging_type` is not `OFF``.",
				Computed:            true,
			},
			"ftp_server_host_id": schema.StringAttribute{
				MarkdownDescription: "Id of host object representing the FTP server.",
				Computed:            true,
			},
			"ftp_server_username": schema.StringAttribute{
				MarkdownDescription: "User name to log in to the FTP server.",
				Computed:            true,
			},
			"ftp_server_password": schema.StringAttribute{
				MarkdownDescription: "Password to log in to the FTP server.",
				Computed:            true,
				Sensitive:           true,
			},
			"ftp_server_path": schema.StringAttribute{
				MarkdownDescription: "Path on the FTP server to which the logs are uploaded.",
				Computed:            true,
			},
			"ftp_server_interface_groups": schema.ListNestedAttribute{
				MarkdownDescription: "Interface Groups through which the FTP server is reachable.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Id of the Interface Group.",
							Computed:            true,
						},
					},
				},
			},
			"flash": schema.BoolAttribute{
				MarkdownDescription: "Save buffer contents to the flash memory before it is overwritten.",
				Computed:            true,
			},
			"flash_maximum_space": schema.Int64Attribute{
				MarkdownDescription: "Maximum space to be used in the flash memory for logging (in kilobytes).",
				Computed:            true,
			},
			"flash_minimum_free_space": schema.Int64Attribute{
				MarkdownDescription: "Minimum free space to be preserved in flash memory (in kilobytes).",
				Computed:            true,
			},
		},
	}
}

func (d *FTDPlatformSettingsSyslogLoggingSetupDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *FTDPlatformSettingsSyslogLoggingSetupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// Get FMC version
	fmcVersion, _ := version.NewVersion(strings.Split(d.client.FMCVersion, " ")[0])

	// Check if FMC client is connected to supports this object
	if fmcVersion.LessThan(minFMCVersionFTDPlatformSettingsSyslogLoggingSetup) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support FTD Platform Settings Syslog Logging Setup, minimum required version is 7.7", d.client.FMCVersion))
		return
	}
	var config FTDPlatformSettingsSyslogLoggingSetup

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
