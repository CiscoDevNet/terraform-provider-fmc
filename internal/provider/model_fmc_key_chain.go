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
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type KeyChain struct {
	Id          types.String   `tfsdk:"id"`
	Domain      types.String   `tfsdk:"domain"`
	Name        types.String   `tfsdk:"name"`
	Type        types.String   `tfsdk:"type"`
	Description types.String   `tfsdk:"description"`
	Keys        []KeyChainKeys `tfsdk:"keys"`
}

type KeyChainKeys struct {
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

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data KeyChain) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/keychains"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data KeyChain) toBody(ctx context.Context, state KeyChain) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	body, _ = sjson.Set(body, "type", "KeyChainObject")
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if len(data.Keys) > 0 {
		body, _ = sjson.Set(body, "keys", []any{})
		for _, item := range data.Keys {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "keyId", item.Id.ValueInt64())
			}
			if !item.Key.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "authString.cryptoKeyString", item.Key.ValueString())
			}
			itemBody, _ = sjson.Set(itemBody, "authString.cryptoEncryptionType", "PLAINTEXT")
			itemBody, _ = sjson.Set(itemBody, "authAlgorithm", "md5")
			if !item.AcceptLifetimeStart.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "acceptLifeTime.startLifeTimeValue", item.AcceptLifetimeStart.ValueString())
			}
			if !item.AcceptLifetimeEndType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "acceptLifeTime.endLifetimeType", item.AcceptLifetimeEndType.ValueString())
			}
			if !item.AcceptLifetimeEnd.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "acceptLifeTime.endLifeTimeValue", item.AcceptLifetimeEnd.ValueString())
			}
			if !item.SendLifetimeStart.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sendLifeTime.startLifeTimeValue", item.SendLifetimeStart.ValueString())
			}
			if !item.SendLifetimeEndType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sendLifeTime.endLifetimeType", item.SendLifetimeEndType.ValueString())
			}
			if !item.SendLifetimeEnd.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sendLifeTime.endLifeTimeValue", item.SendLifetimeEnd.ValueString())
			}
			body, _ = sjson.SetRaw(body, "keys.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *KeyChain) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("keys"); value.Exists() {
		data.Keys = make([]KeyChainKeys, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := KeyChainKeys{}
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
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *KeyChain) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *KeyChain) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
