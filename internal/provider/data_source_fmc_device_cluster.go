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
	_ datasource.DataSource              = &DeviceClusterDataSource{}
	_ datasource.DataSourceWithConfigure = &DeviceClusterDataSource{}
)

func NewDeviceClusterDataSource() datasource.DataSource {
	return &DeviceClusterDataSource{}
}

type DeviceClusterDataSource struct {
	client *fmc.Client
}

func (d *DeviceClusterDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_cluster"
}

func (d *DeviceClusterDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the Device Cluster.").String,

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
				MarkdownDescription: "Name of the FTD Cluster.",
				Optional:            true,
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the resource; This is always `DeviceCluster`.",
				Computed:            true,
			},
			"cluster_key": schema.StringAttribute{
				MarkdownDescription: "Secret key for the cluster, between 1 nd 63 characters.",
				Computed:            true,
			},
			"control_node_vni_prefix": schema.StringAttribute{
				MarkdownDescription: "Cluster Control VXLAN Network Identifier (VNI) Network",
				Computed:            true,
			},
			"control_node_ccl_prefix": schema.StringAttribute{
				MarkdownDescription: "Cluster Control Link Network / Virtual Tunnel Endpoint (VTEP) Network",
				Computed:            true,
			},
			"control_node_interface_id": schema.StringAttribute{
				MarkdownDescription: "Cluster control link interface ID.",
				Computed:            true,
			},
			"control_node_interface_name": schema.StringAttribute{
				MarkdownDescription: "Cluster control link interface Name.",
				Computed:            true,
			},
			"control_node_interface_type": schema.StringAttribute{
				MarkdownDescription: "Cluster control link interface Type.",
				Computed:            true,
			},
			"control_node_device_id": schema.StringAttribute{
				MarkdownDescription: "Cluster Control Node device ID.",
				Computed:            true,
			},
			"control_node_ccl_ipv4_address": schema.StringAttribute{
				MarkdownDescription: "Cluster control link IPv4 address / VTEP IPv4 address.",
				Computed:            true,
			},
			"control_node_priority": schema.Int64Attribute{
				MarkdownDescription: "Priority of cluster controle node.",
				Computed:            true,
			},
			"data_devices": schema.ListNestedAttribute{
				MarkdownDescription: "List of cluster data nodes.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"data_node_device_id": schema.StringAttribute{
							MarkdownDescription: "Cluster Data Node device ID.",
							Computed:            true,
						},
						"data_node_ccl_ipv4_address": schema.StringAttribute{
							MarkdownDescription: "Cluster Data Node link IPv4 address / VTEP IPv4 address.",
							Computed:            true,
						},
						"data_node_priority": schema.Int64Attribute{
							MarkdownDescription: "Priority of cluster data node.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}
func (d *DeviceClusterDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *DeviceClusterDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *DeviceClusterDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DeviceCluster

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
