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
	"strings"

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
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
	_ resource.Resource                = &RealmADLDAPResource{}
	_ resource.ResourceWithImportState = &RealmADLDAPResource{}
)

func NewRealmADLDAPResource() resource.Resource {
	return &RealmADLDAPResource{}
}

type RealmADLDAPResource struct {
	client *fmc.Client
}

func (r *RealmADLDAPResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_realm_ad_ldap"
}

func (r *RealmADLDAPResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages a Realm AD LDAP.").String,

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
				MarkdownDescription: helpers.NewAttributeDescription("Name of the Realm object.").String,
				Required:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Realm object.").String,
				Optional:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this value is always 'Realm'.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Internal API parameter.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description of the Realm object.").String,
				Optional:            true,
			},
			"realm_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the Realm.").AddStringEnumDescription("AD", "LDAP").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("AD", "LDAP"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"ad_primary_domain": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Domain for the Active Directory server where users should be authenticated.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"ad_join_username": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Username of any Active Directory user with rights to create a Domain Computer account in the Active Directory domain for Kerberos captive portal active authentication.").String,
				Optional:            true,
			},
			"ad_join_password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Password for `ad_join_username` user.").String,
				Optional:            true,
				Sensitive:           true,
			},
			"directory_username": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Username used to connect to the directory.").String,
				Required:            true,
			},
			"directory_password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Password for the `directory_username`.").String,
				Required:            true,
				Sensitive:           true,
			},
			"base_dn": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Directory tree where the search for user data should begin.").String,
				Required:            true,
			},
			"group_dn": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Directory tree where the search for group data should begin.").String,
				Required:            true,
			},
			"included_users": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Add users to Included Users.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"included_groups": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Add groups to Included Groups.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"excluded_users": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Add users to Excluded Users.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"excluded_groups": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Add groups to Excluded Groups.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"update_hour": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Hour where the sync (download) from the directory starts.").AddIntegerRangeDescription(0, 23).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 23),
				},
			},
			"update_interval": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interval in hours for the sync (download) from the directory.").AddStringEnumDescription("1", "2", "3", "4", "6", "8", "12", "24").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1", "2", "3", "4", "6", "8", "12", "24"),
				},
			},
			"group_attribute": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Attribute used to identify the group in the LDAP directory. Use 'uniqueMember', 'member' or any custom attribute name.").String,
				Optional:            true,
			},
			"timeout_ise_and_passive_indentity_users": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Timeout (in minutes) for ISE/ISE-PIC or Passive Identity Agent users.").AddIntegerRangeDescription(0, 35791394).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 35791394),
				},
			},
			"timeout_terminal_server_agent_users": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Timeout (in minutes) for Terminal Server Agent users.").AddIntegerRangeDescription(0, 35791394).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 35791394),
				},
			},
			"timeout_captive_portal_users": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Timeout (in minutes) for Captive Portal users.").AddIntegerRangeDescription(0, 35791394).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 35791394),
				},
			},
			"timeout_failed_captive_portal_users": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Timeout (in minutes) for Failed Captive Portal users.").AddIntegerRangeDescription(0, 35791394).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 35791394),
				},
			},
			"timeout_guest_captive_portal_users": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Timeout (in minutes) for Guest Captive Portal Users.").AddIntegerRangeDescription(0, 35791394).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 35791394),
				},
			},
			"directory_servers": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of directory servers.").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"hostname": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Hostname or IP address.").String,
							Required:            true,
						},
						"port": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Port number.").String,
							Required:            true,
						},
						"encryption_protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Encryption method.").AddStringEnumDescription("NONE", "LDAPS", "STARTTLS").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("NONE", "LDAPS", "STARTTLS"),
							},
						},
						"ca_certificate_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("CA certificate ID. Required if `encryption_protocol` is LDAPS/STARTTLS.").String,
							Optional:            true,
						},
						"use_routing_to_select_interface": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Use routing to select the interface for directory communication.").AddDefaultValueDescription("false").String,
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"interface_group_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of the interface group to use for LDAP communication, when `use_routing_to_select_interface` is set to `false`. If not configured, Management interface is used.").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *RealmADLDAPResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *RealmADLDAPResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RealmADLDAP

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
	body := plan.toBody(ctx, RealmADLDAP{})
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

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *RealmADLDAPResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RealmADLDAP

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

func (r *RealmADLDAPResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RealmADLDAP

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

func (r *RealmADLDAPResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RealmADLDAP

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

func (r *RealmADLDAPResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin createSubresources

// End of section. //template:end createSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin deleteSubresources

// End of section. //template:end deleteSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin updateSubresources

// End of section. //template:end updateSubresources
