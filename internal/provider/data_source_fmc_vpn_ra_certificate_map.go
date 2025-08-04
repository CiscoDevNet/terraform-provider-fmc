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
	_ datasource.DataSource              = &VPNRACertificateMapDataSource{}
	_ datasource.DataSourceWithConfigure = &VPNRACertificateMapDataSource{}
)

func NewVPNRACertificateMapDataSource() datasource.DataSource {
	return &VPNRACertificateMapDataSource{}
}

type VPNRACertificateMapDataSource struct {
	client *fmc.Client
}

func (d *VPNRACertificateMapDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_ra_certificate_map"
}

func (d *VPNRACertificateMapDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the VPN RA Certificate Map.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the object",
				Required:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "Name of the FMC domain",
				Optional:            true,
			},
			"vpn_ra_id": schema.StringAttribute{
				MarkdownDescription: "Id of the parent VPN RA Configuration.",
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the object; this value is always 'RaVpnCertificateMapSetting'.",
				Computed:            true,
			},
			"use_group_url": schema.BoolAttribute{
				MarkdownDescription: "Use group URL if group URL and Certificate Map match different Connection Profiles",
				Computed:            true,
			},
			"use_certificate_to_connection_profile_mappings": schema.BoolAttribute{
				MarkdownDescription: "Use the configured rules to match Certificate to Connection Profile",
				Computed:            true,
			},
			"certificate_to_connection_profile_mappings": schema.ListNestedAttribute{
				MarkdownDescription: "Mapping of certificates to connection profiles.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"certificate_map_id": schema.StringAttribute{
							MarkdownDescription: "Id of the Certificate Map.",
							Computed:            true,
						},
						"connection_profile_id": schema.StringAttribute{
							MarkdownDescription: "Id of the Connection Profile.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *VPNRACertificateMapDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *VPNRACertificateMapDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config VPNRACertificateMap

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
