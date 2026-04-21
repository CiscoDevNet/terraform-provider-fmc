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
	"sort"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &IdentityPolicyResource{}
	_ resource.ResourceWithImportState = &IdentityPolicyResource{}
)

func NewIdentityPolicyResource() resource.Resource {
	return &IdentityPolicyResource{}
}

type IdentityPolicyResource struct {
	client *fmc.Client
}

func (r *IdentityPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_identity_policy"
}

func (r *IdentityPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages Identity Policy with corresponding Categories and Rules.").AddMinimumVersionHeaderDescription().AddMinimumVersionDescription("7.2").AddMinimumVersionCreateDescription("10.0").String,

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
				MarkdownDescription: helpers.NewAttributeDescription("Name of the Identity Policy.").String,
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this value is always 'IdentityPolicy'.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description of the Identity Policy.").String,
				Optional:            true,
			},
			"identity_mapping_filter_network_object_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the Host, Network or Network Group used as an Identity Mapping Filter.").String,
				Optional:            true,
			},
			"identity_mapping_filter_network_object_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the Host, Network or Network Group used as an Identity Mapping Filter.").String,
				Optional:            true,
			},
			"active_authentication_server_certificate_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the Internal Certificate object used as an active authentication server certificate.").String,
				Optional:            true,
			},
			"active_authentication_server_certificate_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the Internal Certificate object used as an active authentication server certificate.").AddStringEnumDescription("PKI_InternalCert").AddDefaultValueDescription("PKI_InternalCert").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("PKI_InternalCert"),
				},
				Default: stringdefault.StaticString("PKI_InternalCert"),
			},
			"active_authentication_redirect_fqdn_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the FQDN object used for active authentication redirection (Redirect to Host Name). Supported only in Snort 3.0 or later.").String,
				Optional:            true,
			},
			"active_authentication_redirect_fqdn_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the FQDN object used for active authentication redirection.").AddStringEnumDescription("FQDN").AddDefaultValueDescription("FQDN").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("FQDN"),
				},
				Default: stringdefault.StaticString("FQDN"),
			},
			"active_authentication_redirect_port": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Port used for active authentication redirection. Valid values are 885 or 1025-65535.").AddIntegerRangeDescription(885, 65535).AddDefaultValueDescription("885").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(885, 65535),
				},
				Default: int64default.StaticInt64(885),
			},
			"active_authentication_maximum_login_attempts": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Maximum login attempts. Use 0 for unlimited attempts.").AddDefaultValueDescription("3").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(3),
			},
			"active_authentication_session_sharing": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Share active authentication sessions across firewalls. Supported only in 7.4.1 and later.").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"active_authentication_page": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Active authentication page type.").AddStringEnumDescription("DEFAULT", "CUSTOM").AddDefaultValueDescription("DEFAULT").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("DEFAULT", "CUSTOM"),
				},
				Default: stringdefault.StaticString("DEFAULT"),
			},
			"active_authentication_page_html": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Custom HTML content for the active authentication login page. Only used when `active_authentication_page` is set to 'CUSTOM'.").AddDefaultValueDescription("<!DOCTYPE html>\n<html>\n<head>\n<meta http-equiv=\"content-type\" content=\"text/html; charset=UTF-8\" />\n<title>Login</title>\n<style type=\"text/css\">\nbody {\n    margin:0;\n    font-family:verdana,sans-serif;\n}\nh1 {\n    margin:0;\n    padding:12px 25px;\n    background-color:#343434;\n    color:#ddd;\n}\np {\n    margin:12px 25px;\n}\nstrong {\n    color:#E0042D;\n}\ndiv {\n    padding-left:23px;\n    font-weight: normal;\n    font-size: 8pt;\n}\ninput {\n    margin:12px 25px;\n}\n</style>\n</head>\n<body>\n    <form action=\"\" id=\"loginForm\" method=\"post\" name=\"loginForm\">\n        <h1>Login</h1>\n        <p><strong>Please enter your username and password or log in as a guest.</strong></p>\n        <div class=\"label\">Username\n            <input id=\"name\" maxlength=\"100\" name=\"login\" type=\"text\" value=\"\"/>\n            <b>realm</b>\n            <select name=\"realm\" id=\"realm\"></select>\n        </div>\n        <div class=\"label\" id=\"label-password\">Password\n            <input id=\"pass\" name=\"pass\" type=\"password\" value=\"\"/>\n        </div>\n        <input id=\"submit-btn\" type=\"submit\" name=\"login_action\" value=\"Submit\"/>\n        <p><font size=2 >-or-</font></p>\n        <input id=\"login-btn\" type=\"submit\" name=\"guest_action\" value=\"Login as guest\"/>\n    </form>\n</body>\n</html>\n").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("<!DOCTYPE html>\n<html>\n<head>\n<meta http-equiv=\"content-type\" content=\"text/html; charset=UTF-8\" />\n<title>Login</title>\n<style type=\"text/css\">\nbody {\n    margin:0;\n    font-family:verdana,sans-serif;\n}\nh1 {\n    margin:0;\n    padding:12px 25px;\n    background-color:#343434;\n    color:#ddd;\n}\np {\n    margin:12px 25px;\n}\nstrong {\n    color:#E0042D;\n}\ndiv {\n    padding-left:23px;\n    font-weight: normal;\n    font-size: 8pt;\n}\ninput {\n    margin:12px 25px;\n}\n</style>\n</head>\n<body>\n    <form action=\"\" id=\"loginForm\" method=\"post\" name=\"loginForm\">\n        <h1>Login</h1>\n        <p><strong>Please enter your username and password or log in as a guest.</strong></p>\n        <div class=\"label\">Username\n            <input id=\"name\" maxlength=\"100\" name=\"login\" type=\"text\" value=\"\"/>\n            <b>realm</b>\n            <select name=\"realm\" id=\"realm\"></select>\n        </div>\n        <div class=\"label\" id=\"label-password\">Password\n            <input id=\"pass\" name=\"pass\" type=\"password\" value=\"\"/>\n        </div>\n        <input id=\"submit-btn\" type=\"submit\" name=\"login_action\" value=\"Submit\"/>\n        <p><font size=2 >-or-</font></p>\n        <input id=\"login-btn\" type=\"submit\" name=\"guest_action\" value=\"Login as guest\"/>\n    </form>\n</body>\n</html>\n"),
			},
			"categories": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Ordered list of Categories created between default categories - Standard Rules and Root Rules. Do not include default categories in the list.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the Category.").String,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the Category.").String,
							Required:            true,
						},
					},
				},
			},
			"rules": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Ordered list of Rules.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the Identity Rule.").String,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the Identity Rule.").String,
							Required:            true,
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable rule.").String,
							Optional:            true,
						},
						"category": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the Category the rule belongs to. Can be one of the default categories (Administrator, Standard or Root Rules) or user-defined one.").String,
							Required:            true,
						},
						"authentication_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Authentication type action for the rule.").AddStringEnumDescription("PASSIVE", "ACTIVE", "NO").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("PASSIVE", "ACTIVE", "NO"),
							},
						},
						"authentication_realm_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the Authentication Realm (AD/LDAP).").String,
							Optional:            true,
						},
						"authentication_realm_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of the Authentication Realm.").AddStringEnumDescription("IdentityRealm").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("IdentityRealm"),
							},
						},
						"guest_access_fallback": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Identify as Special Identities/Guest if user cannot be authenticated.").String,
							Optional:            true,
						},
						"active_authentication_fallback": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Use active authentication if passive or VPN identity cannot be established.").String,
							Optional:            true,
						},
						"authentication_protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Protocol used for active authentication.").AddStringEnumDescription("HTTP_BASIC", "NTLM", "KERBEROS", "HTTP NEGOTIATE", "HTTP RESPONSE PAGE", "SAML").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("HTTP_BASIC", "NTLM", "KERBEROS", "HTTP NEGOTIATE", "HTTP RESPONSE PAGE", "SAML"),
							},
						},
						"excluded_user_agents": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of excluded User Agents (Applications).").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the Application object.").String,
										Required:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of the Application object.").AddStringEnumDescription("Application").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("Application"),
										},
									},
								},
							},
						},
						"source_zones": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source Security Zones.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the Security Zone object.").String,
										Required:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of the Security Zone object.").AddStringEnumDescription("SecurityZone").AddDefaultValueDescription("SecurityZone").String,
										Optional:            true,
										Computed:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("SecurityZone"),
										},
										Default: stringdefault.StaticString("SecurityZone"),
									},
								},
							},
						},
						"destination_zones": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Destination Security Zones.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the Security Zone object.").String,
										Required:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of the Security Zone object.").AddStringEnumDescription("SecurityZone").AddDefaultValueDescription("SecurityZone").String,
										Optional:            true,
										Computed:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("SecurityZone"),
										},
										Default: stringdefault.StaticString("SecurityZone"),
									},
								},
							},
						},
						"source_network_literals": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source Network literals.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of the Network Object.").AddStringEnumDescription("Host", "Network").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("Host", "Network"),
										},
									},
									"value": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("IP address or network in CIDR format.").String,
										Required:            true,
									},
								},
							},
						},
						"source_network_objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Objects that represent sources of traffic (Host, Network, Range, FQDN, Network Group, Country, Continent or Geolocation).").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the Network Object.").String,
										Required:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of the Network Object.").String,
										Required:            true,
									},
								},
							},
						},
						"destination_network_literals": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Destination Network literals.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of the Network Object.").AddStringEnumDescription("Host", "Network").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("Host", "Network"),
										},
									},
									"value": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("IP address or network in CIDR format.").String,
										Required:            true,
									},
								},
							},
						},
						"destination_network_objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Objects that represent destinations of traffic (Host, Network, Range, FQDN, Network Group, Country, Continent or Geolocation).").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the Network Object.").String,
										Required:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of the Network Object.").String,
										Required:            true,
									},
								},
							},
						},
						"vlan_tag_literals": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("VLAN Tag literals.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of the VLAN tag literal.").AddStringEnumDescription("VlanTagLiteral").AddDefaultValueDescription("VlanTagLiteral").String,
										Optional:            true,
										Computed:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("VlanTagLiteral"),
										},
										Default: stringdefault.StaticString("VlanTagLiteral"),
									},
									"start_tag": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Start of the VLAN tag range.").String,
										Required:            true,
									},
									"end_tag": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("End of the VLAN tag range.").String,
										Required:            true,
									},
								},
							},
						},
						"vlan_tag_objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("VLAN Tag objects.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the VLAN Tag object.").String,
										Required:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of the VLAN Tag object.").AddStringEnumDescription("VlanTag", "VlanGroupTag").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("VlanTag", "VlanGroupTag"),
										},
									},
								},
							},
						},
						"source_port_literals": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source Port literals.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of the port literal.").AddStringEnumDescription("PortLiteral").AddDefaultValueDescription("PortLiteral").String,
										Optional:            true,
										Computed:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("PortLiteral"),
										},
										Default: stringdefault.StaticString("PortLiteral"),
									},
									"protocol": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("IANA protocol number.").AddStringEnumDescription("6", "17").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("6", "17"),
										},
									},
									"port": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Port number.").String,
										Optional:            true,
									},
								},
							},
						},
						"source_port_objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source Port objects.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the Port object.").String,
										Required:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of the Port object.").AddStringEnumDescription("ProtocolPortObject").AddDefaultValueDescription("ProtocolPortObject").String,
										Optional:            true,
										Computed:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("ProtocolPortObject"),
										},
										Default: stringdefault.StaticString("ProtocolPortObject"),
									},
								},
							},
						},
						"destination_port_literals": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Destination Port literals.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of the port literal.").AddStringEnumDescription("PortLiteral").AddDefaultValueDescription("PortLiteral").String,
										Optional:            true,
										Computed:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("PortLiteral"),
										},
										Default: stringdefault.StaticString("PortLiteral"),
									},
									"port": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Port number.").String,
										Optional:            true,
									},
									"protocol": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("IANA protocol number.").String,
										Required:            true,
									},
								},
							},
						},
						"destination_port_objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Destination Port objects.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the Port object.").String,
										Required:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of the Port object.").AddStringEnumDescription("ProtocolPortObject").AddDefaultValueDescription("ProtocolPortObject").String,
										Optional:            true,
										Computed:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("ProtocolPortObject"),
										},
										Default: stringdefault.StaticString("ProtocolPortObject"),
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *IdentityPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

func (r *IdentityPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	// Check if FMC client is connected to supports this object
	if r.client.FMCVersionParsed.LessThan(minFMCVersionIdentityPolicy) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support Identity Policy creation, minumum required version is 10.0", r.client.FMCVersion))
		return
	}
	var plan IdentityPolicy

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

	planBody := plan.toBody(ctx, IdentityPolicy{})

	// Create object
	body := planBody
	body, _ = sjson.Delete(body, "dummy_categories")
	body, _ = sjson.Delete(body, "dummy_rules")

	res, err := r.client.Post(plan.getPath(), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("id").String())
	plan.fromBodyUnknowns(ctx, res)

	// FMCBUG: Identity policy configuration is not applied on inital POST request
	if !plan.ActiveAuthenticationServerCertificateId.IsNull() || !plan.ActiveAuthenticationRedirectFqdnId.IsNull() || !plan.IdentityMappingFilterNetworkObjectId.IsNull() {
		updateBody, _ := sjson.Set(body, "id", plan.Id.ValueString())
		updateRes, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), updateBody, reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, updateRes.String()))

			_, err := r.client.Delete(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), reqMods...)
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("%s: Failed to delete object after unsuccessful creation, got error: %s", plan.Id.ValueString(), err))
			}
			return
		}
	}

	state := plan
	state.Rules = nil
	state.Categories = nil

	state, diags = r.updateSubresources(ctx, req.Plan, plan, planBody, tfsdk.State{}, state)
	resp.Diagnostics.Append(diags...)

	// On error we do Set anyway. Terraform taints our resource, and the next run is responsible to call Delete() for us.
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished", state.Id.ValueString()))

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *IdentityPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Check if FMC client is connected to supports this object
	if r.client.FMCVersionParsed.LessThan(minFMCVersionIdentityPolicy) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support Identity Policy, minimum required version is 10.0", r.client.FMCVersion))
		return
	}
	var state IdentityPolicy

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
	resGet, err := r.client.Get(urlPath, reqMods...)

	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, resGet.String()))
		return
	}

	resCategories, err := r.client.Get(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString())+"/identitycategories?expanded=true", reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve categories object (GET), got error: %s, %s", err, resCategories.String()))
		return
	}

	resRules, err := r.client.Get(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString())+"/identityrules?expanded=true", reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve rules object (GET), got error: %s, %s", err, resRules.String()))
		return
	}

	// API does not return rules in order they exist in policy, so we need to re-order that
	rulesItems := resRules.Get("items").Array()
	sort.Slice(rulesItems, func(i, j int) bool {
		return rulesItems[i].Get("metadata.index").Int() < rulesItems[j].Get("metadata.index").Int()
	})
	var sb strings.Builder
	sb.WriteString("[")
	for i, rule := range rulesItems {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(rule.Raw)
	}
	sb.WriteString("]")
	sortedRulesJSON := sb.String()

	s := resGet.String()

	replaceCategories := resCategories.Get("items").String()

	// Remove system categories (two fist and the last one)
	tmp := resCategories.Get("items").Array()
	if len(tmp) == 3 {
		replaceCategories = "[]"
	} else {
		tmp2 := tmp[2 : len(tmp)-1]
		replaceCategories = gjson.Parse(fmt.Sprintf(`{"items":%s}`, tmp2)).Get("items").String()
	}

	s, _ = sjson.SetRaw(s, "dummy_categories", replaceCategories)

	replaceRules := sortedRulesJSON
	if replaceRules == "" || replaceRules == "[]" {
		replaceRules = "[]"
	}
	s, _ = sjson.SetRaw(s, "dummy_rules", replaceRules)

	res := gjson.Parse(s)

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *IdentityPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IdentityPolicy

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

	planBody := plan.toBody(ctx, state)
	body := planBody
	body, _ = sjson.Delete(body, "dummy_rules")
	body, _ = sjson.Delete(body, "dummy_categories")

	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.fromBodyUnknowns(ctx, res)

	orig := state
	state = plan
	state.Rules = orig.Rules
	state.Categories = orig.Categories

	state, diags = r.updateSubresources(ctx, req.Plan, plan, planBody, req.State, state)
	resp.Diagnostics.Append(diags...)

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *IdentityPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IdentityPolicy

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
func (r *IdentityPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

// updateSubresources returns a coherent state whether it fails or succeeds. Caller should always set that state
// into the Response (UpdateResponse, CreateResponse, ...), otherwise the API's UUIDs may go out-of-sync with
// terraform.tfstate, which is always a big user-facing problem.
func (r *IdentityPolicyResource) updateSubresources(ctx context.Context, tfsdkPlan tfsdk.Plan, plan IdentityPolicy, planBody string, tfsdkState tfsdk.State, state IdentityPolicy) (IdentityPolicy, diag.Diagnostics) {
	var diags diag.Diagnostics
	var err error

	p := gjson.Parse(planBody)

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !plan.Domain.IsNull() && plan.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(plan.Domain.ValueString()))
	}

	keptCats, keptRules := r.countKept(ctx, state, plan)

	err = r.truncateRulesAt(ctx, &state, keptRules, reqMods...)
	if err != nil {
		diags.AddError("Client Error", err.Error())
		return state, diags
	}

	err = r.truncateCatsAt(ctx, &state, keptCats, reqMods...)
	if err != nil {
		diags.AddError("Client Error", err.Error())
		return state, diags
	}

	if len(plan.Categories) == 0 {
		state.Categories = plan.Categories
	}

	if len(plan.Rules) == 0 {
		state.Rules = plan.Rules
	}

	// Retrieve system categories
	res, err := r.client.Get(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString())+"/identitycategories", reqMods...)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Failed to retrieve system categories, got error: %v, %s", err, res))
		return state, diags
	}

	systemCategories := make(map[string]string)
	items := res.Get("items").Array()
	if len(items) != 3 {
		diags.AddError("Client Error", fmt.Sprintf("Unexpected number of system categories, expected 3 but got %d, response: %s", len(items), res))
		return state, diags
	}

	systemCategories[items[0].Get("name").String()] = items[0].Get("id").String()
	systemCategories[items[1].Get("name").String()] = items[1].Get("id").String()
	systemCategories[items[2].Get("name").String()] = items[2].Get("id").String()

	bodyCats := p.Get("dummy_categories").Array()
	err = r.createCatsAt(ctx, plan, bodyCats, keptCats, &state, reqMods...)
	if err != nil {
		diags.AddError("Client Error", err.Error())
		return state, diags
	}

	bodyRules := p.Get("dummy_rules").Array()
	err = r.createRulesAt(ctx, plan, bodyRules, keptRules, &state, systemCategories, reqMods...)
	if err != nil {
		diags.AddError("Client Error", err.Error())
		return state, diags
	}

	return state, diags
}

// countKept compares the state with the plan starting from index 0, and returns:
// how many categories to keep: they remain identical as to content and order
// how many rules to keep:
// - kept rules must belong to some category that is itself kept,
// - and must themselves remain identical as to id, content, and order
func (r *IdentityPolicyResource) countKept(ctx context.Context, state, plan IdentityPolicy) (int, int) {
	return 0, 0 // TODO
}

func (r *IdentityPolicyResource) truncateRulesAt(ctx context.Context, state *IdentityPolicy, kept int, reqMods ...func(*fmc.Req)) error {
	baseUrl := state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()) + "/identityrules/"

	for i := len(state.Rules) - 1; i >= kept; i-- {
		res, err := r.client.Delete(baseUrl+url.QueryEscape(state.Rules[i].Id.ValueString()), reqMods...)
		if err != nil {
			return fmt.Errorf("Failed to delete rule, got error: %v, %s", err, res.String())
		}

		state.Rules[i] = IdentityPolicyRules{}
		state.Rules = state.Rules[:i]
	}

	return nil
}

func (r *IdentityPolicyResource) truncateCatsAt(ctx context.Context, state *IdentityPolicy, kept int, reqMods ...func(*fmc.Req)) error {
	baseUrl := state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()) + "/identitycategories/"

	for i := len(state.Categories) - 1; i >= kept; i-- {
		res, err := r.client.Delete(baseUrl+url.QueryEscape(state.Categories[i].Id.ValueString()), reqMods...)
		if err != nil {
			return fmt.Errorf("Failed to delete a category, got error: %v, %s", err, res)
		}

		state.Categories[i] = IdentityPolicyCategories{}
		state.Categories = state.Categories[:i]
	}
	return nil
}

func (r *IdentityPolicyResource) createCatsAt(ctx context.Context, plan IdentityPolicy, body []gjson.Result, startIndex int, state *IdentityPolicy, reqMods ...func(*fmc.Req)) error {
	var baseUrl string = ""

	if startIndex < len(plan.Categories) {
		res, err := r.client.Get(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString())+"/identitycategories?filter=name:root_category", reqMods...)
		if err != nil {
			return fmt.Errorf("Failed to retrieve Root Rules category, got error: %v, %s", err, res)
		}
		if catId := res.Get("items.0.id").String(); catId != "" {
			baseUrl = plan.getPath() + "/" + url.QueryEscape(plan.Id.ValueString()) + "/identitycategories?aboveCategory=" + url.QueryEscape(catId)
		} else {
			return fmt.Errorf("Failed to find Root Rules category, got response: %s", res)
		}
	}

	for i := startIndex; i < len(plan.Categories); i++ {
		cat := body[i].String()
		cat, _ = sjson.Delete(cat, "id")

		res, err := r.client.Post(baseUrl, cat, reqMods...)
		if err != nil {
			return fmt.Errorf("Failed to create a category (POST), got error: %v, %s", err, res)
		}

		item := plan.Categories[i]
		item.Id = types.StringValue(res.Get("id").String())

		if len(state.Categories) <= i {
			state.Categories = append(state.Categories, item)
		} else {
			state.Categories[i] = item
		}
	}

	return nil
}

func (r *IdentityPolicyResource) createRulesAt(ctx context.Context, plan IdentityPolicy, body []gjson.Result, startIndex int, state *IdentityPolicy, systemCategories map[string]string, reqMods ...func(*fmc.Req)) error {
	baseUrl := plan.getPath() + "/" + url.QueryEscape(plan.Id.ValueString()) + "/identityrules"
	mapCatNameToId := make(map[string]string)
	for _, cat := range state.Categories {
		mapCatNameToId[cat.Name.ValueString()] = url.QueryEscape(cat.Id.ValueString())
	}
	for cat := range systemCategories {
		mapCatNameToId[identityPolicyCategoriesApiToGuiNames[cat]] = url.QueryEscape(systemCategories[cat])
	}

	for i := startIndex; i < len(plan.Rules); i++ {
		cat := body[i].String()
		cat, _ = sjson.Delete(cat, "id")

		res, err := r.client.Post(baseUrl+"?intoCategory="+mapCatNameToId[plan.Rules[i].Category.ValueString()], cat, reqMods...)
		if err != nil {
			return fmt.Errorf("Failed to create a rule (POST), got error: %v, %s", err, res)
		}

		item := plan.Rules[i]
		item.Id = types.StringValue(res.Get("id").String())

		if len(state.Rules) <= i {
			state.Rules = append(state.Rules, item)
		} else {
			state.Rules[i] = item
		}
	}

	return nil
}
