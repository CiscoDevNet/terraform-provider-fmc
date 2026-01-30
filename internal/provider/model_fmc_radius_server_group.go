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

type RadiusServerGroup struct {
	Id                               types.String                     `tfsdk:"id"`
	Domain                           types.String                     `tfsdk:"domain"`
	Name                             types.String                     `tfsdk:"name"`
	Type                             types.String                     `tfsdk:"type"`
	Description                      types.String                     `tfsdk:"description"`
	GroupAccountingMode              types.String                     `tfsdk:"group_accounting_mode"`
	RetryInterval                    types.Int64                      `tfsdk:"retry_interval"`
	AdRealmId                        types.String                     `tfsdk:"ad_realm_id"`
	AuthorizeOnly                    types.Bool                       `tfsdk:"authorize_only"`
	InterimAccountUpdateInterval     types.Int64                      `tfsdk:"interim_account_update_interval"`
	DynamicAuthorization             types.Bool                       `tfsdk:"dynamic_authorization"`
	DynamicAuthorizationPort         types.Int64                      `tfsdk:"dynamic_authorization_port"`
	MergeDownloadableAccessListOrder types.String                     `tfsdk:"merge_downloadable_access_list_order"`
	RadiusServers                    []RadiusServerGroupRadiusServers `tfsdk:"radius_servers"`
}

type RadiusServerGroupRadiusServers struct {
	Hostname                    types.String `tfsdk:"hostname"`
	MessageAuthenticator        types.Bool   `tfsdk:"message_authenticator"`
	AuthenticationPort          types.Int64  `tfsdk:"authentication_port"`
	Key                         types.String `tfsdk:"key"`
	AccountingPort              types.Int64  `tfsdk:"accounting_port"`
	Timeout                     types.Int64  `tfsdk:"timeout"`
	UseRoutingToSelectInterface types.Bool   `tfsdk:"use_routing_to_select_interface"`
	InterfaceId                 types.String `tfsdk:"interface_id"`
	RedirectAccessListId        types.String `tfsdk:"redirect_access_list_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data RadiusServerGroup) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/radiusservergroups"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data RadiusServerGroup) toBody(ctx context.Context, state RadiusServerGroup) string {
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
	if !data.GroupAccountingMode.IsNull() {
		body, _ = sjson.Set(body, "groupAccountingMode", data.GroupAccountingMode.ValueString())
	}
	if !data.RetryInterval.IsNull() {
		body, _ = sjson.Set(body, "retryInterval", data.RetryInterval.ValueInt64())
	}
	if !data.AdRealmId.IsNull() {
		body, _ = sjson.Set(body, "realm.id", data.AdRealmId.ValueString())
	}
	if !data.AuthorizeOnly.IsNull() {
		body, _ = sjson.Set(body, "enableAuthorizeOnly", data.AuthorizeOnly.ValueBool())
	}
	if !data.InterimAccountUpdateInterval.IsNull() {
		body, _ = sjson.Set(body, "interimAccountUpdateInterval", data.InterimAccountUpdateInterval.ValueInt64())
	}
	if !data.DynamicAuthorization.IsNull() {
		body, _ = sjson.Set(body, "enableDynamicAuthorization", data.DynamicAuthorization.ValueBool())
	}
	if !data.DynamicAuthorizationPort.IsNull() {
		body, _ = sjson.Set(body, "dynamicAuthorizationPort", data.DynamicAuthorizationPort.ValueInt64())
	}
	if !data.MergeDownloadableAccessListOrder.IsNull() {
		body, _ = sjson.Set(body, "mergeDaclPlacementOrder", data.MergeDownloadableAccessListOrder.ValueString())
	}
	if len(data.RadiusServers) > 0 {
		body, _ = sjson.Set(body, "radiusServers", []any{})
		for _, item := range data.RadiusServers {
			itemBody := ""
			if !item.Hostname.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "host", item.Hostname.ValueString())
			}
			if !item.MessageAuthenticator.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enableMessageAuthenticator", item.MessageAuthenticator.ValueBool())
			}
			if !item.AuthenticationPort.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "authenticationPort", item.AuthenticationPort.ValueInt64())
			}
			if !item.Key.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "secretKey", item.Key.ValueString())
			}
			if !item.AccountingPort.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "accountingPort", item.AccountingPort.ValueInt64())
			}
			if !item.Timeout.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "timeout", item.Timeout.ValueInt64())
			}
			if !item.UseRoutingToSelectInterface.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "useRoutingToSelectInterface", item.UseRoutingToSelectInterface.ValueBool())
			}
			if !item.InterfaceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "interface.id", item.InterfaceId.ValueString())
			}
			if !item.RedirectAccessListId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "redirectACL.id", item.RedirectAccessListId.ValueString())
			}
			body, _ = sjson.SetRaw(body, "radiusServers.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *RadiusServerGroup) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("groupAccountingMode"); value.Exists() {
		data.GroupAccountingMode = types.StringValue(value.String())
	} else {
		data.GroupAccountingMode = types.StringValue("SINGLE")
	}
	if value := res.Get("retryInterval"); value.Exists() {
		data.RetryInterval = types.Int64Value(value.Int())
	} else {
		data.RetryInterval = types.Int64Value(10)
	}
	if value := res.Get("realm.id"); value.Exists() {
		data.AdRealmId = types.StringValue(value.String())
	} else {
		data.AdRealmId = types.StringNull()
	}
	if value := res.Get("enableAuthorizeOnly"); value.Exists() {
		data.AuthorizeOnly = types.BoolValue(value.Bool())
	} else {
		data.AuthorizeOnly = types.BoolNull()
	}
	if value := res.Get("interimAccountUpdateInterval"); value.Exists() {
		data.InterimAccountUpdateInterval = types.Int64Value(value.Int())
	} else {
		data.InterimAccountUpdateInterval = types.Int64Null()
	}
	if value := res.Get("enableDynamicAuthorization"); value.Exists() {
		data.DynamicAuthorization = types.BoolValue(value.Bool())
	} else {
		data.DynamicAuthorization = types.BoolNull()
	}
	if value := res.Get("dynamicAuthorizationPort"); value.Exists() {
		data.DynamicAuthorizationPort = types.Int64Value(value.Int())
	} else {
		data.DynamicAuthorizationPort = types.Int64Null()
	}
	if value := res.Get("mergeDaclPlacementOrder"); value.Exists() {
		data.MergeDownloadableAccessListOrder = types.StringValue(value.String())
	} else {
		data.MergeDownloadableAccessListOrder = types.StringNull()
	}
	if value := res.Get("radiusServers"); value.Exists() {
		data.RadiusServers = make([]RadiusServerGroupRadiusServers, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := RadiusServerGroupRadiusServers{}
			if value := res.Get("host"); value.Exists() {
				data.Hostname = types.StringValue(value.String())
			} else {
				data.Hostname = types.StringNull()
			}
			if value := res.Get("enableMessageAuthenticator"); value.Exists() {
				data.MessageAuthenticator = types.BoolValue(value.Bool())
			} else {
				data.MessageAuthenticator = types.BoolValue(true)
			}
			if value := res.Get("authenticationPort"); value.Exists() {
				data.AuthenticationPort = types.Int64Value(value.Int())
			} else {
				data.AuthenticationPort = types.Int64Value(1812)
			}
			if value := res.Get("accountingPort"); value.Exists() {
				data.AccountingPort = types.Int64Value(value.Int())
			} else {
				data.AccountingPort = types.Int64Value(1813)
			}
			if value := res.Get("timeout"); value.Exists() {
				data.Timeout = types.Int64Value(value.Int())
			} else {
				data.Timeout = types.Int64Value(10)
			}
			if value := res.Get("useRoutingToSelectInterface"); value.Exists() {
				data.UseRoutingToSelectInterface = types.BoolValue(value.Bool())
			} else {
				data.UseRoutingToSelectInterface = types.BoolValue(true)
			}
			if value := res.Get("interface.id"); value.Exists() {
				data.InterfaceId = types.StringValue(value.String())
			} else {
				data.InterfaceId = types.StringNull()
			}
			if value := res.Get("redirectACL.id"); value.Exists() {
				data.RedirectAccessListId = types.StringValue(value.String())
			} else {
				data.RedirectAccessListId = types.StringNull()
			}
			(*parent).RadiusServers = append((*parent).RadiusServers, data)
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
func (data *RadiusServerGroup) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("groupAccountingMode"); value.Exists() && !data.GroupAccountingMode.IsNull() {
		data.GroupAccountingMode = types.StringValue(value.String())
	} else if data.GroupAccountingMode.ValueString() != "SINGLE" {
		data.GroupAccountingMode = types.StringNull()
	}
	if value := res.Get("retryInterval"); value.Exists() && !data.RetryInterval.IsNull() {
		data.RetryInterval = types.Int64Value(value.Int())
	} else if data.RetryInterval.ValueInt64() != 10 {
		data.RetryInterval = types.Int64Null()
	}
	if value := res.Get("realm.id"); value.Exists() && !data.AdRealmId.IsNull() {
		data.AdRealmId = types.StringValue(value.String())
	} else {
		data.AdRealmId = types.StringNull()
	}
	if value := res.Get("enableAuthorizeOnly"); value.Exists() && !data.AuthorizeOnly.IsNull() {
		data.AuthorizeOnly = types.BoolValue(value.Bool())
	} else {
		data.AuthorizeOnly = types.BoolNull()
	}
	if value := res.Get("interimAccountUpdateInterval"); value.Exists() && !data.InterimAccountUpdateInterval.IsNull() {
		data.InterimAccountUpdateInterval = types.Int64Value(value.Int())
	} else {
		data.InterimAccountUpdateInterval = types.Int64Null()
	}
	if value := res.Get("enableDynamicAuthorization"); value.Exists() && !data.DynamicAuthorization.IsNull() {
		data.DynamicAuthorization = types.BoolValue(value.Bool())
	} else {
		data.DynamicAuthorization = types.BoolNull()
	}
	if value := res.Get("dynamicAuthorizationPort"); value.Exists() && !data.DynamicAuthorizationPort.IsNull() {
		data.DynamicAuthorizationPort = types.Int64Value(value.Int())
	} else {
		data.DynamicAuthorizationPort = types.Int64Null()
	}
	if value := res.Get("mergeDaclPlacementOrder"); value.Exists() && !data.MergeDownloadableAccessListOrder.IsNull() {
		data.MergeDownloadableAccessListOrder = types.StringValue(value.String())
	} else {
		data.MergeDownloadableAccessListOrder = types.StringNull()
	}
	for i := 0; i < len(data.RadiusServers); i++ {
		keys := [...]string{"host"}
		keyValues := [...]string{data.RadiusServers[i].Hostname.ValueString()}

		parent := &data
		data := (*parent).RadiusServers[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("radiusServers").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing RadiusServers[%d] = %+v",
				i,
				(*parent).RadiusServers[i],
			))
			(*parent).RadiusServers = slices.Delete((*parent).RadiusServers, i, i+1)
			i--

			continue
		}
		if value := res.Get("host"); value.Exists() && !data.Hostname.IsNull() {
			data.Hostname = types.StringValue(value.String())
		} else {
			data.Hostname = types.StringNull()
		}
		if value := res.Get("enableMessageAuthenticator"); value.Exists() && !data.MessageAuthenticator.IsNull() {
			data.MessageAuthenticator = types.BoolValue(value.Bool())
		} else if data.MessageAuthenticator.ValueBool() != true {
			data.MessageAuthenticator = types.BoolNull()
		}
		if value := res.Get("authenticationPort"); value.Exists() && !data.AuthenticationPort.IsNull() {
			data.AuthenticationPort = types.Int64Value(value.Int())
		} else if data.AuthenticationPort.ValueInt64() != 1812 {
			data.AuthenticationPort = types.Int64Null()
		}
		if value := res.Get("accountingPort"); value.Exists() && !data.AccountingPort.IsNull() {
			data.AccountingPort = types.Int64Value(value.Int())
		} else if data.AccountingPort.ValueInt64() != 1813 {
			data.AccountingPort = types.Int64Null()
		}
		if value := res.Get("timeout"); value.Exists() && !data.Timeout.IsNull() {
			data.Timeout = types.Int64Value(value.Int())
		} else if data.Timeout.ValueInt64() != 10 {
			data.Timeout = types.Int64Null()
		}
		if value := res.Get("useRoutingToSelectInterface"); value.Exists() && !data.UseRoutingToSelectInterface.IsNull() {
			data.UseRoutingToSelectInterface = types.BoolValue(value.Bool())
		} else if data.UseRoutingToSelectInterface.ValueBool() != true {
			data.UseRoutingToSelectInterface = types.BoolNull()
		}
		if value := res.Get("interface.id"); value.Exists() && !data.InterfaceId.IsNull() {
			data.InterfaceId = types.StringValue(value.String())
		} else {
			data.InterfaceId = types.StringNull()
		}
		if value := res.Get("redirectACL.id"); value.Exists() && !data.RedirectAccessListId.IsNull() {
			data.RedirectAccessListId = types.StringValue(value.String())
		} else {
			data.RedirectAccessListId = types.StringNull()
		}
		(*parent).RadiusServers[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *RadiusServerGroup) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

func (data RadiusServerGroup) adjustBody(ctx context.Context, req string) string {

	if !data.InterimAccountUpdateInterval.IsNull() {
		req, _ = sjson.Set(req, "enableInterimAccountUpdate", "true")
	}

	// API would still return port, even if the dynamic authorization is not enabled.
	// if !data.DynamicAuthorizationPort.IsNull() {
	// 	req, _ = sjson.Set(req, "enableDynamicAuthorization", "true")
	// }

	if !data.DynamicAuthorization.IsUnknown() && !data.DynamicAuthorization.ValueBool() {
		req, _ = sjson.Delete(req, "dynamicAuthorizationPort")
	}

	if !data.MergeDownloadableAccessListOrder.IsNull() {
		req, _ = sjson.Set(req, "enableMergeDacl", "true")
	}

	return req
}
