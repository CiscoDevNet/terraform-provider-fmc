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
	"github.com/google/uuid"
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
	_ resource.Resource                = &DeviceOSPFResource{}
	_ resource.ResourceWithImportState = &DeviceOSPFResource{}
)

func NewDeviceOSPFResource() resource.Resource {
	return &DeviceOSPFResource{}
}

type DeviceOSPFResource struct {
	client *fmc.Client
}

func (r *DeviceOSPFResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_ospf"
}

func (r *DeviceOSPFResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages a Device OSPF.").AddMinimumVersionHeaderDescription().AddMinimumVersionDescription("7.6").String,

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
			"vrf_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the parent VRF.").String,
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
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this is always 'OspfRoute'").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"process_id": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("OSPF process ID. The numbers 1 and 2 are reserved for the OSPF Process IDs in global VRF. The next two numbers, 3 and 4, are allocated to the two OSPF Process IDs in the first user-defined VRFs. This incremental pattern continues whenever OSPF is enabled in the next user-defined VRF.").AddIntegerRangeDescription(1, 300).String,
				Required:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 300),
				},
			},
			"router_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv4 address used as the router ID. Do not configure for AUTOMATIC router ID selection.").String,
				Optional:            true,
			},
			"rfc_1583_compatible": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable RFC 1583 compatibility as the method used to calculate summary route costs.").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"log_adjacency_changes": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Log adjacency changes.").AddStringEnumDescription("DEFAULT", "DETAILED").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("DEFAULT", "DETAILED"),
				},
			},
			"ignore_lsa_mospf": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Suppresses syslog messages when the route receives unsupported LSA Type 6 multicast OSPF (MOSPF) packets.").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"administrative_distance_inter_area": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Administrative distance for inter-area routes.").AddIntegerRangeDescription(1, 254).AddDefaultValueDescription("110").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 254),
				},
				Default: int64default.StaticInt64(110),
			},
			"administrative_distance_intra_area": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Administrative distance for intra-area routes.").AddIntegerRangeDescription(1, 254).AddDefaultValueDescription("110").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 254),
				},
				Default: int64default.StaticInt64(110),
			},
			"administrative_distance_external": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Administrative distance for external routes.").AddIntegerRangeDescription(1, 254).AddDefaultValueDescription("110").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 254),
				},
				Default: int64default.StaticInt64(110),
			},
			"timer_lsa_group": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interval in seconds at which LSAs are collected into a group and refreshed, check summed, or aged.").AddIntegerRangeDescription(10, 1800).AddDefaultValueDescription("240").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(10, 1800),
				},
				Default: int64default.StaticInt64(240),
			},
			"default_route_always_advertise": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Always advertise default route. When configure to any value, enables Default Information Originate as well.").String,
				Optional:            true,
			},
			"default_route_metric": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Metric for the default route.").AddIntegerRangeDescription(0, 16777214).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 16777214),
				},
			},
			"default_route_metric_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Metric type for the default route.").AddStringEnumDescription("TYPE_1", "TYPE_2").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("TYPE_1", "TYPE_2"),
				},
			},
			"default_route_route_map_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Route Map ID for choosing the process that generates the default route.").String,
				Optional:            true,
			},
			"non_stop_forwarding": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Non-Stop Forwarding (NSF).").String,
				Optional:            true,
			},
			"non_stop_forwarding_mechanism": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Non-Stop Forwarding mechanism.").AddStringEnumDescription("CISCO", "IETF", "NONE").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("CISCO", "IETF", "NONE"),
				},
			},
			"non_stop_forwarding_helper_mode": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Non-Stop Forwarding helper mode.").String,
				Optional:            true,
			},
			"non_stop_forwarding_restart_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Length of graceful restart interval (seconds).").AddIntegerRangeDescription(1, 1800).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 1800),
				},
			},
			"non_stop_forwarding_capability": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Non-Stop Forwarding capability.").String,
				Optional:            true,
			},
			"non_stop_forwarding_strict_mode": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IETF Strict Link State advertisement checking or Cisco Cancel NSF restart when non-NSF-aware neighboring networking devices are detected").String,
				Optional:            true,
			},
			"areas": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("OSPF areas.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Area ID in either Integer or IPv4 format.").String,
							Required:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Area type.").AddStringEnumDescription("normal", "stub", "nssa").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("normal", "stub", "nssa"),
							},
						},
						"no_summary": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("No Summary flag for Stub / NSSA areas.").String,
							Optional:            true,
						},
						"no_redistribution": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("No Redistribution flag for NSSA areas.").String,
							Optional:            true,
						},
						"default_route_metric_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Default route metric type for NSSA areas.").AddStringEnumDescription("TYPE_1", "TYPE_2").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("TYPE_1", "TYPE_2"),
							},
						},
						"default_route_metric": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Default route metric value for NSSA areas.").AddIntegerRangeDescription(0, 16777214).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 16777214),
							},
						},
						"networks": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Area networks.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Network object ID.").String,
										Required:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Network object name.").String,
										Optional:            true,
									},
								},
							},
						},
						"authentication": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Authentication type.").AddStringEnumDescription("PASSWORD", "MESSAGE_DIGEST").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("PASSWORD", "MESSAGE_DIGEST"),
							},
						},
						"default_cost": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Default cost for the area.").AddIntegerRangeDescription(0, 65535).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65535),
							},
						},
						"ranges": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Ranges.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"network_object_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Network object ID.").String,
										Required:            true,
									},
									"advertise": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Whether to advertise this area range.").String,
										Optional:            true,
									},
								},
							},
						},
						"virtual_links": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Virtual links.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"peer_router_host_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Host object Id representing the peer router.").String,
										Required:            true,
									},
									"hello_interval": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Hello interval in seconds.").AddIntegerRangeDescription(1, 8192).AddDefaultValueDescription("10").String,
										Optional:            true,
										Computed:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 8192),
										},
										Default: int64default.StaticInt64(10),
									},
									"transmit_delay": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Transmit delay in seconds.").AddIntegerRangeDescription(1, 8192).AddDefaultValueDescription("1").String,
										Optional:            true,
										Computed:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 8192),
										},
										Default: int64default.StaticInt64(1),
									},
									"retransmit_interval": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Retransmit interval in seconds.").AddIntegerRangeDescription(1, 8192).AddDefaultValueDescription("5").String,
										Optional:            true,
										Computed:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 8192),
										},
										Default: int64default.StaticInt64(5),
									},
									"dead_interval": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Dead interval in seconds.").AddIntegerRangeDescription(1, 8192).AddDefaultValueDescription("40").String,
										Optional:            true,
										Computed:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 8192),
										},
										Default: int64default.StaticInt64(40),
									},
									"authentication_password": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Password for authentication.").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.LengthBetween(0, 8),
										},
										Sensitive: true,
									},
									"authentication_area_password": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Password for authentication.").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.LengthBetween(0, 8),
										},
										Sensitive: true,
									},
									"authentication_area_md5s": schema.ListNestedAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("List of MD5 authentication keys.").String,
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"id": schema.Int64Attribute{
													MarkdownDescription: helpers.NewAttributeDescription("Key ID for the MD5 authentication key.").String,
													Required:            true,
												},
												"key": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("MD5 authentication key.").String,
													Required:            true,
													Sensitive:           true,
												},
											},
										},
									},
									"authentication_md5s": schema.ListNestedAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("List of MD5 authentication keys.").String,
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"id": schema.Int64Attribute{
													MarkdownDescription: helpers.NewAttributeDescription("Key ID for the MD5 authentication key.").String,
													Required:            true,
												},
												"key": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("MD5 authentication key.").String,
													Required:            true,
													Sensitive:           true,
												},
											},
										},
									},
									"authentication_key_chain_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Key chain object ID for authentication.").String,
										Optional:            true,
									},
								},
							},
						},
						"inter_area_filters": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Inter-area filters.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"prefix_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Prefix list object ID.").String,
										Required:            true,
									},
									"prefix_list_name": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Prefix list object name.").String,
										Optional:            true,
									},
									"filter_direction": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Filter direction.").AddStringEnumDescription("IN", "OUT").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("IN", "OUT"),
										},
									},
								},
							},
						},
					},
				},
			},
			"redistributions": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable protocol redistribution.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"route_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Protocol to redistribute.").AddStringEnumDescription("RedistributeConnected", "RedistributeStatic", "RedistributeOSPF", "RedistributeBGP", "RedistributeRIP", "RedistributeEIGRP").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("RedistributeConnected", "RedistributeStatic", "RedistributeOSPF", "RedistributeBGP", "RedistributeRIP", "RedistributeEIGRP"),
							},
						},
						"as_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Autonomous System Number (ASN) for BGP / EIGRP redistribution.").String,
							Optional:            true,
						},
						"ospf_match_external_1": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether to match external type 1 routes.").String,
							Optional:            true,
						},
						"ospf_match_external_2": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether to match external type 2 routes.").String,
							Optional:            true,
						},
						"ospf_match_internal": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether to match internal routes.").String,
							Optional:            true,
						},
						"ospf_match_nssa_external_1": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether to match NSSA external type 1 routes.").String,
							Optional:            true,
						},
						"ospf_match_nssa_external_2": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether to match NSSA external type 2 routes.").String,
							Optional:            true,
						},
						"process_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("OSPF process ID.").String,
							Optional:            true,
						},
						"subnets": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether to redistribute subnets.").String,
							Optional:            true,
						},
						"metric": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Metric value for the routes being distributed.").AddIntegerRangeDescription(0, 16777214).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 16777214),
							},
						},
						"metric_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Metric type for the default route.").AddStringEnumDescription("TYPE_1", "TYPE_2").AddDefaultValueDescription("TYPE_2").String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("TYPE_1", "TYPE_2"),
							},
							Default: stringdefault.StaticString("TYPE_2"),
						},
						"tag": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Tag number.").AddIntegerRangeDescription(0, 4294967295).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 4294967295),
							},
						},
						"route_map_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Route map ID for route filtering.").String,
							Optional:            true,
						},
					},
				},
			},
			"filter_rules": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Filter prefix advertisement between areas.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"access_list_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Standard Access List ID").String,
							Required:            true,
						},
						"traffic_direction": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Filter rule direction.").AddStringEnumDescription("incomingroutefilter", "outgoingroutefilter").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("incomingroutefilter", "outgoingroutefilter"),
							},
						},
						"routing_process": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Protocol for the filter rule. Applicable only for `outgoingroutefilter` direction.").AddStringEnumDescription("CONNECTED", "STATIC", "RIP", "OSPF", "EIGRP", "NONE", "BGP").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("CONNECTED", "STATIC", "RIP", "OSPF", "EIGRP", "NONE", "BGP"),
							},
						},
						"routing_process_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Routing process ID for the filter rule. Applicable for OSPF, BGP, and EIGRP protocols.").String,
							Optional:            true,
						},
						"interface_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Interface ID for the filter rule. Applicable only for `incomingroutefilter` direction.").String,
							Optional:            true,
						},
					},
				},
			},
			"summary_addresses": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Addresses summarization configuration.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"networks": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Summary Networks.").String,
							Required:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Network object ID.").String,
										Required:            true,
									},
								},
							},
						},
						"tag": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Tag number.").AddIntegerRangeDescription(0, 4294967295).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 4294967295),
							},
						},
						"advertise": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether to advertise this summary route.").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *DeviceOSPFResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *DeviceOSPFResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	// Check if FMC client is connected to supports this object
	if r.client.FMCVersionParsed.LessThan(minFMCVersionDeviceOSPF) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support Device OSPF creation, minumum required version is 7.6", r.client.FMCVersion))
		return
	}
	var plan DeviceOSPF

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
	body := plan.toBody(ctx, DeviceOSPF{})
	body = plan.adjustBody(ctx, body)
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

func (r *DeviceOSPFResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Check if FMC client is connected to supports this object
	if r.client.FMCVersionParsed.LessThan(minFMCVersionDeviceOSPF) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support Device OSPF, minimum required version is 7.6", r.client.FMCVersion))
		return
	}
	var state DeviceOSPF

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

func (r *DeviceOSPFResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state DeviceOSPF

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
	body = plan.adjustBody(ctx, body)
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

func (r *DeviceOSPFResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DeviceOSPF

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
func (r *DeviceOSPFResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	errMsg := "Failed to parse import parameters.\nPlease provide import string in the following format: <domain>,<device_id>,<vrf_id>,<id>\n<domain> is optional. If not provided, `Global` is used implicitly and resource's `domain` attribute is not set.\n<vrf_id> is optional.\n" + fmt.Sprintf("Got: %q", req.ID)
	parts := strings.Split(req.ID, ",")
	if len(parts) < 2 || len(parts) > 4 {
		resp.Diagnostics.AddError("Import error", errMsg)
		return
	}

	for i := range parts {
		if parts[i] == "" {
			resp.Diagnostics.AddError("Import error", errMsg)
			return
		}
	}

	if len(parts) == 2 {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("device_id"), parts[0])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), parts[1])...)
	} else if len(parts) == 3 {
		if err := uuid.Validate(parts[0]); err == nil {
			// First part is UUID, so it's device_id
			resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("device_id"), parts[0])...)
			resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("vrf_id"), parts[1])...)
		} else {
			// First part is domain
			resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("domain"), parts[0])...)
			resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("device_id"), parts[1])...)
		}
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), parts[2])...)

	} else if len(parts) == 4 {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("domain"), parts[0])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("device_id"), parts[1])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("vrf_id"), parts[2])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), parts[3])...)
	}

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin createSubresources

// End of section. //template:end createSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin deleteSubresources

// End of section. //template:end deleteSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin updateSubresources

// End of section. //template:end updateSubresources
