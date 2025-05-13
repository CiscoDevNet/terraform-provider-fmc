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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DeviceDeploy struct {
	Id             types.String `tfsdk:"id"`
	Domain         types.String `tfsdk:"domain"`
	Version        types.String `tfsdk:"version"`
	IgnoreWarning  types.Bool   `tfsdk:"ignore_warning"`
	DeviceIdList   types.List   `tfsdk:"device_id_list"`
	DeploymentNote types.String `tfsdk:"deployment_note"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DeviceDeploy) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/deployment/deploymentrequests"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DeviceDeploy) toBody(ctx context.Context, state DeviceDeploy) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	body, _ = sjson.Set(body, "type", "DeploymentRequest")
	if !data.Version.IsNull() {
		body, _ = sjson.Set(body, "version", data.Version.ValueString())
	}
	if !data.IgnoreWarning.IsNull() {
		body, _ = sjson.Set(body, "ignoreWarning", data.IgnoreWarning.ValueBool())
	}
	if !data.DeviceIdList.IsNull() {
		var values []string
		data.DeviceIdList.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "deviceList", values)
	}
	if !data.DeploymentNote.IsNull() {
		body, _ = sjson.Set(body, "deploymentNote", data.DeploymentNote.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin Clone

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// End of section. //template:end toBodyNonBulk
