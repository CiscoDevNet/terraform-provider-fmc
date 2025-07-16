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

type RealmADLDAP struct {
	Id                      types.String                         `tfsdk:"id"`
	Domain                  types.String                         `tfsdk:"domain"`
	Name                    types.String                         `tfsdk:"name"`
	Type                    types.String                         `tfsdk:"type"`
	Description             types.String                         `tfsdk:"description"`
	RealmType               types.String                         `tfsdk:"realm_type"`
	AdPrimaryDomain         types.String                         `tfsdk:"ad_primary_domain"`
	DirectoryUsername       types.String                         `tfsdk:"directory_username"`
	DirectoryPassword       types.String                         `tfsdk:"directory_password"`
	BaseDn                  types.String                         `tfsdk:"base_dn"`
	GroupDn                 types.String                         `tfsdk:"group_dn"`
	DirectoryConfigurations []RealmADLDAPDirectoryConfigurations `tfsdk:"directory_configurations"`
}

type RealmADLDAPDirectoryConfigurations struct {
	Hostname                    types.String `tfsdk:"hostname"`
	Port                        types.Int64  `tfsdk:"port"`
	Encryption                  types.String `tfsdk:"encryption"`
	EncryptionCertificate       types.String `tfsdk:"encryption_certificate"`
	UseRoutingToSelectInterface types.Bool   `tfsdk:"use_routing_to_select_interface"`
	InterfaceGroupId            types.String `tfsdk:"interface_group_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data RealmADLDAP) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/hosts"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data RealmADLDAP) toBody(ctx context.Context, state RealmADLDAP) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.RealmType.IsNull() {
		body, _ = sjson.Set(body, "realmType", data.RealmType.ValueString())
	}
	if !data.AdPrimaryDomain.IsNull() {
		body, _ = sjson.Set(body, "adPrimaryDomain", data.AdPrimaryDomain.ValueString())
	}
	if !data.DirectoryUsername.IsNull() {
		body, _ = sjson.Set(body, "adJoinUsername", data.DirectoryUsername.ValueString())
	}
	if !data.DirectoryPassword.IsNull() {
		body, _ = sjson.Set(body, "dirPassword", data.DirectoryPassword.ValueString())
	}
	if !data.BaseDn.IsNull() {
		body, _ = sjson.Set(body, "baseDn", data.BaseDn.ValueString())
	}
	if !data.GroupDn.IsNull() {
		body, _ = sjson.Set(body, "groupDn", data.GroupDn.ValueString())
	}
	if len(data.DirectoryConfigurations) > 0 {
		body, _ = sjson.Set(body, "directoryConfigurations", []interface{}{})
		for _, item := range data.DirectoryConfigurations {
			itemBody := ""
			if !item.Hostname.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "hostname", item.Hostname.ValueString())
			}
			if !item.Port.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "port", item.Port.ValueInt64())
			}
			if !item.Encryption.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "encryption", item.Encryption.ValueString())
			}
			if !item.EncryptionCertificate.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "encryptionCert.id", item.EncryptionCertificate.ValueString())
			}
			if !item.UseRoutingToSelectInterface.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "useRoutingToSelectInterface", item.UseRoutingToSelectInterface.ValueBool())
			}
			if !item.InterfaceGroupId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "interface.id", item.InterfaceGroupId.ValueString())
			}
			body, _ = sjson.SetRaw(body, "directoryConfigurations.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *RealmADLDAP) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("realmType"); value.Exists() {
		data.RealmType = types.StringValue(value.String())
	} else {
		data.RealmType = types.StringNull()
	}
	if value := res.Get("adPrimaryDomain"); value.Exists() {
		data.AdPrimaryDomain = types.StringValue(value.String())
	} else {
		data.AdPrimaryDomain = types.StringNull()
	}
	if value := res.Get("adJoinUsername"); value.Exists() {
		data.DirectoryUsername = types.StringValue(value.String())
	} else {
		data.DirectoryUsername = types.StringNull()
	}
	if value := res.Get("dirPassword"); value.Exists() {
		data.DirectoryPassword = types.StringValue(value.String())
	} else {
		data.DirectoryPassword = types.StringNull()
	}
	if value := res.Get("baseDn"); value.Exists() {
		data.BaseDn = types.StringValue(value.String())
	} else {
		data.BaseDn = types.StringNull()
	}
	if value := res.Get("groupDn"); value.Exists() {
		data.GroupDn = types.StringValue(value.String())
	} else {
		data.GroupDn = types.StringNull()
	}
	if value := res.Get("directoryConfigurations"); value.Exists() {
		data.DirectoryConfigurations = make([]RealmADLDAPDirectoryConfigurations, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := RealmADLDAPDirectoryConfigurations{}
			if value := res.Get("hostname"); value.Exists() {
				data.Hostname = types.StringValue(value.String())
			} else {
				data.Hostname = types.StringNull()
			}
			if value := res.Get("port"); value.Exists() {
				data.Port = types.Int64Value(value.Int())
			} else {
				data.Port = types.Int64Null()
			}
			if value := res.Get("encryption"); value.Exists() {
				data.Encryption = types.StringValue(value.String())
			} else {
				data.Encryption = types.StringNull()
			}
			if value := res.Get("encryptionCert.id"); value.Exists() {
				data.EncryptionCertificate = types.StringValue(value.String())
			} else {
				data.EncryptionCertificate = types.StringNull()
			}
			if value := res.Get("useRoutingToSelectInterface"); value.Exists() {
				data.UseRoutingToSelectInterface = types.BoolValue(value.Bool())
			} else {
				data.UseRoutingToSelectInterface = types.BoolNull()
			}
			if value := res.Get("interface.id"); value.Exists() {
				data.InterfaceGroupId = types.StringValue(value.String())
			} else {
				data.InterfaceGroupId = types.StringNull()
			}
			(*parent).DirectoryConfigurations = append((*parent).DirectoryConfigurations, data)
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
func (data *RealmADLDAP) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("realmType"); value.Exists() && !data.RealmType.IsNull() {
		data.RealmType = types.StringValue(value.String())
	} else {
		data.RealmType = types.StringNull()
	}
	if value := res.Get("adPrimaryDomain"); value.Exists() && !data.AdPrimaryDomain.IsNull() {
		data.AdPrimaryDomain = types.StringValue(value.String())
	} else {
		data.AdPrimaryDomain = types.StringNull()
	}
	if value := res.Get("adJoinUsername"); value.Exists() && !data.DirectoryUsername.IsNull() {
		data.DirectoryUsername = types.StringValue(value.String())
	} else {
		data.DirectoryUsername = types.StringNull()
	}
	if value := res.Get("dirPassword"); value.Exists() && !data.DirectoryPassword.IsNull() {
		data.DirectoryPassword = types.StringValue(value.String())
	} else {
		data.DirectoryPassword = types.StringNull()
	}
	if value := res.Get("baseDn"); value.Exists() && !data.BaseDn.IsNull() {
		data.BaseDn = types.StringValue(value.String())
	} else {
		data.BaseDn = types.StringNull()
	}
	if value := res.Get("groupDn"); value.Exists() && !data.GroupDn.IsNull() {
		data.GroupDn = types.StringValue(value.String())
	} else {
		data.GroupDn = types.StringNull()
	}
	for i := 0; i < len(data.DirectoryConfigurations); i++ {
		keys := [...]string{"hostname", "port", "encryption", "encryptionCert.id", "useRoutingToSelectInterface", "interface.id"}
		keyValues := [...]string{data.DirectoryConfigurations[i].Hostname.ValueString(), strconv.FormatInt(data.DirectoryConfigurations[i].Port.ValueInt64(), 10), data.DirectoryConfigurations[i].Encryption.ValueString(), data.DirectoryConfigurations[i].EncryptionCertificate.ValueString(), strconv.FormatBool(data.DirectoryConfigurations[i].UseRoutingToSelectInterface.ValueBool()), data.DirectoryConfigurations[i].InterfaceGroupId.ValueString()}

		parent := &data
		data := (*parent).DirectoryConfigurations[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("directoryConfigurations").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing DirectoryConfigurations[%d] = %+v",
				i,
				(*parent).DirectoryConfigurations[i],
			))
			(*parent).DirectoryConfigurations = slices.Delete((*parent).DirectoryConfigurations, i, i+1)
			i--

			continue
		}
		if value := res.Get("hostname"); value.Exists() && !data.Hostname.IsNull() {
			data.Hostname = types.StringValue(value.String())
		} else {
			data.Hostname = types.StringNull()
		}
		if value := res.Get("port"); value.Exists() && !data.Port.IsNull() {
			data.Port = types.Int64Value(value.Int())
		} else {
			data.Port = types.Int64Null()
		}
		if value := res.Get("encryption"); value.Exists() && !data.Encryption.IsNull() {
			data.Encryption = types.StringValue(value.String())
		} else {
			data.Encryption = types.StringNull()
		}
		if value := res.Get("encryptionCert.id"); value.Exists() && !data.EncryptionCertificate.IsNull() {
			data.EncryptionCertificate = types.StringValue(value.String())
		} else {
			data.EncryptionCertificate = types.StringNull()
		}
		if value := res.Get("useRoutingToSelectInterface"); value.Exists() && !data.UseRoutingToSelectInterface.IsNull() {
			data.UseRoutingToSelectInterface = types.BoolValue(value.Bool())
		} else {
			data.UseRoutingToSelectInterface = types.BoolNull()
		}
		if value := res.Get("interface.id"); value.Exists() && !data.InterfaceGroupId.IsNull() {
			data.InterfaceGroupId = types.StringValue(value.String())
		} else {
			data.InterfaceGroupId = types.StringNull()
		}
		(*parent).DirectoryConfigurations[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *RealmADLDAP) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
