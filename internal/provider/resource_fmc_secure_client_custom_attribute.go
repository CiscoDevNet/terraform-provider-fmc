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
	_ resource.Resource                = &SecureClientCustomAttributeResource{}
	_ resource.ResourceWithImportState = &SecureClientCustomAttributeResource{}
)

func NewSecureClientCustomAttributeResource() resource.Resource {
	return &SecureClientCustomAttributeResource{}
}

type SecureClientCustomAttributeResource struct {
	client *fmc.Client
}

func (r *SecureClientCustomAttributeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_secure_client_custom_attribute"
}

func (r *SecureClientCustomAttributeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages a Secure Client Custom Attribute.").AddMinimumVersionHeaderDescription().AddMinimumVersionAnyDescription().AddMinimumVersionCreateDescription("7.4").String,

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
				MarkdownDescription: helpers.NewAttributeDescription("Name of the Secure Client Custom Attribute object.").String,
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this value is always 'AnyConnectCustomAttribute'.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description of the Secure Client Custom Attribute object.").String,
				Optional:            true,
			},
			"attribute_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the attribute.").AddStringEnumDescription("ALLOW_DEFER_UPDATE", "PER_APP_VPN", "DYNAMIC_SPLIT_TUNNELING", "USER_DEFINED_CUSTOM_ATTR").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("ALLOW_DEFER_UPDATE", "PER_APP_VPN", "DYNAMIC_SPLIT_TUNNELING", "USER_DEFINED_CUSTOM_ATTR"),
				},
			},
			"user_defined_attribute_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the user-defined attribute. Applicable only when `attribute_type` is USER_DEFINED_CUSTOM_ATTR.").String,
				Optional:            true,
			},
			"user_defined_attribute_value": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Value of the user-defined attribute. Applicable only when `attribute_type` is USER_DEFINED_CUSTOM_ATTR.").String,
				Optional:            true,
			},
			"per_app_vpn_value": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Base64 encoded value for Per App VPN. Applicable only when `attribute_type` is PER_APP_VPN.").String,
				Optional:            true,
			},
			"dynamic_split_tunnel_included_domains": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Domain names that will be included in the remote access VPN tunnel. Applicable only when `attribute_type` is DYNAMIC_SPLIT_TUNNELING.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"dynamic_split_tunnel_excluded_domains": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Domain names that will be excluded from the remote access VPN tunnel. Applicable only when `attribute_type` is DYNAMIC_SPLIT_TUNNELING.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"defer_update_prompt_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Prompt type.").AddStringEnumDescription("SHOW_UNTIL_USER_ACTION", "SHOW_UNTIL_TIMEOUT", "NO_PROMPT_AUTO_ACTION").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("SHOW_UNTIL_USER_ACTION", "SHOW_UNTIL_TIMEOUT", "NO_PROMPT_AUTO_ACTION"),
				},
			},
			"defer_update_default_action": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Default action to be taken when the user does not respond, or when you want to configure an automatic action without the user's intervention.").AddStringEnumDescription("DEFER", "UPDATE").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("DEFER", "UPDATE"),
				},
			},
			"defer_update_minimum_secure_client_version": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Minimum Secure Client version (in x.x.x format) to be present on the client system to allow or defer the update.").String,
				Optional:            true,
			},
			"defer_update_prompt_dismiss_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Timeout (in seconds) for the prompt dismissal.").AddIntegerRangeDescription(1, 300).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 300),
				},
			},
		},
	}
}

func (r *SecureClientCustomAttributeResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *SecureClientCustomAttributeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	// Check if FMC client is connected to supports this object
	if r.client.FMCVersionParsed.LessThan(minFMCVersionCreateSecureClientCustomAttribute) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support Secure Client Custom Attribute creation, minumum required version is 7.4", r.client.FMCVersion))
		return
	}
	var plan SecureClientCustomAttribute

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
	body := plan.toBody(ctx, SecureClientCustomAttribute{})
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

func (r *SecureClientCustomAttributeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SecureClientCustomAttribute

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

func (r *SecureClientCustomAttributeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SecureClientCustomAttribute

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

func (r *SecureClientCustomAttributeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SecureClientCustomAttribute

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
func (r *SecureClientCustomAttributeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var config SecureClientCustomAttribute

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
		config.Domain = types.StringValue(tmpDomain)
	}
	config.Id = types.StringValue(match[inputPattern.SubexpIndex("id")])

	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
