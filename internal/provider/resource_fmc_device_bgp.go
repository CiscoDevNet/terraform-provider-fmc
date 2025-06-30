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
	"strings"

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &DeviceBGPResource{}
	_ resource.ResourceWithImportState = &DeviceBGPResource{}
)

func NewDeviceBGPResource() resource.Resource {
	return &DeviceBGPResource{}
}

type DeviceBGPResource struct {
	client *fmc.Client
}

func (r *DeviceBGPResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_bgp"
}

func (r *DeviceBGPResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resouce manages Device BGP Settings. As pre-requisite, BGP has to be enabled and AS Number assigned under BGP General Settings (fmc_device_bgp_general_settings).").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "Name of the FMC domain",
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"device_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the parent device.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the object; this is always 'bgp'").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this is always 'bgp'").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"as_number": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Autonomus System (AS) number").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"ipv4_address_family_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv4 Address Family Id").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"ipv4_address_family_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv4 Address Family Type").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"ipv4_learned_route_map_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Learned Route Map ID").String,
				Optional:            true,
			},
			"ipv4_default_information_orginate": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Generate default route").String,
				Optional:            true,
			},
			"ipv4_auto_summary": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Summarize subnet routes into network level routes").String,
				Optional:            true,
			},
			"ipv4_suppress_inactive_routes": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Suppressing advertisement of inactive routes").String,
				Optional:            true,
			},
			"ipv4_synchronization": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Synchronize between BGP and IGP systems").String,
				Optional:            true,
			},
			"ipv4_redistribute_ibgp_into_igp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Redistribute IBGP into IGP. Use filtering to limit the number of prefixes that are redistributed.").String,
				Optional:            true,
			},
			"ipv4_external_distance": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Administrative route distance for external routes").AddIntegerRangeDescription(1, 255).AddDefaultValueDescription("20").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
				Default: int64default.StaticInt64(20),
			},
			"ipv4_internal_distance": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Administrative route distance for internal routes").AddIntegerRangeDescription(1, 255).AddDefaultValueDescription("200").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
				Default: int64default.StaticInt64(200),
			},
			"ipv4_local_distance": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Administrative route distance for local routes").AddIntegerRangeDescription(1, 255).AddDefaultValueDescription("200").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
				Default: int64default.StaticInt64(200),
			},
			"ipv4_number_of_ibgp_paths": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Number of paths to use for IBGP").AddIntegerRangeDescription(1, 8).AddDefaultValueDescription("1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 8),
				},
				Default: int64default.StaticInt64(1),
			},
			"ipv4_number_of_ebgp_paths": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Number of paths to use for EBGP").AddIntegerRangeDescription(1, 8).AddDefaultValueDescription("1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 8),
				},
				Default: int64default.StaticInt64(1),
			},
			"ipv4_neighbors": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP address of the BGP neighbor").String,
							Optional:            true,
						},
						"remote_as": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("AS number of the BGP neighbor").String,
							Optional:            true,
						},
						"bfd_fallover": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("BFD Fallover").AddStringEnumDescription("SINGLE_HOP", "MULTI_HOP", "AUTO_DETECT_HOP", "NONE").AddDefaultValueDescription("NONE").String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("SINGLE_HOP", "MULTI_HOP", "AUTO_DETECT_HOP", "NONE"),
							},
							Default: stringdefault.StaticString("NONE"),
						},
						"update_source_interface_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Interface ID for the update source").String,
							Optional:            true,
						},
						"enable_address": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable communication with this BGP neighbor").AddDefaultValueDescription("false").String,
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"as_override": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable overriding of the AS number of the originating router with the AS number of the sending BGP router.").String,
							Optional:            true,
						},
						"graceful_restart": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable graceful restart for the neighbor.").String,
							Optional:            true,
						},
						"shutdown_administratively": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Disable a neighbor or peer group").AddDefaultValueDescription("false").String,
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"description": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Description of the neighbor").String,
							Optional:            true,
						},
						"filter_access_lists": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set incoming or outgoing Access List to distribute BGP neighbor information.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"access_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Access List ID").String,
										Required:            true,
									},
									"update_direction": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Filter direction").AddStringEnumDescription("IN", "OUT").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("IN", "OUT"),
										},
									},
								},
							},
						},
						"filter_route_maps": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set incoming or outgoing Route Maps to apply a route map to incoming or outgoing routes.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"route_map_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Route Map ID").String,
										Required:            true,
									},
									"update_direction": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Filter direction").AddStringEnumDescription("IN", "OUT").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("IN", "OUT"),
										},
									},
								},
							},
						},
						"filter_prefix_lists": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set incoming or outgoing Prefix List to distribute BGP neighbor information.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"prefix_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Prefix List ID").String,
										Required:            true,
									},
									"update_direction": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Filter direction").AddStringEnumDescription("IN", "OUT").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("IN", "OUT"),
										},
									},
								},
							},
						},
						"filter_as_paths": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set incoming or outgoing AS path filter to distribute BGP neighbor information.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"as_path_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("AS Path ID").String,
										Required:            true,
									},
									"as_path_name": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("AS Path Name").String,
										Required:            true,
									},
									"update_direction": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Filter direction").AddStringEnumDescription("IN", "OUT").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("IN", "OUT"),
										},
									},
								},
							},
						},
						"filter_maximum_prefixes": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Maximum number of prefixes allowed from the neighbor").AddIntegerRangeDescription(1, 2147483647).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 2147483647),
							},
						},
						"filter_warning_only": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Give only warning message when prefix limit exceeded. Can be set to `true` only. Use `neighbor_filter_threshold_value` to set mode to session termination.").String,
							Optional:            true,
						},
						"filter_threshold_value": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Threshold value for the maximum number of prefixes allowed from the neighbor (implies session termination).").AddIntegerRangeDescription(1, 100).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 100),
							},
						},
						"filter_restart_interval": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Time interval to restart the maximum prefix limit in Minutes").AddIntegerRangeDescription(1, 65535).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65535),
							},
						},
						"routes_advertisement_interval": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Time interval to advertise routes in seconds").AddIntegerRangeDescription(0, 600).AddDefaultValueDescription("0").String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 600),
							},
							Default: int64default.StaticInt64(0),
						},
						"routes_remove_private_as": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Remove private AS numbers from outgoing routing updates").AddDefaultValueDescription("false").String,
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"routes_generate_default_route_map_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Generate default routes - Route Map").String,
							Optional:            true,
						},
						"routes_advertise_maps": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Define conditionally advertised routes.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"advertise_map_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Specified route maps are advertised when conditions of the exist map or the non-exist map are met.").String,
										Required:            true,
									},
									"use_exist_map": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set mode to Exist Map (true) or Non-Exist Map (false).").String,
										Required:            true,
									},
									"exist_nonexist_map_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Specify exist / non-exist route map ID.").String,
										Required:            true,
									},
								},
							},
						},
						"keepalive_interval": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Time interval to send keepalive messages in seconds").AddIntegerRangeDescription(0, 65535).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65535),
							},
						},
						"hold_time": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Time interval to hold the neighbor in seconds").AddIntegerRangeDescription(3, 65535).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(3, 65535),
							},
						},
						"minimum_hold_time": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Minimum hold time in seconds").AddIntegerRangeDescription(3, 65535).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(3, 65535),
							},
						},
						"authentication_password": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Setting password enables authentication.").String,
							Optional:            true,
							Sensitive:           true,
						},
						"send_community_attribute": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Send Community attribute to this neighbor").AddDefaultValueDescription("false").String,
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"next_hop_self": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Use itself as next hop for this neighbor").AddDefaultValueDescription("false").String,
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"disable_connection_verification": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Disable Connection Verification").AddDefaultValueDescription("false").String,
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"tcp_path_mtu_discovery": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Use TCP path MTU discovery.").AddDefaultValueDescription("false").String,
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"max_hop_count": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Maximum number of hops to reach the neighbor").AddIntegerRangeDescription(1, 255).AddDefaultValueDescription("1").String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 255),
							},
							Default: int64default.StaticInt64(1),
						},
						"tcp_transport_mode": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("True set it to active, False to passive.").AddDefaultValueDescription("false").String,
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"weight": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Weight of the neighbor").AddIntegerRangeDescription(0, 65535).AddDefaultValueDescription("0").String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65535),
							},
							Default: int64default.StaticInt64(0),
						},
						"version": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set BPG version: 0 - default, 4 - IPv4").AddStringEnumDescription("0", "4").AddDefaultValueDescription("0").String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("0", "4"),
							},
							Default: stringdefault.StaticString("0"),
						},
						"customized_local_as_number": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Customize the AS number for the routes received from neighbor").String,
							Optional:            true,
						},
						"customized_no_prepend": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Do not prepend local AS number to routes received from neighbor").String,
							Optional:            true,
						},
						"customized_replace_as": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Replace real AS number with local AS number in routes received from neighbor").String,
							Optional:            true,
						},
						"customized_accept_both_as": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Accept either real AS number or local AS number in routes experienced from neighbor").String,
							Optional:            true,
						},
					},
				},
			},
			"ipv4_aggregate_addresses": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Generate aggregate address information for IPv4.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"generate_as": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Generate AS set path information").String,
							Optional:            true,
						},
						"filter": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Filter all routes from updates (summary only)").String,
							Optional:            true,
						},
						"network_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Network ID (desired network/host object)").String,
							Optional:            true,
						},
						"advertise_map_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Advertise Route Map ID (select the routes to create AS_SET origin communities)").String,
							Optional:            true,
						},
						"attribute_map_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Attribute Route Map ID (set the attribute of the aggregate route).").String,
							Optional:            true,
						},
						"suppress_map_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Suppress Route Map ID (select the routes to be suppressed).").String,
							Optional:            true,
						},
					},
				},
			},
			"ipv4_filterings": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Filter routes or networks received in incoming BGP updates").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"access_list_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Standard Access List ID that defines which networks are to be received and which are to be suppressed in routing updates.").String,
							Required:            true,
						},
						"direction": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Determine if the filter should be applied to inbound updates or outbound updates").AddStringEnumDescription("incomingroutefilter", "outgoingroutefilter").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("incomingroutefilter", "outgoingroutefilter"),
							},
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Routing process for which you want to filter").AddStringEnumDescription("CONNECTED", "BGP", "OSPF", "RIP", "STATIC").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("CONNECTED", "BGP", "OSPF", "RIP", "STATIC"),
							},
						},
						"process_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Process ID for the OSPF routing protocol.").String,
							Optional:            true,
						},
					},
				},
			},
			"ipv4_networks": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Add networks that will be advertised by the BGP routing process").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"network_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Network to be advertised by the BGP routing processes.").String,
							Required:            true,
						},
						"route_map_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Route Map ID that should be examined to filter the networks to be advertised").String,
							Optional:            true,
						},
					},
				},
			},
			"ipv4_redistributions": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Define the conditions for redistributing routes from another routing domain into BGP.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"source_protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Protocol to redistribute").AddStringEnumDescription("RedistributeConnected", "RedistributeStatic", "RedistributeOSPF", "RedistributeEIGRP").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("RedistributeConnected", "RedistributeStatic", "RedistributeOSPF", "RedistributeEIGRP"),
							},
						},
						"route_map_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Route Map ID to filter the networks to be redistributed").String,
							Optional:            true,
						},
						"metric": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Metric for the redistributed route.").AddIntegerRangeDescription(0, 4294967295).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 4294967295),
							},
						},
						"process_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("OSPF process ID").String,
							Optional:            true,
						},
						"match_external1": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match OSPF External 1 metrics").String,
							Optional:            true,
						},
						"match_external2": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match OSPF External 2 metrics").String,
							Optional:            true,
						},
						"match_internal": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match OSPF Internal metrics").String,
							Optional:            true,
						},
						"match_nssa_external1": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match OSPF NSSA External 1 metrics").String,
							Optional:            true,
						},
						"match_nssa_external2": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match OSPF NSSA External 2 metrics").String,
							Optional:            true,
						},
					},
				},
			},
			"ipv4_route_injections": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Define routes to be conditionally injected into the BGP routing table.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"inject_route_map_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Inject Route Map ID (prefixes to inject into the local BGP routing table)").String,
							Required:            true,
						},
						"exist_route_map_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Exist Route Map ID containing the prefixes that the BGP speaker will track").String,
							Required:            true,
						},
						"inherit_attributes": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Injected route will inherit the attributes of the aggregate route").AddDefaultValueDescription("true").String,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("true"),
						},
					},
				},
			},
		},
	}
}

func (r *DeviceBGPResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *DeviceBGPResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DeviceBGP

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !plan.Domain.IsNull() && plan.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(plan.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, DeviceBGP{})
	res, err := r.client.Post(plan.getPath(), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("id").String())
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *DeviceBGPResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DeviceBGP

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !state.Domain.IsNull() && state.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(state.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	urlPath := state.getPath() + "/" + url.QueryEscape(state.Id.ValueString())
	res, err := r.client.Get(urlPath, reqMods...)

	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	imp, diags := helpers.IsFlagImporting(ctx, req)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// After `terraform import` we switch to a full read.
	if imp {
		state.fromBody(ctx, res)
	} else {
		state.fromBodyPartial(ctx, res)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update

func (r *DeviceBGPResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state DeviceBGP

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Read state
	diags = req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !plan.Domain.IsNull() && plan.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(plan.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	body := plan.toBody(ctx, state)
	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *DeviceBGPResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DeviceBGP

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !state.Domain.IsNull() && state.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(state.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	res, err := r.client.Delete(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString()), reqMods...)
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import

func (r *DeviceBGPResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <device_id>,<id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("device_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin createSubresources

// End of section. //template:end createSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin deleteSubresources

// End of section. //template:end deleteSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin updateSubresources

// End of section. //template:end updateSubresources
