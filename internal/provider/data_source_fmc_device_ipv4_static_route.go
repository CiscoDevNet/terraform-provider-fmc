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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &DeviceIPv4StaticRouteDataSource{}
	_ datasource.DataSourceWithConfigure = &DeviceIPv4StaticRouteDataSource{}
)

func NewDeviceIPv4StaticRouteDataSource() datasource.DataSource {
	return &DeviceIPv4StaticRouteDataSource{}
}

type DeviceIPv4StaticRouteDataSource struct {
	client *fmc.Client
}

func (d *DeviceIPv4StaticRouteDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_ipv4_static_route"
}

func (d *DeviceIPv4StaticRouteDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Device IPv4 Static Route.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "The name of the FMC domain",
				Optional:            true,
			},
			"device_id": schema.StringAttribute{
				MarkdownDescription: "UUID of the parent device (fmc_device.example.id).",
				Required:            true,
			},
			"interface_logical_name": schema.StringAttribute{
				MarkdownDescription: "Logical name of the parent interface (fmc_device_physical_interface.example.logical_name or fmc_device_subinterface.example.logical_name). For transparent mode, any bridge group member interface. For routed mode with bridge groups, any bridge group member interface for the BVI name.",
				Computed:            true,
			},
			"interface_id": schema.StringAttribute{
				MarkdownDescription: "UUID of the same interface which has been given by `interface_logical_name` (e.g. fmc_device_physical_interface.example.id or fmc_device_subinterface.example.id). The value is ignored, but the attribute itself is useful for ensuring that Terraform creates interface resource before the static route resource (and destroys the interface resource only after the static route has been destroyed).",
				Computed:            true,
			},
			"destination_networks": schema.SetNestedAttribute{
				MarkdownDescription: "Set of the destination networks matching this route (fmc_network, fmc_host, but not fmc_range).",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "UUID of the object (such as fmc_network.example.id, etc.).",
							Computed:            true,
						},
					},
				},
			},
			"metric_value": schema.Int64Attribute{
				MarkdownDescription: "The cost of the route. The metric is used to compare routes among different routing protocols. The default administrative distance for static routes is 1, giving it precedence over routes discovered by dynamic routing protocols but not directly connected routes.",
				Computed:            true,
			},
			"gateway_object_id": schema.StringAttribute{
				MarkdownDescription: "UUID of the next hop for this route (such as fmc_host.example.id). Exactly one of `gateway_object_id` or `gateway_literal` must be present.",
				Computed:            true,
			},
			"gateway_literal": schema.StringAttribute{
				MarkdownDescription: "The next hop for this route as a literal IPv4 address. Exactly one of `gateway_object_id` or `gateway_literal` must be present.",
				Computed:            true,
			},
			"is_tunneled": schema.BoolAttribute{
				MarkdownDescription: "Indicates whether this route is a separate default route for VPN traffic. Should be used for default route only (such as when the destination_networks points to a builtin network 'any-ipv4'). Useful if you want VPN traffic to use a different default route than non-VPN traffic. When a tunnel terminates on the device, all traffic from it that cannot be routed using learned or static routes is sent to this route. You can configure only one default tunneled gateway per device. ECMP for tunneled traffic is not supported. This attribute conflicts with `metric_value` attribute.",
				Computed:            true,
			},
		},
	}
}

func (d *DeviceIPv4StaticRouteDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *DeviceIPv4StaticRouteDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DeviceIPv4StaticRoute

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
