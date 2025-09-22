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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &FTDPlatformSettingsDNSDataSource{}
	_ datasource.DataSourceWithConfigure = &FTDPlatformSettingsDNSDataSource{}
)

func NewFTDPlatformSettingsDNSDataSource() datasource.DataSource {
	return &FTDPlatformSettingsDNSDataSource{}
}

type FTDPlatformSettingsDNSDataSource struct {
	client *fmc.Client
}

func (d *FTDPlatformSettingsDNSDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ftd_platform_settings_dns"
}

func (d *FTDPlatformSettingsDNSDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the FTD Platform Settings DNS.").AddMinimumVersionHeaderDescription().AddMinimumVersionDescription("7.7").String,

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
				MarkdownDescription: "Type of the object; this value is always 'DNSSetting'.",
				Computed:            true,
			},
			"server_groups": schema.ListNestedAttribute{
				MarkdownDescription: "List of DNS servers that will be used by device. It is mandatory to define at least one DNS server group marked as default.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "ID of the DNS server group object.",
							Computed:            true,
						},
						"is_default": schema.BoolAttribute{
							MarkdownDescription: "Set DNS server group as default. Any DNS resolution request that does not match the filters for other groups will be resolved using the servers in this group.",
							Computed:            true,
						},
						"filter_domains": schema.ListAttribute{
							MarkdownDescription: "Mandatory for non-default groups. The group will be used for DNS resolutions for these domains only.",
							ElementType:         types.StringType,
							Computed:            true,
						},
					},
				},
			},
			"expire_entry_timer": schema.Int64Attribute{
				MarkdownDescription: "Minimum time-to-live (TTL) for the DNS entry (in minutes).",
				Computed:            true,
			},
			"poll_timer": schema.Int64Attribute{
				MarkdownDescription: "Time limit after which the device queries the DNS server to resolve the name.",
				Computed:            true,
			},
			"interface_objects": schema.ListNestedAttribute{
				MarkdownDescription: "List of interface objects (Security Zones or Interface Groups) to be used for DNS resolution. If not specified, the device uses all interfaces for DNS resolution.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "ID of the interface object (Security Zone or Interface Group).",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Type of the interface object.",
							Computed:            true,
						},
					},
				},
			},
			"use_management_interface": schema.BoolAttribute{
				MarkdownDescription: "Enable lookup via management/diagnostic interface.",
				Computed:            true,
			},
		},
	}
}

func (d *FTDPlatformSettingsDNSDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *FTDPlatformSettingsDNSDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	// Check if FMC client is connected to supports this object
	if d.client.FMCVersionParsed.LessThan(minFMCVersionFTDPlatformSettingsDNS) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support FTD Platform Settings DNS, minimum required version is 7.7", d.client.FMCVersion))
		return
	}
	var config FTDPlatformSettingsDNS

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
