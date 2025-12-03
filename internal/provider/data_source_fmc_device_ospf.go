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
	_ datasource.DataSource              = &DeviceOSPFDataSource{}
	_ datasource.DataSourceWithConfigure = &DeviceOSPFDataSource{}
)

func NewDeviceOSPFDataSource() datasource.DataSource {
	return &DeviceOSPFDataSource{}
}

type DeviceOSPFDataSource struct {
	client *fmc.Client
}

func (d *DeviceOSPFDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_ospf"
}

func (d *DeviceOSPFDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the Device OSPF.").AddMinimumVersionHeaderDescription().AddMinimumVersionDescription("7.6").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the object",
				Required:            true,
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
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the object; this is always 'OspfRoute'",
				Computed:            true,
			},
			"process_id": schema.Int64Attribute{
				MarkdownDescription: "OSPF process ID. The numbers 1 and 2 are reserved for the OSPF Process IDs in global VRF. The next two numbers, 3 and 4, are allocated to the two OSPF Process IDs in the first user-defined VRFs. This incremental pattern continues whenever OSPF is enabled in the next user-defined VRF.",
				Computed:            true,
			},
			"router_id": schema.StringAttribute{
				MarkdownDescription: "IPv4 address used as the router ID. Do not configure for AUTOMATIC router ID selection.",
				Computed:            true,
			},
			"rfc_1583_compatible": schema.BoolAttribute{
				MarkdownDescription: "Enable RFC 1583 compatibility as the method used to calculate summary route costs.",
				Computed:            true,
			},
			"log_adjacency_changes": schema.StringAttribute{
				MarkdownDescription: "Log adjacency changes.",
				Computed:            true,
			},
			"ignore_lsa_mospf": schema.BoolAttribute{
				MarkdownDescription: "Suppresses syslog messages when the route receives unsupported LSA Type 6 multicast OSPF (MOSPF) packets.",
				Computed:            true,
			},
			"administrative_distance_inter_area": schema.Int64Attribute{
				MarkdownDescription: "Administrative distance for inter-area routes.",
				Computed:            true,
			},
			"administrative_distance_intra_area": schema.Int64Attribute{
				MarkdownDescription: "Administrative distance for intra-area routes.",
				Computed:            true,
			},
			"administrative_distance_external": schema.Int64Attribute{
				MarkdownDescription: "Administrative distance for external routes.",
				Computed:            true,
			},
			"timer_lsa_group": schema.Int64Attribute{
				MarkdownDescription: "Interval in seconds at which LSAs are collected into a group and refreshed, check summed, or aged.",
				Computed:            true,
			},
			"default_route_always_advertise": schema.BoolAttribute{
				MarkdownDescription: "Always advertise default route. When configure to any value, enables Default Information Originate as well.",
				Computed:            true,
			},
			"default_route_metric": schema.Int64Attribute{
				MarkdownDescription: "Metric for the default route.",
				Computed:            true,
			},
			"default_route_metric_type": schema.StringAttribute{
				MarkdownDescription: "Metric type for the default route.",
				Computed:            true,
			},
			"default_route_route_map_id": schema.StringAttribute{
				MarkdownDescription: "Route Map ID for choosing the process that generates the default route.",
				Computed:            true,
			},
			"non_stop_forwarding": schema.BoolAttribute{
				MarkdownDescription: "Enable Non-Stop Forwarding (NSF).",
				Computed:            true,
			},
			"non_stop_forwarding_mechanism": schema.StringAttribute{
				MarkdownDescription: "Non-Stop Forwarding mechanism.",
				Computed:            true,
			},
			"non_stop_forwarding_helper_mode": schema.BoolAttribute{
				MarkdownDescription: "Enable Non-Stop Forwarding helper mode.",
				Computed:            true,
			},
			"non_stop_forwarding_restart_interval": schema.Int64Attribute{
				MarkdownDescription: "Length of graceful restart interval (seconds).",
				Computed:            true,
			},
			"non_stop_forwarding_capability": schema.BoolAttribute{
				MarkdownDescription: "Enable Non-Stop Forwarding capability.",
				Computed:            true,
			},
			"non_stop_forwarding_strict_mode": schema.BoolAttribute{
				MarkdownDescription: "IETF Strict Link State advertisement checking or Cisco Cancel NSF restart when non-NSF-aware neighboring networking devices are detected",
				Computed:            true,
			},
			"areas": schema.ListNestedAttribute{
				MarkdownDescription: "OSPF areas.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Area ID in either Integer or IPv4 format.",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Area type.",
							Computed:            true,
						},
						"no_summary": schema.BoolAttribute{
							MarkdownDescription: "No Summary flag for Stub / NSSA areas.",
							Computed:            true,
						},
						"no_redistribution": schema.BoolAttribute{
							MarkdownDescription: "No Redistribution flag for NSSA areas.",
							Computed:            true,
						},
						"default_route_metric_type": schema.StringAttribute{
							MarkdownDescription: "Default route metric type for NSSA areas.",
							Computed:            true,
						},
						"default_route_metric": schema.Int64Attribute{
							MarkdownDescription: "Default route metric value for NSSA areas.",
							Computed:            true,
						},
						"networks": schema.ListNestedAttribute{
							MarkdownDescription: "Area networks.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Network object ID.",
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: "Network object name.",
										Computed:            true,
									},
								},
							},
						},
						"authentication": schema.StringAttribute{
							MarkdownDescription: "Authentication type.",
							Computed:            true,
						},
						"default_cost": schema.Int64Attribute{
							MarkdownDescription: "Default cost for the area.",
							Computed:            true,
						},
						"ranges": schema.ListNestedAttribute{
							MarkdownDescription: "Ranges.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"network_object_id": schema.StringAttribute{
										MarkdownDescription: "Network object ID.",
										Computed:            true,
									},
									"advertise": schema.BoolAttribute{
										MarkdownDescription: "Whether to advertise this area range.",
										Computed:            true,
									},
								},
							},
						},
						"virtual_links": schema.ListNestedAttribute{
							MarkdownDescription: "Virtual links.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"peer_router_host_id": schema.StringAttribute{
										MarkdownDescription: "Host object Id representing the peer router.",
										Computed:            true,
									},
									"hello_interval": schema.Int64Attribute{
										MarkdownDescription: "Hello interval in seconds.",
										Computed:            true,
									},
									"transmit_delay": schema.Int64Attribute{
										MarkdownDescription: "Transmit delay in seconds.",
										Computed:            true,
									},
									"retransmit_interval": schema.Int64Attribute{
										MarkdownDescription: "Retransmit interval in seconds.",
										Computed:            true,
									},
									"dead_interval": schema.Int64Attribute{
										MarkdownDescription: "Dead interval in seconds.",
										Computed:            true,
									},
									"authentication_password": schema.StringAttribute{
										MarkdownDescription: "Password for authentication.",
										Computed:            true,
										Sensitive:           true,
									},
									"authentication_area_password": schema.StringAttribute{
										MarkdownDescription: "Password for authentication.",
										Computed:            true,
										Sensitive:           true,
									},
									"authentication_area_md5s": schema.ListNestedAttribute{
										MarkdownDescription: "List of MD5 authentication keys.",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"id": schema.Int64Attribute{
													MarkdownDescription: "Key ID for the MD5 authentication key.",
													Computed:            true,
												},
												"key": schema.StringAttribute{
													MarkdownDescription: "MD5 authentication key.",
													Computed:            true,
													Sensitive:           true,
												},
											},
										},
									},
									"authentication_md5s": schema.ListNestedAttribute{
										MarkdownDescription: "List of MD5 authentication keys.",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"id": schema.Int64Attribute{
													MarkdownDescription: "Key ID for the MD5 authentication key.",
													Computed:            true,
												},
												"key": schema.StringAttribute{
													MarkdownDescription: "MD5 authentication key.",
													Computed:            true,
													Sensitive:           true,
												},
											},
										},
									},
									"authentication_key_chain_id": schema.StringAttribute{
										MarkdownDescription: "Key chain object ID for authentication.",
										Computed:            true,
									},
								},
							},
						},
						"inter_area_filters": schema.ListNestedAttribute{
							MarkdownDescription: "Inter-area filters.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"prefix_list_id": schema.StringAttribute{
										MarkdownDescription: "Prefix list object ID.",
										Computed:            true,
									},
									"prefix_list_name": schema.StringAttribute{
										MarkdownDescription: "Prefix list object name.",
										Computed:            true,
									},
									"filter_direction": schema.StringAttribute{
										MarkdownDescription: "Filter direction.",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
			"redistributions": schema.ListNestedAttribute{
				MarkdownDescription: "Enable protocol redistribution.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"route_type": schema.StringAttribute{
							MarkdownDescription: "Protocol to redistribute.",
							Computed:            true,
						},
						"as_number": schema.Int64Attribute{
							MarkdownDescription: "Autonomous System Number (ASN) for BGP / EIGRP redistribution.",
							Computed:            true,
						},
						"ospf_match_external_1": schema.BoolAttribute{
							MarkdownDescription: "Whether to match external type 1 routes.",
							Computed:            true,
						},
						"ospf_match_external_2": schema.BoolAttribute{
							MarkdownDescription: "Whether to match external type 2 routes.",
							Computed:            true,
						},
						"ospf_match_internal": schema.BoolAttribute{
							MarkdownDescription: "Whether to match internal routes.",
							Computed:            true,
						},
						"ospf_match_nssa_external_1": schema.BoolAttribute{
							MarkdownDescription: "Whether to match NSSA external type 1 routes.",
							Computed:            true,
						},
						"ospf_match_nssa_external_2": schema.BoolAttribute{
							MarkdownDescription: "Whether to match NSSA external type 2 routes.",
							Computed:            true,
						},
						"process_id": schema.Int64Attribute{
							MarkdownDescription: "OSPF process ID.",
							Computed:            true,
						},
						"subnets": schema.BoolAttribute{
							MarkdownDescription: "Whether to redistribute subnets.",
							Computed:            true,
						},
						"metric": schema.Int64Attribute{
							MarkdownDescription: "Metric value for the routes being distributed.",
							Computed:            true,
						},
						"metric_type": schema.StringAttribute{
							MarkdownDescription: "Metric type for the default route.",
							Computed:            true,
						},
						"tag": schema.Int64Attribute{
							MarkdownDescription: "Tag number.",
							Computed:            true,
						},
						"route_map_id": schema.StringAttribute{
							MarkdownDescription: "Route map ID for route filtering.",
							Computed:            true,
						},
					},
				},
			},
			"filter_rules": schema.ListNestedAttribute{
				MarkdownDescription: "Filter prefix advertisement between areas.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"access_list_id": schema.StringAttribute{
							MarkdownDescription: "Standard Access List ID",
							Computed:            true,
						},
						"traffic_direction": schema.StringAttribute{
							MarkdownDescription: "Filter rule direction.",
							Computed:            true,
						},
						"routing_process": schema.StringAttribute{
							MarkdownDescription: "Protocol for the filter rule. Applicable only for `outgoingroutefilter` direction.",
							Computed:            true,
						},
						"routing_process_id": schema.Int64Attribute{
							MarkdownDescription: "Routing process ID for the filter rule. Applicable for OSPF, BGP, and EIGRP protocols.",
							Computed:            true,
						},
						"interface_id": schema.StringAttribute{
							MarkdownDescription: "Interface ID for the filter rule. Applicable only for `incomingroutefilter` direction.",
							Computed:            true,
						},
					},
				},
			},
			"summary_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "Addresses summarization configuration.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"networks": schema.ListNestedAttribute{
							MarkdownDescription: "Summary Networks.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Network object ID.",
										Computed:            true,
									},
								},
							},
						},
						"tag": schema.Int64Attribute{
							MarkdownDescription: "Tag number.",
							Computed:            true,
						},
						"advertise": schema.BoolAttribute{
							MarkdownDescription: "Whether to advertise this summary route.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *DeviceOSPFDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *DeviceOSPFDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	// Check if FMC client is connected to supports this object
	if d.client.FMCVersionParsed.LessThan(minFMCVersionDeviceOSPF) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support Device OSPF, minimum required version is 7.6", d.client.FMCVersion))
		return
	}
	var config DeviceOSPF

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
