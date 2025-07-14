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
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &AccessRuleDataSource{}
	_ datasource.DataSourceWithConfigure = &AccessRuleDataSource{}
)

func NewAccessRuleDataSource() datasource.DataSource {
	return &AccessRuleDataSource{}
}

type AccessRuleDataSource struct {
	client *fmc.Client
}

func (d *AccessRuleDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_access_rule"
}

func (d *AccessRuleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the Access Rule.").String,

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
			"access_policy_id": schema.StringAttribute{
				MarkdownDescription: "Id of the Access Control Policy.",
				Required:            true,
			},
			"category_name": schema.StringAttribute{
				MarkdownDescription: "Name of the category that owns this rule. Either 'section' or 'category_name' can be set.",
				Computed:            true,
			},
			"section": schema.StringAttribute{
				MarkdownDescription: "The section of the policy to which the rule belongs. Either 'section' or 'category_name' can be set.",
				Computed:            true,
			},
			"action": schema.StringAttribute{
				MarkdownDescription: "Rule action.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the Access Rule. This name needs to be uqique within the policy.",
				Optional:            true,
				Computed:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Indicates whether the access rule is in effect (true) or not (false).",
				Computed:            true,
			},
			"source_network_literals": schema.SetNestedAttribute{
				MarkdownDescription: "Set of objects that represent sources of traffic (literally specified).",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"value": schema.StringAttribute{
							MarkdownDescription: "IP address or network in CIDR format.",
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
							MarkdownDescription: "IP address or network in CIDR format.",
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
							MarkdownDescription: "Start of the VLAN tag range.",
							Computed:            true,
						},
						"end_tag": schema.StringAttribute{
							MarkdownDescription: "End of the VLAN tag range.",
							Computed:            true,
						},
					},
				},
			},
			"vlan_tag_objects": schema.SetNestedAttribute{
				MarkdownDescription: "Set of objects that represent vlan tags or vlan tags group.",
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
			"source_network_objects": schema.SetNestedAttribute{
				MarkdownDescription: "Set of objects that represent sources of traffic (Host, Network, Range, FQDN or Network Group).",
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
			"destination_network_objects": schema.SetNestedAttribute{
				MarkdownDescription: "Set of objects that represent destinations of traffic (Host, Network, Range, FQDN or Network Group).",
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
			"source_dynamic_objects": schema.SetNestedAttribute{
				MarkdownDescription: "Set of objects that represent dynamic sources of traffic.",
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
			"destination_dynamic_objects": schema.SetNestedAttribute{
				MarkdownDescription: "Set of objects that represent dynamic destinations of traffic.",
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
				MarkdownDescription: "Set of objects that represent protocol/port (literally specified).",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Type of the object.",
							Computed:            true,
						},
						"port": schema.StringAttribute{
							MarkdownDescription: "Port number.",
							Computed:            true,
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: "IANA protocol number.",
							Computed:            true,
						},
						"icmp_type": schema.StringAttribute{
							MarkdownDescription: "ICMP type.",
							Computed:            true,
						},
						"icmp_code": schema.StringAttribute{
							MarkdownDescription: "ICMP code.",
							Computed:            true,
						},
					},
				},
			},
			"source_port_objects": schema.SetNestedAttribute{
				MarkdownDescription: "Set of objects representing source ports associated with the rule.",
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
				MarkdownDescription: "Set of objects that represent protocol/port (literally specified).",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Type of the object.",
							Computed:            true,
						},
						"port": schema.StringAttribute{
							MarkdownDescription: "Port number.",
							Computed:            true,
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: "IANA protocol number.",
							Computed:            true,
						},
						"icmp_type": schema.StringAttribute{
							MarkdownDescription: "ICMP type.",
							Computed:            true,
						},
						"icmp_code": schema.StringAttribute{
							MarkdownDescription: "ICMP code.",
							Computed:            true,
						},
					},
				},
			},
			"destination_port_objects": schema.SetNestedAttribute{
				MarkdownDescription: "Set of objects representing destination ports associated with the rule.",
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
			"source_sgt_objects": schema.SetNestedAttribute{
				MarkdownDescription: "Set of objects representing the source Security Group Tags (SGT) or ISE Security Group Tags.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Name of the object.",
							Computed:            true,
						},
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
			"endpoint_device_types": schema.SetNestedAttribute{
				MarkdownDescription: "Set of objects representing the source Endpoint Device Types.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Name of the object.",
							Computed:            true,
						},
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
			"destination_sgt_objects": schema.SetNestedAttribute{
				MarkdownDescription: "Set of objects representing the destination ISE Security Group Tags (SGT).",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Name of the object.",
							Computed:            true,
						},
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
			"source_zones": schema.SetNestedAttribute{
				MarkdownDescription: "Set of objects representing source Security Zones associated with the access rule.",
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
			"destination_zones": schema.SetNestedAttribute{
				MarkdownDescription: "Set of objects representing destination Security Zones associated with the access rule.",
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
				MarkdownDescription: "Set of objects representing the URLs associated with the rule.",
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
			"url_categories": schema.SetNestedAttribute{
				MarkdownDescription: "Set of objects representing the URL Categories associated with the rule.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Id of the object.",
							Computed:            true,
						},
						"reputation": schema.StringAttribute{
							MarkdownDescription: "Reputation applicable to the URL Category.",
							Computed:            true,
						},
					},
				},
			},
			"log_begin": schema.BoolAttribute{
				MarkdownDescription: "Log events at the beginning of the connection. If 'MONITOR' action is selected for access rule, log_begin must be false or absent.",
				Computed:            true,
			},
			"log_end": schema.BoolAttribute{
				MarkdownDescription: "Log events at the end of the connection. If 'MONITOR' action is selected for access rule, log_end must be true.",
				Computed:            true,
			},
			"log_files": schema.BoolAttribute{
				MarkdownDescription: "Log file events.",
				Computed:            true,
			},
			"send_events_to_fmc": schema.BoolAttribute{
				MarkdownDescription: "Send events to the Firepower Management Center event viewer. If 'MONITOR' action is selected for access rule, send_events_to_fmc must be true.",
				Computed:            true,
			},
			"send_syslog": schema.BoolAttribute{
				MarkdownDescription: "Send alerts to syslog.",
				Computed:            true,
			},
			"syslog_config_id": schema.StringAttribute{
				MarkdownDescription: "Id of Syslog Config. Can be set only when send_syslog is true and either log_begin or log_end is true. If not set, the default syslog configuration in Access Control Policy Logging applies.",
				Computed:            true,
			},
			"syslog_severity": schema.StringAttribute{
				MarkdownDescription: "Override the Severity of syslog alerts.",
				Computed:            true,
			},
			"snmp_config_id": schema.StringAttribute{
				MarkdownDescription: "Id of the SNMP alert associated with the access rule. Can be set only when either log_begin or log_end is true.",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Rule description.",
				Computed:            true,
			},
			"file_policy_id": schema.StringAttribute{
				MarkdownDescription: "Id of the File Policy for the rule action. Cannot be set when action is BLOCK, BLOCK_RESET, TRUST, MONITOR.",
				Computed:            true,
			},
			"intrusion_policy_id": schema.StringAttribute{
				MarkdownDescription: "Id of the Intrusion Policy for the rule action. Cannot be set when action is BLOCK, BLOCK_RESET, TRUST, MONITOR.",
				Computed:            true,
			},
			"time_range_id": schema.StringAttribute{
				MarkdownDescription: "Id of Time Range object applied to the rule.",
				Computed:            true,
			},
			"variable_set_id": schema.StringAttribute{
				MarkdownDescription: "Id of the Variable Set for the rule action.",
				Computed:            true,
			},
			"applications": schema.SetNestedAttribute{
				MarkdownDescription: "Set of applications.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Id of the Application.",
							Computed:            true,
						},
					},
				},
			},
			"application_filter_objects": schema.SetNestedAttribute{
				MarkdownDescription: "Set of Application Filtering objects.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Id of the Application Filter.",
							Computed:            true,
						},
					},
				},
			},
			"application_filters": schema.ListNestedAttribute{
				MarkdownDescription: "List of Application Filtering conditions.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"types": schema.SetNestedAttribute{
							MarkdownDescription: "Set of Application Types.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the Application Type.",
										Computed:            true,
									},
								},
							},
						},
						"risks": schema.SetNestedAttribute{
							MarkdownDescription: "Set of Application Risk levels.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the Application Risk level.",
										Computed:            true,
									},
								},
							},
						},
						"business_relevances": schema.SetNestedAttribute{
							MarkdownDescription: "Set of Application Business Relevance levels.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the Application Business Relevance level.",
										Computed:            true,
									},
								},
							},
						},
						"categories": schema.SetNestedAttribute{
							MarkdownDescription: "Set of Application Categories.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the Application Category.",
										Computed:            true,
									},
								},
							},
						},
						"tags": schema.SetNestedAttribute{
							MarkdownDescription: "Set of Application Tags.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Id of the Application Tag.",
										Computed:            true,
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
func (d *AccessRuleDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *AccessRuleDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *AccessRuleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config AccessRule

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
			queryString := fmt.Sprintf("?limit=%d&offset=%d&expanded=true", limit, offset)
			res, err := d.client.Get(config.getPath()+queryString, reqMods...)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects, got error: %s", err))
				return
			}
			if value := res.Get("items"); len(value.Array()) > 0 {
				value.ForEach(func(k, v gjson.Result) bool {
					if config.Name.ValueString() == v.Get("name").String() {
						config.Id = types.StringValue(v.Get("id").String())
						tflog.Debug(ctx, fmt.Sprintf("%s: Found object with name '%v', id: %v", config.Id.ValueString(), config.Name.ValueString(), config.Id.ValueString()))
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
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with name: %v", config.Name.ValueString()))
			return
		}
	}
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
