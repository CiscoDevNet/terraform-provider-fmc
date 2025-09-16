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
	_ datasource.DataSource              = &FTDPlatformSettingsSyslogSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &FTDPlatformSettingsSyslogSettingsDataSource{}
)

func NewFTDPlatformSettingsSyslogSettingsDataSource() datasource.DataSource {
	return &FTDPlatformSettingsSyslogSettingsDataSource{}
}

type FTDPlatformSettingsSyslogSettingsDataSource struct {
	client *fmc.Client
}

func (d *FTDPlatformSettingsSyslogSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ftd_platform_settings_syslog_settings"
}

func (d *FTDPlatformSettingsSyslogSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the FTD Platform Settings Syslog Settings.").AddMinimumVersionHeaderDescription().AddMinimumVersionDescription("7.7").String,

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
			"facility": schema.StringAttribute{
				MarkdownDescription: "System log facility for syslog servers to use as a basis to file messages.",
				Computed:            true,
			},
			"timestamp_format": schema.StringAttribute{
				MarkdownDescription: "Include timestamp to generated syslog messages in the specified format.",
				Computed:            true,
			},
			"device_id_type": schema.StringAttribute{
				MarkdownDescription: "Include device identifier to syslog messages.",
				Computed:            true,
			},
			"device_id_user_defined_id": schema.StringAttribute{
				MarkdownDescription: "User defined device identifier. This is required when `device_id_type` is set to `USERDEFINEDID`.",
				Computed:            true,
			},
			"device_id_interface_id": schema.StringAttribute{
				MarkdownDescription: "Use the IP address of the selected interface (Security Zone or Interface Group that maps to a single interface). This is required when `device_id_type` is set to `INTERFACE`.",
				Computed:            true,
			},
			"all_syslog_messages": schema.BoolAttribute{
				MarkdownDescription: "Enable all syslog messages.",
				Computed:            true,
			},
			"all_syslog_messages_logging_level": schema.StringAttribute{
				MarkdownDescription: "Logging level for all syslog messages. This is required when `all_syslog_messages` is set to `true`.",
				Computed:            true,
			},
		},
	}
}

func (d *FTDPlatformSettingsSyslogSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *FTDPlatformSettingsSyslogSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// Get FMC version
	fmcVersion, _ := version.NewVersion(strings.Split(d.client.FMCVersion, " ")[0])

	// Check if FMC client is connected to supports this object
	if fmcVersion.LessThan(minFMCVersionFTDPlatformSettingsSyslogSettings) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support FTD Platform Settings Syslog Settings, minimum required version is 7.7", d.client.FMCVersion))
		return
	}
	var config FTDPlatformSettingsSyslogSettings

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
