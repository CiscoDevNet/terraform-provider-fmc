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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type AccessRule struct {
	Id                         types.String                           `tfsdk:"id"`
	Domain                     types.String                           `tfsdk:"domain"`
	AccessControlPolicyId      types.String                           `tfsdk:"access_control_policy_id"`
	CategoryName               types.String                           `tfsdk:"category_name"`
	Section                    types.String                           `tfsdk:"section"`
	Action                     types.String                           `tfsdk:"action"`
	Name                       types.String                           `tfsdk:"name"`
	Enabled                    types.Bool                             `tfsdk:"enabled"`
	SourceNetworkLiterals      []AccessRuleSourceNetworkLiterals      `tfsdk:"source_network_literals"`
	DestinationNetworkLiterals []AccessRuleDestinationNetworkLiterals `tfsdk:"destination_network_literals"`
	VlanTagLiterals            []AccessRuleVlanTagLiterals            `tfsdk:"vlan_tag_literals"`
	VlanTagObjects             []AccessRuleVlanTagObjects             `tfsdk:"vlan_tag_objects"`
	SourceNetworkObjects       []AccessRuleSourceNetworkObjects       `tfsdk:"source_network_objects"`
	DestinationNetworkObjects  []AccessRuleDestinationNetworkObjects  `tfsdk:"destination_network_objects"`
	SourceDynamicObjects       []AccessRuleSourceDynamicObjects       `tfsdk:"source_dynamic_objects"`
	DestinationDynamicObjects  []AccessRuleDestinationDynamicObjects  `tfsdk:"destination_dynamic_objects"`
	SourcePortLiterals         []AccessRuleSourcePortLiterals         `tfsdk:"source_port_literals"`
	SourcePortObjects          []AccessRuleSourcePortObjects          `tfsdk:"source_port_objects"`
	DestinationPortLiterals    []AccessRuleDestinationPortLiterals    `tfsdk:"destination_port_literals"`
	DestinationPortObjects     []AccessRuleDestinationPortObjects     `tfsdk:"destination_port_objects"`
	SourceSgtObjects           []AccessRuleSourceSgtObjects           `tfsdk:"source_sgt_objects"`
	EndpointDeviceTypes        []AccessRuleEndpointDeviceTypes        `tfsdk:"endpoint_device_types"`
	DestinationSgtObjects      []AccessRuleDestinationSgtObjects      `tfsdk:"destination_sgt_objects"`
	SourceZones                []AccessRuleSourceZones                `tfsdk:"source_zones"`
	DestinationZones           []AccessRuleDestinationZones           `tfsdk:"destination_zones"`
	UrlLiterals                []AccessRuleUrlLiterals                `tfsdk:"url_literals"`
	UrlObjects                 []AccessRuleUrlObjects                 `tfsdk:"url_objects"`
	UrlCategories              []AccessRuleUrlCategories              `tfsdk:"url_categories"`
	LogBegin                   types.Bool                             `tfsdk:"log_begin"`
	LogEnd                     types.Bool                             `tfsdk:"log_end"`
	LogFiles                   types.Bool                             `tfsdk:"log_files"`
	SendEventsToFmc            types.Bool                             `tfsdk:"send_events_to_fmc"`
	SendSyslog                 types.Bool                             `tfsdk:"send_syslog"`
	SyslogConfigId             types.String                           `tfsdk:"syslog_config_id"`
	SyslogSeverity             types.String                           `tfsdk:"syslog_severity"`
	SnmpConfigId               types.String                           `tfsdk:"snmp_config_id"`
	Description                types.String                           `tfsdk:"description"`
	FilePolicyId               types.String                           `tfsdk:"file_policy_id"`
	IntrusionPolicyId          types.String                           `tfsdk:"intrusion_policy_id"`
	TimeRangeId                types.String                           `tfsdk:"time_range_id"`
	VariableSetId              types.String                           `tfsdk:"variable_set_id"`
	Applications               []AccessRuleApplications               `tfsdk:"applications"`
	ApplicationFilterObjects   []AccessRuleApplicationFilterObjects   `tfsdk:"application_filter_objects"`
	ApplicationFilters         []AccessRuleApplicationFilters         `tfsdk:"application_filters"`
}

type AccessRuleSourceNetworkLiterals struct {
	Value types.String `tfsdk:"value"`
}

type AccessRuleDestinationNetworkLiterals struct {
	Value types.String `tfsdk:"value"`
}

type AccessRuleVlanTagLiterals struct {
	StartTag types.String `tfsdk:"start_tag"`
	EndTag   types.String `tfsdk:"end_tag"`
}

type AccessRuleVlanTagObjects struct {
	Id types.String `tfsdk:"id"`
}

type AccessRuleSourceNetworkObjects struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}

type AccessRuleDestinationNetworkObjects struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}

type AccessRuleSourceDynamicObjects struct {
	Id types.String `tfsdk:"id"`
}

type AccessRuleDestinationDynamicObjects struct {
	Id types.String `tfsdk:"id"`
}

type AccessRuleSourcePortLiterals struct {
	Type     types.String `tfsdk:"type"`
	Port     types.String `tfsdk:"port"`
	Protocol types.String `tfsdk:"protocol"`
	IcmpType types.String `tfsdk:"icmp_type"`
	IcmpCode types.String `tfsdk:"icmp_code"`
}

type AccessRuleSourcePortObjects struct {
	Id types.String `tfsdk:"id"`
}

type AccessRuleDestinationPortLiterals struct {
	Type     types.String `tfsdk:"type"`
	Port     types.String `tfsdk:"port"`
	Protocol types.String `tfsdk:"protocol"`
	IcmpType types.String `tfsdk:"icmp_type"`
	IcmpCode types.String `tfsdk:"icmp_code"`
}

type AccessRuleDestinationPortObjects struct {
	Id types.String `tfsdk:"id"`
}

type AccessRuleSourceSgtObjects struct {
	Name types.String `tfsdk:"name"`
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}

type AccessRuleEndpointDeviceTypes struct {
	Name types.String `tfsdk:"name"`
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}

type AccessRuleDestinationSgtObjects struct {
	Name types.String `tfsdk:"name"`
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}

type AccessRuleSourceZones struct {
	Id types.String `tfsdk:"id"`
}

type AccessRuleDestinationZones struct {
	Id types.String `tfsdk:"id"`
}

type AccessRuleUrlLiterals struct {
	Url types.String `tfsdk:"url"`
}

type AccessRuleUrlObjects struct {
	Id types.String `tfsdk:"id"`
}

type AccessRuleUrlCategories struct {
	Id         types.String `tfsdk:"id"`
	Reputation types.String `tfsdk:"reputation"`
}

type AccessRuleApplications struct {
	Id types.String `tfsdk:"id"`
}

type AccessRuleApplicationFilterObjects struct {
	Id types.String `tfsdk:"id"`
}

type AccessRuleApplicationFilters struct {
	Types              []AccessRuleApplicationFiltersTypes              `tfsdk:"types"`
	Risks              []AccessRuleApplicationFiltersRisks              `tfsdk:"risks"`
	BusinessRelevances []AccessRuleApplicationFiltersBusinessRelevances `tfsdk:"business_relevances"`
	Categories         []AccessRuleApplicationFiltersCategories         `tfsdk:"categories"`
	Tags               []AccessRuleApplicationFiltersTags               `tfsdk:"tags"`
}

type AccessRuleApplicationFiltersTypes struct {
	Id types.String `tfsdk:"id"`
}
type AccessRuleApplicationFiltersRisks struct {
	Id types.String `tfsdk:"id"`
}
type AccessRuleApplicationFiltersBusinessRelevances struct {
	Id types.String `tfsdk:"id"`
}
type AccessRuleApplicationFiltersCategories struct {
	Id types.String `tfsdk:"id"`
}
type AccessRuleApplicationFiltersTags struct {
	Id types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data AccessRule) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/accesspolicies/%v/accessrules", url.QueryEscape(data.AccessControlPolicyId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data AccessRule) toBody(ctx context.Context, state AccessRule) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.CategoryName.IsNull() {
		body, _ = sjson.Set(body, "metadata.category", data.CategoryName.ValueString())
	}
	if !data.Section.IsNull() {
		body, _ = sjson.Set(body, "metadata.section", data.Section.ValueString())
	}
	if !data.Action.IsNull() {
		body, _ = sjson.Set(body, "action", data.Action.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "enabled", data.Enabled.ValueBool())
	}
	if len(data.SourceNetworkLiterals) > 0 {
		body, _ = sjson.Set(body, "sourceNetworks.literals", []any{})
		for _, item := range data.SourceNetworkLiterals {
			itemBody := ""
			itemBody, _ = sjson.Set(itemBody, "type", "AnyNonEmptyString")
			if !item.Value.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "value", item.Value.ValueString())
			}
			body, _ = sjson.SetRaw(body, "sourceNetworks.literals.-1", itemBody)
		}
	}
	if len(data.DestinationNetworkLiterals) > 0 {
		body, _ = sjson.Set(body, "destinationNetworks.literals", []any{})
		for _, item := range data.DestinationNetworkLiterals {
			itemBody := ""
			itemBody, _ = sjson.Set(itemBody, "type", "AnyNonEmptyString")
			if !item.Value.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "value", item.Value.ValueString())
			}
			body, _ = sjson.SetRaw(body, "destinationNetworks.literals.-1", itemBody)
		}
	}
	if len(data.VlanTagLiterals) > 0 {
		body, _ = sjson.Set(body, "vlanTags.literals", []any{})
		for _, item := range data.VlanTagLiterals {
			itemBody := ""
			itemBody, _ = sjson.Set(itemBody, "type", "AnyNonEmptyString")
			if !item.StartTag.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "startTag", item.StartTag.ValueString())
			}
			if !item.EndTag.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "endTag", item.EndTag.ValueString())
			}
			body, _ = sjson.SetRaw(body, "vlanTags.literals.-1", itemBody)
		}
	}
	if len(data.VlanTagObjects) > 0 {
		body, _ = sjson.Set(body, "vlanTags.objects", []any{})
		for _, item := range data.VlanTagObjects {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "vlanTags.objects.-1", itemBody)
		}
	}
	if len(data.SourceNetworkObjects) > 0 {
		body, _ = sjson.Set(body, "sourceNetworks.objects", []any{})
		for _, item := range data.SourceNetworkObjects {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
			}
			body, _ = sjson.SetRaw(body, "sourceNetworks.objects.-1", itemBody)
		}
	}
	if len(data.DestinationNetworkObjects) > 0 {
		body, _ = sjson.Set(body, "destinationNetworks.objects", []any{})
		for _, item := range data.DestinationNetworkObjects {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
			}
			body, _ = sjson.SetRaw(body, "destinationNetworks.objects.-1", itemBody)
		}
	}
	if len(data.SourceDynamicObjects) > 0 {
		body, _ = sjson.Set(body, "sourceDynamicObjects.objects", []any{})
		for _, item := range data.SourceDynamicObjects {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			itemBody, _ = sjson.Set(itemBody, "type", "DynamicObject")
			body, _ = sjson.SetRaw(body, "sourceDynamicObjects.objects.-1", itemBody)
		}
	}
	if len(data.DestinationDynamicObjects) > 0 {
		body, _ = sjson.Set(body, "destinationDynamicObjects.objects", []any{})
		for _, item := range data.DestinationDynamicObjects {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			itemBody, _ = sjson.Set(itemBody, "type", "DynamicObject")
			body, _ = sjson.SetRaw(body, "destinationDynamicObjects.objects.-1", itemBody)
		}
	}
	if len(data.SourcePortLiterals) > 0 {
		body, _ = sjson.Set(body, "sourcePorts.literals", []any{})
		for _, item := range data.SourcePortLiterals {
			itemBody := ""
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
			}
			if !item.Port.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "port", item.Port.ValueString())
			}
			if !item.Protocol.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "protocol", item.Protocol.ValueString())
			}
			if !item.IcmpType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "icmpType", item.IcmpType.ValueString())
			}
			if !item.IcmpCode.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "code", item.IcmpCode.ValueString())
			}
			body, _ = sjson.SetRaw(body, "sourcePorts.literals.-1", itemBody)
		}
	}
	if len(data.SourcePortObjects) > 0 {
		body, _ = sjson.Set(body, "sourcePorts.objects", []any{})
		for _, item := range data.SourcePortObjects {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			itemBody, _ = sjson.Set(itemBody, "type", "AnyNonEmptyString")
			body, _ = sjson.SetRaw(body, "sourcePorts.objects.-1", itemBody)
		}
	}
	if len(data.DestinationPortLiterals) > 0 {
		body, _ = sjson.Set(body, "destinationPorts.literals", []any{})
		for _, item := range data.DestinationPortLiterals {
			itemBody := ""
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
			}
			if !item.Port.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "port", item.Port.ValueString())
			}
			if !item.Protocol.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "protocol", item.Protocol.ValueString())
			}
			if !item.IcmpType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "icmpType", item.IcmpType.ValueString())
			}
			if !item.IcmpCode.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "code", item.IcmpCode.ValueString())
			}
			body, _ = sjson.SetRaw(body, "destinationPorts.literals.-1", itemBody)
		}
	}
	if len(data.DestinationPortObjects) > 0 {
		body, _ = sjson.Set(body, "destinationPorts.objects", []any{})
		for _, item := range data.DestinationPortObjects {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			itemBody, _ = sjson.Set(itemBody, "type", "AnyNonEmptyString")
			body, _ = sjson.SetRaw(body, "destinationPorts.objects.-1", itemBody)
		}
	}
	if len(data.SourceSgtObjects) > 0 {
		body, _ = sjson.Set(body, "sourceSecurityGroupTags.objects", []any{})
		for _, item := range data.SourceSgtObjects {
			itemBody := ""
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
			}
			body, _ = sjson.SetRaw(body, "sourceSecurityGroupTags.objects.-1", itemBody)
		}
	}
	if len(data.EndpointDeviceTypes) > 0 {
		body, _ = sjson.Set(body, "endPointDeviceTypes", []any{})
		for _, item := range data.EndpointDeviceTypes {
			itemBody := ""
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
			}
			body, _ = sjson.SetRaw(body, "endPointDeviceTypes.-1", itemBody)
		}
	}
	if len(data.DestinationSgtObjects) > 0 {
		body, _ = sjson.Set(body, "destinationSecurityGroupTags.objects", []any{})
		for _, item := range data.DestinationSgtObjects {
			itemBody := ""
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
			}
			body, _ = sjson.SetRaw(body, "destinationSecurityGroupTags.objects.-1", itemBody)
		}
	}
	if len(data.SourceZones) > 0 {
		body, _ = sjson.Set(body, "sourceZones.objects", []any{})
		for _, item := range data.SourceZones {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			itemBody, _ = sjson.Set(itemBody, "type", "SecurityZone")
			body, _ = sjson.SetRaw(body, "sourceZones.objects.-1", itemBody)
		}
	}
	if len(data.DestinationZones) > 0 {
		body, _ = sjson.Set(body, "destinationZones.objects", []any{})
		for _, item := range data.DestinationZones {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			itemBody, _ = sjson.Set(itemBody, "type", "SecurityZone")
			body, _ = sjson.SetRaw(body, "destinationZones.objects.-1", itemBody)
		}
	}
	if len(data.UrlLiterals) > 0 {
		body, _ = sjson.Set(body, "urls.literals", []any{})
		for _, item := range data.UrlLiterals {
			itemBody := ""
			if !item.Url.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "url", item.Url.ValueString())
			}
			body, _ = sjson.SetRaw(body, "urls.literals.-1", itemBody)
		}
	}
	if len(data.UrlObjects) > 0 {
		body, _ = sjson.Set(body, "urls.objects", []any{})
		for _, item := range data.UrlObjects {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			itemBody, _ = sjson.Set(itemBody, "type", "AnyNonEmptyString")
			body, _ = sjson.SetRaw(body, "urls.objects.-1", itemBody)
		}
	}
	if len(data.UrlCategories) > 0 {
		body, _ = sjson.Set(body, "urls.urlCategoriesWithReputation", []any{})
		for _, item := range data.UrlCategories {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "category.id", item.Id.ValueString())
			}
			itemBody, _ = sjson.Set(itemBody, "category.type", "AnyNonEmptyString")
			if !item.Reputation.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "reputation", item.Reputation.ValueString())
			}
			body, _ = sjson.SetRaw(body, "urls.urlCategoriesWithReputation.-1", itemBody)
		}
	}
	if !data.LogBegin.IsNull() {
		body, _ = sjson.Set(body, "logBegin", data.LogBegin.ValueBool())
	}
	if !data.LogEnd.IsNull() {
		body, _ = sjson.Set(body, "logEnd", data.LogEnd.ValueBool())
	}
	if !data.LogFiles.IsNull() {
		body, _ = sjson.Set(body, "logFiles", data.LogFiles.ValueBool())
	}
	if !data.SendEventsToFmc.IsNull() {
		body, _ = sjson.Set(body, "sendEventsToFMC", data.SendEventsToFmc.ValueBool())
	}
	if !data.SendSyslog.IsNull() {
		body, _ = sjson.Set(body, "enableSyslog", data.SendSyslog.ValueBool())
	}
	if !data.SyslogConfigId.IsNull() {
		body, _ = sjson.Set(body, "syslogConfig.id", data.SyslogConfigId.ValueString())
	}
	if !data.SyslogSeverity.IsNull() {
		body, _ = sjson.Set(body, "syslogSeverity", data.SyslogSeverity.ValueString())
	}
	if !data.SnmpConfigId.IsNull() {
		body, _ = sjson.Set(body, "snmpConfig.id", data.SnmpConfigId.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.FilePolicyId.IsNull() {
		body, _ = sjson.Set(body, "filePolicy.id", data.FilePolicyId.ValueString())
	}
	if !data.IntrusionPolicyId.IsNull() {
		body, _ = sjson.Set(body, "ipsPolicy.id", data.IntrusionPolicyId.ValueString())
	}
	if !data.TimeRangeId.IsNull() {
		body, _ = sjson.Set(body, "timeRangeObjects.0.id", data.TimeRangeId.ValueString())
	}
	if !data.VariableSetId.IsNull() {
		body, _ = sjson.Set(body, "variableSet.id", data.VariableSetId.ValueString())
	}
	if len(data.Applications) > 0 {
		body, _ = sjson.Set(body, "applications.applications", []any{})
		for _, item := range data.Applications {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "applications.applications.-1", itemBody)
		}
	}
	if len(data.ApplicationFilterObjects) > 0 {
		body, _ = sjson.Set(body, "applications.applicationFilters", []any{})
		for _, item := range data.ApplicationFilterObjects {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "applications.applicationFilters.-1", itemBody)
		}
	}
	if len(data.ApplicationFilters) > 0 {
		body, _ = sjson.Set(body, "applications.inlineApplicationFilters", []any{})
		for _, item := range data.ApplicationFilters {
			itemBody := ""
			if len(item.Types) > 0 {
				itemBody, _ = sjson.Set(itemBody, "applicationTypes", []any{})
				for _, childItem := range item.Types {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "applicationTypes.-1", itemChildBody)
				}
			}
			if len(item.Risks) > 0 {
				itemBody, _ = sjson.Set(itemBody, "risks", []any{})
				for _, childItem := range item.Risks {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "risks.-1", itemChildBody)
				}
			}
			if len(item.BusinessRelevances) > 0 {
				itemBody, _ = sjson.Set(itemBody, "productivities", []any{})
				for _, childItem := range item.BusinessRelevances {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "productivities.-1", itemChildBody)
				}
			}
			if len(item.Categories) > 0 {
				itemBody, _ = sjson.Set(itemBody, "categories", []any{})
				for _, childItem := range item.Categories {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "categories.-1", itemChildBody)
				}
			}
			if len(item.Tags) > 0 {
				itemBody, _ = sjson.Set(itemBody, "tags", []any{})
				for _, childItem := range item.Tags {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "tags.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "applications.inlineApplicationFilters.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *AccessRule) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("metadata.category"); value.Exists() {
		data.CategoryName = types.StringValue(value.String())
	} else {
		data.CategoryName = types.StringNull()
	}
	if value := res.Get("metadata.section"); value.Exists() {
		data.Section = types.StringValue(value.String())
	} else {
		data.Section = types.StringNull()
	}
	if value := res.Get("action"); value.Exists() {
		data.Action = types.StringValue(value.String())
	} else {
		data.Action = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("enabled"); value.Exists() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolValue(true)
	}
	if value := res.Get("sourceNetworks.literals"); value.Exists() {
		data.SourceNetworkLiterals = make([]AccessRuleSourceNetworkLiterals, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := AccessRuleSourceNetworkLiterals{}
			if value := res.Get("value"); value.Exists() {
				data.Value = types.StringValue(value.String())
			} else {
				data.Value = types.StringNull()
			}
			(*parent).SourceNetworkLiterals = append((*parent).SourceNetworkLiterals, data)
			return true
		})
	}
	if value := res.Get("destinationNetworks.literals"); value.Exists() {
		data.DestinationNetworkLiterals = make([]AccessRuleDestinationNetworkLiterals, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := AccessRuleDestinationNetworkLiterals{}
			if value := res.Get("value"); value.Exists() {
				data.Value = types.StringValue(value.String())
			} else {
				data.Value = types.StringNull()
			}
			(*parent).DestinationNetworkLiterals = append((*parent).DestinationNetworkLiterals, data)
			return true
		})
	}
	if value := res.Get("vlanTags.literals"); value.Exists() {
		data.VlanTagLiterals = make([]AccessRuleVlanTagLiterals, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := AccessRuleVlanTagLiterals{}
			if value := res.Get("startTag"); value.Exists() {
				data.StartTag = types.StringValue(value.String())
			} else {
				data.StartTag = types.StringNull()
			}
			if value := res.Get("endTag"); value.Exists() {
				data.EndTag = types.StringValue(value.String())
			} else {
				data.EndTag = types.StringNull()
			}
			(*parent).VlanTagLiterals = append((*parent).VlanTagLiterals, data)
			return true
		})
	}
	if value := res.Get("vlanTags.objects"); value.Exists() {
		data.VlanTagObjects = make([]AccessRuleVlanTagObjects, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := AccessRuleVlanTagObjects{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).VlanTagObjects = append((*parent).VlanTagObjects, data)
			return true
		})
	}
	if value := res.Get("sourceNetworks.objects"); value.Exists() {
		data.SourceNetworkObjects = make([]AccessRuleSourceNetworkObjects, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := AccessRuleSourceNetworkObjects{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			(*parent).SourceNetworkObjects = append((*parent).SourceNetworkObjects, data)
			return true
		})
	}
	if value := res.Get("destinationNetworks.objects"); value.Exists() {
		data.DestinationNetworkObjects = make([]AccessRuleDestinationNetworkObjects, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := AccessRuleDestinationNetworkObjects{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			(*parent).DestinationNetworkObjects = append((*parent).DestinationNetworkObjects, data)
			return true
		})
	}
	if value := res.Get("sourceDynamicObjects.objects"); value.Exists() {
		data.SourceDynamicObjects = make([]AccessRuleSourceDynamicObjects, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := AccessRuleSourceDynamicObjects{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).SourceDynamicObjects = append((*parent).SourceDynamicObjects, data)
			return true
		})
	}
	if value := res.Get("destinationDynamicObjects.objects"); value.Exists() {
		data.DestinationDynamicObjects = make([]AccessRuleDestinationDynamicObjects, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := AccessRuleDestinationDynamicObjects{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).DestinationDynamicObjects = append((*parent).DestinationDynamicObjects, data)
			return true
		})
	}
	if value := res.Get("sourcePorts.literals"); value.Exists() {
		data.SourcePortLiterals = make([]AccessRuleSourcePortLiterals, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := AccessRuleSourcePortLiterals{}
			if value := res.Get("type"); value.Exists() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			if value := res.Get("port"); value.Exists() {
				data.Port = types.StringValue(value.String())
			} else {
				data.Port = types.StringNull()
			}
			if value := res.Get("protocol"); value.Exists() {
				data.Protocol = types.StringValue(value.String())
			} else {
				data.Protocol = types.StringNull()
			}
			if value := res.Get("icmpType"); value.Exists() {
				data.IcmpType = types.StringValue(value.String())
			} else {
				data.IcmpType = types.StringNull()
			}
			if value := res.Get("code"); value.Exists() {
				data.IcmpCode = types.StringValue(value.String())
			} else {
				data.IcmpCode = types.StringNull()
			}
			(*parent).SourcePortLiterals = append((*parent).SourcePortLiterals, data)
			return true
		})
	}
	if value := res.Get("sourcePorts.objects"); value.Exists() {
		data.SourcePortObjects = make([]AccessRuleSourcePortObjects, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := AccessRuleSourcePortObjects{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).SourcePortObjects = append((*parent).SourcePortObjects, data)
			return true
		})
	}
	if value := res.Get("destinationPorts.literals"); value.Exists() {
		data.DestinationPortLiterals = make([]AccessRuleDestinationPortLiterals, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := AccessRuleDestinationPortLiterals{}
			if value := res.Get("type"); value.Exists() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			if value := res.Get("port"); value.Exists() {
				data.Port = types.StringValue(value.String())
			} else {
				data.Port = types.StringNull()
			}
			if value := res.Get("protocol"); value.Exists() {
				data.Protocol = types.StringValue(value.String())
			} else {
				data.Protocol = types.StringNull()
			}
			if value := res.Get("icmpType"); value.Exists() {
				data.IcmpType = types.StringValue(value.String())
			} else {
				data.IcmpType = types.StringNull()
			}
			if value := res.Get("code"); value.Exists() {
				data.IcmpCode = types.StringValue(value.String())
			} else {
				data.IcmpCode = types.StringNull()
			}
			(*parent).DestinationPortLiterals = append((*parent).DestinationPortLiterals, data)
			return true
		})
	}
	if value := res.Get("destinationPorts.objects"); value.Exists() {
		data.DestinationPortObjects = make([]AccessRuleDestinationPortObjects, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := AccessRuleDestinationPortObjects{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).DestinationPortObjects = append((*parent).DestinationPortObjects, data)
			return true
		})
	}
	if value := res.Get("sourceSecurityGroupTags.objects"); value.Exists() {
		data.SourceSgtObjects = make([]AccessRuleSourceSgtObjects, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := AccessRuleSourceSgtObjects{}
			if value := res.Get("name"); value.Exists() {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			(*parent).SourceSgtObjects = append((*parent).SourceSgtObjects, data)
			return true
		})
	}
	if value := res.Get("endPointDeviceTypes"); value.Exists() {
		data.EndpointDeviceTypes = make([]AccessRuleEndpointDeviceTypes, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := AccessRuleEndpointDeviceTypes{}
			if value := res.Get("name"); value.Exists() {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			(*parent).EndpointDeviceTypes = append((*parent).EndpointDeviceTypes, data)
			return true
		})
	}
	if value := res.Get("destinationSecurityGroupTags.objects"); value.Exists() {
		data.DestinationSgtObjects = make([]AccessRuleDestinationSgtObjects, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := AccessRuleDestinationSgtObjects{}
			if value := res.Get("name"); value.Exists() {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			(*parent).DestinationSgtObjects = append((*parent).DestinationSgtObjects, data)
			return true
		})
	}
	if value := res.Get("sourceZones.objects"); value.Exists() {
		data.SourceZones = make([]AccessRuleSourceZones, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := AccessRuleSourceZones{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).SourceZones = append((*parent).SourceZones, data)
			return true
		})
	}
	if value := res.Get("destinationZones.objects"); value.Exists() {
		data.DestinationZones = make([]AccessRuleDestinationZones, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := AccessRuleDestinationZones{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).DestinationZones = append((*parent).DestinationZones, data)
			return true
		})
	}
	if value := res.Get("urls.literals"); value.Exists() {
		data.UrlLiterals = make([]AccessRuleUrlLiterals, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := AccessRuleUrlLiterals{}
			if value := res.Get("url"); value.Exists() {
				data.Url = types.StringValue(value.String())
			} else {
				data.Url = types.StringNull()
			}
			(*parent).UrlLiterals = append((*parent).UrlLiterals, data)
			return true
		})
	}
	if value := res.Get("urls.objects"); value.Exists() {
		data.UrlObjects = make([]AccessRuleUrlObjects, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := AccessRuleUrlObjects{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).UrlObjects = append((*parent).UrlObjects, data)
			return true
		})
	}
	if value := res.Get("urls.urlCategoriesWithReputation"); value.Exists() {
		data.UrlCategories = make([]AccessRuleUrlCategories, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := AccessRuleUrlCategories{}
			if value := res.Get("category.id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("reputation"); value.Exists() {
				data.Reputation = types.StringValue(value.String())
			} else {
				data.Reputation = types.StringNull()
			}
			(*parent).UrlCategories = append((*parent).UrlCategories, data)
			return true
		})
	}
	if value := res.Get("logBegin"); value.Exists() {
		data.LogBegin = types.BoolValue(value.Bool())
	} else {
		data.LogBegin = types.BoolValue(false)
	}
	if value := res.Get("logEnd"); value.Exists() {
		data.LogEnd = types.BoolValue(value.Bool())
	} else {
		data.LogEnd = types.BoolValue(false)
	}
	if value := res.Get("logFiles"); value.Exists() {
		data.LogFiles = types.BoolValue(value.Bool())
	} else {
		data.LogFiles = types.BoolValue(false)
	}
	if value := res.Get("sendEventsToFMC"); value.Exists() {
		data.SendEventsToFmc = types.BoolValue(value.Bool())
	} else {
		data.SendEventsToFmc = types.BoolValue(false)
	}
	if value := res.Get("enableSyslog"); value.Exists() {
		data.SendSyslog = types.BoolValue(value.Bool())
	} else {
		data.SendSyslog = types.BoolValue(false)
	}
	if value := res.Get("syslogConfig.id"); value.Exists() {
		data.SyslogConfigId = types.StringValue(value.String())
	} else {
		data.SyslogConfigId = types.StringNull()
	}
	if value := res.Get("syslogSeverity"); value.Exists() {
		data.SyslogSeverity = types.StringValue(value.String())
	} else {
		data.SyslogSeverity = types.StringNull()
	}
	if value := res.Get("snmpConfig.id"); value.Exists() {
		data.SnmpConfigId = types.StringValue(value.String())
	} else {
		data.SnmpConfigId = types.StringNull()
	}
	if value := res.Get("filePolicy.id"); value.Exists() {
		data.FilePolicyId = types.StringValue(value.String())
	} else {
		data.FilePolicyId = types.StringNull()
	}
	if value := res.Get("ipsPolicy.id"); value.Exists() {
		data.IntrusionPolicyId = types.StringValue(value.String())
	} else {
		data.IntrusionPolicyId = types.StringNull()
	}
	if value := res.Get("timeRangeObjects.0.id"); value.Exists() {
		data.TimeRangeId = types.StringValue(value.String())
	} else {
		data.TimeRangeId = types.StringNull()
	}
	if value := res.Get("variableSet.id"); value.Exists() {
		data.VariableSetId = types.StringValue(value.String())
	} else {
		data.VariableSetId = types.StringNull()
	}
	if value := res.Get("applications.applications"); value.Exists() {
		data.Applications = make([]AccessRuleApplications, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := AccessRuleApplications{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).Applications = append((*parent).Applications, data)
			return true
		})
	}
	if value := res.Get("applications.applicationFilters"); value.Exists() {
		data.ApplicationFilterObjects = make([]AccessRuleApplicationFilterObjects, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := AccessRuleApplicationFilterObjects{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).ApplicationFilterObjects = append((*parent).ApplicationFilterObjects, data)
			return true
		})
	}
	if value := res.Get("applications.inlineApplicationFilters"); value.Exists() {
		data.ApplicationFilters = make([]AccessRuleApplicationFilters, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := AccessRuleApplicationFilters{}
			if value := res.Get("applicationTypes"); value.Exists() {
				data.Types = make([]AccessRuleApplicationFiltersTypes, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := AccessRuleApplicationFiltersTypes{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					(*parent).Types = append((*parent).Types, data)
					return true
				})
			}
			if value := res.Get("risks"); value.Exists() {
				data.Risks = make([]AccessRuleApplicationFiltersRisks, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := AccessRuleApplicationFiltersRisks{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					(*parent).Risks = append((*parent).Risks, data)
					return true
				})
			}
			if value := res.Get("productivities"); value.Exists() {
				data.BusinessRelevances = make([]AccessRuleApplicationFiltersBusinessRelevances, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := AccessRuleApplicationFiltersBusinessRelevances{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					(*parent).BusinessRelevances = append((*parent).BusinessRelevances, data)
					return true
				})
			}
			if value := res.Get("categories"); value.Exists() {
				data.Categories = make([]AccessRuleApplicationFiltersCategories, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := AccessRuleApplicationFiltersCategories{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					(*parent).Categories = append((*parent).Categories, data)
					return true
				})
			}
			if value := res.Get("tags"); value.Exists() {
				data.Tags = make([]AccessRuleApplicationFiltersTags, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := AccessRuleApplicationFiltersTags{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					(*parent).Tags = append((*parent).Tags, data)
					return true
				})
			}
			(*parent).ApplicationFilters = append((*parent).ApplicationFilters, data)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *AccessRule) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("metadata.category"); value.Exists() && !data.CategoryName.IsNull() {
		data.CategoryName = types.StringValue(value.String())
	} else {
		data.CategoryName = types.StringNull()
	}
	if value := res.Get("metadata.section"); value.Exists() && !data.Section.IsNull() {
		data.Section = types.StringValue(value.String())
	} else {
		data.Section = types.StringNull()
	}
	if value := res.Get("action"); value.Exists() && !data.Action.IsNull() {
		data.Action = types.StringValue(value.String())
	} else {
		data.Action = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = types.BoolValue(value.Bool())
	} else if data.Enabled.ValueBool() != true {
		data.Enabled = types.BoolNull()
	}
	for i := 0; i < len(data.SourceNetworkLiterals); i++ {
		keys := [...]string{"value"}
		keyValues := [...]string{data.SourceNetworkLiterals[i].Value.ValueString()}

		parent := &data
		data := (*parent).SourceNetworkLiterals[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("sourceNetworks.literals").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing SourceNetworkLiterals[%d] = %+v",
				i,
				(*parent).SourceNetworkLiterals[i],
			))
			(*parent).SourceNetworkLiterals = slices.Delete((*parent).SourceNetworkLiterals, i, i+1)
			i--

			continue
		}
		if value := res.Get("value"); value.Exists() && !data.Value.IsNull() {
			data.Value = types.StringValue(value.String())
		} else {
			data.Value = types.StringNull()
		}
		(*parent).SourceNetworkLiterals[i] = data
	}
	for i := 0; i < len(data.DestinationNetworkLiterals); i++ {
		keys := [...]string{"value"}
		keyValues := [...]string{data.DestinationNetworkLiterals[i].Value.ValueString()}

		parent := &data
		data := (*parent).DestinationNetworkLiterals[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("destinationNetworks.literals").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing DestinationNetworkLiterals[%d] = %+v",
				i,
				(*parent).DestinationNetworkLiterals[i],
			))
			(*parent).DestinationNetworkLiterals = slices.Delete((*parent).DestinationNetworkLiterals, i, i+1)
			i--

			continue
		}
		if value := res.Get("value"); value.Exists() && !data.Value.IsNull() {
			data.Value = types.StringValue(value.String())
		} else {
			data.Value = types.StringNull()
		}
		(*parent).DestinationNetworkLiterals[i] = data
	}
	for i := 0; i < len(data.VlanTagLiterals); i++ {
		keys := [...]string{"startTag", "endTag"}
		keyValues := [...]string{data.VlanTagLiterals[i].StartTag.ValueString(), data.VlanTagLiterals[i].EndTag.ValueString()}

		parent := &data
		data := (*parent).VlanTagLiterals[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("vlanTags.literals").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing VlanTagLiterals[%d] = %+v",
				i,
				(*parent).VlanTagLiterals[i],
			))
			(*parent).VlanTagLiterals = slices.Delete((*parent).VlanTagLiterals, i, i+1)
			i--

			continue
		}
		if value := res.Get("startTag"); value.Exists() && !data.StartTag.IsNull() {
			data.StartTag = types.StringValue(value.String())
		} else {
			data.StartTag = types.StringNull()
		}
		if value := res.Get("endTag"); value.Exists() && !data.EndTag.IsNull() {
			data.EndTag = types.StringValue(value.String())
		} else {
			data.EndTag = types.StringNull()
		}
		(*parent).VlanTagLiterals[i] = data
	}
	for i := 0; i < len(data.VlanTagObjects); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.VlanTagObjects[i].Id.ValueString()}

		parent := &data
		data := (*parent).VlanTagObjects[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("vlanTags.objects").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing VlanTagObjects[%d] = %+v",
				i,
				(*parent).VlanTagObjects[i],
			))
			(*parent).VlanTagObjects = slices.Delete((*parent).VlanTagObjects, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).VlanTagObjects[i] = data
	}
	for i := 0; i < len(data.SourceNetworkObjects); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.SourceNetworkObjects[i].Id.ValueString()}

		parent := &data
		data := (*parent).SourceNetworkObjects[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("sourceNetworks.objects").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing SourceNetworkObjects[%d] = %+v",
				i,
				(*parent).SourceNetworkObjects[i],
			))
			(*parent).SourceNetworkObjects = slices.Delete((*parent).SourceNetworkObjects, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		(*parent).SourceNetworkObjects[i] = data
	}
	for i := 0; i < len(data.DestinationNetworkObjects); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.DestinationNetworkObjects[i].Id.ValueString()}

		parent := &data
		data := (*parent).DestinationNetworkObjects[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("destinationNetworks.objects").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing DestinationNetworkObjects[%d] = %+v",
				i,
				(*parent).DestinationNetworkObjects[i],
			))
			(*parent).DestinationNetworkObjects = slices.Delete((*parent).DestinationNetworkObjects, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		(*parent).DestinationNetworkObjects[i] = data
	}
	for i := 0; i < len(data.SourceDynamicObjects); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.SourceDynamicObjects[i].Id.ValueString()}

		parent := &data
		data := (*parent).SourceDynamicObjects[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("sourceDynamicObjects.objects").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing SourceDynamicObjects[%d] = %+v",
				i,
				(*parent).SourceDynamicObjects[i],
			))
			(*parent).SourceDynamicObjects = slices.Delete((*parent).SourceDynamicObjects, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).SourceDynamicObjects[i] = data
	}
	for i := 0; i < len(data.DestinationDynamicObjects); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.DestinationDynamicObjects[i].Id.ValueString()}

		parent := &data
		data := (*parent).DestinationDynamicObjects[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("destinationDynamicObjects.objects").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing DestinationDynamicObjects[%d] = %+v",
				i,
				(*parent).DestinationDynamicObjects[i],
			))
			(*parent).DestinationDynamicObjects = slices.Delete((*parent).DestinationDynamicObjects, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).DestinationDynamicObjects[i] = data
	}
	for i := 0; i < len(data.SourcePortLiterals); i++ {
		keys := [...]string{"type", "port", "protocol", "icmpType", "code"}
		keyValues := [...]string{data.SourcePortLiterals[i].Type.ValueString(), data.SourcePortLiterals[i].Port.ValueString(), data.SourcePortLiterals[i].Protocol.ValueString(), data.SourcePortLiterals[i].IcmpType.ValueString(), data.SourcePortLiterals[i].IcmpCode.ValueString()}

		parent := &data
		data := (*parent).SourcePortLiterals[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("sourcePorts.literals").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing SourcePortLiterals[%d] = %+v",
				i,
				(*parent).SourcePortLiterals[i],
			))
			(*parent).SourcePortLiterals = slices.Delete((*parent).SourcePortLiterals, i, i+1)
			i--

			continue
		}
		if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		if value := res.Get("port"); value.Exists() && !data.Port.IsNull() {
			data.Port = types.StringValue(value.String())
		} else {
			data.Port = types.StringNull()
		}
		if value := res.Get("protocol"); value.Exists() && !data.Protocol.IsNull() {
			data.Protocol = types.StringValue(value.String())
		} else {
			data.Protocol = types.StringNull()
		}
		if value := res.Get("icmpType"); value.Exists() && !data.IcmpType.IsNull() {
			data.IcmpType = types.StringValue(value.String())
		} else {
			data.IcmpType = types.StringNull()
		}
		if value := res.Get("code"); value.Exists() && !data.IcmpCode.IsNull() {
			data.IcmpCode = types.StringValue(value.String())
		} else {
			data.IcmpCode = types.StringNull()
		}
		(*parent).SourcePortLiterals[i] = data
	}
	for i := 0; i < len(data.SourcePortObjects); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.SourcePortObjects[i].Id.ValueString()}

		parent := &data
		data := (*parent).SourcePortObjects[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("sourcePorts.objects").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing SourcePortObjects[%d] = %+v",
				i,
				(*parent).SourcePortObjects[i],
			))
			(*parent).SourcePortObjects = slices.Delete((*parent).SourcePortObjects, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).SourcePortObjects[i] = data
	}
	for i := 0; i < len(data.DestinationPortLiterals); i++ {
		keys := [...]string{"type", "port", "protocol", "icmpType", "code"}
		keyValues := [...]string{data.DestinationPortLiterals[i].Type.ValueString(), data.DestinationPortLiterals[i].Port.ValueString(), data.DestinationPortLiterals[i].Protocol.ValueString(), data.DestinationPortLiterals[i].IcmpType.ValueString(), data.DestinationPortLiterals[i].IcmpCode.ValueString()}

		parent := &data
		data := (*parent).DestinationPortLiterals[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("destinationPorts.literals").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing DestinationPortLiterals[%d] = %+v",
				i,
				(*parent).DestinationPortLiterals[i],
			))
			(*parent).DestinationPortLiterals = slices.Delete((*parent).DestinationPortLiterals, i, i+1)
			i--

			continue
		}
		if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		if value := res.Get("port"); value.Exists() && !data.Port.IsNull() {
			data.Port = types.StringValue(value.String())
		} else {
			data.Port = types.StringNull()
		}
		if value := res.Get("protocol"); value.Exists() && !data.Protocol.IsNull() {
			data.Protocol = types.StringValue(value.String())
		} else {
			data.Protocol = types.StringNull()
		}
		if value := res.Get("icmpType"); value.Exists() && !data.IcmpType.IsNull() {
			data.IcmpType = types.StringValue(value.String())
		} else {
			data.IcmpType = types.StringNull()
		}
		if value := res.Get("code"); value.Exists() && !data.IcmpCode.IsNull() {
			data.IcmpCode = types.StringValue(value.String())
		} else {
			data.IcmpCode = types.StringNull()
		}
		(*parent).DestinationPortLiterals[i] = data
	}
	for i := 0; i < len(data.DestinationPortObjects); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.DestinationPortObjects[i].Id.ValueString()}

		parent := &data
		data := (*parent).DestinationPortObjects[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("destinationPorts.objects").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing DestinationPortObjects[%d] = %+v",
				i,
				(*parent).DestinationPortObjects[i],
			))
			(*parent).DestinationPortObjects = slices.Delete((*parent).DestinationPortObjects, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).DestinationPortObjects[i] = data
	}
	for i := 0; i < len(data.SourceSgtObjects); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.SourceSgtObjects[i].Id.ValueString()}

		parent := &data
		data := (*parent).SourceSgtObjects[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("sourceSecurityGroupTags.objects").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing SourceSgtObjects[%d] = %+v",
				i,
				(*parent).SourceSgtObjects[i],
			))
			(*parent).SourceSgtObjects = slices.Delete((*parent).SourceSgtObjects, i, i+1)
			i--

			continue
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		(*parent).SourceSgtObjects[i] = data
	}
	for i := 0; i < len(data.EndpointDeviceTypes); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.EndpointDeviceTypes[i].Id.ValueString()}

		parent := &data
		data := (*parent).EndpointDeviceTypes[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("endPointDeviceTypes").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing EndpointDeviceTypes[%d] = %+v",
				i,
				(*parent).EndpointDeviceTypes[i],
			))
			(*parent).EndpointDeviceTypes = slices.Delete((*parent).EndpointDeviceTypes, i, i+1)
			i--

			continue
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		(*parent).EndpointDeviceTypes[i] = data
	}
	for i := 0; i < len(data.DestinationSgtObjects); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.DestinationSgtObjects[i].Id.ValueString()}

		parent := &data
		data := (*parent).DestinationSgtObjects[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("destinationSecurityGroupTags.objects").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing DestinationSgtObjects[%d] = %+v",
				i,
				(*parent).DestinationSgtObjects[i],
			))
			(*parent).DestinationSgtObjects = slices.Delete((*parent).DestinationSgtObjects, i, i+1)
			i--

			continue
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		(*parent).DestinationSgtObjects[i] = data
	}
	for i := 0; i < len(data.SourceZones); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.SourceZones[i].Id.ValueString()}

		parent := &data
		data := (*parent).SourceZones[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("sourceZones.objects").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing SourceZones[%d] = %+v",
				i,
				(*parent).SourceZones[i],
			))
			(*parent).SourceZones = slices.Delete((*parent).SourceZones, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).SourceZones[i] = data
	}
	for i := 0; i < len(data.DestinationZones); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.DestinationZones[i].Id.ValueString()}

		parent := &data
		data := (*parent).DestinationZones[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("destinationZones.objects").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing DestinationZones[%d] = %+v",
				i,
				(*parent).DestinationZones[i],
			))
			(*parent).DestinationZones = slices.Delete((*parent).DestinationZones, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).DestinationZones[i] = data
	}
	for i := 0; i < len(data.UrlLiterals); i++ {
		keys := [...]string{"url"}
		keyValues := [...]string{data.UrlLiterals[i].Url.ValueString()}

		parent := &data
		data := (*parent).UrlLiterals[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("urls.literals").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing UrlLiterals[%d] = %+v",
				i,
				(*parent).UrlLiterals[i],
			))
			(*parent).UrlLiterals = slices.Delete((*parent).UrlLiterals, i, i+1)
			i--

			continue
		}
		if value := res.Get("url"); value.Exists() && !data.Url.IsNull() {
			data.Url = types.StringValue(value.String())
		} else {
			data.Url = types.StringNull()
		}
		(*parent).UrlLiterals[i] = data
	}
	for i := 0; i < len(data.UrlObjects); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.UrlObjects[i].Id.ValueString()}

		parent := &data
		data := (*parent).UrlObjects[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("urls.objects").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing UrlObjects[%d] = %+v",
				i,
				(*parent).UrlObjects[i],
			))
			(*parent).UrlObjects = slices.Delete((*parent).UrlObjects, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).UrlObjects[i] = data
	}
	for i := 0; i < len(data.UrlCategories); i++ {
		keys := [...]string{"category.id"}
		keyValues := [...]string{data.UrlCategories[i].Id.ValueString()}

		parent := &data
		data := (*parent).UrlCategories[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("urls.urlCategoriesWithReputation").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing UrlCategories[%d] = %+v",
				i,
				(*parent).UrlCategories[i],
			))
			(*parent).UrlCategories = slices.Delete((*parent).UrlCategories, i, i+1)
			i--

			continue
		}
		if value := res.Get("category.id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("reputation"); value.Exists() && !data.Reputation.IsNull() {
			data.Reputation = types.StringValue(value.String())
		} else {
			data.Reputation = types.StringNull()
		}
		(*parent).UrlCategories[i] = data
	}
	if value := res.Get("logBegin"); value.Exists() && !data.LogBegin.IsNull() {
		data.LogBegin = types.BoolValue(value.Bool())
	} else if data.LogBegin.ValueBool() != false {
		data.LogBegin = types.BoolNull()
	}
	if value := res.Get("logEnd"); value.Exists() && !data.LogEnd.IsNull() {
		data.LogEnd = types.BoolValue(value.Bool())
	} else if data.LogEnd.ValueBool() != false {
		data.LogEnd = types.BoolNull()
	}
	if value := res.Get("logFiles"); value.Exists() && !data.LogFiles.IsNull() {
		data.LogFiles = types.BoolValue(value.Bool())
	} else if data.LogFiles.ValueBool() != false {
		data.LogFiles = types.BoolNull()
	}
	if value := res.Get("sendEventsToFMC"); value.Exists() && !data.SendEventsToFmc.IsNull() {
		data.SendEventsToFmc = types.BoolValue(value.Bool())
	} else if data.SendEventsToFmc.ValueBool() != false {
		data.SendEventsToFmc = types.BoolNull()
	}
	if value := res.Get("enableSyslog"); value.Exists() && !data.SendSyslog.IsNull() {
		data.SendSyslog = types.BoolValue(value.Bool())
	} else if data.SendSyslog.ValueBool() != false {
		data.SendSyslog = types.BoolNull()
	}
	if value := res.Get("syslogConfig.id"); value.Exists() && !data.SyslogConfigId.IsNull() {
		data.SyslogConfigId = types.StringValue(value.String())
	} else {
		data.SyslogConfigId = types.StringNull()
	}
	if value := res.Get("syslogSeverity"); value.Exists() && !data.SyslogSeverity.IsNull() {
		data.SyslogSeverity = types.StringValue(value.String())
	} else {
		data.SyslogSeverity = types.StringNull()
	}
	if value := res.Get("snmpConfig.id"); value.Exists() && !data.SnmpConfigId.IsNull() {
		data.SnmpConfigId = types.StringValue(value.String())
	} else {
		data.SnmpConfigId = types.StringNull()
	}
	if value := res.Get("filePolicy.id"); value.Exists() && !data.FilePolicyId.IsNull() {
		data.FilePolicyId = types.StringValue(value.String())
	} else {
		data.FilePolicyId = types.StringNull()
	}
	if value := res.Get("ipsPolicy.id"); value.Exists() && !data.IntrusionPolicyId.IsNull() {
		data.IntrusionPolicyId = types.StringValue(value.String())
	} else {
		data.IntrusionPolicyId = types.StringNull()
	}
	if value := res.Get("timeRangeObjects.0.id"); value.Exists() && !data.TimeRangeId.IsNull() {
		data.TimeRangeId = types.StringValue(value.String())
	} else {
		data.TimeRangeId = types.StringNull()
	}
	if value := res.Get("variableSet.id"); value.Exists() && !data.VariableSetId.IsNull() {
		data.VariableSetId = types.StringValue(value.String())
	} else {
		data.VariableSetId = types.StringNull()
	}
	for i := 0; i < len(data.Applications); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.Applications[i].Id.ValueString()}

		parent := &data
		data := (*parent).Applications[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("applications.applications").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing Applications[%d] = %+v",
				i,
				(*parent).Applications[i],
			))
			(*parent).Applications = slices.Delete((*parent).Applications, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).Applications[i] = data
	}
	for i := 0; i < len(data.ApplicationFilterObjects); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.ApplicationFilterObjects[i].Id.ValueString()}

		parent := &data
		data := (*parent).ApplicationFilterObjects[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("applications.applicationFilters").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing ApplicationFilterObjects[%d] = %+v",
				i,
				(*parent).ApplicationFilterObjects[i],
			))
			(*parent).ApplicationFilterObjects = slices.Delete((*parent).ApplicationFilterObjects, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).ApplicationFilterObjects[i] = data
	}
	{
		l := len(res.Get("applications.inlineApplicationFilters").Array())
		tflog.Debug(ctx, fmt.Sprintf("applications.inlineApplicationFilters array resizing from %d to %d", len(data.ApplicationFilters), l))
		for i := len(data.ApplicationFilters); i < l; i++ {
			data.ApplicationFilters = append(data.ApplicationFilters, AccessRuleApplicationFilters{})
		}
		if len(data.ApplicationFilters) > l {
			data.ApplicationFilters = data.ApplicationFilters[:l]
		}
	}
	for i := range data.ApplicationFilters {
		parent := &data
		data := (*parent).ApplicationFilters[i]
		parentRes := &res
		res := parentRes.Get(fmt.Sprintf("applications.inlineApplicationFilters.%d", i))
		for i := 0; i < len(data.Types); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.Types[i].Id.ValueString()}

			parent := &data
			data := (*parent).Types[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("applicationTypes").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() != keyValues[ik] {
							found = false
							break
						}
						found = true
					}
					if found {
						res = v
						return false
					}
					return true
				},
			)
			if !res.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing Types[%d] = %+v",
					i,
					(*parent).Types[i],
				))
				(*parent).Types = slices.Delete((*parent).Types, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).Types[i] = data
		}
		for i := 0; i < len(data.Risks); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.Risks[i].Id.ValueString()}

			parent := &data
			data := (*parent).Risks[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("risks").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() != keyValues[ik] {
							found = false
							break
						}
						found = true
					}
					if found {
						res = v
						return false
					}
					return true
				},
			)
			if !res.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing Risks[%d] = %+v",
					i,
					(*parent).Risks[i],
				))
				(*parent).Risks = slices.Delete((*parent).Risks, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).Risks[i] = data
		}
		for i := 0; i < len(data.BusinessRelevances); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.BusinessRelevances[i].Id.ValueString()}

			parent := &data
			data := (*parent).BusinessRelevances[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("productivities").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() != keyValues[ik] {
							found = false
							break
						}
						found = true
					}
					if found {
						res = v
						return false
					}
					return true
				},
			)
			if !res.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing BusinessRelevances[%d] = %+v",
					i,
					(*parent).BusinessRelevances[i],
				))
				(*parent).BusinessRelevances = slices.Delete((*parent).BusinessRelevances, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).BusinessRelevances[i] = data
		}
		for i := 0; i < len(data.Categories); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.Categories[i].Id.ValueString()}

			parent := &data
			data := (*parent).Categories[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("categories").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() != keyValues[ik] {
							found = false
							break
						}
						found = true
					}
					if found {
						res = v
						return false
					}
					return true
				},
			)
			if !res.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing Categories[%d] = %+v",
					i,
					(*parent).Categories[i],
				))
				(*parent).Categories = slices.Delete((*parent).Categories, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).Categories[i] = data
		}
		for i := 0; i < len(data.Tags); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.Tags[i].Id.ValueString()}

			parent := &data
			data := (*parent).Tags[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("tags").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() != keyValues[ik] {
							found = false
							break
						}
						found = true
					}
					if found {
						res = v
						return false
					}
					return true
				},
			)
			if !res.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing Tags[%d] = %+v",
					i,
					(*parent).Tags[i],
				))
				(*parent).Tags = slices.Delete((*parent).Tags, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).Tags[i] = data
		}
		(*parent).ApplicationFilters[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *AccessRule) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin Clone

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// End of section. //template:end toBodyNonBulk

// Section below is generated&owned by "gen/generator.go". //template:begin findObjectsToBeReplaced

// End of section. //template:end findObjectsToBeReplaced

// Section below is generated&owned by "gen/generator.go". //template:begin clearItemIds

// End of section. //template:end clearItemIds

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyPutDelete

// End of section. //template:end toBodyPutDelete

func (data AccessRule) adjustBody(ctx context.Context, req string) string {
	// metadata includes category_name and section
	req, _ = sjson.Delete(req, "metadata")
	return req
}

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBodyBulk

// End of section. //template:end adjustBodyBulk
