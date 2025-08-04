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
	_ datasource.DataSource              = &VPNRAIPSecIKEParametersDataSource{}
	_ datasource.DataSourceWithConfigure = &VPNRAIPSecIKEParametersDataSource{}
)

func NewVPNRAIPSecIKEParametersDataSource() datasource.DataSource {
	return &VPNRAIPSecIKEParametersDataSource{}
}

type VPNRAIPSecIKEParametersDataSource struct {
	client *fmc.Client
}

func (d *VPNRAIPSecIKEParametersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_ra_ipsec_ike_parameters"
}

func (d *VPNRAIPSecIKEParametersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the VPN RA IPSec IKE Parameters.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the object",
				Required:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "Name of the FMC domain",
				Optional:            true,
			},
			"vpn_ra_id": schema.StringAttribute{
				MarkdownDescription: "Id of the parent VPN RA Configuration.",
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the object; this value is always 'RaVpnIPsecAdvancedSetting'.",
				Computed:            true,
			},
			"ikev2_identity_sent_to_peer": schema.StringAttribute{
				MarkdownDescription: "Identity sent to the peer during IKEv2 session establishment.",
				Computed:            true,
			},
			"ikev2_notification_on_tunnel_disconnect": schema.BoolAttribute{
				MarkdownDescription: "Enable notification on tunnel disconnect.",
				Computed:            true,
			},
			"ikev2_do_not_reboot_until_all_sessions_are_terminated": schema.BoolAttribute{
				MarkdownDescription: "Wait for all active sessions to voluntarily terminate before the system reboots.",
				Computed:            true,
			},
			"ikev2_cookie_challenge": schema.StringAttribute{
				MarkdownDescription: "Whether to send cookie challenges to peer devices in response to SA initiated packets.",
				Computed:            true,
			},
			"ikev2_threshold_to_challenge_incoming_cookies": schema.Int64Attribute{
				MarkdownDescription: "Percentage of the total allowed SAs that are in-negotiation.",
				Computed:            true,
			},
			"ikev2_number_of_sas_allowed_in_negotiation": schema.Int64Attribute{
				MarkdownDescription: "Limits the maximum number of SAs that can be in negotiation at any time.",
				Computed:            true,
			},
			"ikev2_maximum_number_of_sas_allowed": schema.Int64Attribute{
				MarkdownDescription: "Maximum number of Security Associations (SAs) allowed.",
				Computed:            true,
			},
			"ipsec_path_maximum_transmission_unit_aging_reset_interval": schema.Int64Attribute{
				MarkdownDescription: "Enter the number of minutes at which the Path Maximum Transission Unit (PMTU) value of an SA is reset to its original value.",
				Computed:            true,
			},
			"nat_keepalive_message_traversal": schema.BoolAttribute{
				MarkdownDescription: "Enable NAT keepalive message traversal.",
				Computed:            true,
			},
			"nat_keepalive_message_traversal_interval": schema.Int64Attribute{
				MarkdownDescription: "NAT keepalive message traversal interval in seconds.",
				Computed:            true,
			},
		},
	}
}

func (d *VPNRAIPSecIKEParametersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *VPNRAIPSecIKEParametersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config VPNRAIPSecIKEParameters

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
