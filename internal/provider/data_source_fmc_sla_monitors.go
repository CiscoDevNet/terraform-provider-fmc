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
	_ datasource.DataSource              = &SLAMonitorsDataSource{}
	_ datasource.DataSourceWithConfigure = &SLAMonitorsDataSource{}
)

func NewSLAMonitorsDataSource() datasource.DataSource {
	return &SLAMonitorsDataSource{}
}

type SLAMonitorsDataSource struct {
	client *fmc.Client
}

func (d *SLAMonitorsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sla_monitors"
}

func (d *SLAMonitorsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the SLA Monitors.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the object",
				Computed:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "Name of the FMC domain",
				Optional:            true,
			},
			"items": schema.MapNestedAttribute{
				MarkdownDescription: "Map of SLA Monitors. The key of the map is the name of the individual SLA monitor.",
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Id of the managed SLA Monitor.",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Type of the object; this value is always 'SLAMonitor'.",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "Description of the object.",
							Computed:            true,
						},
						"sla_monitor_id": schema.Int64Attribute{
							MarkdownDescription: "ID number of the SLA operation.",
							Computed:            true,
						},
						"timeout": schema.Int64Attribute{
							MarkdownDescription: "Amount of time (in milliseconds) that the SLA operation waits for a response to the ICMP echo requests.",
							Computed:            true,
						},
						"frequency": schema.Int64Attribute{
							MarkdownDescription: "Frequency (in seconds) of ICMP echo request transmissions.",
							Computed:            true,
						},
						"threshold": schema.Int64Attribute{
							MarkdownDescription: "Amount of time (in milliseconds) that must pass after an ICMP echo request before a rising threshold is declared.",
							Computed:            true,
						},
						"data_size": schema.Int64Attribute{
							MarkdownDescription: "Size (in bytes) of the ICMP request packet payload.",
							Computed:            true,
						},
						"tos": schema.Int64Attribute{
							MarkdownDescription: "Type of Service (ToS) defined in the IP header of the ICMP request packet.",
							Computed:            true,
						},
						"number_of_packets": schema.Int64Attribute{
							MarkdownDescription: "Number of packets that are sent.",
							Computed:            true,
						},
						"monitor_address": schema.StringAttribute{
							MarkdownDescription: "IP address to monitor.",
							Computed:            true,
						},
						"selected_interfaces": schema.ListNestedAttribute{
							MarkdownDescription: "Security zones or interface groups that contain the interfaces through which the device communicates with the management station.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the security zone or interface object.",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *SLAMonitorsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *SLAMonitorsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SLAMonitors

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

	// Get all objects from FMC
	urlPath := config.getPath() + "?expanded=true"
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
