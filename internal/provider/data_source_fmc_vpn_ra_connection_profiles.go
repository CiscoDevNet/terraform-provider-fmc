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
	"strings"
	"time"

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
	_ datasource.DataSource              = &VPNRAConnectionProfilesDataSource{}
	_ datasource.DataSourceWithConfigure = &VPNRAConnectionProfilesDataSource{}
)

func NewVPNRAConnectionProfilesDataSource() datasource.DataSource {
	return &VPNRAConnectionProfilesDataSource{}
}

type VPNRAConnectionProfilesDataSource struct {
	client *fmc.Client
}

func (d *VPNRAConnectionProfilesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_ra_connection_profiles"
}

func (d *VPNRAConnectionProfilesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the VPN RA Connection Profiles.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the object",
				Computed:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "Name of the FMC domain",
				Optional:            true,
			},
			"vpn_ra_id": schema.StringAttribute{
				MarkdownDescription: "Id of the parent VPN RA Configuration.",
				Required:            true,
			},
			"items": schema.MapNestedAttribute{
				MarkdownDescription: "Map of Connection Profiles. The key of the map is the name of the Connection Profile. Use `DefaultWEBVPNGroup` to manage the default connection profile. On destruction, the default connection profile will not be deleted and its configuration will not be erased.",
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Id of the Connection Profile.",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Type of the object; this value is always 'RaVpnConnectionProfile'.",
							Computed:            true,
						},
						"group_policy_id": schema.StringAttribute{
							MarkdownDescription: "Id of the Group Policy.",
							Computed:            true,
						},
						"ipv4_address_pools": schema.ListNestedAttribute{
							MarkdownDescription: "IPv4 Address Pools for client address assignment. Either `ipv4_address_pools` or `ipv6_address_pools` or `dhcp_servers` must be specified.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the IPv4 Address Pool.",
										Computed:            true,
									},
								},
							},
						},
						"ipv6_address_pools": schema.ListNestedAttribute{
							MarkdownDescription: "IPv6 Address Pools for client address assignment. Either `ipv4_address_pools` or `ipv6_address_pools` or `dhcp_servers` must be specified.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the IPv6 Address Pool.",
										Computed:            true,
									},
								},
							},
						},
						"dhcp_servers": schema.ListNestedAttribute{
							MarkdownDescription: "List of DHCP Servers for client address assignment. Either `ipv4_address_pools` or `ipv6_address_pools` or `dhcp_servers` must be specified.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of Host representing the DHCP Server.",
										Computed:            true,
									},
								},
							},
						},
						"authentication_method": schema.StringAttribute{
							MarkdownDescription: "User authentication method.",
							Computed:            true,
						},
						"multiple_certificate_authentication": schema.BoolAttribute{
							MarkdownDescription: "Authenticate the VPN client using the machine and user certificates.",
							Computed:            true,
						},
						"primary_authentication_server_use_local": schema.BoolAttribute{
							MarkdownDescription: "Use LOCAL FMC as primary authentication server.",
							Computed:            true,
						},
						"primary_authentication_server_id": schema.StringAttribute{
							MarkdownDescription: "Id of the primary authentication RADIUS Server Group or Realm. Use if `primary_authentication_server_use_local` is not set to `true`.",
							Computed:            true,
						},
						"primary_authentication_server_type": schema.StringAttribute{
							MarkdownDescription: "Type of the primary authentication RADIUS Server Group or Realm, like `RadiusServerGroup` or `Realm`.",
							Computed:            true,
						},
						"primary_authentication_fallback_to_local": schema.BoolAttribute{
							MarkdownDescription: "Fallback to LOCAL FMC if primary authentication Server/Realm is not reachable.",
							Computed:            true,
						},
						"saml_and_certificate_username_must_match": schema.BoolAttribute{
							MarkdownDescription: "Allow VPN connection only if the username from the certificate matches the SAML single sign-on username.",
							Computed:            true,
						},
						"primary_authentication_prefill_username_from_certificate": schema.BoolAttribute{
							MarkdownDescription: "Prefill username from certificate.",
							Computed:            true,
						},
						"primary_authentication_prefill_username_from_certificate_map_primary_field": schema.StringAttribute{
							MarkdownDescription: "Map primary field for username.",
							Computed:            true,
						},
						"primary_authentication_prefill_username_from_certificate_map_secondary_field": schema.StringAttribute{
							MarkdownDescription: "Map secondary field for username.",
							Computed:            true,
						},
						"primary_authentication_prefill_username_from_certificate_map_entire_dn": schema.BoolAttribute{
							MarkdownDescription: "Map entire Distinguished Name (DN) as username.",
							Computed:            true,
						},
						"primary_authentication_hide_username_in_login_window": schema.BoolAttribute{
							MarkdownDescription: "Username is pre-filled from the client certificate, but hidden to the user.",
							Computed:            true,
						},
						"secondary_authentication": schema.BoolAttribute{
							MarkdownDescription: "Enable secondary authentication.",
							Computed:            true,
						},
						"secondary_authentication_server_use_local": schema.BoolAttribute{
							MarkdownDescription: "Use LOCAL FMC as secondary authentication server.",
							Computed:            true,
						},
						"secondary_authentication_server_id": schema.StringAttribute{
							MarkdownDescription: "Id of the secondary authentication RADIUS Server Group or Realm. Use if `secondary_authentication_server_use_local` is not set to `true`.",
							Computed:            true,
						},
						"secondary_authentication_server_type": schema.StringAttribute{
							MarkdownDescription: "Type of the secondary authentication RADIUS Server Group or Realm, like `RadiusServerGroup` or `Realm`.",
							Computed:            true,
						},
						"secondary_authentication_fallback_to_local": schema.BoolAttribute{
							MarkdownDescription: "Fallback to LOCAL FMC if secondary authentication Server/Realm is not reachable.",
							Computed:            true,
						},
						"secondary_authentication_prompt_for_username": schema.BoolAttribute{
							MarkdownDescription: "Prompt for username during secondary authentication.",
							Computed:            true,
						},
						"secondary_authentication_use_primary_authentication_username": schema.BoolAttribute{
							MarkdownDescription: "Use primary username for secondary authentication.",
							Computed:            true,
						},
						"use_secondary_authentication_username_for_reporting": schema.BoolAttribute{
							MarkdownDescription: "Secondary username is used for reporting user activity during a VPN session.",
							Computed:            true,
						},
						"saml_use_external_browser": schema.BoolAttribute{
							MarkdownDescription: "Use default OS browser (`true`) or embedded browser (`false`) for SAML authentication.",
							Computed:            true,
						},
						"authorization_server_id": schema.StringAttribute{
							MarkdownDescription: "Id of the authorization RADIUS Server Group or Realm.",
							Computed:            true,
						},
						"authorization_server_type": schema.StringAttribute{
							MarkdownDescription: "Type of the authorization RADIUS Server Group or Realm, like `RadiusServerGroup` or `Realm`.",
							Computed:            true,
						},
						"allow_connection_only_if_user_exists_in_authorization_database": schema.BoolAttribute{
							MarkdownDescription: "Allow connection only if the username of client exists in the authorization database.",
							Computed:            true,
						},
						"accounting_server_id": schema.StringAttribute{
							MarkdownDescription: "Id of the RADIUS accounting server.",
							Computed:            true,
						},
						"accounting_server_type": schema.StringAttribute{
							MarkdownDescription: "Type of the RADIUS accounting server (`RadiusServerGroup`).",
							Computed:            true,
						},
						"strip_realm_from_username": schema.BoolAttribute{
							MarkdownDescription: "Remove the realm from the username before passing the username on to the AAA server.",
							Computed:            true,
						},
						"strip_group_from_username": schema.BoolAttribute{
							MarkdownDescription: "Remove the group name from the username before passing the username on to the AAA server.",
							Computed:            true,
						},
						"password_management": schema.BoolAttribute{
							MarkdownDescription: "Enable managing the password for the remote access VPN users.",
							Computed:            true,
						},
						"password_management_notify_user_on_password_expiry_day": schema.BoolAttribute{
							MarkdownDescription: "Notify user on the day the password expires.",
							Computed:            true,
						},
						"password_management_advance_password_expiration_notification": schema.Int64Attribute{
							MarkdownDescription: "Notify user on the number of days before password expiration.",
							Computed:            true,
						},
						"alias_names": schema.ListNestedAttribute{
							MarkdownDescription: "List of Alias Names.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Name of the Alias.",
										Computed:            true,
									},
									"enabled": schema.BoolAttribute{
										MarkdownDescription: "Enable the alias.",
										Computed:            true,
									},
								},
							},
						},
						"alias_urls": schema.ListNestedAttribute{
							MarkdownDescription: "List of Alias URLs (group URLs).",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"url_object_id": schema.StringAttribute{
										MarkdownDescription: "Id of the URL object.",
										Computed:            true,
									},
									"enabled": schema.BoolAttribute{
										MarkdownDescription: "Enable the alias.",
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

func (d *VPNRAConnectionProfilesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

func (d *VPNRAConnectionProfilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config VPNRAConnectionProfiles

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

	// FMCBUG CSCwq61583 FMC API: RAVPN sub-endpoints are unstable
	var res fmc.Res
	var err error
	for range 5 {
		res, err = d.client.Get(urlPath, reqMods...)
		if err == nil {
			break
		}
		if !strings.Contains(err.Error(), "StatusCode 404") {
			break
		}
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
