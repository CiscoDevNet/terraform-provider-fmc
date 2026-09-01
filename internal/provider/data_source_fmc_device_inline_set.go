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
	_ datasource.DataSource              = &DeviceInlineSetDataSource{}
	_ datasource.DataSourceWithConfigure = &DeviceInlineSetDataSource{}
)

func NewDeviceInlineSetDataSource() datasource.DataSource {
	return &DeviceInlineSetDataSource{}
}

type DeviceInlineSetDataSource struct {
	client *fmc.Client
}

func (d *DeviceInlineSetDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_inline_set"
}

func (d *DeviceInlineSetDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the Device Inline Set.").String,

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
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the object; this value is always 'InlineSet'.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the Inline Set.",
				Optional:            true,
				Computed:            true,
			},
			"mtu": schema.Int64Attribute{
				MarkdownDescription: "Maximum Transmission Unit (MTU) of the Inline Set.",
				Computed:            true,
			},
			"fail_safe": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"bypass_standby": schema.BoolAttribute{
				MarkdownDescription: "Put the hardware bypass of the Inline Set into standby mode. Available only on interfaces of devices equipped with a bypass-capable network module.",
				Computed:            true,
			},
			"bypass_force": schema.BoolAttribute{
				MarkdownDescription: "Enable hardware bypass (bypass-force) for the Inline Set. Available only on interfaces of devices equipped with a bypass-capable network module.",
				Computed:            true,
			},
			"interface_pairs": schema.SetNestedAttribute{
				MarkdownDescription: "Interface pairs that are members of the Inline Set.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"first_interface_id": schema.StringAttribute{
							MarkdownDescription: "Id of the first interface of the interface pair.",
							Computed:            true,
						},
						"first_interface_name": schema.StringAttribute{
							MarkdownDescription: "Name of the first interface of the interface pair.",
							Computed:            true,
						},
						"first_interface_type": schema.StringAttribute{
							MarkdownDescription: "Type of the first interface of the interface pair.",
							Computed:            true,
						},
						"second_interface_id": schema.StringAttribute{
							MarkdownDescription: "Id of the second interface of the interface pair.",
							Computed:            true,
						},
						"second_interface_name": schema.StringAttribute{
							MarkdownDescription: "Name of the second interface of the interface pair.",
							Computed:            true,
						},
						"second_interface_type": schema.StringAttribute{
							MarkdownDescription: "Type of the second interface of the interface pair.",
							Computed:            true,
						},
					},
				},
			},
			"tap_mode": schema.BoolAttribute{
				MarkdownDescription: "Enable tap mode, so that a copy of the traffic is inspected while the original traffic is forwarded without being affected by the inspection results.",
				Computed:            true,
			},
			"propagate_link_state": schema.BoolAttribute{
				MarkdownDescription: "Automatically bring down the second interface of an interface pair when one of the interfaces in the pair goes down.",
				Computed:            true,
			},
			"strict_tcp_enforcement": schema.BoolAttribute{
				MarkdownDescription: "Block all TCP packets that are not part of a properly established TCP session.",
				Computed:            true,
			},
			"snort_fail_open_busy": schema.BoolAttribute{
				MarkdownDescription: "Allow new traffic to pass without inspection when the Snort process is busy.",
				Computed:            true,
			},
			"snort_fail_open_down": schema.BoolAttribute{
				MarkdownDescription: "Allow new traffic to pass without inspection when the Snort process is down.",
				Computed:            true,
			},
		},
	}
}
func (d *DeviceInlineSetDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *DeviceInlineSetDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *DeviceInlineSetDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DeviceInlineSet

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
