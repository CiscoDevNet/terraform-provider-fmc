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
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
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
	"github.com/netascode/terraform-provider-fmc/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource = &SmartLicenseResource{}
)

func NewSmartLicenseResource() resource.Resource {
	return &SmartLicenseResource{}
}

type SmartLicenseResource struct {
	client *fmc.Client
}

func (r *SmartLicenseResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_smart_license"
}

func (r *SmartLicenseResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Smart License.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"registration_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Smart license mode").AddStringEnumDescription("REGISTER", "EVALUATION").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("REGISTER", "EVALUATION"),
				},
			},
			"token": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Registration token. Mandatory when registration_type is set to REGISTER.").String,
				Optional:            true,
			},
			"registration_status": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Smart license registration status").String,
				Computed:            true,
			},
			"force": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set to true to force Smart License re-registration. This will take effect on each apply.").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"retain_registration": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set to true to keep registration after the resource is destroyed.").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
}

func (r *SmartLicenseResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

var _ resource.ResourceWithValidateConfig = &SmartLicenseResource{}

func (p *SmartLicenseResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data SmartLicense

	diags := req.Config.Get(ctx, &data)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	if data.RegistrationType.ValueString() == "REGISTER" && (data.Token.IsNull() || data.Token.IsUnknown()) {
		resp.Diagnostics.AddAttributeError(
			path.Root("token"),
			"Missing Attribute Configuration",
			"Token is required when registration_type is set to REGISTER",
		)
		return
	}
}

func (r *SmartLicenseResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, state SmartLicense

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Read state before create, so we can check what is the current registration status
	reqMods := [](func(*fmc.Req)){}
	res, err := r.client.Get(state.getPath(), reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	state.fromBody(ctx, res.Get("items.0"))

	// If smart license is already in evaluation mode and user requests evaluation mode - do nothing. Sending a request in such case would result in an error.
	if state.RegistrationStatus.ValueString() == "EVALUATION" && plan.RegistrationType.ValueString() == "EVALUATION" {
		tflog.Debug(ctx, fmt.Sprintf("%s: Smart license is already in evaluation mode, no action required", state.Id.ValueString()))
		plan.RegistrationStatus = state.RegistrationStatus
		plan.Id = types.StringValue(uuid.New().String())
		diags = resp.State.Set(ctx, &plan)
		resp.Diagnostics.Append(diags...)
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, SmartLicense{})
	body, _ = sjson.Delete(body, "dummy_retain_registration")
	body, _ = sjson.Delete(body, "dummy_force")
	res, err = r.client.Post(plan.getPath(), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}

	// Wait until registration is finished
	if plan.RegistrationType.ValueString() == "REGISTER" {
		res, diags = r.waitForRegistration(ctx, "REGISTERED")
	} else {
		res, diags = r.waitForRegistration(ctx, "EVALUATION")
	}
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	plan.Id = types.StringValue(uuid.New().String())
	plan.fromBodyUnknowns(ctx, res.Get("items.0"))

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *SmartLicenseResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SmartLicense

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	reqMods := [](func(*fmc.Req)){}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	res, err := r.client.Get(state.getPath(), reqMods...)
	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	state.fromBodyPartial(ctx, res.Get("items.0"))

	// Trigger re-creation if `force` flag is set
	if state.Force.ValueBool() {
		state.Force = types.BoolNull()
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *SmartLicenseResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SmartLicense

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

	reqMods := [](func(*fmc.Req)){}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	// If smart license is already in evaluation mode and user requests evaluation mode - do nothing. Sending a request in such case would result in an error.
	// It's not automatically detected by terraform, because two different fields keep status (RegistrationStatus) and requested status (RegistrationType)
	if state.RegistrationStatus.ValueString() == "EVALUATION" && plan.RegistrationType.ValueString() == "EVALUATION" {
		tflog.Debug(ctx, fmt.Sprintf("%s: Smart license is already in evaluation mode, no action required", state.Id.ValueString()))
		plan.RegistrationStatus = state.RegistrationStatus
		diags = resp.State.Set(ctx, &plan)
		resp.Diagnostics.Append(diags...)
		return
	}

	// If re-registration is forced, only deregister it the device is already registered
	if plan.Force.ValueBool() && state.RegistrationStatus.ValueString() == "REGISTERED" {
		tflog.Debug(ctx, fmt.Sprintf("%s: Force flag is set, deregistering smart license", plan.Id.ValueString()))
		res, err := r.deregisterSmartLicense(ctx, reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
			return
		}
		res, err = r.client.Get(state.getPath(), reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
			return
		}
		state.fromBody(ctx, res.Get("items.0"))
	}

	// Register FMC only if it is not already registered
	if state.RegistrationStatus.ValueString() != "REGISTERED" {
		tflog.Debug(ctx, fmt.Sprintf("%s: Starting registration", plan.Id.ValueString()))
		body := plan.toBody(ctx, SmartLicense{})
		body, _ = sjson.Delete(body, "dummy_retain_registration")
		body, _ = sjson.Delete(body, "dummy_force")
		res, err := r.client.Post(plan.getPath(), body, reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
			return
		}

		res, diags = r.waitForRegistration(ctx, "REGISTERED")
		if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
			return
		}
		state.fromBody(ctx, res.Get("items.0"))
	}
	plan.Id = state.Id
	plan.RegistrationStatus = state.RegistrationStatus

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *SmartLicenseResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SmartLicense

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	if state.RegistrationStatus.ValueString() == "EVALUATION" {
		tflog.Debug(ctx, fmt.Sprintf("%s: Smart license is in evaluation mode, cannot be deregistered", state.Id.ValueString()))
	} else if state.RetainRegistration.ValueBool() {
		tflog.Debug(ctx, fmt.Sprintf("%s: Retain registration flag is set, device is not unregistered", state.Id.ValueString()))
	} else {
		reqMods := [](func(*fmc.Req)){}

		tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
		res, err := r.deregisterSmartLicense(ctx, reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
			return
		}
		tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))
	}

	resp.State.RemoveResource(ctx)
}

// Deregister FMC from Smart Licensing
func (r *SmartLicenseResource) deregisterSmartLicense(ctx context.Context, reqMods ...func(*fmc.Req)) (gjson.Result, error) {
	plan := SmartLicense{
		RegistrationType: types.StringValue("DEREGISTER"),
	}
	body := plan.toBody(ctx, SmartLicense{})
	return r.client.Post(plan.getPath(), body, reqMods...)
}

// Wait for registration to finish
func (r *SmartLicenseResource) waitForRegistration(ctx context.Context, expectedState string) (gjson.Result, diag.Diagnostics) {
	var diag diag.Diagnostics
	const atom time.Duration = 5 * time.Second

	reqMods := [](func(*fmc.Req)){}

	for i := time.Duration(0); i < 5*time.Minute; i += atom {
		task, err := r.client.Get("/api/fmc_platform/v1/license/smartlicenses", reqMods...)
		if err != nil {
			diag.AddError("Client Error", fmt.Sprintf("Failed to read object (GET), got error: %s, %s", err, task.String()))
			return task, diag
		}
		stat := strings.ToUpper(task.Get("items.0.regStatus").String())
		if stat == expectedState {
			return task, nil
		}
		time.Sleep(atom)
	}
	diag.AddError("Client Error", "Registration failed, FMC did not finish registration in time")
	return gjson.Result{}, diag
}
