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
	_ datasource.DataSource              = &FTDManualNATRuleDataSource{}
	_ datasource.DataSourceWithConfigure = &FTDManualNATRuleDataSource{}
)

func NewFTDManualNATRuleDataSource() datasource.DataSource {
	return &FTDManualNATRuleDataSource{}
}

type FTDManualNATRuleDataSource struct {
	client *fmc.Client
}

func (d *FTDManualNATRuleDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ftd_manual_nat_rule"
}

func (d *FTDManualNATRuleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the FTD Manual NAT Rule.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the object",
				Required:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "Name of the FMC domain",
				Optional:            true,
			},
			"ftd_nat_policy_id": schema.StringAttribute{
				MarkdownDescription: "Id of the FTD NAT Policy.",
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the object; this value is always 'FTDManualNatRule'.",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of Manual NAT rule.",
				Computed:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Indicates if the rule is enabled.",
				Computed:            true,
			},
			"section": schema.StringAttribute{
				MarkdownDescription: "Name of section to which the rule belongs.",
				Computed:            true,
			},
			"nat_type": schema.StringAttribute{
				MarkdownDescription: "Type of the rule",
				Computed:            true,
			},
			"fall_through": schema.BoolAttribute{
				MarkdownDescription: "Fallthrough to Interface PAT (Destination Interface)",
				Computed:            true,
			},
			"interface_in_original_destination": schema.BoolAttribute{
				MarkdownDescription: "Use interface address as original destination",
				Computed:            true,
			},
			"interface_in_translated_source": schema.BoolAttribute{
				MarkdownDescription: "Translate source network to destination interface address",
				Computed:            true,
			},
			"ipv6": schema.BoolAttribute{
				MarkdownDescription: "Use the IPv6 address of the destination interface for interface PAT.",
				Computed:            true,
			},
			"net_to_net": schema.BoolAttribute{
				MarkdownDescription: "Net to Net Mapping",
				Computed:            true,
			},
			"no_proxy_arp": schema.BoolAttribute{
				MarkdownDescription: "Do not proxy ARP on Destination Interface",
				Computed:            true,
			},
			"unidirectional": schema.BoolAttribute{
				MarkdownDescription: "Whether the rule is unidirectional",
				Computed:            true,
			},
			"source_interface_id": schema.StringAttribute{
				MarkdownDescription: "ID of source security zone or interface group",
				Computed:            true,
			},
			"original_source_id": schema.StringAttribute{
				MarkdownDescription: "ID of original source network object (host, network or range)",
				Computed:            true,
			},
			"original_source_port_id": schema.StringAttribute{
				MarkdownDescription: "ID of original source port object",
				Computed:            true,
			},
			"original_destination_id": schema.StringAttribute{
				MarkdownDescription: "ID of original destination network object (host, network or range)",
				Computed:            true,
			},
			"original_destination_port_id": schema.StringAttribute{
				MarkdownDescription: "ID of original destination port object",
				Computed:            true,
			},
			"route_lookup": schema.BoolAttribute{
				MarkdownDescription: "Perform Route Lookup for Destination Interface",
				Computed:            true,
			},
			"destination_interface_id": schema.StringAttribute{
				MarkdownDescription: "ID of destination security zone or interface group",
				Computed:            true,
			},
			"translated_source_id": schema.StringAttribute{
				MarkdownDescription: "ID of translated source network object (host, network or range)",
				Computed:            true,
			},
			"translated_source_port_id": schema.StringAttribute{
				MarkdownDescription: "ID of translated source port object",
				Computed:            true,
			},
			"translate_dns": schema.BoolAttribute{
				MarkdownDescription: "Translate DNS replies that match this rule",
				Computed:            true,
			},
			"translated_destination_id": schema.StringAttribute{
				MarkdownDescription: "ID of translated destination network object (host, network or range)",
				Computed:            true,
			},
			"translated_destination_port_id": schema.StringAttribute{
				MarkdownDescription: "ID of translated destination port object",
				Computed:            true,
			},
		},
	}
}

func (d *FTDManualNATRuleDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *FTDManualNATRuleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config FTDManualNATRule

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
