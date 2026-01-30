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
	_ datasource.DataSource              = &RadiusServerGroupDataSource{}
	_ datasource.DataSourceWithConfigure = &RadiusServerGroupDataSource{}
)

func NewRadiusServerGroupDataSource() datasource.DataSource {
	return &RadiusServerGroupDataSource{}
}

type RadiusServerGroupDataSource struct {
	client *fmc.Client
}

func (d *RadiusServerGroupDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_radius_server_group"
}

func (d *RadiusServerGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the Radius Server Group.").String,

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
				MarkdownDescription: "Name of the RADIUS Server Group object.",
				Optional:            true,
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the object; this value is always 'RadiusServerGroup'.",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the object.",
				Computed:            true,
			},
			"group_accounting_mode": schema.StringAttribute{
				MarkdownDescription: "Indicates whether accounting messages are sent to a single server (SINGLE) or sent to all servers in the group simultaneously (MULTIPLE).",
				Computed:            true,
			},
			"retry_interval": schema.Int64Attribute{
				MarkdownDescription: "Retry interval (in seconds) for the request.",
				Computed:            true,
			},
			"ad_realm_id": schema.StringAttribute{
				MarkdownDescription: "Id of Active Directory (AD) realm this RADIUS server group is associated with.",
				Computed:            true,
			},
			"authorize_only": schema.BoolAttribute{
				MarkdownDescription: "This RADIUS server group is used for authorization or accounting only.",
				Computed:            true,
			},
			"interim_account_update_interval": schema.Int64Attribute{
				MarkdownDescription: "Interval (in hours) for interim accounting updates.",
				Computed:            true,
			},
			"dynamic_authorization": schema.BoolAttribute{
				MarkdownDescription: "Enables the RADIUS dynamic authorization or Change of Authorization (CoA) services for this RADIUS server group.",
				Computed:            true,
			},
			"dynamic_authorization_port": schema.Int64Attribute{
				MarkdownDescription: "Port number for the RADIUS dynamic authorization services.",
				Computed:            true,
			},
			"merge_downloadable_access_list_order": schema.StringAttribute{
				MarkdownDescription: "Placement order of the downloadable Access List with the Cisco AV pair Access List.",
				Computed:            true,
			},
			"radius_servers": schema.ListNestedAttribute{
				MarkdownDescription: "List of RADIUS servers in the group.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"hostname": schema.StringAttribute{
							MarkdownDescription: "IP Address or hostname of the RADIUS server.",
							Computed:            true,
						},
						"message_authenticator": schema.BoolAttribute{
							MarkdownDescription: "Enables RADIUS Server-Enabled Message Authenticator.",
							Computed:            true,
						},
						"authentication_port": schema.Int64Attribute{
							MarkdownDescription: "Port number for the RADIUS authentication services.",
							Computed:            true,
						},
						"key": schema.StringAttribute{
							MarkdownDescription: "Shared secret that is used for data encryption.",
							Computed:            true,
							Sensitive:           true,
						},
						"accounting_port": schema.Int64Attribute{
							MarkdownDescription: "Port number for the RADIUS accounting services.",
							Computed:            true,
						},
						"timeout": schema.Int64Attribute{
							MarkdownDescription: "Timeout (in seconds) for the RADIUS server.",
							Computed:            true,
						},
						"use_routing_to_select_interface": schema.BoolAttribute{
							MarkdownDescription: "Use routing to select the interface for the RADIUS server (true) or use specified interface (false).",
							Computed:            true,
						},
						"interface_id": schema.StringAttribute{
							MarkdownDescription: "Id of Security Zone or Interface Group for the RADIUS server communication.",
							Computed:            true,
						},
						"redirect_access_list_id": schema.StringAttribute{
							MarkdownDescription: "Id of the redirect Extended Access List.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}
func (d *RadiusServerGroupDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *RadiusServerGroupDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *RadiusServerGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config RadiusServerGroup

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
