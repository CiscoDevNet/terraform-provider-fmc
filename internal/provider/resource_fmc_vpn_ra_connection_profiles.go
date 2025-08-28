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
	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/planmodifiers"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
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
	_ resource.Resource                = &VPNRAConnectionProfilesResource{}
	_ resource.ResourceWithImportState = &VPNRAConnectionProfilesResource{}
)

func NewVPNRAConnectionProfilesResource() resource.Resource {
	return &VPNRAConnectionProfilesResource{}
}

type VPNRAConnectionProfilesResource struct {
	client *fmc.Client
}

func (r *VPNRAConnectionProfilesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_ra_connection_profiles"
}

func (r *VPNRAConnectionProfilesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages FTD Remote Access (RA) Virtual Private Networks (VPNs) Connection Profiles.").AddMinimumVersionHeaderDescription().AddMinimumVersionBulkCreateDescription("999").AddMinimumVersionBulkDeleteDescription("999").AddMinimumVersionBulkUpdateDescription().String,

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
			"vpn_ra_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the parent VPN RA Configuration.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"items": schema.MapNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Map of Connection Profiles. The key of the map is the name of the Connection Profile. Use `DefaultWEBVPNGroup` to manage the default connection profile. On destruction, the default connection profile will not be deleted and its configuration will not be erased.").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the Connection Profile.").String,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								planmodifiers.UseStateForUnknownKeepNonNullStateString(),
							},
						},
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this value is always 'RaVpnConnectionProfile'.").String,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								planmodifiers.UseStateForUnknownKeepNonNullStateString(),
							},
						},
						"group_policy_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the Group Policy.").String,
							Optional:            true,
						},
						"ipv4_address_pools": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv4 Address Pools for client address assignment. Either `ipv4_address_pools` or `ipv6_address_pools` or `dhcp_servers` must be specified.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the IPv4 Address Pool.").String,
										Required:            true,
									},
								},
							},
						},
						"ipv6_address_pools": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv6 Address Pools for client address assignment. Either `ipv4_address_pools` or `ipv6_address_pools` or `dhcp_servers` must be specified.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the IPv6 Address Pool.").String,
										Required:            true,
									},
								},
							},
						},
						"dhcp_servers": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of DHCP Servers for client address assignment. Either `ipv4_address_pools` or `ipv6_address_pools` or `dhcp_servers` must be specified.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of Host representing the DHCP Server.").String,
										Required:            true,
									},
								},
							},
						},
						"authentication_method": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("User authentication method.").AddStringEnumDescription("AAA_ONLY", "SAML", "CLIENT_CERTIFICATE_ONLY", "AAA_AND_CLIENT_CERTIFICATE", "SAML_AND_CLIENT_CERTIFICATE").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("AAA_ONLY", "SAML", "CLIENT_CERTIFICATE_ONLY", "AAA_AND_CLIENT_CERTIFICATE", "SAML_AND_CLIENT_CERTIFICATE"),
							},
						},
						"multiple_certificate_authentication": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Authenticate the VPN client using the machine and user certificates.").String,
							Optional:            true,
						},
						"primary_authentication_server_use_local": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Use LOCAL FMC as primary authentication server.").String,
							Optional:            true,
						},
						"primary_authentication_server_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the primary authentication RADIUS Server Group or Realm. Use if `primary_authentication_server_use_local` is not set to `true`.").String,
							Optional:            true,
						},
						"primary_authentication_server_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of the primary authentication RADIUS Server Group or Realm, like `RadiusServerGroup` or `Realm`.").String,
							Optional:            true,
						},
						"primary_authentication_fallback_to_local": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Fallback to LOCAL FMC if primary authentication Server/Realm is not reachable.").String,
							Optional:            true,
						},
						"saml_and_certificate_username_must_match": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Allow VPN connection only if the username from the certificate matches the SAML single sign-on username.").String,
							Optional:            true,
						},
						"primary_authentication_prefill_username_from_certificate": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Prefill username from certificate.").String,
							Optional:            true,
						},
						"primary_authentication_prefill_username_from_certificate_map_primary_field": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Map primary field for username.").AddStringEnumDescription("CN_COMMMON_NAME", "C_COUNTRY", "DNQ_DN_QUALIFIER", "EA_EMAIL_ADDRESS", "GENQ_GENERATIONAL_QUALIFIER", "GN_GIVEN_NAME", "I_INITIAL", "L_LOCALITY", "N_NAME", "O_ORGANISATION", "OU_ORGANISATIONAL_UNIT", "SER_SERIAL_NUMBER", "SN_SURNAME", "SP_STATE_PROVINCE", "T_TITLE", "UID_USER_ID", "UPN_USER_PRINCIPAL_NAME", "NONE").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("CN_COMMMON_NAME", "C_COUNTRY", "DNQ_DN_QUALIFIER", "EA_EMAIL_ADDRESS", "GENQ_GENERATIONAL_QUALIFIER", "GN_GIVEN_NAME", "I_INITIAL", "L_LOCALITY", "N_NAME", "O_ORGANISATION", "OU_ORGANISATIONAL_UNIT", "SER_SERIAL_NUMBER", "SN_SURNAME", "SP_STATE_PROVINCE", "T_TITLE", "UID_USER_ID", "UPN_USER_PRINCIPAL_NAME", "NONE"),
							},
						},
						"primary_authentication_prefill_username_from_certificate_map_secondary_field": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Map secondary field for username.").AddStringEnumDescription("CN_COMMMON_NAME", "C_COUNTRY", "DNQ_DN_QUALIFIER", "EA_EMAIL_ADDRESS", "GENQ_GENERATIONAL_QUALIFIER", "GN_GIVEN_NAME", "I_INITIAL", "L_LOCALITY", "N_NAME", "O_ORGANISATION", "OU_ORGANISATIONAL_UNIT", "SER_SERIAL_NUMBER", "SN_SURNAME", "SP_STATE_PROVINCE", "T_TITLE", "UID_USER_ID", "UPN_USER_PRINCIPAL_NAME", "NONE").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("CN_COMMMON_NAME", "C_COUNTRY", "DNQ_DN_QUALIFIER", "EA_EMAIL_ADDRESS", "GENQ_GENERATIONAL_QUALIFIER", "GN_GIVEN_NAME", "I_INITIAL", "L_LOCALITY", "N_NAME", "O_ORGANISATION", "OU_ORGANISATIONAL_UNIT", "SER_SERIAL_NUMBER", "SN_SURNAME", "SP_STATE_PROVINCE", "T_TITLE", "UID_USER_ID", "UPN_USER_PRINCIPAL_NAME", "NONE"),
							},
						},
						"primary_authentication_prefill_username_from_certificate_map_entire_dn": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Map entire Distinguished Name (DN) as username.").String,
							Optional:            true,
						},
						"primary_authentication_hide_username_in_login_window": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Username is pre-filled from the client certificate, but hidden to the user.").String,
							Optional:            true,
						},
						"secondary_authentication": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable secondary authentication.").String,
							Optional:            true,
						},
						"secondary_authentication_server_use_local": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Use LOCAL FMC as secondary authentication server.").String,
							Optional:            true,
						},
						"secondary_authentication_server_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the secondary authentication RADIUS Server Group or Realm. Use if `secondary_authentication_server_use_local` is not set to `true`.").String,
							Optional:            true,
						},
						"secondary_authentication_server_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of the secondary authentication RADIUS Server Group or Realm, like `RadiusServerGroup` or `Realm`.").String,
							Optional:            true,
						},
						"secondary_authentication_fallback_to_local": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Fallback to LOCAL FMC if secondary authentication Server/Realm is not reachable.").String,
							Optional:            true,
						},
						"secondary_authentication_prompt_for_username": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Prompt for username during secondary authentication.").String,
							Optional:            true,
						},
						"secondary_authentication_use_primary_authentication_username": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Use primary username for secondary authentication.").String,
							Optional:            true,
						},
						"use_secondary_authentication_username_for_reporting": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Secondary username is used for reporting user activity during a VPN session.").String,
							Optional:            true,
						},
						"saml_use_external_browser": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Use default OS browser (`true`) or embedded browser (`false`) for SAML authentication.").String,
							Optional:            true,
						},
						"authorization_server_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the authorization RADIUS Server Group or Realm.").String,
							Optional:            true,
						},
						"authorization_server_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of the authorization RADIUS Server Group or Realm, like `RadiusServerGroup` or `Realm`.").String,
							Optional:            true,
						},
						"allow_connection_only_if_user_exists_in_authorization_database": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Allow connection only if the username of client exists in the authorization database.").String,
							Optional:            true,
						},
						"accounting_server_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the RADIUS accounting server.").String,
							Optional:            true,
						},
						"accounting_server_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of the RADIUS accounting server (`RadiusServerGroup`).").String,
							Optional:            true,
						},
						"strip_realm_from_username": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Remove the realm from the username before passing the username on to the AAA server.").String,
							Optional:            true,
						},
						"strip_group_from_username": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Remove the group name from the username before passing the username on to the AAA server.").String,
							Optional:            true,
						},
						"password_management": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable managing the password for the remote access VPN users.").String,
							Optional:            true,
						},
						"password_management_notify_user_on_password_expiry_day": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Notify user on the day the password expires.").String,
							Optional:            true,
						},
						"password_management_advance_password_expiration_notification": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Notify user on the number of days before password expiration.").AddIntegerRangeDescription(1, 180).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 180),
							},
						},
						"alias_names": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of Alias Names.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Name of the Alias.").String,
										Required:            true,
									},
									"enabled": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Enable the alias.").String,
										Optional:            true,
									},
								},
							},
						},
						"alias_urls": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of Alias URLs (group URLs).").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"url_object_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the URL object.").String,
										Required:            true,
									},
									"enabled": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Enable the alias.").String,
										Optional:            true,
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

func (r *VPNRAConnectionProfilesResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

var _ resource.ResourceWithValidateConfig = &VPNRAConnectionProfilesResource{}

func (p *VPNRAConnectionProfilesResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data VPNRAConnectionProfiles

	diags := req.Config.Get(ctx, &data)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	for key, item := range data.Items {
		if !item.PrimaryAuthenticationPrefillUsernameFromCertificate.IsNull() && item.PrimaryAuthenticationPrefillUsernameFromCertificateMapEntireDn.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("items").AtMapKey(key).AtName("primary_authentication_prefill_username_from_certificate_map_entire_dn"),
				"Missing Required Attribute",
				fmt.Sprintf("%s: The `primary_authentication_prefill_username_from_certificate_map_entire_dn` attribute is required when `primary_authentication_prefill_username_from_certificate` is set.", path.Root("items").AtMapKey(key)),
			)
		}

		// if !item.SecondaryAuthenticationPrefillUsernameFromCertificate.IsNull() && item.SecondaryAuthenticationPrefillUsernameFromCertificateMapEntireDn.IsNull() {
		// 	resp.Diagnostics.AddAttributeError(
		// 		path.Root("items").AtMapKey(key).AtName("secondary_authentication_prefill_username_from_certificate_map_entire_dn"),
		// 		"Missing Required Attribute",
		// 		fmt.Sprintf("%s: The `secondary_authentication_prefill_username_from_certificate_map_entire_dn` attribute is required when `secondary_authentication_prefill_username_from_certificate` is set.", path.Root("items").AtMapKey(key)),
		// 	)
		// }

		if !item.SecondaryAuthentication.IsNull() && item.SecondaryAuthentication.ValueBool() {

			if item.SecondaryAuthenticationPromptForUsername.IsNull() {
				resp.Diagnostics.AddAttributeError(
					path.Root("items").AtMapKey(key).AtName("secondary_authentication_prompt_for_username"),
					"Missing Required Attribute",
					fmt.Sprintf("%s: The `secondary_authentication_prompt_for_username` attribute is required when `secondary_authentication` is set.", path.Root("items").AtMapKey(key)),
				)
			}

			if item.SecondaryAuthenticationUsePrimaryAuthenticationUsername.IsNull() {
				resp.Diagnostics.AddAttributeError(
					path.Root("items").AtMapKey(key).AtName("secondary_authentication_use_primary_authentication_username"),
					"Missing Required Attribute",
					fmt.Sprintf("%s: The `secondary_authentication_use_primary_authentication_username` attribute is required when `secondary_authentication` is set.", path.Root("items").AtMapKey(key)),
				)
			}

			if item.UseSecondaryAuthenticationUsernameForReporting.IsNull() {
				resp.Diagnostics.AddAttributeError(
					path.Root("items").AtMapKey(key).AtName("use_secondary_authentication_username_for_reporting"),
					"Missing Required Attribute",
					fmt.Sprintf("%s: The `use_secondary_authentication_username_for_reporting` attribute is required when `secondary_authentication` is set.", path.Root("items").AtMapKey(key)),
				)
			}

		}
	}
}

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *VPNRAConnectionProfilesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan VPNRAConnectionProfiles

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

	//// Prepare state to track creation process. Create request is split to multiple requests, where just subset of them may be successful
	// Copy fields, as those may contain domain information or other references
	state := plan
	// Create random ID to track bulk resource. This does not relate to FMC in any way
	state.Id = types.StringValue(uuid.New().String())
	// Erase all Items, those will be filled in after creation
	state.Items = make(map[string]VPNRAConnectionProfilesItems, len(plan.Items))
	// Creation process is put in a separate function, as that same proces will be needed with `Update`
	plan, diags = r.createSubresources(ctx, state, plan, reqMods...)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		// Save state for whatever was already created
		diags = resp.State.Set(ctx, &plan)
		tflog.Debug(ctx, fmt.Sprintf("%s: Create failed, some items might have been created", plan.Id.ValueString()))
		resp.Diagnostics.Append(diags...)
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *VPNRAConnectionProfilesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state VPNRAConnectionProfiles

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

	// Get all objects from FMC
	urlPath := state.getPath() + "?expanded=true"
	res, err := r.client.Get(urlPath, reqMods...)
	if err != nil {
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

func (r *VPNRAConnectionProfilesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state VPNRAConnectionProfiles

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

	// DELETE
	// Delete objects (that are present in state, but missing in plan)
	var toDelete VPNRAConnectionProfiles
	toDelete.Items = make(map[string]VPNRAConnectionProfilesItems)
	planOwnedIDs := make(map[string]string, len(plan.Items))

	// Prepare list of ID that are in plan
	for k, v := range plan.Items {
		if !v.Id.IsUnknown() && v.Id.ValueString() != "" {
			planOwnedIDs[v.Id.ValueString()] = k
		}
	}

	// Check if ID from state list is in plan as well. If not, mark it for delete
	for k, v := range state.Items {
		if _, ok := planOwnedIDs[v.Id.ValueString()]; !ok {
			toDelete.Items[k] = v
		}
	}

	// If there are objects marked to be deleted
	if len(toDelete.Items) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to delete: %d", state.Id.ValueString(), len(toDelete.Items)))
		state, diags = r.deleteSubresources(ctx, state, toDelete, reqMods...)
		if diags != nil {
			resp.Diagnostics.Append(diags...)
			diags = resp.State.Set(ctx, &state)
			resp.Diagnostics.Append(diags...)
			return
		}
	}

	// CREATE
	// Create new objects (objects that have missing IDs in plan)
	var toCreate VPNRAConnectionProfiles
	toCreate.Items = make(map[string]VPNRAConnectionProfilesItems)
	// Scan plan for items with no ID
	for k, v := range plan.Items {
		if v.Id.IsUnknown() || v.Id.IsNull() {
			toCreate.Items[k] = v
		}
	}

	// If there are objects marked for create
	if len(toCreate.Items) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to create: %d", state.Id.ValueString(), len(toCreate.Items)))
		state, diags = r.createSubresources(ctx, state, toCreate, reqMods...)
		if diags != nil {
			resp.Diagnostics.Append(diags...)
			diags = resp.State.Set(ctx, &state)
			resp.Diagnostics.Append(diags...)
			return
		}
	}

	// UPDATE
	// Update objects (objects that have different definition in plan and state)
	var notEqual bool
	var toUpdate VPNRAConnectionProfiles
	toUpdate.Items = make(map[string]VPNRAConnectionProfilesItems)

	for _, valueState := range state.Items {

		// Check if the ID from plan exists on list of ID owned by state
		if keyState, ok := planOwnedIDs[valueState.Id.ValueString()]; ok {

			// Check if items in state and plan are qual
			notEqual, diags = helpers.IsConfigUpdatingAt(ctx, req.Plan, req.State, path.Root("items").AtMapKey(keyState))
			if diags != nil {
				resp.Diagnostics.Append(diags...)
				diags = resp.State.Set(ctx, &state)
				resp.Diagnostics.Append(diags...)
				return
			}

			// If definitions differ, add object to update list
			if notEqual {
				toUpdate.Items[keyState] = plan.Items[keyState]
			}
		}
	}

	// If there are objects marked for update
	if len(toUpdate.Items) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to update: %d", state.Id.ValueString(), len(toUpdate.Items)))
		state, diags = r.updateSubresources(ctx, state, toUpdate, reqMods...)
		if diags != nil {
			resp.Diagnostics.Append(diags...)
			diags = resp.State.Set(ctx, &state)
			resp.Diagnostics.Append(diags...)
			return
		}
	}
	plan = state

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *VPNRAConnectionProfilesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state VPNRAConnectionProfiles

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

	// Execute delete
	state, diags = r.deleteSubresources(ctx, state, state, reqMods...)
	resp.Diagnostics.Append(diags...)

	// Check if every element was removed
	if len(state.Items) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Not all elements have been removed", state.Id.ValueString()))
		diags = resp.State.Set(ctx, &state)
		resp.Diagnostics.Append(diags...)
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *VPNRAConnectionProfilesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import looks for string in the following format: <domain_name>,<ref_id>,[<object1_name>,<object2_name>,...]
	// <domain_name> is optional
	// <ref_id> for objects that have `reference` attributes
	// <object1_name>,<object2_name>,... is coma-separated list of object names
	var config VPNRAConnectionProfiles

	// Compile pattern for import command parsing
	var inputPattern = regexp.MustCompile(`^(?:(?P<domain>[^\s,]+),)?(?P<vpn_ra_id>[^\s,]+),\[(?P<names>.*?)\]$`)

	// Parse parameter
	match := inputPattern.FindStringSubmatch(req.ID)

	// Check if regex matched
	if match == nil {
		resp.Diagnostics.AddError("Import error", "Failed to parse import parameters. Please provide import string in the following format: <domain_name>,[<object1_name>,<object2_name>,...]")
		return
	}

	// Extract values
	if tmpDomain := match[inputPattern.SubexpIndex("domain")]; tmpDomain != "" {
		config.Domain = types.StringValue(tmpDomain)
	}
	names := strings.Split(match[inputPattern.SubexpIndex("names")], ",")

	// Fill state with names of objects to import
	config.Items = make(map[string]VPNRAConnectionProfilesItems, len(names))
	for _, v := range names {
		config.Items[v] = VPNRAConnectionProfilesItems{}
	}

	// Set reference attributes
	config.VpnRaId = types.StringValue(match[inputPattern.SubexpIndex("vpn_ra_id")])

	// Generate new ID
	config.Id = types.StringValue(uuid.New().String())

	// Set filled in structure
	diags := resp.State.Set(ctx, config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set import flag
	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import

// createSubresources takes list of objects, splits them into bulks and creates them
// We want to save the state after each create event, to be able track already created resources
func (r *VPNRAConnectionProfilesResource) createSubresources(ctx context.Context, state, plan VPNRAConnectionProfiles, reqMods ...func(*fmc.Req)) (VPNRAConnectionProfiles, diag.Diagnostics) {

	var tmpObject VPNRAConnectionProfiles
	tmpObject.Items = make(map[string]VPNRAConnectionProfilesItems, 1)
	var res gjson.Result
	for k, v := range plan.Items {

		if k == "DefaultWEBVPNGroup" {
			// This is a default connection profile, we will update it instead of creating a new one
			tflog.Debug(ctx, fmt.Sprintf("%s: Default Connection Profile found, updating instead of creating", state.Id.ValueString()))
			// Get id of the default connection profile
			resTmp, err := r.client.Get(state.getPath(), reqMods...)
			if err != nil {
				return state, diag.Diagnostics{
					diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("%s: Failed to retrieve default object (GET), got error: %s, %s", state.Id.ValueString(), err, resTmp.String())),
				}
			}
			query := "items.#(name==\"DefaultWEBVPNGroup\").id"
			if defaultWebVPNGroupId := resTmp.Get(query); defaultWebVPNGroupId.Exists() {
				// Set the ID of the default connection profile
				v.Id = types.StringValue(defaultWebVPNGroupId.String())
			} else {
				return state, diag.Diagnostics{
					diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("%s: Failed extract DefaultWEBVPNGroup object ID: %s", state.Id.ValueString(), resTmp.String())),
				}
			}

			// Update object
			tmpObject.Items[k] = v
			body := tmpObject.toBodyNonBulk(ctx, state)
			res, err = r.client.Put(state.getPath()+"/"+url.QueryEscape(v.Id.ValueString()), body, reqMods...)
			if err != nil {
				return state, diag.Diagnostics{
					diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("%s: Failed to update object (PUT) id %s, got error: %s, %s", state.Id.ValueString(), v.Id.ValueString(), err, res.String())),
				}
			}

		} else {
			var err error
			tmpObject.Items[k] = v
			body := tmpObject.toBodyNonBulk(ctx, state)
			res, err = r.client.Post(state.getPath(), body, reqMods...)
			if err != nil {
				return state, diag.Diagnostics{
					diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("%s: Failed to create object (POST) id %s, got error: %s, %s", state.Id.ValueString(), v.Id.ValueString(), err, res.String())),
				}
			}
		}

		// fromBodyUnknowns expect result to be listed under "items" key
		body, _ := sjson.SetRaw("{items:[]}", "items.-1", res.String())
		res = gjson.Parse(body)

		// Read computed values
		tmpObject.fromBodyUnknowns(ctx, res)

		// Save object to plan
		state.Items[k] = tmpObject.Items[k]

		// Clear tmpObject.Items
		delete(tmpObject.Items, k)
	}

	return state, nil
}

// deleteSubresources takes list of objects and deletes them either in bulk, or one-by-one, depending on FMC version
func (r *VPNRAConnectionProfilesResource) deleteSubresources(ctx context.Context, state, plan VPNRAConnectionProfiles, reqMods ...func(*fmc.Req)) (VPNRAConnectionProfiles, diag.Diagnostics) {
	objectsToRemove := plan.Clone()

	tflog.Debug(ctx, fmt.Sprintf("%s: One-by-one deletion mode (VPN RA Connection Profiles)", state.Id.ValueString()))
	for k, v := range objectsToRemove.Items {
		// Check if the object was not already deleted
		if v.Id.IsNull() {
			delete(state.Items, k)
			continue
		}

		// If this is default connection profile, we will not delete it
		if k == "DefaultWEBVPNGroup" {
			delete(state.Items, k)
			continue
		}

		urlPath := state.getPath() + "/" + url.QueryEscape(v.Id.ValueString())
		res, err := r.client.Delete(urlPath, reqMods...)
		if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
			return state, diag.Diagnostics{
				diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("%s: Failed to delete object (DELETE) id %s, got error: %s, %s", state.Id.ValueString(), v.Id.ValueString(), err, res.String())),
			}
		}

		// Remove deleted item from state
		delete(state.Items, k)
	}

	return state, nil
}

// Section below is generated&owned by "gen/generator.go". //template:begin updateSubresources

// updateSubresources take elements one-by-one and updates them, as bulks are not supported
func (r *VPNRAConnectionProfilesResource) updateSubresources(ctx context.Context, state, plan VPNRAConnectionProfiles, reqMods ...func(*fmc.Req)) (VPNRAConnectionProfiles, diag.Diagnostics) {
	var tmpObject VPNRAConnectionProfiles
	tmpObject.Items = make(map[string]VPNRAConnectionProfilesItems, 1)

	tflog.Debug(ctx, fmt.Sprintf("%s: One-by-one update mode (VPN RA Connection Profiles)", state.Id.ValueString()))

	for k, v := range plan.Items {
		tmpObject.Items[k] = v

		body := tmpObject.toBodyNonBulk(ctx, state)
		urlPath := state.getPath() + "/" + url.QueryEscape(v.Id.ValueString())
		res, err := r.client.Put(urlPath, body, reqMods...)
		if err != nil {
			return state, diag.Diagnostics{
				diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Failed to update object (PUT) id %s, got error: %s, %s", state.Id.ValueString(), err, res.String())),
			}
		}

		// Update state
		state.Items[k] = v

		// Clear tmpObject.Items
		delete(tmpObject.Items, k)
	}

	return state, nil
}

// End of section. //template:end updateSubresources
