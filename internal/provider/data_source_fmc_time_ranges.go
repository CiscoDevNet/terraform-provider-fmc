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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/netascode/terraform-provider-fmc/internal/provider/helpers"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &TimeRangesDataSource{}
	_ datasource.DataSourceWithConfigure = &TimeRangesDataSource{}
)

func NewTimeRangesDataSource() datasource.DataSource {
	return &TimeRangesDataSource{}
}

type TimeRangesDataSource struct {
	client *fmc.Client
}

func (d *TimeRangesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_time_ranges"
}

func (d *TimeRangesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the Time Ranges.").String,

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
				MarkdownDescription: "Map of Time Ranges. The key of the map is the name of the individual Time Range.",
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Id of the managed Time Range.",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Type of the object; this is always 'TimeRange'.",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "Description of the object.",
							Computed:            true,
						},
						"start_time": schema.StringAttribute{
							MarkdownDescription: "Date and time at which the time range object starts being effective. If not specified 'starts now' is assumed.",
							Computed:            true,
						},
						"end_time": schema.StringAttribute{
							MarkdownDescription: "Date and time at which the time range object stops being effective. If not specified 'never ends' is assumed.",
							Computed:            true,
						},
						"recurrence_list": schema.ListNestedAttribute{
							MarkdownDescription: "List of recurring intervals during which the time range is effective. These intervals are valid only between start_time and end_time.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"recurrence_type": schema.StringAttribute{
										MarkdownDescription: "Type of the recurrence interval.",
										Computed:            true,
									},
									"range_start_time": schema.StringAttribute{
										MarkdownDescription: "Time (in HH:MM format) at which the time range starts being effective. This field must be used if recurrence_type is specified as RANGE.",
										Computed:            true,
									},
									"range_end_time": schema.StringAttribute{
										MarkdownDescription: "Time (in HH:MM format) at which the time range stops being effective. This field must be used if recurrence_type is specified as RANGE.",
										Computed:            true,
									},
									"range_start_day": schema.StringAttribute{
										MarkdownDescription: "Day of week at which the time range starts being effective. This field must be used if recurrence_type is specified as RANGE.",
										Computed:            true,
									},
									"range_end_day": schema.StringAttribute{
										MarkdownDescription: "Day of week at which the time range stops being effective. This field must be used if recurrence_type is specified as RANGE.",
										Computed:            true,
									},
									"daily_start_time": schema.StringAttribute{
										MarkdownDescription: "Time (in HH:MM format) at which the time range starts being effective on selected days. This field must be used if recurrence_type is specified as DAILY_INTERVAL.",
										Computed:            true,
									},
									"daily_end_time": schema.StringAttribute{
										MarkdownDescription: "Time (in HH:MM format) at which the time range stops being effective on selected days. This field must be used if recurrence_type is specified as DAILY_INTERVAL.",
										Computed:            true,
									},
									"daily_days": schema.SetAttribute{
										MarkdownDescription: "List of days on which the time range is effective. This field must be used if recurrence_type is specified as DAILY_INTERVAL.",
										ElementType:         types.StringType,
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

func (d *TimeRangesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *TimeRangesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config TimeRanges

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
