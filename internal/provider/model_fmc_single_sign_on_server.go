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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type SingleSignOnServer struct {
	Id                                                 types.String `tfsdk:"id"`
	Domain                                             types.String `tfsdk:"domain"`
	Name                                               types.String `tfsdk:"name"`
	Type                                               types.String `tfsdk:"type"`
	IdentityProviderEntityIdUrl                        types.String `tfsdk:"identity_provider_entity_id_url"`
	SsoUrl                                             types.String `tfsdk:"sso_url"`
	LogoutUrl                                          types.String `tfsdk:"logout_url"`
	BaseUrl                                            types.String `tfsdk:"base_url"`
	IdentityProviderCertificateId                      types.String `tfsdk:"identity_provider_certificate_id"`
	IdentityProviderCertificateName                    types.String `tfsdk:"identity_provider_certificate_name"`
	ServiceProviderCertificateId                       types.String `tfsdk:"service_provider_certificate_id"`
	RequestSignatureType                               types.String `tfsdk:"request_signature_type"`
	RequestTimeout                                     types.Int64  `tfsdk:"request_timeout"`
	IdentityProviderAccessibleOnlyOnInternalNetwork    types.Bool   `tfsdk:"identity_provider_accessible_only_on_internal_network"`
	RequestIdentityProviderReauthenticationAtEachLogin types.Bool   `tfsdk:"request_identity_provider_reauthentication_at_each_login"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SingleSignOnServer) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/ssoservers"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data SingleSignOnServer) toBody(ctx context.Context, state SingleSignOnServer) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	body, _ = sjson.Set(body, "type", "SSOServer")
	if !data.IdentityProviderEntityIdUrl.IsNull() {
		body, _ = sjson.Set(body, "idpEntityId", data.IdentityProviderEntityIdUrl.ValueString())
	}
	if !data.SsoUrl.IsNull() {
		body, _ = sjson.Set(body, "ssoURL", data.SsoUrl.ValueString())
	}
	if !data.LogoutUrl.IsNull() {
		body, _ = sjson.Set(body, "logoutURL", data.LogoutUrl.ValueString())
	}
	if !data.BaseUrl.IsNull() {
		body, _ = sjson.Set(body, "baseUrl", data.BaseUrl.ValueString())
	}
	if !data.IdentityProviderCertificateId.IsNull() {
		body, _ = sjson.Set(body, "identityProviderCert.id", data.IdentityProviderCertificateId.ValueString())
	}
	if !data.IdentityProviderCertificateName.IsNull() {
		body, _ = sjson.Set(body, "identityProviderCert.name", data.IdentityProviderCertificateName.ValueString())
	}
	if !data.ServiceProviderCertificateId.IsNull() {
		body, _ = sjson.Set(body, "serviceProviderCert.id", data.ServiceProviderCertificateId.ValueString())
	}
	if !data.RequestSignatureType.IsNull() {
		body, _ = sjson.Set(body, "requestSignatureType", data.RequestSignatureType.ValueString())
	}
	if !data.RequestTimeout.IsNull() {
		body, _ = sjson.Set(body, "requestTimeout", data.RequestTimeout.ValueInt64())
	}
	if !data.IdentityProviderAccessibleOnlyOnInternalNetwork.IsNull() {
		body, _ = sjson.Set(body, "serverOnInternalNetwork", data.IdentityProviderAccessibleOnlyOnInternalNetwork.ValueBool())
	}
	if !data.RequestIdentityProviderReauthenticationAtEachLogin.IsNull() {
		body, _ = sjson.Set(body, "reAuthAtLogin", data.RequestIdentityProviderReauthenticationAtEachLogin.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SingleSignOnServer) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("idpEntityId"); value.Exists() {
		data.IdentityProviderEntityIdUrl = types.StringValue(value.String())
	} else {
		data.IdentityProviderEntityIdUrl = types.StringNull()
	}
	if value := res.Get("ssoURL"); value.Exists() {
		data.SsoUrl = types.StringValue(value.String())
	} else {
		data.SsoUrl = types.StringNull()
	}
	if value := res.Get("logoutURL"); value.Exists() {
		data.LogoutUrl = types.StringValue(value.String())
	} else {
		data.LogoutUrl = types.StringNull()
	}
	if value := res.Get("baseUrl"); value.Exists() {
		data.BaseUrl = types.StringValue(value.String())
	} else {
		data.BaseUrl = types.StringNull()
	}
	if value := res.Get("identityProviderCert.id"); value.Exists() {
		data.IdentityProviderCertificateId = types.StringValue(value.String())
	} else {
		data.IdentityProviderCertificateId = types.StringNull()
	}
	if value := res.Get("identityProviderCert.name"); value.Exists() {
		data.IdentityProviderCertificateName = types.StringValue(value.String())
	} else {
		data.IdentityProviderCertificateName = types.StringNull()
	}
	if value := res.Get("serviceProviderCert.id"); value.Exists() {
		data.ServiceProviderCertificateId = types.StringValue(value.String())
	} else {
		data.ServiceProviderCertificateId = types.StringNull()
	}
	if value := res.Get("requestSignatureType"); value.Exists() {
		data.RequestSignatureType = types.StringValue(value.String())
	} else {
		data.RequestSignatureType = types.StringNull()
	}
	if value := res.Get("requestTimeout"); value.Exists() {
		data.RequestTimeout = types.Int64Value(value.Int())
	} else {
		data.RequestTimeout = types.Int64Value(300)
	}
	if value := res.Get("serverOnInternalNetwork"); value.Exists() {
		data.IdentityProviderAccessibleOnlyOnInternalNetwork = types.BoolValue(value.Bool())
	} else {
		data.IdentityProviderAccessibleOnlyOnInternalNetwork = types.BoolNull()
	}
	if value := res.Get("reAuthAtLogin"); value.Exists() {
		data.RequestIdentityProviderReauthenticationAtEachLogin = types.BoolValue(value.Bool())
	} else {
		data.RequestIdentityProviderReauthenticationAtEachLogin = types.BoolNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *SingleSignOnServer) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("idpEntityId"); value.Exists() && !data.IdentityProviderEntityIdUrl.IsNull() {
		data.IdentityProviderEntityIdUrl = types.StringValue(value.String())
	} else {
		data.IdentityProviderEntityIdUrl = types.StringNull()
	}
	if value := res.Get("ssoURL"); value.Exists() && !data.SsoUrl.IsNull() {
		data.SsoUrl = types.StringValue(value.String())
	} else {
		data.SsoUrl = types.StringNull()
	}
	if value := res.Get("logoutURL"); value.Exists() && !data.LogoutUrl.IsNull() {
		data.LogoutUrl = types.StringValue(value.String())
	} else {
		data.LogoutUrl = types.StringNull()
	}
	if value := res.Get("baseUrl"); value.Exists() && !data.BaseUrl.IsNull() {
		data.BaseUrl = types.StringValue(value.String())
	} else {
		data.BaseUrl = types.StringNull()
	}
	if value := res.Get("identityProviderCert.id"); value.Exists() && !data.IdentityProviderCertificateId.IsNull() {
		data.IdentityProviderCertificateId = types.StringValue(value.String())
	} else {
		data.IdentityProviderCertificateId = types.StringNull()
	}
	if value := res.Get("identityProviderCert.name"); value.Exists() && !data.IdentityProviderCertificateName.IsNull() {
		data.IdentityProviderCertificateName = types.StringValue(value.String())
	} else {
		data.IdentityProviderCertificateName = types.StringNull()
	}
	if value := res.Get("serviceProviderCert.id"); value.Exists() && !data.ServiceProviderCertificateId.IsNull() {
		data.ServiceProviderCertificateId = types.StringValue(value.String())
	} else {
		data.ServiceProviderCertificateId = types.StringNull()
	}
	if value := res.Get("requestSignatureType"); value.Exists() && !data.RequestSignatureType.IsNull() {
		data.RequestSignatureType = types.StringValue(value.String())
	} else {
		data.RequestSignatureType = types.StringNull()
	}
	if value := res.Get("requestTimeout"); value.Exists() && !data.RequestTimeout.IsNull() {
		data.RequestTimeout = types.Int64Value(value.Int())
	} else if data.RequestTimeout.ValueInt64() != 300 {
		data.RequestTimeout = types.Int64Null()
	}
	if value := res.Get("serverOnInternalNetwork"); value.Exists() && !data.IdentityProviderAccessibleOnlyOnInternalNetwork.IsNull() {
		data.IdentityProviderAccessibleOnlyOnInternalNetwork = types.BoolValue(value.Bool())
	} else {
		data.IdentityProviderAccessibleOnlyOnInternalNetwork = types.BoolNull()
	}
	if value := res.Get("reAuthAtLogin"); value.Exists() && !data.RequestIdentityProviderReauthenticationAtEachLogin.IsNull() {
		data.RequestIdentityProviderReauthenticationAtEachLogin = types.BoolValue(value.Bool())
	} else {
		data.RequestIdentityProviderReauthenticationAtEachLogin = types.BoolNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *SingleSignOnServer) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
