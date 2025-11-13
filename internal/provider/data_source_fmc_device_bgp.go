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
	_ datasource.DataSource              = &DeviceBGPDataSource{}
	_ datasource.DataSourceWithConfigure = &DeviceBGPDataSource{}
)

func NewDeviceBGPDataSource() datasource.DataSource {
	return &DeviceBGPDataSource{}
}

type DeviceBGPDataSource struct {
	client *fmc.Client
}

func (d *DeviceBGPDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_bgp"
}

func (d *DeviceBGPDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the Device BGP.").String,

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
			"vrf_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the parent VRF.").String,
				Optional:            true,
			},
			"device_id": schema.StringAttribute{
				MarkdownDescription: "Id of the parent device.",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the object; this is always 'bgp'",
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the object; this is always 'bgp'",
				Computed:            true,
			},
			"as_number": schema.StringAttribute{
				MarkdownDescription: "Autonomus System (AS) number",
				Optional:            true,
				Computed:            true,
			},
			"ipv4_address_family_id": schema.StringAttribute{
				MarkdownDescription: "IPv4 Address Family Id",
				Computed:            true,
			},
			"ipv4_address_family_type": schema.StringAttribute{
				MarkdownDescription: "IPv4 Address Family Type",
				Computed:            true,
			},
			"ipv4_learned_route_map_id": schema.StringAttribute{
				MarkdownDescription: "Learned Route Map ID",
				Computed:            true,
			},
			"ipv4_default_information_orginate": schema.BoolAttribute{
				MarkdownDescription: "Generate default route",
				Computed:            true,
			},
			"ipv4_auto_summary": schema.BoolAttribute{
				MarkdownDescription: "Summarize subnet routes into network level routes",
				Computed:            true,
			},
			"ipv4_suppress_inactive_routes": schema.BoolAttribute{
				MarkdownDescription: "Suppressing advertisement of inactive routes",
				Computed:            true,
			},
			"ipv4_synchronization": schema.BoolAttribute{
				MarkdownDescription: "Synchronize between BGP and IGP systems",
				Computed:            true,
			},
			"ipv4_redistribute_ibgp_into_igp": schema.BoolAttribute{
				MarkdownDescription: "Redistribute IBGP into IGP. Use filtering to limit the number of prefixes that are redistributed.",
				Computed:            true,
			},
			"ipv4_external_distance": schema.Int64Attribute{
				MarkdownDescription: "Administrative route distance for external routes",
				Computed:            true,
			},
			"ipv4_internal_distance": schema.Int64Attribute{
				MarkdownDescription: "Administrative route distance for internal routes",
				Computed:            true,
			},
			"ipv4_local_distance": schema.Int64Attribute{
				MarkdownDescription: "Administrative route distance for local routes",
				Computed:            true,
			},
			"ipv4_number_of_ibgp_paths": schema.Int64Attribute{
				MarkdownDescription: "Number of paths to use for IBGP",
				Computed:            true,
			},
			"ipv4_number_of_ebgp_paths": schema.Int64Attribute{
				MarkdownDescription: "Number of paths to use for EBGP",
				Computed:            true,
			},
			"ipv4_neighbors": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: "IP address of the BGP neighbor",
							Computed:            true,
						},
						"remote_as": schema.StringAttribute{
							MarkdownDescription: "AS number of the BGP neighbor",
							Computed:            true,
						},
						"bfd_fallover": schema.StringAttribute{
							MarkdownDescription: "BFD Fallover",
							Computed:            true,
						},
						"update_source_interface_id": schema.StringAttribute{
							MarkdownDescription: "Interface ID for the update source",
							Computed:            true,
						},
						"enable_address": schema.BoolAttribute{
							MarkdownDescription: "Enable communication with this BGP neighbor",
							Computed:            true,
						},
						"as_override": schema.BoolAttribute{
							MarkdownDescription: "Enable overriding of the AS number of the originating router with the AS number of the sending BGP router.",
							Computed:            true,
						},
						"graceful_restart": schema.BoolAttribute{
							MarkdownDescription: "Enable graceful restart for the neighbor.",
							Computed:            true,
						},
						"shutdown_administratively": schema.BoolAttribute{
							MarkdownDescription: "Disable a neighbor or peer group",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "Description of the neighbor",
							Computed:            true,
						},
						"filter_access_lists": schema.ListNestedAttribute{
							MarkdownDescription: "Set incoming or outgoing Access List to distribute BGP neighbor information.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"access_list_id": schema.StringAttribute{
										MarkdownDescription: "Access List ID",
										Computed:            true,
									},
									"update_direction": schema.StringAttribute{
										MarkdownDescription: "Filter direction",
										Computed:            true,
									},
								},
							},
						},
						"filter_route_maps": schema.ListNestedAttribute{
							MarkdownDescription: "Set incoming or outgoing Route Maps to apply a route map to incoming or outgoing routes.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"route_map_id": schema.StringAttribute{
										MarkdownDescription: "Route Map ID",
										Computed:            true,
									},
									"update_direction": schema.StringAttribute{
										MarkdownDescription: "Filter direction",
										Computed:            true,
									},
								},
							},
						},
						"filter_prefix_lists": schema.ListNestedAttribute{
							MarkdownDescription: "Set incoming or outgoing Prefix List to distribute BGP neighbor information.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"prefix_list_id": schema.StringAttribute{
										MarkdownDescription: "Prefix List ID",
										Computed:            true,
									},
									"update_direction": schema.StringAttribute{
										MarkdownDescription: "Filter direction",
										Computed:            true,
									},
								},
							},
						},
						"filter_as_paths": schema.ListNestedAttribute{
							MarkdownDescription: "Set incoming or outgoing AS path filter to distribute BGP neighbor information.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"as_path_id": schema.StringAttribute{
										MarkdownDescription: "AS Path ID",
										Computed:            true,
									},
									"as_path_name": schema.StringAttribute{
										MarkdownDescription: "AS Path Name",
										Computed:            true,
									},
									"update_direction": schema.StringAttribute{
										MarkdownDescription: "Filter direction",
										Computed:            true,
									},
								},
							},
						},
						"filter_maximum_prefixes": schema.Int64Attribute{
							MarkdownDescription: "Maximum number of prefixes allowed from the neighbor",
							Computed:            true,
						},
						"filter_warning_only": schema.BoolAttribute{
							MarkdownDescription: "Give only warning message when prefix limit exceeded. Can be set to `true` only. Use `neighbor_filter_threshold_value` to set mode to session termination.",
							Computed:            true,
						},
						"filter_threshold_value": schema.Int64Attribute{
							MarkdownDescription: "Threshold value for the maximum number of prefixes allowed from the neighbor (implies session termination).",
							Computed:            true,
						},
						"filter_restart_interval": schema.Int64Attribute{
							MarkdownDescription: "Time interval to restart the maximum prefix limit in Minutes",
							Computed:            true,
						},
						"routes_advertisement_interval": schema.Int64Attribute{
							MarkdownDescription: "Time interval to advertise routes in seconds",
							Computed:            true,
						},
						"routes_remove_private_as": schema.BoolAttribute{
							MarkdownDescription: "Remove private AS numbers from outgoing routing updates",
							Computed:            true,
						},
						"routes_generate_default_route_map_id": schema.StringAttribute{
							MarkdownDescription: "Generate default routes - Route Map",
							Computed:            true,
						},
						"routes_advertise_maps": schema.ListNestedAttribute{
							MarkdownDescription: "Define conditionally advertised routes.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"advertise_map_id": schema.StringAttribute{
										MarkdownDescription: "Specified route maps are advertised when conditions of the exist map or the non-exist map are met.",
										Computed:            true,
									},
									"use_exist_map": schema.BoolAttribute{
										MarkdownDescription: "Set mode to Exist Map (true) or Non-Exist Map (false).",
										Computed:            true,
									},
									"exist_nonexist_map_id": schema.StringAttribute{
										MarkdownDescription: "Specify exist / non-exist route map ID.",
										Computed:            true,
									},
								},
							},
						},
						"keepalive_interval": schema.Int64Attribute{
							MarkdownDescription: "Time interval to send keepalive messages in seconds",
							Computed:            true,
						},
						"hold_time": schema.Int64Attribute{
							MarkdownDescription: "Time interval to hold the neighbor in seconds",
							Computed:            true,
						},
						"minimum_hold_time": schema.Int64Attribute{
							MarkdownDescription: "Minimum hold time in seconds",
							Computed:            true,
						},
						"authentication_password": schema.StringAttribute{
							MarkdownDescription: "Setting password enables authentication.",
							Computed:            true,
							Sensitive:           true,
						},
						"send_community_attribute": schema.BoolAttribute{
							MarkdownDescription: "Send Community attribute to this neighbor",
							Computed:            true,
						},
						"next_hop_self": schema.BoolAttribute{
							MarkdownDescription: "Use itself as next hop for this neighbor",
							Computed:            true,
						},
						"disable_connection_verification": schema.BoolAttribute{
							MarkdownDescription: "Disable Connection Verification",
							Computed:            true,
						},
						"tcp_path_mtu_discovery": schema.BoolAttribute{
							MarkdownDescription: "Use TCP path MTU discovery.",
							Computed:            true,
						},
						"max_hop_count": schema.Int64Attribute{
							MarkdownDescription: "Maximum number of hops to reach the neighbor",
							Computed:            true,
						},
						"tcp_transport_mode": schema.BoolAttribute{
							MarkdownDescription: "True set it to active, False to passive.",
							Computed:            true,
						},
						"weight": schema.Int64Attribute{
							MarkdownDescription: "Weight of the neighbor",
							Computed:            true,
						},
						"version": schema.StringAttribute{
							MarkdownDescription: "Set BGP version: 0 - default, 4 - IPv4",
							Computed:            true,
						},
						"customized_local_as_number": schema.StringAttribute{
							MarkdownDescription: "Customize the AS number for the routes received from neighbor",
							Computed:            true,
						},
						"customized_no_prepend": schema.BoolAttribute{
							MarkdownDescription: "Do not prepend local AS number to routes received from neighbor",
							Computed:            true,
						},
						"customized_replace_as": schema.BoolAttribute{
							MarkdownDescription: "Replace real AS number with local AS number in routes received from neighbor",
							Computed:            true,
						},
						"customized_accept_both_as": schema.BoolAttribute{
							MarkdownDescription: "Accept either real AS number or local AS number in routes experienced from neighbor",
							Computed:            true,
						},
					},
				},
			},
			"ipv4_aggregate_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "Generate aggregate address information for IPv4.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"generate_as": schema.BoolAttribute{
							MarkdownDescription: "Generate AS set path information",
							Computed:            true,
						},
						"filter": schema.BoolAttribute{
							MarkdownDescription: "Filter all routes from updates (summary only)",
							Computed:            true,
						},
						"network_id": schema.StringAttribute{
							MarkdownDescription: "Network ID (desired network/host object)",
							Computed:            true,
						},
						"advertise_map_id": schema.StringAttribute{
							MarkdownDescription: "Advertise Route Map ID (select the routes to create AS_SET origin communities)",
							Computed:            true,
						},
						"attribute_map_id": schema.StringAttribute{
							MarkdownDescription: "Attribute Route Map ID (set the attribute of the aggregate route).",
							Computed:            true,
						},
						"suppress_map_id": schema.StringAttribute{
							MarkdownDescription: "Suppress Route Map ID (select the routes to be suppressed).",
							Computed:            true,
						},
					},
				},
			},
			"ipv4_filterings": schema.ListNestedAttribute{
				MarkdownDescription: "Filter routes or networks received in incoming BGP updates",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"access_list_id": schema.StringAttribute{
							MarkdownDescription: "Standard Access List ID that defines which networks are to be received and which are to be suppressed in routing updates.",
							Computed:            true,
						},
						"direction": schema.StringAttribute{
							MarkdownDescription: "Determine if the filter should be applied to inbound updates or outbound updates",
							Computed:            true,
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: "Routing process for which you want to filter",
							Computed:            true,
						},
						"process_id": schema.StringAttribute{
							MarkdownDescription: "Process ID for the OSPF routing protocol.",
							Computed:            true,
						},
					},
				},
			},
			"ipv4_networks": schema.ListNestedAttribute{
				MarkdownDescription: "Add networks that will be advertised by the BGP routing process",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"network_id": schema.StringAttribute{
							MarkdownDescription: "Network to be advertised by the BGP routing processes.",
							Computed:            true,
						},
						"route_map_id": schema.StringAttribute{
							MarkdownDescription: "Route Map ID that should be examined to filter the networks to be advertised",
							Computed:            true,
						},
					},
				},
			},
			"ipv4_redistributions": schema.ListNestedAttribute{
				MarkdownDescription: "Define the conditions for redistributing routes from another routing domain into BGP.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"source_protocol": schema.StringAttribute{
							MarkdownDescription: "Protocol to redistribute",
							Computed:            true,
						},
						"route_map_id": schema.StringAttribute{
							MarkdownDescription: "Route Map ID to filter the networks to be redistributed",
							Computed:            true,
						},
						"metric": schema.Int64Attribute{
							MarkdownDescription: "Metric for the redistributed route.",
							Computed:            true,
						},
						"process_id": schema.StringAttribute{
							MarkdownDescription: "OSPF process ID",
							Computed:            true,
						},
						"match_external1": schema.BoolAttribute{
							MarkdownDescription: "Match OSPF External 1 metrics",
							Computed:            true,
						},
						"match_external2": schema.BoolAttribute{
							MarkdownDescription: "Match OSPF External 2 metrics",
							Computed:            true,
						},
						"match_internal": schema.BoolAttribute{
							MarkdownDescription: "Match OSPF Internal metrics",
							Computed:            true,
						},
						"match_nssa_external1": schema.BoolAttribute{
							MarkdownDescription: "Match OSPF NSSA External 1 metrics",
							Computed:            true,
						},
						"match_nssa_external2": schema.BoolAttribute{
							MarkdownDescription: "Match OSPF NSSA External 2 metrics",
							Computed:            true,
						},
					},
				},
			},
			"ipv4_route_injections": schema.ListNestedAttribute{
				MarkdownDescription: "Define routes to be conditionally injected into the BGP routing table.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"inject_route_map_id": schema.StringAttribute{
							MarkdownDescription: "Inject Route Map ID (prefixes to inject into the local BGP routing table)",
							Computed:            true,
						},
						"exist_route_map_id": schema.StringAttribute{
							MarkdownDescription: "Exist Route Map ID containing the prefixes that the BGP speaker will track",
							Computed:            true,
						},
						"inherit_attributes": schema.StringAttribute{
							MarkdownDescription: "Injected route will inherit the attributes of the aggregate route",
							Computed:            true,
						},
					},
				},
			},
			"ipv4_import_route_targets": schema.ListAttribute{
				MarkdownDescription: "Route Target extended community that you want to match for the routes to be imported. Applicable only for BGP in VRF context.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"ipv4_export_route_targets": schema.ListAttribute{
				MarkdownDescription: "Route Target extended community to tag the source virtual router's routes with the route target value. Applicable only for BGP in VRF context.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"ipv4_import_global_vrf_route_map_id": schema.StringAttribute{
				MarkdownDescription: "Leak the global virtual router routes to the user-defined virtual router based on the specified Route Map. Applicable only for BGP in VRF context.",
				Computed:            true,
			},
			"ipv4_export_global_vrf_route_map_id": schema.StringAttribute{
				MarkdownDescription: "Leak the user-defined virtual router routes to the global virtual router based on specified Route Map. Applicable only for BGP in VRF context.",
				Computed:            true,
			},
			"ipv4_import_user_vrf_route_map_id": schema.StringAttribute{
				MarkdownDescription: "Filter the routes at the destination virtual router. Applicable only for BGP in VRF context.",
				Computed:            true,
			},
			"ipv4_export_user_vrf_route_map_id": schema.StringAttribute{
				MarkdownDescription: "Filter the routes at the source virtual router before the routes are exported to other virtual routers. Applicable only for BGP in VRF context.",
				Computed:            true,
			},
		},
	}
}
func (d *DeviceBGPDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("as_number"),
		),
	}
}

func (d *DeviceBGPDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *DeviceBGPDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DeviceBGP

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
	if config.Id.IsNull() && !config.AsNumber.IsNull() {
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
					if config.AsNumber.ValueString() == v.Get("asNumber").String() {
						config.Id = types.StringValue(v.Get("id").String())
						tflog.Debug(ctx, fmt.Sprintf("%s: Found object with as_number '%v', id: %v", config.Id.ValueString(), config.AsNumber.ValueString(), config.Id.ValueString()))
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
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with as_number: %v", config.AsNumber.ValueString()))
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
