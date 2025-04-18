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

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func FMCWaitForJobToFinish(ctx context.Context, client *fmc.Client, jobId string, reqMods [](func(*fmc.Req))) diag.Diagnostics {
	var diags diag.Diagnostics
	const atom time.Duration = 5 * time.Second

	for i := time.Duration(0); i < 5*time.Minute; i += atom {
		task, err := client.Get("/api/fmc_config/v1/domain/{DOMAIN_UUID}/job/taskstatuses/"+url.QueryEscape(jobId), reqMods...)
		if err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Failed to get task status, got error: %s, %s", err, task.String()))
			return diags
		}
		stat := strings.ToUpper(task.Get("status").String())
		if stat == "FAILED" {
			diags.AddError("Client Error", fmt.Sprintf("API task for the new device failed: %s, %s", task.Get("message"), task.Get("description")))
			return diags
		}
		if stat != "PENDING" && stat != "RUNNING" && stat != "IN_PROGRESS" && stat != "DEPLOYING" && stat != "UNKNOWN" {
			break
		}
		time.Sleep(atom)
	}
	return nil
}

func FMCWaitForDeploymentToFinish(ctx context.Context, client *fmc.Client, deviceIds []string, reqMods [](func(*fmc.Req))) diag.Diagnostics {
	var diags diag.Diagnostics
	const atom time.Duration = 5 * time.Second

Outerloop:
	for i := time.Duration(0); i < 5*time.Minute; i += atom {
		underDeployment, err := client.Get("/api/fmc_config/v1/domain/{DOMAIN_UUID}/deployment/jobhistories?filter=jobType:DEPLOYMENT;status:DEPLOYING&expanded=true", reqMods...)
		if err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Failed to read object (GET), got error: %s, %s", err, underDeployment.String()))
			return diags
		}

		// If there are no deployments in progress, exit
		if !underDeployment.Get("items").Exists() {
			break
		}

		// Extract IDs of devices, that are currently under deployment
		var devicesUnderDeploymentIds []string
		query := "items.#.deviceList.#.deviceUUID"
		res := underDeployment.Get(query).Array()

		for _, v := range res {
			for _, vv := range v.Array() {
				devicesUnderDeploymentIds = append(devicesUnderDeploymentIds, vv.String())
			}
		}

		// If any of the devices from provided list is under deployment, wait
		for _, deviceId := range deviceIds {
			if helpers.Contains(devicesUnderDeploymentIds, deviceId) {
				tflog.Debug(ctx, fmt.Sprintf("Device %s is still under deployment. Waiting.", deviceId))
				time.Sleep(atom)
				continue Outerloop
			}
		}

		// If none of the devices from provided list is under deployment, exit
		break Outerloop
	}
	return nil
}

func FMCDeviceDeploy(ctx context.Context, client *fmc.Client, plan DeviceDeploy, reqMods [](func(*fmc.Req))) diag.Diagnostics {
	var diags diag.Diagnostics

	// List of devices that are going to be deployed (requested by the user and in deployable state)
	var devicesToBeDeployed []string

	// List of device IDs that are requested to be deployed
	var planDevices []string
	plan.DeviceIdList.ElementsAs(ctx, &planDevices, false)

	// Wait for any deployments in progress to finish
	diags = FMCWaitForDeploymentToFinish(ctx, client, planDevices, reqMods)
	if diags.HasError() {
		return diags
	}

	// Get list of deployable devices
	resDeployable, err := client.Get("/api/fmc_config/v1/domain/{DOMAIN_UUID}/deployment/deployabledevices?expanded=true", reqMods...)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Failed to obtain list of deployable devices object (GET), got error: %s, %s", err, resDeployable.String()))
		return diags
	}

	// Get version (common across all the devices)
	// Force Deploy would require different way of obtaining version, as no devices may be in deployable state
	plan.Version = types.StringValue(resDeployable.Get("items.0.version").String())

	// Extract IDs or deployable devices
	tmpResDeployableIds := resDeployable.Get("items.#.device.id")
	var deployableDeviceIds []string
	tmpResDeployableIds.ForEach(func(_, value gjson.Result) bool {
		deployableDeviceIds = append(deployableDeviceIds, value.String())
		return true
	})

	tflog.Debug(ctx, fmt.Sprintf("%s: Deployable devices: %v", plan.Id.ValueString(), deployableDeviceIds))

	// Clear the list of devices to deploy
	plan.DeviceIdList = types.ListNull(types.StringType)

	// List of devices that actually need deployment
	for _, device := range planDevices {
		if helpers.Contains(deployableDeviceIds, device) {
			devicesToBeDeployed = append(devicesToBeDeployed, device)
		}
	}

	// Put device list into deployment payload
	plan.DeviceIdList = helpers.GetStringListFromStringSlice(devicesToBeDeployed)

	// Prepare payload for deployment
	if len(devicesToBeDeployed) > 0 {
		// Trigger deployment if any devices are in deployable state
		tflog.Debug(ctx, fmt.Sprintf("%s: Deploying devices: %v", plan.Id.ValueString(), devicesToBeDeployed))
		body := plan.toBody(ctx, DeviceDeploy{})
		res, err := client.Post(plan.getPath(), body, reqMods...)
		if err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
			return diags
		}

		if res.Get("metadata.task.id").Exists() {
			taskID := res.Get("metadata.task.id").String()
			tflog.Debug(ctx, fmt.Sprintf("%s: Async task initiated successfully (id: %s)", plan.Id.ValueString(), taskID))
			// Wait for deployment to finish
			diags = FMCWaitForJobToFinish(ctx, client, taskID, reqMods)
			if diags.HasError() {
				return diags
			}
		} else {
			tflog.Debug(ctx, fmt.Sprintf("%s: No task ID returned", plan.Id.ValueString()))
		}
		tflog.Debug(ctx, fmt.Sprintf("%s: Deploy completed", plan.Id.ValueString()))
	} else {
		tflog.Debug(ctx, fmt.Sprintf("%s: No devices in deployable state", plan.Id.ValueString()))
	}

	return diags
}

func FMCupdateDeviceGroup(ctx context.Context, client *fmc.Client, device basetypes.StringValue, plan tfsdk.Plan, state tfsdk.State, reqMods [](func(*fmc.Req))) diag.Diagnostics {
	deviceId := device.ValueString()

	// Extract IDs of current and planned device group
	var planDeviceGroup types.String
	if diags := plan.GetAttribute(ctx, path.Root("device_group_id"), &planDeviceGroup); diags.HasError() {
		return diags
	}

	var stateDeviceGroup types.String
	if diags := state.GetAttribute(ctx, path.Root("device_group_id"), &stateDeviceGroup); diags.HasError() {
		return diags
	}

	// When device group changes, we need to remove the device from the current group and add it to the new one

	if !stateDeviceGroup.IsNull() {
		// Get Device Group to which device currently belongs
		res, err := client.Get("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devicegroups/devicegrouprecords/"+url.QueryEscape(stateDeviceGroup.ValueString()), reqMods...)
		if err != nil {
			return diag.Diagnostics{diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Failed to get current device group (GET), got error: %s, %s", err, res.String()))}
		}

		// Remove device from current device group
		// Device Group endpoint returns 'metadata' twice in the response, which breaks sjson.Delete. Hence we copy needed fields to a new JSON.
		var request = ""
		resString := res.String()
		request, _ = sjson.Set(request, "id", gjson.Get(resString, "id").String())
		request, _ = sjson.Set(request, "type", gjson.Get(resString, "type").String())
		request, _ = sjson.Set(request, "name", gjson.Get(resString, "name").String())

		// Get all members
		members := gjson.Get(resString, "members").Array()
		filteredMembers := []string{}

		// Filter out the device to be removed
		for _, member := range members {
			if member.Get("id").String() != deviceId {
				filteredMembers = append(filteredMembers, member.Raw)
			}
		}

		// Update request
		request, _ = sjson.SetRaw(request, "members", fmt.Sprintf("[%s]", strings.Join(filteredMembers, ",")))

		// Make the PUT request
		res, err = client.Put("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devicegroups/devicegrouprecords/"+url.QueryEscape(stateDeviceGroup.ValueString()), request, reqMods...)
		if err != nil {
			return diag.Diagnostics{diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Failed to remove (PUT) device from current device group, got error: %s, %s", err, res.String()))}
		}
	}

	if !planDeviceGroup.IsNull() {
		// Get Device Group to which device needs to be assigned
		res, err := client.Get("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devicegroups/devicegrouprecords/"+url.QueryEscape(planDeviceGroup.ValueString()), reqMods...)
		if err != nil {
			return diag.Diagnostics{diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Failed to get destination device group (GET), got error: %s, %s", err, res.String()))}
		}

		// Put device to new device group
		// Device Group endpoint returns 'metadata' twice in the response, which breaks sjson.Delete. Hence we copy needed fields to a new JSON.
		var request = ""
		resString := res.String()
		request, _ = sjson.Set(request, "id", gjson.Get(resString, "id").String())
		request, _ = sjson.Set(request, "type", gjson.Get(resString, "type").String())
		request, _ = sjson.Set(request, "name", gjson.Get(resString, "name").String())
		members := gjson.Get(resString, "members")
		if members.Exists() {
			request, _ = sjson.SetRaw(request, "members", members.Raw)
		} else {
			request, _ = sjson.SetRaw(request, "members", "[]")
		}
		request, _ = sjson.SetRaw(request, "members.-1", fmt.Sprintf(`{"id":"%s"}`, deviceId))
		res, err = client.Put("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devicegroups/devicegrouprecords/"+url.QueryEscape(planDeviceGroup.ValueString()), request, reqMods...)
		if err != nil {
			return diag.Diagnostics{diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Failed to add (PUT) device to new device group, got error: %s, %s", err, res.String()))}
		}
	}

	return nil
}

// Check if device of a given ID is multi-instance
func FMCIsDeviceMultiInstance(ctx context.Context, client *fmc.Client, deviceId string, reqMods [](func(*fmc.Req))) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	res, err := client.Get(fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords/%v", url.QueryEscape(deviceId)), reqMods...)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Failed to retrieve devicerecord (GET), got error: %s, %s", err, res.String()))
		return false, diags
	}
	if res.Get("metadata.isMultiInstance").Exists() && res.Get("metadata.isMultiInstance").Bool() {
		tflog.Debug(ctx, fmt.Sprintf("Device %s is multi-instance", deviceId))
		return true, diags
	}
	return false, diags
}
