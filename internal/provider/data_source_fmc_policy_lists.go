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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &PolicyListsDataSource{}
	_ datasource.DataSourceWithConfigure = &PolicyListsDataSource{}
)

func NewPolicyListsDataSource() datasource.DataSource {
	return &PolicyListsDataSource{}
}

type PolicyListsDataSource struct {
	client *fmc.Client
}

func (d *PolicyListsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policy_lists"
}

func (d *PolicyListsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the Policy Lists.").String,

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
				MarkdownDescription: "Map of Policy Lists. The key of the map is the name of the individual Policy List.",
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Id of the Policy List.",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Type of the object; this value is always 'PolicyList'.",
							Computed:            true,
						},
						"action": schema.StringAttribute{
							MarkdownDescription: "Select whether to allow or block access for matching conditions.",
							Computed:            true,
						},
						"interfaces": schema.SetNestedAttribute{
							MarkdownDescription: "Security zones/interface groups that contain the interfaces through which the device communicates with the management station.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the object.",
										Computed:            true,
									},
								},
							},
						},
						"interface_names": schema.ListAttribute{
							MarkdownDescription: "List of interface names that are not in the zones.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"address_standard_access_lists": schema.SetNestedAttribute{
							MarkdownDescription: "Redistribute any routes that have a destination address that is permitted by a standard access list. `address_standard_access_lists` and `address_ipv4_prefix_lists` are mutually exclusive.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the object.",
										Computed:            true,
									},
								},
							},
						},
						"address_ipv4_prefix_lists": schema.SetNestedAttribute{
							MarkdownDescription: "Redistribute any routes that have a destination address that is permitted by a prefix list. `address_standard_access_lists` and `address_ipv4_prefix_lists` are mutually exclusive.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the object.",
										Computed:            true,
									},
								},
							},
						},
						"next_hop_standard_access_lists": schema.SetNestedAttribute{
							MarkdownDescription: "Redistribute any routes that have a next hop router address passed by a standard access list. `next_hop_standard_access_lists` and `next_hop_ipv4_prefix_lists` are mutually exclusive.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the object.",
										Computed:            true,
									},
								},
							},
						},
						"next_hop_ipv4_prefix_lists": schema.SetNestedAttribute{
							MarkdownDescription: "Redistribute any routes that have a next hop router address passed by a prefix list. `next_hop_standard_access_lists` and `next_hop_ipv4_prefix_lists` are mutually exclusive.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the object.",
										Computed:            true,
									},
								},
							},
						},
						"route_source_standard_access_lists": schema.SetNestedAttribute{
							MarkdownDescription: "Redistribute routes that have been advertised by routers at the address specified by the access list. `route_source_standard_access_lists` and `route_source_ipv4_prefix_lists` are mutually exclusive.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the object.",
										Computed:            true,
									},
								},
							},
						},
						"route_source_ipv4_prefix_lists": schema.SetNestedAttribute{
							MarkdownDescription: "Redistribute routes that have been advertised by routers at the address specified by the prefix list. `route_source_standard_access_lists` and `route_source_ipv4_prefix_lists` are mutually exclusive.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the object.",
										Computed:            true,
									},
								},
							},
						},
						"as_paths": schema.SetNestedAttribute{
							MarkdownDescription: "Match a BGP autonomous system path.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of object.",
										Computed:            true,
									},
								},
							},
						},
						"community_lists": schema.SetNestedAttribute{
							MarkdownDescription: "Standard/Expanded Community Lists.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the object.",
										Computed:            true,
									},
								},
							},
						},
						"extended_community_lists": schema.SetNestedAttribute{
							MarkdownDescription: "Extended Community Lists.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the object.",
										Computed:            true,
									},
								},
							},
						},
						"match_community_exactly": schema.BoolAttribute{
							MarkdownDescription: "Match BGP community exactly with the specified community.",
							Computed:            true,
						},
						"metric": schema.Int64Attribute{
							MarkdownDescription: "Match routes that have a specified metric. Specyfing multiple values is not supported due to FMC API bug.",
							Computed:            true,
						},
						"tag": schema.Int64Attribute{
							MarkdownDescription: "Match routes that have a specified security group tag. Specyfing multiple values is not supported due to FMC API bug.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *PolicyListsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *PolicyListsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config PolicyLists

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
