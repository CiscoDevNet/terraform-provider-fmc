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
	_ datasource.DataSource              = &DeviceOSPFInterfaceDataSource{}
	_ datasource.DataSourceWithConfigure = &DeviceOSPFInterfaceDataSource{}
)

func NewDeviceOSPFInterfaceDataSource() datasource.DataSource {
	return &DeviceOSPFInterfaceDataSource{}
}

type DeviceOSPFInterfaceDataSource struct {
	client *fmc.Client
}

func (d *DeviceOSPFInterfaceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_ospf_interface"
}

func (d *DeviceOSPFInterfaceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the Device OSPF Interface.").AddMinimumVersionHeaderDescription().AddMinimumVersionDescription("7.6").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the object",
				Required:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "Name of the FMC domain",
				Optional:            true,
			},
			"vrf_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the parent VRF.").String,
				Optional:            true,
			},
			"device_id": schema.StringAttribute{
				MarkdownDescription: "Id of the parent device.",
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the object; this is always 'OspfInterface'",
				Computed:            true,
			},
			"interface_id": schema.StringAttribute{
				MarkdownDescription: "ID of the interface associated with this OSPF interface.",
				Computed:            true,
			},
			"default_cost": schema.Int64Attribute{
				MarkdownDescription: "Cost of sending a packet through the interface.",
				Computed:            true,
			},
			"priority": schema.Int64Attribute{
				MarkdownDescription: "Designated router for a network.",
				Computed:            true,
			},
			"mtu_missmatch_ignore": schema.BoolAttribute{
				MarkdownDescription: "Ignore MTU mismatch on the interface.",
				Computed:            true,
			},
			"hello_interval": schema.Int64Attribute{
				MarkdownDescription: "Interval, in seconds, between hello packets sent on an interface.",
				Computed:            true,
			},
			"transmit_delay": schema.Int64Attribute{
				MarkdownDescription: "Estimated time in seconds to send an LSA packet on the interface.",
				Computed:            true,
			},
			"retransmit_interval": schema.Int64Attribute{
				MarkdownDescription: "Time in seconds between LSA retransmissions for adjacencies that belong to the interface.",
				Computed:            true,
			},
			"dead_interval": schema.Int64Attribute{
				MarkdownDescription: "Time period in seconds for which hello packets must not be seen before neighbors indicate that the router is down.",
				Computed:            true,
			},
			"hello_multiplier": schema.Int64Attribute{
				MarkdownDescription: "Number of Hello packets to be sent per second.",
				Computed:            true,
			},
			"point_to_point": schema.BoolAttribute{
				MarkdownDescription: "Configure the interface as point-to-point non-broadcast.",
				Computed:            true,
			},
			"bfd": schema.BoolAttribute{
				MarkdownDescription: "Enable Bidirectional Forwarding Detection (BFD) on the interface.",
				Computed:            true,
			},
			"authentication_password": schema.StringAttribute{
				MarkdownDescription: "Password for authentication.",
				Computed:            true,
				Sensitive:           true,
			},
			"authentication_area_password": schema.StringAttribute{
				MarkdownDescription: "Password for authentication.",
				Computed:            true,
				Sensitive:           true,
			},
			"authentication_area_md5s": schema.ListNestedAttribute{
				MarkdownDescription: "List of MD5 authentication keys.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							MarkdownDescription: "Key ID for the MD5 authentication key.",
							Computed:            true,
						},
						"key": schema.StringAttribute{
							MarkdownDescription: "MD5 authentication key.",
							Computed:            true,
							Sensitive:           true,
						},
					},
				},
			},
			"authentication_md5s": schema.ListNestedAttribute{
				MarkdownDescription: "List of MD5 authentication keys.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							MarkdownDescription: "Key ID for the MD5 authentication key.",
							Computed:            true,
						},
						"key": schema.StringAttribute{
							MarkdownDescription: "MD5 authentication key.",
							Computed:            true,
							Sensitive:           true,
						},
					},
				},
			},
			"authentication_key_chain_id": schema.StringAttribute{
				MarkdownDescription: "Key chain object ID for authentication.",
				Computed:            true,
			},
		},
	}
}

func (d *DeviceOSPFInterfaceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *DeviceOSPFInterfaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	// Check if FMC client is connected to supports this object
	if d.client.FMCVersionParsed.LessThan(minFMCVersionDeviceOSPFInterface) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support Device OSPF Interface, minimum required version is 7.6", d.client.FMCVersion))
		return
	}
	var config DeviceOSPFInterface

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
