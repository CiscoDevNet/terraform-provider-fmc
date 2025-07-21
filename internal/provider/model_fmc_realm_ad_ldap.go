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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type RealmADLDAP struct {
	Id                              types.String                               `tfsdk:"id"`
	Domain                          types.String                               `tfsdk:"domain"`
	Name                            types.String                               `tfsdk:"name"`
	Type                            types.String                               `tfsdk:"type"`
	Version                         types.String                               `tfsdk:"version"`
	Description                     types.String                               `tfsdk:"description"`
	RealmType                       types.String                               `tfsdk:"realm_type"`
	AdPrimaryDomain                 types.String                               `tfsdk:"ad_primary_domain"`
	AdJoinUsername                  types.String                               `tfsdk:"ad_join_username"`
	AdJoinPassword                  types.String                               `tfsdk:"ad_join_password"`
	DirectoryUsername               types.String                               `tfsdk:"directory_username"`
	DirectoryPassword               types.String                               `tfsdk:"directory_password"`
	BaseDn                          types.String                               `tfsdk:"base_dn"`
	GroupDn                         types.String                               `tfsdk:"group_dn"`
	UpdateHour                      types.Int64                                `tfsdk:"update_hour"`
	UpdateInterval                  types.String                               `tfsdk:"update_interval"`
	GroupAttribute                  types.String                               `tfsdk:"group_attribute"`
	TimeoutIseUsers                 types.Int64                                `tfsdk:"timeout_ise_users"`
	TimeoutTerminalServerAgentUsers types.Int64                                `tfsdk:"timeout_terminal_server_agent_users"`
	TimeoutCaptivePortalUsers       types.Int64                                `tfsdk:"timeout_captive_portal_users"`
	TimeoutFailedCaptivePortalUsers types.Int64                                `tfsdk:"timeout_failed_captive_portal_users"`
	TimeoutGuestCaptivePortalUsers  types.Int64                                `tfsdk:"timeout_guest_captive_portal_users"`
	DirectoryServerConfigurations   []RealmADLDAPDirectoryServerConfigurations `tfsdk:"directory_server_configurations"`
}

type RealmADLDAPDirectoryServerConfigurations struct {
	Hostname                    types.String `tfsdk:"hostname"`
	Port                        types.Int64  `tfsdk:"port"`
	EncryptionProtocol          types.String `tfsdk:"encryption_protocol"`
	EncryptionCertificate       types.String `tfsdk:"encryption_certificate"`
	UseRoutingToSelectInterface types.Bool   `tfsdk:"use_routing_to_select_interface"`
	InterfaceGroupId            types.String `tfsdk:"interface_group_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data RealmADLDAP) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/realms"
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
	if !data.Version.IsNull() && !data.Version.IsUnknown() {
		body, _ = sjson.Set(body, "version", data.Version.ValueString())
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
	if !data.AdJoinUsername.IsNull() {
		body, _ = sjson.Set(body, "adJoinUsername", data.AdJoinUsername.ValueString())
	}
	if !data.AdJoinPassword.IsNull() {
		body, _ = sjson.Set(body, "adJoinPassword", data.AdJoinPassword.ValueString())
	}
	if !data.DirectoryUsername.IsNull() {
		body, _ = sjson.Set(body, "dirUsername", data.DirectoryUsername.ValueString())
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
	if !data.UpdateHour.IsNull() {
		body, _ = sjson.Set(body, "updateHour", data.UpdateHour.ValueInt64())
	}
	if !data.UpdateInterval.IsNull() {
		body, _ = sjson.Set(body, "updateInterval", data.UpdateInterval.ValueString())
	}
	if !data.GroupAttribute.IsNull() {
		body, _ = sjson.Set(body, "groupAttribute", data.GroupAttribute.ValueString())
	}
	if !data.TimeoutIseUsers.IsNull() {
		body, _ = sjson.Set(body, "authSessionTimeout", data.TimeoutIseUsers.ValueInt64())
	}
	if !data.TimeoutTerminalServerAgentUsers.IsNull() {
		body, _ = sjson.Set(body, "tsAgentSessionTimeout", data.TimeoutTerminalServerAgentUsers.ValueInt64())
	}
	if !data.TimeoutCaptivePortalUsers.IsNull() {
		body, _ = sjson.Set(body, "activeAuthSessionTimeout", data.TimeoutCaptivePortalUsers.ValueInt64())
	}
	if !data.TimeoutFailedCaptivePortalUsers.IsNull() {
		body, _ = sjson.Set(body, "failedAuthSessionTimeout", data.TimeoutFailedCaptivePortalUsers.ValueInt64())
	}
	if !data.TimeoutGuestCaptivePortalUsers.IsNull() {
		body, _ = sjson.Set(body, "guestSessionTimeout", data.TimeoutGuestCaptivePortalUsers.ValueInt64())
	}
	if len(data.DirectoryServerConfigurations) > 0 {
		body, _ = sjson.Set(body, "directoryConfigurations", []interface{}{})
		for _, item := range data.DirectoryServerConfigurations {
			itemBody := ""
			if !item.Hostname.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "hostname", item.Hostname.ValueString())
			}
			if !item.Port.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "port", item.Port.ValueInt64())
			}
			if !item.EncryptionProtocol.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "encryptionProtocol", item.EncryptionProtocol.ValueString())
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
	if value := res.Get("version"); value.Exists() {
		data.Version = types.StringValue(value.String())
	} else {
		data.Version = types.StringNull()
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
		data.AdJoinUsername = types.StringValue(value.String())
	} else {
		data.AdJoinUsername = types.StringNull()
	}
	if value := res.Get("dirUsername"); value.Exists() {
		data.DirectoryUsername = types.StringValue(value.String())
	} else {
		data.DirectoryUsername = types.StringNull()
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
	if value := res.Get("updateHour"); value.Exists() {
		data.UpdateHour = types.Int64Value(value.Int())
	} else {
		data.UpdateHour = types.Int64Null()
	}
	if value := res.Get("updateInterval"); value.Exists() {
		data.UpdateInterval = types.StringValue(value.String())
	} else {
		data.UpdateInterval = types.StringNull()
	}
	if value := res.Get("groupAttribute"); value.Exists() {
		data.GroupAttribute = types.StringValue(value.String())
	} else {
		data.GroupAttribute = types.StringNull()
	}
	if value := res.Get("authSessionTimeout"); value.Exists() {
		data.TimeoutIseUsers = types.Int64Value(value.Int())
	} else {
		data.TimeoutIseUsers = types.Int64Null()
	}
	if value := res.Get("tsAgentSessionTimeout"); value.Exists() {
		data.TimeoutTerminalServerAgentUsers = types.Int64Value(value.Int())
	} else {
		data.TimeoutTerminalServerAgentUsers = types.Int64Null()
	}
	if value := res.Get("activeAuthSessionTimeout"); value.Exists() {
		data.TimeoutCaptivePortalUsers = types.Int64Value(value.Int())
	} else {
		data.TimeoutCaptivePortalUsers = types.Int64Null()
	}
	if value := res.Get("failedAuthSessionTimeout"); value.Exists() {
		data.TimeoutFailedCaptivePortalUsers = types.Int64Value(value.Int())
	} else {
		data.TimeoutFailedCaptivePortalUsers = types.Int64Null()
	}
	if value := res.Get("guestSessionTimeout"); value.Exists() {
		data.TimeoutGuestCaptivePortalUsers = types.Int64Value(value.Int())
	} else {
		data.TimeoutGuestCaptivePortalUsers = types.Int64Null()
	}
	if value := res.Get("directoryConfigurations"); value.Exists() {
		data.DirectoryServerConfigurations = make([]RealmADLDAPDirectoryServerConfigurations, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := RealmADLDAPDirectoryServerConfigurations{}
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
			if value := res.Get("encryptionProtocol"); value.Exists() {
				data.EncryptionProtocol = types.StringValue(value.String())
			} else {
				data.EncryptionProtocol = types.StringNull()
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
			(*parent).DirectoryServerConfigurations = append((*parent).DirectoryServerConfigurations, data)
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
	if value := res.Get("version"); value.Exists() && !data.Version.IsNull() {
		data.Version = types.StringValue(value.String())
	} else {
		data.Version = types.StringNull()
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
	if value := res.Get("adJoinUsername"); value.Exists() && !data.AdJoinUsername.IsNull() {
		data.AdJoinUsername = types.StringValue(value.String())
	} else {
		data.AdJoinUsername = types.StringNull()
	}
	if value := res.Get("dirUsername"); value.Exists() && !data.DirectoryUsername.IsNull() {
		data.DirectoryUsername = types.StringValue(value.String())
	} else {
		data.DirectoryUsername = types.StringNull()
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
	if value := res.Get("updateHour"); value.Exists() && !data.UpdateHour.IsNull() {
		data.UpdateHour = types.Int64Value(value.Int())
	} else {
		data.UpdateHour = types.Int64Null()
	}
	if value := res.Get("updateInterval"); value.Exists() && !data.UpdateInterval.IsNull() {
		data.UpdateInterval = types.StringValue(value.String())
	} else {
		data.UpdateInterval = types.StringNull()
	}
	if value := res.Get("groupAttribute"); value.Exists() && !data.GroupAttribute.IsNull() {
		data.GroupAttribute = types.StringValue(value.String())
	} else {
		data.GroupAttribute = types.StringNull()
	}
	if value := res.Get("authSessionTimeout"); value.Exists() && !data.TimeoutIseUsers.IsNull() {
		data.TimeoutIseUsers = types.Int64Value(value.Int())
	} else {
		data.TimeoutIseUsers = types.Int64Null()
	}
	if value := res.Get("tsAgentSessionTimeout"); value.Exists() && !data.TimeoutTerminalServerAgentUsers.IsNull() {
		data.TimeoutTerminalServerAgentUsers = types.Int64Value(value.Int())
	} else {
		data.TimeoutTerminalServerAgentUsers = types.Int64Null()
	}
	if value := res.Get("activeAuthSessionTimeout"); value.Exists() && !data.TimeoutCaptivePortalUsers.IsNull() {
		data.TimeoutCaptivePortalUsers = types.Int64Value(value.Int())
	} else {
		data.TimeoutCaptivePortalUsers = types.Int64Null()
	}
	if value := res.Get("failedAuthSessionTimeout"); value.Exists() && !data.TimeoutFailedCaptivePortalUsers.IsNull() {
		data.TimeoutFailedCaptivePortalUsers = types.Int64Value(value.Int())
	} else {
		data.TimeoutFailedCaptivePortalUsers = types.Int64Null()
	}
	if value := res.Get("guestSessionTimeout"); value.Exists() && !data.TimeoutGuestCaptivePortalUsers.IsNull() {
		data.TimeoutGuestCaptivePortalUsers = types.Int64Value(value.Int())
	} else {
		data.TimeoutGuestCaptivePortalUsers = types.Int64Null()
	}
	for i := 0; i < len(data.DirectoryServerConfigurations); i++ {
		keys := [...]string{"hostname"}
		keyValues := [...]string{data.DirectoryServerConfigurations[i].Hostname.ValueString()}

		parent := &data
		data := (*parent).DirectoryServerConfigurations[i]
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
			tflog.Debug(ctx, fmt.Sprintf("removing DirectoryServerConfigurations[%d] = %+v",
				i,
				(*parent).DirectoryServerConfigurations[i],
			))
			(*parent).DirectoryServerConfigurations = slices.Delete((*parent).DirectoryServerConfigurations, i, i+1)
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
		if value := res.Get("encryptionProtocol"); value.Exists() && !data.EncryptionProtocol.IsNull() {
			data.EncryptionProtocol = types.StringValue(value.String())
		} else {
			data.EncryptionProtocol = types.StringNull()
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
		(*parent).DirectoryServerConfigurations[i] = data
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
	if data.Version.IsUnknown() {
		if value := res.Get("version"); value.Exists() {
			data.Version = types.StringValue(value.String())
		} else {
			data.Version = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

func (data RealmADLDAP) adjustBody(ctx context.Context, req string) string {
	// Add sequence numbers to the entities
	for i := range len(data.DirectoryServerConfigurations) {
		req, _ = sjson.Set(req, fmt.Sprintf("directoryConfigurations.%d.id", i), i+1)
	}

	return req
}
