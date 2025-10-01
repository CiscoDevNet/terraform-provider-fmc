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
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
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
	_ resource.Resource                = &HealthPolicyResource{}
	_ resource.ResourceWithImportState = &HealthPolicyResource{}
)

func NewHealthPolicyResource() resource.Resource {
	return &HealthPolicyResource{}
}

type HealthPolicyResource struct {
	client *fmc.Client
}

func (r *HealthPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_health_policy"
}

func (r *HealthPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages a Health Policy.\n Due to bug in certain FMC versions, updates are not supported; to change a policy, delete and recreate it.\n Any not configured health module will be created with its default settings.\n Only one ManagementCenterPolicy Health Policy can exist. Due to unability to update policies, ManagementCenterPolicy is not supported.").AddMinimumVersionHeaderDescription().AddMinimumVersionDescription("7.7").String,

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
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the Health Policy.").String,
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description of the Health Policy.").String,
				Optional:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Object type; This is always `HealthPolicy`").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"policy_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Define, if the policy is for the device or for the management center.").AddStringEnumDescription("DevicePolicy").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("DevicePolicy"),
				},
			},
			"is_default_policy": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify if this is the default policy.").String,
				Optional:            true,
			},
			"health_modules": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configuration of health modules.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the health module.").AddStringEnumDescription("hm_disk_status", "hm_asp_drop", "hm_tds", "hm_db", "hm_conn_status_sse", "hm_is5800_powersupply", "hm_card_reset", "hm_ntp_server", "hm_threat_grid_amp", "hm_ftd_ha", "hm_talosagent", "hm_conn_status_amp", "hm_snort_stats", "hm_critical_process", "hm_fsic", "hm_ftd_csdac_identity_services", "hm_deployed_configuration", "hm_ftd_config_resource", "hm_routing_stats", "hm_vpn_stats", "hm_snortstats", "hm_is5800_alarm", "hm_cluster", "hm_bypass", "hm_mu", "hm_reconfig_detection", "hm_xTLS", "hm_pathmonitoring", "hm_cpu", "hm_process", "hm_simu", "hm_fmcaccess_config_change", "hm_linkstate_propagation", "hm_adv_snort_stats", "hm_sdwan", "hm_static_analysis", "hm_chm", "hm_platform_faults", "hm_conn_stats", "hm_fxos_health", "hm_chassis_status_ftd", "hm_du", "hm_ifconfig", "hm_flow_offload", "hm_inlinelink_alarm").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("hm_disk_status", "hm_asp_drop", "hm_tds", "hm_db", "hm_conn_status_sse", "hm_is5800_powersupply", "hm_card_reset", "hm_ntp_server", "hm_threat_grid_amp", "hm_ftd_ha", "hm_talosagent", "hm_conn_status_amp", "hm_snort_stats", "hm_critical_process", "hm_fsic", "hm_ftd_csdac_identity_services", "hm_deployed_configuration", "hm_ftd_config_resource", "hm_routing_stats", "hm_vpn_stats", "hm_snortstats", "hm_is5800_alarm", "hm_cluster", "hm_bypass", "hm_mu", "hm_reconfig_detection", "hm_xTLS", "hm_pathmonitoring", "hm_cpu", "hm_process", "hm_simu", "hm_fmcaccess_config_change", "hm_linkstate_propagation", "hm_adv_snort_stats", "hm_sdwan", "hm_static_analysis", "hm_chm", "hm_platform_faults", "hm_conn_stats", "hm_fxos_health", "hm_chassis_status_ftd", "hm_du", "hm_ifconfig", "hm_flow_offload", "hm_inlinelink_alarm"),
							},
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable health module.").String,
							Optional:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of health module.").AddStringEnumDescription("FTD", "FMC_FTD", "SENSOR", "FMC").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("FTD", "FMC_FTD", "SENSOR", "FMC"),
							},
						},
						"alert_severity": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Severity of health module alert.").AddStringEnumDescription("Critical", "Major", "Warning", "Minor", "Info").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("Critical", "Major", "Warning", "Minor", "Info"),
							},
						},
						"critical_threshold": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Critical threshold value for health module.").AddIntegerRangeDescription(1, 99).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 99),
							},
						},
						"warning_threshold": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Warning threshold value for health module.").AddIntegerRangeDescription(1, 99).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 99),
							},
						},
						"custom_thresholds": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Custom threshold configuration for health module.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of custom threshold.").AddStringEnumDescription("Red-FC", "Yellow-FC").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("Red-FC", "Yellow-FC"),
										},
									},
									"threshold": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Threshold level.").AddIntegerRangeDescription(1, 99).String,
										Required:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 99),
										},
									},
								},
							},
						},
						"alert_configs": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Alert configuration for health module.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Name of the alert configuration.").String,
										Required:            true,
									},
									"enabled": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Enable alert configuration.").String,
										Required:            true,
									},
									"thresholds": schema.ListNestedAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Thresholds for alert configuration.").String,
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"type": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Type of threshold.").AddStringEnumDescription("red", "yellow").String,
													Required:            true,
													Validators: []validator.String{
														stringvalidator.OneOf("red", "yellow"),
													},
												},
												"threshold": schema.Int64Attribute{
													MarkdownDescription: helpers.NewAttributeDescription("Threshold level.").AddIntegerRangeDescription(1, 99).String,
													Required:            true,
													Validators: []validator.Int64{
														int64validator.Between(1, 99),
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"health_module_run_time_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify the interval (in minutes) for running the health modules.").AddIntegerRangeDescription(5, 60).AddDefaultValueDescription("5").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(5, 60),
				},
				Default: int64default.StaticInt64(5),
			},
			"metric_collection_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify the frequency (in minutes) for collection of time series data.").AddIntegerRangeDescription(1, 60).AddDefaultValueDescription("1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 60),
				},
				Default: int64default.StaticInt64(1),
			},
		},
	}
}

func (r *HealthPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *HealthPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	// Check if FMC client is connected to supports this object
	if r.client.FMCVersionParsed.LessThan(minFMCVersionHealthPolicy) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support Health Policy creation, minumum required version is 7.7", r.client.FMCVersion))
		return
	}
	var plan HealthPolicy

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
	body := plan.toBody(ctx, HealthPolicy{})
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

func (r *HealthPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Check if FMC client is connected to supports this object
	if r.client.FMCVersionParsed.LessThan(minFMCVersionHealthPolicy) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support Health Policy, minimum required version is 7.7", r.client.FMCVersion))
		return
	}
	var state HealthPolicy

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

func (r *HealthPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state HealthPolicy

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

func (r *HealthPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state HealthPolicy

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
func (r *HealthPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var config HealthPolicy

	// Parse import ID
	var inputPattern = regexp.MustCompile(`^(?:(?P<domain>[^\s,]+),)?(?P<id>[^\s,]+?)$`)
	match := inputPattern.FindStringSubmatch(req.ID)
	if match == nil {
		errMsg := "Failed to parse import parameters.\nPlease provide import string in the following format: <domain>,<id>\n<domain> is optional.\n" + fmt.Sprintf("Got: %q", req.ID)
		resp.Diagnostics.AddError("Import error", errMsg)
		return
	}

	// Set domain, if provided
	if tmpDomain := match[inputPattern.SubexpIndex("domain")]; tmpDomain != "" {
		config.Domain = types.StringValue(tmpDomain)
	}
	config.Id = types.StringValue(match[inputPattern.SubexpIndex("id")])

	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
