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
	_ datasource.DataSource              = &DevicePhysicalInterfaceDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicePhysicalInterfaceDataSource{}
)

func NewDevicePhysicalInterfaceDataSource() datasource.DataSource {
	return &DevicePhysicalInterfaceDataSource{}
}

type DevicePhysicalInterfaceDataSource struct {
	client *fmc.Client
}

func (d *DevicePhysicalInterfaceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_physical_interface"
}

func (d *DevicePhysicalInterfaceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Device Physical Interface.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Optional:            true,
				Computed:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "The name of the FMC domain",
				Optional:            true,
			},
			"device_id": schema.StringAttribute{
				MarkdownDescription: "UUID of the parent device (fmc_device.example.id).",
				Required:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Indicates whether to enable the interface.",
				Computed:            true,
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "Mode of the interface. Use INLINE if, and only if, the interface is part of fmc_inline_set with tap_mode=false or tap_mode unset. Use TAP if, and only if, the interface is part of fmc_inline_set with tap_mode = true. Use ERSPAN only when both erspan_source_ip and erspan_flow_id are set.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the interface; it must already be present on the device.",
				Optional:            true,
				Computed:            true,
			},
			"logical_name": schema.StringAttribute{
				MarkdownDescription: "Customizable logical name of the interface, unique on the device. Should not contain whitespace or slash characters. Must be non-empty in order to set security_zone_id, mtu, inline sets, etc.",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"management_only": schema.BoolAttribute{
				MarkdownDescription: "Indicates whether this interface limits traffic to management traffic; when true, through-the-box traffic is disallowed. Value true conflicts with mode INLINE, PASSIVE, TAP, ERSPAN, or with security_zone_id.",
				Computed:            true,
			},
			"mtu": schema.Int64Attribute{
				MarkdownDescription: "Maximum transmission unit. Can only be used when logical_name is set.",
				Computed:            true,
			},
			"priority": schema.Int64Attribute{
				MarkdownDescription: "Priority 0-65535. Can only be set for routed interfaces.",
				Computed:            true,
			},
			"security_zone_id": schema.StringAttribute{
				MarkdownDescription: "UUID of the assigned security zone (fmc_security_zone.example.id). Can only be used when logical_name is set.",
				Computed:            true,
			},
			"ipv4_static_address": schema.StringAttribute{
				MarkdownDescription: "Static IPv4 address. Conflicts with mode INLINE, PASSIVE, TAP, ERSPAN.",
				Computed:            true,
			},
			"ipv4_static_netmask": schema.StringAttribute{
				MarkdownDescription: "Netmask (width) for ipv4_static_address.",
				Computed:            true,
			},
			"ipv4_dhcp_obtain_route": schema.BoolAttribute{
				MarkdownDescription: "Any non-null value here indicates to enable DHCPv4. Value `false` indicates to enable DHCPv4 without obtaining from there the default IPv4 route but anyway requires also ipv4_dhcp_route_metric to be set to exactly 1. Value `true` indicates to enable DHCPv4 and obtain the route and also requires ipv4_dhcp_route_metric to be non-null. The ipv4_dhcp_obtain_route must be null when using ipv4_static_netmask.",
				Computed:            true,
			},
			"ipv4_dhcp_route_metric": schema.Int64Attribute{
				MarkdownDescription: "The metric for ipv4_dhcp_obtain_route. Any non-null value enables DHCP as a side effect. Must be null when using ipv4_static_netmask.",
				Computed:            true,
			},
			"ipv6_enable": schema.BoolAttribute{
				MarkdownDescription: "Indicates whether to enable IPv6.",
				Computed:            true,
			},
			"ipv6_enforce_eui": schema.BoolAttribute{
				MarkdownDescription: "Indicates whether to enforce IPv6 Extended Unique Identifier (EUI64 from RFC2373).",
				Computed:            true,
			},
			"ipv6_enable_auto_config": schema.BoolAttribute{
				MarkdownDescription: "Indicates whether to enable IPv6 autoconfiguration.",
				Computed:            true,
			},
			"ipv6_enable_dhcp_address": schema.BoolAttribute{
				MarkdownDescription: "Indicates whether to enable DHCPv6 for address config.",
				Computed:            true,
			},
			"ipv6_enable_dhcp_nonaddress": schema.BoolAttribute{
				MarkdownDescription: "Indicates whether to enable DHCPv6 for non-address config.",
				Computed:            true,
			},
			"ipv6_enable_ra": schema.BoolAttribute{
				MarkdownDescription: "Indicates whether to enable IPv6 router advertisement (RA).",
				Computed:            true,
			},
			"ipv6_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: "IPv6 address without a slash and prefix.",
							Computed:            true,
						},
						"prefix": schema.StringAttribute{
							MarkdownDescription: "Prefix width for the IPv6 address.",
							Computed:            true,
						},
						"enforce_eui": schema.BoolAttribute{
							MarkdownDescription: "Indicates whether to enforce IPv6 Extended Unique Identifier (EUI64 from RFC2373).",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}
func (d *DevicePhysicalInterfaceDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *DevicePhysicalInterfaceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *DevicePhysicalInterfaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DevicePhysicalInterface

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
			queryString := fmt.Sprintf("?limit=%d&offset=%d", limit, offset)
			res, err := d.client.Get(config.getPath()+queryString, reqMods...)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects, got error: %s", err))
				return
			}
			if value := res.Get("items"); len(value.Array()) > 0 {
				value.ForEach(func(k, v gjson.Result) bool {
					if config.Name.ValueString() == v.Get("name").String() {
						config.Id = types.StringValue(v.Get("id").String())
						tflog.Debug(ctx, fmt.Sprintf("%s: Found object with name '%v', id: %v", config.Id.String(), config.Name.ValueString(), config.Id.String()))
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
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with name: %s", config.Name.ValueString()))
			return
		}
	}

	res, err := d.client.Get(config.getPath()+"/"+url.QueryEscape(config.Id.ValueString()), reqMods...)
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
