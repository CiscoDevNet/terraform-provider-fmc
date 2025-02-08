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
	_ datasource.DataSource              = &DeviceVTEPPolicyDataSource{}
	_ datasource.DataSourceWithConfigure = &DeviceVTEPPolicyDataSource{}
)

func NewDeviceVTEPPolicyDataSource() datasource.DataSource {
	return &DeviceVTEPPolicyDataSource{}
}

type DeviceVTEPPolicyDataSource struct {
	client *fmc.Client
}

func (d *DeviceVTEPPolicyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_vtep_policy"
}

func (d *DeviceVTEPPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the Device VTEP Policy.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the object",
				Required:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "Name of the FMC domain",
				Optional:            true,
			},
			"device_id": schema.StringAttribute{
				MarkdownDescription: "Id of the parent device.",
				Required:            true,
			},
			"nve_enabled": schema.BoolAttribute{
				MarkdownDescription: "Indicates whether to enable NVE on the `device_id`. Can only be false if `vteps` are empty.",
				Computed:            true,
			},
			"vteps": schema.ListNestedAttribute{
				MarkdownDescription: "List that can either be empty or contain one VTEP object.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"source_interface_id": schema.StringAttribute{
							MarkdownDescription: "Id of the source interface. It cannot refer to a subinterface.",
							Computed:            true,
						},
						"nve_number": schema.Int64Attribute{
							MarkdownDescription: "VTEP NVE number, currently must always be 1.",
							Computed:            true,
						},
						"encapsulation_port": schema.Int64Attribute{
							MarkdownDescription: "Encapsulation port number. For VXLAN suggested 4789 (default), for GENEVE suggested 6081.",
							Computed:            true,
						},
						"encapsulation_type": schema.StringAttribute{
							MarkdownDescription: "Encapsulation type.",
							Computed:            true,
						},
						"neighbor_discovery": schema.StringAttribute{
							MarkdownDescription: "How to discover addresses of the neighbor VTEPs for the VTEP-to-VTEP communication. For STATIC_PEER_IP and DEFAULT_MULTICAST_GROUP you must set `neighbor_address_literal` to a single IP address. For STATIC_PEER_GROUP you must however set `neighbor_address_id` to a UUID of a network group and such network group can contain only IPv4 Hosts and IPv4 Ranges (but not Networks, etc.).",
							Computed:            true,
						},
						"neighbor_address_literal": schema.StringAttribute{
							MarkdownDescription: "Used for neighbor_discovery STATIC_PEER_IP, where it holds any unicast IP address. Used for neighbor_discovery DEFAULT_MULTICAST_GROUP, where it holds IP address in range 224.0.0.0 to 239.255.255.255.",
							Computed:            true,
						},
						"neighbor_address_id": schema.StringAttribute{
							MarkdownDescription: "Used for neighbor_discovery STATIC_PEER_GROUP, where it holds UUID of the network group and such network group can contain only IPv4 Hosts and IPv4 Ranges (but not Networks, etc.).",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *DeviceVTEPPolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *DeviceVTEPPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DeviceVTEPPolicy

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
