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
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &VPNS2SIPSECSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &VPNS2SIPSECSettingsDataSource{}
)

func NewVPNS2SIPSECSettingsDataSource() datasource.DataSource {
	return &VPNS2SIPSECSettingsDataSource{}
}

type VPNS2SIPSECSettingsDataSource struct {
	client *fmc.Client
}

func (d *VPNS2SIPSECSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_s2s_ipsec_settings"
}

func (d *VPNS2SIPSECSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the VPN S2S IPSEC Settings.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the object",
				Required:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "Name of the FMC domain",
				Optional:            true,
			},
			"vpn_s2s_id": schema.StringAttribute{
				MarkdownDescription: "Id of the parent VPN S2S Topology.",
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the object; this value is always 'IPSecSetting'.",
				Computed:            true,
			},
			"crypto_map_type": schema.StringAttribute{
				MarkdownDescription: "Type of the crypto map.",
				Computed:            true,
			},
			"ikev2_mode": schema.StringAttribute{
				MarkdownDescription: "IKEv2 mode.",
				Computed:            true,
			},
			"ikev1_ipsec_proposals": schema.SetNestedAttribute{
				MarkdownDescription: "Set of IKEv1 IPSEC proposals.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Id of the IKEv1 IPSEC proposal",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Name of the IKEv1 IPSEC proposal",
							Computed:            true,
						},
					},
				},
			},
			"ikev2_ipsec_proposals": schema.SetNestedAttribute{
				MarkdownDescription: "Set of IKEv2 IPSEC proposals.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Id of the IKEv2 IPSEC proposal",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Name of the IKEv2 IPSEC proposal",
							Computed:            true,
						},
					},
				},
			},
			"security_association_strength_enforcement": schema.BoolAttribute{
				MarkdownDescription: "Enable Security Association (SA) strength enforcement.",
				Computed:            true,
			},
			"reverse_route_injection": schema.BoolAttribute{
				MarkdownDescription: "Enable Reverse Route Injection (RRI).",
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

func (d *VPNS2SIPSECSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *VPNS2SIPSECSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config VPNS2SIPSECSettings

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
