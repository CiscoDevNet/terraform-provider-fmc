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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &FTDPlatformSettingsSNMPDataSource{}
	_ datasource.DataSourceWithConfigure = &FTDPlatformSettingsSNMPDataSource{}
)

func NewFTDPlatformSettingsSNMPDataSource() datasource.DataSource {
	return &FTDPlatformSettingsSNMPDataSource{}
}

type FTDPlatformSettingsSNMPDataSource struct {
	client *fmc.Client
}

func (d *FTDPlatformSettingsSNMPDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ftd_platform_settings_snmp"
}

func (d *FTDPlatformSettingsSNMPDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the FTD Platform Settings SNMP.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the object",
				Required:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "Name of the FMC domain",
				Optional:            true,
			},
			"ftd_platform_settings_id": schema.StringAttribute{
				MarkdownDescription: "Id of the parent FTD Platform Settings.",
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the object; this value is always 'FTDSNMPPlatformSettings'.",
				Computed:            true,
			},
			"enable_snmp_servers": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP servers.",
				Computed:            true,
			},
			"read_community_string": schema.StringAttribute{
				MarkdownDescription: "Password used by a SNMP management station when sending requests to the threat defense device.",
				Computed:            true,
				Sensitive:           true,
			},
			"system_administrator_name": schema.StringAttribute{
				MarkdownDescription: "Name of the device administrator or other contact person.",
				Computed:            true,
			},
			"location": schema.StringAttribute{
				MarkdownDescription: "Location of the device.",
				Computed:            true,
			},
			"listen_port": schema.Int64Attribute{
				MarkdownDescription: "UDP port on which incoming requests will be accepted.",
				Computed:            true,
			},
			"snmp_management_hosts": schema.ListNestedAttribute{
				MarkdownDescription: "List of SNMP management hosts.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_object_id": schema.StringAttribute{
							MarkdownDescription: "Id of the network object that defines the SNMP management station's host address. This can be an IPv6 host, IPv4 host, IPv4 range or IPv4 subnet.",
							Computed:            true,
						},
						"snmp_version": schema.StringAttribute{
							MarkdownDescription: "SNMP version to be used.",
							Computed:            true,
						},
						"username": schema.StringAttribute{
							MarkdownDescription: "(SNMPv3 only) Select SNMPv3 username.",
							Computed:            true,
						},
						"read_community_string": schema.StringAttribute{
							MarkdownDescription: "(SNMPv1, 2c only) Read community string.",
							Computed:            true,
							Sensitive:           true,
						},
						"poll": schema.BoolAttribute{
							MarkdownDescription: "The management station periodically requests information from the device.",
							Computed:            true,
						},
						"trap": schema.BoolAttribute{
							MarkdownDescription: "The device sends trap events to the management station as they occur.",
							Computed:            true,
						},
						"trap_port": schema.Int64Attribute{
							MarkdownDescription: "SNMP trap UDP port.",
							Computed:            true,
						},
						"use_device_management_interface": schema.BoolAttribute{
							MarkdownDescription: "Use the device management interface to reach SNMP management station.",
							Computed:            true,
						},
						"interface_literals": schema.SetAttribute{
							MarkdownDescription: "List of interfaces to reach SNMP management host.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"interface_objects": schema.SetNestedAttribute{
							MarkdownDescription: "List of interface objects (Security Zones or Interface Groups) to reach SNMP management host.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the interface object.",
										Computed:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of the interface object; either 'SecurityZone' or 'InterfaceGroup'.",
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: "Name of the interface object.",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
			"snmpv3_users": schema.ListNestedAttribute{
				MarkdownDescription: "List of SNMPv3 users.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"security_level": schema.StringAttribute{
							MarkdownDescription: "Select desired security level.",
							Computed:            true,
						},
						"username": schema.StringAttribute{
							MarkdownDescription: "SNMPv3 username.",
							Computed:            true,
						},
						"password_type": schema.StringAttribute{
							MarkdownDescription: "Whether `authentication_password` and `encryption_password` are in clear text or encrypted.",
							Computed:            true,
						},
						"authentication_algorithm_type": schema.StringAttribute{
							MarkdownDescription: "Type of authentication algorithm.",
							Computed:            true,
						},
						"authentication_password": schema.StringAttribute{
							MarkdownDescription: "SNMPv3 authentication password. If you selected Encrypted as the `password_type`, the password must be formatted as xx:xx:xx..., where xx are hexadecimal values.",
							Computed:            true,
							Sensitive:           true,
						},
						"encryption_type": schema.StringAttribute{
							MarkdownDescription: "Type of encryption algorithm.",
							Computed:            true,
						},
						"encryption_password": schema.StringAttribute{
							MarkdownDescription: "SNMPv3 encryption password. If you selected Encrypted as the `password_type`, the password must be formatted as xx:xx:xx..., where xx are hexadecimal values.",
							Computed:            true,
							Sensitive:           true,
						},
					},
				},
			},
			"trap_syslog": schema.BoolAttribute{
				MarkdownDescription: "Enable transmission of trap-related syslog messages.",
				Computed:            true,
			},
			"trap_authentication": schema.BoolAttribute{
				MarkdownDescription: "Unauthorized SNMP access. This authentication failure occurs for packets with an incorrect community string.",
				Computed:            true,
			},
			"trap_link_up": schema.BoolAttribute{
				MarkdownDescription: "One of the device's communication links has become available.",
				Computed:            true,
			},
			"trap_link_down": schema.BoolAttribute{
				MarkdownDescription: "One of the device's communication links has failed.",
				Computed:            true,
			},
			"trap_cold_start": schema.BoolAttribute{
				MarkdownDescription: "The device is reinitializing itself such that its configuration or the protocol entity implementation may be altered.",
				Computed:            true,
			},
			"trap_warm_start": schema.BoolAttribute{
				MarkdownDescription: "The device is reinitializing itself such that its configuration and the protocol entity implementation is unaltered.",
				Computed:            true,
			},
			"trap_field_replacement_unit_insert": schema.BoolAttribute{
				MarkdownDescription: "Field Replaceable Unit (FRU) has been inserted.",
				Computed:            true,
			},
			"trap_field_replacement_unit_delete": schema.BoolAttribute{
				MarkdownDescription: "Field Replaceable Unit (FRU) has been removed.",
				Computed:            true,
			},
			"trap_configuration_change": schema.BoolAttribute{
				MarkdownDescription: "There has been a hardware change.",
				Computed:            true,
			},
			"trap_connection_limit_reached": schema.BoolAttribute{
				MarkdownDescription: "Connection attempt was rejected because the configured connections limit has been reached.",
				Computed:            true,
			},
			"trap_nat_packet_discard": schema.BoolAttribute{
				MarkdownDescription: "IP packets are discarded by the NAT.",
				Computed:            true,
			},
			"trap_cpu_rising_threshold": schema.BoolAttribute{
				MarkdownDescription: "CPU utilization exceeds a predefined threshold for a configured period of time.",
				Computed:            true,
			},
			"trap_cpu_rising_threshold_value": schema.Int64Attribute{
				MarkdownDescription: "Percent of CPU utilization that triggers a trap.",
				Computed:            true,
			},
			"trap_cpu_rising_threshold_interval": schema.Int64Attribute{
				MarkdownDescription: "Time interval (in minutes) to evaluate the CPU rising threshold.",
				Computed:            true,
			},
			"trap_memory_rising_threshold": schema.BoolAttribute{
				MarkdownDescription: "Memory utilization exceeds a predefined threshold",
				Computed:            true,
			},
			"trap_memory_rising_threshold_value": schema.Int64Attribute{
				MarkdownDescription: "Percent of memory utilization that triggers a trap.",
				Computed:            true,
			},
			"trap_failover": schema.BoolAttribute{
				MarkdownDescription: "Change in the failover state.",
				Computed:            true,
			},
			"trap_cluster": schema.BoolAttribute{
				MarkdownDescription: "Change in the cluster state.",
				Computed:            true,
			},
			"trap_peer_flap": schema.BoolAttribute{
				MarkdownDescription: "BGP route flapping detected.",
				Computed:            true,
			},
		},
	}
}

func (d *FTDPlatformSettingsSNMPDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *FTDPlatformSettingsSNMPDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config FTDPlatformSettingsSNMP

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
