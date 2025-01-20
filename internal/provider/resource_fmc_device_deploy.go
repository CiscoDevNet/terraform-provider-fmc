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

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/netascode/terraform-provider-fmc/internal/provider/helpers"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource = &DeviceDeployResource{}
)

func NewDeviceDeployResource() resource.Resource {
	return &DeviceDeployResource{}
}

type DeviceDeployResource struct {
	client *fmc.Client
}

func (r *DeviceDeployResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_deploy"
}

func (r *DeviceDeployResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Device Deploy.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "The name of the FMC domain",
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"version": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Version to which the deployment should be done in milliseconds unix timestamp. If not provided, the latest version will be used.").String,
				Optional:            true,
			},
			"ignore_warning": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Ignore warnings during deployment.").String,
				Optional:            true,
			},
			"device_list": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of device ids to be deployed.").String,
				ElementType:         types.StringType,
				Required:            true,
			},
			"deployment_note": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("User note.").String,
				Optional:            true,
			},
		},
	}
}

func (r *DeviceDeployResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

func (r *DeviceDeployResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DeviceDeploy

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

	// Create ID
	plan.Id = types.StringValue(uuid.New().String())

	// Trigger deployment
	diags = r.triggerDeployment(ctx, plan, reqMods)
	resp.Diagnostics.Append(diags...)

	// Save state
	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *DeviceDeployResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DeviceDeploy

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Clear device list, this will trigger update
	state.DeviceList = types.ListNull(types.StringType)

	// Save state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *DeviceDeployResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan DeviceDeploy

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

	// Trigger deployment
	diags = r.triggerDeployment(ctx, plan, reqMods)
	resp.Diagnostics.Append(diags...)

	// Save state
	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *DeviceDeployResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DeviceDeploy

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

func (r *DeviceDeployResource) triggerDeployment(ctx context.Context, plan DeviceDeploy, reqMods [](func(*fmc.Req))) diag.Diagnostics {
	var diags diag.Diagnostics

	tflog.Debug(ctx, fmt.Sprintf("%s: Triggering deployment", plan.Id.ValueString()))

	// Save original device list
	origPlanDeviceList := plan.DeviceList
	// Restore original list of devices
	defer func() {
		plan.DeviceList = origPlanDeviceList
	}()

	// Get list of deployable devices
	resDeployable, err := r.client.Get("/api/fmc_config/v1/domain/{DOMAIN_UUID}/deployment/deployabledevices?expanded=true", reqMods...)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Failed to obtain list of deployable devices object (GET), got error: %s, %s", err, resDeployable.String()))
		return diags
	}

	// Get version (common across all the devices)
	// Force Deploy would require different way of obtaining version, as no devices may be in deployable state
	plan.Version = types.StringValue(resDeployable.Get("items.0.version").String())

	// If force_deploy is not set, make sure that all devices in plan are in deployable state
	//if !plan.ForceDeploy.ValueBool() {

	// Extract IDs or deployable devices
	tmpResDeployableIds := resDeployable.Get("items.#.device.id")
	var deployableDeviceIds []string
	tmpResDeployableIds.ForEach(func(_, value gjson.Result) bool {
		deployableDeviceIds = append(deployableDeviceIds, value.String())
		return true
	})

	tflog.Debug(ctx, fmt.Sprintf("%s: Deployable devices: %v", plan.Id.ValueString(), deployableDeviceIds))

	// List of devices that actually need deployment
	var devicesToBeDeployed []string
	var planDevices []string
	plan.DeviceList.ElementsAs(ctx, &planDevices, false)

	// Clear the list of devices to deploy
	plan.DeviceList = types.ListNull(types.StringType)

	// Check which devices from original list are deployable
	for _, device := range planDevices {
		if helpers.Contains(deployableDeviceIds, device) {
			devicesToBeDeployed = append(devicesToBeDeployed, device)
		}
	}

	plan.DeviceList = helpers.GetStringListFromStringSlice(devicesToBeDeployed)
	//}

	var deviceIdsSlice []string
	plan.DeviceList.ElementsAs(ctx, &deviceIdsSlice, false)

	if len(deviceIdsSlice) > 0 {
		// Trigger deployment if any devices are in deployable state
		tflog.Debug(ctx, fmt.Sprintf("%s: Deploying devices: %v", plan.Id.ValueString(), deviceIdsSlice))
		body := plan.toBody(ctx, DeviceDeploy{})
		res, err := r.client.Post(plan.getPath(), body, reqMods...)
		if err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
			return diags
		}

		// Wait for deployment to finish
		diags = helpers.FMCWaitForJobToFinish(ctx, r.client, res.Get("metadata.task.id").String(), reqMods...)
		if diags.HasError() {
			return diags
		}

		tflog.Debug(ctx, fmt.Sprintf("%s: Deploy completed", plan.Id.ValueString()))
	} else {
		tflog.Debug(ctx, fmt.Sprintf("%s: No devices in deployable state", plan.Id.ValueString()))
	}

	return nil
}
