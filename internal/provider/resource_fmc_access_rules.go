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
	"slices"
	"strings"
	"time"

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework-validators/resourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
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
	_ resource.Resource = &AccessRulesResource{}
)

func NewAccessRulesResource() resource.Resource {
	return &AccessRulesResource{}
}

type AccessRulesResource struct {
	client *fmc.Client
}

func (r *AccessRulesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_access_rules"
}

func (r *AccessRulesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This is BETA resource and it's behaviour may change in the future releases.\n This resource manages Access Rules in Access Control Policies in bulks.\n Order of the rules is meant to be preserved within the resource, however NOT between multiple `fmc_access_rule` resources that create rules within a single category/section.\n Any change to the rule set will trigger recreation of all the rules. This is done to preserve the order of the rules. This usually means, that re-created rules will be placed at the end of the policy section/category.\n").String,

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
			"access_policy_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the Access Control Policy.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"category_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the category that owns this rule. Either 'section' or 'category_name' can be set.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"section": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The section of the policy to which the rule belongs. Either 'section' or 'category_name' can be set.").AddStringEnumDescription("default", "mandatory").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("default", "mandatory"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"items": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Ordered list of Access Rules. Rules must be sorted in the order of the corresponding categories, if they have `category_name`. Uncategorized non-mandatory rules must be below all other rules.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the Access Rule.").String,
							Computed:            true,
						},
						"action": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Rule action.").AddStringEnumDescription("ALLOW", "TRUST", "BLOCK", "MONITOR", "BLOCK_RESET", "BLOCK_INTERACTIVE", "BLOCK_RESET_INTERACTIVE").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("ALLOW", "TRUST", "BLOCK", "MONITOR", "BLOCK_RESET", "BLOCK_INTERACTIVE", "BLOCK_RESET_INTERACTIVE"),
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the Access Rule. This name needs to be uqique within the policy.").String,
							Required:            true,
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Indicates whether the access rule is in effect (true) or not (false).").AddDefaultValueDescription("true").String,
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(true),
						},
						"source_network_literals": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects that represent sources of traffic (literally specified).").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"value": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("IP address or network in CIDR format.").String,
										Optional:            true,
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
										MarkdownDescription: helpers.NewAttributeDescription("IP address or network in CIDR format.").String,
										Optional:            true,
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
										MarkdownDescription: helpers.NewAttributeDescription("Start of the VLAN tag range.").String,
										Optional:            true,
									},
									"end_tag": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("End of the VLAN tag range.").String,
										Optional:            true,
									},
								},
							},
						},
						"vlan_tag_objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects that represent vlan tags or vlan tags group.").String,
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
						"source_network_objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects that represent sources of traffic (Host, Network, Range, FQDN or Network Group).").String,
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
						"destination_network_objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects that represent destinations of traffic (Host, Network, Range, FQDN or Network Group).").String,
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
						"source_dynamic_objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects that represent dynamic sources of traffic.").String,
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
						"destination_dynamic_objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects that represent dynamic destinations of traffic.").String,
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
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects that represent protocol/port (literally specified).").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of the object.").AddStringEnumDescription("PortLiteral", "ICMPv4PortLiteral").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("PortLiteral", "ICMPv4PortLiteral"),
										},
									},
									"port": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Port number.").String,
										Optional:            true,
									},
									"protocol": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("IANA protocol number.").String,
										Required:            true,
									},
									"icmp_type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("ICMP type.").String,
										Optional:            true,
									},
									"icmp_code": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("ICMP code.").String,
										Optional:            true,
									},
								},
							},
						},
						"source_port_objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects representing source ports associated with the rule.").String,
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
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects that represent protocol/port (literally specified).").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of the object.").AddStringEnumDescription("PortLiteral", "ICMPv4PortLiteral").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("PortLiteral", "ICMPv4PortLiteral"),
										},
									},
									"port": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Port number.").String,
										Optional:            true,
									},
									"protocol": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("IANA protocol number.").String,
										Required:            true,
									},
									"icmp_type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("ICMP type.").String,
										Optional:            true,
									},
									"icmp_code": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("ICMP code.").String,
										Optional:            true,
									},
								},
							},
						},
						"destination_port_objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects representing destination ports associated with the rule.").String,
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
						"source_sgt_objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects representing the source Security Group Tags (SGT) or ISE Security Group Tags.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Name of the object.").String,
										Required:            true,
									},
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
						"endpoint_device_types": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects representing the source Endpoint Device Types.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Name of the object.").String,
										Required:            true,
									},
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
						"destination_sgt_objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects representing the destination ISE Security Group Tags (SGT).").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Name of the object.").String,
										Required:            true,
									},
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
						"source_zones": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects representing source Security Zones associated with the access rule.").String,
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
						"destination_zones": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects representing destination Security Zones associated with the access rule.").String,
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
						"url_literals": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects representing the URLs associated with the rule (literally specified).").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"url": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("URL such as https://www.example.com/app").String,
										Required:            true,
									},
								},
							},
						},
						"url_objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects representing the URLs associated with the rule.").String,
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
						"url_categories": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects representing the URL Categories associated with the rule.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the object.").String,
										Optional:            true,
									},
									"reputation": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Reputation applicable to the URL Category.").AddStringEnumDescription("ANY_EXCEPT_UNKNOWN", "TRUSTED", "FAVORABLE", "NEUTRAL", "QUESTIONABLE", "UNTRUSTED", "ANY_AND_UNKNOWN", "TRUSTED_AND_UNKNOWN", "FAVORABLE_AND_UNKNOWN", "NEUTRAL_AND_UNKNOWN", "QUESTIONABLE_AND_UNKNOWN", "UNTRUSTED_AND_UNKNOWN").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("ANY_EXCEPT_UNKNOWN", "TRUSTED", "FAVORABLE", "NEUTRAL", "QUESTIONABLE", "UNTRUSTED", "ANY_AND_UNKNOWN", "TRUSTED_AND_UNKNOWN", "FAVORABLE_AND_UNKNOWN", "NEUTRAL_AND_UNKNOWN", "QUESTIONABLE_AND_UNKNOWN", "UNTRUSTED_AND_UNKNOWN"),
										},
									},
								},
							},
						},
						"log_begin": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Log events at the beginning of the connection. If 'MONITOR' action is selected for access rule, log_begin must be false or absent.").AddDefaultValueDescription("false").String,
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"log_end": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Log events at the end of the connection. If 'MONITOR' action is selected for access rule, log_end must be true.").AddDefaultValueDescription("false").String,
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"log_files": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Log file events.").AddDefaultValueDescription("false").String,
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"send_events_to_fmc": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Send events to the Firepower Management Center event viewer. If 'MONITOR' action is selected for access rule, send_events_to_fmc must be true.").AddDefaultValueDescription("false").String,
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"send_syslog": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Send alerts to syslog.").AddDefaultValueDescription("false").String,
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"syslog_config_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of Syslog Config. Can be set only when send_syslog is true and either log_begin or log_end is true. If not set, the default syslog configuration in Access Control Policy Logging applies.").String,
							Optional:            true,
						},
						"syslog_severity": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Override the Severity of syslog alerts.").AddStringEnumDescription("ALERT", "CRIT", "DEBUG", "EMERG", "ERR", "INFO", "NOTICE", "WARNING").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("ALERT", "CRIT", "DEBUG", "EMERG", "ERR", "INFO", "NOTICE", "WARNING"),
							},
						},
						"snmp_config_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the SNMP alert associated with the access rule. Can be set only when either log_begin or log_end is true.").String,
							Optional:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Rule description.").String,
							Optional:            true,
						},
						"file_policy_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the File Policy for the rule action. Cannot be set when action is BLOCK, BLOCK_RESET, TRUST, MONITOR.").String,
							Optional:            true,
						},
						"intrusion_policy_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the Intrusion Policy for the rule action. Cannot be set when action is BLOCK, BLOCK_RESET, TRUST, MONITOR.").String,
							Optional:            true,
						},
						"time_range_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of Time Range object applied to the rule.").String,
							Optional:            true,
						},
						"variable_set_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the Variable Set for the rule action.").String,
							Optional:            true,
						},
						"applications": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of applications.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the Application.").String,
										Required:            true,
									},
								},
							},
						},
						"application_filter_objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of Application Filtering objects.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the Application Filter.").String,
										Required:            true,
									},
								},
							},
						},
						"application_filters": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of Application Filtering conditions.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"types": schema.SetNestedAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set of Application Types.").String,
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Id of the Application Type.").String,
													Required:            true,
												},
											},
										},
									},
									"risks": schema.SetNestedAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set of Application Risk levels.").String,
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Id of the Application Risk level.").String,
													Required:            true,
												},
											},
										},
									},
									"business_relevances": schema.SetNestedAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set of Application Business Relevance levels.").String,
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Id of the Application Business Relevance level.").String,
													Required:            true,
												},
											},
										},
									},
									"categories": schema.SetNestedAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set of Application Categories.").String,
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Id of the Application Category.").String,
													Required:            true,
												},
											},
										},
									},
									"tags": schema.SetNestedAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set of Application Tags.").String,
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Id of the Application Tag.").String,
													Required:            true,
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
	}
}

func (r *AccessRulesResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

func (r AccessRulesResource) ConfigValidators(ctx context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{
		resourcevalidator.Conflicting(
			path.MatchRoot("category_name"),
			path.MatchRoot("section"),
		),
	}
}

func (r *AccessRulesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan AccessRules

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

	// Create new UUID for the bulk resource
	plan.Id = types.StringValue(uuid.NewString())

	// Create rules
	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))
	err := r.createRulesAt(ctx, plan, &plan, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %v", err))
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *AccessRulesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state AccessRules

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

	// Create URL path for the request
	var ruleNames strings.Builder
	var bulks []string
	var resAccessRules string = ""
	for i := 0; i < len(state.Items); i++ {
		if ruleNames.Len() > 0 {
			ruleNames.WriteString(",")
		}
		ruleNames.WriteString(url.QueryEscape(state.Items[i].Name.ValueString()))
		if ruleNames.Len() >= maxUrlParamLength || (i+1)%1000 == 0 {
			bulks = append(bulks, ruleNames.String())
			ruleNames.Reset()
		}
	}
	if ruleNames.Len() > 0 {
		bulks = append(bulks, ruleNames.String())
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))
	for _, bulk := range bulks {
		urlPath := state.getPath() + "?expanded=true&limit=1000&filter=name:" + bulk

		// Read Access Rules
		resTemp, err := r.client.Get(urlPath, reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, resAccessRules))
			return
		}
		if resAccessRules == "" {
			resAccessRules = resTemp.String()
		} else {
			items := gjson.Get(resTemp.String(), "items")
			if !items.Exists() {
				continue
			}
			if resItems := items.String()[1 : len(items.String())-1]; resItems != "" {
				resAccessRules, _ = sjson.SetRaw(resAccessRules, "items.-1", resItems)
			}
		}
	}

	var tmp string

	gjson.AddModifier("case", func(json, arg string) string {
		if arg == "lower" {
			return strings.ToLower(json)
		}
		return json
	})

	if !state.CategoryName.IsUnknown() && state.CategoryName.ValueString() != "" {
		var category string
		// Get the categories from the response
		res := gjson.Get(resAccessRules, "items.#.metadata.category")
		// Check if all categories are the same
		unique := make(map[string]struct{})
		for _, item := range res.Array() {
			unique[item.String()] = struct{}{}
		}
		if len(unique) > 1 {
			// If there are multiple categories, set a warning that would trigger resource recreation
			category = "[WARN] Inconsistent categories detected"
		} else {
			category = res.Array()[0].String()
		}
		// Assign the category to the temporary JSON string
		tmp, _ = sjson.Set(tmp, "category", category)
	} else if !state.Section.IsUnknown() && state.Section.ValueString() != "" {
		var section string
		// Get all sections from the response
		res := gjson.Get(resAccessRules, "items.#.metadata.section|@case:lower")
		// Check if all sections are the same
		unique := make(map[string]struct{})
		for _, item := range res.Array() {
			unique[item.String()] = struct{}{}
		}
		if len(unique) > 1 {
			// If there are multiple sections, set a warning that would trigger resource recreation
			section = "[WARN] Inconsistent sections detected"
		} else {
			section = res.Array()[0].String()
		}
		// Assign the section to the temporary JSON string
		tmp, _ = sjson.Set(tmp, "section", section)
	}

	// Merge computed category/section into the response
	tmp, _ = sjson.SetRaw(tmp, "items", gjson.Get(resAccessRules, "items").String())
	res := gjson.Parse(tmp)

	tflog.Debug(ctx, fmt.Sprintf("res: %s", res.String()))

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

func (r *AccessRulesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state AccessRules

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

	err := r.truncateRulesAt(ctx, &state, 0, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %v", err))
		return
	}

	err = r.createRulesAt(ctx, plan, &plan, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %v", err))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *AccessRulesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state AccessRules

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
	err := r.truncateRulesAt(ctx, &state, 0, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %v", err))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// Section below is generated&owned by "gen/generator.go". //template:begin import
// End of section. //template:end import

func (r *AccessRulesResource) createRulesAt(ctx context.Context, plan AccessRules, state *AccessRules, reqMods ...func(*fmc.Req)) error {
	var idx = 0
	bulk := plan
	bulk.Items = make([]AccessRulesItems, 0, bulkSizeCreate)
	state.Items = state.Items[:0] // reset state items

	// Mutex ensures that all rules, even if split into multiple bulks, are not mixed with
	// other rules being created at the same time.
	accessRulesCreateMu.Lock()
	defer accessRulesCreateMu.Unlock()

	// iterate over all items
	for _, v := range plan.Items {
		// count loops
		idx++

		// add object to current bulk
		bulk.Items = append(bulk.Items, v)

		// If bulk size was reached or all entries have been processed
		if idx%bulkSizeCreate == 0 || idx == len(plan.Items) {

			// Create object
			body := bulk.toBody(ctx, AccessRules{})
			body = gjson.Get(body, "items").String()

			// Execute request
			urlParams := "?bulk=true"
			if c := plan.CategoryName.ValueString(); c != "" {
				urlParams += "&category=" + url.QueryEscape(c)
			} else if s := plan.Section.ValueString(); s != "" {
				urlParams += "&section=" + url.QueryEscape(s)
			}

			res, err := r.client.Post(plan.getPath()+urlParams, body, reqMods...)
			if err != nil {
				return err
			}

			// Read result and save it to state
			bulk.fromBodyUnknowns(ctx, res)
			state.Items = append(state.Items, bulk.Items...)
			bulk.Items = bulk.Items[:0] // reset bulk items
		}
	}
	return nil
}

func (r *AccessRulesResource) truncateRulesAt(ctx context.Context, state *AccessRules, kept int, reqMods ...func(*fmc.Req)) error {
	var b strings.Builder
	var bulks []string
	var counts []int
	count := 0

	for i := kept; i < len(state.Items); i++ {
		if b.Len() != 0 {
			// URL encoded comma
			b.WriteString(",")
		}
		b.WriteString(state.Items[i].Id.ValueString())
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
		res, err := r.client.Delete(state.getPath()+"?bulk=true&filter=ids:"+bulk, reqMods...)
		if err != nil {
			return fmt.Errorf("failed to bulk-delete rules, got error: %v, %s", err, res.String())
		}

		state.Items = slices.Delete(state.Items, kept, kept+counts[i])
	}

	return nil
}
