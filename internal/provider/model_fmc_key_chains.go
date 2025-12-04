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
	"strconv"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type KeyChains struct {
	Id     types.String              `tfsdk:"id"`
	Domain types.String              `tfsdk:"domain"`
	Items  map[string]KeyChainsItems `tfsdk:"items"`
}

type KeyChainsItems struct {
	Id          types.String         `tfsdk:"id"`
	Type        types.String         `tfsdk:"type"`
	Description types.String         `tfsdk:"description"`
	Keys        []KeyChainsItemsKeys `tfsdk:"keys"`
}

type KeyChainsItemsKeys struct {
	Id                    types.Int64  `tfsdk:"id"`
	Key                   types.String `tfsdk:"key"`
	AcceptLifetimeStart   types.String `tfsdk:"accept_lifetime_start"`
	AcceptLifetimeEndType types.String `tfsdk:"accept_lifetime_end_type"`
	AcceptLifetimeEnd     types.String `tfsdk:"accept_lifetime_end"`
	SendLifetimeStart     types.String `tfsdk:"send_lifetime_start"`
	SendLifetimeEndType   types.String `tfsdk:"send_lifetime_end_type"`
	SendLifetimeEnd       types.String `tfsdk:"send_lifetime_end"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionBulkDeleteKeyChains = version.Must(version.NewVersion("999"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data KeyChains) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/keychains"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data KeyChains) toBody(ctx context.Context, state KeyChains) string {
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
			itemBody, _ = sjson.Set(itemBody, "type", "KeyChainObject")
			if !item.Description.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "description", item.Description.ValueString())
			}
			if len(item.Keys) > 0 {
				itemBody, _ = sjson.Set(itemBody, "keys", []any{})
				for _, childItem := range item.Keys {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "keyId", childItem.Id.ValueInt64())
					}
					if !childItem.Key.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "authString.cryptoKeyString", childItem.Key.ValueString())
					}
					itemChildBody, _ = sjson.Set(itemChildBody, "authString.cryptoEncryptionType", "PLAINTEXT")
					itemChildBody, _ = sjson.Set(itemChildBody, "authAlgorithm", "md5")
					if !childItem.AcceptLifetimeStart.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "acceptLifeTime.startLifeTimeValue", childItem.AcceptLifetimeStart.ValueString())
					}
					if !childItem.AcceptLifetimeEndType.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "acceptLifeTime.endLifetimeType", childItem.AcceptLifetimeEndType.ValueString())
					}
					if !childItem.AcceptLifetimeEnd.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "acceptLifeTime.endLifeTimeValue", childItem.AcceptLifetimeEnd.ValueString())
					}
					if !childItem.SendLifetimeStart.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "sendLifeTime.startLifeTimeValue", childItem.SendLifetimeStart.ValueString())
					}
					if !childItem.SendLifetimeEndType.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "sendLifeTime.endLifetimeType", childItem.SendLifetimeEndType.ValueString())
					}
					if !childItem.SendLifetimeEnd.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "sendLifeTime.endLifeTimeValue", childItem.SendLifetimeEnd.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "keys.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "items.-1", itemBody)
		}
	}
	return gjson.Get(body, "items").String()
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *KeyChains) fromBody(ctx context.Context, res gjson.Result) {
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
		if value := res.Get("description"); value.Exists() {
			data.Description = types.StringValue(value.String())
		} else {
			data.Description = types.StringNull()
		}
		if value := res.Get("keys"); value.Exists() {
			data.Keys = make([]KeyChainsItemsKeys, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := KeyChainsItemsKeys{}
				if value := res.Get("keyId"); value.Exists() {
					data.Id = types.Int64Value(value.Int())
				} else {
					data.Id = types.Int64Null()
				}
				if value := res.Get("authString.cryptoKeyString"); value.Exists() {
					data.Key = types.StringValue(value.String())
				} else {
					data.Key = types.StringNull()
				}
				if value := res.Get("acceptLifeTime.startLifeTimeValue"); value.Exists() {
					data.AcceptLifetimeStart = types.StringValue(value.String())
				} else {
					data.AcceptLifetimeStart = types.StringNull()
				}
				if value := res.Get("acceptLifeTime.endLifetimeType"); value.Exists() {
					data.AcceptLifetimeEndType = types.StringValue(value.String())
				} else {
					data.AcceptLifetimeEndType = types.StringNull()
				}
				if value := res.Get("acceptLifeTime.endLifeTimeValue"); value.Exists() {
					data.AcceptLifetimeEnd = types.StringValue(value.String())
				} else {
					data.AcceptLifetimeEnd = types.StringNull()
				}
				if value := res.Get("sendLifeTime.startLifeTimeValue"); value.Exists() {
					data.SendLifetimeStart = types.StringValue(value.String())
				} else {
					data.SendLifetimeStart = types.StringNull()
				}
				if value := res.Get("sendLifeTime.endLifetimeType"); value.Exists() {
					data.SendLifetimeEndType = types.StringValue(value.String())
				} else {
					data.SendLifetimeEndType = types.StringNull()
				}
				if value := res.Get("sendLifeTime.endLifeTimeValue"); value.Exists() {
					data.SendLifetimeEnd = types.StringValue(value.String())
				} else {
					data.SendLifetimeEnd = types.StringNull()
				}
				(*parent).Keys = append((*parent).Keys, data)
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
func (data *KeyChains) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
		if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
			data.Description = types.StringValue(value.String())
		} else {
			if !data.Description.IsNull() && data.Description.ValueString() == "" {
				data.Description = types.StringValue("")
			} else {
				data.Description = types.StringNull()
			}
		}
		for i := 0; i < len(data.Keys); i++ {
			keys := [...]string{"keyId"}
			keyValues := [...]string{strconv.FormatInt(data.Keys[i].Id.ValueInt64(), 10)}

			parent := &data
			data := (*parent).Keys[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("keys").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing Keys[%d] = %+v",
					i,
					(*parent).Keys[i],
				))
				(*parent).Keys = slices.Delete((*parent).Keys, i, i+1)
				i--

				continue
			}
			if value := res.Get("keyId"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.Int64Value(value.Int())
			} else {
				data.Id = types.Int64Null()
			}
			if value := res.Get("authString.cryptoKeyString"); value.Exists() && !data.Key.IsNull() {
				data.Key = types.StringValue(value.String())
			} else {
				data.Key = types.StringNull()
			}
			if value := res.Get("acceptLifeTime.startLifeTimeValue"); value.Exists() && !data.AcceptLifetimeStart.IsNull() {
				data.AcceptLifetimeStart = types.StringValue(value.String())
			} else {
				data.AcceptLifetimeStart = types.StringNull()
			}
			if value := res.Get("acceptLifeTime.endLifetimeType"); value.Exists() && !data.AcceptLifetimeEndType.IsNull() {
				data.AcceptLifetimeEndType = types.StringValue(value.String())
			} else {
				data.AcceptLifetimeEndType = types.StringNull()
			}
			if value := res.Get("acceptLifeTime.endLifeTimeValue"); value.Exists() && !data.AcceptLifetimeEnd.IsNull() {
				data.AcceptLifetimeEnd = types.StringValue(value.String())
			} else {
				data.AcceptLifetimeEnd = types.StringNull()
			}
			if value := res.Get("sendLifeTime.startLifeTimeValue"); value.Exists() && !data.SendLifetimeStart.IsNull() {
				data.SendLifetimeStart = types.StringValue(value.String())
			} else {
				data.SendLifetimeStart = types.StringNull()
			}
			if value := res.Get("sendLifeTime.endLifetimeType"); value.Exists() && !data.SendLifetimeEndType.IsNull() {
				data.SendLifetimeEndType = types.StringValue(value.String())
			} else {
				data.SendLifetimeEndType = types.StringNull()
			}
			if value := res.Get("sendLifeTime.endLifeTimeValue"); value.Exists() && !data.SendLifetimeEnd.IsNull() {
				data.SendLifetimeEnd = types.StringValue(value.String())
			} else {
				data.SendLifetimeEnd = types.StringNull()
			}
			(*parent).Keys[i] = data
		}
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *KeyChains) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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

func (data *KeyChains) Clone() KeyChains {
	ret := *data
	ret.Items = maps.Clone(data.Items)

	return ret
}

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// Updates done one-by-one require different API body
func (data KeyChains) toBodyNonBulk(ctx context.Context, state KeyChains) string {
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
