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
	_ datasource.DataSource              = &DeviceEIGRPDataSource{}
	_ datasource.DataSourceWithConfigure = &DeviceEIGRPDataSource{}
)

func NewDeviceEIGRPDataSource() datasource.DataSource {
	return &DeviceEIGRPDataSource{}
}

type DeviceEIGRPDataSource struct {
	client *fmc.Client
}

func (d *DeviceEIGRPDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_eigrp"
}

func (d *DeviceEIGRPDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the Device EIGRP.").String,

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
			"device_id": schema.StringAttribute{
				MarkdownDescription: "Id of the parent device.",
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the object; this value is always 'EigrpRoute'.",
				Computed:            true,
			},
			"as_number": schema.StringAttribute{
				MarkdownDescription: "Autonomous System (AS) number of the EIGRP process.",
				Optional:            true,
				Computed:            true,
			},
			"networks": schema.ListNestedAttribute{
				MarkdownDescription: "Networks or Hosts that participate in the EIGRP routing process.",
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
			"auto_summary": schema.BoolAttribute{
				MarkdownDescription: "Advertise summary routes to other EIGRP routers.",
				Computed:            true,
			},
			"passive_interfaces_all": schema.BoolAttribute{
				MarkdownDescription: "Set all interfaces as passive (do not send or receive routing updates).",
				Computed:            true,
			},
			"passive_interfaces": schema.ListNestedAttribute{
				MarkdownDescription: "Interfaces as passive (do not send or receive routing updates). Applicable only when `passive_interfaces_all` is not set to true.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Interface ID.",
							Computed:            true,
						},
					},
				},
			},
			"neighbors": schema.ListNestedAttribute{
				MarkdownDescription: "EIGRP neighbors.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_id": schema.StringAttribute{
							MarkdownDescription: "Id of interface through which the neighbor is reachable.",
							Computed:            true,
						},
						"host_id": schema.StringAttribute{
							MarkdownDescription: "Id of Host object representing the neighbor.",
							Computed:            true,
						},
					},
				},
			},
			"filter_rules": schema.ListNestedAttribute{
				MarkdownDescription: "Filter rules (distribute lists) for EIGRP routing updates.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"filter_direction": schema.StringAttribute{
							MarkdownDescription: "Filter direction.",
							Computed:            true,
						},
						"standard_access_list_id": schema.StringAttribute{
							MarkdownDescription: "Standard Access List ID, that defines the networks that are to be received and suppressed in routing updates.",
							Computed:            true,
						},
						"interface_id": schema.StringAttribute{
							MarkdownDescription: "Id of interface to which the filter rule applies.",
							Computed:            true,
						},
						"routing_process": schema.StringAttribute{
							MarkdownDescription: "Routing process whose routes are filtered. Applicable only for `OUTBOUND` direction.",
							Computed:            true,
						},
						"routing_process_id": schema.StringAttribute{
							MarkdownDescription: "OSPF process ID or BGP AS number.",
							Computed:            true,
						},
					},
				},
			},
			"redistributions": schema.ListNestedAttribute{
				MarkdownDescription: "Redistributions to the EIGRP routing process.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"redistribute_protocol": schema.StringAttribute{
							MarkdownDescription: "Protocol to redistribute.",
							Computed:            true,
						},
						"process_id": schema.StringAttribute{
							MarkdownDescription: "OSPF process ID or BGP AS number.",
							Computed:            true,
						},
						"route_map_id": schema.StringAttribute{
							MarkdownDescription: "Route map ID for route filtering.",
							Computed:            true,
						},
						"match_internal": schema.BoolAttribute{
							MarkdownDescription: "Whether to match OSPF internal routes.",
							Computed:            true,
						},
						"match_external_1": schema.BoolAttribute{
							MarkdownDescription: "Whether to match OSPF external type 1 routes.",
							Computed:            true,
						},
						"match_external_2": schema.BoolAttribute{
							MarkdownDescription: "Whether to match OSPF external type 2 routes.",
							Computed:            true,
						},
						"match_nssa_external_1": schema.BoolAttribute{
							MarkdownDescription: "Whether to match OSPF NSSA external type 1 routes.",
							Computed:            true,
						},
						"match_nssa_external_2": schema.BoolAttribute{
							MarkdownDescription: "Whether to match OSPF NSSA external type 2 routes.",
							Computed:            true,
						},
						"metric_bandwidth": schema.Int64Attribute{
							MarkdownDescription: "Metric bandwidth (in Kbps).",
							Computed:            true,
						},
						"metric_delay": schema.Int64Attribute{
							MarkdownDescription: "Metric delay (in tens of microseconds).",
							Computed:            true,
						},
						"metric_reliability": schema.Int64Attribute{
							MarkdownDescription: "Metric reliability. The value 255 indicates 100 percent reliability; 0 means no reliability.",
							Computed:            true,
						},
						"metric_loading": schema.Int64Attribute{
							MarkdownDescription: "Metric effective bandwidth (loading). 255 indicates 100 percent loading.",
							Computed:            true,
						},
						"metric_mtu": schema.Int64Attribute{
							MarkdownDescription: "Metric path MTU (in bytes).",
							Computed:            true,
						},
					},
				},
			},
			"summary_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "Summary addresses advertised on interfaces.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_id": schema.StringAttribute{
							MarkdownDescription: "Id of interface on which the summary address is advertised.",
							Computed:            true,
						},
						"network_id": schema.StringAttribute{
							MarkdownDescription: "Network object ID of the summary address.",
							Computed:            true,
						},
						"administrative_distance": schema.Int64Attribute{
							MarkdownDescription: "Administrative distance of the summary route.",
							Computed:            true,
						},
					},
				},
			},
			"interfaces": schema.ListNestedAttribute{
				MarkdownDescription: "EIGRP interface settings.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_id": schema.StringAttribute{
							MarkdownDescription: "Interface ID.",
							Computed:            true,
						},
						"interface_type": schema.StringAttribute{
							MarkdownDescription: "Interface type.",
							Computed:            true,
						},
						"hello_interval": schema.Int64Attribute{
							MarkdownDescription: "Interval (in seconds) between hello transmissions.",
							Computed:            true,
						},
						"hold_time": schema.Int64Attribute{
							MarkdownDescription: "Time (in seconds) before the neighbor is considered down.",
							Computed:            true,
						},
						"split_horizon": schema.BoolAttribute{
							MarkdownDescription: "Enable split horizon.",
							Computed:            true,
						},
						"delay_time": schema.Int64Attribute{
							MarkdownDescription: "Throughput delay (in tens of microseconds).",
							Computed:            true,
						},
						"md5_authentication": schema.BoolAttribute{
							MarkdownDescription: "Enable MD5 authentication.",
							Computed:            true,
						},
						"authentication_key_type": schema.StringAttribute{
							MarkdownDescription: "Authentication key encryption.",
							Computed:            true,
						},
						"authentication_key_id": schema.Int64Attribute{
							MarkdownDescription: "Authentication key ID.",
							Computed:            true,
						},
						"authentication_key": schema.StringAttribute{
							MarkdownDescription: "Authentication key.",
							Computed:            true,
							Sensitive:           true,
						},
					},
				},
			},
			"router_id": schema.StringAttribute{
				MarkdownDescription: "IPv4 address used as the router ID. Do not configure for AUTOMATIC router ID selection.",
				Computed:            true,
			},
			"default_information_in_standard_access_list_id": schema.StringAttribute{
				MarkdownDescription: "Standard Access List ID used to accept exterior default routing information.",
				Computed:            true,
			},
			"default_information_out_standard_access_list_id": schema.StringAttribute{
				MarkdownDescription: "Standard Access List ID used to advertise exterior default routing information.",
				Computed:            true,
			},
			"administrative_distance_internal": schema.Int64Attribute{
				MarkdownDescription: "Administrative distance for EIGRP internal routes.",
				Computed:            true,
			},
			"administrative_distance_external": schema.Int64Attribute{
				MarkdownDescription: "Administrative distance for EIGRP external routes.",
				Computed:            true,
			},
			"log_neighbor_changes": schema.BoolAttribute{
				MarkdownDescription: "Enable logging of EIGRP neighbor adjacency changes.",
				Computed:            true,
			},
			"log_neighbor_warnings": schema.BoolAttribute{
				MarkdownDescription: "Enable logging of EIGRP neighbor warning messages.",
				Computed:            true,
			},
			"log_neighbor_warnings_interval": schema.Int64Attribute{
				MarkdownDescription: "Interval (in seconds) between repeated neighbor warning messages.",
				Computed:            true,
			},
			"stub_receive_only": schema.BoolAttribute{
				MarkdownDescription: "Configure the device as a receive-only EIGRP stub router.",
				Computed:            true,
			},
			"stub_advertise_connected": schema.BoolAttribute{
				MarkdownDescription: "EIGRP stub router advertises connected routes.",
				Computed:            true,
			},
			"stub_advertise_redistributed": schema.BoolAttribute{
				MarkdownDescription: "EIGRP stub router advertises redistributed routes.",
				Computed:            true,
			},
			"stub_advertise_static": schema.BoolAttribute{
				MarkdownDescription: "EIGRP stub router advertises static routes.",
				Computed:            true,
			},
			"stub_advertise_summary": schema.BoolAttribute{
				MarkdownDescription: "EIGRP stub router advertises summary routes.",
				Computed:            true,
			},
			"default_metric_bandwidth": schema.Int64Attribute{
				MarkdownDescription: "Default metric bandwidth (in Kbps) for redistributed routes.",
				Computed:            true,
			},
			"default_metric_delay_time": schema.Int64Attribute{
				MarkdownDescription: "Default metric delay (in tens of microseconds) for redistributed routes.",
				Computed:            true,
			},
			"default_metric_reliability": schema.Int64Attribute{
				MarkdownDescription: "Default metric reliability for redistributed routes.",
				Computed:            true,
			},
			"default_metric_loading": schema.Int64Attribute{
				MarkdownDescription: "Default metric effective bandwidth (loading) for redistributed routes.",
				Computed:            true,
			},
			"default_metric_mtu": schema.Int64Attribute{
				MarkdownDescription: "Default metric path MTU (in bytes) for redistributed routes.",
				Computed:            true,
			},
			"cluster_ipv4_pool_id": schema.StringAttribute{
				MarkdownDescription: "IPv4 Address Pool ID used for router ID assignment on clustered devices.",
				Computed:            true,
			},
		},
	}
}
func (d *DeviceEIGRPDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("as_number"),
		),
	}
}

func (d *DeviceEIGRPDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *DeviceEIGRPDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DeviceEIGRP

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
