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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/tidwall/gjson"
)

func FMCWaitForJobToFinish(ctx context.Context, client *fmc.Client, jobId string, reqMods [](func(*fmc.Req))) diag.Diagnostics {
	var diags diag.Diagnostics
	const atom time.Duration = 5 * time.Second

	for i := time.Duration(0); i < 5*time.Minute; i += atom {
		task, err := client.Get("/api/fmc_config/v1/domain/{DOMAIN_UUID}/job/taskstatuses/"+url.QueryEscape(jobId), reqMods...)
		if err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Failed to read object (GET), got error: %s, %s", err, task.String()))
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

	// List of devices that actually need deployment
	var devicesToBeDeployed []string

	// List of device IDs that should be deployed
	var planDevices []string
	plan.DeviceIdList.ElementsAs(ctx, &planDevices, false)

	// Check if any of the devices is under deployment
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

		// Wait for deployment to finish
		diags = FMCWaitForJobToFinish(ctx, client, res.Get("metadata.task.id").String(), reqMods)
		if diags.HasError() {
			return diags
		}

		tflog.Debug(ctx, fmt.Sprintf("%s: Deploy completed", plan.Id.ValueString()))
	} else {
		tflog.Debug(ctx, fmt.Sprintf("%s: No devices in deployable state", plan.Id.ValueString()))
	}

	return diags
}
