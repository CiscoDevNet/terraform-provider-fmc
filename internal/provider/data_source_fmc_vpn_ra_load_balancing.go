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
	_ datasource.DataSource              = &VPNRALoadBalancingDataSource{}
	_ datasource.DataSourceWithConfigure = &VPNRALoadBalancingDataSource{}
)

func NewVPNRALoadBalancingDataSource() datasource.DataSource {
	return &VPNRALoadBalancingDataSource{}
}

type VPNRALoadBalancingDataSource struct {
	client *fmc.Client
}

func (d *VPNRALoadBalancingDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_ra_load_balancing"
}

func (d *VPNRALoadBalancingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the VPN RA Load Balancing.").String,

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
				MarkdownDescription: "Id of the parent VPN RA Topology.",
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the object; this value is always 'LoadBalanacing'.",
				Computed:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable load balancing for VPN RA.",
				Computed:            true,
			},
			"ipv4_group_address": schema.StringAttribute{
				MarkdownDescription: "IPv4 Group Address.",
				Computed:            true,
			},
			"ipv6_group_address": schema.StringAttribute{
				MarkdownDescription: "IPv6 Group Address.",
				Computed:            true,
			},
			"interface_id": schema.StringAttribute{
				MarkdownDescription: "Id of the Security Zone or Interface Group which the director and members use to communicate.",
				Computed:            true,
			},
			"port": schema.Int64Attribute{
				MarkdownDescription: "UDP Port for communication between the director and members in a group. The default port is 9023.",
				Computed:            true,
			},
			"ipsec": schema.BoolAttribute{
				MarkdownDescription: "Enable encryption between the director and members.",
				Computed:            true,
			},
			"ipsec_encryption_key": schema.StringAttribute{
				MarkdownDescription: "Encryption key for the IPsec tunnel.",
				Computed:            true,
				Sensitive:           true,
			},
			"send_fqdn_to_peer_devices_instead_of_ip": schema.BoolAttribute{
				MarkdownDescription: "Enable redirection using a Fully Qualified Domain Name (FQDN) instead of an IP address.",
				Computed:            true,
			},
			"ikev2_redirect_phase": schema.StringAttribute{
				MarkdownDescription: "Phase during which the redirection occurs in IKEv2.",
				Computed:            true,
			},
		},
	}
}

func (d *VPNRALoadBalancingDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

func (d *VPNRALoadBalancingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config VPNRALoadBalancing

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

	// FMCBUG CSCwq61583 FMC API: RAVPN sub-endpoints are unstable
	var res fmc.Res
	var err error
	for range 5 {
		res, err = d.client.Get(urlPath, reqMods...)
		if err == nil {
			break
		}
		if !strings.Contains(err.Error(), "StatusCode 404") && !strings.Contains(err.Error(), "StatusCode 400") {
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
