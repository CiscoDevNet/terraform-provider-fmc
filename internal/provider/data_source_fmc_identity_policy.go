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
	_ datasource.DataSource              = &IdentityPolicyDataSource{}
	_ datasource.DataSourceWithConfigure = &IdentityPolicyDataSource{}
)

func NewIdentityPolicyDataSource() datasource.DataSource {
	return &IdentityPolicyDataSource{}
}

type IdentityPolicyDataSource struct {
	client *fmc.Client
}

func (d *IdentityPolicyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_identity_policy"
}

func (d *IdentityPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the Identity Policy.").AddMinimumVersionHeaderDescription().AddMinimumVersionDescription("7.2").String,

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
				MarkdownDescription: "Name of the Identity Policy.",
				Optional:            true,
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the object; this value is always 'IdentityPolicy'.",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the Identity Policy.",
				Computed:            true,
			},
			"identity_mapping_filter_network_object_id": schema.StringAttribute{
				MarkdownDescription: "Id of the Host, Network or Network Group used as an Identity Mapping Filter.",
				Computed:            true,
			},
			"identity_mapping_filter_network_object_type": schema.StringAttribute{
				MarkdownDescription: "Type of the Host, Network or Network Group used as an Identity Mapping Filter.",
				Computed:            true,
			},
			"active_authentication_server_certificate_id": schema.StringAttribute{
				MarkdownDescription: "Id of the Internal Certificate object used as an active authentication server certificate.",
				Computed:            true,
			},
			"active_authentication_server_certificate_type": schema.StringAttribute{
				MarkdownDescription: "Type of the Internal Certificate object used as an active authentication server certificate.",
				Computed:            true,
			},
			"active_authentication_redirect_fqdn_id": schema.StringAttribute{
				MarkdownDescription: "Id of the FQDN object used for active authentication redirection (Redirect to Host Name). Supported only in Snort 3.0 or later.",
				Computed:            true,
			},
			"active_authentication_redirect_fqdn_type": schema.StringAttribute{
				MarkdownDescription: "Type of the FQDN object used for active authentication redirection.",
				Computed:            true,
			},
			"active_authentication_redirect_port": schema.Int64Attribute{
				MarkdownDescription: "Port used for active authentication redirection. Valid values are 885 or 1025-65535.",
				Computed:            true,
			},
			"active_authentication_maximum_login_attempts": schema.Int64Attribute{
				MarkdownDescription: "Maximum login attempts. Use 0 for unlimited attempts.",
				Computed:            true,
			},
			"active_authentication_session_sharing": schema.BoolAttribute{
				MarkdownDescription: "Share active authentication sessions across firewalls. Supported only in 7.4.1 and later.",
				Computed:            true,
			},
			"active_authentication_page": schema.StringAttribute{
				MarkdownDescription: "Active authentication page type.",
				Computed:            true,
			},
			"active_authentication_page_html": schema.StringAttribute{
				MarkdownDescription: "Custom HTML content for the active authentication login page. Only used when `active_authentication_page` is set to 'CUSTOM'.",
				Computed:            true,
			},
			"categories": schema.ListNestedAttribute{
				MarkdownDescription: "Ordered list of Categories created between default categories - Standard Rules and Root Rules. Do not include default categories in the list.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Id of the Category.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Name of the Category.",
							Computed:            true,
						},
					},
				},
			},
			"rules": schema.ListNestedAttribute{
				MarkdownDescription: "Ordered list of Rules.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Id of the Identity Rule.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Name of the Identity Rule.",
							Computed:            true,
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: "Enable rule.",
							Computed:            true,
						},
						"category": schema.StringAttribute{
							MarkdownDescription: "Name of the Category the rule belongs to. Can be one of the default categories (Administrator, Standard or Root Rules) or a user-defined one.",
							Computed:            true,
						},
						"authentication_type": schema.StringAttribute{
							MarkdownDescription: "Authentication type action for the rule.",
							Computed:            true,
						},
						"authentication_realm_id": schema.StringAttribute{
							MarkdownDescription: "Id of the Authentication Realm (AD/LDAP).",
							Computed:            true,
						},
						"authentication_realm_type": schema.StringAttribute{
							MarkdownDescription: "Type of the Authentication Realm.",
							Computed:            true,
						},
						"guest_access_fallback": schema.BoolAttribute{
							MarkdownDescription: "Identify the user as Special Identities/Guest if user cannot be authenticated.",
							Computed:            true,
						},
						"active_authentication_fallback": schema.BoolAttribute{
							MarkdownDescription: "Use active authentication if passive or VPN identity cannot be established.",
							Computed:            true,
						},
						"authentication_protocol": schema.StringAttribute{
							MarkdownDescription: "Protocol used for active authentication.",
							Computed:            true,
						},
						"excluded_user_agents": schema.ListNestedAttribute{
							MarkdownDescription: "List of excluded User Agents (Applications).",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the Application object.",
										Computed:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of the Application object.",
										Computed:            true,
									},
								},
							},
						},
						"source_zones": schema.SetNestedAttribute{
							MarkdownDescription: "Source Security Zones.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the Security Zone object.",
										Computed:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of the Security Zone object.",
										Computed:            true,
									},
								},
							},
						},
						"destination_zones": schema.SetNestedAttribute{
							MarkdownDescription: "Destination Security Zones.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the Security Zone object.",
										Computed:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of the Security Zone object.",
										Computed:            true,
									},
								},
							},
						},
						"source_network_literals": schema.SetNestedAttribute{
							MarkdownDescription: "Source Network literals.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of the Network Object.",
										Computed:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "IP address or network in CIDR format.",
										Computed:            true,
									},
								},
							},
						},
						"source_network_objects": schema.SetNestedAttribute{
							MarkdownDescription: "Objects that represent sources of traffic (Host, Network, Range, FQDN, Network Group, Country, Continent or Geolocation).",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the Network Object.",
										Computed:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of the Network Object.",
										Computed:            true,
									},
								},
							},
						},
						"destination_network_literals": schema.SetNestedAttribute{
							MarkdownDescription: "Destination Network literals.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of the Network Object.",
										Computed:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "IP address or network in CIDR format.",
										Computed:            true,
									},
								},
							},
						},
						"destination_network_objects": schema.SetNestedAttribute{
							MarkdownDescription: "Objects that represent destinations of traffic (Host, Network, Range, FQDN, Network Group, Country, Continent or Geolocation).",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the Network Object.",
										Computed:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of the Network Object.",
										Computed:            true,
									},
								},
							},
						},
						"vlan_tag_literals": schema.SetNestedAttribute{
							MarkdownDescription: "VLAN Tag literals.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of the VLAN tag literal.",
										Computed:            true,
									},
									"start_tag": schema.StringAttribute{
										MarkdownDescription: "Start of the VLAN tag range.",
										Computed:            true,
									},
									"end_tag": schema.StringAttribute{
										MarkdownDescription: "End of the VLAN tag range.",
										Computed:            true,
									},
								},
							},
						},
						"vlan_tag_objects": schema.SetNestedAttribute{
							MarkdownDescription: "VLAN Tag objects.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the VLAN Tag object.",
										Computed:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of the VLAN Tag object.",
										Computed:            true,
									},
								},
							},
						},
						"source_port_literals": schema.SetNestedAttribute{
							MarkdownDescription: "Source Port literals.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of the port literal.",
										Computed:            true,
									},
									"protocol": schema.StringAttribute{
										MarkdownDescription: "IANA protocol number.",
										Computed:            true,
									},
									"port": schema.StringAttribute{
										MarkdownDescription: "Port number.",
										Computed:            true,
									},
								},
							},
						},
						"source_port_objects": schema.SetNestedAttribute{
							MarkdownDescription: "Source Port objects.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the Port object.",
										Computed:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of the Port object.",
										Computed:            true,
									},
								},
							},
						},
						"destination_port_literals": schema.SetNestedAttribute{
							MarkdownDescription: "Destination Port literals.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of the port literal.",
										Computed:            true,
									},
									"port": schema.StringAttribute{
										MarkdownDescription: "Port number.",
										Computed:            true,
									},
									"protocol": schema.StringAttribute{
										MarkdownDescription: "IANA protocol number.",
										Computed:            true,
									},
								},
							},
						},
						"destination_port_objects": schema.SetNestedAttribute{
							MarkdownDescription: "Destination Port objects.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the Port object.",
										Computed:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of the Port object.",
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
func (d *IdentityPolicyDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *IdentityPolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *IdentityPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	// Check if FMC client is connected to supports this object
	if d.client.FMCVersionParsed.LessThan(minFMCVersionIdentityPolicy) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support Identity Policy, minimum required version is 7.2", d.client.FMCVersion))
		return
	}
	var config IdentityPolicy

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
