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
	_ datasource.DataSource              = &DeviceBGPGeneralSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &DeviceBGPGeneralSettingsDataSource{}
)

func NewDeviceBGPGeneralSettingsDataSource() datasource.DataSource {
	return &DeviceBGPGeneralSettingsDataSource{}
}

type DeviceBGPGeneralSettingsDataSource struct {
	client *fmc.Client
}

func (d *DeviceBGPGeneralSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_bgp_general_settings"
}

func (d *DeviceBGPGeneralSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the Device BGP General Settings.").String,

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
			"name": schema.StringAttribute{
				MarkdownDescription: "Object name; Always set to 'AsaBGPGeneralTable'",
				Computed:            true,
			},
			"as_number": schema.StringAttribute{
				MarkdownDescription: "Autonomous System (AS) number in 'asplain' or 'asdot' format",
				Optional:            true,
				Computed:            true,
			},
			"router_id": schema.StringAttribute{
				MarkdownDescription: "String value for the routerID. Possible values can be 'AUTOMATIC' or valid ipv4 address",
				Computed:            true,
			},
			"scanning_interval": schema.Int64Attribute{
				MarkdownDescription: "Scanning interval of BGP routers for next hop validation in Seconds.",
				Computed:            true,
			},
			"as_number_in_path_attribute": schema.Int64Attribute{
				MarkdownDescription: "Range to discard routes that have as-path segments that exceed a specified value.",
				Computed:            true,
			},
			"log_neighbor_changes": schema.BoolAttribute{
				MarkdownDescription: "Logging of BGP neighbor status changes.",
				Computed:            true,
			},
			"tcp_path_mtu_discovery": schema.BoolAttribute{
				MarkdownDescription: "Use TCP path MTU discovery.",
				Computed:            true,
			},
			"reset_session_upon_failover": schema.BoolAttribute{
				MarkdownDescription: "Reset session upon failover",
				Computed:            true,
			},
			"enforce_first_peer_as": schema.BoolAttribute{
				MarkdownDescription: "Discard updates received from an external BGP (eBGP) peers that do not list their autonomous system (AS) number.",
				Computed:            true,
			},
			"use_dot_notation": schema.BoolAttribute{
				MarkdownDescription: "Change format of BGP 4-byte autonomous system numbers from asplain (decimal values) to dot notation.",
				Computed:            true,
			},
			"aggregate_timer": schema.Int64Attribute{
				MarkdownDescription: "Interval at which BGP routes will be aggregated or to disable timer-based router aggregation (in seconds).",
				Computed:            true,
			},
			"default_local_preference": schema.Int64Attribute{
				MarkdownDescription: "Default local preference",
				Computed:            true,
			},
			"compare_med_from_different_neighbors": schema.BoolAttribute{
				MarkdownDescription: "Allow comparing MED from different neighbors",
				Computed:            true,
			},
			"compare_router_id_in_path": schema.BoolAttribute{
				MarkdownDescription: "Compare Router ID for identical EBGP paths",
				Computed:            true,
			},
			"pick_best_med": schema.BoolAttribute{
				MarkdownDescription: "Pick the best-MED path among paths advertised by neighbor AS",
				Computed:            true,
			},
			"missing_med_as_best": schema.BoolAttribute{
				MarkdownDescription: "Treat missing MED as the best preferred path",
				Computed:            true,
			},
			"keepalive_interval": schema.Int64Attribute{
				MarkdownDescription: "Keepalive interval in seconds",
				Computed:            true,
			},
			"hold_time": schema.Int64Attribute{
				MarkdownDescription: "Hold time in seconds",
				Computed:            true,
			},
			"min_hold_time": schema.Int64Attribute{
				MarkdownDescription: "Minimum hold time (0 or 3-65535 seconds)",
				Computed:            true,
			},
			"next_hop_address_tracking": schema.BoolAttribute{
				MarkdownDescription: "Enable next hop address tracking",
				Computed:            true,
			},
			"next_hop_delay_interval": schema.Int64Attribute{
				MarkdownDescription: "Next hop delay interval in seconds",
				Computed:            true,
			},
			"graceful_restart": schema.BoolAttribute{
				MarkdownDescription: "Enable graceful restart",
				Computed:            true,
			},
			"graceful_restart_restart_time": schema.Int64Attribute{
				MarkdownDescription: "Graceful Restart Time in seconds",
				Computed:            true,
			},
			"graceful_restart_stale_path_time": schema.Int64Attribute{
				MarkdownDescription: "Stalepath Time in seconds",
				Computed:            true,
			},
		},
	}
}
func (d *DeviceBGPGeneralSettingsDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("as_number"),
		),
	}
}

func (d *DeviceBGPGeneralSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *DeviceBGPGeneralSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DeviceBGPGeneralSettings

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
