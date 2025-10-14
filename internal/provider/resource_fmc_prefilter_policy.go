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
	"slices"
	"strings"
	"time"

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &PrefilterPolicyResource{}
	_ resource.ResourceWithImportState = &PrefilterPolicyResource{}
)

func NewPrefilterPolicyResource() resource.Resource {
	return &PrefilterPolicyResource{}
}

type PrefilterPolicyResource struct {
	client *fmc.Client
}

func (r *PrefilterPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_prefilter_policy"
}

func (r *PrefilterPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages Prefilter Policy with corresponding rules.").String,

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
				MarkdownDescription: helpers.NewAttributeDescription("Name of the Prefilter policy.").String,
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description of the policy.").String,
				Optional:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this value is always `PrefilterPolicy`.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"default_action": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specifies the default action to take when none of the rules meet the conditions.").AddStringEnumDescription("BLOCK_TUNNELS", "ANALYZE_TUNNELS").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("BLOCK_TUNNELS", "ANALYZE_TUNNELS"),
				},
			},
			"default_action_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Default action Id.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"default_action_log_begin": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Log events at the beginning of the connection for default action.").String,
				Optional:            true,
			},
			"default_action_log_end": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Log events at the end of the connection for default action.").String,
				Optional:            true,
			},
			"default_action_send_events_to_fmc": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Send events to the Firepower Management Center event viewer for default action.").String,
				Optional:            true,
			},
			"default_action_syslog_alert_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of syslog alert. Can be set only when either `default_action_log_begin` or `default_action_log_end` is true.").String,
				Optional:            true,
			},
			"default_action_snmp_alert_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the SNMP alert. Can be set only when either `default_action_log_begin` or `default_action_log_end` is true.").String,
				Optional:            true,
			},
			"rules": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The ordered list of rules.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the prefilter rule.").String,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the rule.").String,
							Required:            true,
						},
						"rule_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of the rule. At least one Encapsulation Port Object (`encapsulation_ports`) is mandatory to be specified for TUNNEL Rules.").AddStringEnumDescription("PREFILTER", "TUNNEL").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("PREFILTER", "TUNNEL"),
							},
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable the rule.").AddDefaultValueDescription("true").String,
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(true),
						},
						"action": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("What to do when the conditions defined by the rule are met.").AddStringEnumDescription("FASTPATH", "ANALYZE", "BLOCK").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("FASTPATH", "ANALYZE", "BLOCK"),
							},
						},
						"bidirectional": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Indicates whether the rule is bidirectional. Used for TUNNEL rules.").String,
							Optional:            true,
						},
						"tunnel_zone_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of Tunnel Zone. Can be only set for TUNNEL rules with ANALYZE action.").String,
							Optional:            true,
						},
						"time_range_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of Time Range object applied to the rule.").String,
							Optional:            true,
						},
						"source_interfaces": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects that represent source interfaces.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the object.").String,
										Optional:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of the object.").String,
										Optional:            true,
									},
								},
							},
						},
						"destination_interfaces": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects that represent destination interfaces.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the object.").String,
										Optional:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of the object.").String,
										Optional:            true,
									},
								},
							},
						},
						"source_network_literals": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects that represent sources of traffic (literally specified).").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"value": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Network literal value.").String,
										Optional:            true,
									},
								},
							},
						},
						"source_network_objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects that represent sources of traffic (Networks, Hosts, Ranges or Network Groups).").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the object.").String,
										Required:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of the object.").String,
										Required:            true,
									},
								},
							},
						},
						"destination_network_literals": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects that represent destinations of traffic (literally specified).").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"value": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
									},
								},
							},
						},
						"destination_network_objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects that represent destinations of traffic (Networks, Hosts, Ranges or Network Groups).").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the object.").String,
										Required:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of the object").String,
										Required:            true,
									},
								},
							},
						},
						"vlan_tag_literals": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects that represent vlan tags (literally specified).").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"start_tag": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Required:            true,
									},
									"end_tag": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Required:            true,
									},
								},
							},
						},
						"vlan_tag_objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects representing vlan tags or vlan tag groups.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the object.").String,
										Optional:            true,
									},
								},
							},
						},
						"source_port_literals": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects that represent protocol/port (literally specified). Can be only set for PREFILTER rules.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"protocol": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Protocol number.").String,
										Required:            true,
									},
									"port": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Port number.").String,
										Optional:            true,
									},
								},
							},
						},
						"source_port_objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects representing source ports associated with the rule (Port or Port Groups). Can be only set for PREFILTER rules.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the object.").String,
										Optional:            true,
									},
								},
							},
						},
						"destination_port_literals": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects that represent protocol/port (literally specified). Can be only set for PREFILTER rules.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"protocol": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Required:            true,
									},
									"port": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
									},
								},
							},
						},
						"destination_port_objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects representing destination ports associated with the rule (Port or Port Groups). Can be only set for PREFILTER rules.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the object.").String,
										Optional:            true,
									},
								},
							},
						},
						"encapsulation_ports": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of encapsulation ports to be used. Mandatory for TUNNEL rules.").AddStringEnumDescription("GRE", "IP_IN_IP", "IPV6_IN_IP", "TEREDO").String,
							ElementType:         types.StringType,
							Optional:            true,
							Validators: []validator.Set{
								setvalidator.ValueStringsAre(
									stringvalidator.OneOf("GRE", "IP_IN_IP", "IPV6_IN_IP", "TEREDO"),
								),
							},
						},
						"log_begin": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Log events at the beginning of the connection.").String,
							Optional:            true,
						},
						"log_end": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Log events at the end of the connection.").String,
							Optional:            true,
						},
						"send_events_to_fmc": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Send events to the Firepower Management Center event viewer.").String,
							Optional:            true,
						},
						"send_syslog": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Send alerts associated with the prefilter rule to default syslog configuration in Prefilter Logging.").String,
							Optional:            true,
						},
						"syslog_alert_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the syslog alert. Can be set only when `syslog_enabled` is true and either `log_begin` or `log_end` is true. If not set, the default policy syslog configuration in Access Control Logging applies.").String,
							Optional:            true,
						},
						"syslog_severity": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Override the Severity of syslog alerts.").AddStringEnumDescription("ALERT", "CRIT", "DEBUG", "EMERG", "ERR", "INFO", "NOTICE", "WARNING").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("ALERT", "CRIT", "DEBUG", "EMERG", "ERR", "INFO", "NOTICE", "WARNING"),
							},
						},
						"snmp_alert_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the SNMP alert associated with the prefilter rule. Can be set only when either `log_begin` or `log_end` is true.").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *PrefilterPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

func (r *PrefilterPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PrefilterPolicy

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

	planBody := plan.toBody(ctx, PrefilterPolicy{})

	// Create object
	body := planBody
	body, _ = sjson.Delete(body, "dummy_rules")

	res, err := r.client.Post(plan.getPath(), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("id").String())

	read, err := r.client.Get(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))

		res, err := r.client.Delete(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), reqMods...)
		if err != nil {
			resp.Diagnostics.AddWarning("Client Error", fmt.Sprintf("Also, cannot DELETE a hanging policy object, got error: %s, %s", err, res.String()))
		}
		return
	}

	plan.fromBodyUnknowns(ctx, read)

	state := plan
	state.Rules = nil

	state, diags = r.updateSubresources(ctx, req.Plan, plan, planBody, tfsdk.State{}, state)
	resp.Diagnostics.Append(diags...)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished", state.Id.ValueString()))

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *PrefilterPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PrefilterPolicy

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
	resGet, err := r.client.Get(urlPath, reqMods...)

	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve prefilter policy object (GET), got error: %s, %s", err, resGet.String()))
		return
	}

	// Default action properties (like logging settings) are not part of the main object
	resDefaultAction, err := r.client.Get(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString())+"/defaultactions/"+state.DefaultActionId.String(), reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve default action object (GET), got error: %s, %s", err, resDefaultAction.String()))
		return
	}

	resRules, err := r.client.Get(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString())+"/prefilterrules?expanded=true", reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve rules object (GET), got error: %s, %s", err, resRules.String()))
		return
	}

	s := resGet.String()

	replaceRules := resRules.Get("items").String()
	if replaceRules == "" {
		replaceRules = "[]"
	}
	s, _ = sjson.SetRaw(s, "dummy_rules", replaceRules)
	s, _ = sjson.SetRaw(s, "defaultAction", resDefaultAction.String())

	res := gjson.Parse(s)

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

func (r *PrefilterPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state PrefilterPolicy

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

	planBody := plan.toBody(ctx, state)
	body := planBody
	body, _ = sjson.Delete(body, "dummy_rules")
	// default_action_id is computed, but still needs to be part of update request
	body, _ = sjson.Set(body, "defaultAction.id", state.DefaultActionId.ValueString())

	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.fromBodyUnknowns(ctx, res)

	orig := state
	state = plan
	state.Rules = orig.Rules

	state, diags = r.updateSubresources(ctx, req.Plan, plan, planBody, req.State, state)
	resp.Diagnostics.Append(diags...)

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *PrefilterPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PrefilterPolicy

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
func (r *PrefilterPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var config PrefilterPolicy

	// Parse import ID
	var inputPattern = regexp.MustCompile(`^(?:(?P<domain>[^\s,]+),)?(?P<id>[^\s,]+?)$`)
	match := inputPattern.FindStringSubmatch(req.ID)
	if match == nil {
		errMsg := "Failed to parse import parameters.\nPlease provide import string in the following format: <domain>,<id>\n<domain> is optional. If not provided, `Global` is used implicitly and resource's `domain` attribute is not set.\n" + fmt.Sprintf("Got: %q", req.ID)
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

func (r *PrefilterPolicyResource) updateSubresources(ctx context.Context, tfsdkPlan tfsdk.Plan, plan PrefilterPolicy, planBody string, tfsdkState tfsdk.State, state PrefilterPolicy) (PrefilterPolicy, diag.Diagnostics) {
	var diags diag.Diagnostics

	p := gjson.Parse(planBody)
	bodyRules := p.Get("dummy_rules").Array()

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !plan.Domain.IsNull() && plan.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(plan.Domain.ValueString()))
	}

	keptRules := 0

	err := r.truncateRulesAt(ctx, &state, keptRules, reqMods...)
	if err != nil {
		diags.AddError("Client Error", err.Error())
		return state, diags
	}

	if len(plan.Rules) == 0 {
		state.Rules = plan.Rules
	}

	err = r.createRulesAt(ctx, plan, bodyRules, keptRules, &state, reqMods...)
	if err != nil {
		diags.AddError("Client Error", err.Error())
		return state, diags
	}

	return state, diags
}

func (r *PrefilterPolicyResource) truncateRulesAt(ctx context.Context, state *PrefilterPolicy, kept int, reqMods ...func(*fmc.Req)) error {
	var b strings.Builder
	var bulks []string
	var counts []int
	count := 0

	b.Grow(maxUrlParamLength + 100)

	for i := kept; i < len(state.Rules); i++ {
		if b.Len() != 0 {
			b.WriteString(",")
		}
		b.WriteString(state.Rules[i].Id.ValueString())
		count++
		if b.Len() >= maxUrlParamLength {
			bulks = append(bulks, b.String())
			b.Reset()
			counts = append(counts, count)
			count = 0
		}
	}
	if b.Len() > 0 {
		bulks = append(bulks, b.String())
		counts = append(counts, count)
	}

	defer func() {
		// Apparently, the bulk DELETE has a race. Stabilize:
		time.Sleep(2 * time.Second)
	}()

	for i, bulk := range bulks {
		res, err := r.client.Delete(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString())+
			"/prefilterrules?bulk=true&filter=ids:"+url.QueryEscape(bulk), reqMods...)
		if err != nil {
			return fmt.Errorf("Failed to bulk-delete rules, got error: %v, %s", err, res.String())
		}

		state.Rules = slices.Delete(state.Rules, 0, counts[i])
	}

	return nil
}

func (r *PrefilterPolicyResource) createRulesAt(ctx context.Context, plan PrefilterPolicy, body []gjson.Result, startIndex int, state *PrefilterPolicy, reqMods ...func(*fmc.Req)) error {
	for i := startIndex; i < len(body); i++ {
		bulk := `{"dummy_rules":[]}`
		j := i
		bulkCount := 0
		bodyLength := 0
		for ; i < len(body); i++ {
			rule := body[i].String()
			rule, _ = sjson.Delete(rule, "id")

			// Check if the body is too big for a single POST
			bodyLength += len(rule)
			if bodyLength >= maxPayloadSize {
				i--
				break
			}

			bulk, _ = sjson.SetRaw(bulk, "dummy_rules.-1", rule)

			// Count the number of rules in the bulk
			bulkCount++
			if bulkCount >= bulkSizeCreate {
				break
			}
		}

		res, err := r.client.Post(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString())+"/prefilterrules?bulk=true",
			gjson.Parse(bulk).Get("dummy_rules").String(),
			reqMods...)
		if err != nil {
			return fmt.Errorf("failed to configure object (POST), got error: %s, %s", err, res.String())
		}

		for _, v := range res.Get("items").Array() {
			if v.Get("name").String() == "" {
				tflog.Debug(ctx, fmt.Sprintf("ignoring a bogus item in POST reply, it can happen: %s", v))
				continue
			}

			item := plan.Rules[j]
			item.Id = types.StringValue(v.Get("id").String())

			if len(state.Rules) <= j {
				state.Rules = append(state.Rules, item)
			} else {
				state.Rules[j] = item
			}

			j++
		}
	}

	return nil
}
