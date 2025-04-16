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
	_ datasource.DataSource              = &ChassisLogicalDeviceDataSource{}
	_ datasource.DataSourceWithConfigure = &ChassisLogicalDeviceDataSource{}
)

func NewChassisLogicalDeviceDataSource() datasource.DataSource {
	return &ChassisLogicalDeviceDataSource{}
}

type ChassisLogicalDeviceDataSource struct {
	client *fmc.Client
}

func (d *ChassisLogicalDeviceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_chassis_logical_device"
}

func (d *ChassisLogicalDeviceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the Chassis Logical Device.").String,

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
			"chassis_id": schema.StringAttribute{
				MarkdownDescription: "Id of the parent chassis.",
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the device; this value is always 'LogicalDevice'.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the logical device.",
				Optional:            true,
				Computed:            true,
			},
			"ftd_version": schema.StringAttribute{
				MarkdownDescription: "Version of the logical device, that should be deployed. Image should be pre-deployed to the chassis.",
				Computed:            true,
			},
			"ipv4_address": schema.StringAttribute{
				MarkdownDescription: "Management IPv4 address of the logical device.",
				Computed:            true,
			},
			"ipv4_netmask": schema.StringAttribute{
				MarkdownDescription: "Netmask of Management IPv4 address.",
				Computed:            true,
			},
			"ipv4_gateway": schema.StringAttribute{
				MarkdownDescription: "Gateway for Management IPv4 address.",
				Computed:            true,
			},
			"ipv6_address": schema.StringAttribute{
				MarkdownDescription: "Management IPv6 address of the logical device.",
				Computed:            true,
			},
			"ipv6_prefix_length": schema.Int64Attribute{
				MarkdownDescription: "Prefix length of Management IPv6 address.",
				Computed:            true,
			},
			"ipv6_gateway": schema.StringAttribute{
				MarkdownDescription: "Gateway for Management IPv6 address.",
				Computed:            true,
			},
			"search_domain": schema.StringAttribute{
				MarkdownDescription: "Search domain for the logical device.",
				Computed:            true,
			},
			"fqdn": schema.StringAttribute{
				MarkdownDescription: "Fully qualified domain name (FQDN) of the logical device.",
				Computed:            true,
			},
			"firewall_mode": schema.StringAttribute{
				MarkdownDescription: "Firewall mode of the logical device.",
				Computed:            true,
			},
			"dns_servers": schema.StringAttribute{
				MarkdownDescription: "DNS servers for the logical device. Up to three, comma-separated DNS servers can be specified.",
				Computed:            true,
			},
			"device_password": schema.StringAttribute{
				MarkdownDescription: "Admin password for the logical device.",
				Computed:            true,
			},
			"permit_expert_mode": schema.StringAttribute{
				MarkdownDescription: "Permit expert mode for the logical device.",
				Computed:            true,
			},
			"resource_profile_id": schema.StringAttribute{
				MarkdownDescription: "Id of the resource profile.",
				Computed:            true,
			},
			"resource_profile_name": schema.StringAttribute{
				MarkdownDescription: "Name of the resource profile.",
				Computed:            true,
			},
			"assigned_interfaces": schema.SetNestedAttribute{
				MarkdownDescription: "Interface assignment for the logical device.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Id of the interface.",
							Computed:            true,
						},
					},
				},
			},
			"device_group_id": schema.StringAttribute{
				MarkdownDescription: "Id of the device group.",
				Computed:            true,
			},
			"access_policy_id": schema.StringAttribute{
				MarkdownDescription: "Id of the Access Control Policy to be assigned to the logical device.",
				Computed:            true,
			},
			"platform_settings_id": schema.StringAttribute{
				MarkdownDescription: "Id of the platform settings.",
				Computed:            true,
			},
			"license_capabilities": schema.SetAttribute{
				MarkdownDescription: "Array of strings representing the license capabilities on the managed device.",
				ElementType:         types.StringType,
				Computed:            true,
			},
		},
	}
}
func (d *ChassisLogicalDeviceDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *ChassisLogicalDeviceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *ChassisLogicalDeviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ChassisLogicalDevice

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
