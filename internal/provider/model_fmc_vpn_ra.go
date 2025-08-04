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

type VPNRA struct {
	Id                                           types.String              `tfsdk:"id"`
	Domain                                       types.String              `tfsdk:"domain"`
	Name                                         types.String              `tfsdk:"name"`
	Description                                  types.String              `tfsdk:"description"`
	Type                                         types.String              `tfsdk:"type"`
	ProtocolSsl                                  types.Bool                `tfsdk:"protocol_ssl"`
	ProtocolIpsecIkev2                           types.Bool                `tfsdk:"protocol_ipsec_ikev2"`
	LocalRealmId                                 types.String              `tfsdk:"local_realm_id"`
	DapPolicyId                                  types.String              `tfsdk:"dap_policy_id"`
	AccessInterfaces                             []VPNRAAccessInterfaces   `tfsdk:"access_interfaces"`
	AllowUsersToSelectConnectionProfile          types.Bool                `tfsdk:"allow_users_to_select_connection_profile"`
	WebAccessPortNumber                          types.Int64               `tfsdk:"web_access_port_number"`
	DtlsPortNumber                               types.Int64               `tfsdk:"dtls_port_number"`
	SslGlobalIdentityCertificateId               types.String              `tfsdk:"ssl_global_identity_certificate_id"`
	IpsecIkev2IdentityCertificateId              types.String              `tfsdk:"ipsec_ikev2_identity_certificate_id"`
	ServiceAccessObjectId                        types.String              `tfsdk:"service_access_object_id"`
	BypassAccessControlPolicyForDecryptedTraffic types.Bool                `tfsdk:"bypass_access_control_policy_for_decrypted_traffic"`
	SecureClientImages                           []VPNRASecureClientImages `tfsdk:"secure_client_images"`
	ExternalBrowserPackageId                     types.String              `tfsdk:"external_browser_package_id"`
	SecureClientCustomizationId                  types.String              `tfsdk:"secure_client_customization_id"`
	AddressAssignmentPolicyId                    types.String              `tfsdk:"address_assignment_policy_id"`
	CertificateMapId                             types.String              `tfsdk:"certificate_map_id"`
	GroupPolicies                                []VPNRAGroupPolicies      `tfsdk:"group_policies"`
	LdapAttributeMapId                           types.String              `tfsdk:"ldap_attribute_map_id"`
	LoadBalancingId                              types.String              `tfsdk:"load_balancing_id"`
	Ikev2Policies                                []VPNRAIkev2Policies      `tfsdk:"ikev2_policies"`
	IpsecIkeParametersId                         types.String              `tfsdk:"ipsec_ike_parameters_id"`
}

type VPNRAAccessInterfaces struct {
	Id                             types.String `tfsdk:"id"`
	ProtocolIpsecIkev2             types.Bool   `tfsdk:"protocol_ipsec_ikev2"`
	ProtocolSsl                    types.Bool   `tfsdk:"protocol_ssl"`
	ProtocolSslDtls                types.Bool   `tfsdk:"protocol_ssl_dtls"`
	InterfaceSpecificCertificateId types.String `tfsdk:"interface_specific_certificate_id"`
}

type VPNRASecureClientImages struct {
	Id              types.String `tfsdk:"id"`
	OperatingSystem types.String `tfsdk:"operating_system"`
}

type VPNRAGroupPolicies struct {
	Id types.String `tfsdk:"id"`
}

type VPNRAIkev2Policies struct {
	Id types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data VPNRA) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/ravpns"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data VPNRA) toBody(ctx context.Context, state VPNRA) string {
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
	if !data.ProtocolSsl.IsNull() {
		body, _ = sjson.Set(body, "configureSSL", data.ProtocolSsl.ValueBool())
	}
	if !data.ProtocolIpsecIkev2.IsNull() {
		body, _ = sjson.Set(body, "configureIpsec", data.ProtocolIpsecIkev2.ValueBool())
	}
	if !data.LocalRealmId.IsNull() {
		body, _ = sjson.Set(body, "localRealmServer.id", data.LocalRealmId.ValueString())
	}
	if !data.DapPolicyId.IsNull() {
		body, _ = sjson.Set(body, "dapPolicy.id", data.DapPolicyId.ValueString())
	}
	if len(data.AccessInterfaces) > 0 {
		body, _ = sjson.Set(body, "accessInterfaceSettings.interfaceSettings", []interface{}{})
		for _, item := range data.AccessInterfaces {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "accessInterface.id", item.Id.ValueString())
			}
			if !item.ProtocolIpsecIkev2.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enableIPSecIkev2", item.ProtocolIpsecIkev2.ValueBool())
			}
			if !item.ProtocolSsl.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enableSSL", item.ProtocolSsl.ValueBool())
			}
			if !item.ProtocolSslDtls.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enableDTLS", item.ProtocolSslDtls.ValueBool())
			}
			if !item.InterfaceSpecificCertificateId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "idCertificate.id", item.InterfaceSpecificCertificateId.ValueString())
			}
			body, _ = sjson.SetRaw(body, "accessInterfaceSettings.interfaceSettings.-1", itemBody)
		}
	}
	if !data.AllowUsersToSelectConnectionProfile.IsNull() {
		body, _ = sjson.Set(body, "accessInterfaceSettings.allowConnectionProfileSelection", data.AllowUsersToSelectConnectionProfile.ValueBool())
	}
	if !data.WebAccessPortNumber.IsNull() {
		body, _ = sjson.Set(body, "accessInterfaceSettings.webPort", data.WebAccessPortNumber.ValueInt64())
	}
	if !data.DtlsPortNumber.IsNull() {
		body, _ = sjson.Set(body, "accessInterfaceSettings.dtlsPort", data.DtlsPortNumber.ValueInt64())
	}
	if !data.SslGlobalIdentityCertificateId.IsNull() {
		body, _ = sjson.Set(body, "accessInterfaceSettings.sslIdCertificate.id", data.SslGlobalIdentityCertificateId.ValueString())
	}
	if !data.IpsecIkev2IdentityCertificateId.IsNull() {
		body, _ = sjson.Set(body, "accessInterfaceSettings.ipsecIdCertificate.id", data.IpsecIkev2IdentityCertificateId.ValueString())
	}
	if !data.ServiceAccessObjectId.IsNull() {
		body, _ = sjson.Set(body, "accessInterfaceSettings.serviceAccessObject.id", data.ServiceAccessObjectId.ValueString())
	}
	if !data.BypassAccessControlPolicyForDecryptedTraffic.IsNull() {
		body, _ = sjson.Set(body, "accessInterfaceSettings.bypassACPolicyForDecryptTraffic", data.BypassAccessControlPolicyForDecryptedTraffic.ValueBool())
	}
	if len(data.SecureClientImages) > 0 {
		body, _ = sjson.Set(body, "anyConnectClientImages", []interface{}{})
		for _, item := range data.SecureClientImages {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "anyconnectImage.id", item.Id.ValueString())
			}
			if !item.OperatingSystem.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "operatingSystem", item.OperatingSystem.ValueString())
			}
			body, _ = sjson.SetRaw(body, "anyConnectClientImages.-1", itemBody)
		}
	}
	if !data.ExternalBrowserPackageId.IsNull() {
		body, _ = sjson.Set(body, "externalBrowserPackage.id", data.ExternalBrowserPackageId.ValueString())
	}
	if len(data.GroupPolicies) > 0 {
		body, _ = sjson.Set(body, "groupPolicies", []interface{}{})
		for _, item := range data.GroupPolicies {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "groupPolicies.-1", itemBody)
		}
	}
	if len(data.Ikev2Policies) > 0 {
		body, _ = sjson.Set(body, "ikev2Policies", []interface{}{})
		for _, item := range data.Ikev2Policies {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "ikev2Policies.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *VPNRA) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("configureSSL"); value.Exists() {
		data.ProtocolSsl = types.BoolValue(value.Bool())
	} else {
		data.ProtocolSsl = types.BoolValue(true)
	}
	if value := res.Get("configureIpsec"); value.Exists() {
		data.ProtocolIpsecIkev2 = types.BoolValue(value.Bool())
	} else {
		data.ProtocolIpsecIkev2 = types.BoolValue(true)
	}
	if value := res.Get("localRealmServer.id"); value.Exists() {
		data.LocalRealmId = types.StringValue(value.String())
	} else {
		data.LocalRealmId = types.StringNull()
	}
	if value := res.Get("dapPolicy.id"); value.Exists() {
		data.DapPolicyId = types.StringValue(value.String())
	} else {
		data.DapPolicyId = types.StringNull()
	}
	if value := res.Get("accessInterfaceSettings.interfaceSettings"); value.Exists() {
		data.AccessInterfaces = make([]VPNRAAccessInterfaces, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := VPNRAAccessInterfaces{}
			if value := res.Get("accessInterface.id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("enableIPSecIkev2"); value.Exists() {
				data.ProtocolIpsecIkev2 = types.BoolValue(value.Bool())
			} else {
				data.ProtocolIpsecIkev2 = types.BoolValue(true)
			}
			if value := res.Get("enableSSL"); value.Exists() {
				data.ProtocolSsl = types.BoolValue(value.Bool())
			} else {
				data.ProtocolSsl = types.BoolValue(true)
			}
			if value := res.Get("enableDTLS"); value.Exists() {
				data.ProtocolSslDtls = types.BoolValue(value.Bool())
			} else {
				data.ProtocolSslDtls = types.BoolValue(true)
			}
			if value := res.Get("idCertificate.id"); value.Exists() {
				data.InterfaceSpecificCertificateId = types.StringValue(value.String())
			} else {
				data.InterfaceSpecificCertificateId = types.StringNull()
			}
			(*parent).AccessInterfaces = append((*parent).AccessInterfaces, data)
			return true
		})
	}
	if value := res.Get("accessInterfaceSettings.allowConnectionProfileSelection"); value.Exists() {
		data.AllowUsersToSelectConnectionProfile = types.BoolValue(value.Bool())
	} else {
		data.AllowUsersToSelectConnectionProfile = types.BoolNull()
	}
	if value := res.Get("accessInterfaceSettings.webPort"); value.Exists() {
		data.WebAccessPortNumber = types.Int64Value(value.Int())
	} else {
		data.WebAccessPortNumber = types.Int64Value(443)
	}
	if value := res.Get("accessInterfaceSettings.dtlsPort"); value.Exists() {
		data.DtlsPortNumber = types.Int64Value(value.Int())
	} else {
		data.DtlsPortNumber = types.Int64Value(443)
	}
	if value := res.Get("accessInterfaceSettings.sslIdCertificate.id"); value.Exists() {
		data.SslGlobalIdentityCertificateId = types.StringValue(value.String())
	} else {
		data.SslGlobalIdentityCertificateId = types.StringNull()
	}
	if value := res.Get("accessInterfaceSettings.ipsecIdCertificate.id"); value.Exists() {
		data.IpsecIkev2IdentityCertificateId = types.StringValue(value.String())
	} else {
		data.IpsecIkev2IdentityCertificateId = types.StringNull()
	}
	if value := res.Get("accessInterfaceSettings.serviceAccessObject.id"); value.Exists() {
		data.ServiceAccessObjectId = types.StringValue(value.String())
	} else {
		data.ServiceAccessObjectId = types.StringNull()
	}
	if value := res.Get("accessInterfaceSettings.bypassACPolicyForDecryptTraffic"); value.Exists() {
		data.BypassAccessControlPolicyForDecryptedTraffic = types.BoolValue(value.Bool())
	} else {
		data.BypassAccessControlPolicyForDecryptedTraffic = types.BoolValue(false)
	}
	if value := res.Get("anyConnectClientImages"); value.Exists() {
		data.SecureClientImages = make([]VPNRASecureClientImages, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := VPNRASecureClientImages{}
			if value := res.Get("anyconnectImage.id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("operatingSystem"); value.Exists() {
				data.OperatingSystem = types.StringValue(value.String())
			} else {
				data.OperatingSystem = types.StringNull()
			}
			(*parent).SecureClientImages = append((*parent).SecureClientImages, data)
			return true
		})
	}
	if value := res.Get("externalBrowserPackage.id"); value.Exists() {
		data.ExternalBrowserPackageId = types.StringValue(value.String())
	} else {
		data.ExternalBrowserPackageId = types.StringNull()
	}
	if value := res.Get("secureClientCustomizationSettings.id"); value.Exists() {
		data.SecureClientCustomizationId = types.StringValue(value.String())
	} else {
		data.SecureClientCustomizationId = types.StringNull()
	}
	if value := res.Get("addressAssignmentSettings.id"); value.Exists() {
		data.AddressAssignmentPolicyId = types.StringValue(value.String())
	} else {
		data.AddressAssignmentPolicyId = types.StringNull()
	}
	if value := res.Get("certificateMapSettings.id"); value.Exists() {
		data.CertificateMapId = types.StringValue(value.String())
	} else {
		data.CertificateMapId = types.StringNull()
	}
	if value := res.Get("groupPolicies"); value.Exists() {
		data.GroupPolicies = make([]VPNRAGroupPolicies, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := VPNRAGroupPolicies{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).GroupPolicies = append((*parent).GroupPolicies, data)
			return true
		})
	}
	if value := res.Get("ldapAttributeMaps.id"); value.Exists() {
		data.LdapAttributeMapId = types.StringValue(value.String())
	} else {
		data.LdapAttributeMapId = types.StringNull()
	}
	if value := res.Get("loadBalanceSettings.id"); value.Exists() {
		data.LoadBalancingId = types.StringValue(value.String())
	} else {
		data.LoadBalancingId = types.StringNull()
	}
	if value := res.Get("ikev2Policies"); value.Exists() {
		data.Ikev2Policies = make([]VPNRAIkev2Policies, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := VPNRAIkev2Policies{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).Ikev2Policies = append((*parent).Ikev2Policies, data)
			return true
		})
	}
	if value := res.Get("ipsecAdvancedSettings.id"); value.Exists() {
		data.IpsecIkeParametersId = types.StringValue(value.String())
	} else {
		data.IpsecIkeParametersId = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *VPNRA) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("configureSSL"); value.Exists() && !data.ProtocolSsl.IsNull() {
		data.ProtocolSsl = types.BoolValue(value.Bool())
	} else if data.ProtocolSsl.ValueBool() != true {
		data.ProtocolSsl = types.BoolNull()
	}
	if value := res.Get("configureIpsec"); value.Exists() && !data.ProtocolIpsecIkev2.IsNull() {
		data.ProtocolIpsecIkev2 = types.BoolValue(value.Bool())
	} else if data.ProtocolIpsecIkev2.ValueBool() != true {
		data.ProtocolIpsecIkev2 = types.BoolNull()
	}
	if value := res.Get("localRealmServer.id"); value.Exists() && !data.LocalRealmId.IsNull() {
		data.LocalRealmId = types.StringValue(value.String())
	} else {
		data.LocalRealmId = types.StringNull()
	}
	if value := res.Get("dapPolicy.id"); value.Exists() && !data.DapPolicyId.IsNull() {
		data.DapPolicyId = types.StringValue(value.String())
	} else {
		data.DapPolicyId = types.StringNull()
	}
	for i := 0; i < len(data.AccessInterfaces); i++ {
		keys := [...]string{"accessInterface.id"}
		keyValues := [...]string{data.AccessInterfaces[i].Id.ValueString()}

		parent := &data
		data := (*parent).AccessInterfaces[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("accessInterfaceSettings.interfaceSettings").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing AccessInterfaces[%d] = %+v",
				i,
				(*parent).AccessInterfaces[i],
			))
			(*parent).AccessInterfaces = slices.Delete((*parent).AccessInterfaces, i, i+1)
			i--

			continue
		}
		if value := res.Get("accessInterface.id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("enableIPSecIkev2"); value.Exists() && !data.ProtocolIpsecIkev2.IsNull() {
			data.ProtocolIpsecIkev2 = types.BoolValue(value.Bool())
		} else if data.ProtocolIpsecIkev2.ValueBool() != true {
			data.ProtocolIpsecIkev2 = types.BoolNull()
		}
		if value := res.Get("enableSSL"); value.Exists() && !data.ProtocolSsl.IsNull() {
			data.ProtocolSsl = types.BoolValue(value.Bool())
		} else if data.ProtocolSsl.ValueBool() != true {
			data.ProtocolSsl = types.BoolNull()
		}
		if value := res.Get("enableDTLS"); value.Exists() && !data.ProtocolSslDtls.IsNull() {
			data.ProtocolSslDtls = types.BoolValue(value.Bool())
		} else if data.ProtocolSslDtls.ValueBool() != true {
			data.ProtocolSslDtls = types.BoolNull()
		}
		if value := res.Get("idCertificate.id"); value.Exists() && !data.InterfaceSpecificCertificateId.IsNull() {
			data.InterfaceSpecificCertificateId = types.StringValue(value.String())
		} else {
			data.InterfaceSpecificCertificateId = types.StringNull()
		}
		(*parent).AccessInterfaces[i] = data
	}
	if value := res.Get("accessInterfaceSettings.allowConnectionProfileSelection"); value.Exists() && !data.AllowUsersToSelectConnectionProfile.IsNull() {
		data.AllowUsersToSelectConnectionProfile = types.BoolValue(value.Bool())
	} else {
		data.AllowUsersToSelectConnectionProfile = types.BoolNull()
	}
	if value := res.Get("accessInterfaceSettings.webPort"); value.Exists() && !data.WebAccessPortNumber.IsNull() {
		data.WebAccessPortNumber = types.Int64Value(value.Int())
	} else if data.WebAccessPortNumber.ValueInt64() != 443 {
		data.WebAccessPortNumber = types.Int64Null()
	}
	if value := res.Get("accessInterfaceSettings.dtlsPort"); value.Exists() && !data.DtlsPortNumber.IsNull() {
		data.DtlsPortNumber = types.Int64Value(value.Int())
	} else if data.DtlsPortNumber.ValueInt64() != 443 {
		data.DtlsPortNumber = types.Int64Null()
	}
	if value := res.Get("accessInterfaceSettings.sslIdCertificate.id"); value.Exists() && !data.SslGlobalIdentityCertificateId.IsNull() {
		data.SslGlobalIdentityCertificateId = types.StringValue(value.String())
	} else {
		data.SslGlobalIdentityCertificateId = types.StringNull()
	}
	if value := res.Get("accessInterfaceSettings.ipsecIdCertificate.id"); value.Exists() && !data.IpsecIkev2IdentityCertificateId.IsNull() {
		data.IpsecIkev2IdentityCertificateId = types.StringValue(value.String())
	} else {
		data.IpsecIkev2IdentityCertificateId = types.StringNull()
	}
	if value := res.Get("accessInterfaceSettings.serviceAccessObject.id"); value.Exists() && !data.ServiceAccessObjectId.IsNull() {
		data.ServiceAccessObjectId = types.StringValue(value.String())
	} else {
		data.ServiceAccessObjectId = types.StringNull()
	}
	if value := res.Get("accessInterfaceSettings.bypassACPolicyForDecryptTraffic"); value.Exists() && !data.BypassAccessControlPolicyForDecryptedTraffic.IsNull() {
		data.BypassAccessControlPolicyForDecryptedTraffic = types.BoolValue(value.Bool())
	} else if data.BypassAccessControlPolicyForDecryptedTraffic.ValueBool() != false {
		data.BypassAccessControlPolicyForDecryptedTraffic = types.BoolNull()
	}
	{
		l := len(res.Get("anyConnectClientImages").Array())
		tflog.Debug(ctx, fmt.Sprintf("anyConnectClientImages array resizing from %d to %d", len(data.SecureClientImages), l))
		for i := len(data.SecureClientImages); i < l; i++ {
			data.SecureClientImages = append(data.SecureClientImages, VPNRASecureClientImages{})
		}
		if len(data.SecureClientImages) > l {
			data.SecureClientImages = data.SecureClientImages[:l]
		}
	}
	for i := range data.SecureClientImages {
		parent := &data
		data := (*parent).SecureClientImages[i]
		parentRes := &res
		res := parentRes.Get(fmt.Sprintf("anyConnectClientImages.%d", i))
		if value := res.Get("anyconnectImage.id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("operatingSystem"); value.Exists() && !data.OperatingSystem.IsNull() {
			data.OperatingSystem = types.StringValue(value.String())
		} else {
			data.OperatingSystem = types.StringNull()
		}
		(*parent).SecureClientImages[i] = data
	}
	if value := res.Get("externalBrowserPackage.id"); value.Exists() && !data.ExternalBrowserPackageId.IsNull() {
		data.ExternalBrowserPackageId = types.StringValue(value.String())
	} else {
		data.ExternalBrowserPackageId = types.StringNull()
	}
	if value := res.Get("secureClientCustomizationSettings.id"); value.Exists() && !data.SecureClientCustomizationId.IsNull() {
		data.SecureClientCustomizationId = types.StringValue(value.String())
	} else {
		data.SecureClientCustomizationId = types.StringNull()
	}
	if value := res.Get("addressAssignmentSettings.id"); value.Exists() && !data.AddressAssignmentPolicyId.IsNull() {
		data.AddressAssignmentPolicyId = types.StringValue(value.String())
	} else {
		data.AddressAssignmentPolicyId = types.StringNull()
	}
	if value := res.Get("certificateMapSettings.id"); value.Exists() && !data.CertificateMapId.IsNull() {
		data.CertificateMapId = types.StringValue(value.String())
	} else {
		data.CertificateMapId = types.StringNull()
	}
	for i := 0; i < len(data.GroupPolicies); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.GroupPolicies[i].Id.ValueString()}

		parent := &data
		data := (*parent).GroupPolicies[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("groupPolicies").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing GroupPolicies[%d] = %+v",
				i,
				(*parent).GroupPolicies[i],
			))
			(*parent).GroupPolicies = slices.Delete((*parent).GroupPolicies, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).GroupPolicies[i] = data
	}
	if value := res.Get("ldapAttributeMaps.id"); value.Exists() && !data.LdapAttributeMapId.IsNull() {
		data.LdapAttributeMapId = types.StringValue(value.String())
	} else {
		data.LdapAttributeMapId = types.StringNull()
	}
	if value := res.Get("loadBalanceSettings.id"); value.Exists() && !data.LoadBalancingId.IsNull() {
		data.LoadBalancingId = types.StringValue(value.String())
	} else {
		data.LoadBalancingId = types.StringNull()
	}
	for i := 0; i < len(data.Ikev2Policies); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.Ikev2Policies[i].Id.ValueString()}

		parent := &data
		data := (*parent).Ikev2Policies[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("ikev2Policies").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Ikev2Policies[%d] = %+v",
				i,
				(*parent).Ikev2Policies[i],
			))
			(*parent).Ikev2Policies = slices.Delete((*parent).Ikev2Policies, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).Ikev2Policies[i] = data
	}
	if value := res.Get("ipsecAdvancedSettings.id"); value.Exists() && !data.IpsecIkeParametersId.IsNull() {
		data.IpsecIkeParametersId = types.StringValue(value.String())
	} else {
		data.IpsecIkeParametersId = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *VPNRA) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
	if data.SecureClientCustomizationId.IsUnknown() {
		if value := res.Get("secureClientCustomizationSettings.id"); value.Exists() {
			data.SecureClientCustomizationId = types.StringValue(value.String())
		} else {
			data.SecureClientCustomizationId = types.StringNull()
		}
	}
	if data.AddressAssignmentPolicyId.IsUnknown() {
		if value := res.Get("addressAssignmentSettings.id"); value.Exists() {
			data.AddressAssignmentPolicyId = types.StringValue(value.String())
		} else {
			data.AddressAssignmentPolicyId = types.StringNull()
		}
	}
	if data.CertificateMapId.IsUnknown() {
		if value := res.Get("certificateMapSettings.id"); value.Exists() {
			data.CertificateMapId = types.StringValue(value.String())
		} else {
			data.CertificateMapId = types.StringNull()
		}
	}
	if data.LdapAttributeMapId.IsUnknown() {
		if value := res.Get("ldapAttributeMaps.id"); value.Exists() {
			data.LdapAttributeMapId = types.StringValue(value.String())
		} else {
			data.LdapAttributeMapId = types.StringNull()
		}
	}
	if data.LoadBalancingId.IsUnknown() {
		if value := res.Get("loadBalanceSettings.id"); value.Exists() {
			data.LoadBalancingId = types.StringValue(value.String())
		} else {
			data.LoadBalancingId = types.StringNull()
		}
	}
	if data.IpsecIkeParametersId.IsUnknown() {
		if value := res.Get("ipsecAdvancedSettings.id"); value.Exists() {
			data.IpsecIkeParametersId = types.StringValue(value.String())
		} else {
			data.IpsecIkeParametersId = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns
