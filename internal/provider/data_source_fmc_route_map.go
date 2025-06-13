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
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &RouteMapDataSource{}
	_ datasource.DataSourceWithConfigure = &RouteMapDataSource{}
)

func NewRouteMapDataSource() datasource.DataSource {
	return &RouteMapDataSource{}
}

type RouteMapDataSource struct {
	client *fmc.Client
}

func (d *RouteMapDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_route_map"
}

func (d *RouteMapDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the Route Map.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the object",
				Optional:            true,
				Computed:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "Name of the FMC domain",
				Optional:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the Route Map object.",
				Optional:            true,
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the object; this value is always 'RouteMap'.",
				Computed:            true,
			},
			"overridable": schema.BoolAttribute{
				MarkdownDescription: "Whether the object values can be overridden.",
				Computed:            true,
			},
			"entries": schema.ListNestedAttribute{
				MarkdownDescription: "List of route map entries.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"action": schema.StringAttribute{
							MarkdownDescription: "Indicate the redistribution access.",
							Computed:            true,
						},
						"match_security_zones": schema.ListNestedAttribute{
							MarkdownDescription: "Match traffic based on the (ingress/egress) security_zones.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "ID of the security zone.",
										Computed:            true,
									},
								},
							},
						},
						"match_interface_names": schema.ListAttribute{
							MarkdownDescription: "List of interface names that are not in the zones.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"match_ipv4_address_access_lists": schema.ListNestedAttribute{
							MarkdownDescription: "Match routes based on the route address.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the object.",
										Computed:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of the access list.",
										Computed:            true,
									},
								},
							},
						},
						"match_ipv4_address_prefix_lists": schema.ListNestedAttribute{
							MarkdownDescription: "Match routes based on the route address.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the object.",
										Computed:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of the access list.",
										Computed:            true,
									},
								},
							},
						},
						"match_ipv4_next_hop_access_lists": schema.ListNestedAttribute{
							MarkdownDescription: "Match routes based on the next hop address of a route.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the object.",
										Computed:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of the access list.",
										Computed:            true,
									},
								},
							},
						},
						"match_ipv4_next_hop_prefix_lists": schema.ListNestedAttribute{
							MarkdownDescription: "Match routes based on the next hop address of a route.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the object.",
										Computed:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of the access list.",
										Computed:            true,
									},
								},
							},
						},
						"match_ipv4_route_source_access_lists": schema.ListNestedAttribute{
							MarkdownDescription: "Match routes based on the advertising source address of the route.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the object.",
										Computed:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of the access list.",
										Computed:            true,
									},
								},
							},
						},
						"match_ipv4_route_source_prefix_lists": schema.ListNestedAttribute{
							MarkdownDescription: "Match routes based on the advertising source address of the route",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the object.",
										Computed:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of the access list.",
										Computed:            true,
									},
								},
							},
						},
						"match_ipv6_address_extended_access_list_id": schema.StringAttribute{
							MarkdownDescription: "Match routes based on the route address.",
							Computed:            true,
						},
						"match_ipv6_address_prefix_list_id": schema.StringAttribute{
							MarkdownDescription: "Match routes based on the route address.",
							Computed:            true,
						},
						"match_ipv6_next_hop_extended_access_list_id": schema.StringAttribute{
							MarkdownDescription: "Match routes based on the next hop address of a route.",
							Computed:            true,
						},
						"match_ipv6_next_hop_prefix_list_id": schema.StringAttribute{
							MarkdownDescription: "Match routes based on the next hop address of a route.",
							Computed:            true,
						},
						"match_ipv6_route_source_extended_access_list_id": schema.StringAttribute{
							MarkdownDescription: "Match routes based on the advertising source address of the route.",
							Computed:            true,
						},
						"match_ipv6_route_source_prefix_list_id": schema.StringAttribute{
							MarkdownDescription: "Match routes based on the advertising source address of the route",
							Computed:            true,
						},
						"match_bgp_as_path_lists": schema.ListNestedAttribute{
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
						"match_bgp_community_lists": schema.ListNestedAttribute{
							MarkdownDescription: "List of Standard/Expanded Community Lists.",
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
						"match_bgp_extended_community_lists": schema.ListNestedAttribute{
							MarkdownDescription: "List of Extended Community Lists.",
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
						"match_bgp_policy_lists": schema.ListNestedAttribute{
							MarkdownDescription: "List of Policy Lists.",
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
						"match_metric_route_values": schema.ListAttribute{
							MarkdownDescription: "List of metric values to match.",
							ElementType:         types.Int64Type,
							Computed:            true,
						},
						"match_tag_values": schema.ListAttribute{
							MarkdownDescription: "Tag values.",
							ElementType:         types.Int64Type,
							Computed:            true,
						},
						"match_route_type_external_1": schema.BoolAttribute{
							MarkdownDescription: "Match external type 1 routes.",
							Computed:            true,
						},
						"match_route_type_external_2": schema.BoolAttribute{
							MarkdownDescription: "Match external type 2 routes.",
							Computed:            true,
						},
						"match_route_type_internal": schema.BoolAttribute{
							MarkdownDescription: "Match internal routes.",
							Computed:            true,
						},
						"match_route_type_local": schema.BoolAttribute{
							MarkdownDescription: "Match local routes.",
							Computed:            true,
						},
						"match_route_type_nssa_external_1": schema.BoolAttribute{
							MarkdownDescription: "Match NSSA external type 1 routes.",
							Computed:            true,
						},
						"match_route_type_nssa_external_2": schema.BoolAttribute{
							MarkdownDescription: "Match NSSA external type 2 routes.",
							Computed:            true,
						},
						"set_metric_bandwidth": schema.Int64Attribute{
							MarkdownDescription: "Set the metric bandwidth value in Kbits per second.",
							Computed:            true,
						},
						"set_metric_type": schema.StringAttribute{
							MarkdownDescription: "Set the metric type.",
							Computed:            true,
						},
						"set_bgp_as_path_prepend": schema.ListAttribute{
							MarkdownDescription: "Set the AS path prepend value.",
							ElementType:         types.Int64Type,
							Computed:            true,
						},
						"set_bgp_as_path_prepend_last_as": schema.Int64Attribute{
							MarkdownDescription: "Set the AS path prepend value.",
							Computed:            true,
						},
						"set_bgp_as_path_convert_route_tag_into_as_path": schema.BoolAttribute{
							MarkdownDescription: "Convert the route tag into an AS path.",
							Computed:            true,
						},
						"set_bgp_community_none": schema.BoolAttribute{
							MarkdownDescription: "Set the specific community to none.",
							Computed:            true,
						},
						"set_bgp_community_specific_community": schema.Int64Attribute{
							MarkdownDescription: "Set the specific community.",
							Computed:            true,
						},
						"set_bgp_community_add_to_existing_communities": schema.BoolAttribute{
							MarkdownDescription: "Set the specific community to none.",
							Computed:            true,
						},
						"set_bgp_community_internet": schema.BoolAttribute{
							MarkdownDescription: "Set the specific community to none.",
							Computed:            true,
						},
						"set_bgp_community_no_advertise": schema.BoolAttribute{
							MarkdownDescription: "Set the specific community to none.",
							Computed:            true,
						},
						"set_bgp_community_no_export": schema.BoolAttribute{
							MarkdownDescription: "Set the specific community to none.",
							Computed:            true,
						},
						"set_bgp_community_route_target": schema.StringAttribute{
							MarkdownDescription: "Set the extended community route target.",
							Computed:            true,
						},
						"set_bgp_community_add_to_existing_extended_communities": schema.BoolAttribute{
							MarkdownDescription: "Set the extended community additive.",
							Computed:            true,
						},
						"set_bgp_automatic_tag": schema.BoolAttribute{
							MarkdownDescription: "Set the automatic tag setting.",
							Computed:            true,
						},
						"set_bgp_local_preference": schema.Int64Attribute{
							MarkdownDescription: "Set the local preference value.",
							Computed:            true,
						},
						"set_bgp_weight": schema.Int64Attribute{
							MarkdownDescription: "Set the weight value.",
							Computed:            true,
						},
						"set_bgp_origin": schema.StringAttribute{
							MarkdownDescription: "Set the origin value.",
							Computed:            true,
						},
						"set_bgp_ipv4_next_hop": schema.StringAttribute{
							MarkdownDescription: "Set the next hop IPv4 address.",
							Computed:            true,
						},
						"set_bgp_ipv4_next_hop_specific_ip": schema.ListAttribute{
							MarkdownDescription: "Set the next hop IPv4 address.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"set_bgp_ipv4_prefix_list_id": schema.StringAttribute{
							MarkdownDescription: "Set the prefix list for IPv4.",
							Computed:            true,
						},
						"set_bgp_ipv6_next_hop": schema.StringAttribute{
							MarkdownDescription: "Set the next hop IPv6 address.",
							Computed:            true,
						},
						"set_bgp_ipv6_next_hop_specific_ip": schema.ListAttribute{
							MarkdownDescription: "Set the next hop IPv6 address.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"set_bgp_ipv6_prefix_list_id": schema.StringAttribute{
							MarkdownDescription: "Set the prefix list for IPv6.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}
func (d *RouteMapDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *RouteMapDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *RouteMapDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config RouteMap

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
	if config.Id.IsNull() && !config.Name.IsNull() {
		offset := 0
		limit := 1000
		for page := 1; ; page++ {
			queryString := fmt.Sprintf("?limit=%d&offset=%d&expanded=true", limit, offset)
			res, err := d.client.Get(config.getPath()+queryString, reqMods...)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects, got error: %s", err))
				return
			}
			if value := res.Get("items"); len(value.Array()) > 0 {
				value.ForEach(func(k, v gjson.Result) bool {
					if config.Name.ValueString() == v.Get("name").String() {
						config.Id = types.StringValue(v.Get("id").String())
						tflog.Debug(ctx, fmt.Sprintf("%s: Found object with name '%v', id: %v", config.Id.ValueString(), config.Name.ValueString(), config.Id.ValueString()))
						return false
					}
					return true
				})
			}
			if !config.Id.IsNull() || !res.Get("paging.next.0").Exists() {
				break
			}
			offset += limit
		}

		if config.Id.IsNull() {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with name: %v", config.Name.ValueString()))
			return
		}
	}
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
