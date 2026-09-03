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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type AccessRules struct {
	Id                    types.String       `tfsdk:"id"`
	Domain                types.String       `tfsdk:"domain"`
	AccessControlPolicyId types.String       `tfsdk:"access_control_policy_id"`
	CategoryName          types.String       `tfsdk:"category_name"`
	Section               types.String       `tfsdk:"section"`
	Items                 []AccessRulesItems `tfsdk:"items"`
}

type AccessRulesItems struct {
	Id                         types.String                                 `tfsdk:"id"`
	Action                     types.String                                 `tfsdk:"action"`
	Name                       types.String                                 `tfsdk:"name"`
	Enabled                    types.Bool                                   `tfsdk:"enabled"`
	SourceNetworkLiterals      []AccessRulesItemsSourceNetworkLiterals      `tfsdk:"source_network_literals"`
	DestinationNetworkLiterals []AccessRulesItemsDestinationNetworkLiterals `tfsdk:"destination_network_literals"`
	VlanTagLiterals            []AccessRulesItemsVlanTagLiterals            `tfsdk:"vlan_tag_literals"`
	VlanTagObjects             []AccessRulesItemsVlanTagObjects             `tfsdk:"vlan_tag_objects"`
	SourceNetworkObjects       []AccessRulesItemsSourceNetworkObjects       `tfsdk:"source_network_objects"`
	DestinationNetworkObjects  []AccessRulesItemsDestinationNetworkObjects  `tfsdk:"destination_network_objects"`
	SourceDynamicObjects       []AccessRulesItemsSourceDynamicObjects       `tfsdk:"source_dynamic_objects"`
	DestinationDynamicObjects  []AccessRulesItemsDestinationDynamicObjects  `tfsdk:"destination_dynamic_objects"`
	SourcePortLiterals         []AccessRulesItemsSourcePortLiterals         `tfsdk:"source_port_literals"`
	SourcePortObjects          []AccessRulesItemsSourcePortObjects          `tfsdk:"source_port_objects"`
	DestinationPortLiterals    []AccessRulesItemsDestinationPortLiterals    `tfsdk:"destination_port_literals"`
	DestinationPortObjects     []AccessRulesItemsDestinationPortObjects     `tfsdk:"destination_port_objects"`
	SourceSgtObjects           []AccessRulesItemsSourceSgtObjects           `tfsdk:"source_sgt_objects"`
	EndpointDeviceTypes        []AccessRulesItemsEndpointDeviceTypes        `tfsdk:"endpoint_device_types"`
	DestinationSgtObjects      []AccessRulesItemsDestinationSgtObjects      `tfsdk:"destination_sgt_objects"`
	SourceZones                []AccessRulesItemsSourceZones                `tfsdk:"source_zones"`
	DestinationZones           []AccessRulesItemsDestinationZones           `tfsdk:"destination_zones"`
	UrlLiterals                []AccessRulesItemsUrlLiterals                `tfsdk:"url_literals"`
	UrlObjects                 []AccessRulesItemsUrlObjects                 `tfsdk:"url_objects"`
	UrlCategories              []AccessRulesItemsUrlCategories              `tfsdk:"url_categories"`
	LogConnectionBegin         types.Bool                                   `tfsdk:"log_connection_begin"`
	LogConnectionEnd           types.Bool                                   `tfsdk:"log_connection_end"`
	LogFiles                   types.Bool                                   `tfsdk:"log_files"`
	SendEventsToFmc            types.Bool                                   `tfsdk:"send_events_to_fmc"`
	SendSyslog                 types.Bool                                   `tfsdk:"send_syslog"`
	SyslogAlertId              types.String                                 `tfsdk:"syslog_alert_id"`
	SyslogSeverity             types.String                                 `tfsdk:"syslog_severity"`
	SnmpAlertId                types.String                                 `tfsdk:"snmp_alert_id"`
	Description                types.String                                 `tfsdk:"description"`
	FilePolicyId               types.String                                 `tfsdk:"file_policy_id"`
	IntrusionPolicyId          types.String                                 `tfsdk:"intrusion_policy_id"`
	TimeRangeId                types.String                                 `tfsdk:"time_range_id"`
	VariableSetId              types.String                                 `tfsdk:"variable_set_id"`
	Applications               []AccessRulesItemsApplications               `tfsdk:"applications"`
	ApplicationFilterObjects   []AccessRulesItemsApplicationFilterObjects   `tfsdk:"application_filter_objects"`
	ApplicationFilters         []AccessRulesItemsApplicationFilters         `tfsdk:"application_filters"`
	UserObjects                []AccessRulesItemsUserObjects                `tfsdk:"user_objects"`
}

type AccessRulesItemsSourceNetworkLiterals struct {
	Value types.String `tfsdk:"value"`
}
type AccessRulesItemsDestinationNetworkLiterals struct {
	Value types.String `tfsdk:"value"`
}
type AccessRulesItemsVlanTagLiterals struct {
	StartTag types.String `tfsdk:"start_tag"`
	EndTag   types.String `tfsdk:"end_tag"`
}
type AccessRulesItemsVlanTagObjects struct {
	Id types.String `tfsdk:"id"`
}
type AccessRulesItemsSourceNetworkObjects struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}
type AccessRulesItemsDestinationNetworkObjects struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}
type AccessRulesItemsSourceDynamicObjects struct {
	Id types.String `tfsdk:"id"`
}
type AccessRulesItemsDestinationDynamicObjects struct {
	Id types.String `tfsdk:"id"`
}
type AccessRulesItemsSourcePortLiterals struct {
	Protocol types.String `tfsdk:"protocol"`
	Port     types.String `tfsdk:"port"`
}
type AccessRulesItemsSourcePortObjects struct {
	Id types.String `tfsdk:"id"`
}
type AccessRulesItemsDestinationPortLiterals struct {
	Type     types.String `tfsdk:"type"`
	Port     types.String `tfsdk:"port"`
	Protocol types.String `tfsdk:"protocol"`
	IcmpType types.String `tfsdk:"icmp_type"`
	IcmpCode types.String `tfsdk:"icmp_code"`
}
type AccessRulesItemsDestinationPortObjects struct {
	Id types.String `tfsdk:"id"`
}
type AccessRulesItemsSourceSgtObjects struct {
	Name types.String `tfsdk:"name"`
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}
type AccessRulesItemsEndpointDeviceTypes struct {
	Name types.String `tfsdk:"name"`
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}
type AccessRulesItemsDestinationSgtObjects struct {
	Name types.String `tfsdk:"name"`
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}
type AccessRulesItemsSourceZones struct {
	Id types.String `tfsdk:"id"`
}
type AccessRulesItemsDestinationZones struct {
	Id types.String `tfsdk:"id"`
}
type AccessRulesItemsUrlLiterals struct {
	Url types.String `tfsdk:"url"`
}
type AccessRulesItemsUrlObjects struct {
	Id types.String `tfsdk:"id"`
}
type AccessRulesItemsUrlCategories struct {
	Id         types.String `tfsdk:"id"`
	Reputation types.String `tfsdk:"reputation"`
}
type AccessRulesItemsApplications struct {
	Id types.String `tfsdk:"id"`
}
type AccessRulesItemsApplicationFilterObjects struct {
	Id types.String `tfsdk:"id"`
}
type AccessRulesItemsApplicationFilters struct {
	Types              []AccessRulesItemsApplicationFiltersTypes              `tfsdk:"types"`
	Risks              []AccessRulesItemsApplicationFiltersRisks              `tfsdk:"risks"`
	BusinessRelevances []AccessRulesItemsApplicationFiltersBusinessRelevances `tfsdk:"business_relevances"`
	Categories         []AccessRulesItemsApplicationFiltersCategories         `tfsdk:"categories"`
	Tags               []AccessRulesItemsApplicationFiltersTags               `tfsdk:"tags"`
}
type AccessRulesItemsUserObjects struct {
	Id        types.String `tfsdk:"id"`
	Type      types.String `tfsdk:"type"`
	Name      types.String `tfsdk:"name"`
	RealmId   types.String `tfsdk:"realm_id"`
	RealmType types.String `tfsdk:"realm_type"`
	RealmName types.String `tfsdk:"realm_name"`
}

type AccessRulesItemsApplicationFiltersTypes struct {
	Id types.String `tfsdk:"id"`
}
type AccessRulesItemsApplicationFiltersRisks struct {
	Id types.String `tfsdk:"id"`
}
type AccessRulesItemsApplicationFiltersBusinessRelevances struct {
	Id types.String `tfsdk:"id"`
}
type AccessRulesItemsApplicationFiltersCategories struct {
	Id types.String `tfsdk:"id"`
}
type AccessRulesItemsApplicationFiltersTags struct {
	Id types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data AccessRules) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/accesspolicies/%v/accessrules", url.QueryEscape(data.AccessControlPolicyId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data AccessRules) toBody(ctx context.Context, state AccessRules) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.CategoryName.IsNull() {
		body, _ = sjson.Set(body, "category", data.CategoryName.ValueString())
	}
	if !data.Section.IsNull() {
		body, _ = sjson.Set(body, "section", data.Section.ValueString())
	}
	if len(data.Items) > 0 {
		var itemsBody strings.Builder
		itemsBody.WriteString("[")
		for _, item := range data.Items {
			itemBody := ""
			if !item.Id.IsNull() && !item.Id.IsUnknown() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Action.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "action", item.Action.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.Enabled.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enabled", item.Enabled.ValueBool())
			}
			if len(item.SourceNetworkLiterals) > 0 {
				var sourceNetworkLiteralsChildBody strings.Builder
				sourceNetworkLiteralsChildBody.WriteString("[")
				for _, childItem := range item.SourceNetworkLiterals {
					itemChildBody := ""
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "AnyNonEmptyString")
					if !childItem.Value.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Value.ValueString())
					}
					if itemChildBody != "" {
						if sourceNetworkLiteralsChildBody.Len() > 1 {
							sourceNetworkLiteralsChildBody.WriteString(",")
						}
						sourceNetworkLiteralsChildBody.WriteString(itemChildBody)
					}
				}
				sourceNetworkLiteralsChildBody.WriteString("]")
				itemBody, _ = sjson.SetRaw(itemBody, "sourceNetworks.literals", sourceNetworkLiteralsChildBody.String())
			}
			if len(item.DestinationNetworkLiterals) > 0 {
				var destinationNetworkLiteralsChildBody strings.Builder
				destinationNetworkLiteralsChildBody.WriteString("[")
				for _, childItem := range item.DestinationNetworkLiterals {
					itemChildBody := ""
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "AnyNonEmptyString")
					if !childItem.Value.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Value.ValueString())
					}
					if itemChildBody != "" {
						if destinationNetworkLiteralsChildBody.Len() > 1 {
							destinationNetworkLiteralsChildBody.WriteString(",")
						}
						destinationNetworkLiteralsChildBody.WriteString(itemChildBody)
					}
				}
				destinationNetworkLiteralsChildBody.WriteString("]")
				itemBody, _ = sjson.SetRaw(itemBody, "destinationNetworks.literals", destinationNetworkLiteralsChildBody.String())
			}
			if len(item.VlanTagLiterals) > 0 {
				var vlanTagLiteralsChildBody strings.Builder
				vlanTagLiteralsChildBody.WriteString("[")
				for _, childItem := range item.VlanTagLiterals {
					itemChildBody := ""
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "AnyNonEmptyString")
					if !childItem.StartTag.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "startTag", childItem.StartTag.ValueString())
					}
					if !childItem.EndTag.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "endTag", childItem.EndTag.ValueString())
					}
					if itemChildBody != "" {
						if vlanTagLiteralsChildBody.Len() > 1 {
							vlanTagLiteralsChildBody.WriteString(",")
						}
						vlanTagLiteralsChildBody.WriteString(itemChildBody)
					}
				}
				vlanTagLiteralsChildBody.WriteString("]")
				itemBody, _ = sjson.SetRaw(itemBody, "vlanTags.literals", vlanTagLiteralsChildBody.String())
			}
			if len(item.VlanTagObjects) > 0 {
				var vlanTagObjectsChildBody strings.Builder
				vlanTagObjectsChildBody.WriteString("[")
				for _, childItem := range item.VlanTagObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if itemChildBody != "" {
						if vlanTagObjectsChildBody.Len() > 1 {
							vlanTagObjectsChildBody.WriteString(",")
						}
						vlanTagObjectsChildBody.WriteString(itemChildBody)
					}
				}
				vlanTagObjectsChildBody.WriteString("]")
				itemBody, _ = sjson.SetRaw(itemBody, "vlanTags.objects", vlanTagObjectsChildBody.String())
			}
			if len(item.SourceNetworkObjects) > 0 {
				var sourceNetworkObjectsChildBody strings.Builder
				sourceNetworkObjectsChildBody.WriteString("[")
				for _, childItem := range item.SourceNetworkObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					if itemChildBody != "" {
						if sourceNetworkObjectsChildBody.Len() > 1 {
							sourceNetworkObjectsChildBody.WriteString(",")
						}
						sourceNetworkObjectsChildBody.WriteString(itemChildBody)
					}
				}
				sourceNetworkObjectsChildBody.WriteString("]")
				itemBody, _ = sjson.SetRaw(itemBody, "sourceNetworks.objects", sourceNetworkObjectsChildBody.String())
			}
			if len(item.DestinationNetworkObjects) > 0 {
				var destinationNetworkObjectsChildBody strings.Builder
				destinationNetworkObjectsChildBody.WriteString("[")
				for _, childItem := range item.DestinationNetworkObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					if itemChildBody != "" {
						if destinationNetworkObjectsChildBody.Len() > 1 {
							destinationNetworkObjectsChildBody.WriteString(",")
						}
						destinationNetworkObjectsChildBody.WriteString(itemChildBody)
					}
				}
				destinationNetworkObjectsChildBody.WriteString("]")
				itemBody, _ = sjson.SetRaw(itemBody, "destinationNetworks.objects", destinationNetworkObjectsChildBody.String())
			}
			if len(item.SourceDynamicObjects) > 0 {
				var sourceDynamicObjectsChildBody strings.Builder
				sourceDynamicObjectsChildBody.WriteString("[")
				for _, childItem := range item.SourceDynamicObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "DynamicObject")
					if itemChildBody != "" {
						if sourceDynamicObjectsChildBody.Len() > 1 {
							sourceDynamicObjectsChildBody.WriteString(",")
						}
						sourceDynamicObjectsChildBody.WriteString(itemChildBody)
					}
				}
				sourceDynamicObjectsChildBody.WriteString("]")
				itemBody, _ = sjson.SetRaw(itemBody, "sourceDynamicObjects.objects", sourceDynamicObjectsChildBody.String())
			}
			if len(item.DestinationDynamicObjects) > 0 {
				var destinationDynamicObjectsChildBody strings.Builder
				destinationDynamicObjectsChildBody.WriteString("[")
				for _, childItem := range item.DestinationDynamicObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "DynamicObject")
					if itemChildBody != "" {
						if destinationDynamicObjectsChildBody.Len() > 1 {
							destinationDynamicObjectsChildBody.WriteString(",")
						}
						destinationDynamicObjectsChildBody.WriteString(itemChildBody)
					}
				}
				destinationDynamicObjectsChildBody.WriteString("]")
				itemBody, _ = sjson.SetRaw(itemBody, "destinationDynamicObjects.objects", destinationDynamicObjectsChildBody.String())
			}
			if len(item.SourcePortLiterals) > 0 {
				var sourcePortLiteralsChildBody strings.Builder
				sourcePortLiteralsChildBody.WriteString("[")
				for _, childItem := range item.SourcePortLiterals {
					itemChildBody := ""
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "PortLiteral")
					if !childItem.Protocol.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "protocol", childItem.Protocol.ValueString())
					}
					if !childItem.Port.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "port", childItem.Port.ValueString())
					}
					if itemChildBody != "" {
						if sourcePortLiteralsChildBody.Len() > 1 {
							sourcePortLiteralsChildBody.WriteString(",")
						}
						sourcePortLiteralsChildBody.WriteString(itemChildBody)
					}
				}
				sourcePortLiteralsChildBody.WriteString("]")
				itemBody, _ = sjson.SetRaw(itemBody, "sourcePorts.literals", sourcePortLiteralsChildBody.String())
			}
			if len(item.SourcePortObjects) > 0 {
				var sourcePortObjectsChildBody strings.Builder
				sourcePortObjectsChildBody.WriteString("[")
				for _, childItem := range item.SourcePortObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "AnyNonEmptyString")
					if itemChildBody != "" {
						if sourcePortObjectsChildBody.Len() > 1 {
							sourcePortObjectsChildBody.WriteString(",")
						}
						sourcePortObjectsChildBody.WriteString(itemChildBody)
					}
				}
				sourcePortObjectsChildBody.WriteString("]")
				itemBody, _ = sjson.SetRaw(itemBody, "sourcePorts.objects", sourcePortObjectsChildBody.String())
			}
			if len(item.DestinationPortLiterals) > 0 {
				var destinationPortLiteralsChildBody strings.Builder
				destinationPortLiteralsChildBody.WriteString("[")
				for _, childItem := range item.DestinationPortLiterals {
					itemChildBody := ""
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					if !childItem.Port.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "port", childItem.Port.ValueString())
					}
					if !childItem.Protocol.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "protocol", childItem.Protocol.ValueString())
					}
					if !childItem.IcmpType.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "icmpType", childItem.IcmpType.ValueString())
					}
					if !childItem.IcmpCode.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "code", childItem.IcmpCode.ValueString())
					}
					if itemChildBody != "" {
						if destinationPortLiteralsChildBody.Len() > 1 {
							destinationPortLiteralsChildBody.WriteString(",")
						}
						destinationPortLiteralsChildBody.WriteString(itemChildBody)
					}
				}
				destinationPortLiteralsChildBody.WriteString("]")
				itemBody, _ = sjson.SetRaw(itemBody, "destinationPorts.literals", destinationPortLiteralsChildBody.String())
			}
			if len(item.DestinationPortObjects) > 0 {
				var destinationPortObjectsChildBody strings.Builder
				destinationPortObjectsChildBody.WriteString("[")
				for _, childItem := range item.DestinationPortObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "AnyNonEmptyString")
					if itemChildBody != "" {
						if destinationPortObjectsChildBody.Len() > 1 {
							destinationPortObjectsChildBody.WriteString(",")
						}
						destinationPortObjectsChildBody.WriteString(itemChildBody)
					}
				}
				destinationPortObjectsChildBody.WriteString("]")
				itemBody, _ = sjson.SetRaw(itemBody, "destinationPorts.objects", destinationPortObjectsChildBody.String())
			}
			if len(item.SourceSgtObjects) > 0 {
				var sourceSgtObjectsChildBody strings.Builder
				sourceSgtObjectsChildBody.WriteString("[")
				for _, childItem := range item.SourceSgtObjects {
					itemChildBody := ""
					if !childItem.Name.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "name", childItem.Name.ValueString())
					}
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					if itemChildBody != "" {
						if sourceSgtObjectsChildBody.Len() > 1 {
							sourceSgtObjectsChildBody.WriteString(",")
						}
						sourceSgtObjectsChildBody.WriteString(itemChildBody)
					}
				}
				sourceSgtObjectsChildBody.WriteString("]")
				itemBody, _ = sjson.SetRaw(itemBody, "sourceSecurityGroupTags.objects", sourceSgtObjectsChildBody.String())
			}
			if len(item.EndpointDeviceTypes) > 0 {
				var endpointDeviceTypesChildBody strings.Builder
				endpointDeviceTypesChildBody.WriteString("[")
				for _, childItem := range item.EndpointDeviceTypes {
					itemChildBody := ""
					if !childItem.Name.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "name", childItem.Name.ValueString())
					}
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					if itemChildBody != "" {
						if endpointDeviceTypesChildBody.Len() > 1 {
							endpointDeviceTypesChildBody.WriteString(",")
						}
						endpointDeviceTypesChildBody.WriteString(itemChildBody)
					}
				}
				endpointDeviceTypesChildBody.WriteString("]")
				itemBody, _ = sjson.SetRaw(itemBody, "endPointDeviceTypes", endpointDeviceTypesChildBody.String())
			}
			if len(item.DestinationSgtObjects) > 0 {
				var destinationSgtObjectsChildBody strings.Builder
				destinationSgtObjectsChildBody.WriteString("[")
				for _, childItem := range item.DestinationSgtObjects {
					itemChildBody := ""
					if !childItem.Name.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "name", childItem.Name.ValueString())
					}
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					if itemChildBody != "" {
						if destinationSgtObjectsChildBody.Len() > 1 {
							destinationSgtObjectsChildBody.WriteString(",")
						}
						destinationSgtObjectsChildBody.WriteString(itemChildBody)
					}
				}
				destinationSgtObjectsChildBody.WriteString("]")
				itemBody, _ = sjson.SetRaw(itemBody, "destinationSecurityGroupTags.objects", destinationSgtObjectsChildBody.String())
			}
			if len(item.SourceZones) > 0 {
				var sourceZonesChildBody strings.Builder
				sourceZonesChildBody.WriteString("[")
				for _, childItem := range item.SourceZones {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "SecurityZone")
					if itemChildBody != "" {
						if sourceZonesChildBody.Len() > 1 {
							sourceZonesChildBody.WriteString(",")
						}
						sourceZonesChildBody.WriteString(itemChildBody)
					}
				}
				sourceZonesChildBody.WriteString("]")
				itemBody, _ = sjson.SetRaw(itemBody, "sourceZones.objects", sourceZonesChildBody.String())
			}
			if len(item.DestinationZones) > 0 {
				var destinationZonesChildBody strings.Builder
				destinationZonesChildBody.WriteString("[")
				for _, childItem := range item.DestinationZones {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "SecurityZone")
					if itemChildBody != "" {
						if destinationZonesChildBody.Len() > 1 {
							destinationZonesChildBody.WriteString(",")
						}
						destinationZonesChildBody.WriteString(itemChildBody)
					}
				}
				destinationZonesChildBody.WriteString("]")
				itemBody, _ = sjson.SetRaw(itemBody, "destinationZones.objects", destinationZonesChildBody.String())
			}
			if len(item.UrlLiterals) > 0 {
				var urlLiteralsChildBody strings.Builder
				urlLiteralsChildBody.WriteString("[")
				for _, childItem := range item.UrlLiterals {
					itemChildBody := ""
					if !childItem.Url.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "url", childItem.Url.ValueString())
					}
					if itemChildBody != "" {
						if urlLiteralsChildBody.Len() > 1 {
							urlLiteralsChildBody.WriteString(",")
						}
						urlLiteralsChildBody.WriteString(itemChildBody)
					}
				}
				urlLiteralsChildBody.WriteString("]")
				itemBody, _ = sjson.SetRaw(itemBody, "urls.literals", urlLiteralsChildBody.String())
			}
			if len(item.UrlObjects) > 0 {
				var urlObjectsChildBody strings.Builder
				urlObjectsChildBody.WriteString("[")
				for _, childItem := range item.UrlObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "AnyNonEmptyString")
					if itemChildBody != "" {
						if urlObjectsChildBody.Len() > 1 {
							urlObjectsChildBody.WriteString(",")
						}
						urlObjectsChildBody.WriteString(itemChildBody)
					}
				}
				urlObjectsChildBody.WriteString("]")
				itemBody, _ = sjson.SetRaw(itemBody, "urls.objects", urlObjectsChildBody.String())
			}
			if len(item.UrlCategories) > 0 {
				var urlCategoriesChildBody strings.Builder
				urlCategoriesChildBody.WriteString("[")
				for _, childItem := range item.UrlCategories {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "category.id", childItem.Id.ValueString())
					}
					itemChildBody, _ = sjson.Set(itemChildBody, "category.type", "AnyNonEmptyString")
					if !childItem.Reputation.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "reputation", childItem.Reputation.ValueString())
					}
					if itemChildBody != "" {
						if urlCategoriesChildBody.Len() > 1 {
							urlCategoriesChildBody.WriteString(",")
						}
						urlCategoriesChildBody.WriteString(itemChildBody)
					}
				}
				urlCategoriesChildBody.WriteString("]")
				itemBody, _ = sjson.SetRaw(itemBody, "urls.urlCategoriesWithReputation", urlCategoriesChildBody.String())
			}
			if !item.LogConnectionBegin.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "logBegin", item.LogConnectionBegin.ValueBool())
			}
			if !item.LogConnectionEnd.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "logEnd", item.LogConnectionEnd.ValueBool())
			}
			if !item.LogFiles.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "logFiles", item.LogFiles.ValueBool())
			}
			if !item.SendEventsToFmc.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sendEventsToFMC", item.SendEventsToFmc.ValueBool())
			}
			if !item.SendSyslog.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enableSyslog", item.SendSyslog.ValueBool())
			}
			if !item.SyslogAlertId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "syslogConfig.id", item.SyslogAlertId.ValueString())
			}
			if !item.SyslogSeverity.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "syslogSeverity", item.SyslogSeverity.ValueString())
			}
			if !item.SnmpAlertId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "snmpConfig.id", item.SnmpAlertId.ValueString())
			}
			if !item.Description.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "description", item.Description.ValueString())
			}
			if !item.FilePolicyId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "filePolicy.id", item.FilePolicyId.ValueString())
			}
			if !item.IntrusionPolicyId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipsPolicy.id", item.IntrusionPolicyId.ValueString())
			}
			if !item.TimeRangeId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "timeRangeObjects.0.id", item.TimeRangeId.ValueString())
			}
			if !item.VariableSetId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "variableSet.id", item.VariableSetId.ValueString())
			}
			if len(item.Applications) > 0 {
				var applicationsChildBody strings.Builder
				applicationsChildBody.WriteString("[")
				for _, childItem := range item.Applications {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if itemChildBody != "" {
						if applicationsChildBody.Len() > 1 {
							applicationsChildBody.WriteString(",")
						}
						applicationsChildBody.WriteString(itemChildBody)
					}
				}
				applicationsChildBody.WriteString("]")
				itemBody, _ = sjson.SetRaw(itemBody, "applications.applications", applicationsChildBody.String())
			}
			if len(item.ApplicationFilterObjects) > 0 {
				var applicationFilterObjectsChildBody strings.Builder
				applicationFilterObjectsChildBody.WriteString("[")
				for _, childItem := range item.ApplicationFilterObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if itemChildBody != "" {
						if applicationFilterObjectsChildBody.Len() > 1 {
							applicationFilterObjectsChildBody.WriteString(",")
						}
						applicationFilterObjectsChildBody.WriteString(itemChildBody)
					}
				}
				applicationFilterObjectsChildBody.WriteString("]")
				itemBody, _ = sjson.SetRaw(itemBody, "applications.applicationFilters", applicationFilterObjectsChildBody.String())
			}
			if len(item.ApplicationFilters) > 0 {
				var applicationFiltersChildBody strings.Builder
				applicationFiltersChildBody.WriteString("[")
				for _, childItem := range item.ApplicationFilters {
					itemChildBody := ""
					if len(childItem.Types) > 0 {
						var typesChildChildBody strings.Builder
						typesChildChildBody.WriteString("[")
						for _, childChildItem := range childItem.Types {
							itemChildChildBody := ""
							if !childChildItem.Id.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "id", childChildItem.Id.ValueString())
							}
							if itemChildChildBody != "" {
								if typesChildChildBody.Len() > 1 {
									typesChildChildBody.WriteString(",")
								}
								typesChildChildBody.WriteString(itemChildChildBody)
							}
						}
						typesChildChildBody.WriteString("]")
						itemChildBody, _ = sjson.SetRaw(itemChildBody, "applicationTypes", typesChildChildBody.String())
					}
					if len(childItem.Risks) > 0 {
						var risksChildChildBody strings.Builder
						risksChildChildBody.WriteString("[")
						for _, childChildItem := range childItem.Risks {
							itemChildChildBody := ""
							if !childChildItem.Id.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "id", childChildItem.Id.ValueString())
							}
							if itemChildChildBody != "" {
								if risksChildChildBody.Len() > 1 {
									risksChildChildBody.WriteString(",")
								}
								risksChildChildBody.WriteString(itemChildChildBody)
							}
						}
						risksChildChildBody.WriteString("]")
						itemChildBody, _ = sjson.SetRaw(itemChildBody, "risks", risksChildChildBody.String())
					}
					if len(childItem.BusinessRelevances) > 0 {
						var businessRelevancesChildChildBody strings.Builder
						businessRelevancesChildChildBody.WriteString("[")
						for _, childChildItem := range childItem.BusinessRelevances {
							itemChildChildBody := ""
							if !childChildItem.Id.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "id", childChildItem.Id.ValueString())
							}
							if itemChildChildBody != "" {
								if businessRelevancesChildChildBody.Len() > 1 {
									businessRelevancesChildChildBody.WriteString(",")
								}
								businessRelevancesChildChildBody.WriteString(itemChildChildBody)
							}
						}
						businessRelevancesChildChildBody.WriteString("]")
						itemChildBody, _ = sjson.SetRaw(itemChildBody, "productivities", businessRelevancesChildChildBody.String())
					}
					if len(childItem.Categories) > 0 {
						var categoriesChildChildBody strings.Builder
						categoriesChildChildBody.WriteString("[")
						for _, childChildItem := range childItem.Categories {
							itemChildChildBody := ""
							if !childChildItem.Id.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "id", childChildItem.Id.ValueString())
							}
							if itemChildChildBody != "" {
								if categoriesChildChildBody.Len() > 1 {
									categoriesChildChildBody.WriteString(",")
								}
								categoriesChildChildBody.WriteString(itemChildChildBody)
							}
						}
						categoriesChildChildBody.WriteString("]")
						itemChildBody, _ = sjson.SetRaw(itemChildBody, "categories", categoriesChildChildBody.String())
					}
					if len(childItem.Tags) > 0 {
						var tagsChildChildBody strings.Builder
						tagsChildChildBody.WriteString("[")
						for _, childChildItem := range childItem.Tags {
							itemChildChildBody := ""
							if !childChildItem.Id.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "id", childChildItem.Id.ValueString())
							}
							if itemChildChildBody != "" {
								if tagsChildChildBody.Len() > 1 {
									tagsChildChildBody.WriteString(",")
								}
								tagsChildChildBody.WriteString(itemChildChildBody)
							}
						}
						tagsChildChildBody.WriteString("]")
						itemChildBody, _ = sjson.SetRaw(itemChildBody, "tags", tagsChildChildBody.String())
					}
					if itemChildBody != "" {
						if applicationFiltersChildBody.Len() > 1 {
							applicationFiltersChildBody.WriteString(",")
						}
						applicationFiltersChildBody.WriteString(itemChildBody)
					}
				}
				applicationFiltersChildBody.WriteString("]")
				itemBody, _ = sjson.SetRaw(itemBody, "applications.inlineApplicationFilters", applicationFiltersChildBody.String())
			}
			if len(item.UserObjects) > 0 {
				var userObjectsChildBody strings.Builder
				userObjectsChildBody.WriteString("[")
				for _, childItem := range item.UserObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					if !childItem.Name.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "name", childItem.Name.ValueString())
					}
					if !childItem.RealmId.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "realm.id", childItem.RealmId.ValueString())
					}
					if !childItem.RealmType.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "realm.type", childItem.RealmType.ValueString())
					}
					if !childItem.RealmName.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "realm.name", childItem.RealmName.ValueString())
					}
					if itemChildBody != "" {
						if userObjectsChildBody.Len() > 1 {
							userObjectsChildBody.WriteString(",")
						}
						userObjectsChildBody.WriteString(itemChildBody)
					}
				}
				userObjectsChildBody.WriteString("]")
				itemBody, _ = sjson.SetRaw(itemBody, "users.objects", userObjectsChildBody.String())
			}
			if itemBody != "" {
				if itemsBody.Len() > 1 {
					itemsBody.WriteString(",")
				}
				itemsBody.WriteString(itemBody)
			}
		}
		itemsBody.WriteString("]")
		body, _ = sjson.SetRaw(body, "items", itemsBody.String())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *AccessRules) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("category"); value.Exists() {
		data.CategoryName = types.StringValue(value.String())
	} else {
		data.CategoryName = types.StringNull()
	}
	if value := res.Get("section"); value.Exists() {
		data.Section = types.StringValue(value.String())
	} else {
		data.Section = types.StringNull()
	}
	if value := res.Get("items"); value.Exists() {
		data.Items = make([]AccessRulesItems, 0, int(value.Get("#").Int()))
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := AccessRulesItems{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
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
				data.SourceNetworkLiterals = make([]AccessRulesItemsSourceNetworkLiterals, 0, int(value.Get("#").Int()))
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := AccessRulesItemsSourceNetworkLiterals{}
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
				data.DestinationNetworkLiterals = make([]AccessRulesItemsDestinationNetworkLiterals, 0, int(value.Get("#").Int()))
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := AccessRulesItemsDestinationNetworkLiterals{}
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
				data.VlanTagLiterals = make([]AccessRulesItemsVlanTagLiterals, 0, int(value.Get("#").Int()))
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := AccessRulesItemsVlanTagLiterals{}
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
				data.VlanTagObjects = make([]AccessRulesItemsVlanTagObjects, 0, int(value.Get("#").Int()))
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := AccessRulesItemsVlanTagObjects{}
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
				data.SourceNetworkObjects = make([]AccessRulesItemsSourceNetworkObjects, 0, int(value.Get("#").Int()))
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := AccessRulesItemsSourceNetworkObjects{}
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
				data.DestinationNetworkObjects = make([]AccessRulesItemsDestinationNetworkObjects, 0, int(value.Get("#").Int()))
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := AccessRulesItemsDestinationNetworkObjects{}
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
				data.SourceDynamicObjects = make([]AccessRulesItemsSourceDynamicObjects, 0, int(value.Get("#").Int()))
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := AccessRulesItemsSourceDynamicObjects{}
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
				data.DestinationDynamicObjects = make([]AccessRulesItemsDestinationDynamicObjects, 0, int(value.Get("#").Int()))
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := AccessRulesItemsDestinationDynamicObjects{}
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
				data.SourcePortLiterals = make([]AccessRulesItemsSourcePortLiterals, 0, int(value.Get("#").Int()))
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := AccessRulesItemsSourcePortLiterals{}
					if value := res.Get("protocol"); value.Exists() {
						data.Protocol = types.StringValue(value.String())
					} else {
						data.Protocol = types.StringNull()
					}
					if value := res.Get("port"); value.Exists() {
						data.Port = types.StringValue(value.String())
					} else {
						data.Port = types.StringNull()
					}
					(*parent).SourcePortLiterals = append((*parent).SourcePortLiterals, data)
					return true
				})
			}
			if value := res.Get("sourcePorts.objects"); value.Exists() {
				data.SourcePortObjects = make([]AccessRulesItemsSourcePortObjects, 0, int(value.Get("#").Int()))
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := AccessRulesItemsSourcePortObjects{}
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
				data.DestinationPortLiterals = make([]AccessRulesItemsDestinationPortLiterals, 0, int(value.Get("#").Int()))
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := AccessRulesItemsDestinationPortLiterals{}
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
				data.DestinationPortObjects = make([]AccessRulesItemsDestinationPortObjects, 0, int(value.Get("#").Int()))
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := AccessRulesItemsDestinationPortObjects{}
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
				data.SourceSgtObjects = make([]AccessRulesItemsSourceSgtObjects, 0, int(value.Get("#").Int()))
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := AccessRulesItemsSourceSgtObjects{}
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
				data.EndpointDeviceTypes = make([]AccessRulesItemsEndpointDeviceTypes, 0, int(value.Get("#").Int()))
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := AccessRulesItemsEndpointDeviceTypes{}
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
				data.DestinationSgtObjects = make([]AccessRulesItemsDestinationSgtObjects, 0, int(value.Get("#").Int()))
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := AccessRulesItemsDestinationSgtObjects{}
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
				data.SourceZones = make([]AccessRulesItemsSourceZones, 0, int(value.Get("#").Int()))
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := AccessRulesItemsSourceZones{}
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
				data.DestinationZones = make([]AccessRulesItemsDestinationZones, 0, int(value.Get("#").Int()))
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := AccessRulesItemsDestinationZones{}
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
				data.UrlLiterals = make([]AccessRulesItemsUrlLiterals, 0, int(value.Get("#").Int()))
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := AccessRulesItemsUrlLiterals{}
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
				data.UrlObjects = make([]AccessRulesItemsUrlObjects, 0, int(value.Get("#").Int()))
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := AccessRulesItemsUrlObjects{}
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
				data.UrlCategories = make([]AccessRulesItemsUrlCategories, 0, int(value.Get("#").Int()))
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := AccessRulesItemsUrlCategories{}
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
				data.LogConnectionBegin = types.BoolValue(value.Bool())
			} else {
				data.LogConnectionBegin = types.BoolValue(false)
			}
			if value := res.Get("logEnd"); value.Exists() {
				data.LogConnectionEnd = types.BoolValue(value.Bool())
			} else {
				data.LogConnectionEnd = types.BoolValue(false)
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
				data.SyslogAlertId = types.StringValue(value.String())
			} else {
				data.SyslogAlertId = types.StringNull()
			}
			if value := res.Get("syslogSeverity"); value.Exists() {
				data.SyslogSeverity = types.StringValue(value.String())
			} else {
				data.SyslogSeverity = types.StringNull()
			}
			if value := res.Get("snmpConfig.id"); value.Exists() {
				data.SnmpAlertId = types.StringValue(value.String())
			} else {
				data.SnmpAlertId = types.StringNull()
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
				data.Applications = make([]AccessRulesItemsApplications, 0, int(value.Get("#").Int()))
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := AccessRulesItemsApplications{}
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
				data.ApplicationFilterObjects = make([]AccessRulesItemsApplicationFilterObjects, 0, int(value.Get("#").Int()))
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := AccessRulesItemsApplicationFilterObjects{}
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
				data.ApplicationFilters = make([]AccessRulesItemsApplicationFilters, 0, int(value.Get("#").Int()))
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := AccessRulesItemsApplicationFilters{}
					if value := res.Get("applicationTypes"); value.Exists() {
						data.Types = make([]AccessRulesItemsApplicationFiltersTypes, 0, int(value.Get("#").Int()))
						value.ForEach(func(k, res gjson.Result) bool {
							parent := &data
							data := AccessRulesItemsApplicationFiltersTypes{}
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
						data.Risks = make([]AccessRulesItemsApplicationFiltersRisks, 0, int(value.Get("#").Int()))
						value.ForEach(func(k, res gjson.Result) bool {
							parent := &data
							data := AccessRulesItemsApplicationFiltersRisks{}
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
						data.BusinessRelevances = make([]AccessRulesItemsApplicationFiltersBusinessRelevances, 0, int(value.Get("#").Int()))
						value.ForEach(func(k, res gjson.Result) bool {
							parent := &data
							data := AccessRulesItemsApplicationFiltersBusinessRelevances{}
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
						data.Categories = make([]AccessRulesItemsApplicationFiltersCategories, 0, int(value.Get("#").Int()))
						value.ForEach(func(k, res gjson.Result) bool {
							parent := &data
							data := AccessRulesItemsApplicationFiltersCategories{}
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
						data.Tags = make([]AccessRulesItemsApplicationFiltersTags, 0, int(value.Get("#").Int()))
						value.ForEach(func(k, res gjson.Result) bool {
							parent := &data
							data := AccessRulesItemsApplicationFiltersTags{}
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
			if value := res.Get("users.objects"); value.Exists() {
				data.UserObjects = make([]AccessRulesItemsUserObjects, 0, int(value.Get("#").Int()))
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := AccessRulesItemsUserObjects{}
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
					if value := res.Get("name"); value.Exists() {
						data.Name = types.StringValue(value.String())
					} else {
						data.Name = types.StringNull()
					}
					if value := res.Get("realm.id"); value.Exists() {
						data.RealmId = types.StringValue(value.String())
					} else {
						data.RealmId = types.StringNull()
					}
					if value := res.Get("realm.type"); value.Exists() {
						data.RealmType = types.StringValue(value.String())
					} else {
						data.RealmType = types.StringNull()
					}
					if value := res.Get("realm.name"); value.Exists() {
						data.RealmName = types.StringValue(value.String())
					} else {
						data.RealmName = types.StringNull()
					}
					(*parent).UserObjects = append((*parent).UserObjects, data)
					return true
				})
			}
			(*parent).Items = append((*parent).Items, data)
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
func (data *AccessRules) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("category"); value.Exists() && !data.CategoryName.IsNull() {
		data.CategoryName = types.StringValue(value.String())
	} else {
		data.CategoryName = types.StringNull()
	}
	if value := res.Get("section"); value.Exists() && !data.Section.IsNull() {
		data.Section = types.StringValue(value.String())
	} else {
		data.Section = types.StringNull()
	}
	itemsArray := res.Get("items").Array()
	{
		l := len(itemsArray)
		tflog.Debug(ctx, fmt.Sprintf("items array resizing from %d to %d", len(data.Items), l))
		for i := len(data.Items); i < l; i++ {
			data.Items = append(data.Items, AccessRulesItems{})
		}
		if len(data.Items) > l {
			data.Items = data.Items[:l]
		}
	}
	for i := range data.Items {
		parent := &data
		data := (*parent).Items[i]
		res := itemsArray[i]
		if value := res.Get("id"); value.Exists() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
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
		sourceNetworkLiteralsArray := res.Get("sourceNetworks.literals")
		for i := 0; i < len(data.SourceNetworkLiterals); i++ {
			keys := [...]string{"value"}
			keyValues := [...]string{data.SourceNetworkLiterals[i].Value.ValueString()}

			parent := &data
			data := (*parent).SourceNetworkLiterals[i]
			var res gjson.Result

			sourceNetworkLiteralsArray.ForEach(
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
		destinationNetworkLiteralsArray := res.Get("destinationNetworks.literals")
		for i := 0; i < len(data.DestinationNetworkLiterals); i++ {
			keys := [...]string{"value"}
			keyValues := [...]string{data.DestinationNetworkLiterals[i].Value.ValueString()}

			parent := &data
			data := (*parent).DestinationNetworkLiterals[i]
			var res gjson.Result

			destinationNetworkLiteralsArray.ForEach(
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
		vlanTagLiteralsArray := res.Get("vlanTags.literals")
		for i := 0; i < len(data.VlanTagLiterals); i++ {
			keys := [...]string{"startTag", "endTag"}
			keyValues := [...]string{data.VlanTagLiterals[i].StartTag.ValueString(), data.VlanTagLiterals[i].EndTag.ValueString()}

			parent := &data
			data := (*parent).VlanTagLiterals[i]
			var res gjson.Result

			vlanTagLiteralsArray.ForEach(
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
		vlanTagObjectsArray := res.Get("vlanTags.objects")
		for i := 0; i < len(data.VlanTagObjects); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.VlanTagObjects[i].Id.ValueString()}

			parent := &data
			data := (*parent).VlanTagObjects[i]
			var res gjson.Result

			vlanTagObjectsArray.ForEach(
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
		sourceNetworkObjectsArray := res.Get("sourceNetworks.objects")
		for i := 0; i < len(data.SourceNetworkObjects); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.SourceNetworkObjects[i].Id.ValueString()}

			parent := &data
			data := (*parent).SourceNetworkObjects[i]
			var res gjson.Result

			sourceNetworkObjectsArray.ForEach(
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
		destinationNetworkObjectsArray := res.Get("destinationNetworks.objects")
		for i := 0; i < len(data.DestinationNetworkObjects); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.DestinationNetworkObjects[i].Id.ValueString()}

			parent := &data
			data := (*parent).DestinationNetworkObjects[i]
			var res gjson.Result

			destinationNetworkObjectsArray.ForEach(
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
		sourceDynamicObjectsArray := res.Get("sourceDynamicObjects.objects")
		for i := 0; i < len(data.SourceDynamicObjects); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.SourceDynamicObjects[i].Id.ValueString()}

			parent := &data
			data := (*parent).SourceDynamicObjects[i]
			var res gjson.Result

			sourceDynamicObjectsArray.ForEach(
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
		destinationDynamicObjectsArray := res.Get("destinationDynamicObjects.objects")
		for i := 0; i < len(data.DestinationDynamicObjects); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.DestinationDynamicObjects[i].Id.ValueString()}

			parent := &data
			data := (*parent).DestinationDynamicObjects[i]
			var res gjson.Result

			destinationDynamicObjectsArray.ForEach(
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
		sourcePortLiteralsArray := res.Get("sourcePorts.literals")
		for i := 0; i < len(data.SourcePortLiterals); i++ {
			keys := [...]string{"protocol", "port"}
			keyValues := [...]string{data.SourcePortLiterals[i].Protocol.ValueString(), data.SourcePortLiterals[i].Port.ValueString()}

			parent := &data
			data := (*parent).SourcePortLiterals[i]
			var res gjson.Result

			sourcePortLiteralsArray.ForEach(
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
			if value := res.Get("protocol"); value.Exists() && !data.Protocol.IsNull() {
				data.Protocol = types.StringValue(value.String())
			} else {
				data.Protocol = types.StringNull()
			}
			if value := res.Get("port"); value.Exists() && !data.Port.IsNull() {
				data.Port = types.StringValue(value.String())
			} else {
				data.Port = types.StringNull()
			}
			(*parent).SourcePortLiterals[i] = data
		}
		sourcePortObjectsArray := res.Get("sourcePorts.objects")
		for i := 0; i < len(data.SourcePortObjects); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.SourcePortObjects[i].Id.ValueString()}

			parent := &data
			data := (*parent).SourcePortObjects[i]
			var res gjson.Result

			sourcePortObjectsArray.ForEach(
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
		destinationPortLiteralsArray := res.Get("destinationPorts.literals")
		for i := 0; i < len(data.DestinationPortLiterals); i++ {
			keys := [...]string{"type", "port", "protocol", "icmpType", "code"}
			keyValues := [...]string{data.DestinationPortLiterals[i].Type.ValueString(), data.DestinationPortLiterals[i].Port.ValueString(), data.DestinationPortLiterals[i].Protocol.ValueString(), data.DestinationPortLiterals[i].IcmpType.ValueString(), data.DestinationPortLiterals[i].IcmpCode.ValueString()}

			parent := &data
			data := (*parent).DestinationPortLiterals[i]
			var res gjson.Result

			destinationPortLiteralsArray.ForEach(
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
		destinationPortObjectsArray := res.Get("destinationPorts.objects")
		for i := 0; i < len(data.DestinationPortObjects); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.DestinationPortObjects[i].Id.ValueString()}

			parent := &data
			data := (*parent).DestinationPortObjects[i]
			var res gjson.Result

			destinationPortObjectsArray.ForEach(
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
		sourceSgtObjectsArray := res.Get("sourceSecurityGroupTags.objects")
		for i := 0; i < len(data.SourceSgtObjects); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.SourceSgtObjects[i].Id.ValueString()}

			parent := &data
			data := (*parent).SourceSgtObjects[i]
			var res gjson.Result

			sourceSgtObjectsArray.ForEach(
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
		endpointDeviceTypesArray := res.Get("endPointDeviceTypes")
		for i := 0; i < len(data.EndpointDeviceTypes); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.EndpointDeviceTypes[i].Id.ValueString()}

			parent := &data
			data := (*parent).EndpointDeviceTypes[i]
			var res gjson.Result

			endpointDeviceTypesArray.ForEach(
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
		destinationSgtObjectsArray := res.Get("destinationSecurityGroupTags.objects")
		for i := 0; i < len(data.DestinationSgtObjects); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.DestinationSgtObjects[i].Id.ValueString()}

			parent := &data
			data := (*parent).DestinationSgtObjects[i]
			var res gjson.Result

			destinationSgtObjectsArray.ForEach(
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
		sourceZonesArray := res.Get("sourceZones.objects")
		for i := 0; i < len(data.SourceZones); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.SourceZones[i].Id.ValueString()}

			parent := &data
			data := (*parent).SourceZones[i]
			var res gjson.Result

			sourceZonesArray.ForEach(
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
		destinationZonesArray := res.Get("destinationZones.objects")
		for i := 0; i < len(data.DestinationZones); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.DestinationZones[i].Id.ValueString()}

			parent := &data
			data := (*parent).DestinationZones[i]
			var res gjson.Result

			destinationZonesArray.ForEach(
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
		urlLiteralsArray := res.Get("urls.literals")
		for i := 0; i < len(data.UrlLiterals); i++ {
			keys := [...]string{"url"}
			keyValues := [...]string{data.UrlLiterals[i].Url.ValueString()}

			parent := &data
			data := (*parent).UrlLiterals[i]
			var res gjson.Result

			urlLiteralsArray.ForEach(
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
		urlObjectsArray := res.Get("urls.objects")
		for i := 0; i < len(data.UrlObjects); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.UrlObjects[i].Id.ValueString()}

			parent := &data
			data := (*parent).UrlObjects[i]
			var res gjson.Result

			urlObjectsArray.ForEach(
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
		urlCategoriesArray := res.Get("urls.urlCategoriesWithReputation")
		for i := 0; i < len(data.UrlCategories); i++ {
			keys := [...]string{"category.id"}
			keyValues := [...]string{data.UrlCategories[i].Id.ValueString()}

			parent := &data
			data := (*parent).UrlCategories[i]
			var res gjson.Result

			urlCategoriesArray.ForEach(
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
		if value := res.Get("logBegin"); value.Exists() && !data.LogConnectionBegin.IsNull() {
			data.LogConnectionBegin = types.BoolValue(value.Bool())
		} else if data.LogConnectionBegin.ValueBool() != false {
			data.LogConnectionBegin = types.BoolNull()
		}
		if value := res.Get("logEnd"); value.Exists() && !data.LogConnectionEnd.IsNull() {
			data.LogConnectionEnd = types.BoolValue(value.Bool())
		} else if data.LogConnectionEnd.ValueBool() != false {
			data.LogConnectionEnd = types.BoolNull()
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
		if value := res.Get("syslogConfig.id"); value.Exists() && !data.SyslogAlertId.IsNull() {
			data.SyslogAlertId = types.StringValue(value.String())
		} else {
			data.SyslogAlertId = types.StringNull()
		}
		if value := res.Get("syslogSeverity"); value.Exists() && !data.SyslogSeverity.IsNull() {
			data.SyslogSeverity = types.StringValue(value.String())
		} else {
			data.SyslogSeverity = types.StringNull()
		}
		if value := res.Get("snmpConfig.id"); value.Exists() && !data.SnmpAlertId.IsNull() {
			data.SnmpAlertId = types.StringValue(value.String())
		} else {
			data.SnmpAlertId = types.StringNull()
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
		applicationsArray := res.Get("applications.applications")
		for i := 0; i < len(data.Applications); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.Applications[i].Id.ValueString()}

			parent := &data
			data := (*parent).Applications[i]
			var res gjson.Result

			applicationsArray.ForEach(
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
		applicationFilterObjectsArray := res.Get("applications.applicationFilters")
		for i := 0; i < len(data.ApplicationFilterObjects); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.ApplicationFilterObjects[i].Id.ValueString()}

			parent := &data
			data := (*parent).ApplicationFilterObjects[i]
			var res gjson.Result

			applicationFilterObjectsArray.ForEach(
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
		applicationFiltersArray := res.Get("applications.inlineApplicationFilters").Array()
		{
			l := len(applicationFiltersArray)
			tflog.Debug(ctx, fmt.Sprintf("applications.inlineApplicationFilters array resizing from %d to %d", len(data.ApplicationFilters), l))
			for i := len(data.ApplicationFilters); i < l; i++ {
				data.ApplicationFilters = append(data.ApplicationFilters, AccessRulesItemsApplicationFilters{})
			}
			if len(data.ApplicationFilters) > l {
				data.ApplicationFilters = data.ApplicationFilters[:l]
			}
		}
		for i := range data.ApplicationFilters {
			parent := &data
			data := (*parent).ApplicationFilters[i]
			res := applicationFiltersArray[i]
			typesArray := res.Get("applicationTypes")
			for i := 0; i < len(data.Types); i++ {
				keys := [...]string{"id"}
				keyValues := [...]string{data.Types[i].Id.ValueString()}

				parent := &data
				data := (*parent).Types[i]
				var res gjson.Result

				typesArray.ForEach(
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
			risksArray := res.Get("risks")
			for i := 0; i < len(data.Risks); i++ {
				keys := [...]string{"id"}
				keyValues := [...]string{data.Risks[i].Id.ValueString()}

				parent := &data
				data := (*parent).Risks[i]
				var res gjson.Result

				risksArray.ForEach(
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
			businessRelevancesArray := res.Get("productivities")
			for i := 0; i < len(data.BusinessRelevances); i++ {
				keys := [...]string{"id"}
				keyValues := [...]string{data.BusinessRelevances[i].Id.ValueString()}

				parent := &data
				data := (*parent).BusinessRelevances[i]
				var res gjson.Result

				businessRelevancesArray.ForEach(
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
			categoriesArray := res.Get("categories")
			for i := 0; i < len(data.Categories); i++ {
				keys := [...]string{"id"}
				keyValues := [...]string{data.Categories[i].Id.ValueString()}

				parent := &data
				data := (*parent).Categories[i]
				var res gjson.Result

				categoriesArray.ForEach(
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
			tagsArray := res.Get("tags")
			for i := 0; i < len(data.Tags); i++ {
				keys := [...]string{"id"}
				keyValues := [...]string{data.Tags[i].Id.ValueString()}

				parent := &data
				data := (*parent).Tags[i]
				var res gjson.Result

				tagsArray.ForEach(
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
		userObjectsArray := res.Get("users.objects")
		for i := 0; i < len(data.UserObjects); i++ {
			keys := [...]string{"id", "realm.id"}
			keyValues := [...]string{data.UserObjects[i].Id.ValueString(), data.UserObjects[i].RealmId.ValueString()}

			parent := &data
			data := (*parent).UserObjects[i]
			var res gjson.Result

			userObjectsArray.ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing UserObjects[%d] = %+v",
					i,
					(*parent).UserObjects[i],
				))
				(*parent).UserObjects = slices.Delete((*parent).UserObjects, i, i+1)
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
			if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			if value := res.Get("realm.id"); value.Exists() && !data.RealmId.IsNull() {
				data.RealmId = types.StringValue(value.String())
			} else {
				data.RealmId = types.StringNull()
			}
			if value := res.Get("realm.type"); value.Exists() && !data.RealmType.IsNull() {
				data.RealmType = types.StringValue(value.String())
			} else {
				data.RealmType = types.StringNull()
			}
			if value := res.Get("realm.name"); value.Exists() && !data.RealmName.IsNull() {
				data.RealmName = types.StringValue(value.String())
			} else {
				data.RealmName = types.StringNull()
			}
			(*parent).UserObjects[i] = data
		}
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *AccessRules) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	itemsArray := res.Get("items").Array()
	for i := range data.Items {
		var r gjson.Result
		if i < len(itemsArray) {
			r = itemsArray[i]
		}
		if v := data.Items[i]; v.Id.IsUnknown() {
			if value := r.Get("id"); value.Exists() {
				v.Id = types.StringValue(value.String())
			} else {
				v.Id = types.StringNull()
			}
			data.Items[i] = v
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin Clone

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// End of section. //template:end toBodyNonBulk
