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
	_ datasource.DataSource              = &StandardCommunityListsDataSource{}
	_ datasource.DataSourceWithConfigure = &StandardCommunityListsDataSource{}
)

func NewStandardCommunityListsDataSource() datasource.DataSource {
	return &StandardCommunityListsDataSource{}
}

type StandardCommunityListsDataSource struct {
	client *fmc.Client
}

func (d *StandardCommunityListsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_standard_community_lists"
}

func (d *StandardCommunityListsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the Standard Community Lists.").String,

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
				MarkdownDescription: "Map of Standard Community Lists. The key of the map is the name of the individual Standard Community List.",
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Id of the managed Standard Community List.",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Type of the object; this value is always 'CommunityListObject'.",
							Computed:            true,
						},
						"entries": schema.ListNestedAttribute{
							MarkdownDescription: "List of entries",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"action": schema.StringAttribute{
										MarkdownDescription: "Indicate redistribution access.",
										Computed:            true,
									},
									"communities": schema.StringAttribute{
										MarkdownDescription: "List of communities. Separate multiple values by space. Valid values can be from 1 to 4294967295 or from 0:1 to 65534:65535",
										Computed:            true,
									},
									"internet": schema.BoolAttribute{
										MarkdownDescription: "Specify Internet well-known community.",
										Computed:            true,
									},
									"no_advertise": schema.BoolAttribute{
										MarkdownDescription: "Specify No-advertise well-known community.",
										Computed:            true,
									},
									"no_export": schema.BoolAttribute{
										MarkdownDescription: "Specify No-export well-known community.",
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

func (d *StandardCommunityListsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *StandardCommunityListsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config StandardCommunityLists

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
