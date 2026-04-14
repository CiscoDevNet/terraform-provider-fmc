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
	"slices"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type IdentityPolicy struct {
	Id                                        types.String               `tfsdk:"id"`
	Domain                                    types.String               `tfsdk:"domain"`
	Name                                      types.String               `tfsdk:"name"`
	Type                                      types.String               `tfsdk:"type"`
	Description                               types.String               `tfsdk:"description"`
	IdentityMappingFilterNetworkObjectId      types.String               `tfsdk:"identity_mapping_filter_network_object_id"`
	IdentityMappingFilterNetworkObjectType    types.String               `tfsdk:"identity_mapping_filter_network_object_type"`
	ActiveAuthenticationServerCertificateId   types.String               `tfsdk:"active_authentication_server_certificate_id"`
	ActiveAuthenticationServerCertificateType types.String               `tfsdk:"active_authentication_server_certificate_type"`
	ActiveAuthenticationRedirectFqdnId        types.String               `tfsdk:"active_authentication_redirect_fqdn_id"`
	ActiveAuthenticationRedirectFqdnType      types.String               `tfsdk:"active_authentication_redirect_fqdn_type"`
	ActiveAuthenticationRedirectPort          types.Int64                `tfsdk:"active_authentication_redirect_port"`
	ActiveAuthenticationMaximumLoginAttempts  types.Int64                `tfsdk:"active_authentication_maximum_login_attempts"`
	ActiveAuthenticationSessionSharing        types.Bool                 `tfsdk:"active_authentication_session_sharing"`
	ActiveAuthenticationPage                  types.String               `tfsdk:"active_authentication_page"`
	ActiveAuthenticationPageHtml              types.String               `tfsdk:"active_authentication_page_html"`
	Categories                                []IdentityPolicyCategories `tfsdk:"categories"`
	Rules                                     []IdentityPolicyRules      `tfsdk:"rules"`
}

type IdentityPolicyCategories struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type IdentityPolicyRules struct {
	Id                           types.String                                    `tfsdk:"id"`
	Name                         types.String                                    `tfsdk:"name"`
	Enabled                      types.Bool                                      `tfsdk:"enabled"`
	Category                     types.String                                    `tfsdk:"category"`
	AuthenticationType           types.String                                    `tfsdk:"authentication_type"`
	AuthenticationRealmId        types.String                                    `tfsdk:"authentication_realm_id"`
	AuthenticationRealmType      types.String                                    `tfsdk:"authentication_realm_type"`
	GuestAccessFallback          types.Bool                                      `tfsdk:"guest_access_fallback"`
	ActiveAuthenticationFallback types.Bool                                      `tfsdk:"active_authentication_fallback"`
	AuthenticationProtocol       types.String                                    `tfsdk:"authentication_protocol"`
	ExcludedUserAgents           []IdentityPolicyRulesExcludedUserAgents         `tfsdk:"excluded_user_agents"`
	SourceZones                  []IdentityPolicyRulesSourceZones                `tfsdk:"source_zones"`
	DestinationZones             []IdentityPolicyRulesDestinationZones           `tfsdk:"destination_zones"`
	SourceNetworkLiterals        []IdentityPolicyRulesSourceNetworkLiterals      `tfsdk:"source_network_literals"`
	SourceNetworkObjects         []IdentityPolicyRulesSourceNetworkObjects       `tfsdk:"source_network_objects"`
	DestinationNetworkLiterals   []IdentityPolicyRulesDestinationNetworkLiterals `tfsdk:"destination_network_literals"`
	DestinationNetworkObjects    []IdentityPolicyRulesDestinationNetworkObjects  `tfsdk:"destination_network_objects"`
	VlanTagLiterals              []IdentityPolicyRulesVlanTagLiterals            `tfsdk:"vlan_tag_literals"`
	VlanTagObjects               []IdentityPolicyRulesVlanTagObjects             `tfsdk:"vlan_tag_objects"`
	SourcePortLiterals           []IdentityPolicyRulesSourcePortLiterals         `tfsdk:"source_port_literals"`
	SourcePortObjects            []IdentityPolicyRulesSourcePortObjects          `tfsdk:"source_port_objects"`
	DestinationPortLiterals      []IdentityPolicyRulesDestinationPortLiterals    `tfsdk:"destination_port_literals"`
	DestinationPortObjects       []IdentityPolicyRulesDestinationPortObjects     `tfsdk:"destination_port_objects"`
}

type IdentityPolicyRulesExcludedUserAgents struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}
type IdentityPolicyRulesSourceZones struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}
type IdentityPolicyRulesDestinationZones struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}
type IdentityPolicyRulesSourceNetworkLiterals struct {
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}
type IdentityPolicyRulesSourceNetworkObjects struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}
type IdentityPolicyRulesDestinationNetworkLiterals struct {
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}
type IdentityPolicyRulesDestinationNetworkObjects struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}
type IdentityPolicyRulesVlanTagLiterals struct {
	Type     types.String `tfsdk:"type"`
	StartTag types.String `tfsdk:"start_tag"`
	EndTag   types.String `tfsdk:"end_tag"`
}
type IdentityPolicyRulesVlanTagObjects struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}
type IdentityPolicyRulesSourcePortLiterals struct {
	Type     types.String `tfsdk:"type"`
	Protocol types.String `tfsdk:"protocol"`
	Port     types.String `tfsdk:"port"`
}
type IdentityPolicyRulesSourcePortObjects struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}
type IdentityPolicyRulesDestinationPortLiterals struct {
	Type     types.String `tfsdk:"type"`
	Port     types.String `tfsdk:"port"`
	Protocol types.String `tfsdk:"protocol"`
}
type IdentityPolicyRulesDestinationPortObjects struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionIdentityPolicy = version.Must(version.NewVersion("7.2"))
var minFMCVersionCreateIdentityPolicy = version.Must(version.NewVersion("10.0"))

// End of section. //template:end minimumVersions

var identityPolicyCategoriesApiToGuiNames = map[string]string{
	"admin_category":    "Administrator Rules",
	"standard_category": "Standard Rules",
	"root_category":     "Root Rules",
}

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data IdentityPolicy) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/identitypolicies"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data IdentityPolicy) toBody(ctx context.Context, state IdentityPolicy) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	body, _ = sjson.Set(body, "type", "IdentityPolicy")
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.IdentityMappingFilterNetworkObjectId.IsNull() {
		body, _ = sjson.Set(body, "mappingFilter.id", data.IdentityMappingFilterNetworkObjectId.ValueString())
	}
	if !data.IdentityMappingFilterNetworkObjectType.IsNull() {
		body, _ = sjson.Set(body, "mappingFilter.type", data.IdentityMappingFilterNetworkObjectType.ValueString())
	}
	if !data.ActiveAuthenticationServerCertificateId.IsNull() {
		body, _ = sjson.Set(body, "authCert.id", data.ActiveAuthenticationServerCertificateId.ValueString())
	}
	if !data.ActiveAuthenticationServerCertificateType.IsNull() {
		body, _ = sjson.Set(body, "authCert.type", data.ActiveAuthenticationServerCertificateType.ValueString())
	}
	if !data.ActiveAuthenticationRedirectFqdnId.IsNull() {
		body, _ = sjson.Set(body, "redirectHost.id", data.ActiveAuthenticationRedirectFqdnId.ValueString())
	}
	if !data.ActiveAuthenticationRedirectFqdnType.IsNull() {
		body, _ = sjson.Set(body, "redirectHost.type", data.ActiveAuthenticationRedirectFqdnType.ValueString())
	}
	if !data.ActiveAuthenticationRedirectPort.IsNull() {
		body, _ = sjson.Set(body, "redirectPort", data.ActiveAuthenticationRedirectPort.ValueInt64())
	}
	if !data.ActiveAuthenticationMaximumLoginAttempts.IsNull() {
		body, _ = sjson.Set(body, "maxLoginAttempts", data.ActiveAuthenticationMaximumLoginAttempts.ValueInt64())
	}
	if !data.ActiveAuthenticationSessionSharing.IsNull() {
		body, _ = sjson.Set(body, "activeSessionSharing", data.ActiveAuthenticationSessionSharing.ValueBool())
	}
	if !data.ActiveAuthenticationPage.IsNull() {
		body, _ = sjson.Set(body, "activeAuthPage", data.ActiveAuthenticationPage.ValueString())
	}
	if !data.ActiveAuthenticationPageHtml.IsNull() {
		body, _ = sjson.Set(body, "htmlActiveAuthPage", data.ActiveAuthenticationPageHtml.ValueString())
	}
	if len(data.Categories) > 0 {
		body, _ = sjson.Set(body, "dummy_categories", []any{})
		for _, item := range data.Categories {
			itemBody := ""
			if !item.Id.IsNull() && !item.Id.IsUnknown() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			itemBody, _ = sjson.Set(itemBody, "type", "IdentityPolicyCategory")
			body, _ = sjson.SetRaw(body, "dummy_categories.-1", itemBody)
		}
	}
	if len(data.Rules) > 0 {
		body, _ = sjson.Set(body, "dummy_rules", []any{})
		for _, item := range data.Rules {
			itemBody := ""
			if !item.Id.IsNull() && !item.Id.IsUnknown() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			itemBody, _ = sjson.Set(itemBody, "type", "IdentityPolicyRule")
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.Enabled.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enabled", item.Enabled.ValueBool())
			}
			if !item.AuthenticationType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "authenticationType", item.AuthenticationType.ValueString())
			}
			if !item.AuthenticationRealmId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "realm.id", item.AuthenticationRealmId.ValueString())
			}
			if !item.AuthenticationRealmType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "realm.type", item.AuthenticationRealmType.ValueString())
			}
			if !item.GuestAccessFallback.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "guestAccessFallback", item.GuestAccessFallback.ValueBool())
			}
			if !item.ActiveAuthenticationFallback.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "captivePortalFallback", item.ActiveAuthenticationFallback.ValueBool())
			}
			if !item.AuthenticationProtocol.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "authenticationProtocol", item.AuthenticationProtocol.ValueString())
			}
			if len(item.ExcludedUserAgents) > 0 {
				itemBody, _ = sjson.Set(itemBody, "excludedUserAgent", []any{})
				for _, childItem := range item.ExcludedUserAgents {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "excludedUserAgent.-1", itemChildBody)
				}
			}
			if len(item.SourceZones) > 0 {
				itemBody, _ = sjson.Set(itemBody, "sourceZones", []any{})
				for _, childItem := range item.SourceZones {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "sourceZones.-1", itemChildBody)
				}
			}
			if len(item.DestinationZones) > 0 {
				itemBody, _ = sjson.Set(itemBody, "destinationZones", []any{})
				for _, childItem := range item.DestinationZones {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "destinationZones.-1", itemChildBody)
				}
			}
			if len(item.SourceNetworkLiterals) > 0 {
				itemBody, _ = sjson.Set(itemBody, "sourceNetworks.literals", []any{})
				for _, childItem := range item.SourceNetworkLiterals {
					itemChildBody := ""
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					if !childItem.Value.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Value.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "sourceNetworks.literals.-1", itemChildBody)
				}
			}
			if len(item.SourceNetworkObjects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "sourceNetworks.objects", []any{})
				for _, childItem := range item.SourceNetworkObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "sourceNetworks.objects.-1", itemChildBody)
				}
			}
			if len(item.DestinationNetworkLiterals) > 0 {
				itemBody, _ = sjson.Set(itemBody, "destinationNetworks.literals", []any{})
				for _, childItem := range item.DestinationNetworkLiterals {
					itemChildBody := ""
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					if !childItem.Value.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Value.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "destinationNetworks.literals.-1", itemChildBody)
				}
			}
			if len(item.DestinationNetworkObjects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "destinationNetworks.objects", []any{})
				for _, childItem := range item.DestinationNetworkObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "destinationNetworks.objects.-1", itemChildBody)
				}
			}
			if len(item.VlanTagLiterals) > 0 {
				itemBody, _ = sjson.Set(itemBody, "vlanTags.literals", []any{})
				for _, childItem := range item.VlanTagLiterals {
					itemChildBody := ""
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					if !childItem.StartTag.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "startTag", childItem.StartTag.ValueString())
					}
					if !childItem.EndTag.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "endTag", childItem.EndTag.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "vlanTags.literals.-1", itemChildBody)
				}
			}
			if len(item.VlanTagObjects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "vlanTags.objects", []any{})
				for _, childItem := range item.VlanTagObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "vlanTags.objects.-1", itemChildBody)
				}
			}
			if len(item.SourcePortLiterals) > 0 {
				itemBody, _ = sjson.Set(itemBody, "sourcePorts.literals", []any{})
				for _, childItem := range item.SourcePortLiterals {
					itemChildBody := ""
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					if !childItem.Protocol.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "protocol", childItem.Protocol.ValueString())
					}
					if !childItem.Port.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "port", childItem.Port.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "sourcePorts.literals.-1", itemChildBody)
				}
			}
			if len(item.SourcePortObjects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "sourcePorts.objects", []any{})
				for _, childItem := range item.SourcePortObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "sourcePorts.objects.-1", itemChildBody)
				}
			}
			if len(item.DestinationPortLiterals) > 0 {
				itemBody, _ = sjson.Set(itemBody, "destinationPorts.literals", []any{})
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
					itemBody, _ = sjson.SetRaw(itemBody, "destinationPorts.literals.-1", itemChildBody)
				}
			}
			if len(item.DestinationPortObjects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "destinationPorts.objects", []any{})
				for _, childItem := range item.DestinationPortObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "destinationPorts.objects.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "dummy_rules.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *IdentityPolicy) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("mappingFilter.id"); value.Exists() {
		data.IdentityMappingFilterNetworkObjectId = types.StringValue(value.String())
	} else {
		data.IdentityMappingFilterNetworkObjectId = types.StringNull()
	}
	if value := res.Get("mappingFilter.type"); value.Exists() {
		data.IdentityMappingFilterNetworkObjectType = types.StringValue(value.String())
	} else {
		data.IdentityMappingFilterNetworkObjectType = types.StringNull()
	}
	if value := res.Get("authCert.id"); value.Exists() {
		data.ActiveAuthenticationServerCertificateId = types.StringValue(value.String())
	} else {
		data.ActiveAuthenticationServerCertificateId = types.StringNull()
	}
	if value := res.Get("authCert.type"); value.Exists() {
		data.ActiveAuthenticationServerCertificateType = types.StringValue(value.String())
	} else {
		data.ActiveAuthenticationServerCertificateType = types.StringValue("PKI_InternalCert")
	}
	if value := res.Get("redirectHost.id"); value.Exists() {
		data.ActiveAuthenticationRedirectFqdnId = types.StringValue(value.String())
	} else {
		data.ActiveAuthenticationRedirectFqdnId = types.StringNull()
	}
	if value := res.Get("redirectHost.type"); value.Exists() {
		data.ActiveAuthenticationRedirectFqdnType = types.StringValue(value.String())
	} else {
		data.ActiveAuthenticationRedirectFqdnType = types.StringValue("FQDN")
	}
	if value := res.Get("redirectPort"); value.Exists() {
		data.ActiveAuthenticationRedirectPort = types.Int64Value(value.Int())
	} else {
		data.ActiveAuthenticationRedirectPort = types.Int64Value(885)
	}
	if value := res.Get("maxLoginAttempts"); value.Exists() {
		data.ActiveAuthenticationMaximumLoginAttempts = types.Int64Value(value.Int())
	} else {
		data.ActiveAuthenticationMaximumLoginAttempts = types.Int64Value(3)
	}
	if value := res.Get("activeSessionSharing"); value.Exists() {
		data.ActiveAuthenticationSessionSharing = types.BoolValue(value.Bool())
	} else {
		data.ActiveAuthenticationSessionSharing = types.BoolValue(true)
	}
	if value := res.Get("activeAuthPage"); value.Exists() {
		data.ActiveAuthenticationPage = types.StringValue(value.String())
	} else {
		data.ActiveAuthenticationPage = types.StringValue("DEFAULT")
	}
	if value := res.Get("htmlActiveAuthPage"); value.Exists() {
		data.ActiveAuthenticationPageHtml = types.StringValue(value.String())
	} else {
		data.ActiveAuthenticationPageHtml = types.StringValue("<!DOCTYPE html>\n<html>\n<head>\n<meta http-equiv=\"content-type\" content=\"text/html; charset=UTF-8\" />\n<title>Login</title>\n<style type=\"text/css\">\nbody {\n    margin:0;\n    font-family:verdana,sans-serif;\n}\nh1 {\n    margin:0;\n    padding:12px 25px;\n    background-color:#343434;\n    color:#ddd;\n}\np {\n    margin:12px 25px;\n}\nstrong {\n    color:#E0042D;\n}\ndiv {\n    padding-left:23px;\n    font-weight: normal;\n    font-size: 8pt;\n}\ninput {\n    margin:12px 25px;\n}\n</style>\n</head>\n<body>\n    <form action=\"\" id=\"loginForm\" method=\"post\" name=\"loginForm\">\n        <h1>Login</h1>\n        <p><strong>Please enter your username and password or log in as a guest.</strong></p>\n        <div class=\"label\">Username\n            <input id=\"name\" maxlength=\"100\" name=\"login\" type=\"text\" value=\"\"/>\n            <b>realm</b>\n            <select name=\"realm\" id=\"realm\"></select>\n        </div>\n        <div class=\"label\" id=\"label-password\">Password\n            <input id=\"pass\" name=\"pass\" type=\"password\" value=\"\"/>\n        </div>\n        <input id=\"submit-btn\" type=\"submit\" name=\"login_action\" value=\"Submit\"/>\n        <p><font size=2 >-or-</font></p>\n        <input id=\"login-btn\" type=\"submit\" name=\"guest_action\" value=\"Login as guest\"/>\n    </form>\n</body>\n</html>\n")
	}
	if value := res.Get("dummy_categories"); value.Exists() {
		data.Categories = make([]IdentityPolicyCategories, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := IdentityPolicyCategories{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("name"); value.Exists() {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			(*parent).Categories = append((*parent).Categories, data)
			return true
		})
	}
	if value := res.Get("dummy_rules"); value.Exists() {
		data.Rules = make([]IdentityPolicyRules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := IdentityPolicyRules{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("name"); value.Exists() {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			if value := res.Get("enabled"); value.Exists() {
				data.Enabled = types.BoolValue(value.Bool())
			} else {
				data.Enabled = types.BoolNull()
			}
			if value := res.Get("authenticationType"); value.Exists() {
				data.AuthenticationType = types.StringValue(value.String())
			} else {
				data.AuthenticationType = types.StringNull()
			}
			if value := res.Get("realm.id"); value.Exists() {
				data.AuthenticationRealmId = types.StringValue(value.String())
			} else {
				data.AuthenticationRealmId = types.StringNull()
			}
			if value := res.Get("realm.type"); value.Exists() {
				data.AuthenticationRealmType = types.StringValue(value.String())
			} else {
				data.AuthenticationRealmType = types.StringNull()
			}
			if value := res.Get("guestAccessFallback"); value.Exists() {
				data.GuestAccessFallback = types.BoolValue(value.Bool())
			} else {
				data.GuestAccessFallback = types.BoolNull()
			}
			if value := res.Get("captivePortalFallback"); value.Exists() {
				data.ActiveAuthenticationFallback = types.BoolValue(value.Bool())
			} else {
				data.ActiveAuthenticationFallback = types.BoolNull()
			}
			if value := res.Get("authenticationProtocol"); value.Exists() {
				data.AuthenticationProtocol = types.StringValue(value.String())
			} else {
				data.AuthenticationProtocol = types.StringNull()
			}
			if value := res.Get("excludedUserAgent"); value.Exists() {
				data.ExcludedUserAgents = make([]IdentityPolicyRulesExcludedUserAgents, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := IdentityPolicyRulesExcludedUserAgents{}
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
					(*parent).ExcludedUserAgents = append((*parent).ExcludedUserAgents, data)
					return true
				})
			}
			if value := res.Get("sourceZones"); value.Exists() {
				data.SourceZones = make([]IdentityPolicyRulesSourceZones, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := IdentityPolicyRulesSourceZones{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					if value := res.Get("type"); value.Exists() {
						data.Type = types.StringValue(value.String())
					} else {
						data.Type = types.StringValue("SecurityZone")
					}
					(*parent).SourceZones = append((*parent).SourceZones, data)
					return true
				})
			}
			if value := res.Get("destinationZones"); value.Exists() {
				data.DestinationZones = make([]IdentityPolicyRulesDestinationZones, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := IdentityPolicyRulesDestinationZones{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					if value := res.Get("type"); value.Exists() {
						data.Type = types.StringValue(value.String())
					} else {
						data.Type = types.StringValue("SecurityZone")
					}
					(*parent).DestinationZones = append((*parent).DestinationZones, data)
					return true
				})
			}
			if value := res.Get("sourceNetworks.literals"); value.Exists() {
				data.SourceNetworkLiterals = make([]IdentityPolicyRulesSourceNetworkLiterals, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := IdentityPolicyRulesSourceNetworkLiterals{}
					if value := res.Get("type"); value.Exists() {
						data.Type = types.StringValue(value.String())
					} else {
						data.Type = types.StringNull()
					}
					if value := res.Get("value"); value.Exists() {
						data.Value = types.StringValue(value.String())
					} else {
						data.Value = types.StringNull()
					}
					(*parent).SourceNetworkLiterals = append((*parent).SourceNetworkLiterals, data)
					return true
				})
			}
			if value := res.Get("sourceNetworks.objects"); value.Exists() {
				data.SourceNetworkObjects = make([]IdentityPolicyRulesSourceNetworkObjects, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := IdentityPolicyRulesSourceNetworkObjects{}
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
			if value := res.Get("destinationNetworks.literals"); value.Exists() {
				data.DestinationNetworkLiterals = make([]IdentityPolicyRulesDestinationNetworkLiterals, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := IdentityPolicyRulesDestinationNetworkLiterals{}
					if value := res.Get("type"); value.Exists() {
						data.Type = types.StringValue(value.String())
					} else {
						data.Type = types.StringNull()
					}
					if value := res.Get("value"); value.Exists() {
						data.Value = types.StringValue(value.String())
					} else {
						data.Value = types.StringNull()
					}
					(*parent).DestinationNetworkLiterals = append((*parent).DestinationNetworkLiterals, data)
					return true
				})
			}
			if value := res.Get("destinationNetworks.objects"); value.Exists() {
				data.DestinationNetworkObjects = make([]IdentityPolicyRulesDestinationNetworkObjects, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := IdentityPolicyRulesDestinationNetworkObjects{}
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
			if value := res.Get("vlanTags.literals"); value.Exists() {
				data.VlanTagLiterals = make([]IdentityPolicyRulesVlanTagLiterals, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := IdentityPolicyRulesVlanTagLiterals{}
					if value := res.Get("type"); value.Exists() {
						data.Type = types.StringValue(value.String())
					} else {
						data.Type = types.StringValue("VlanTagLiteral")
					}
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
				data.VlanTagObjects = make([]IdentityPolicyRulesVlanTagObjects, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := IdentityPolicyRulesVlanTagObjects{}
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
					(*parent).VlanTagObjects = append((*parent).VlanTagObjects, data)
					return true
				})
			}
			if value := res.Get("sourcePorts.literals"); value.Exists() {
				data.SourcePortLiterals = make([]IdentityPolicyRulesSourcePortLiterals, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := IdentityPolicyRulesSourcePortLiterals{}
					if value := res.Get("type"); value.Exists() {
						data.Type = types.StringValue(value.String())
					} else {
						data.Type = types.StringValue("PortLiteral")
					}
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
				data.SourcePortObjects = make([]IdentityPolicyRulesSourcePortObjects, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := IdentityPolicyRulesSourcePortObjects{}
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
					(*parent).SourcePortObjects = append((*parent).SourcePortObjects, data)
					return true
				})
			}
			if value := res.Get("destinationPorts.literals"); value.Exists() {
				data.DestinationPortLiterals = make([]IdentityPolicyRulesDestinationPortLiterals, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := IdentityPolicyRulesDestinationPortLiterals{}
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
					(*parent).DestinationPortLiterals = append((*parent).DestinationPortLiterals, data)
					return true
				})
			}
			if value := res.Get("destinationPorts.objects"); value.Exists() {
				data.DestinationPortObjects = make([]IdentityPolicyRulesDestinationPortObjects, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := IdentityPolicyRulesDestinationPortObjects{}
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
					(*parent).DestinationPortObjects = append((*parent).DestinationPortObjects, data)
					return true
				})
			}
			(*parent).Rules = append((*parent).Rules, data)
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
func (data *IdentityPolicy) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("mappingFilter.id"); value.Exists() && !data.IdentityMappingFilterNetworkObjectId.IsNull() {
		data.IdentityMappingFilterNetworkObjectId = types.StringValue(value.String())
	} else {
		data.IdentityMappingFilterNetworkObjectId = types.StringNull()
	}
	if value := res.Get("mappingFilter.type"); value.Exists() && !data.IdentityMappingFilterNetworkObjectType.IsNull() {
		data.IdentityMappingFilterNetworkObjectType = types.StringValue(value.String())
	} else {
		data.IdentityMappingFilterNetworkObjectType = types.StringNull()
	}
	if value := res.Get("authCert.id"); value.Exists() && !data.ActiveAuthenticationServerCertificateId.IsNull() {
		data.ActiveAuthenticationServerCertificateId = types.StringValue(value.String())
	} else {
		data.ActiveAuthenticationServerCertificateId = types.StringNull()
	}
	if value := res.Get("authCert.type"); value.Exists() && !data.ActiveAuthenticationServerCertificateType.IsNull() {
		data.ActiveAuthenticationServerCertificateType = types.StringValue(value.String())
	} else if data.ActiveAuthenticationServerCertificateType.ValueString() != "PKI_InternalCert" {
		data.ActiveAuthenticationServerCertificateType = types.StringNull()
	}
	if value := res.Get("redirectHost.id"); value.Exists() && !data.ActiveAuthenticationRedirectFqdnId.IsNull() {
		data.ActiveAuthenticationRedirectFqdnId = types.StringValue(value.String())
	} else {
		data.ActiveAuthenticationRedirectFqdnId = types.StringNull()
	}
	if value := res.Get("redirectHost.type"); value.Exists() && !data.ActiveAuthenticationRedirectFqdnType.IsNull() {
		data.ActiveAuthenticationRedirectFqdnType = types.StringValue(value.String())
	} else if data.ActiveAuthenticationRedirectFqdnType.ValueString() != "FQDN" {
		data.ActiveAuthenticationRedirectFqdnType = types.StringNull()
	}
	if value := res.Get("redirectPort"); value.Exists() && !data.ActiveAuthenticationRedirectPort.IsNull() {
		data.ActiveAuthenticationRedirectPort = types.Int64Value(value.Int())
	} else if data.ActiveAuthenticationRedirectPort.ValueInt64() != 885 {
		data.ActiveAuthenticationRedirectPort = types.Int64Null()
	}
	if value := res.Get("maxLoginAttempts"); value.Exists() && !data.ActiveAuthenticationMaximumLoginAttempts.IsNull() {
		data.ActiveAuthenticationMaximumLoginAttempts = types.Int64Value(value.Int())
	} else if data.ActiveAuthenticationMaximumLoginAttempts.ValueInt64() != 3 {
		data.ActiveAuthenticationMaximumLoginAttempts = types.Int64Null()
	}
	if value := res.Get("activeSessionSharing"); value.Exists() && !data.ActiveAuthenticationSessionSharing.IsNull() {
		data.ActiveAuthenticationSessionSharing = types.BoolValue(value.Bool())
	} else if data.ActiveAuthenticationSessionSharing.ValueBool() != true {
		data.ActiveAuthenticationSessionSharing = types.BoolNull()
	}
	if value := res.Get("activeAuthPage"); value.Exists() && !data.ActiveAuthenticationPage.IsNull() {
		data.ActiveAuthenticationPage = types.StringValue(value.String())
	} else if data.ActiveAuthenticationPage.ValueString() != "DEFAULT" {
		data.ActiveAuthenticationPage = types.StringNull()
	}
	if value := res.Get("htmlActiveAuthPage"); value.Exists() && !data.ActiveAuthenticationPageHtml.IsNull() {
		data.ActiveAuthenticationPageHtml = types.StringValue(value.String())
	} else if data.ActiveAuthenticationPageHtml.ValueString() != "<!DOCTYPE html>\n<html>\n<head>\n<meta http-equiv=\"content-type\" content=\"text/html; charset=UTF-8\" />\n<title>Login</title>\n<style type=\"text/css\">\nbody {\n    margin:0;\n    font-family:verdana,sans-serif;\n}\nh1 {\n    margin:0;\n    padding:12px 25px;\n    background-color:#343434;\n    color:#ddd;\n}\np {\n    margin:12px 25px;\n}\nstrong {\n    color:#E0042D;\n}\ndiv {\n    padding-left:23px;\n    font-weight: normal;\n    font-size: 8pt;\n}\ninput {\n    margin:12px 25px;\n}\n</style>\n</head>\n<body>\n    <form action=\"\" id=\"loginForm\" method=\"post\" name=\"loginForm\">\n        <h1>Login</h1>\n        <p><strong>Please enter your username and password or log in as a guest.</strong></p>\n        <div class=\"label\">Username\n            <input id=\"name\" maxlength=\"100\" name=\"login\" type=\"text\" value=\"\"/>\n            <b>realm</b>\n            <select name=\"realm\" id=\"realm\"></select>\n        </div>\n        <div class=\"label\" id=\"label-password\">Password\n            <input id=\"pass\" name=\"pass\" type=\"password\" value=\"\"/>\n        </div>\n        <input id=\"submit-btn\" type=\"submit\" name=\"login_action\" value=\"Submit\"/>\n        <p><font size=2 >-or-</font></p>\n        <input id=\"login-btn\" type=\"submit\" name=\"guest_action\" value=\"Login as guest\"/>\n    </form>\n</body>\n</html>\n" {
		data.ActiveAuthenticationPageHtml = types.StringNull()
	}
	{
		l := len(res.Get("dummy_categories").Array())
		tflog.Debug(ctx, fmt.Sprintf("dummy_categories array resizing from %d to %d", len(data.Categories), l))
		for i := len(data.Categories); i < l; i++ {
			data.Categories = append(data.Categories, IdentityPolicyCategories{})
		}
		if len(data.Categories) > l {
			data.Categories = data.Categories[:l]
		}
	}
	for i := range data.Categories {
		parent := &data
		data := (*parent).Categories[i]
		parentRes := &res
		res := parentRes.Get(fmt.Sprintf("dummy_categories.%d", i))
		if value := res.Get("id"); value.Exists() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		(*parent).Categories[i] = data
	}
	{
		l := len(res.Get("dummy_rules").Array())
		tflog.Debug(ctx, fmt.Sprintf("dummy_rules array resizing from %d to %d", len(data.Rules), l))
		for i := len(data.Rules); i < l; i++ {
			data.Rules = append(data.Rules, IdentityPolicyRules{})
		}
		if len(data.Rules) > l {
			data.Rules = data.Rules[:l]
		}
	}
	for i := range data.Rules {
		parent := &data
		data := (*parent).Rules[i]
		parentRes := &res
		res := parentRes.Get(fmt.Sprintf("dummy_rules.%d", i))
		if value := res.Get("id"); value.Exists() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
			data.Enabled = types.BoolValue(value.Bool())
		} else {
			data.Enabled = types.BoolNull()
		}
		if value := res.Get("authenticationType"); value.Exists() && !data.AuthenticationType.IsNull() {
			data.AuthenticationType = types.StringValue(value.String())
		} else {
			data.AuthenticationType = types.StringNull()
		}
		if value := res.Get("realm.id"); value.Exists() && !data.AuthenticationRealmId.IsNull() {
			data.AuthenticationRealmId = types.StringValue(value.String())
		} else {
			data.AuthenticationRealmId = types.StringNull()
		}
		if value := res.Get("realm.type"); value.Exists() && !data.AuthenticationRealmType.IsNull() {
			data.AuthenticationRealmType = types.StringValue(value.String())
		} else {
			data.AuthenticationRealmType = types.StringNull()
		}
		if value := res.Get("guestAccessFallback"); value.Exists() && !data.GuestAccessFallback.IsNull() {
			data.GuestAccessFallback = types.BoolValue(value.Bool())
		} else {
			data.GuestAccessFallback = types.BoolNull()
		}
		if value := res.Get("captivePortalFallback"); value.Exists() && !data.ActiveAuthenticationFallback.IsNull() {
			data.ActiveAuthenticationFallback = types.BoolValue(value.Bool())
		} else {
			data.ActiveAuthenticationFallback = types.BoolNull()
		}
		if value := res.Get("authenticationProtocol"); value.Exists() && !data.AuthenticationProtocol.IsNull() {
			data.AuthenticationProtocol = types.StringValue(value.String())
		} else {
			data.AuthenticationProtocol = types.StringNull()
		}
		for i := 0; i < len(data.ExcludedUserAgents); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.ExcludedUserAgents[i].Id.ValueString()}

			parent := &data
			data := (*parent).ExcludedUserAgents[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("excludedUserAgent").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing ExcludedUserAgents[%d] = %+v",
					i,
					(*parent).ExcludedUserAgents[i],
				))
				(*parent).ExcludedUserAgents = slices.Delete((*parent).ExcludedUserAgents, i, i+1)
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
			(*parent).ExcludedUserAgents[i] = data
		}
		for i := 0; i < len(data.SourceZones); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.SourceZones[i].Id.ValueString()}

			parent := &data
			data := (*parent).SourceZones[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("sourceZones").ForEach(
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
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else if data.Type.ValueString() != "SecurityZone" {
				data.Type = types.StringNull()
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

			parentRes.Get("destinationZones").ForEach(
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
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else if data.Type.ValueString() != "SecurityZone" {
				data.Type = types.StringNull()
			}
			(*parent).DestinationZones[i] = data
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
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			if value := res.Get("value"); value.Exists() && !data.Value.IsNull() {
				data.Value = types.StringValue(value.String())
			} else {
				data.Value = types.StringNull()
			}
			(*parent).SourceNetworkLiterals[i] = data
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
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			if value := res.Get("value"); value.Exists() && !data.Value.IsNull() {
				data.Value = types.StringValue(value.String())
			} else {
				data.Value = types.StringNull()
			}
			(*parent).DestinationNetworkLiterals[i] = data
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
		for i := 0; i < len(data.VlanTagLiterals); i++ {
			keys := [...]string{"type", "startTag", "endTag"}
			keyValues := [...]string{data.VlanTagLiterals[i].Type.ValueString(), data.VlanTagLiterals[i].StartTag.ValueString(), data.VlanTagLiterals[i].EndTag.ValueString()}

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
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else if data.Type.ValueString() != "VlanTagLiteral" {
				data.Type = types.StringNull()
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
			keys := [...]string{"id", "type"}
			keyValues := [...]string{data.VlanTagObjects[i].Id.ValueString(), data.VlanTagObjects[i].Type.ValueString()}

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
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			(*parent).VlanTagObjects[i] = data
		}
		for i := 0; i < len(data.SourcePortLiterals); i++ {
			keys := [...]string{"type", "protocol", "port"}
			keyValues := [...]string{data.SourcePortLiterals[i].Type.ValueString(), data.SourcePortLiterals[i].Protocol.ValueString(), data.SourcePortLiterals[i].Port.ValueString()}

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
			} else if data.Type.ValueString() != "PortLiteral" {
				data.Type = types.StringNull()
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
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			(*parent).SourcePortObjects[i] = data
		}
		for i := 0; i < len(data.DestinationPortLiterals); i++ {
			keys := [...]string{"type", "port", "protocol"}
			keyValues := [...]string{data.DestinationPortLiterals[i].Type.ValueString(), data.DestinationPortLiterals[i].Port.ValueString(), data.DestinationPortLiterals[i].Protocol.ValueString()}

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
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			(*parent).DestinationPortObjects[i] = data
		}
		(*parent).Rules[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *IdentityPolicy) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
	for i := range data.Categories {
		r := res.Get(fmt.Sprintf("dummy_categories.%d", i))
		if v := data.Categories[i]; v.Id.IsUnknown() {
			if value := r.Get("id"); value.Exists() {
				v.Id = types.StringValue(value.String())
			} else {
				v.Id = types.StringNull()
			}
			data.Categories[i] = v
		}
	}
	for i := range data.Rules {
		r := res.Get(fmt.Sprintf("dummy_rules.%d", i))
		if v := data.Rules[i]; v.Id.IsUnknown() {
			if value := r.Get("id"); value.Exists() {
				v.Id = types.StringValue(value.String())
			} else {
				v.Id = types.StringNull()
			}
			data.Rules[i] = v
		}
	}
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

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBody

// End of section. //template:end adjustBody

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBodyBulk

// End of section. //template:end adjustBodyBulk

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyOverrides

// End of section. //template:end toBodyOverrides

// Section below is generated&owned by "gen/generator.go". //template:begin synthesizeOverrides

// End of section. //template:end synthesizeOverrides
