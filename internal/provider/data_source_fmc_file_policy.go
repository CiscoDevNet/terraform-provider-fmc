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
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &FilePolicyDataSource{}
	_ datasource.DataSourceWithConfigure = &FilePolicyDataSource{}
)

func NewFilePolicyDataSource() datasource.DataSource {
	return &FilePolicyDataSource{}
}

type FilePolicyDataSource struct {
	client *fmc.Client
}

func (d *FilePolicyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_file_policy"
}

func (d *FilePolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the File Policy.",

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
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of file policy.",
				Optional:            true,
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the object",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "File policy description.",
				Computed:            true,
			},
			"first_time_file_analysis": schema.BoolAttribute{
				MarkdownDescription: "Analyze first-seen files while AMP cloud disposition is pending",
				Computed:            true,
			},
			"custom_detection_list": schema.BoolAttribute{
				MarkdownDescription: "Enable custom detection list",
				Computed:            true,
			},
			"clean_list": schema.BoolAttribute{
				MarkdownDescription: "Enable clean list",
				Computed:            true,
			},
			"threat_score": schema.StringAttribute{
				MarkdownDescription: "If AMP Cloud disposition is Unknown, override disposition based upon threat score.",
				Computed:            true,
			},
			"inspect_archives": schema.BoolAttribute{
				MarkdownDescription: "Inspect Archives",
				Computed:            true,
			},
			"block_encrypted_archives": schema.BoolAttribute{
				MarkdownDescription: "Block encrypted archives",
				Computed:            true,
			},
			"block_uninspectable_archives": schema.BoolAttribute{
				MarkdownDescription: "Block uninspectable Archives",
				Computed:            true,
			},
			"max_archive_depth": schema.Int64Attribute{
				MarkdownDescription: "Max archive depth",
				Computed:            true,
			},
			"file_rules": schema.ListNestedAttribute{
				MarkdownDescription: "The ordered list of file rules.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Unique identifier representing the File Rule.",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "The name of file rule type.",
							Computed:            true,
						},
						"application_protocol": schema.StringAttribute{
							MarkdownDescription: "Defines a protocol for file inspection.",
							Computed:            true,
						},
						"action": schema.StringAttribute{
							MarkdownDescription: "Action to be performed on a file.",
							Computed:            true,
						},
						"store_files": schema.SetAttribute{
							MarkdownDescription: "List of file dispositions that should be stored (MALWARE, CUSTOM, CLEAN, UNKNOWN).",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"direction_of_transfer": schema.StringAttribute{
							MarkdownDescription: "Direction of file transfer.",
							Computed:            true,
						},
						"file_categories": schema.SetNestedAttribute{
							MarkdownDescription: "Defines a list of file categories for inspection.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "The id of file category.",
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: "The name of file category.",
										Computed:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: "The type of file category.",
										Computed:            true,
									},
								},
							},
						},
						"file_types": schema.SetNestedAttribute{
							MarkdownDescription: "Defines a list of file types for inspection.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "The id of file type.",
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: "The name of file type.",
										Computed:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: "The name of file type.",
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
func (d *FilePolicyDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *FilePolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

func (d *FilePolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config FilePolicy

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
	urlPath := config.getPath() + "/" + url.QueryEscape(config.Id.ValueString())
	res, err := d.client.Get(urlPath, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	resFileRules, err := d.client.Get(config.getPath()+"/"+url.QueryEscape(config.Id.ValueString())+"/filerules?expanded=true", reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	replaceFileRules := resFileRules.Get("items").String()
	if replaceFileRules == "" {
		replaceFileRules = "[]"
	}

	replace, _ := sjson.SetRaw(res.String(), "dummy_file_rules", replaceFileRules)
	res = gjson.Parse(replace)

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
