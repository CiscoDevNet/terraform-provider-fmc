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
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &RadiusExternalAuthenticationObjectResource{}
	_ resource.ResourceWithImportState = &RadiusExternalAuthenticationObjectResource{}
)

func NewRadiusExternalAuthenticationObjectResource() resource.Resource {
	return &RadiusExternalAuthenticationObjectResource{}
}

type RadiusExternalAuthenticationObjectResource struct {
	client *fmc.Client
}

func (r *RadiusExternalAuthenticationObjectResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_radius_external_authentication_object"
}

func (r *RadiusExternalAuthenticationObjectResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages a RADIUS External Authentication object. It defines the RADIUS server(s) used to authenticate users, and can be referenced by `fmc_ftd_platform_settings_external_authentication` to enable RADIUS-based SSH/CLI authentication on FTD devices, or configured directly on the FMC (System > Users > External Authentication) for FMC login.\n User privileges (e.g. Administrator vs. read-only access) are derived from attributes returned by the RADIUS server for the authenticating user - most commonly the `Service-Type` attribute for FTD CLI access. This can be driven by Active Directory group membership if the RADIUS server (e.g. ISE or NPS) is configured to authenticate against AD and map groups to the appropriate attribute values. FMC does not configure this mapping; it is done entirely on the RADIUS server.").String,

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
				MarkdownDescription: helpers.NewAttributeDescription("Name of the RADIUS External Authentication object.").String,
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this value is always 'RADIUSConfigObject'.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description of the object.").String,
				Optional:            true,
			},
			"server_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP address or hostname of the primary RADIUS server.").String,
				Required:            true,
			},
			"server_port": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Port number of the primary RADIUS server.").AddDefaultValueDescription("1812").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("1812"),
			},
			"key": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Shared secret used to communicate with the primary RADIUS server.").String,
				Required:            true,
				Sensitive:           true,
			},
			"backup_server_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP address or hostname of the backup RADIUS server.").String,
				Optional:            true,
			},
			"backup_server_port": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Port number of the backup RADIUS server.").String,
				Optional:            true,
			},
			"backup_key": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Shared secret used to communicate with the backup RADIUS server.").String,
				Optional:            true,
				Sensitive:           true,
			},
			"timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Timeout (in seconds) before retrying the primary server.").AddIntegerRangeDescription(1, 1024).AddDefaultValueDescription("30").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 1024),
				},
				Default: int64default.StaticInt64(30),
			},
			"retries": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Number of retries before rolling over to the backup RADIUS server.").AddIntegerRangeDescription(0, 10).AddDefaultValueDescription("3").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 10),
				},
				Default: int64default.StaticInt64(3),
			},
			"message_authenticator_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enables RADIUS Server-Enabled Message Authenticator, requiring the Message-Authenticator attribute in all RADIUS responses.").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"cli_access_user_list": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Comma-separated list of usernames that should have CLI access, when using the predefined user list method instead of defining users on the RADIUS server. Leave unset when users and their privileges are managed on the RADIUS server (recommended when privileges are based on Active Directory group membership).").String,
				Optional:            true,
			},
		},
	}
}

func (r *RadiusExternalAuthenticationObjectResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *RadiusExternalAuthenticationObjectResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RadiusExternalAuthenticationObject

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
	body := plan.toBody(ctx, RadiusExternalAuthenticationObject{})
	res, err := r.client.Post(plan.getPath(), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}

	if retrievedId := res.Get("id"); retrievedId.Exists() {
		// FMC returned the id directly in the POST response (documented/expected behavior)
		plan.Id = types.StringValue(retrievedId.String())
	} else {
		// Some FMC/cdFMC deployments (observed via the Security Cloud Control API gateway) omit
		// the `id` field from the POST response for this endpoint. Fall back to looking up the
		// newly created object by its (unique) name.
		tflog.Debug(ctx, fmt.Sprintf("%s: id missing from POST response, looking up object by name %q", plan.Id.ValueString(), plan.Name.ValueString()))
		resList, err := r.client.Get(plan.getPath(), reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object after create, got error: %s", err))
			return
		}
		resList.Get("items").ForEach(func(k, v gjson.Result) bool {
			if v.Get("name").String() == plan.Name.ValueString() {
				plan.Id = types.StringValue(v.Get("id").String())
				return false
			}
			return true
		})
		if plan.Id.ValueString() == "" {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find created object with name: %s", plan.Name.ValueString()))
			return
		}

		// Fetch the full object, since the POST response body did not include all fields either.
		res, err = r.client.Get(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
			return
		}
	}
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *RadiusExternalAuthenticationObjectResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RadiusExternalAuthenticationObject

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

func (r *RadiusExternalAuthenticationObjectResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RadiusExternalAuthenticationObject

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

	if plan.Id.ValueString() == "" {
		resp.Diagnostics.AddError("Client Error", "Cannot update object: id is empty, refusing to send request to the collection endpoint")
		return
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

func (r *RadiusExternalAuthenticationObjectResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RadiusExternalAuthenticationObject

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

	if state.Id.ValueString() == "" {
		resp.Diagnostics.AddError("Client Error", "Cannot delete object: id is empty, refusing to send request to the collection endpoint")
		return
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
func (r *RadiusExternalAuthenticationObjectResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

// Section below is generated&owned by "gen/generator.go". //template:begin createSubresources

// End of section. //template:end createSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin deleteSubresources

// End of section. //template:end deleteSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin updateSubresources

// End of section. //template:end updateSubresources
