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

type VPNRALDAPAttributeMap struct {
	Id      types.String                  `tfsdk:"id"`
	Domain  types.String                  `tfsdk:"domain"`
	VpnRaId types.String                  `tfsdk:"vpn_ra_id"`
	Type    types.String                  `tfsdk:"type"`
	Realms  []VPNRALDAPAttributeMapRealms `tfsdk:"realms"`
}

type VPNRALDAPAttributeMapRealms struct {
	RealmAdLdapId types.String                               `tfsdk:"realm_ad_ldap_id"`
	AttributeMaps []VPNRALDAPAttributeMapRealmsAttributeMaps `tfsdk:"attribute_maps"`
}

type VPNRALDAPAttributeMapRealmsAttributeMaps struct {
	LdapAttributeName  types.String                                        `tfsdk:"ldap_attribute_name"`
	CiscoAttributeName types.String                                        `tfsdk:"cisco_attribute_name"`
	ValueMaps          []VPNRALDAPAttributeMapRealmsAttributeMapsValueMaps `tfsdk:"value_maps"`
}

type VPNRALDAPAttributeMapRealmsAttributeMapsValueMaps struct {
	LdapValue  types.String `tfsdk:"ldap_value"`
	CiscoValue types.String `tfsdk:"cisco_value"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data VPNRALDAPAttributeMap) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/ravpns/%v/ldapattributemaps", url.QueryEscape(data.VpnRaId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data VPNRALDAPAttributeMap) toBody(ctx context.Context, state VPNRALDAPAttributeMap) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	body, _ = sjson.Set(body, "type", "RaVpnLdapAttributeMap")
	if len(data.Realms) > 0 {
		body, _ = sjson.Set(body, "ldapAttributeMapList", []any{})
		for _, item := range data.Realms {
			itemBody := ""
			if !item.RealmAdLdapId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ldapServer.id", item.RealmAdLdapId.ValueString())
			}
			if len(item.AttributeMaps) > 0 {
				itemBody, _ = sjson.Set(itemBody, "ldapAttributeMaps", []any{})
				for _, childItem := range item.AttributeMaps {
					itemChildBody := ""
					if !childItem.LdapAttributeName.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "ldapName", childItem.LdapAttributeName.ValueString())
					}
					if !childItem.CiscoAttributeName.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "ciscoName", childItem.CiscoAttributeName.ValueString())
					}
					if len(childItem.ValueMaps) > 0 {
						itemChildBody, _ = sjson.Set(itemChildBody, "valueMappings", []any{})
						for _, childChildItem := range childItem.ValueMaps {
							itemChildChildBody := ""
							if !childChildItem.LdapValue.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "ldapValue", childChildItem.LdapValue.ValueString())
							}
							if !childChildItem.CiscoValue.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "ciscoValue", childChildItem.CiscoValue.ValueString())
							}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "valueMappings.-1", itemChildChildBody)
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "ldapAttributeMaps.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "ldapAttributeMapList.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *VPNRALDAPAttributeMap) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("ldapAttributeMapList"); value.Exists() {
		data.Realms = make([]VPNRALDAPAttributeMapRealms, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := VPNRALDAPAttributeMapRealms{}
			if value := res.Get("ldapServer.id"); value.Exists() {
				data.RealmAdLdapId = types.StringValue(value.String())
			} else {
				data.RealmAdLdapId = types.StringNull()
			}
			if value := res.Get("ldapAttributeMaps"); value.Exists() {
				data.AttributeMaps = make([]VPNRALDAPAttributeMapRealmsAttributeMaps, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := VPNRALDAPAttributeMapRealmsAttributeMaps{}
					if value := res.Get("ldapName"); value.Exists() {
						data.LdapAttributeName = types.StringValue(value.String())
					} else {
						data.LdapAttributeName = types.StringNull()
					}
					if value := res.Get("ciscoName"); value.Exists() {
						data.CiscoAttributeName = types.StringValue(value.String())
					} else {
						data.CiscoAttributeName = types.StringNull()
					}
					if value := res.Get("valueMappings"); value.Exists() {
						data.ValueMaps = make([]VPNRALDAPAttributeMapRealmsAttributeMapsValueMaps, 0)
						value.ForEach(func(k, res gjson.Result) bool {
							parent := &data
							data := VPNRALDAPAttributeMapRealmsAttributeMapsValueMaps{}
							if value := res.Get("ldapValue"); value.Exists() {
								data.LdapValue = types.StringValue(value.String())
							} else {
								data.LdapValue = types.StringNull()
							}
							if value := res.Get("ciscoValue"); value.Exists() {
								data.CiscoValue = types.StringValue(value.String())
							} else {
								data.CiscoValue = types.StringNull()
							}
							(*parent).ValueMaps = append((*parent).ValueMaps, data)
							return true
						})
					}
					(*parent).AttributeMaps = append((*parent).AttributeMaps, data)
					return true
				})
			}
			(*parent).Realms = append((*parent).Realms, data)
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
func (data *VPNRALDAPAttributeMap) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	for i := 0; i < len(data.Realms); i++ {
		keys := [...]string{"ldapServer.id"}
		keyValues := [...]string{data.Realms[i].RealmAdLdapId.ValueString()}

		parent := &data
		data := (*parent).Realms[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("ldapAttributeMapList").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Realms[%d] = %+v",
				i,
				(*parent).Realms[i],
			))
			(*parent).Realms = slices.Delete((*parent).Realms, i, i+1)
			i--

			continue
		}
		if value := res.Get("ldapServer.id"); value.Exists() && !data.RealmAdLdapId.IsNull() {
			data.RealmAdLdapId = types.StringValue(value.String())
		} else {
			data.RealmAdLdapId = types.StringNull()
		}
		for i := 0; i < len(data.AttributeMaps); i++ {
			keys := [...]string{"ldapName", "ciscoName"}
			keyValues := [...]string{data.AttributeMaps[i].LdapAttributeName.ValueString(), data.AttributeMaps[i].CiscoAttributeName.ValueString()}

			parent := &data
			data := (*parent).AttributeMaps[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("ldapAttributeMaps").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing AttributeMaps[%d] = %+v",
					i,
					(*parent).AttributeMaps[i],
				))
				(*parent).AttributeMaps = slices.Delete((*parent).AttributeMaps, i, i+1)
				i--

				continue
			}
			if value := res.Get("ldapName"); value.Exists() && !data.LdapAttributeName.IsNull() {
				data.LdapAttributeName = types.StringValue(value.String())
			} else {
				data.LdapAttributeName = types.StringNull()
			}
			if value := res.Get("ciscoName"); value.Exists() && !data.CiscoAttributeName.IsNull() {
				data.CiscoAttributeName = types.StringValue(value.String())
			} else {
				data.CiscoAttributeName = types.StringNull()
			}
			for i := 0; i < len(data.ValueMaps); i++ {
				keys := [...]string{"ldapValue", "ciscoValue"}
				keyValues := [...]string{data.ValueMaps[i].LdapValue.ValueString(), data.ValueMaps[i].CiscoValue.ValueString()}

				parent := &data
				data := (*parent).ValueMaps[i]
				parentRes := &res
				var res gjson.Result

				parentRes.Get("valueMappings").ForEach(
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
					tflog.Debug(ctx, fmt.Sprintf("removing ValueMaps[%d] = %+v",
						i,
						(*parent).ValueMaps[i],
					))
					(*parent).ValueMaps = slices.Delete((*parent).ValueMaps, i, i+1)
					i--

					continue
				}
				if value := res.Get("ldapValue"); value.Exists() && !data.LdapValue.IsNull() {
					data.LdapValue = types.StringValue(value.String())
				} else {
					data.LdapValue = types.StringNull()
				}
				if value := res.Get("ciscoValue"); value.Exists() && !data.CiscoValue.IsNull() {
					data.CiscoValue = types.StringValue(value.String())
				} else {
					data.CiscoValue = types.StringNull()
				}
				(*parent).ValueMaps[i] = data
			}
			(*parent).AttributeMaps[i] = data
		}
		(*parent).Realms[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *VPNRALDAPAttributeMap) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
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

// toBodyPutDelete is used to create the body for PUT requests to clear the resource state
func (data VPNRALDAPAttributeMap) toBodyPutDelete(ctx context.Context) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if data.Type.ValueString() != "" {
		body, _ = sjson.Set(body, "type", data.Type.ValueString())
	}
	return body
}

// End of section. //template:end toBodyPutDelete

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBody

// End of section. //template:end adjustBody

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBodyBulk

// End of section. //template:end adjustBodyBulk
