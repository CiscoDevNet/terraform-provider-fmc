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
	_ datasource.DataSource              = &VPNS2SAdvancedSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &VPNS2SAdvancedSettingsDataSource{}
)

func NewVPNS2SAdvancedSettingsDataSource() datasource.DataSource {
	return &VPNS2SAdvancedSettingsDataSource{}
}

type VPNS2SAdvancedSettingsDataSource struct {
	client *fmc.Client
}

func (d *VPNS2SAdvancedSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_s2s_advanced_settings"
}

func (d *VPNS2SAdvancedSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the VPN S2S Advanced Settings.").String,

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
				MarkdownDescription: "Type of the object; this value is always 'AdvancedSetting'.",
				Computed:            true,
			},
			"ike_keepalive": schema.StringAttribute{
				MarkdownDescription: "IKE keepalive mode.",
				Computed:            true,
			},
			"ike_keepalive_threshold": schema.Int64Attribute{
				MarkdownDescription: "IKE keepalive threshold in seconds.",
				Computed:            true,
			},
			"ike_keepalive_retry_interval": schema.Int64Attribute{
				MarkdownDescription: "IKE keepalive retry interval in seconds.",
				Computed:            true,
			},
			"ike_identity_sent_to_peers": schema.StringAttribute{
				MarkdownDescription: "Identity sent to peer.",
				Computed:            true,
			},
			"ike_peer_identity_validation": schema.StringAttribute{
				MarkdownDescription: "Peer identity validation.",
				Computed:            true,
			},
			"ike_enable_aggressive_mode": schema.BoolAttribute{
				MarkdownDescription: "Enable aggressive mode.",
				Computed:            true,
			},
			"ike_enable_notification_on_tunnel_disconnect": schema.BoolAttribute{
				MarkdownDescription: "Enable notification on tunnel disconnect.",
				Computed:            true,
			},
			"ikev2_cookie_challenge": schema.StringAttribute{
				MarkdownDescription: "Cookie challenge.",
				Computed:            true,
			},
			"ikev2_threshold_to_challenge_incoming_cookies": schema.Int64Attribute{
				MarkdownDescription: "Threshold to challenge incoming cookies in percent.",
				Computed:            true,
			},
			"ikev2_percentage_of_sas_allowed_in_negotiation": schema.Int64Attribute{
				MarkdownDescription: "Percentage of SAs allowed in negotiation.",
				Computed:            true,
			},
			"ikev2_maximum_number_of_sas_allowed_in_negotiation": schema.Int64Attribute{
				MarkdownDescription: "Maximum number of SAs allowed in negotiation.",
				Computed:            true,
			},
			"ipsec_enable_fragmentation_before_encryption": schema.BoolAttribute{
				MarkdownDescription: "Enable fragmentation before encryption.",
				Computed:            true,
			},
			"ipsec_path_maximum_transmission_unit_aging_reset_interval": schema.Int64Attribute{
				MarkdownDescription: "Reset interval in minutes.",
				Computed:            true,
			},
			"nat_keepalive_message_traversal_interval": schema.Int64Attribute{
				MarkdownDescription: "NAT keepalive message traversal interval in seconds.",
				Computed:            true,
			},
			"vpn_idle_timeout": schema.Int64Attribute{
				MarkdownDescription: "VPN idle timeout in minutes.",
				Computed:            true,
			},
			"bypass_access_control_traffic_for_decrypted_traffic": schema.BoolAttribute{
				MarkdownDescription: "Bypass access control traffic for decrypted traffic (sysopt permit-vpn).",
				Computed:            true,
			},
			"use_cert_map_configured_in_endpoint_to_determine_tunnel": schema.BoolAttribute{
				MarkdownDescription: "Use certificate map configured in endpoint to determine tunnel.",
				Computed:            true,
			},
			"use_certificate_ou_to_determine_tunnel": schema.BoolAttribute{
				MarkdownDescription: "Use certificate OU to determine tunnel.",
				Computed:            true,
			},
			"use_ike_identity_ou_to_determine_tunnel": schema.BoolAttribute{
				MarkdownDescription: "Use IKE identity OU to determine tunnel.",
				Computed:            true,
			},
			"use_peer_ip_address_to_determine_tunnel": schema.BoolAttribute{
				MarkdownDescription: "Use peer IP address to determine tunnel.",
				Computed:            true,
			},
		},
	}
}

func (d *VPNS2SAdvancedSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *VPNS2SAdvancedSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config VPNS2SAdvancedSettings

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
