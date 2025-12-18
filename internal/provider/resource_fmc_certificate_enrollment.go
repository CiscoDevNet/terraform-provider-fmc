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
	"regexp"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &CertificateEnrollmentResource{}
	_ resource.ResourceWithImportState = &CertificateEnrollmentResource{}
)

func NewCertificateEnrollmentResource() resource.Resource {
	return &CertificateEnrollmentResource{}
}

type CertificateEnrollmentResource struct {
	client *fmc.Client
}

func (r *CertificateEnrollmentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_certificate_enrollment"
}

func (r *CertificateEnrollmentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This device manages Certificate Enrollment configuration.\n Manual certificate creation is supported only with `manual_ca_only` set to `true`.\n (FMC 7.2) Only PKCS12 based certificate enrollment object is supported.\n (FMC 7.4 and FMC 7.6) Only PKCS12 and MANUAL with CA only based certificate enrollment object is supported.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "Name of the FMC domain",
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the Certificate Enrollment object.").String,
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this value is always 'CertEnrollment'.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description of the Certificate Enrollment object.").String,
				Optional:            true,
			},
			"enrollment_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Certificate enrollment type.").AddStringEnumDescription("SCEP", "ACME", "EST", "MANUAL", "SELF_SIGNED_CERTFICATE", "PKCS12").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("SCEP", "ACME", "EST", "MANUAL", "SELF_SIGNED_CERTFICATE", "PKCS12"),
				},
			},
			"validation_usage_ipsec_client": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Validate an IPsec client certificate for a site-to-site VPN connection.").String,
				Optional:            true,
			},
			"validation_usage_ssl_server": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Validate an SSL server certificate.").String,
				Optional:            true,
			},
			"validation_usage_ssl_client": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Validate an SSL client certificate during a remote access VPN connection attempt.").String,
				Optional:            true,
			},
			"skip_ca_flag_check": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Skip checking the basic constraints extension and the CA flag in a trustpoint certificate.").String,
				Optional:            true,
			},
			"est_enrollment_url": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("EST enrollment CA server URL.").String,
				Optional:            true,
			},
			"est_username": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("EST enrollment CA server username.").String,
				Optional:            true,
			},
			"est_password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("EST enrollment CA server password.").String,
				Optional:            true,
				Sensitive:           true,
			},
			"est_fingerprint": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("EST enrollment CA server fingerprint.").String,
				Optional:            true,
			},
			"est_source_interface_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of interface group or security zone that interacts with the CA server.").String,
				Optional:            true,
			},
			"est_source_interface_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of interface group or security zone that interacts with the CA server.").String,
				Optional:            true,
			},
			"est_ignore_server_certificate_validation": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Ignore EST server certificate validations.").String,
				Optional:            true,
			},
			"scep_enrollment_url": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SCEP enrollment CA server URL.").String,
				Optional:            true,
			},
			"scep_challenge_password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SCEP enrollment challenge password.").String,
				Optional:            true,
				Sensitive:           true,
			},
			"scep_retry_period": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interval (in minutes) between certificate request attempts.").AddIntegerRangeDescription(1, 60).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 60),
				},
			},
			"scep_retry_count": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Number of retries that should be made if no certificate is issued upon the first request.").AddIntegerRangeDescription(0, 100).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 100),
				},
			},
			"scep_fingerprint": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SCEP enrollment CA server fingerprint.").String,
				Optional:            true,
			},
			"manual_ca_only": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Create only the CA certificate from the selected CA. An identity certificate will not be created for this certificate. Must be set to `true`.").String,
				Optional:            true,
			},
			"manual_ca_certificate": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Base64 encoded certificate in PEM format.").String,
				Optional:            true,
			},
			"pkcs12_certificate": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Base64 encoded certificate in PKCS12 format.").String,
				Optional:            true,
			},
			"pkcs12_certificate_passphrase": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Passphrase for the PKCS12 certificate.").String,
				Optional:            true,
				Sensitive:           true,
			},
			"acme_enrollment_url": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ACME enrollment CA server URL.").String,
				Optional:            true,
			},
			"acme_authentication_protocol": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ACME authentication protocol.").AddStringEnumDescription("HTTP01").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("HTTP01"),
				},
			},
			"acme_authentication_interface_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of security zone or interface group that has the interface through which the ACME server communicates with the device to verify ownership of the domain.").String,
				Optional:            true,
			},
			"acme_authentication_interface_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of security zone or interface group that has the interface through which the ACME server communicates with the device to verify ownership of the domain.").String,
				Optional:            true,
			},
			"acme_source_interface_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of security zone or interface group that has the interface through which the device interacts with the ACME server to request and receive the enrolled ACME certificate.").String,
				Optional:            true,
			},
			"acme_source_interface_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of security zone or interface group that has the interface through which the device interacts with the ACME server to request and receive the enrolled ACME certificate.").String,
				Optional:            true,
			},
			"acme_ca_only_certificate_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of manually installed CA-only certificate that authenticates the ACME server.").String,
				Optional:            true,
			},
			"acme_auto_enrollment": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable enable automatic enrollment of the ACME certificates based on the configured lifetime.").String,
				Optional:            true,
			},
			"acme_auto_enrollment_lifetime": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddIntegerRangeDescription(10, 99).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(10, 99),
				},
			},
			"acme_auto_enrollment_key_regeneration": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Regenerate a new key for each ACME enrollment.").String,
				Optional:            true,
			},
			"certificate_include_fqdn": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Include the device's fully qualified domain name (FQDN) in the certificate request").AddStringEnumDescription("DEVICE_HOSTNAME", "NONE", "CUSTOM", "DEFAULT").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("DEVICE_HOSTNAME", "NONE", "CUSTOM", "DEFAULT"),
				},
			},
			"certificate_custom_fqdn": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Device's custom FQDN to be included in the certificate.").String,
				Optional:            true,
			},
			"certificate_alternate_fqdn": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Alternate FQDN.").String,
				Optional:            true,
			},
			"certificate_include_device_ip": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Device IP in the certificate.").String,
				Optional:            true,
			},
			"certificate_common_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Common Name (CN) for the certificate.").String,
				Optional:            true,
			},
			"certificate_organizational_unit": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Organizational Unit (OU) for the certificate.").String,
				Optional:            true,
			},
			"certificate_organization": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Organization (O) for the certificate.").String,
				Optional:            true,
			},
			"certificate_locality": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Locality (L) for the certificate.").String,
				Optional:            true,
			},
			"certificate_state": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("State (ST) for the certificate.").String,
				Optional:            true,
			},
			"certificate_country_code": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Country Code (C) for the certificate.").String,
				Optional:            true,
			},
			"certificate_email": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Email (E) for the certificate.").String,
				Optional:            true,
			},
			"certificate_include_device_serial_number": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Include the device serial in the certificate.").String,
				Optional:            true,
			},
			"key_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of key pair.").AddStringEnumDescription("RSA", "ECDSA", "EdDSA").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("RSA", "ECDSA", "EdDSA"),
				},
			},
			"key_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the key pair used.").String,
				Optional:            true,
			},
			"key_size": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Desired key size (modulus), in bits.").AddStringEnumDescription("CertKey_512", "CertKey_768", "CertKey_1024", "CertKey_2048", "CertKey_3072", "CertKey_4096", "CertKey_256", "CertKey_384", "CertKey_521").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("CertKey_512", "CertKey_768", "CertKey_1024", "CertKey_2048", "CertKey_3072", "CertKey_4096", "CertKey_256", "CertKey_384", "CertKey_521"),
				},
			},
			"ignore_ipsec_key_usage": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Do not validate values in the key usage and extended key usage extensions of IPsec remote client certificates.").String,
				Optional:            true,
			},
			"crl_use_distribution_point_from_the_certificate": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Obtain the revocation lists distribution URL from the certificate.").String,
				Optional:            true,
			},
			"crl_static_urls": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Static URL list for certificate revocation.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"ocsp_url": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("URL for the Online Certificate Status Protocol (OCSP).").String,
				Optional:            true,
			},
			"revocation_evaluation_priority": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Priority for certificate revocation evaluation. Needs to be set if both CRL and OCSP are enabled.").AddStringEnumDescription("CRL", "OCSP", "NONE").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("CRL", "OCSP", "NONE"),
				},
			},
			"consider_certificate_valid_if_revocation_information_not_reachable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Consider the certificate valid if revocation information can not be reached.").String,
				Optional:            true,
			},
		},
	}
}

func (r *CertificateEnrollmentResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *CertificateEnrollmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CertificateEnrollment

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !plan.Domain.IsNull() && plan.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(plan.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, CertificateEnrollment{})
	body = plan.adjustBody(ctx, body)
	res, err := r.client.Post(plan.getPath(), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("id").String())
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

func (r *CertificateEnrollmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CertificateEnrollment

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !state.Domain.IsNull() && state.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(state.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	urlPath := state.getPath() + "/" + url.QueryEscape(state.Id.ValueString())
	res, err := r.client.Get(urlPath, reqMods...)

	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	imp, diags := helpers.IsFlagImporting(ctx, req)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// After `terraform import` we switch to a full read.
	if imp {
		state.fromBody(ctx, res)
	} else {
		state.fromBodyPartial(ctx, res)
	}

	// FMCBUG CSCwq44439
	// FMC API adds '\r\n' to the end of PKCS12 certificate, so we need to remove it.
	if !state.Pkcs12Certificate.IsNull() && state.Pkcs12Certificate.ValueString() != "" {
		state.Pkcs12Certificate = types.StringValue(strings.TrimSuffix(state.Pkcs12Certificate.ValueString(), "\r\n"))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// Section below is generated&owned by "gen/generator.go". //template:begin update

func (r *CertificateEnrollmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state CertificateEnrollment

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Read state
	diags = req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !plan.Domain.IsNull() && plan.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(plan.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	body := plan.toBody(ctx, state)
	body = plan.adjustBody(ctx, body)
	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *CertificateEnrollmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CertificateEnrollment

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !state.Domain.IsNull() && state.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(state.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	res, err := r.client.Delete(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString()), reqMods...)
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *CertificateEnrollmentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Parse import ID
	var inputPattern = regexp.MustCompile(`^(?:(?P<domain>[^\s,]+),)?(?P<id>[^\s,]+?)$`)
	match := inputPattern.FindStringSubmatch(req.ID)
	if match == nil {
		errMsg := "Failed to parse import parameters.\nPlease provide import string in the following format: <domain>,<id>\n<domain> is optional. If not provided, `Global` is used implicitly and resource's `domain` attribute is not set.\n" + fmt.Sprintf("Got: %q", req.ID)
		resp.Diagnostics.AddError("Import error", errMsg)
		return
	}

	// Set domain, if provided
	if tmpDomain := match[inputPattern.SubexpIndex("domain")]; tmpDomain != "" {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("domain"), tmpDomain)...)
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), match[inputPattern.SubexpIndex("id")])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
