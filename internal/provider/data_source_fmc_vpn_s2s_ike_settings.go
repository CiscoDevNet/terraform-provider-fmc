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
	_ datasource.DataSource              = &VPNS2SIKESettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &VPNS2SIKESettingsDataSource{}
)

func NewVPNS2SIKESettingsDataSource() datasource.DataSource {
	return &VPNS2SIKESettingsDataSource{}
}

type VPNS2SIKESettingsDataSource struct {
	client *fmc.Client
}

func (d *VPNS2SIKESettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_s2s_ike_settings"
}

func (d *VPNS2SIKESettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the VPN S2S IKE Settings.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the object",
				Required:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "Name of the FMC domain",
				Optional:            true,
			},
			"vpn_s2s_id": schema.StringAttribute{
				MarkdownDescription: "Id of the parent VPN S2S Topology.",
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the object; this value is always 'IkeSetting'.",
				Computed:            true,
			},
			"ikev1_authentication_type": schema.StringAttribute{
				MarkdownDescription: "Authentication method for IKEv1.",
				Computed:            true,
			},
			"ikev1_automatic_pre_shared_key_length": schema.Int64Attribute{
				MarkdownDescription: "Length of the automatically generated pre-shared key for IKEv1.",
				Computed:            true,
			},
			"ikev1_manual_pre_shared_key": schema.StringAttribute{
				MarkdownDescription: "Manually configured pre-shared key for IKEv1.",
				Computed:            true,
			},
			"ikev1_policies": schema.SetNestedAttribute{
				MarkdownDescription: "Set of policies for IKEv1.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Id of the IKEv1 policy",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Name of the IKEv1 policy",
							Computed:            true,
						},
					},
				},
			},
			"ikev2_authentication_type": schema.StringAttribute{
				MarkdownDescription: "Authentication method for IKEv2.",
				Computed:            true,
			},
			"ikev2_automatic_pre_shared_key_length": schema.Int64Attribute{
				MarkdownDescription: "Length of the automatically generated pre-shared key for IKEv2.",
				Computed:            true,
			},
			"ikev2_manual_pre_shared_key": schema.StringAttribute{
				MarkdownDescription: "Manually configured pre-shared key for IKEv2.",
				Computed:            true,
			},
			"ikev2_enforce_hex_based_pre_shared_key": schema.BoolAttribute{
				MarkdownDescription: "Enforce use of a hex-based pre-shared key for IKEv2.",
				Computed:            true,
			},
			"ikev2_policies": schema.SetNestedAttribute{
				MarkdownDescription: "Set of policies for IKEv2 settings.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Id of the IKEv2 policy",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Name of the IKEv2 policy",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *VPNS2SIKESettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *VPNS2SIKESettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config VPNS2SIKESettings

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
