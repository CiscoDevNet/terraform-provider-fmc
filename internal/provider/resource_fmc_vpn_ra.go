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
	"time"

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
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
	_ resource.Resource                = &VPNRAResource{}
	_ resource.ResourceWithImportState = &VPNRAResource{}
)

func NewVPNRAResource() resource.Resource {
	return &VPNRAResource{}
}

type VPNRAResource struct {
	client *fmc.Client
}

func (r *VPNRAResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_ra"
}

func (r *VPNRAResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages FTD Remote Access (RA) Virtual Private Networks (VPNs).").String,

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
				MarkdownDescription: helpers.NewAttributeDescription("Name of the VPN Remote Access (RA) Configuration.").String,
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description of the object.").String,
				Optional:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this value is always 'RAVpn'.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"protocol_ssl": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable SSL protocol.").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"protocol_ipsec_ikev2": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable IPsec IKEv2 protocol.").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"local_realm_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of local realm server. This can be set only after relevant connection profiles are configured.").String,
				Optional:            true,
			},
			"dynamic_access_policy_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of Dynamic Access Policy (DAP).").String,
				Optional:            true,
			},
			"access_interfaces": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of interfaces that will support incoming Remote Access VPN connections.").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of Interface Group or Security Zone.").String,
							Required:            true,
						},
						"protocol_ipsec_ikev2": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable IPsec IKEv2.").String,
							Optional:            true,
						},
						"protocol_ssl": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable SSL.").String,
							Optional:            true,
						},
						"protocol_ssl_dtls": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable DTLS for the VPN.").String,
							Optional:            true,
						},
						"interface_specific_certificate_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of interface specific identity certificate.").String,
							Optional:            true,
						},
					},
				},
			},
			"allow_users_to_select_connection_profile": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allow Users to select connection profile while logging in.").String,
				Optional:            true,
			},
			"web_access_port_number": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Web Access Port Number.").AddDefaultValueDescription("443").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(443),
			},
			"dtls_port_number": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("DTLS Port Number.").AddDefaultValueDescription("443").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(443),
			},
			"ssl_global_identity_certificate_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of SSL Global Identity Certificate.").String,
				Optional:            true,
			},
			"ipsec_ikev2_identity_certificate_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of IPsec IKEv2 Identity Certificate.").String,
				Optional:            true,
			},
			"service_access_object_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the Service Access object.").String,
				Optional:            true,
			},
			"bypass_access_control_policy_for_decrypted_traffic": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Bypass Access Control policy for decrypted traffic (sysopt permit-vpn).").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"secure_client_images": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of Secure Client images.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of Secure Client image.").String,
							Required:            true,
						},
						"operating_system": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Operating system for which the Secure Client image is intended.").AddStringEnumDescription("WINDOWS", "LINUX", "MAC").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("WINDOWS", "LINUX", "MAC"),
							},
						},
					},
				},
			},
			"external_browser_package_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of Secure Client External Browser Package.").String,
				Optional:            true,
			},
			"secure_client_customization_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of Secure Client Customization.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"address_assignment_policy_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of Address Assignment Policy.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"certificate_map_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of Certificate Map.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"group_policies": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of Group Policies associated with the VPN. It is mandatory to include at least 'DfltGrpPolicy' in the list.").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of group policy.").String,
							Required:            true,
						},
					},
				},
			},
			"ldap_attribute_map_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of LDAP Attribute Mapping.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"load_balancing_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of Load Balancing settings.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"ikev2_policies": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of IKEv2 policies.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of IKEv2 policy.").String,
							Required:            true,
						},
					},
				},
			},
			"ipsec_ike_parameters_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of IPsec/IKEv2 parameters.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *VPNRAResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

func (r *VPNRAResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan VPNRA

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
	body := plan.toBody(ctx, VPNRA{})
	res, err := r.client.Post(plan.getPath(), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("id").String())
	plan.fromBodyUnknowns(ctx, res)

	// Wait for the FMC to process the new object and create subresources
	time.Sleep(5 * time.Second)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *VPNRAResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state VPNRA

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update

func (r *VPNRAResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state VPNRA

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

func (r *VPNRAResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state VPNRA

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
func (r *VPNRAResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
