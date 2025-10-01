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
	"regexp"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
	_ resource.Resource                = &FTDPlatformSettingsSyslogEventListResource{}
	_ resource.ResourceWithImportState = &FTDPlatformSettingsSyslogEventListResource{}
)

func NewFTDPlatformSettingsSyslogEventListResource() resource.Resource {
	return &FTDPlatformSettingsSyslogEventListResource{}
}

type FTDPlatformSettingsSyslogEventListResource struct {
	client *fmc.Client
}

func (r *FTDPlatformSettingsSyslogEventListResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ftd_platform_settings_syslog_event_list"
}

func (r *FTDPlatformSettingsSyslogEventListResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages FTD Platform Settings - Syslog - Event Lists.").AddMinimumVersionHeaderDescription().AddMinimumVersionDescription("7.7").String,

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
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this value is always 'EventList'.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the event list.").String,
				Required:            true,
			},
			"event_classes": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of event classes.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"class": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Event class.").AddStringEnumDescription("ACCESS_LIST", "APPLICATION_FIREWALL", "AUTH", "BOTNET_TRAFFIC_FILTERING", "BRIDGE", "CA", "CARD_MANAGEMENT", "CLUSTERING", "CONFIG", "CSD", "CTS", "DAP", "EAPOUDP", "EIGRP", "EMAIL", "ENVIRONMENT_MONITORING", "HA", "IDENTITY_BASED_FIREWALL", "IDS", "IKEV2_TOOLKIT", "IP", "IPAA", "IPS", "IPV6", "LICENSING", "MDM_PROXY", "NACPOLICY", "NACSETTINGS", "NAT_AND_PAT", "NETWORK_ACCESS_POINT", "NP", "NP_SSL", "OSPF", "PASSWORD_ENCRYPTION", "PHONE_PROXY", "RIP", "RM", "RULE_ENGINE", "SCANSAFE", "SESSION", "SMART_CALL_HOME", "SNMP", "SSL", "SVC", "SYS", "TAG_SWITCHING", "THREAT_DETECTION", "TRANSACTIONAL_RULE_ENGINE_TRE", "UC_IMS", "VM", "VPDN", "VPN", "VPNC", "VPNFO", "VPNLB", "VXLAN", "WEBFO", "WEBVPN").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("ACCESS_LIST", "APPLICATION_FIREWALL", "AUTH", "BOTNET_TRAFFIC_FILTERING", "BRIDGE", "CA", "CARD_MANAGEMENT", "CLUSTERING", "CONFIG", "CSD", "CTS", "DAP", "EAPOUDP", "EIGRP", "EMAIL", "ENVIRONMENT_MONITORING", "HA", "IDENTITY_BASED_FIREWALL", "IDS", "IKEV2_TOOLKIT", "IP", "IPAA", "IPS", "IPV6", "LICENSING", "MDM_PROXY", "NACPOLICY", "NACSETTINGS", "NAT_AND_PAT", "NETWORK_ACCESS_POINT", "NP", "NP_SSL", "OSPF", "PASSWORD_ENCRYPTION", "PHONE_PROXY", "RIP", "RM", "RULE_ENGINE", "SCANSAFE", "SESSION", "SMART_CALL_HOME", "SNMP", "SSL", "SVC", "SYS", "TAG_SWITCHING", "THREAT_DETECTION", "TRANSACTIONAL_RULE_ENGINE_TRE", "UC_IMS", "VM", "VPDN", "VPN", "VPNC", "VPNFO", "VPNLB", "VXLAN", "WEBFO", "WEBVPN"),
							},
						},
						"severity": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Syslog severity level.").AddStringEnumDescription("EMERG", "ALERT", "CRIT", "ERR", "WARNING", "NOTICE", "INFO", "DEBUG").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("EMERG", "ALERT", "CRIT", "ERR", "WARNING", "NOTICE", "INFO", "DEBUG"),
							},
						},
					},
				},
			},
			"message_ids": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Syslog Message IDs. Use hyphen to specify range of IDs.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
		},
	}
}

func (r *FTDPlatformSettingsSyslogEventListResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *FTDPlatformSettingsSyslogEventListResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	// Check if FMC client is connected to supports this object
	if r.client.FMCVersionParsed.LessThan(minFMCVersionFTDPlatformSettingsSyslogEventList) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support FTD Platform Settings Syslog Event List creation, minumum required version is 7.7", r.client.FMCVersion))
		return
	}
	var plan FTDPlatformSettingsSyslogEventList

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, FTDPlatformSettingsSyslogEventList{})
	res, err := r.client.Post(plan.getPath(), body, reqMods...)
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

func (r *FTDPlatformSettingsSyslogEventListResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Check if FMC client is connected to supports this object
	if r.client.FMCVersionParsed.LessThan(minFMCVersionFTDPlatformSettingsSyslogEventList) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support FTD Platform Settings Syslog Event List, minimum required version is 7.7", r.client.FMCVersion))
		return
	}
	var state FTDPlatformSettingsSyslogEventList

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

func (r *FTDPlatformSettingsSyslogEventListResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state FTDPlatformSettingsSyslogEventList

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

func (r *FTDPlatformSettingsSyslogEventListResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state FTDPlatformSettingsSyslogEventList

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
	res, err := r.client.Delete(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString()), reqMods...)
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *FTDPlatformSettingsSyslogEventListResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var config FTDPlatformSettingsSyslogEventList

	// Parse import ID
	var inputPattern = regexp.MustCompile(`^(?:(?P<domain>[^\s,]+),)?(?P<ftd_platform_settings_id>[^\s,]+),(?P<id>[^\s,]+?)$`)
	match := inputPattern.FindStringSubmatch(req.ID)
	if match == nil {
		errMsg := "Failed to parse import parameters.\nPlease provide import string in the following format: <domain>,<ftd_platform_settings_id>,<id>\n<domain> is optional.\n" + fmt.Sprintf("Got: %q", req.ID)
		resp.Diagnostics.AddError("Import error", errMsg)
		return
	}

	// Set domain, if provided
	if tmpDomain := match[inputPattern.SubexpIndex("domain")]; tmpDomain != "" {
		config.Domain = types.StringValue(tmpDomain)
	}
	config.Id = types.StringValue(match[inputPattern.SubexpIndex("id")])
	config.FtdPlatformSettingsId = types.StringValue(match[inputPattern.SubexpIndex("ftd_platform_settings_id")])

	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
