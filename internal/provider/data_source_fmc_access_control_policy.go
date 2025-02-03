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

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &AccessControlPolicyDataSource{}
	_ datasource.DataSourceWithConfigure = &AccessControlPolicyDataSource{}
)

func NewAccessControlPolicyDataSource() datasource.DataSource {
	return &AccessControlPolicyDataSource{}
}

type AccessControlPolicyDataSource struct {
	client *fmc.Client
}

func (d *AccessControlPolicyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_access_control_policy"
}

func (d *AccessControlPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Access Control Policy.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Optional:            true,
				Computed:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "The name of the FMC domain",
				Optional:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the access control policy.",
				Optional:            true,
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description",
				Computed:            true,
			},
			"default_action": schema.StringAttribute{
				MarkdownDescription: "Specifies the default action to take when none of the rules meet the conditions.",
				Computed:            true,
			},
			"default_action_id": schema.StringAttribute{
				MarkdownDescription: "Default action ID.",
				Computed:            true,
			},
			"default_action_log_begin": schema.BoolAttribute{
				MarkdownDescription: "Indicating whether the device will log events at the beginning of the connection.",
				Computed:            true,
			},
			"default_action_log_end": schema.BoolAttribute{
				MarkdownDescription: "Indicating whether the device will log events at the end of the connection.",
				Computed:            true,
			},
			"default_action_send_events_to_fmc": schema.BoolAttribute{
				MarkdownDescription: "Indicating whether the device will send events to the Firepower Management Center event viewer.",
				Computed:            true,
			},
			"default_action_send_syslog": schema.BoolAttribute{
				MarkdownDescription: "Indicating whether the device will send events to a syslog server.",
				Computed:            true,
			},
			"default_action_syslog_config_id": schema.StringAttribute{
				MarkdownDescription: "UUID of the syslog config. Can be set only when default_action_send_syslog is true and either default_action_log_begin or default_action_log_end is true. If not set, the default policy syslog configuration in Access Control Logging applies.",
				Computed:            true,
			},
			"default_action_syslog_severity": schema.StringAttribute{
				MarkdownDescription: "Override the Severity of syslog alerts.",
				Computed:            true,
			},
			"default_action_snmp_config_id": schema.StringAttribute{
				MarkdownDescription: "UUID of the SNMP alert. Can be set only when either default_action_log_begin or default_action_log_end is true.",
				Computed:            true,
			},
			"default_action_intrusion_policy_id": schema.StringAttribute{
				MarkdownDescription: "UUID of the existing intrusion policy (e.g. fmc_intrusion_policy.example.id). Cannot be set when default action is BLOCK, TRUST, NETWORK_DISCOVERY.",
				Computed:            true,
			},
			"categories": schema.ListNestedAttribute{
				MarkdownDescription: "The ordered list of categories.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Identifier of the category.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "User-specified unique string.",
							Computed:            true,
						},
						"section": schema.StringAttribute{
							MarkdownDescription: "The section of the policy to which the category belongs. Categories must be ordered so that entire section 'mandatory' comes above the section 'default'. If you use inheritance, the mandatory section applies before child policy's own rules, while the default section applies after child policy's own rules.",
							Computed:            true,
						},
					},
				},
			},
			"rules": schema.ListNestedAttribute{
				MarkdownDescription: "The ordered list of rules. Rules must be sorted in the order of the corresponding categories, if they have `category_name`. Uncategorized non-mandatory rules must be below all other rules. The first matching rule is selected. Except for MONITOR rules, the system does not continue to evaluate traffic against additional rules after that traffic matches a rule.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Unique identifier (UUID) of the access rule.",
							Computed:            true,
						},
						"action": schema.StringAttribute{
							MarkdownDescription: "What to do when the conditions defined by the rule are met.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "User-specified unique string.",
							Computed:            true,
						},
						"category_name": schema.StringAttribute{
							MarkdownDescription: "Name of the category that owns this rule (a `name` from `categories` list).",
							Computed:            true,
						},
						"section": schema.StringAttribute{
							MarkdownDescription: "The section of the policy to which the rule belongs. Can only be used when the `category_name` is null. Rules must be ordered so that entire section 'mandatory' comes above the section 'default'. Null value means 'default'. If you use inheritance, the mandatory section applies before child policy's own rules, while the default section applies after child policy's own rules.",
							Computed:            true,
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: "Indicates whether the access rule is in effect (true) or not (false). Default is true.",
							Computed:            true,
						},
						"source_network_literals": schema.SetNestedAttribute{
							MarkdownDescription: "Set of objects that represent sources of traffic (literally specified).",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"value": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
								},
							},
						},
						"destination_network_literals": schema.SetNestedAttribute{
							MarkdownDescription: "Set of objects that represent destinations of traffic (literally specified).",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"value": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
								},
							},
						},
						"vlan_tag_literals": schema.SetNestedAttribute{
							MarkdownDescription: "Set of objects that represent vlan tags (literally specified).",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"start_tag": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"end_tag": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
								},
							},
						},
						"vlan_tag_objects": schema.SetNestedAttribute{
							MarkdownDescription: "Set of objects that represent vlan tags (fmc_vlan_tag, fmc_vlan_tag_group, ...).",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "UUID of the object (such as fmc_vlan_tag.example.id, etc.).",
										Computed:            true,
									},
								},
							},
						},
						"source_network_objects": schema.SetNestedAttribute{
							MarkdownDescription: "Set of objects that represent sources of traffic (fmc_network, fmc_host, ...).",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "UUID of the object (such as fmc_network.example.id, etc.).",
										Computed:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of the object (such as fmc_network.example.type, etc.).",
										Computed:            true,
									},
								},
							},
						},
						"destination_network_objects": schema.SetNestedAttribute{
							MarkdownDescription: "Set of objects that represent destinations of traffic (fmc_network, fmc_host, ...).",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "UUID of the object (such as fmc_network.example.id, etc.).",
										Computed:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of the object (such as fmc_network.example.type, etc.).",
										Computed:            true,
									},
								},
							},
						},
						"source_dynamic_objects": schema.SetNestedAttribute{
							MarkdownDescription: "Set of objects that represent dynamic sources of traffic (fmc_dynamic_object).",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "UUID of the object (such as fmc_dynamic_object.example.id, etc.).",
										Computed:            true,
									},
								},
							},
						},
						"destination_dynamic_objects": schema.SetNestedAttribute{
							MarkdownDescription: "Set of objects that represent dynamic destinations of traffic (fmc_dynamic_object).",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "UUID of the object (such as fmc_dynamic_object.example.id, etc.).",
										Computed:            true,
									},
								},
							},
						},
						"source_port_literals": schema.SetNestedAttribute{
							MarkdownDescription: "Set of objects that represent protocol/port (literally specified).",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"port": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"protocol": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"icmp_type": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"icmp_code": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
								},
							},
						},
						"source_port_objects": schema.SetNestedAttribute{
							MarkdownDescription: "Set of objects representing source ports associated with the rule (fmc_port or fmc_port_group).",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "UUID of the object (such as fmc_port.example.id, fmc_port_group.example.id, ...).",
										Computed:            true,
									},
								},
							},
						},
						"destination_port_literals": schema.SetNestedAttribute{
							MarkdownDescription: "Set of objects that represent protocol/port (literally specified).",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"port": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"protocol": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"icmp_type": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"icmp_code": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
								},
							},
						},
						"destination_port_objects": schema.SetNestedAttribute{
							MarkdownDescription: "Set of objects representing destination ports associated with the rule (fmc_port or fmc_port_group).",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "UUID of the object (such as fmc_port.example.id, fmc_port_group.example.id, ...).",
										Computed:            true,
									},
								},
							},
						},
						"source_sgt_objects": schema.SetNestedAttribute{
							MarkdownDescription: "Set of objects representing the source Security Group Tags (fmc_sgt - part of the dynamic attributes).",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "UUID of the object (such as fmc_sgt.example.id, etc.).",
										Computed:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of the object (such as fmc_security_group_tag.example.type, etc.).",
										Computed:            true,
									},
								},
							},
						},
						"source_zones": schema.SetNestedAttribute{
							MarkdownDescription: "Set of objects representing source security zones associated with the access rule (fmc_security_zone).",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "UUID of the object (such as fmc_security_zone.example.id, etc.).",
										Computed:            true,
									},
								},
							},
						},
						"destination_zones": schema.SetNestedAttribute{
							MarkdownDescription: "Set of objects representing destination security zones associated with the access rule (fmc_security_zone).",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "UUID of the object (such as fmc_security_zone.example.id, etc.).",
										Computed:            true,
									},
								},
							},
						},
						"url_literals": schema.SetNestedAttribute{
							MarkdownDescription: "Set of objects representing the URLs associated with the rule (literally specified).",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"url": schema.StringAttribute{
										MarkdownDescription: "URL such as https://www.example.com/app",
										Computed:            true,
									},
								},
							},
						},
						"url_objects": schema.SetNestedAttribute{
							MarkdownDescription: "Set of objects representing the URLs associated with the rule (fmc_url or fmc_url_group).",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "UUID of the object (such as fmc_url.example.id, fmc_url_group.id, etc.).",
										Computed:            true,
									},
								},
							},
						},
						"url_categories": schema.SetNestedAttribute{
							MarkdownDescription: "Set of objects representing the URL Categories associated with the rule (fmc_url_category).",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "UUID of the object (such as fmc_url_category.example.id, etc.).",
										Computed:            true,
									},
									"reputation": schema.StringAttribute{
										MarkdownDescription: "Reputation applicable to the category.",
										Computed:            true,
									},
								},
							},
						},
						"log_begin": schema.BoolAttribute{
							MarkdownDescription: "Indicates whether the device will log events at the beginning of the connection. If 'MONITOR' action is selected for access rule, log_begin must be false or absent.",
							Computed:            true,
						},
						"log_end": schema.BoolAttribute{
							MarkdownDescription: "Indicates whether the device will log events at the end of the connection. If 'MONITOR' action is selected for access rule, log_end must be true.",
							Computed:            true,
						},
						"log_files": schema.BoolAttribute{
							MarkdownDescription: "Indicates whether the device will log file events.",
							Computed:            true,
						},
						"send_events_to_fmc": schema.BoolAttribute{
							MarkdownDescription: "Indicates whether the device will send events to the Firepower Management Center event viewer. If 'MONITOR' action is selected for access rule, send_events_to_fmc must be true.",
							Computed:            true,
						},
						"send_syslog": schema.BoolAttribute{
							MarkdownDescription: "Indicates whether the alerts associated with the access rule are sent to syslog.",
							Computed:            true,
						},
						"syslog_config_id": schema.StringAttribute{
							MarkdownDescription: "UUID of the syslog config. Can be set only when send_syslog is true and either log_begin or log_end is true. If not set, the default policy syslog configuration in Access Control Logging applies.",
							Computed:            true,
						},
						"syslog_severity": schema.StringAttribute{
							MarkdownDescription: "Override the Severity of syslog alerts.",
							Computed:            true,
						},
						"snmp_config_id": schema.StringAttribute{
							MarkdownDescription: "UUID of the SNMP alert associated with the access rule. Can be set only when either log_begin or log_end is true.",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "User-specified string.",
							Computed:            true,
						},
						"file_policy_id": schema.StringAttribute{
							MarkdownDescription: "Identifier (UUID) of the File Policy for the rule action. Cannot be set when action is BLOCK, BLOCK_RESET, TRUST, MONITOR.",
							Computed:            true,
						},
						"intrusion_policy_id": schema.StringAttribute{
							MarkdownDescription: "Identifier (UUID) of the fmc_intrusion_policy for the rule action. Cannot be set when action is BLOCK, BLOCK_RESET, TRUST, MONITOR.",
							Computed:            true,
						},
						"variable_set_id": schema.StringAttribute{
							MarkdownDescription: "Identifier (UUID) of the Variable Set for the rule action.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}
func (d *AccessControlPolicyDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *AccessControlPolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

func (d *AccessControlPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config AccessControlPolicy

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
	if config.Id.IsNull() && !config.Name.IsNull() {
		offset := 0
		limit := 1000
		for page := 1; ; page++ {
			queryString := fmt.Sprintf("?limit=%d&offset=%d", limit, offset)
			res, err := d.client.Get(config.getPath()+queryString, reqMods...)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects, got error: %s", err))
				return
			}
			if value := res.Get("items"); len(value.Array()) > 0 {
				value.ForEach(func(k, v gjson.Result) bool {
					if config.Name.ValueString() == v.Get("name").String() {
						config.Id = types.StringValue(v.Get("id").String())
						tflog.Debug(ctx, fmt.Sprintf("%s: Found object with name '%v', id: %v", config.Id.String(), config.Name.ValueString(), config.Id.String()))
						return false
					}
					return true
				})
			}
			if !config.Id.IsNull() || !res.Get("paging.next.0").Exists() {
				break
			}
			offset += limit
		}

		if config.Id.IsNull() {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with name: %s", config.Name.ValueString()))
			return
		}
	}

	res, err := d.client.Get(config.getPath()+"/"+url.QueryEscape(config.Id.ValueString()), reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	resCats, err := d.client.Get(config.getPath()+"/"+url.QueryEscape(config.Id.ValueString())+"/categories?expanded=true&offset=0&limit=1000", reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	resRules, err := d.client.Get(config.getPath()+"/"+url.QueryEscape(config.Id.ValueString())+"/accessrules?expanded=true&offset=0&limit=1000", reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	replaceCats := resCats.Get("items").String()
	if replaceCats == "" {
		replaceCats = "[]"
	}
	replaceRules := resRules.Get("items").String()
	if replaceRules == "" {
		replaceRules = "[]"
	}
	replace, _ := sjson.SetRaw(res.String(), "dummy_categories", replaceCats)
	replace, _ = sjson.SetRaw(replace, "dummy_rules", replaceRules)
	res = gjson.Parse(replace)

	config.fromBody(ctx, res)
	config.adjustFromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
