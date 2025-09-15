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

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type FTDPlatformSettingsTrustedDNSServers struct {
	Id                        types.String                                                   `tfsdk:"id"`
	Domain                    types.String                                                   `tfsdk:"domain"`
	FtdPlatformSettingsId     types.String                                                   `tfsdk:"ftd_platform_settings_id"`
	Type                      types.String                                                   `tfsdk:"type"`
	TrustAnyDnsServer         types.Bool                                                     `tfsdk:"trust_any_dns_server"`
	TrustDhcpPool             types.Bool                                                     `tfsdk:"trust_dhcp_pool"`
	TrustDhcpRelay            types.Bool                                                     `tfsdk:"trust_dhcp_relay"`
	TrustDhcpClient           types.Bool                                                     `tfsdk:"trust_dhcp_client"`
	TrustDnsServerGroup       types.Bool                                                     `tfsdk:"trust_dns_server_group"`
	TrustedDnsServersLiterals types.Set                                                      `tfsdk:"trusted_dns_servers_literals"`
	TrustedDnsServersObjects  []FTDPlatformSettingsTrustedDNSServersTrustedDnsServersObjects `tfsdk:"trusted_dns_servers_objects"`
}

type FTDPlatformSettingsTrustedDNSServersTrustedDnsServersObjects struct {
	Id types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionCreateFTDPlatformSettingsTrustedDNSServers = version.Must(version.NewVersion("7.7"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data FTDPlatformSettingsTrustedDNSServers) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/ftdplatformsettingspolicies/%v/trusteddnssettings", url.QueryEscape(data.FtdPlatformSettingsId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data FTDPlatformSettingsTrustedDNSServers) toBody(ctx context.Context, state FTDPlatformSettingsTrustedDNSServers) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.TrustAnyDnsServer.IsNull() {
		body, _ = sjson.Set(body, "trustAnyDNSServer", data.TrustAnyDnsServer.ValueBool())
	}
	if !data.TrustDhcpPool.IsNull() {
		body, _ = sjson.Set(body, "trustDhcpPool", data.TrustDhcpPool.ValueBool())
	}
	if !data.TrustDhcpRelay.IsNull() {
		body, _ = sjson.Set(body, "trustDhcpRelay", data.TrustDhcpRelay.ValueBool())
	}
	if !data.TrustDhcpClient.IsNull() {
		body, _ = sjson.Set(body, "trustDhcpClient", data.TrustDhcpClient.ValueBool())
	}
	if !data.TrustDnsServerGroup.IsNull() {
		body, _ = sjson.Set(body, "trustDNSServerGroup", data.TrustDnsServerGroup.ValueBool())
	}
	if !data.TrustedDnsServersLiterals.IsNull() {
		var values []string
		data.TrustedDnsServersLiterals.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "dnsServers.literals", values)
	}
	if len(data.TrustedDnsServersObjects) > 0 {
		body, _ = sjson.Set(body, "dnsServers.objects", []any{})
		for _, item := range data.TrustedDnsServersObjects {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "dnsServers.objects.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *FTDPlatformSettingsTrustedDNSServers) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("trustAnyDNSServer"); value.Exists() {
		data.TrustAnyDnsServer = types.BoolValue(value.Bool())
	} else {
		data.TrustAnyDnsServer = types.BoolNull()
	}
	if value := res.Get("trustDhcpPool"); value.Exists() {
		data.TrustDhcpPool = types.BoolValue(value.Bool())
	} else {
		data.TrustDhcpPool = types.BoolNull()
	}
	if value := res.Get("trustDhcpRelay"); value.Exists() {
		data.TrustDhcpRelay = types.BoolValue(value.Bool())
	} else {
		data.TrustDhcpRelay = types.BoolNull()
	}
	if value := res.Get("trustDhcpClient"); value.Exists() {
		data.TrustDhcpClient = types.BoolValue(value.Bool())
	} else {
		data.TrustDhcpClient = types.BoolNull()
	}
	if value := res.Get("trustDNSServerGroup"); value.Exists() {
		data.TrustDnsServerGroup = types.BoolValue(value.Bool())
	} else {
		data.TrustDnsServerGroup = types.BoolNull()
	}
	if value := res.Get("dnsServers.literals"); value.Exists() {
		data.TrustedDnsServersLiterals = helpers.GetStringSet(value.Array())
	} else {
		data.TrustedDnsServersLiterals = types.SetNull(types.StringType)
	}
	if value := res.Get("dnsServers.objects"); value.Exists() {
		data.TrustedDnsServersObjects = make([]FTDPlatformSettingsTrustedDNSServersTrustedDnsServersObjects, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := FTDPlatformSettingsTrustedDNSServersTrustedDnsServersObjects{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).TrustedDnsServersObjects = append((*parent).TrustedDnsServersObjects, data)
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
func (data *FTDPlatformSettingsTrustedDNSServers) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("trustAnyDNSServer"); value.Exists() && !data.TrustAnyDnsServer.IsNull() {
		data.TrustAnyDnsServer = types.BoolValue(value.Bool())
	} else {
		data.TrustAnyDnsServer = types.BoolNull()
	}
	if value := res.Get("trustDhcpPool"); value.Exists() && !data.TrustDhcpPool.IsNull() {
		data.TrustDhcpPool = types.BoolValue(value.Bool())
	} else {
		data.TrustDhcpPool = types.BoolNull()
	}
	if value := res.Get("trustDhcpRelay"); value.Exists() && !data.TrustDhcpRelay.IsNull() {
		data.TrustDhcpRelay = types.BoolValue(value.Bool())
	} else {
		data.TrustDhcpRelay = types.BoolNull()
	}
	if value := res.Get("trustDhcpClient"); value.Exists() && !data.TrustDhcpClient.IsNull() {
		data.TrustDhcpClient = types.BoolValue(value.Bool())
	} else {
		data.TrustDhcpClient = types.BoolNull()
	}
	if value := res.Get("trustDNSServerGroup"); value.Exists() && !data.TrustDnsServerGroup.IsNull() {
		data.TrustDnsServerGroup = types.BoolValue(value.Bool())
	} else {
		data.TrustDnsServerGroup = types.BoolNull()
	}
	if value := res.Get("dnsServers.literals"); value.Exists() && !data.TrustedDnsServersLiterals.IsNull() {
		data.TrustedDnsServersLiterals = helpers.GetStringSet(value.Array())
	} else {
		data.TrustedDnsServersLiterals = types.SetNull(types.StringType)
	}
	for i := 0; i < len(data.TrustedDnsServersObjects); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.TrustedDnsServersObjects[i].Id.ValueString()}

		parent := &data
		data := (*parent).TrustedDnsServersObjects[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("dnsServers.objects").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing TrustedDnsServersObjects[%d] = %+v",
				i,
				(*parent).TrustedDnsServersObjects[i],
			))
			(*parent).TrustedDnsServersObjects = slices.Delete((*parent).TrustedDnsServersObjects, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).TrustedDnsServersObjects[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *FTDPlatformSettingsTrustedDNSServers) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyPutDelete

// toBodyPutDelete is used to create the body for PUT requests to clear the resource state
func (data FTDPlatformSettingsTrustedDNSServers) toBodyPutDelete(ctx context.Context) string {
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
