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

type VPNRACertificateMap struct {
	Id                                        types.String                                                `tfsdk:"id"`
	Domain                                    types.String                                                `tfsdk:"domain"`
	VpnRaId                                   types.String                                                `tfsdk:"vpn_ra_id"`
	Type                                      types.String                                                `tfsdk:"type"`
	UseGroupUrl                               types.Bool                                                  `tfsdk:"use_group_url"`
	UseCertificateToConnectionProfileMappings types.Bool                                                  `tfsdk:"use_certificate_to_connection_profile_mappings"`
	CertificateToConnectionProfileMappings    []VPNRACertificateMapCertificateToConnectionProfileMappings `tfsdk:"certificate_to_connection_profile_mappings"`
}

type VPNRACertificateMapCertificateToConnectionProfileMappings struct {
	CertificateMapId    types.String `tfsdk:"certificate_map_id"`
	ConnectionProfileId types.String `tfsdk:"connection_profile_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data VPNRACertificateMap) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/ravpns/%v/certificatemapsettings", url.QueryEscape(data.VpnRaId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data VPNRACertificateMap) toBody(ctx context.Context, state VPNRACertificateMap) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	body, _ = sjson.Set(body, "type", "RaVpnCertificateMapSetting")
	if !data.UseGroupUrl.IsNull() {
		body, _ = sjson.Set(body, "useGroupURL", data.UseGroupUrl.ValueBool())
	}
	if !data.UseCertificateToConnectionProfileMappings.IsNull() {
		body, _ = sjson.Set(body, "enableCertificateToConnectionProfileMapping", data.UseCertificateToConnectionProfileMappings.ValueBool())
	}
	if len(data.CertificateToConnectionProfileMappings) > 0 {
		body, _ = sjson.Set(body, "certificateToConnectionProfileMap", []interface{}{})
		for _, item := range data.CertificateToConnectionProfileMappings {
			itemBody := ""
			if !item.CertificateMapId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "certificateMap.id", item.CertificateMapId.ValueString())
			}
			if !item.ConnectionProfileId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "connectionProfile.id", item.ConnectionProfileId.ValueString())
			}
			body, _ = sjson.SetRaw(body, "certificateToConnectionProfileMap.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *VPNRACertificateMap) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("useGroupURL"); value.Exists() {
		data.UseGroupUrl = types.BoolValue(value.Bool())
	} else {
		data.UseGroupUrl = types.BoolNull()
	}
	if value := res.Get("enableCertificateToConnectionProfileMapping"); value.Exists() {
		data.UseCertificateToConnectionProfileMappings = types.BoolValue(value.Bool())
	} else {
		data.UseCertificateToConnectionProfileMappings = types.BoolNull()
	}
	if value := res.Get("certificateToConnectionProfileMap"); value.Exists() {
		data.CertificateToConnectionProfileMappings = make([]VPNRACertificateMapCertificateToConnectionProfileMappings, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := VPNRACertificateMapCertificateToConnectionProfileMappings{}
			if value := res.Get("certificateMap.id"); value.Exists() {
				data.CertificateMapId = types.StringValue(value.String())
			} else {
				data.CertificateMapId = types.StringNull()
			}
			if value := res.Get("connectionProfile.id"); value.Exists() {
				data.ConnectionProfileId = types.StringValue(value.String())
			} else {
				data.ConnectionProfileId = types.StringNull()
			}
			(*parent).CertificateToConnectionProfileMappings = append((*parent).CertificateToConnectionProfileMappings, data)
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
func (data *VPNRACertificateMap) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("useGroupURL"); value.Exists() && !data.UseGroupUrl.IsNull() {
		data.UseGroupUrl = types.BoolValue(value.Bool())
	} else {
		data.UseGroupUrl = types.BoolNull()
	}
	if value := res.Get("enableCertificateToConnectionProfileMapping"); value.Exists() && !data.UseCertificateToConnectionProfileMappings.IsNull() {
		data.UseCertificateToConnectionProfileMappings = types.BoolValue(value.Bool())
	} else {
		data.UseCertificateToConnectionProfileMappings = types.BoolNull()
	}
	for i := 0; i < len(data.CertificateToConnectionProfileMappings); i++ {
		keys := [...]string{"certificateMap.id"}
		keyValues := [...]string{data.CertificateToConnectionProfileMappings[i].CertificateMapId.ValueString()}

		parent := &data
		data := (*parent).CertificateToConnectionProfileMappings[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("certificateToConnectionProfileMap").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing CertificateToConnectionProfileMappings[%d] = %+v",
				i,
				(*parent).CertificateToConnectionProfileMappings[i],
			))
			(*parent).CertificateToConnectionProfileMappings = slices.Delete((*parent).CertificateToConnectionProfileMappings, i, i+1)
			i--

			continue
		}
		if value := res.Get("certificateMap.id"); value.Exists() && !data.CertificateMapId.IsNull() {
			data.CertificateMapId = types.StringValue(value.String())
		} else {
			data.CertificateMapId = types.StringNull()
		}
		if value := res.Get("connectionProfile.id"); value.Exists() && !data.ConnectionProfileId.IsNull() {
			data.ConnectionProfileId = types.StringValue(value.String())
		} else {
			data.ConnectionProfileId = types.StringNull()
		}
		(*parent).CertificateToConnectionProfileMappings[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *VPNRACertificateMap) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// toBodyPutDelete is used to create the body for PUT requests to clear the resource state
func (data VPNRACertificateMap) toBodyPutDelete(ctx context.Context) string {
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
