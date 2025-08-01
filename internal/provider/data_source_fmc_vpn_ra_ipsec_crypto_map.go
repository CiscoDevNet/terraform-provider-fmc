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
	_ datasource.DataSource              = &VPNRAIPSecCryptoMapDataSource{}
	_ datasource.DataSourceWithConfigure = &VPNRAIPSecCryptoMapDataSource{}
)

func NewVPNRAIPSecCryptoMapDataSource() datasource.DataSource {
	return &VPNRAIPSecCryptoMapDataSource{}
}

type VPNRAIPSecCryptoMapDataSource struct {
	client *fmc.Client
}

func (d *VPNRAIPSecCryptoMapDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_ra_ipsec_crypto_map"
}

func (d *VPNRAIPSecCryptoMapDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the VPN RA IPSec Crypto Map.").String,

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
			"vpn_ra_id": schema.StringAttribute{
				MarkdownDescription: "Id of the parent VPN RA Topology.",
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the object; this value is always 'RaVpnIPsecCryptoMap'.",
				Computed:            true,
			},
			"interface_id": schema.StringAttribute{
				MarkdownDescription: "Id of the interface object.",
				Optional:            true,
				Computed:            true,
			},
			"ikev2_ipsec_proposals": schema.ListNestedAttribute{
				MarkdownDescription: "List of IKEv2 IPSec proposals",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Id of the IKEv2 IPSec proposal.",
							Computed:            true,
						},
					},
				},
			},
			"reverse_route_injection": schema.BoolAttribute{
				MarkdownDescription: "Enable Reverse Route Injection (RRI).",
				Computed:            true,
			},
			"client_services": schema.BoolAttribute{
				MarkdownDescription: "Enable Client Services.",
				Computed:            true,
			},
			"client_services_port": schema.Int64Attribute{
				MarkdownDescription: "Port for Client Services.",
				Computed:            true,
			},
			"perfect_forward_secrecy": schema.BoolAttribute{
				MarkdownDescription: "Enable IPSEC Perfect Forward Secrecy (PFS).",
				Computed:            true,
			},
			"perfect_forward_secrecy_modulus_group": schema.StringAttribute{
				MarkdownDescription: "Modulus group for IPSEC Perfect Forward Secrecy (PFS).",
				Computed:            true,
			},
			"lifetime_duration": schema.Int64Attribute{
				MarkdownDescription: "Number of seconds a security association exists before expiring.",
				Computed:            true,
			},
			"lifetime_size": schema.Int64Attribute{
				MarkdownDescription: "Volume of traffic (in kilobytes) that can pass between IPsec peers using a given security association before it expires.",
				Computed:            true,
			},
			"validate_incoming_icmp_error_messages": schema.BoolAttribute{
				MarkdownDescription: "Enable incoming ICMP error messages validation.",
				Computed:            true,
			},
			"do_not_fragment_policy": schema.StringAttribute{
				MarkdownDescription: "Policy for handling Do Not Fragment (DNF) packets.",
				Computed:            true,
			},
			"tfc": schema.BoolAttribute{
				MarkdownDescription: "Enable Traffic Flow Confidentiality (TFC) packets.",
				Computed:            true,
			},
			"tfc_burst_bytes": schema.Int64Attribute{
				MarkdownDescription: "Burst size in bytes for TFC packets. Set 0 for `auto` or value in range 1-16.",
				Computed:            true,
			},
			"tfc_payload_bytes": schema.Int64Attribute{
				MarkdownDescription: "Payload size in bytes for TFC packets. Set 0 for `auto` or value in range 64-1024.",
				Computed:            true,
			},
			"tfc_timeout": schema.Int64Attribute{
				MarkdownDescription: "Timeout duration in seconds for TFC packets. Set 0 for `auto` or value in range 10-60.",
				Computed:            true,
			},
		},
	}
}
func (d *VPNRAIPSecCryptoMapDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("interface_id"),
		),
	}
}

func (d *VPNRAIPSecCryptoMapDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *VPNRAIPSecCryptoMapDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config VPNRAIPSecCryptoMap

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
	if config.Id.IsNull() && !config.InterfaceId.IsNull() {
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
					if config.InterfaceId.ValueString() == v.Get("interfaceObject.id").String() {
						config.Id = types.StringValue(v.Get("id").String())
						tflog.Debug(ctx, fmt.Sprintf("%s: Found object with interface_id '%v', id: %v", config.Id.ValueString(), config.InterfaceId.ValueString(), config.Id.ValueString()))
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
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with interface_id: %v", config.InterfaceId.ValueString()))
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
