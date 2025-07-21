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
	_ datasource.DataSource              = &RealmADLDAPDataSource{}
	_ datasource.DataSourceWithConfigure = &RealmADLDAPDataSource{}
)

func NewRealmADLDAPDataSource() datasource.DataSource {
	return &RealmADLDAPDataSource{}
}

type RealmADLDAPDataSource struct {
	client *fmc.Client
}

func (d *RealmADLDAPDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_realm_ad_ldap"
}

func (d *RealmADLDAPDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the Realm AD LDAP.").String,

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
				MarkdownDescription: "Name of the Realm object.",
				Optional:            true,
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the object; this value is always 'Realm'.",
				Computed:            true,
			},
			"version": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the realm.",
				Computed:            true,
			},
			"realm_type": schema.StringAttribute{
				MarkdownDescription: "Type of the realm",
				Computed:            true,
			},
			"ad_primary_domain": schema.StringAttribute{
				MarkdownDescription: "Primary domain for AD realm.",
				Computed:            true,
			},
			"ad_join_username": schema.StringAttribute{
				MarkdownDescription: "Username for joining the AD domain.",
				Computed:            true,
			},
			"ad_join_password": schema.StringAttribute{
				MarkdownDescription: "Password for joining the AD domain.",
				Computed:            true,
				Sensitive:           true,
			},
			"directory_username": schema.StringAttribute{
				MarkdownDescription: "Username for joining the AD domain.",
				Computed:            true,
			},
			"directory_password": schema.StringAttribute{
				MarkdownDescription: "Password for the AD domain user.",
				Computed:            true,
				Sensitive:           true,
			},
			"base_dn": schema.StringAttribute{
				MarkdownDescription: "Base DN for the LDAP search.",
				Computed:            true,
			},
			"group_dn": schema.StringAttribute{
				MarkdownDescription: "DN of the group to search for users.",
				Computed:            true,
			},
			"update_hour": schema.Int64Attribute{
				MarkdownDescription: "Hour where the sync (download) from the directory starts.",
				Computed:            true,
			},
			"update_interval": schema.StringAttribute{
				MarkdownDescription: "Interval in hours for the sync (download) from the directory.",
				Computed:            true,
			},
			"group_attribute": schema.StringAttribute{
				MarkdownDescription: "Attribute used to identify the group in the LDAP directory. Use uniqueMember, member or any custom attribute name.",
				Computed:            true,
			},
			"timeout_ise_users": schema.Int64Attribute{
				MarkdownDescription: "Timeout for the authentication session in seconds.",
				Computed:            true,
			},
			"timeout_terminal_server_agent_users": schema.Int64Attribute{
				MarkdownDescription: "Timeout for the authentication session in seconds.",
				Computed:            true,
			},
			"timeout_captive_portal_users": schema.Int64Attribute{
				MarkdownDescription: "Timeout for the authentication session in seconds.",
				Computed:            true,
			},
			"timeout_failed_captive_portal_users": schema.Int64Attribute{
				MarkdownDescription: "Timeout for the authentication session in seconds.",
				Computed:            true,
			},
			"timeout_guest_captive_portal_users": schema.Int64Attribute{
				MarkdownDescription: "Timeout for the authentication session in seconds.",
				Computed:            true,
			},
			"directory_server_configurations": schema.ListNestedAttribute{
				MarkdownDescription: "List of directory configurations for the realm.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"hostname": schema.StringAttribute{
							MarkdownDescription: "Hostname or IP address of the LDAP server.",
							Computed:            true,
						},
						"port": schema.Int64Attribute{
							MarkdownDescription: "Port number for the LDAP server.",
							Computed:            true,
						},
						"encryption_protocol": schema.StringAttribute{
							MarkdownDescription: "Encryption method for the LDAP connection.",
							Computed:            true,
						},
						"encryption_certificate": schema.StringAttribute{
							MarkdownDescription: "ID of the encryption certificate for LDAPS.",
							Computed:            true,
						},
						"use_routing_to_select_interface": schema.BoolAttribute{
							MarkdownDescription: "Whether to use routing to select the interface for LDAP communication.",
							Computed:            true,
						},
						"interface_group_id": schema.StringAttribute{
							MarkdownDescription: "ID of the interface group to use for LDAP communication.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}
func (d *RealmADLDAPDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *RealmADLDAPDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *RealmADLDAPDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config RealmADLDAP

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
