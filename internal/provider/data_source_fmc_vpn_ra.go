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
	_ datasource.DataSource              = &VPNRADataSource{}
	_ datasource.DataSourceWithConfigure = &VPNRADataSource{}
)

func NewVPNRADataSource() datasource.DataSource {
	return &VPNRADataSource{}
}

type VPNRADataSource struct {
	client *fmc.Client
}

func (d *VPNRADataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_ra"
}

func (d *VPNRADataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the VPN RA.").String,

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
				MarkdownDescription: "Name of the VPN Remote Access (RA) Topology.",
				Optional:            true,
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the object.",
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the object; this value is always 'RAVpn'.",
				Computed:            true,
			},
			"protocol_ssl": schema.BoolAttribute{
				MarkdownDescription: "Enable SSL protocol for the VPN.",
				Computed:            true,
			},
			"protocol_ipsec_ikev2": schema.BoolAttribute{
				MarkdownDescription: "Enable IPsec IKEv2 protocol for the VPN.",
				Computed:            true,
			},
			"local_realm_id": schema.StringAttribute{
				MarkdownDescription: "Local realm server for the VPN.",
				Computed:            true,
			},
			"dap_policy_id": schema.StringAttribute{
				MarkdownDescription: "Identifier for the DAP (Dynamic Access Policy) used for the VPN.",
				Computed:            true,
			},
			"access_interfaces": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Id of interface group or security zone.",
							Computed:            true,
						},
						"protocol_ipsec_ikev2": schema.BoolAttribute{
							MarkdownDescription: "Enable IPsec IKEv2 for the VPN.",
							Computed:            true,
						},
						"protocol_ssl": schema.BoolAttribute{
							MarkdownDescription: "Enable SSL for the VPN.",
							Computed:            true,
						},
						"protocol_ssl_dtls": schema.BoolAttribute{
							MarkdownDescription: "Enable DTLS for the VPN.",
							Computed:            true,
						},
						"interface_specific_certificate_id": schema.StringAttribute{
							MarkdownDescription: "Identifier for the ID certificate used for the VPN.",
							Computed:            true,
						},
					},
				},
			},
			"allow_users_to_select_connection_profile": schema.BoolAttribute{
				MarkdownDescription: "Allow users to select a connection profile.",
				Computed:            true,
			},
			"web_port": schema.Int64Attribute{
				MarkdownDescription: "Port number for the web access of the VPN.",
				Computed:            true,
			},
			"dtls_port": schema.Int64Attribute{
				MarkdownDescription: "Port number for the web access of the VPN.",
				Computed:            true,
			},
			"ssl_global_identity_certificate_id": schema.StringAttribute{
				MarkdownDescription: "Identifier for the SSL certificate used for enrollment.",
				Computed:            true,
			},
			"ipsec_global_identity_certificate_id": schema.StringAttribute{
				MarkdownDescription: "Identifier for the IPsec certificate used for enrollment.",
				Computed:            true,
			},
			"service_access_id": schema.StringAttribute{
				MarkdownDescription: "Identifier for the service access object.",
				Computed:            true,
			},
			"bypass_access_control_policy_for_decrypted_traffic": schema.BoolAttribute{
				MarkdownDescription: "Bypass Access Control policy for decrypted traffic (sysopt permit-vpn)",
				Computed:            true,
			},
			"secure_client_images": schema.ListNestedAttribute{
				MarkdownDescription: "List of Secure Client images to be used for the VPN.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Unique identifier of the Secure Client image.",
							Computed:            true,
						},
						"operating_system": schema.StringAttribute{
							MarkdownDescription: "Operating system of the Secure Client image.",
							Computed:            true,
						},
					},
				},
			},
			"external_browser_package_id": schema.StringAttribute{
				MarkdownDescription: "Identifier for the external browser package used for the VPN.",
				Computed:            true,
			},
			"secure_client_customization_id": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"address_assignment_policy_id": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"certificate_map_id": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"group_policies": schema.ListNestedAttribute{
				MarkdownDescription: "List of group policies associated with the VPN. It is mandatory to include at least 'DfltGrpPolicy' in the list.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Unique identifier of the group policy.",
							Computed:            true,
						},
					},
				},
			},
			"ldap_attribute_map_id": schema.StringAttribute{
				MarkdownDescription: "Identifier for the LDAP attribute mapping used for the VPN.",
				Computed:            true,
			},
			"load_balance_id": schema.StringAttribute{
				MarkdownDescription: "Identifier for the load balancing settings used for the VPN.",
				Computed:            true,
			},
			"ikev2_policies": schema.ListNestedAttribute{
				MarkdownDescription: "List of IKEv2 policies associated with the VPN.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Unique identifier of the IKEv2 policy.",
							Computed:            true,
						},
					},
				},
			},
			"ipsec_advanced_settings_id": schema.StringAttribute{
				MarkdownDescription: "Identifier for the IPsec/IKEv2 advanced settings used for the VPN.",
				Computed:            true,
			},
		},
	}
}
func (d *VPNRADataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *VPNRADataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *VPNRADataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config VPNRA

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
