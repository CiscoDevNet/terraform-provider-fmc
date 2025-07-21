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

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type CertificateEnrollment struct {
	Id                                                          types.String `tfsdk:"id"`
	Domain                                                      types.String `tfsdk:"domain"`
	Name                                                        types.String `tfsdk:"name"`
	Type                                                        types.String `tfsdk:"type"`
	Description                                                 types.String `tfsdk:"description"`
	EnrollmentType                                              types.String `tfsdk:"enrollment_type"`
	ValidationUsageIpsecClient                                  types.Bool   `tfsdk:"validation_usage_ipsec_client"`
	ValidationUsageSslServer                                    types.Bool   `tfsdk:"validation_usage_ssl_server"`
	ValidationUsageSslClient                                    types.Bool   `tfsdk:"validation_usage_ssl_client"`
	SkipCaFlagCheck                                             types.Bool   `tfsdk:"skip_ca_flag_check"`
	EstEnrollmentUrl                                            types.String `tfsdk:"est_enrollment_url"`
	EstUsername                                                 types.String `tfsdk:"est_username"`
	EstPassword                                                 types.String `tfsdk:"est_password"`
	EstFingerprint                                              types.String `tfsdk:"est_fingerprint"`
	EstSourceInterfaceId                                        types.String `tfsdk:"est_source_interface_id"`
	EstSourceInterfaceName                                      types.String `tfsdk:"est_source_interface_name"`
	EstIgnoreServerCertificateValidation                        types.Bool   `tfsdk:"est_ignore_server_certificate_validation"`
	ScepEnrollmentUrl                                           types.String `tfsdk:"scep_enrollment_url"`
	ScepChallengePassword                                       types.String `tfsdk:"scep_challenge_password"`
	ScepRetryPeriod                                             types.Int64  `tfsdk:"scep_retry_period"`
	ScepRetryCount                                              types.Int64  `tfsdk:"scep_retry_count"`
	ScepFingerprint                                             types.String `tfsdk:"scep_fingerprint"`
	ManualCaOnly                                                types.Bool   `tfsdk:"manual_ca_only"`
	ManualCaCertificate                                         types.String `tfsdk:"manual_ca_certificate"`
	Pkcs12Certificate                                           types.String `tfsdk:"pkcs12_certificate"`
	Pkcs12CertificatePassphrase                                 types.String `tfsdk:"pkcs12_certificate_passphrase"`
	IncludeFqdn                                                 types.String `tfsdk:"include_fqdn"`
	CustomFqdn                                                  types.String `tfsdk:"custom_fqdn"`
	IncludeDeviceIp                                             types.String `tfsdk:"include_device_ip"`
	CommonName                                                  types.String `tfsdk:"common_name"`
	OrganizationalUnit                                          types.String `tfsdk:"organizational_unit"`
	Organization                                                types.String `tfsdk:"organization"`
	Locality                                                    types.String `tfsdk:"locality"`
	State                                                       types.String `tfsdk:"state"`
	CountryCode                                                 types.String `tfsdk:"country_code"`
	Email                                                       types.String `tfsdk:"email"`
	IncludeDeviceSerialNumber                                   types.Bool   `tfsdk:"include_device_serial_number"`
	KeyType                                                     types.String `tfsdk:"key_type"`
	KeyName                                                     types.String `tfsdk:"key_name"`
	KeySize                                                     types.String `tfsdk:"key_size"`
	IgnoreIpsecKeyUsage                                         types.Bool   `tfsdk:"ignore_ipsec_key_usage"`
	CrlUseDistributionPointFromTheCertificate                   types.Bool   `tfsdk:"crl_use_distribution_point_from_the_certificate"`
	CrlStaticUrlsList                                           types.List   `tfsdk:"crl_static_urls_list"`
	OcspUrl                                                     types.String `tfsdk:"ocsp_url"`
	EvaluationPriority                                          types.String `tfsdk:"evaluation_priority"`
	ConsiderCertificateValidIfRevocationInformationNotReachable types.Bool   `tfsdk:"consider_certificate_valid_if_revocation_information_not_reachable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data CertificateEnrollment) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/certenrollments"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data CertificateEnrollment) toBody(ctx context.Context, state CertificateEnrollment) string {
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
	if !data.EnrollmentType.IsNull() {
		body, _ = sjson.Set(body, "enrollmentType", data.EnrollmentType.ValueString())
	}
	if !data.ValidationUsageIpsecClient.IsNull() {
		body, _ = sjson.Set(body, "validationUsage.ipsecClient", data.ValidationUsageIpsecClient.ValueBool())
	}
	if !data.ValidationUsageSslServer.IsNull() {
		body, _ = sjson.Set(body, "validationUsage.sslServer", data.ValidationUsageSslServer.ValueBool())
	}
	if !data.ValidationUsageSslClient.IsNull() {
		body, _ = sjson.Set(body, "validationUsage.sslClient", data.ValidationUsageSslClient.ValueBool())
	}
	if !data.SkipCaFlagCheck.IsNull() {
		body, _ = sjson.Set(body, "skipCaFlagCheck", data.SkipCaFlagCheck.ValueBool())
	}
	if !data.EstEnrollmentUrl.IsNull() {
		body, _ = sjson.Set(body, "estContent.enrollmentUrl", data.EstEnrollmentUrl.ValueString())
	}
	if !data.EstUsername.IsNull() {
		body, _ = sjson.Set(body, "estContent.username", data.EstUsername.ValueString())
	}
	if !data.EstPassword.IsNull() {
		body, _ = sjson.Set(body, "estContent.password", data.EstPassword.ValueString())
	}
	if !data.EstFingerprint.IsNull() {
		body, _ = sjson.Set(body, "estContent.fingerprint", data.EstFingerprint.ValueString())
	}
	if !data.EstSourceInterfaceId.IsNull() {
		body, _ = sjson.Set(body, "estContent.sourceInterface.id", data.EstSourceInterfaceId.ValueString())
	}
	if !data.EstSourceInterfaceName.IsNull() {
		body, _ = sjson.Set(body, "estContent.sourceInterface.name", data.EstSourceInterfaceName.ValueString())
	}
	if !data.EstIgnoreServerCertificateValidation.IsNull() {
		body, _ = sjson.Set(body, "estContent.ignoreServerCertificateValidations", data.EstIgnoreServerCertificateValidation.ValueBool())
	}
	if !data.ScepEnrollmentUrl.IsNull() {
		body, _ = sjson.Set(body, "scepContent.enrollmentUrl", data.ScepEnrollmentUrl.ValueString())
	}
	if !data.ScepChallengePassword.IsNull() {
		body, _ = sjson.Set(body, "scepContent.challengePassword", data.ScepChallengePassword.ValueString())
	}
	if !data.ScepRetryPeriod.IsNull() {
		body, _ = sjson.Set(body, "scepContent.retryPeriodInMins", data.ScepRetryPeriod.ValueInt64())
	}
	if !data.ScepRetryCount.IsNull() {
		body, _ = sjson.Set(body, "scepContent.retryCount", data.ScepRetryCount.ValueInt64())
	}
	if !data.ScepFingerprint.IsNull() {
		body, _ = sjson.Set(body, "scepContent.fingerprint", data.ScepFingerprint.ValueString())
	}
	if !data.ManualCaOnly.IsNull() {
		body, _ = sjson.Set(body, "caCertificate.caOnly", data.ManualCaOnly.ValueBool())
	}
	if !data.ManualCaCertificate.IsNull() {
		body, _ = sjson.Set(body, "caCertificate.base64Content", data.ManualCaCertificate.ValueString())
	}
	if !data.Pkcs12Certificate.IsNull() {
		body, _ = sjson.Set(body, "pkcs12Content.base64Certificate", data.Pkcs12Certificate.ValueString())
	}
	if !data.Pkcs12CertificatePassphrase.IsNull() {
		body, _ = sjson.Set(body, "pkcs12Content.passPhrase", data.Pkcs12CertificatePassphrase.ValueString())
	}
	if !data.IncludeFqdn.IsNull() {
		body, _ = sjson.Set(body, "certificateParameters.includeFQDN", data.IncludeFqdn.ValueString())
	}
	if !data.CustomFqdn.IsNull() {
		body, _ = sjson.Set(body, "certificateParameters.customFqdn", data.CustomFqdn.ValueString())
	}
	if !data.IncludeDeviceIp.IsNull() {
		body, _ = sjson.Set(body, "certificateParameters.includeDeviceIp", data.IncludeDeviceIp.ValueString())
	}
	if !data.CommonName.IsNull() {
		body, _ = sjson.Set(body, "certificateParameters.commonName", data.CommonName.ValueString())
	}
	if !data.OrganizationalUnit.IsNull() {
		body, _ = sjson.Set(body, "certificateParameters.organizationalUnit", data.OrganizationalUnit.ValueString())
	}
	if !data.Organization.IsNull() {
		body, _ = sjson.Set(body, "certificateParameters.organization", data.Organization.ValueString())
	}
	if !data.Locality.IsNull() {
		body, _ = sjson.Set(body, "certificateParameters.locality", data.Locality.ValueString())
	}
	if !data.State.IsNull() {
		body, _ = sjson.Set(body, "certificateParameters.state", data.State.ValueString())
	}
	if !data.CountryCode.IsNull() {
		body, _ = sjson.Set(body, "certificateParameters.countryCode", data.CountryCode.ValueString())
	}
	if !data.Email.IsNull() {
		body, _ = sjson.Set(body, "certificateParameters.email", data.Email.ValueString())
	}
	if !data.IncludeDeviceSerialNumber.IsNull() {
		body, _ = sjson.Set(body, "certificateParameters.includeDeviceSerial", data.IncludeDeviceSerialNumber.ValueBool())
	}
	if !data.KeyType.IsNull() {
		body, _ = sjson.Set(body, "key.keyType", data.KeyType.ValueString())
	}
	if !data.KeyName.IsNull() {
		body, _ = sjson.Set(body, "key.keyName", data.KeyName.ValueString())
	}
	if !data.KeySize.IsNull() {
		body, _ = sjson.Set(body, "key.keySize", data.KeySize.ValueString())
	}
	if !data.IgnoreIpsecKeyUsage.IsNull() {
		body, _ = sjson.Set(body, "key.ignoreIpsecKeyusage", data.IgnoreIpsecKeyUsage.ValueBool())
	}
	if !data.CrlUseDistributionPointFromTheCertificate.IsNull() {
		body, _ = sjson.Set(body, "revocation.certRevocationListContent.enableCertRevocationListDistributionPoint", data.CrlUseDistributionPointFromTheCertificate.ValueBool())
	}
	if !data.CrlStaticUrlsList.IsNull() {
		var values []string
		data.CrlStaticUrlsList.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "revocation.certRevocationListContent.certRevocationStaticUrlList", values)
	}
	if !data.OcspUrl.IsNull() {
		body, _ = sjson.Set(body, "revocation.onlineCertificateStatusProtocolContent.onlineCertificateStatusProtocolUrl", data.OcspUrl.ValueString())
	}
	if !data.EvaluationPriority.IsNull() {
		body, _ = sjson.Set(body, "revocation.evaluationPriority", data.EvaluationPriority.ValueString())
	}
	if !data.ConsiderCertificateValidIfRevocationInformationNotReachable.IsNull() {
		body, _ = sjson.Set(body, "revocation.ignoreRevocation", data.ConsiderCertificateValidIfRevocationInformationNotReachable.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *CertificateEnrollment) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("enrollmentType"); value.Exists() {
		data.EnrollmentType = types.StringValue(value.String())
	} else {
		data.EnrollmentType = types.StringNull()
	}
	if value := res.Get("validationUsage.ipsecClient"); value.Exists() {
		data.ValidationUsageIpsecClient = types.BoolValue(value.Bool())
	} else {
		data.ValidationUsageIpsecClient = types.BoolNull()
	}
	if value := res.Get("validationUsage.sslServer"); value.Exists() {
		data.ValidationUsageSslServer = types.BoolValue(value.Bool())
	} else {
		data.ValidationUsageSslServer = types.BoolNull()
	}
	if value := res.Get("validationUsage.sslClient"); value.Exists() {
		data.ValidationUsageSslClient = types.BoolValue(value.Bool())
	} else {
		data.ValidationUsageSslClient = types.BoolNull()
	}
	if value := res.Get("estContent.enrollmentUrl"); value.Exists() {
		data.EstEnrollmentUrl = types.StringValue(value.String())
	} else {
		data.EstEnrollmentUrl = types.StringNull()
	}
	if value := res.Get("estContent.username"); value.Exists() {
		data.EstUsername = types.StringValue(value.String())
	} else {
		data.EstUsername = types.StringNull()
	}
	if value := res.Get("estContent.fingerprint"); value.Exists() {
		data.EstFingerprint = types.StringValue(value.String())
	} else {
		data.EstFingerprint = types.StringNull()
	}
	if value := res.Get("estContent.sourceInterface.id"); value.Exists() {
		data.EstSourceInterfaceId = types.StringValue(value.String())
	} else {
		data.EstSourceInterfaceId = types.StringNull()
	}
	if value := res.Get("estContent.sourceInterface.name"); value.Exists() {
		data.EstSourceInterfaceName = types.StringValue(value.String())
	} else {
		data.EstSourceInterfaceName = types.StringNull()
	}
	if value := res.Get("estContent.ignoreServerCertificateValidations"); value.Exists() {
		data.EstIgnoreServerCertificateValidation = types.BoolValue(value.Bool())
	} else {
		data.EstIgnoreServerCertificateValidation = types.BoolNull()
	}
	if value := res.Get("scepContent.enrollmentUrl"); value.Exists() {
		data.ScepEnrollmentUrl = types.StringValue(value.String())
	} else {
		data.ScepEnrollmentUrl = types.StringNull()
	}
	if value := res.Get("scepContent.retryPeriodInMins"); value.Exists() {
		data.ScepRetryPeriod = types.Int64Value(value.Int())
	} else {
		data.ScepRetryPeriod = types.Int64Null()
	}
	if value := res.Get("scepContent.retryCount"); value.Exists() {
		data.ScepRetryCount = types.Int64Value(value.Int())
	} else {
		data.ScepRetryCount = types.Int64Null()
	}
	if value := res.Get("scepContent.fingerprint"); value.Exists() {
		data.ScepFingerprint = types.StringValue(value.String())
	} else {
		data.ScepFingerprint = types.StringNull()
	}
	if value := res.Get("caCertificate.caOnly"); value.Exists() {
		data.ManualCaOnly = types.BoolValue(value.Bool())
	} else {
		data.ManualCaOnly = types.BoolNull()
	}
	if value := res.Get("caCertificate.base64Content"); value.Exists() {
		data.ManualCaCertificate = types.StringValue(value.String())
	} else {
		data.ManualCaCertificate = types.StringNull()
	}
	if value := res.Get("pkcs12Content.base64Certificate"); value.Exists() {
		data.Pkcs12Certificate = types.StringValue(value.String())
	} else {
		data.Pkcs12Certificate = types.StringNull()
	}
	if value := res.Get("certificateParameters.includeFQDN"); value.Exists() {
		data.IncludeFqdn = types.StringValue(value.String())
	} else {
		data.IncludeFqdn = types.StringNull()
	}
	if value := res.Get("certificateParameters.customFqdn"); value.Exists() {
		data.CustomFqdn = types.StringValue(value.String())
	} else {
		data.CustomFqdn = types.StringNull()
	}
	if value := res.Get("certificateParameters.includeDeviceIp"); value.Exists() {
		data.IncludeDeviceIp = types.StringValue(value.String())
	} else {
		data.IncludeDeviceIp = types.StringNull()
	}
	if value := res.Get("certificateParameters.commonName"); value.Exists() {
		data.CommonName = types.StringValue(value.String())
	} else {
		data.CommonName = types.StringNull()
	}
	if value := res.Get("certificateParameters.organizationalUnit"); value.Exists() {
		data.OrganizationalUnit = types.StringValue(value.String())
	} else {
		data.OrganizationalUnit = types.StringNull()
	}
	if value := res.Get("certificateParameters.organization"); value.Exists() {
		data.Organization = types.StringValue(value.String())
	} else {
		data.Organization = types.StringNull()
	}
	if value := res.Get("certificateParameters.locality"); value.Exists() {
		data.Locality = types.StringValue(value.String())
	} else {
		data.Locality = types.StringNull()
	}
	if value := res.Get("certificateParameters.state"); value.Exists() {
		data.State = types.StringValue(value.String())
	} else {
		data.State = types.StringNull()
	}
	if value := res.Get("certificateParameters.countryCode"); value.Exists() {
		data.CountryCode = types.StringValue(value.String())
	} else {
		data.CountryCode = types.StringNull()
	}
	if value := res.Get("certificateParameters.email"); value.Exists() {
		data.Email = types.StringValue(value.String())
	} else {
		data.Email = types.StringNull()
	}
	if value := res.Get("certificateParameters.includeDeviceSerial"); value.Exists() {
		data.IncludeDeviceSerialNumber = types.BoolValue(value.Bool())
	} else {
		data.IncludeDeviceSerialNumber = types.BoolNull()
	}
	if value := res.Get("key.keyType"); value.Exists() {
		data.KeyType = types.StringValue(value.String())
	} else {
		data.KeyType = types.StringNull()
	}
	if value := res.Get("key.keyName"); value.Exists() {
		data.KeyName = types.StringValue(value.String())
	} else {
		data.KeyName = types.StringNull()
	}
	if value := res.Get("key.keySize"); value.Exists() {
		data.KeySize = types.StringValue(value.String())
	} else {
		data.KeySize = types.StringNull()
	}
	if value := res.Get("key.ignoreIpsecKeyusage"); value.Exists() {
		data.IgnoreIpsecKeyUsage = types.BoolValue(value.Bool())
	} else {
		data.IgnoreIpsecKeyUsage = types.BoolNull()
	}
	if value := res.Get("revocation.certRevocationListContent.enableCertRevocationListDistributionPoint"); value.Exists() {
		data.CrlUseDistributionPointFromTheCertificate = types.BoolValue(value.Bool())
	} else {
		data.CrlUseDistributionPointFromTheCertificate = types.BoolNull()
	}
	if value := res.Get("revocation.certRevocationListContent.certRevocationStaticUrlList"); value.Exists() {
		data.CrlStaticUrlsList = helpers.GetStringList(value.Array())
	} else {
		data.CrlStaticUrlsList = types.ListNull(types.StringType)
	}
	if value := res.Get("revocation.onlineCertificateStatusProtocolContent.onlineCertificateStatusProtocolUrl"); value.Exists() {
		data.OcspUrl = types.StringValue(value.String())
	} else {
		data.OcspUrl = types.StringNull()
	}
	if value := res.Get("revocation.evaluationPriority"); value.Exists() {
		data.EvaluationPriority = types.StringValue(value.String())
	} else {
		data.EvaluationPriority = types.StringNull()
	}
	if value := res.Get("revocation.ignoreRevocation"); value.Exists() {
		data.ConsiderCertificateValidIfRevocationInformationNotReachable = types.BoolValue(value.Bool())
	} else {
		data.ConsiderCertificateValidIfRevocationInformationNotReachable = types.BoolNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *CertificateEnrollment) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("enrollmentType"); value.Exists() && !data.EnrollmentType.IsNull() {
		data.EnrollmentType = types.StringValue(value.String())
	} else {
		data.EnrollmentType = types.StringNull()
	}
	if value := res.Get("validationUsage.ipsecClient"); value.Exists() && !data.ValidationUsageIpsecClient.IsNull() {
		data.ValidationUsageIpsecClient = types.BoolValue(value.Bool())
	} else {
		data.ValidationUsageIpsecClient = types.BoolNull()
	}
	if value := res.Get("validationUsage.sslServer"); value.Exists() && !data.ValidationUsageSslServer.IsNull() {
		data.ValidationUsageSslServer = types.BoolValue(value.Bool())
	} else {
		data.ValidationUsageSslServer = types.BoolNull()
	}
	if value := res.Get("validationUsage.sslClient"); value.Exists() && !data.ValidationUsageSslClient.IsNull() {
		data.ValidationUsageSslClient = types.BoolValue(value.Bool())
	} else {
		data.ValidationUsageSslClient = types.BoolNull()
	}
	if value := res.Get("estContent.enrollmentUrl"); value.Exists() && !data.EstEnrollmentUrl.IsNull() {
		data.EstEnrollmentUrl = types.StringValue(value.String())
	} else {
		data.EstEnrollmentUrl = types.StringNull()
	}
	if value := res.Get("estContent.username"); value.Exists() && !data.EstUsername.IsNull() {
		data.EstUsername = types.StringValue(value.String())
	} else {
		data.EstUsername = types.StringNull()
	}
	if value := res.Get("estContent.fingerprint"); value.Exists() && !data.EstFingerprint.IsNull() {
		data.EstFingerprint = types.StringValue(value.String())
	} else {
		data.EstFingerprint = types.StringNull()
	}
	if value := res.Get("estContent.sourceInterface.id"); value.Exists() && !data.EstSourceInterfaceId.IsNull() {
		data.EstSourceInterfaceId = types.StringValue(value.String())
	} else {
		data.EstSourceInterfaceId = types.StringNull()
	}
	if value := res.Get("estContent.sourceInterface.name"); value.Exists() && !data.EstSourceInterfaceName.IsNull() {
		data.EstSourceInterfaceName = types.StringValue(value.String())
	} else {
		data.EstSourceInterfaceName = types.StringNull()
	}
	if value := res.Get("estContent.ignoreServerCertificateValidations"); value.Exists() && !data.EstIgnoreServerCertificateValidation.IsNull() {
		data.EstIgnoreServerCertificateValidation = types.BoolValue(value.Bool())
	} else {
		data.EstIgnoreServerCertificateValidation = types.BoolNull()
	}
	if value := res.Get("scepContent.enrollmentUrl"); value.Exists() && !data.ScepEnrollmentUrl.IsNull() {
		data.ScepEnrollmentUrl = types.StringValue(value.String())
	} else {
		data.ScepEnrollmentUrl = types.StringNull()
	}
	if value := res.Get("scepContent.retryPeriodInMins"); value.Exists() && !data.ScepRetryPeriod.IsNull() {
		data.ScepRetryPeriod = types.Int64Value(value.Int())
	} else {
		data.ScepRetryPeriod = types.Int64Null()
	}
	if value := res.Get("scepContent.retryCount"); value.Exists() && !data.ScepRetryCount.IsNull() {
		data.ScepRetryCount = types.Int64Value(value.Int())
	} else {
		data.ScepRetryCount = types.Int64Null()
	}
	if value := res.Get("scepContent.fingerprint"); value.Exists() && !data.ScepFingerprint.IsNull() {
		data.ScepFingerprint = types.StringValue(value.String())
	} else {
		data.ScepFingerprint = types.StringNull()
	}
	if value := res.Get("caCertificate.caOnly"); value.Exists() && !data.ManualCaOnly.IsNull() {
		data.ManualCaOnly = types.BoolValue(value.Bool())
	} else {
		data.ManualCaOnly = types.BoolNull()
	}
	if value := res.Get("caCertificate.base64Content"); value.Exists() && !data.ManualCaCertificate.IsNull() {
		data.ManualCaCertificate = types.StringValue(value.String())
	} else {
		data.ManualCaCertificate = types.StringNull()
	}
	if value := res.Get("pkcs12Content.base64Certificate"); value.Exists() && !data.Pkcs12Certificate.IsNull() {
		data.Pkcs12Certificate = types.StringValue(value.String())
	} else {
		data.Pkcs12Certificate = types.StringNull()
	}
	if value := res.Get("certificateParameters.includeFQDN"); value.Exists() && !data.IncludeFqdn.IsNull() {
		data.IncludeFqdn = types.StringValue(value.String())
	} else {
		data.IncludeFqdn = types.StringNull()
	}
	if value := res.Get("certificateParameters.customFqdn"); value.Exists() && !data.CustomFqdn.IsNull() {
		data.CustomFqdn = types.StringValue(value.String())
	} else {
		data.CustomFqdn = types.StringNull()
	}
	if value := res.Get("certificateParameters.includeDeviceIp"); value.Exists() && !data.IncludeDeviceIp.IsNull() {
		data.IncludeDeviceIp = types.StringValue(value.String())
	} else {
		data.IncludeDeviceIp = types.StringNull()
	}
	if value := res.Get("certificateParameters.commonName"); value.Exists() && !data.CommonName.IsNull() {
		data.CommonName = types.StringValue(value.String())
	} else {
		data.CommonName = types.StringNull()
	}
	if value := res.Get("certificateParameters.organizationalUnit"); value.Exists() && !data.OrganizationalUnit.IsNull() {
		data.OrganizationalUnit = types.StringValue(value.String())
	} else {
		data.OrganizationalUnit = types.StringNull()
	}
	if value := res.Get("certificateParameters.organization"); value.Exists() && !data.Organization.IsNull() {
		data.Organization = types.StringValue(value.String())
	} else {
		data.Organization = types.StringNull()
	}
	if value := res.Get("certificateParameters.locality"); value.Exists() && !data.Locality.IsNull() {
		data.Locality = types.StringValue(value.String())
	} else {
		data.Locality = types.StringNull()
	}
	if value := res.Get("certificateParameters.state"); value.Exists() && !data.State.IsNull() {
		data.State = types.StringValue(value.String())
	} else {
		data.State = types.StringNull()
	}
	if value := res.Get("certificateParameters.countryCode"); value.Exists() && !data.CountryCode.IsNull() {
		data.CountryCode = types.StringValue(value.String())
	} else {
		data.CountryCode = types.StringNull()
	}
	if value := res.Get("certificateParameters.email"); value.Exists() && !data.Email.IsNull() {
		data.Email = types.StringValue(value.String())
	} else {
		data.Email = types.StringNull()
	}
	if value := res.Get("certificateParameters.includeDeviceSerial"); value.Exists() && !data.IncludeDeviceSerialNumber.IsNull() {
		data.IncludeDeviceSerialNumber = types.BoolValue(value.Bool())
	} else {
		data.IncludeDeviceSerialNumber = types.BoolNull()
	}
	if value := res.Get("key.keyType"); value.Exists() && !data.KeyType.IsNull() {
		data.KeyType = types.StringValue(value.String())
	} else {
		data.KeyType = types.StringNull()
	}
	if value := res.Get("key.keyName"); value.Exists() && !data.KeyName.IsNull() {
		data.KeyName = types.StringValue(value.String())
	} else {
		data.KeyName = types.StringNull()
	}
	if value := res.Get("key.keySize"); value.Exists() && !data.KeySize.IsNull() {
		data.KeySize = types.StringValue(value.String())
	} else {
		data.KeySize = types.StringNull()
	}
	if value := res.Get("key.ignoreIpsecKeyusage"); value.Exists() && !data.IgnoreIpsecKeyUsage.IsNull() {
		data.IgnoreIpsecKeyUsage = types.BoolValue(value.Bool())
	} else {
		data.IgnoreIpsecKeyUsage = types.BoolNull()
	}
	if value := res.Get("revocation.certRevocationListContent.enableCertRevocationListDistributionPoint"); value.Exists() && !data.CrlUseDistributionPointFromTheCertificate.IsNull() {
		data.CrlUseDistributionPointFromTheCertificate = types.BoolValue(value.Bool())
	} else {
		data.CrlUseDistributionPointFromTheCertificate = types.BoolNull()
	}
	if value := res.Get("revocation.certRevocationListContent.certRevocationStaticUrlList"); value.Exists() && !data.CrlStaticUrlsList.IsNull() {
		data.CrlStaticUrlsList = helpers.GetStringList(value.Array())
	} else {
		data.CrlStaticUrlsList = types.ListNull(types.StringType)
	}
	if value := res.Get("revocation.onlineCertificateStatusProtocolContent.onlineCertificateStatusProtocolUrl"); value.Exists() && !data.OcspUrl.IsNull() {
		data.OcspUrl = types.StringValue(value.String())
	} else {
		data.OcspUrl = types.StringNull()
	}
	if value := res.Get("revocation.evaluationPriority"); value.Exists() && !data.EvaluationPriority.IsNull() {
		data.EvaluationPriority = types.StringValue(value.String())
	} else {
		data.EvaluationPriority = types.StringNull()
	}
	if value := res.Get("revocation.ignoreRevocation"); value.Exists() && !data.ConsiderCertificateValidIfRevocationInformationNotReachable.IsNull() {
		data.ConsiderCertificateValidIfRevocationInformationNotReachable = types.BoolValue(value.Bool())
	} else {
		data.ConsiderCertificateValidIfRevocationInformationNotReachable = types.BoolNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *CertificateEnrollment) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

func (data CertificateEnrollment) adjustBody(ctx context.Context, req string) string {

	// Enable CRL if either distribution point or static URLs are provided
	if !data.CrlUseDistributionPointFromTheCertificate.IsNull() || !!data.CrlStaticUrlsList.IsNull() {
		req, _ = sjson.Set(req, "revocation.certRevocationListContent.enableCertRevocationList", "true")
	}

	// If CRL static URLs are provided, enable static CRL
	if !data.CrlStaticUrlsList.IsNull() {
		req, _ = sjson.Set(req, "revocation.certRevocationListContent.enableStaticCertRevocationList", "true")
	}

	// If OSCP URL is provided, enable OCSP
	if !data.OcspUrl.IsNull() {
		req, _ = sjson.Set(req, "revocation.onlineCertificateStatusProtocolContent.enableOnlineCertificateStatusProtocol", "true")
	}

	return req
}
