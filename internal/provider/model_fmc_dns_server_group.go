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

type DNSServerGroup struct {
	Id            types.String               `tfsdk:"id"`
	Domain        types.String               `tfsdk:"domain"`
	Name          types.String               `tfsdk:"name"`
	Type          types.String               `tfsdk:"type"`
	DefaultDomain types.String               `tfsdk:"default_domain"`
	Timeout       types.Int64                `tfsdk:"timeout"`
	Retries       types.Int64                `tfsdk:"retries"`
	DnsServers    []DNSServerGroupDnsServers `tfsdk:"dns_servers"`
}

type DNSServerGroupDnsServers struct {
	Ip types.String `tfsdk:"ip"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionCreateDNSServerGroup = version.Must(version.NewVersion("7.4"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DNSServerGroup) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/dnsservergroups"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DNSServerGroup) toBody(ctx context.Context, state DNSServerGroup) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	body, _ = sjson.Set(body, "type", "DNSServerGroupObject")
	if !data.DefaultDomain.IsNull() {
		body, _ = sjson.Set(body, "defaultdomain", data.DefaultDomain.ValueString())
	}
	if !data.Timeout.IsNull() {
		body, _ = sjson.Set(body, "timeout", data.Timeout.ValueInt64())
	}
	if !data.Retries.IsNull() {
		body, _ = sjson.Set(body, "retries", data.Retries.ValueInt64())
	}
	if len(data.DnsServers) > 0 {
		body, _ = sjson.Set(body, "dnsservers", []any{})
		for _, item := range data.DnsServers {
			itemBody := ""
			if !item.Ip.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name-server", item.Ip.ValueString())
			}
			body, _ = sjson.SetRaw(body, "dnsservers.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DNSServerGroup) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("defaultdomain"); value.Exists() {
		data.DefaultDomain = types.StringValue(value.String())
	} else {
		data.DefaultDomain = types.StringNull()
	}
	if value := res.Get("timeout"); value.Exists() {
		data.Timeout = types.Int64Value(value.Int())
	} else {
		data.Timeout = types.Int64Null()
	}
	if value := res.Get("retries"); value.Exists() {
		data.Retries = types.Int64Value(value.Int())
	} else {
		data.Retries = types.Int64Null()
	}
	if value := res.Get("dnsservers"); value.Exists() {
		data.DnsServers = make([]DNSServerGroupDnsServers, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DNSServerGroupDnsServers{}
			if value := res.Get("name-server"); value.Exists() {
				data.Ip = types.StringValue(value.String())
			} else {
				data.Ip = types.StringNull()
			}
			(*parent).DnsServers = append((*parent).DnsServers, data)
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
func (data *DNSServerGroup) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("defaultdomain"); value.Exists() && !data.DefaultDomain.IsNull() {
		data.DefaultDomain = types.StringValue(value.String())
	} else {
		data.DefaultDomain = types.StringNull()
	}
	if value := res.Get("timeout"); value.Exists() && !data.Timeout.IsNull() {
		data.Timeout = types.Int64Value(value.Int())
	} else {
		data.Timeout = types.Int64Null()
	}
	if value := res.Get("retries"); value.Exists() && !data.Retries.IsNull() {
		data.Retries = types.Int64Value(value.Int())
	} else {
		data.Retries = types.Int64Null()
	}
	for i := 0; i < len(data.DnsServers); i++ {
		keys := [...]string{"name-server"}
		keyValues := [...]string{data.DnsServers[i].Ip.ValueString()}

		parent := &data
		data := (*parent).DnsServers[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("dnsservers").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing DnsServers[%d] = %+v",
				i,
				(*parent).DnsServers[i],
			))
			(*parent).DnsServers = slices.Delete((*parent).DnsServers, i, i+1)
			i--

			continue
		}
		if value := res.Get("name-server"); value.Exists() && !data.Ip.IsNull() {
			data.Ip = types.StringValue(value.String())
		} else {
			data.Ip = types.StringNull()
		}
		(*parent).DnsServers[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DNSServerGroup) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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

// End of section. //template:end toBodyPutDelete

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBody

// End of section. //template:end adjustBody

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBodyBulk

// End of section. //template:end adjustBodyBulk
