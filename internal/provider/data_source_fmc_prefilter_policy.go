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
	_ datasource.DataSource              = &PrefilterPolicyDataSource{}
	_ datasource.DataSourceWithConfigure = &PrefilterPolicyDataSource{}
)

func NewPrefilterPolicyDataSource() datasource.DataSource {
	return &PrefilterPolicyDataSource{}
}

type PrefilterPolicyDataSource struct {
	client *fmc.Client
}

func (d *PrefilterPolicyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_prefilter_policy"
}

func (d *PrefilterPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the Prefilter Policy.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the object",
				Optional:            true,
				Computed:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "Name of the FMC domain",
				Optional:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the prefilter policy.",
				Optional:            true,
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the prefilter policy.",
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Object type; This is always `PrefilterPolicy`",
				Computed:            true,
			},
			"default_action": schema.StringAttribute{
				MarkdownDescription: "Specifies the default action to take when none of the rules meet the conditions.",
				Computed:            true,
			},
			"default_action_id": schema.StringAttribute{
				MarkdownDescription: "Default action ID",
				Computed:            true,
			},
			"default_action_log_begin": schema.BoolAttribute{
				MarkdownDescription: "Log events at the beginning of the connection.",
				Computed:            true,
			},
			"default_action_log_end": schema.BoolAttribute{
				MarkdownDescription: "Log events at the end of the connection.",
				Computed:            true,
			},
			"default_action_send_events_to_fmc": schema.BoolAttribute{
				MarkdownDescription: "Send events to the Firepower Management Center event viewer.",
				Computed:            true,
			},
			"default_action_syslog_config_id": schema.StringAttribute{
				MarkdownDescription: "UUID of the syslog config. Can be set only when either default_action_log_begin or default_action_log_end is true.",
				Computed:            true,
			},
			"default_action_snmp_config_id": schema.StringAttribute{
				MarkdownDescription: "UUID of the SNMP alert. Can be set only when either default_action_log_begin or default_action_log_end is true.",
				Computed:            true,
			},
			"rules": schema.ListNestedAttribute{
				MarkdownDescription: "The ordered list of rules.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Id of the prefilter rule.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Name of the rule.",
							Computed:            true,
						},
						"rule_type": schema.StringAttribute{
							MarkdownDescription: "Type of the rule. At least one Encapsulation Port Object (encapsulation_ports) is mandatory to be specified for TUNNEL Rules.",
							Computed:            true,
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: "Indicates whether the prefilter rule is in effect (true) or not (false). Default is true.",
							Computed:            true,
						},
						"action": schema.StringAttribute{
							MarkdownDescription: "What to do when the conditions defined by the rule are met.",
							Computed:            true,
						},
						"bidirectional": schema.BoolAttribute{
							MarkdownDescription: "Indicates whether the rule is bidirectional. Can be true only for TUNNEL rules. Default is false.",
							Computed:            true,
						},
						"tunnel_zone_id": schema.StringAttribute{
							MarkdownDescription: "Id of Tunnel Zone. Can be only set for TUNNEL rules with ANALYZE action.",
							Computed:            true,
						},
						"time_range_id": schema.StringAttribute{
							MarkdownDescription: "Id of Time Range object applied to the rule.",
							Computed:            true,
						},
						"source_interfaces": schema.SetNestedAttribute{
							MarkdownDescription: "Set of objects that represent source interfaces.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the object.",
										Computed:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of the object.",
										Computed:            true,
									},
								},
							},
						},
						"destination_interfaces": schema.SetNestedAttribute{
							MarkdownDescription: "Set of objects that represent destination interfaces.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the object.",
										Computed:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of the object.",
										Computed:            true,
									},
								},
							},
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
						"source_network_objects": schema.SetNestedAttribute{
							MarkdownDescription: "Set of objects that represent sources of traffic (Networks, Hosts, Ranges or Network Groups).",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the object",
										Computed:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of the object",
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
						"destination_network_objects": schema.SetNestedAttribute{
							MarkdownDescription: "Set of objects that represent destinations of traffic (Networks, Hosts, Ranges or Network Groups).",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the object.",
										Computed:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of the object",
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
							MarkdownDescription: "Set of objects representing vlan tags or vlan tag groups",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the object.",
										Computed:            true,
									},
								},
							},
						},
						"source_port_literals": schema.SetNestedAttribute{
							MarkdownDescription: "Set of objects that represent protocol/port (literally specified). Can be only set for PREFILTER rules.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"protocol": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"port": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
								},
							},
						},
						"source_port_objects": schema.SetNestedAttribute{
							MarkdownDescription: "Set of objects representing source ports associated with the rule (Port or Port Groups). Can be only set for PREFILTER rules.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the object.",
										Computed:            true,
									},
								},
							},
						},
						"destination_port_literals": schema.SetNestedAttribute{
							MarkdownDescription: "Set of objects that represent protocol/port (literally specified). Can be only set for PREFILTER rules.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"protocol": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"port": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
								},
							},
						},
						"destination_port_objects": schema.SetNestedAttribute{
							MarkdownDescription: "Set of objects representing destination ports associated with the rule (Port or Port Groups). Can be only set for PREFILTER rules.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the object.",
										Computed:            true,
									},
								},
							},
						},
						"encapsulation_ports": schema.SetAttribute{
							MarkdownDescription: "List of encapsulation ports to be used.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"log_begin": schema.BoolAttribute{
							MarkdownDescription: "Log events at the beginning of the connection. Default is false.",
							Computed:            true,
						},
						"log_end": schema.BoolAttribute{
							MarkdownDescription: "Log events at the end of the connection. Default is false.",
							Computed:            true,
						},
						"send_events_to_fmc": schema.BoolAttribute{
							MarkdownDescription: "Send events to the Firepower Management Center event viewer. Default is false.",
							Computed:            true,
						},
						"send_syslog": schema.BoolAttribute{
							MarkdownDescription: "Send alerts associated with the prefilter rule to default syslog configuration in Prefilter Logging. Default is false.",
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
							MarkdownDescription: "UUID of the SNMP alert associated with the prefilter rule. Can be set only when either log_begin or log_end is true.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}
func (d *PrefilterPolicyDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *PrefilterPolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

func (d *PrefilterPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config PrefilterPolicy

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
	urlPath := config.getPath() + "/" + url.QueryEscape(config.Id.ValueString())
	resGet, err := d.client.Get(urlPath, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	// Default action properties (like logging settings) are not part of the main object
	resDefaultAction, err := d.client.Get(config.getPath()+"/"+url.QueryEscape(config.Id.ValueString())+"/defaultactions/"+config.DefaultActionId.String(), reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve default action object (GET), got error: %s, %s", err, resDefaultAction.String()))
		return
	}

	resRules, err := d.client.Get(config.getPath()+"/"+url.QueryEscape(config.Id.ValueString())+"/prefilterrules?expanded=true", reqMods...)
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

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
