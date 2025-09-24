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
	_ datasource.DataSource              = &HealthPolicyDataSource{}
	_ datasource.DataSourceWithConfigure = &HealthPolicyDataSource{}
)

func NewHealthPolicyDataSource() datasource.DataSource {
	return &HealthPolicyDataSource{}
}

type HealthPolicyDataSource struct {
	client *fmc.Client
}

func (d *HealthPolicyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_health_policy"
}

func (d *HealthPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the Health Policy.").AddMinimumVersionHeaderDescription().AddMinimumVersionDescription("7.7").String,

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
				MarkdownDescription: "Name of the Health Policy.",
				Optional:            true,
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the Health Policy.",
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Object type; This is always `HealthPolicy`",
				Computed:            true,
			},
			"policy_type": schema.StringAttribute{
				MarkdownDescription: "Define, if the policy is for the device or for the management center.",
				Computed:            true,
			},
			"is_default_policy": schema.BoolAttribute{
				MarkdownDescription: "Specify if this is the default policy.",
				Computed:            true,
			},
			"health_modules": schema.ListNestedAttribute{
				MarkdownDescription: "Configuration of health modules.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Name of the health module.",
							Computed:            true,
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: "Enable health module.",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Type of health module.",
							Computed:            true,
						},
						"alert_severity": schema.StringAttribute{
							MarkdownDescription: "Severity of health module alert.",
							Computed:            true,
						},
						"critical_threshold": schema.Int64Attribute{
							MarkdownDescription: "Critical threshold value for health module.",
							Computed:            true,
						},
						"warning_threshold": schema.Int64Attribute{
							MarkdownDescription: "Warning threshold value for health module.",
							Computed:            true,
						},
						"custom_thresholds": schema.ListNestedAttribute{
							MarkdownDescription: "Custom threshold configuration for health module.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of custom threshold.",
										Computed:            true,
									},
									"threshold": schema.Int64Attribute{
										MarkdownDescription: "Threshold level.",
										Computed:            true,
									},
								},
							},
						},
						"alert_configs": schema.ListNestedAttribute{
							MarkdownDescription: "Alert configuration for health module.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Name of the alert configuration.",
										Computed:            true,
									},
									"enabled": schema.BoolAttribute{
										MarkdownDescription: "Enable alert configuration.",
										Computed:            true,
									},
									"thresholds": schema.ListNestedAttribute{
										MarkdownDescription: "Thresholds for alert configuration.",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"type": schema.StringAttribute{
													MarkdownDescription: "Type of threshold.",
													Computed:            true,
												},
												"threshold": schema.Int64Attribute{
													MarkdownDescription: "Threshold level.",
													Computed:            true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"health_module_run_time_interval": schema.Int64Attribute{
				MarkdownDescription: "Specify the interval (in minutes) for running the health modules.",
				Computed:            true,
			},
			"metric_collection_interval": schema.Int64Attribute{
				MarkdownDescription: "Specify the frequency (in minutes) for collection of time series data.",
				Computed:            true,
			},
		},
	}
}
func (d *HealthPolicyDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *HealthPolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *HealthPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	// Check if FMC client is connected to supports this object
	if d.client.FMCVersionParsed.LessThan(minFMCVersionHealthPolicy) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support Health Policy, minimum required version is 7.7", d.client.FMCVersion))
		return
	}
	var config HealthPolicy

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
						tflog.Debug(ctx, fmt.Sprintf("%s: Found object with name '%v', id: %v", config.Id.ValueString(), config.Name.ValueString(), config.Id.ValueString()))
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
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with name: %v", config.Name.ValueString()))
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
