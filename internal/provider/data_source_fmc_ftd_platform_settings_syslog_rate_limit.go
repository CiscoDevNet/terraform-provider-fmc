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
	_ datasource.DataSource              = &FTDPlatformSettingsSyslogRateLimitDataSource{}
	_ datasource.DataSourceWithConfigure = &FTDPlatformSettingsSyslogRateLimitDataSource{}
)

func NewFTDPlatformSettingsSyslogRateLimitDataSource() datasource.DataSource {
	return &FTDPlatformSettingsSyslogRateLimitDataSource{}
}

type FTDPlatformSettingsSyslogRateLimitDataSource struct {
	client *fmc.Client
}

func (d *FTDPlatformSettingsSyslogRateLimitDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ftd_platform_settings_syslog_rate_limit"
}

func (d *FTDPlatformSettingsSyslogRateLimitDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the FTD Platform Settings Syslog Rate Limit.").AddMinimumVersionHeaderDescription().AddMinimumVersionDescription("7.7").String,

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
				MarkdownDescription: "Type of the object; this value is always 'RateLimit'.",
				Computed:            true,
			},
			"rate_limit_type": schema.StringAttribute{
				MarkdownDescription: "Rate limit type - severity level or syslog ID.",
				Computed:            true,
			},
			"rate_limit_value": schema.StringAttribute{
				MarkdownDescription: "Severity level or syslog message ID as per `rate_limit_type`.",
				Computed:            true,
			},
			"number_of_messages": schema.Int64Attribute{
				MarkdownDescription: "Maximum number of messages of the specified type allowed in the specified time period.",
				Computed:            true,
			},
			"interval": schema.Int64Attribute{
				MarkdownDescription: "Number of seconds before the rate limit counter resets.",
				Computed:            true,
			},
		},
	}
}

func (d *FTDPlatformSettingsSyslogRateLimitDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *FTDPlatformSettingsSyslogRateLimitDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	// Check if FMC client is connected to supports this object
	if d.client.FMCVersionParsed.LessThan(minFMCVersionFTDPlatformSettingsSyslogRateLimit) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support FTD Platform Settings Syslog Rate Limit, minimum required version is 7.7", d.client.FMCVersion))
		return
	}
	var config FTDPlatformSettingsSyslogRateLimit

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
