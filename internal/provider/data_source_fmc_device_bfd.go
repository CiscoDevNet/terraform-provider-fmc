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
	_ datasource.DataSource              = &DeviceBFDDataSource{}
	_ datasource.DataSourceWithConfigure = &DeviceBFDDataSource{}
)

func NewDeviceBFDDataSource() datasource.DataSource {
	return &DeviceBFDDataSource{}
}

type DeviceBFDDataSource struct {
	client *fmc.Client
}

func (d *DeviceBFDDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_bfd"
}

func (d *DeviceBFDDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the Device BFD.").AddMinimumVersionHeaderDescription().AddMinimumVersionDescription("7.4").String,

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
			"vrf_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the parent VRF.").String,
				Optional:            true,
			},
			"device_id": schema.StringAttribute{
				MarkdownDescription: "Id of the parent device.",
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the object; this value is always 'BFDPolicy'",
				Computed:            true,
			},
			"hop_type": schema.StringAttribute{
				MarkdownDescription: "BFD Hop type.",
				Computed:            true,
			},
			"bfd_template_id": schema.StringAttribute{
				MarkdownDescription: "Id of the BFD Template",
				Computed:            true,
			},
			"interface_logical_name": schema.StringAttribute{
				MarkdownDescription: "Logical Name of the interface for BFD assignment if `hop_type` is set to SINGLE_HOP.",
				Optional:            true,
				Computed:            true,
			},
			"destination_host_object_id": schema.StringAttribute{
				MarkdownDescription: "Id of the destination host object if `hop_type` is set to MULTI_HOP.",
				Computed:            true,
			},
			"source_host_object_id": schema.StringAttribute{
				MarkdownDescription: "Id of the source host object if `hop_type` is set to MULTI_HOP.",
				Computed:            true,
			},
			"interface_id": schema.StringAttribute{
				MarkdownDescription: "Id of the interface for BFD assignment if `hop_type` is set to SINGLE_HOP.",
				Computed:            true,
			},
			"slow_timer": schema.Int64Attribute{
				MarkdownDescription: "BFD Slow Timer value in range: 1000-30000, default: 1000",
				Computed:            true,
			},
		},
	}
}
func (d *DeviceBFDDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("interface_logical_name"),
		),
	}
}

func (d *DeviceBFDDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *DeviceBFDDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	// Check if FMC client is connected to supports this object
	if d.client.FMCVersionParsed.LessThan(minFMCVersionDeviceBFD) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support Device BFD, minimum required version is 7.4", d.client.FMCVersion))
		return
	}
	var config DeviceBFD

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
	if config.Id.IsNull() && !config.InterfaceLogicalName.IsNull() {
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
					if config.InterfaceLogicalName.ValueString() == v.Get("interface.ifname").String() {
						config.Id = types.StringValue(v.Get("id").String())
						tflog.Debug(ctx, fmt.Sprintf("%s: Found object with interface_logical_name '%v', id: %v", config.Id.ValueString(), config.InterfaceLogicalName.ValueString(), config.Id.ValueString()))
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
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with interface_logical_name: %v", config.InterfaceLogicalName.ValueString()))
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
