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
	_ datasource.DataSource              = &FTDPlatformSettingsHTTPAccessDataSource{}
	_ datasource.DataSourceWithConfigure = &FTDPlatformSettingsHTTPAccessDataSource{}
)

func NewFTDPlatformSettingsHTTPAccessDataSource() datasource.DataSource {
	return &FTDPlatformSettingsHTTPAccessDataSource{}
}

type FTDPlatformSettingsHTTPAccessDataSource struct {
	client *fmc.Client
}

func (d *FTDPlatformSettingsHTTPAccessDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ftd_platform_settings_http_access"
}

func (d *FTDPlatformSettingsHTTPAccessDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the FTD Platform Settings HTTP Access.").AddMinimumVersionHeaderDescription().AddMinimumVersionDescription("7.7").String,

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
				MarkdownDescription: "Type of the object; this value is always 'HttpAccessSetting'.",
				Computed:            true,
			},
			"server_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable HTTP server.",
				Computed:            true,
			},
			"server_port": schema.Int64Attribute{
				MarkdownDescription: "Port on which the HTTP server will listen. Please don't use 80 or 1443.",
				Computed:            true,
			},
			"configurations": schema.ListNestedAttribute{
				MarkdownDescription: "List of allowed HTTP connections.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"source_network_object_id": schema.StringAttribute{
							MarkdownDescription: "Id of network object (host, network, network group) defining the source IP addresses from which HTTP access is allowed.",
							Computed:            true,
						},
						"interface_literals": schema.SetAttribute{
							MarkdownDescription: "List of interface literals on which HTTP server is available.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"interface_objects": schema.SetNestedAttribute{
							MarkdownDescription: "List of interface objects (Security Zones or Interface Groups) on which HTTP server is available.",
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

func (d *FTDPlatformSettingsHTTPAccessDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *FTDPlatformSettingsHTTPAccessDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	// Check if FMC client is connected to supports this object
	if d.client.FMCVersionParsed.LessThan(minFMCVersionFTDPlatformSettingsHTTPAccess) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support FTD Platform Settings HTTP Access, minimum required version is 7.7", d.client.FMCVersion))
		return
	}
	var config FTDPlatformSettingsHTTPAccess

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
