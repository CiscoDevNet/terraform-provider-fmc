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
	"slices"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DNSServerGroups struct {
	Id     types.String                    `tfsdk:"id"`
	Domain types.String                    `tfsdk:"domain"`
	Items  map[string]DNSServerGroupsItems `tfsdk:"items"`
}

type DNSServerGroupsItems struct {
	Id            types.String                     `tfsdk:"id"`
	Type          types.String                     `tfsdk:"type"`
	DefaultDomain types.String                     `tfsdk:"default_domain"`
	Timeout       types.Int64                      `tfsdk:"timeout"`
	Retries       types.Int64                      `tfsdk:"retries"`
	DnsServers    []DNSServerGroupsItemsDnsServers `tfsdk:"dns_servers"`
}

type DNSServerGroupsItemsDnsServers struct {
	Ip types.String `tfsdk:"ip"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionBulkCreateDNSServerGroups = version.Must(version.NewVersion("999"))
var minFMCVersionBulkDeleteDNSServerGroups = version.Must(version.NewVersion("999"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DNSServerGroups) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/dnsservergroups"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DNSServerGroups) toBody(ctx context.Context, state DNSServerGroups) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if len(data.Items) > 0 {
		body, _ = sjson.Set(body, "items", []interface{}{})
		for key, item := range data.Items {
			itemBody, _ := sjson.Set("{}", "name", key)
			if !item.Id.IsNull() && !item.Id.IsUnknown() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			itemBody, _ = sjson.Set(itemBody, "type", "DNSServerGroupObject")
			if !item.DefaultDomain.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "defaultdomain", item.DefaultDomain.ValueString())
			}
			if !item.Timeout.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "timeout", item.Timeout.ValueInt64())
			}
			if !item.Retries.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "retries", item.Retries.ValueInt64())
			}
			if len(item.DnsServers) > 0 {
				itemBody, _ = sjson.Set(itemBody, "dnsservers", []interface{}{})
				for _, childItem := range item.DnsServers {
					itemChildBody := ""
					if !childItem.Ip.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "name-server", childItem.Ip.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "dnsservers.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "items.-1", itemBody)
		}
	}
	return gjson.Get(body, "items").String()
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DNSServerGroups) fromBody(ctx context.Context, res gjson.Result) {
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
			data.DnsServers = make([]DNSServerGroupsItemsDnsServers, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DNSServerGroupsItemsDnsServers{}
				if value := res.Get("name-server"); value.Exists() {
					data.Ip = types.StringValue(value.String())
				} else {
					data.Ip = types.StringNull()
				}
				(*parent).DnsServers = append((*parent).DnsServers, data)
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
func (data *DNSServerGroups) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DNSServerGroups) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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

func (data *DNSServerGroups) Clone() DNSServerGroups {
	ret := *data
	ret.Items = maps.Clone(data.Items)

	return ret
}

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// Updates done one-by-one require different API body
func (data DNSServerGroups) toBodyNonBulk(ctx context.Context, state DNSServerGroups) string {
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
