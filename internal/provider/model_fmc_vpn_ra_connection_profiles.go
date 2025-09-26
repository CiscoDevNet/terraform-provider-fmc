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
	"maps"
	"net/url"
	"slices"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type VPNRAConnectionProfiles struct {
	Id      types.String                            `tfsdk:"id"`
	Domain  types.String                            `tfsdk:"domain"`
	VpnRaId types.String                            `tfsdk:"vpn_ra_id"`
	Items   map[string]VPNRAConnectionProfilesItems `tfsdk:"items"`
}

type VPNRAConnectionProfilesItems struct {
	Id                                                                   types.String                                   `tfsdk:"id"`
	Type                                                                 types.String                                   `tfsdk:"type"`
	GroupPolicyId                                                        types.String                                   `tfsdk:"group_policy_id"`
	Ipv4AddressPools                                                     []VPNRAConnectionProfilesItemsIpv4AddressPools `tfsdk:"ipv4_address_pools"`
	Ipv6AddressPools                                                     []VPNRAConnectionProfilesItemsIpv6AddressPools `tfsdk:"ipv6_address_pools"`
	DhcpServers                                                          []VPNRAConnectionProfilesItemsDhcpServers      `tfsdk:"dhcp_servers"`
	AuthenticationMethod                                                 types.String                                   `tfsdk:"authentication_method"`
	MultipleCertificateAuthentication                                    types.Bool                                     `tfsdk:"multiple_certificate_authentication"`
	PrimaryAuthenticationServerUseLocal                                  types.Bool                                     `tfsdk:"primary_authentication_server_use_local"`
	PrimaryAuthenticationServerId                                        types.String                                   `tfsdk:"primary_authentication_server_id"`
	PrimaryAuthenticationServerType                                      types.String                                   `tfsdk:"primary_authentication_server_type"`
	PrimaryAuthenticationFallbackToLocal                                 types.Bool                                     `tfsdk:"primary_authentication_fallback_to_local"`
	SamlAndCertificateUsernameMustMatch                                  types.Bool                                     `tfsdk:"saml_and_certificate_username_must_match"`
	PrimaryAuthenticationPrefillUsernameFromCertificate                  types.Bool                                     `tfsdk:"primary_authentication_prefill_username_from_certificate"`
	PrimaryAuthenticationPrefillUsernameFromCertificateMapPrimaryField   types.String                                   `tfsdk:"primary_authentication_prefill_username_from_certificate_map_primary_field"`
	PrimaryAuthenticationPrefillUsernameFromCertificateMapSecondaryField types.String                                   `tfsdk:"primary_authentication_prefill_username_from_certificate_map_secondary_field"`
	PrimaryAuthenticationPrefillUsernameFromCertificateMapEntireDn       types.Bool                                     `tfsdk:"primary_authentication_prefill_username_from_certificate_map_entire_dn"`
	PrimaryAuthenticationHideUsernameInLoginWindow                       types.Bool                                     `tfsdk:"primary_authentication_hide_username_in_login_window"`
	SecondaryAuthentication                                              types.Bool                                     `tfsdk:"secondary_authentication"`
	SecondaryAuthenticationServerUseLocal                                types.Bool                                     `tfsdk:"secondary_authentication_server_use_local"`
	SecondaryAuthenticationServerId                                      types.String                                   `tfsdk:"secondary_authentication_server_id"`
	SecondaryAuthenticationServerType                                    types.String                                   `tfsdk:"secondary_authentication_server_type"`
	SecondaryAuthenticationFallbackToLocal                               types.Bool                                     `tfsdk:"secondary_authentication_fallback_to_local"`
	SecondaryAuthenticationPromptForUsername                             types.Bool                                     `tfsdk:"secondary_authentication_prompt_for_username"`
	SecondaryAuthenticationUsePrimaryAuthenticationUsername              types.Bool                                     `tfsdk:"secondary_authentication_use_primary_authentication_username"`
	UseSecondaryAuthenticationUsernameForReporting                       types.Bool                                     `tfsdk:"use_secondary_authentication_username_for_reporting"`
	SamlUseExternalBrowser                                               types.Bool                                     `tfsdk:"saml_use_external_browser"`
	AuthorizationServerId                                                types.String                                   `tfsdk:"authorization_server_id"`
	AuthorizationServerType                                              types.String                                   `tfsdk:"authorization_server_type"`
	AllowConnectionOnlyIfUserExistsInAuthorizationDatabase               types.Bool                                     `tfsdk:"allow_connection_only_if_user_exists_in_authorization_database"`
	AccountingServerId                                                   types.String                                   `tfsdk:"accounting_server_id"`
	AccountingServerType                                                 types.String                                   `tfsdk:"accounting_server_type"`
	StripRealmFromUsername                                               types.Bool                                     `tfsdk:"strip_realm_from_username"`
	StripGroupFromUsername                                               types.Bool                                     `tfsdk:"strip_group_from_username"`
	PasswordManagement                                                   types.Bool                                     `tfsdk:"password_management"`
	PasswordManagementNotifyUserOnPasswordExpiryDay                      types.Bool                                     `tfsdk:"password_management_notify_user_on_password_expiry_day"`
	PasswordManagementAdvancePasswordExpirationNotification              types.Int64                                    `tfsdk:"password_management_advance_password_expiration_notification"`
	AliasNames                                                           []VPNRAConnectionProfilesItemsAliasNames       `tfsdk:"alias_names"`
	AliasUrls                                                            []VPNRAConnectionProfilesItemsAliasUrls        `tfsdk:"alias_urls"`
}

type VPNRAConnectionProfilesItemsIpv4AddressPools struct {
	Id types.String `tfsdk:"id"`
}
type VPNRAConnectionProfilesItemsIpv6AddressPools struct {
	Id types.String `tfsdk:"id"`
}
type VPNRAConnectionProfilesItemsDhcpServers struct {
	Id types.String `tfsdk:"id"`
}
type VPNRAConnectionProfilesItemsAliasNames struct {
	Name    types.String `tfsdk:"name"`
	Enabled types.Bool   `tfsdk:"enabled"`
}
type VPNRAConnectionProfilesItemsAliasUrls struct {
	UrlObjectId types.String `tfsdk:"url_object_id"`
	Enabled     types.Bool   `tfsdk:"enabled"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionBulkCreateVPNRAConnectionProfiles = version.Must(version.NewVersion("999"))
var minFMCVersionBulkDeleteVPNRAConnectionProfiles = version.Must(version.NewVersion("999"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data VPNRAConnectionProfiles) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/ravpns/%v/connectionprofiles", url.QueryEscape(data.VpnRaId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data VPNRAConnectionProfiles) toBody(ctx context.Context, state VPNRAConnectionProfiles) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if len(data.Items) > 0 {
		body, _ = sjson.Set(body, "items", []any{})
		for key, item := range data.Items {
			itemBody, _ := sjson.Set("{}", "name", key)
			if !item.Id.IsNull() && !item.Id.IsUnknown() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			itemBody, _ = sjson.Set(itemBody, "type", "RaVpnConnectionProfile")
			if !item.GroupPolicyId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "groupPolicy.id", item.GroupPolicyId.ValueString())
			}
			if len(item.Ipv4AddressPools) > 0 {
				itemBody, _ = sjson.Set(itemBody, "ipv4AddressPool", []any{})
				for _, childItem := range item.Ipv4AddressPools {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "ipv4AddressPool.-1", itemChildBody)
				}
			}
			if len(item.Ipv6AddressPools) > 0 {
				itemBody, _ = sjson.Set(itemBody, "ipv6AddressPool", []any{})
				for _, childItem := range item.Ipv6AddressPools {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "ipv6AddressPool.-1", itemChildBody)
				}
			}
			if len(item.DhcpServers) > 0 {
				itemBody, _ = sjson.Set(itemBody, "dhcpServersForAddressAssignment", []any{})
				for _, childItem := range item.DhcpServers {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "dhcpServersForAddressAssignment.-1", itemChildBody)
				}
			}
			if !item.AuthenticationMethod.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "authenticationMethod", item.AuthenticationMethod.ValueString())
			}
			if !item.MultipleCertificateAuthentication.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enableMultipleCertificateAuthentication", item.MultipleCertificateAuthentication.ValueBool())
			}
			if !item.PrimaryAuthenticationServerUseLocal.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "useLocalAsPrimaryAuthServer", item.PrimaryAuthenticationServerUseLocal.ValueBool())
			}
			if !item.PrimaryAuthenticationServerId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "primaryAuthenticationServer.id", item.PrimaryAuthenticationServerId.ValueString())
			}
			if !item.PrimaryAuthenticationServerType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "primaryAuthenticationServer.type", item.PrimaryAuthenticationServerType.ValueString())
			}
			if !item.PrimaryAuthenticationFallbackToLocal.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enablePrimaryAuthFallbackToLocal", item.PrimaryAuthenticationFallbackToLocal.ValueBool())
			}
			if !item.SamlAndCertificateUsernameMustMatch.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "certificateUsernameSettings.matchCertificateAndSamlUsername", item.SamlAndCertificateUsernameMustMatch.ValueBool())
			}
			if !item.PrimaryAuthenticationPrefillUsernameFromCertificate.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "certificateUsernameSettings.prefillUsernameFromCertificate", item.PrimaryAuthenticationPrefillUsernameFromCertificate.ValueBool())
			}
			if !item.PrimaryAuthenticationPrefillUsernameFromCertificateMapPrimaryField.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "certificateUsernameSettings.mapPrimaryField", item.PrimaryAuthenticationPrefillUsernameFromCertificateMapPrimaryField.ValueString())
			}
			if !item.PrimaryAuthenticationPrefillUsernameFromCertificateMapSecondaryField.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "certificateUsernameSettings.mapSecondaryField", item.PrimaryAuthenticationPrefillUsernameFromCertificateMapSecondaryField.ValueString())
			}
			if !item.PrimaryAuthenticationPrefillUsernameFromCertificateMapEntireDn.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "certificateUsernameSettings.mapEntireDNAsUsername", item.PrimaryAuthenticationPrefillUsernameFromCertificateMapEntireDn.ValueBool())
			}
			if !item.PrimaryAuthenticationHideUsernameInLoginWindow.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "certificateUsernameSettings.hideUsername", item.PrimaryAuthenticationHideUsernameInLoginWindow.ValueBool())
			}
			if !item.SecondaryAuthentication.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enableSecondaryAuthentication", item.SecondaryAuthentication.ValueBool())
			}
			if !item.SecondaryAuthenticationServerUseLocal.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "useLocalAsSecondaryAuthServer", item.SecondaryAuthenticationServerUseLocal.ValueBool())
			}
			if !item.SecondaryAuthenticationServerId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "secondaryAuthenticationServer.id", item.SecondaryAuthenticationServerId.ValueString())
			}
			if !item.SecondaryAuthenticationServerType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "secondaryAuthenticationServer.type", item.SecondaryAuthenticationServerType.ValueString())
			}
			if !item.SecondaryAuthenticationFallbackToLocal.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enableSecondaryAuthFallbackToLocal", item.SecondaryAuthenticationFallbackToLocal.ValueBool())
			}
			if !item.SecondaryAuthenticationPromptForUsername.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "secondaryAuthenticationSettings.promptUsername", item.SecondaryAuthenticationPromptForUsername.ValueBool())
			}
			if !item.SecondaryAuthenticationUsePrimaryAuthenticationUsername.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "secondaryAuthenticationSettings.usePrimaryUsername", item.SecondaryAuthenticationUsePrimaryAuthenticationUsername.ValueBool())
			}
			if !item.UseSecondaryAuthenticationUsernameForReporting.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "secondaryAuthenticationSettings.useSecondaryUsernameforSession", item.UseSecondaryAuthenticationUsernameForReporting.ValueBool())
			}
			if !item.SamlUseExternalBrowser.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enableExternalBrowserForSAML", item.SamlUseExternalBrowser.ValueBool())
			}
			if !item.AuthorizationServerId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "authorizationServer.id", item.AuthorizationServerId.ValueString())
			}
			if !item.AuthorizationServerType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "authorizationServer.type", item.AuthorizationServerType.ValueString())
			}
			if !item.AllowConnectionOnlyIfUserExistsInAuthorizationDatabase.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "allowConnectionOnlyIfAuthorized", item.AllowConnectionOnlyIfUserExistsInAuthorizationDatabase.ValueBool())
			}
			if !item.AccountingServerId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "accountingServer.id", item.AccountingServerId.ValueString())
			}
			if !item.AccountingServerType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "accountingServer.type", item.AccountingServerType.ValueString())
			}
			if !item.StripRealmFromUsername.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "stripRealmFromUsername", item.StripRealmFromUsername.ValueBool())
			}
			if !item.StripGroupFromUsername.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "stripGroupFromUsername", item.StripGroupFromUsername.ValueBool())
			}
			if !item.PasswordManagement.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enablePasswordManagement", item.PasswordManagement.ValueBool())
			}
			if !item.PasswordManagementNotifyUserOnPasswordExpiryDay.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "notifyUserOnPasswordExpiryDay", item.PasswordManagementNotifyUserOnPasswordExpiryDay.ValueBool())
			}
			if !item.PasswordManagementAdvancePasswordExpirationNotification.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "passwordExpirationNotificationPeriod", item.PasswordManagementAdvancePasswordExpirationNotification.ValueInt64())
			}
			if len(item.AliasNames) > 0 {
				itemBody, _ = sjson.Set(itemBody, "groupAlias", []any{})
				for _, childItem := range item.AliasNames {
					itemChildBody := ""
					if !childItem.Name.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "aliasName", childItem.Name.ValueString())
					}
					if !childItem.Enabled.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "enabled", childItem.Enabled.ValueBool())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "groupAlias.-1", itemChildBody)
				}
			}
			if len(item.AliasUrls) > 0 {
				itemBody, _ = sjson.Set(itemBody, "groupUrl", []any{})
				for _, childItem := range item.AliasUrls {
					itemChildBody := ""
					if !childItem.UrlObjectId.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "aliasUrl.id", childItem.UrlObjectId.ValueString())
					}
					if !childItem.Enabled.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "enabled", childItem.Enabled.ValueBool())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "groupUrl.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "items.-1", itemBody)
		}
	}
	return gjson.Get(body, "items").String()
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *VPNRAConnectionProfiles) fromBody(ctx context.Context, res gjson.Result) {
	for k := range data.Items {
		parent := &data
		data := (*parent).Items[k]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("items").ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("name").String() == k {
					res = v
					return false // break ForEach
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("subresource not found, removing: name=%v", k))
			delete((*parent).Items, k)
			continue
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
		if value := res.Get("groupPolicy.id"); value.Exists() {
			data.GroupPolicyId = types.StringValue(value.String())
		} else {
			data.GroupPolicyId = types.StringNull()
		}
		if value := res.Get("ipv4AddressPool"); value.Exists() {
			data.Ipv4AddressPools = make([]VPNRAConnectionProfilesItemsIpv4AddressPools, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := VPNRAConnectionProfilesItemsIpv4AddressPools{}
				if value := res.Get("id"); value.Exists() {
					data.Id = types.StringValue(value.String())
				} else {
					data.Id = types.StringNull()
				}
				(*parent).Ipv4AddressPools = append((*parent).Ipv4AddressPools, data)
				return true
			})
		}
		if value := res.Get("ipv6AddressPool"); value.Exists() {
			data.Ipv6AddressPools = make([]VPNRAConnectionProfilesItemsIpv6AddressPools, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := VPNRAConnectionProfilesItemsIpv6AddressPools{}
				if value := res.Get("id"); value.Exists() {
					data.Id = types.StringValue(value.String())
				} else {
					data.Id = types.StringNull()
				}
				(*parent).Ipv6AddressPools = append((*parent).Ipv6AddressPools, data)
				return true
			})
		}
		if value := res.Get("dhcpServersForAddressAssignment"); value.Exists() {
			data.DhcpServers = make([]VPNRAConnectionProfilesItemsDhcpServers, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := VPNRAConnectionProfilesItemsDhcpServers{}
				if value := res.Get("id"); value.Exists() {
					data.Id = types.StringValue(value.String())
				} else {
					data.Id = types.StringNull()
				}
				(*parent).DhcpServers = append((*parent).DhcpServers, data)
				return true
			})
		}
		if value := res.Get("authenticationMethod"); value.Exists() {
			data.AuthenticationMethod = types.StringValue(value.String())
		} else {
			data.AuthenticationMethod = types.StringNull()
		}
		if value := res.Get("enableMultipleCertificateAuthentication"); value.Exists() {
			data.MultipleCertificateAuthentication = types.BoolValue(value.Bool())
		} else {
			data.MultipleCertificateAuthentication = types.BoolNull()
		}
		if value := res.Get("useLocalAsPrimaryAuthServer"); value.Exists() {
			data.PrimaryAuthenticationServerUseLocal = types.BoolValue(value.Bool())
		} else {
			data.PrimaryAuthenticationServerUseLocal = types.BoolNull()
		}
		if value := res.Get("primaryAuthenticationServer.id"); value.Exists() {
			data.PrimaryAuthenticationServerId = types.StringValue(value.String())
		} else {
			data.PrimaryAuthenticationServerId = types.StringNull()
		}
		if value := res.Get("primaryAuthenticationServer.type"); value.Exists() {
			data.PrimaryAuthenticationServerType = types.StringValue(value.String())
		} else {
			data.PrimaryAuthenticationServerType = types.StringNull()
		}
		if value := res.Get("enablePrimaryAuthFallbackToLocal"); value.Exists() {
			data.PrimaryAuthenticationFallbackToLocal = types.BoolValue(value.Bool())
		} else {
			data.PrimaryAuthenticationFallbackToLocal = types.BoolNull()
		}
		if value := res.Get("certificateUsernameSettings.matchCertificateAndSamlUsername"); value.Exists() {
			data.SamlAndCertificateUsernameMustMatch = types.BoolValue(value.Bool())
		} else {
			data.SamlAndCertificateUsernameMustMatch = types.BoolNull()
		}
		if value := res.Get("certificateUsernameSettings.prefillUsernameFromCertificate"); value.Exists() {
			data.PrimaryAuthenticationPrefillUsernameFromCertificate = types.BoolValue(value.Bool())
		} else {
			data.PrimaryAuthenticationPrefillUsernameFromCertificate = types.BoolNull()
		}
		if value := res.Get("certificateUsernameSettings.mapPrimaryField"); value.Exists() {
			data.PrimaryAuthenticationPrefillUsernameFromCertificateMapPrimaryField = types.StringValue(value.String())
		} else {
			data.PrimaryAuthenticationPrefillUsernameFromCertificateMapPrimaryField = types.StringNull()
		}
		if value := res.Get("certificateUsernameSettings.mapSecondaryField"); value.Exists() {
			data.PrimaryAuthenticationPrefillUsernameFromCertificateMapSecondaryField = types.StringValue(value.String())
		} else {
			data.PrimaryAuthenticationPrefillUsernameFromCertificateMapSecondaryField = types.StringNull()
		}
		if value := res.Get("certificateUsernameSettings.mapEntireDNAsUsername"); value.Exists() {
			data.PrimaryAuthenticationPrefillUsernameFromCertificateMapEntireDn = types.BoolValue(value.Bool())
		} else {
			data.PrimaryAuthenticationPrefillUsernameFromCertificateMapEntireDn = types.BoolNull()
		}
		if value := res.Get("certificateUsernameSettings.hideUsername"); value.Exists() {
			data.PrimaryAuthenticationHideUsernameInLoginWindow = types.BoolValue(value.Bool())
		} else {
			data.PrimaryAuthenticationHideUsernameInLoginWindow = types.BoolNull()
		}
		if value := res.Get("enableSecondaryAuthentication"); value.Exists() {
			data.SecondaryAuthentication = types.BoolValue(value.Bool())
		} else {
			data.SecondaryAuthentication = types.BoolNull()
		}
		if value := res.Get("useLocalAsSecondaryAuthServer"); value.Exists() {
			data.SecondaryAuthenticationServerUseLocal = types.BoolValue(value.Bool())
		} else {
			data.SecondaryAuthenticationServerUseLocal = types.BoolNull()
		}
		if value := res.Get("secondaryAuthenticationServer.id"); value.Exists() {
			data.SecondaryAuthenticationServerId = types.StringValue(value.String())
		} else {
			data.SecondaryAuthenticationServerId = types.StringNull()
		}
		if value := res.Get("secondaryAuthenticationServer.type"); value.Exists() {
			data.SecondaryAuthenticationServerType = types.StringValue(value.String())
		} else {
			data.SecondaryAuthenticationServerType = types.StringNull()
		}
		if value := res.Get("enableSecondaryAuthFallbackToLocal"); value.Exists() {
			data.SecondaryAuthenticationFallbackToLocal = types.BoolValue(value.Bool())
		} else {
			data.SecondaryAuthenticationFallbackToLocal = types.BoolNull()
		}
		if value := res.Get("secondaryAuthenticationSettings.promptUsername"); value.Exists() {
			data.SecondaryAuthenticationPromptForUsername = types.BoolValue(value.Bool())
		} else {
			data.SecondaryAuthenticationPromptForUsername = types.BoolNull()
		}
		if value := res.Get("secondaryAuthenticationSettings.usePrimaryUsername"); value.Exists() {
			data.SecondaryAuthenticationUsePrimaryAuthenticationUsername = types.BoolValue(value.Bool())
		} else {
			data.SecondaryAuthenticationUsePrimaryAuthenticationUsername = types.BoolNull()
		}
		if value := res.Get("secondaryAuthenticationSettings.useSecondaryUsernameforSession"); value.Exists() {
			data.UseSecondaryAuthenticationUsernameForReporting = types.BoolValue(value.Bool())
		} else {
			data.UseSecondaryAuthenticationUsernameForReporting = types.BoolNull()
		}
		if value := res.Get("authorizationServer.id"); value.Exists() {
			data.AuthorizationServerId = types.StringValue(value.String())
		} else {
			data.AuthorizationServerId = types.StringNull()
		}
		if value := res.Get("authorizationServer.type"); value.Exists() {
			data.AuthorizationServerType = types.StringValue(value.String())
		} else {
			data.AuthorizationServerType = types.StringNull()
		}
		if value := res.Get("allowConnectionOnlyIfAuthorized"); value.Exists() {
			data.AllowConnectionOnlyIfUserExistsInAuthorizationDatabase = types.BoolValue(value.Bool())
		} else {
			data.AllowConnectionOnlyIfUserExistsInAuthorizationDatabase = types.BoolNull()
		}
		if value := res.Get("accountingServer.id"); value.Exists() {
			data.AccountingServerId = types.StringValue(value.String())
		} else {
			data.AccountingServerId = types.StringNull()
		}
		if value := res.Get("accountingServer.type"); value.Exists() {
			data.AccountingServerType = types.StringValue(value.String())
		} else {
			data.AccountingServerType = types.StringNull()
		}
		if value := res.Get("stripRealmFromUsername"); value.Exists() {
			data.StripRealmFromUsername = types.BoolValue(value.Bool())
		} else {
			data.StripRealmFromUsername = types.BoolNull()
		}
		if value := res.Get("stripGroupFromUsername"); value.Exists() {
			data.StripGroupFromUsername = types.BoolValue(value.Bool())
		} else {
			data.StripGroupFromUsername = types.BoolNull()
		}
		if value := res.Get("enablePasswordManagement"); value.Exists() {
			data.PasswordManagement = types.BoolValue(value.Bool())
		} else {
			data.PasswordManagement = types.BoolNull()
		}
		if value := res.Get("notifyUserOnPasswordExpiryDay"); value.Exists() {
			data.PasswordManagementNotifyUserOnPasswordExpiryDay = types.BoolValue(value.Bool())
		} else {
			data.PasswordManagementNotifyUserOnPasswordExpiryDay = types.BoolNull()
		}
		if value := res.Get("passwordExpirationNotificationPeriod"); value.Exists() {
			data.PasswordManagementAdvancePasswordExpirationNotification = types.Int64Value(value.Int())
		} else {
			data.PasswordManagementAdvancePasswordExpirationNotification = types.Int64Null()
		}
		if value := res.Get("groupAlias"); value.Exists() {
			data.AliasNames = make([]VPNRAConnectionProfilesItemsAliasNames, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := VPNRAConnectionProfilesItemsAliasNames{}
				if value := res.Get("aliasName"); value.Exists() {
					data.Name = types.StringValue(value.String())
				} else {
					data.Name = types.StringNull()
				}
				if value := res.Get("enabled"); value.Exists() {
					data.Enabled = types.BoolValue(value.Bool())
				} else {
					data.Enabled = types.BoolNull()
				}
				(*parent).AliasNames = append((*parent).AliasNames, data)
				return true
			})
		}
		if value := res.Get("groupUrl"); value.Exists() {
			data.AliasUrls = make([]VPNRAConnectionProfilesItemsAliasUrls, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := VPNRAConnectionProfilesItemsAliasUrls{}
				if value := res.Get("aliasUrl.id"); value.Exists() {
					data.UrlObjectId = types.StringValue(value.String())
				} else {
					data.UrlObjectId = types.StringNull()
				}
				if value := res.Get("enabled"); value.Exists() {
					data.Enabled = types.BoolValue(value.Bool())
				} else {
					data.Enabled = types.BoolNull()
				}
				(*parent).AliasUrls = append((*parent).AliasUrls, data)
				return true
			})
		}
		(*parent).Items[k] = data
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *VPNRAConnectionProfiles) fromBodyPartial(ctx context.Context, res gjson.Result) {
	for i := range data.Items {
		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("items").ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("id").String() == data.Id.ValueString() && data.Id.ValueString() != "" {
					res = v
					return false // break ForEach
				}
				return true
			},
		)
		if value := res.Get("id"); value.Exists() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		if value := res.Get("groupPolicy.id"); value.Exists() && !data.GroupPolicyId.IsNull() {
			data.GroupPolicyId = types.StringValue(value.String())
		} else {
			data.GroupPolicyId = types.StringNull()
		}
		for i := 0; i < len(data.Ipv4AddressPools); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.Ipv4AddressPools[i].Id.ValueString()}

			parent := &data
			data := (*parent).Ipv4AddressPools[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("ipv4AddressPool").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing Ipv4AddressPools[%d] = %+v",
					i,
					(*parent).Ipv4AddressPools[i],
				))
				(*parent).Ipv4AddressPools = slices.Delete((*parent).Ipv4AddressPools, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).Ipv4AddressPools[i] = data
		}
		for i := 0; i < len(data.Ipv6AddressPools); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.Ipv6AddressPools[i].Id.ValueString()}

			parent := &data
			data := (*parent).Ipv6AddressPools[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("ipv6AddressPool").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing Ipv6AddressPools[%d] = %+v",
					i,
					(*parent).Ipv6AddressPools[i],
				))
				(*parent).Ipv6AddressPools = slices.Delete((*parent).Ipv6AddressPools, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).Ipv6AddressPools[i] = data
		}
		for i := 0; i < len(data.DhcpServers); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.DhcpServers[i].Id.ValueString()}

			parent := &data
			data := (*parent).DhcpServers[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("dhcpServersForAddressAssignment").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing DhcpServers[%d] = %+v",
					i,
					(*parent).DhcpServers[i],
				))
				(*parent).DhcpServers = slices.Delete((*parent).DhcpServers, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).DhcpServers[i] = data
		}
		if value := res.Get("authenticationMethod"); value.Exists() && !data.AuthenticationMethod.IsNull() {
			data.AuthenticationMethod = types.StringValue(value.String())
		} else {
			data.AuthenticationMethod = types.StringNull()
		}
		if value := res.Get("enableMultipleCertificateAuthentication"); value.Exists() && !data.MultipleCertificateAuthentication.IsNull() {
			data.MultipleCertificateAuthentication = types.BoolValue(value.Bool())
		} else {
			data.MultipleCertificateAuthentication = types.BoolNull()
		}
		if value := res.Get("useLocalAsPrimaryAuthServer"); value.Exists() && !data.PrimaryAuthenticationServerUseLocal.IsNull() {
			data.PrimaryAuthenticationServerUseLocal = types.BoolValue(value.Bool())
		} else {
			data.PrimaryAuthenticationServerUseLocal = types.BoolNull()
		}
		if value := res.Get("primaryAuthenticationServer.id"); value.Exists() && !data.PrimaryAuthenticationServerId.IsNull() {
			data.PrimaryAuthenticationServerId = types.StringValue(value.String())
		} else {
			data.PrimaryAuthenticationServerId = types.StringNull()
		}
		if value := res.Get("primaryAuthenticationServer.type"); value.Exists() && !data.PrimaryAuthenticationServerType.IsNull() {
			data.PrimaryAuthenticationServerType = types.StringValue(value.String())
		} else {
			data.PrimaryAuthenticationServerType = types.StringNull()
		}
		if value := res.Get("enablePrimaryAuthFallbackToLocal"); value.Exists() && !data.PrimaryAuthenticationFallbackToLocal.IsNull() {
			data.PrimaryAuthenticationFallbackToLocal = types.BoolValue(value.Bool())
		} else {
			data.PrimaryAuthenticationFallbackToLocal = types.BoolNull()
		}
		if value := res.Get("certificateUsernameSettings.matchCertificateAndSamlUsername"); value.Exists() && !data.SamlAndCertificateUsernameMustMatch.IsNull() {
			data.SamlAndCertificateUsernameMustMatch = types.BoolValue(value.Bool())
		} else {
			data.SamlAndCertificateUsernameMustMatch = types.BoolNull()
		}
		if value := res.Get("certificateUsernameSettings.prefillUsernameFromCertificate"); value.Exists() && !data.PrimaryAuthenticationPrefillUsernameFromCertificate.IsNull() {
			data.PrimaryAuthenticationPrefillUsernameFromCertificate = types.BoolValue(value.Bool())
		} else {
			data.PrimaryAuthenticationPrefillUsernameFromCertificate = types.BoolNull()
		}
		if value := res.Get("certificateUsernameSettings.mapPrimaryField"); value.Exists() && !data.PrimaryAuthenticationPrefillUsernameFromCertificateMapPrimaryField.IsNull() {
			data.PrimaryAuthenticationPrefillUsernameFromCertificateMapPrimaryField = types.StringValue(value.String())
		} else {
			data.PrimaryAuthenticationPrefillUsernameFromCertificateMapPrimaryField = types.StringNull()
		}
		if value := res.Get("certificateUsernameSettings.mapSecondaryField"); value.Exists() && !data.PrimaryAuthenticationPrefillUsernameFromCertificateMapSecondaryField.IsNull() {
			data.PrimaryAuthenticationPrefillUsernameFromCertificateMapSecondaryField = types.StringValue(value.String())
		} else {
			data.PrimaryAuthenticationPrefillUsernameFromCertificateMapSecondaryField = types.StringNull()
		}
		if value := res.Get("certificateUsernameSettings.mapEntireDNAsUsername"); value.Exists() && !data.PrimaryAuthenticationPrefillUsernameFromCertificateMapEntireDn.IsNull() {
			data.PrimaryAuthenticationPrefillUsernameFromCertificateMapEntireDn = types.BoolValue(value.Bool())
		} else {
			data.PrimaryAuthenticationPrefillUsernameFromCertificateMapEntireDn = types.BoolNull()
		}
		if value := res.Get("certificateUsernameSettings.hideUsername"); value.Exists() && !data.PrimaryAuthenticationHideUsernameInLoginWindow.IsNull() {
			data.PrimaryAuthenticationHideUsernameInLoginWindow = types.BoolValue(value.Bool())
		} else {
			data.PrimaryAuthenticationHideUsernameInLoginWindow = types.BoolNull()
		}
		if value := res.Get("enableSecondaryAuthentication"); value.Exists() && !data.SecondaryAuthentication.IsNull() {
			data.SecondaryAuthentication = types.BoolValue(value.Bool())
		} else {
			data.SecondaryAuthentication = types.BoolNull()
		}
		if value := res.Get("useLocalAsSecondaryAuthServer"); value.Exists() && !data.SecondaryAuthenticationServerUseLocal.IsNull() {
			data.SecondaryAuthenticationServerUseLocal = types.BoolValue(value.Bool())
		} else {
			data.SecondaryAuthenticationServerUseLocal = types.BoolNull()
		}
		if value := res.Get("secondaryAuthenticationServer.id"); value.Exists() && !data.SecondaryAuthenticationServerId.IsNull() {
			data.SecondaryAuthenticationServerId = types.StringValue(value.String())
		} else {
			data.SecondaryAuthenticationServerId = types.StringNull()
		}
		if value := res.Get("secondaryAuthenticationServer.type"); value.Exists() && !data.SecondaryAuthenticationServerType.IsNull() {
			data.SecondaryAuthenticationServerType = types.StringValue(value.String())
		} else {
			data.SecondaryAuthenticationServerType = types.StringNull()
		}
		if value := res.Get("enableSecondaryAuthFallbackToLocal"); value.Exists() && !data.SecondaryAuthenticationFallbackToLocal.IsNull() {
			data.SecondaryAuthenticationFallbackToLocal = types.BoolValue(value.Bool())
		} else {
			data.SecondaryAuthenticationFallbackToLocal = types.BoolNull()
		}
		if value := res.Get("secondaryAuthenticationSettings.promptUsername"); value.Exists() && !data.SecondaryAuthenticationPromptForUsername.IsNull() {
			data.SecondaryAuthenticationPromptForUsername = types.BoolValue(value.Bool())
		} else {
			data.SecondaryAuthenticationPromptForUsername = types.BoolNull()
		}
		if value := res.Get("secondaryAuthenticationSettings.usePrimaryUsername"); value.Exists() && !data.SecondaryAuthenticationUsePrimaryAuthenticationUsername.IsNull() {
			data.SecondaryAuthenticationUsePrimaryAuthenticationUsername = types.BoolValue(value.Bool())
		} else {
			data.SecondaryAuthenticationUsePrimaryAuthenticationUsername = types.BoolNull()
		}
		if value := res.Get("secondaryAuthenticationSettings.useSecondaryUsernameforSession"); value.Exists() && !data.UseSecondaryAuthenticationUsernameForReporting.IsNull() {
			data.UseSecondaryAuthenticationUsernameForReporting = types.BoolValue(value.Bool())
		} else {
			data.UseSecondaryAuthenticationUsernameForReporting = types.BoolNull()
		}
		if value := res.Get("authorizationServer.id"); value.Exists() && !data.AuthorizationServerId.IsNull() {
			data.AuthorizationServerId = types.StringValue(value.String())
		} else {
			data.AuthorizationServerId = types.StringNull()
		}
		if value := res.Get("authorizationServer.type"); value.Exists() && !data.AuthorizationServerType.IsNull() {
			data.AuthorizationServerType = types.StringValue(value.String())
		} else {
			data.AuthorizationServerType = types.StringNull()
		}
		if value := res.Get("allowConnectionOnlyIfAuthorized"); value.Exists() && !data.AllowConnectionOnlyIfUserExistsInAuthorizationDatabase.IsNull() {
			data.AllowConnectionOnlyIfUserExistsInAuthorizationDatabase = types.BoolValue(value.Bool())
		} else {
			data.AllowConnectionOnlyIfUserExistsInAuthorizationDatabase = types.BoolNull()
		}
		if value := res.Get("accountingServer.id"); value.Exists() && !data.AccountingServerId.IsNull() {
			data.AccountingServerId = types.StringValue(value.String())
		} else {
			data.AccountingServerId = types.StringNull()
		}
		if value := res.Get("accountingServer.type"); value.Exists() && !data.AccountingServerType.IsNull() {
			data.AccountingServerType = types.StringValue(value.String())
		} else {
			data.AccountingServerType = types.StringNull()
		}
		if value := res.Get("stripRealmFromUsername"); value.Exists() && !data.StripRealmFromUsername.IsNull() {
			data.StripRealmFromUsername = types.BoolValue(value.Bool())
		} else {
			data.StripRealmFromUsername = types.BoolNull()
		}
		if value := res.Get("stripGroupFromUsername"); value.Exists() && !data.StripGroupFromUsername.IsNull() {
			data.StripGroupFromUsername = types.BoolValue(value.Bool())
		} else {
			data.StripGroupFromUsername = types.BoolNull()
		}
		if value := res.Get("enablePasswordManagement"); value.Exists() && !data.PasswordManagement.IsNull() {
			data.PasswordManagement = types.BoolValue(value.Bool())
		} else {
			data.PasswordManagement = types.BoolNull()
		}
		if value := res.Get("notifyUserOnPasswordExpiryDay"); value.Exists() && !data.PasswordManagementNotifyUserOnPasswordExpiryDay.IsNull() {
			data.PasswordManagementNotifyUserOnPasswordExpiryDay = types.BoolValue(value.Bool())
		} else {
			data.PasswordManagementNotifyUserOnPasswordExpiryDay = types.BoolNull()
		}
		if value := res.Get("passwordExpirationNotificationPeriod"); value.Exists() && !data.PasswordManagementAdvancePasswordExpirationNotification.IsNull() {
			data.PasswordManagementAdvancePasswordExpirationNotification = types.Int64Value(value.Int())
		} else {
			data.PasswordManagementAdvancePasswordExpirationNotification = types.Int64Null()
		}
		for i := 0; i < len(data.AliasNames); i++ {
			keys := [...]string{"aliasName"}
			keyValues := [...]string{data.AliasNames[i].Name.ValueString()}

			parent := &data
			data := (*parent).AliasNames[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("groupAlias").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing AliasNames[%d] = %+v",
					i,
					(*parent).AliasNames[i],
				))
				(*parent).AliasNames = slices.Delete((*parent).AliasNames, i, i+1)
				i--

				continue
			}
			if value := res.Get("aliasName"); value.Exists() && !data.Name.IsNull() {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
				data.Enabled = types.BoolValue(value.Bool())
			} else {
				data.Enabled = types.BoolNull()
			}
			(*parent).AliasNames[i] = data
		}
		for i := 0; i < len(data.AliasUrls); i++ {
			keys := [...]string{"aliasUrl.id"}
			keyValues := [...]string{data.AliasUrls[i].UrlObjectId.ValueString()}

			parent := &data
			data := (*parent).AliasUrls[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("groupUrl").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing AliasUrls[%d] = %+v",
					i,
					(*parent).AliasUrls[i],
				))
				(*parent).AliasUrls = slices.Delete((*parent).AliasUrls, i, i+1)
				i--

				continue
			}
			if value := res.Get("aliasUrl.id"); value.Exists() && !data.UrlObjectId.IsNull() {
				data.UrlObjectId = types.StringValue(value.String())
			} else {
				data.UrlObjectId = types.StringNull()
			}
			if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
				data.Enabled = types.BoolValue(value.Bool())
			} else {
				data.Enabled = types.BoolNull()
			}
			(*parent).AliasUrls[i] = data
		}
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *VPNRAConnectionProfiles) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	for i, val := range data.Items {
		var r gjson.Result
		res.Get("items").ForEach(
			func(_, v gjson.Result) bool {
				if val.Id.IsUnknown() {
					if v.Get("name").String() == i {
						r = v
						return false // break ForEach
					}
				} else {
					if v.Get("id").String() == val.Id.ValueString() && val.Id.ValueString() != "" {
						r = v
						return false // break ForEach
					}
				}

				return true
			},
		)
		if v := data.Items[i]; v.Id.IsUnknown() {
			if value := r.Get("id"); value.Exists() {
				v.Id = types.StringValue(value.String())
			} else {
				v.Id = types.StringNull()
			}
			data.Items[i] = v
		}
		if v := data.Items[i]; v.Type.IsUnknown() {
			if value := r.Get("type"); value.Exists() {
				v.Type = types.StringValue(value.String())
			} else {
				v.Type = types.StringNull()
			}
			data.Items[i] = v
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin Clone

func (data *VPNRAConnectionProfiles) Clone() VPNRAConnectionProfiles {
	ret := *data
	ret.Items = maps.Clone(data.Items)

	return ret
}

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// Updates done one-by-one require different API body
func (data VPNRAConnectionProfiles) toBodyNonBulk(ctx context.Context, state VPNRAConnectionProfiles) string {
	// This is one-by-one update, so only one element to update is expected
	if len(data.Items) > 1 {
		tflog.Error(ctx, "Found more than one element to change. Only one will be changed.")
	}

	// Utilize existing toBody function
	body := data.toBody(ctx, state)

	// Get first element only
	return gjson.Get(body, "0").String()
}

// End of section. //template:end toBodyNonBulk
