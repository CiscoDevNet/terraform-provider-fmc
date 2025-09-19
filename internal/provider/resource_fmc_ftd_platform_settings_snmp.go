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

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &FTDPlatformSettingsSNMPResource{}
	_ resource.ResourceWithImportState = &FTDPlatformSettingsSNMPResource{}
)

func NewFTDPlatformSettingsSNMPResource() resource.Resource {
	return &FTDPlatformSettingsSNMPResource{}
}

type FTDPlatformSettingsSNMPResource struct {
	client *fmc.Client
}

func (r *FTDPlatformSettingsSNMPResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ftd_platform_settings_snmp"
}

func (r *FTDPlatformSettingsSNMPResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages FTD Platform Settings - SNMP.").AddMinimumVersionHeaderDescription().AddMinimumVersionDescription("7.7").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "Name of the FMC domain",
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"ftd_platform_settings_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the parent FTD Platform Settings.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this value is always 'FTDSNMPPlatformSettings'.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"enable_snmp_server": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable SNMP servers.").String,
				Optional:            true,
			},
			"read_community_string": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Password used by a SNMP management station when sending requests to the threat defense device.").String,
				Optional:            true,
				Sensitive:           true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 32),
				},
			},
			"system_administrator_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the device administrator or other contact person.").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 127),
				},
			},
			"location": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Location of the device.").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 127),
				},
			},
			"snmp_server_port": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("UDP port on which incoming requests will be accepted.").AddIntegerRangeDescription(1, 65535).AddDefaultValueDescription("161").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(161),
			},
			"snmp_management_hosts": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of SNMP management hosts.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_object_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the network object that defines the SNMP management station's host address. This can be an IPv6 host, IPv4 host, IPv4 range or IPv4 subnet.").String,
							Required:            true,
						},
						"snmp_version": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("SNMP version to be used.").AddStringEnumDescription("SNMPv1", "SNMPv2c", "SNMPv3").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("SNMPv1", "SNMPv2c", "SNMPv3"),
							},
						},
						"username": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("(SNMPv3 only) Select SNMPv3 username.").String,
							Optional:            true,
						},
						"read_community_string": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("(SNMPv1, 2c only) Read community string.").String,
							Optional:            true,
							Sensitive:           true,
						},
						"poll": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The management station periodically requests information from the device.").String,
							Optional:            true,
						},
						"trap": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The device sends trap events to the management station as they occur.").String,
							Optional:            true,
						},
						"trap_port": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("SNMP trap UDP port.").AddIntegerRangeDescription(1, 65535).AddDefaultValueDescription("162").String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65535),
							},
							Default: int64default.StaticInt64(162),
						},
						"use_management_interface": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Use the device management interface to reach SNMP management station.").String,
							Optional:            true,
						},
						"interface_literals": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of interfaces to reach SNMP management host.").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"interface_objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of interface objects (Security Zones or Interface Groups) to reach SNMP management host.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the interface object.").String,
										Required:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of the interface object; either 'SecurityZone' or 'InterfaceGroup'.").AddStringEnumDescription("SecurityZone", "InterfaceGroup").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("SecurityZone", "InterfaceGroup"),
										},
									},
									"name": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Name of the interface object.").String,
										Required:            true,
									},
								},
							},
						},
					},
				},
			},
			"snmpv3_users": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of SNMPv3 users.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"security_level": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Select desired security level.").AddStringEnumDescription("Auth", "NoAuth", "Priv").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("Auth", "NoAuth", "Priv"),
							},
						},
						"username": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("SNMPv3 username.").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(0, 32),
							},
						},
						"password_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether `authentication_password` and `encryption_password` are in clear text or encrypted.").AddStringEnumDescription("Clear", "Encrypted").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("Clear", "Encrypted"),
							},
						},
						"authentication_algorithm_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of authentication algorithm.").AddStringEnumDescription("SHA", "SHA224", "SHA256", "SHA384").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("SHA", "SHA224", "SHA256", "SHA384"),
							},
						},
						"authentication_password": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("SNMPv3 authentication password. If you selected Encrypted as the `password_type`, the password must be formatted as xx:xx:xx..., where xx are hexadecimal values.").String,
							Optional:            true,
							Sensitive:           true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(0, 256),
							},
						},
						"encryption_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of encryption algorithm.").AddStringEnumDescription("AES128", "AES192", "AES256").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("AES128", "AES192", "AES256"),
							},
						},
						"encryption_password": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("SNMPv3 encryption password. If you selected Encrypted as the `password_type`, the password must be formatted as xx:xx:xx..., where xx are hexadecimal values.").String,
							Optional:            true,
							Sensitive:           true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(0, 256),
							},
						},
					},
				},
			},
			"trap_syslog": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable transmission of trap-related syslog messages.").String,
				Optional:            true,
			},
			"trap_authentication": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Unauthorized SNMP access. This authentication failure occurs for packets with an incorrect community string.").String,
				Optional:            true,
			},
			"trap_link_up": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("One of the device's communication links has become available.").String,
				Optional:            true,
			},
			"trap_link_down": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("One of the device's communication links has failed.").String,
				Optional:            true,
			},
			"trap_cold_start": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The device is reinitializing itself such that its configuration or the protocol entity implementation may be altered.").String,
				Optional:            true,
			},
			"trap_warm_start": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The device is reinitializing itself such that its configuration and the protocol entity implementation is unaltered.").String,
				Optional:            true,
			},
			"trap_field_replacement_unit_insert": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Field Replaceable Unit (FRU) has been inserted.").String,
				Optional:            true,
			},
			"trap_field_replacement_unit_delete": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Field Replaceable Unit (FRU) has been removed.").String,
				Optional:            true,
			},
			"trap_configuration_change": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("There has been a hardware change.").String,
				Optional:            true,
			},
			"trap_connection_limit_reached": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Connection attempt was rejected because the configured connections limit has been reached.").String,
				Optional:            true,
			},
			"trap_nat_packet_discard": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP packets are discarded by the NAT.").String,
				Optional:            true,
			},
			"trap_cpu_rising_threshold": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("CPU utilization exceeds a predefined threshold for a configured period of time.").String,
				Optional:            true,
			},
			"trap_cpu_rising_threshold_value": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Percent of CPU utilization that triggers a trap.").AddIntegerRangeDescription(10, 94).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(10, 94),
				},
			},
			"trap_cpu_rising_threshold_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Time interval (in minutes) to evaluate the CPU rising threshold.").AddIntegerRangeDescription(1, 60).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 60),
				},
			},
			"trap_memory_rising_threshold": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Memory utilization exceeds a predefined threshold").String,
				Optional:            true,
			},
			"trap_memory_rising_threshold_value": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Percent of memory utilization that triggers a trap.").AddIntegerRangeDescription(50, 95).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(50, 95),
				},
			},
			"trap_failover": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Change in the failover state.").String,
				Optional:            true,
			},
			"trap_cluster": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Change in the cluster state.").String,
				Optional:            true,
			},
			"trap_peer_flap": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("BGP route flapping detected.").String,
				Optional:            true,
			},
		},
	}
}

func (r *FTDPlatformSettingsSNMPResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *FTDPlatformSettingsSNMPResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Get FMC version
	fmcVersion, _ := version.NewVersion(strings.Split(r.client.FMCVersion, " ")[0])

	// Check if FMC client is connected to supports this object
	if fmcVersion.LessThan(minFMCVersionFTDPlatformSettingsSNMP) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support FTD Platform Settings SNMP creation, minumum required version is 7.7", r.client.FMCVersion))
		return
	}
	var plan FTDPlatformSettingsSNMP

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !plan.Domain.IsNull() && plan.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(plan.Domain.ValueString()))
	}
	//// ID needs to be retrieved from FMC, however we are expecting exactly one object
	// Get objects from FMC
	resId, err := r.client.Get(plan.getPath(), reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	// Check if exactly one object is returned
	val := resId.Get("items").Array()
	if len(val) != 1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Expected 1 object, got %d", len(val)))
		return
	}

	// Extract ID from the object
	if retrievedId := val[0].Get("id"); retrievedId.Exists() {
		plan.Id = types.StringValue(retrievedId.String())
		tflog.Debug(ctx, fmt.Sprintf("%s: Found object", plan.Id))
	} else {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object id from payload: %s", resId.String()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, FTDPlatformSettingsSNMP{})
	res, err := r.client.Put(plan.getPath()+"/"+url.PathEscape(plan.Id.ValueString()), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("id").String())
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *FTDPlatformSettingsSNMPResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get FMC version
	fmcVersion, _ := version.NewVersion(strings.Split(r.client.FMCVersion, " ")[0])

	// Check if FMC client is connected to supports this object
	if fmcVersion.LessThan(minFMCVersionFTDPlatformSettingsSNMP) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support FTD Platform Settings SNMP, minimum required version is 7.7", r.client.FMCVersion))
		return
	}
	var state FTDPlatformSettingsSNMP

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !state.Domain.IsNull() && state.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(state.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	urlPath := state.getPath() + "/" + url.QueryEscape(state.Id.ValueString())
	res, err := r.client.Get(urlPath, reqMods...)

	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	imp, diags := helpers.IsFlagImporting(ctx, req)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// After `terraform import` we switch to a full read.
	if imp {
		state.fromBody(ctx, res)
	} else {
		state.fromBodyPartial(ctx, res)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update

func (r *FTDPlatformSettingsSNMPResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state FTDPlatformSettingsSNMP

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Read state
	diags = req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !plan.Domain.IsNull() && plan.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(plan.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	body := plan.toBody(ctx, state)
	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *FTDPlatformSettingsSNMPResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state FTDPlatformSettingsSNMP

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !state.Domain.IsNull() && state.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(state.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	body := state.toBodyPutDelete(ctx)
	res, err := r.client.Put(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString()), body, reqMods...)
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import

func (r *FTDPlatformSettingsSNMPResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <ftd_platform_settings_id>,<id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("ftd_platform_settings_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
