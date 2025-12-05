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
	_ datasource.DataSource              = &KeyChainsDataSource{}
	_ datasource.DataSourceWithConfigure = &KeyChainsDataSource{}
)

func NewKeyChainsDataSource() datasource.DataSource {
	return &KeyChainsDataSource{}
}

type KeyChainsDataSource struct {
	client *fmc.Client
}

func (d *KeyChainsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_key_chains"
}

func (d *KeyChainsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the Key Chains.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the object",
				Computed:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "Name of the FMC domain",
				Optional:            true,
			},
			"items": schema.MapNestedAttribute{
				MarkdownDescription: "Map of Key Chains. The key of the map is the name of the individual Key Chain.",
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Id of the Key Chain object.",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Type of the object; this value is always 'KeyChainObject'.",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "Description of the object.",
							Computed:            true,
						},
						"keys": schema.ListNestedAttribute{
							MarkdownDescription: "Keys",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.Int64Attribute{
										MarkdownDescription: "Key ID.",
										Computed:            true,
									},
									"key": schema.StringAttribute{
										MarkdownDescription: "Crypto key string.",
										Computed:            true,
									},
									"accept_lifetime_start": schema.StringAttribute{
										MarkdownDescription: "Start time for key acceptance lifetime in YYYY-MM-DDTHH:mm:ss format.",
										Computed:            true,
									},
									"accept_lifetime_end_type": schema.StringAttribute{
										MarkdownDescription: "Type of end time for key acceptance lifetime.",
										Computed:            true,
									},
									"accept_lifetime_end": schema.StringAttribute{
										MarkdownDescription: "End time for key acceptance lifetime in YYYY-MM-DDTHH:mm:ss format (`DATETIME`) or duration in seconds (`DURATION`).",
										Computed:            true,
									},
									"send_lifetime_start": schema.StringAttribute{
										MarkdownDescription: "Start time for key send lifetime in YYYY-MM-DDTHH:mm:ss format.",
										Computed:            true,
									},
									"send_lifetime_end_type": schema.StringAttribute{
										MarkdownDescription: "Type of end time for key send lifetime.",
										Computed:            true,
									},
									"send_lifetime_end": schema.StringAttribute{
										MarkdownDescription: "End time for key send lifetime in YYYY-MM-DDTHH:mm:ss format (`DATETIME`) or duration in seconds (`DURATION`).",
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

func (d *KeyChainsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *KeyChainsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config KeyChains

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

	// Get all objects from FMC
	urlPath := config.getPath() + "?expanded=true"
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
