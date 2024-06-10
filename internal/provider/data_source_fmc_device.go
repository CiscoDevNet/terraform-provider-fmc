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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &DeviceDataSource{}
	_ datasource.DataSourceWithConfigure = &DeviceDataSource{}
)

func NewDeviceDataSource() datasource.DataSource {
	return &DeviceDataSource{}
}

type DeviceDataSource struct {
	client *fmc.Client
}

func (d *DeviceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device"
}

func (d *DeviceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Device.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "The name of the FMC domain",
				Optional:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "User-specified name, must be unique. Example: 'Device 01 - 192.168.0.152'",
				Computed:            true,
			},
			"host_name": schema.StringAttribute{
				MarkdownDescription: "Hostname or IP address of the device. Either the host_name or nat_id must be present.",
				Computed:            true,
			},
			"nat_id": schema.StringAttribute{
				MarkdownDescription: "(used for device registration behind NAT) If the device to be registered and the Firepower Management Center are separated by network address translation (NAT), set a unique string identifier.",
				Computed:            true,
			},
			"license_capabilities": schema.SetAttribute{
				MarkdownDescription: "Array of strings representing the license capabilities on the managed device. For registering FTD, the allowed values are: BASE (mandatory), THREAT, URLFilter, MALWARE, APEX, PLUS, VPNOnly. For Firepower ASA or NGIPSv devices, allowed values are: BASE, THREAT, PROTECT, CONTROL, URLFilter, MALWARE, VPN, SSL.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"registration_key": schema.StringAttribute{
				MarkdownDescription: "Registration Key identical to the one previously configured on the device (`configure manager`).",
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the device; this value is always 'Device'.",
				Computed:            true,
			},
			"access_policy_id": schema.StringAttribute{
				MarkdownDescription: "The UUID of the assigned access control policy. For example `fmc_access_control_policy.example.id`.",
				Computed:            true,
			},
			"nat_policy_id": schema.StringAttribute{
				MarkdownDescription: "The UUID of the assigned NAT policy.",
				Computed:            true,
			},
			"prohibit_packet_transfer": schema.BoolAttribute{
				MarkdownDescription: "Value true prohibits the device from sending packet data with events to the Firepower Management Center. Value false allows the transfer when a certain event is triggered. Not all traffic data is sent; connection events do not include a payload, only connection metadata.",
				Computed:            true,
			},
			"performance_tier": schema.StringAttribute{
				MarkdownDescription: "Performance tier for the managed device, applicable only to vFTD devices >=6.8.0.",
				Computed:            true,
			},
		},
	}
}

func (d *DeviceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

func (d *DeviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config Device

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

	res, err := d.client.Get(config.getPath()+"/"+url.QueryEscape(config.Id.ValueString()), reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	policies, err := d.client.Get("/api/fmc_config/v1/domain/{DOMAIN_UUID}/assignment/policyassignments?offset=0&limit=1000&expanded=true", reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromPolicyBody(ctx, policies)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
