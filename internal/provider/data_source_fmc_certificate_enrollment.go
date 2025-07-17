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

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &CertificateEnrollmentDataSource{}
	_ datasource.DataSourceWithConfigure = &CertificateEnrollmentDataSource{}
)

func NewCertificateEnrollmentDataSource() datasource.DataSource {
	return &CertificateEnrollmentDataSource{}
}

type CertificateEnrollmentDataSource struct {
	client *fmc.Client
}

func (d *CertificateEnrollmentDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_certificate_enrollment"
}

func (d *CertificateEnrollmentDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source reads the Certificate Enrollment.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the object",
				Optional:            true,
				Computed:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "Name of the FMC domain",
				Optional:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the certificate.",
				Optional:            true,
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the object; this value is always 'CertEnrollment'.",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the certificate.",
				Computed:            true,
			},
			"enrollment_type": schema.StringAttribute{
				MarkdownDescription: "Type of enrollment for the certificate.",
				Computed:            true,
			},
			"validate_ipsec_client": schema.BoolAttribute{
				MarkdownDescription: "Whether the certificate is used for IPsec client validation.",
				Computed:            true,
			},
			"validate_ssl_server": schema.BoolAttribute{
				MarkdownDescription: "Whether the certificate is used for SSL server validation.",
				Computed:            true,
			},
			"validate_ssl_client": schema.BoolAttribute{
				MarkdownDescription: "Whether the certificate is used for SSL client validation.",
				Computed:            true,
			},
			"skip_ca_flag_check": schema.BoolAttribute{
				MarkdownDescription: "Whether to skip the CA flag check for the certificate.",
				Computed:            true,
			},
			"est_enrollment_url": schema.StringAttribute{
				MarkdownDescription: "URL for the EST enrollment.",
				Computed:            true,
			},
			"est_username": schema.StringAttribute{
				MarkdownDescription: "Username for the EST enrollment.",
				Computed:            true,
			},
			"est_password": schema.StringAttribute{
				MarkdownDescription: "Password for the EST enrollment.",
				Computed:            true,
				Sensitive:           true,
			},
			"est_fingerprint": schema.StringAttribute{
				MarkdownDescription: "Fingerprint for the EST enrollment.",
				Computed:            true,
			},
			"est_source_interface_id": schema.StringAttribute{
				MarkdownDescription: "Source interface group or security zone for the EST enrollment.",
				Computed:            true,
			},
			"est_source_interface_name": schema.StringAttribute{
				MarkdownDescription: "Source interface group or security zone for the EST enrollment.",
				Computed:            true,
			},
			"est_ignore_server_certificate_validation": schema.BoolAttribute{
				MarkdownDescription: "Whether to ignore server certificate validations for the EST enrollment.",
				Computed:            true,
			},
			"scep_enrollment_url": schema.StringAttribute{
				MarkdownDescription: "URL for the SCEP enrollment.",
				Computed:            true,
			},
			"scep_challenge_password": schema.StringAttribute{
				MarkdownDescription: "Challenge password for the SCEP enrollment.",
				Computed:            true,
				Sensitive:           true,
			},
			"scep_retry_period": schema.Int64Attribute{
				MarkdownDescription: "Retry period in minutes for the SCEP enrollment.",
				Computed:            true,
			},
			"scep_retry_count": schema.Int64Attribute{
				MarkdownDescription: "Number of retries for the SCEP enrollment.",
				Computed:            true,
			},
			"scep_fingerprint": schema.StringAttribute{
				MarkdownDescription: "Fingerprint for the SCEP enrollment.",
				Computed:            true,
			},
			"manual_ca_only": schema.BoolAttribute{
				MarkdownDescription: "Whether the certificate is a CA only certificate.",
				Computed:            true,
			},
			"manual_ca_certificate": schema.StringAttribute{
				MarkdownDescription: "Base64 encoded content of the CA certificate.",
				Computed:            true,
			},
			"pkcs12_certificate": schema.StringAttribute{
				MarkdownDescription: "Base64 encoded content of the PKCS12 certificate.",
				Computed:            true,
			},
			"pkcs12_certificate_passphrase": schema.StringAttribute{
				MarkdownDescription: "Passphrase for the PKCS12 certificate.",
				Computed:            true,
				Sensitive:           true,
			},
			"include_fqdn": schema.StringAttribute{
				MarkdownDescription: "How the FQDN is included in the certificate.",
				Computed:            true,
			},
			"custom_fqdn": schema.StringAttribute{
				MarkdownDescription: "Custom FQDN to be included in the certificate.",
				Computed:            true,
			},
			"include_device_ip": schema.StringAttribute{
				MarkdownDescription: "Device IP in the certificate.",
				Computed:            true,
			},
			"common_name": schema.StringAttribute{
				MarkdownDescription: "Common Name for the certificate.",
				Computed:            true,
			},
			"organizational_unit": schema.StringAttribute{
				MarkdownDescription: "Organizational Unit for the certificate.",
				Computed:            true,
			},
			"organization": schema.StringAttribute{
				MarkdownDescription: "Organization for the certificate.",
				Computed:            true,
			},
			"locality": schema.StringAttribute{
				MarkdownDescription: "Locality for the certificate.",
				Computed:            true,
			},
			"state": schema.StringAttribute{
				MarkdownDescription: "State for the certificate.",
				Computed:            true,
			},
			"country_code": schema.StringAttribute{
				MarkdownDescription: "Country Code for the certificate.",
				Computed:            true,
			},
			"email": schema.StringAttribute{
				MarkdownDescription: "Email for the certificate.",
				Computed:            true,
			},
			"include_device_serial_number": schema.BoolAttribute{
				MarkdownDescription: "Whether to include the device serial in the certificate.",
				Computed:            true,
			},
			"key_type": schema.StringAttribute{
				MarkdownDescription: "Type of key used for the certificate.",
				Computed:            true,
			},
			"key_name": schema.StringAttribute{
				MarkdownDescription: "Name of the key used for the certificate.",
				Computed:            true,
			},
			"key_size": schema.StringAttribute{
				MarkdownDescription: "Size of the key used for the certificate.",
				Computed:            true,
			},
			"ignore_ipsec_key_usage": schema.BoolAttribute{
				MarkdownDescription: "Whether to ignore IPsec key usage for the certificate.",
				Computed:            true,
			},
			"enable_certificate_revocation_list": schema.BoolAttribute{
				MarkdownDescription: "Whether to enable certificate revocation list.",
				Computed:            true,
			},
			"use_crl_distribution_point_from_the_certificate": schema.BoolAttribute{
				MarkdownDescription: "Whether to use CRL distribution point from the certificate.",
				Computed:            true,
			},
			"enable_static_certificate_revocation_list": schema.BoolAttribute{
				MarkdownDescription: "Whether to enable static certificate revocation list.",
				Computed:            true,
			},
			"cert_revocation_static_url_list": schema.ListAttribute{
				MarkdownDescription: "Static URL list for certificate revocation.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"enable_online_certificate_status_protocol": schema.BoolAttribute{
				MarkdownDescription: "Whether to enable Online Certificate Status Protocol (OCSP).",
				Computed:            true,
			},
			"online_certificate_status_protocol_url": schema.StringAttribute{
				MarkdownDescription: "URL for the Online Certificate Status Protocol (OCSP).",
				Computed:            true,
			},
			"evaluation_priority": schema.StringAttribute{
				MarkdownDescription: "Priority for certificate revocation evaluation. Needs to be set if both CRL and OCSP are enabled.",
				Computed:            true,
			},
			"ignore_revocation": schema.BoolAttribute{
				MarkdownDescription: "Consider the certificate valid if revocation information can not be reached.",
				Computed:            true,
			},
		},
	}
}
func (d *CertificateEnrollmentDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *CertificateEnrollmentDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *CertificateEnrollmentDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config CertificateEnrollment

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !config.Domain.IsNull() && config.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(config.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))
	if config.Id.IsNull() && !config.Name.IsNull() {
		offset := 0
		limit := 1000
		for page := 1; ; page++ {
			queryString := fmt.Sprintf("?limit=%d&offset=%d&expanded=true", limit, offset)
			res, err := d.client.Get(config.getPath()+queryString, reqMods...)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects, got error: %s", err))
				return
			}
			if value := res.Get("items"); len(value.Array()) > 0 {
				value.ForEach(func(k, v gjson.Result) bool {
					if config.Name.ValueString() == v.Get("name").String() {
						config.Id = types.StringValue(v.Get("id").String())
						tflog.Debug(ctx, fmt.Sprintf("%s: Found object with name '%v', id: %v", config.Id.ValueString(), config.Name.ValueString(), config.Id.ValueString()))
						return false
					}
					return true
				})
			}
			if !config.Id.IsNull() || !res.Get("paging.next.0").Exists() {
				break
			}
			offset += limit
		}

		if config.Id.IsNull() {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with name: %v", config.Name.ValueString()))
			return
		}
	}
	urlPath := config.getPath() + "/" + url.QueryEscape(config.Id.ValueString())
	res, err := d.client.Get(urlPath, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
