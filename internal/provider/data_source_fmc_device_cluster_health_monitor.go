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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &DeviceClusterHealthMonitorDataSource{}
	_ datasource.DataSourceWithConfigure = &DeviceClusterHealthMonitorDataSource{}
)

func NewDeviceClusterHealthMonitorDataSource() datasource.DataSource {
	return &DeviceClusterHealthMonitorDataSource{}
}

type DeviceClusterHealthMonitorDataSource struct {
	client *fmc.Client
}

func (d *DeviceClusterHealthMonitorDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_cluster_health_monitor"
}

func (d *DeviceClusterHealthMonitorDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the Device Cluster Health Monitor.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the object",
				Required:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "Name of the FMC domain",
				Optional:            true,
			},
			"cluster_id": schema.StringAttribute{
				MarkdownDescription: "Id of the parent cluster",
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the resource; This is always `ClusterHealthMonitorSetting`.",
				Computed:            true,
			},
			"health_check": schema.BoolAttribute{
				MarkdownDescription: "Enable health check.",
				Computed:            true,
			},
			"hold_time": schema.Float64Attribute{
				MarkdownDescription: "Time (in seconds) to wait before declaring an unresponsive peer as down.",
				Computed:            true,
			},
			"debounce_time": schema.Int64Attribute{
				MarkdownDescription: "The time (in milliseconds) before the interface is considered to have failed.",
				Computed:            true,
			},
			"data_interface_auto_rejoin_attempts": schema.Int64Attribute{
				MarkdownDescription: "Number of rejoin attempts. Use '-1' for unlimited attempts.",
				Computed:            true,
			},
			"data_interface_auto_rejoin_interval": schema.Int64Attribute{
				MarkdownDescription: "Interval duration (in minutes) between rejoin attempts.",
				Computed:            true,
			},
			"data_interface_auto_rejoin_interval_variation": schema.Int64Attribute{
				MarkdownDescription: "Rejoin interval increase strategy. Possible values are 1 (no change); 2 (2 x the previous duration), or 3 (3 x the previous duration)",
				Computed:            true,
			},
			"cluster_interface_auto_rejoin_attempts": schema.Int64Attribute{
				MarkdownDescription: "Number of rejoin attempts. Use '-1' for unlimited attempts.",
				Computed:            true,
			},
			"cluster_interface_auto_rejoin_interval": schema.Int64Attribute{
				MarkdownDescription: "Interval duration (in minutes) between rejoin attempts.",
				Computed:            true,
			},
			"cluster_interface_auto_rejoin_interval_variation": schema.Int64Attribute{
				MarkdownDescription: "Rejoin interval increase strategy. Possible values are 1 (no change); 2 (2 x the previous duration), or 3 (3 x the previous duration)",
				Computed:            true,
			},
			"system_auto_rejoin_attempts": schema.Int64Attribute{
				MarkdownDescription: "Number of rejoin attempts. Use '-1' for unlimited attempts.",
				Computed:            true,
			},
			"system_auto_rejoin_interval": schema.Int64Attribute{
				MarkdownDescription: "Interval duration (in minutes) between rejoin attempts.",
				Computed:            true,
			},
			"system_auto_rejoin_interval_variation": schema.Int64Attribute{
				MarkdownDescription: "Rejoin interval increase strategy. Possible values are 1 (no change); 2 (2 x the previous duration), or 3 (3 x the previous duration)",
				Computed:            true,
			},
			"unmonitored_interfaces": schema.SetAttribute{
				MarkdownDescription: "List of interfaces excluded from monitoring.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"service_application_monitoring": schema.BoolAttribute{
				MarkdownDescription: "Enable service application monitoring (Snort and disk-full processes).",
				Computed:            true,
			},
		},
	}
}

func (d *DeviceClusterHealthMonitorDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *DeviceClusterHealthMonitorDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DeviceClusterHealthMonitor

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
