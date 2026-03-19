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
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type SystemInformation struct {
	Hostname      types.String `tfsdk:"hostname"`
	LspVersion    types.String `tfsdk:"lsp_version"`
	Model         types.String `tfsdk:"model"`
	Platform      types.String `tfsdk:"platform"`
	SerialNumber  types.String `tfsdk:"serial_number"`
	ServerVersion types.String `tfsdk:"server_version"`
	SruVersion    types.String `tfsdk:"sru_version"`
	Type          types.String `tfsdk:"type"`
	Uptime        types.String `tfsdk:"uptime"`
	VdbVersion    types.String `tfsdk:"vdb_version"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SystemInformation) getPath() string {
	return "/api/fmc_platform/v1/info/serverversion"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SystemInformation) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("items.0.hostname"); value.Exists() {
		data.Hostname = types.StringValue(value.String())
	} else {
		data.Hostname = types.StringNull()
	}
	if value := res.Get("items.0.lspVersion"); value.Exists() {
		data.LspVersion = types.StringValue(value.String())
	} else {
		data.LspVersion = types.StringNull()
	}
	if value := res.Get("items.0.model"); value.Exists() {
		data.Model = types.StringValue(value.String())
	} else {
		data.Model = types.StringNull()
	}
	if value := res.Get("items.0.platform"); value.Exists() {
		data.Platform = types.StringValue(value.String())
	} else {
		data.Platform = types.StringNull()
	}
	if value := res.Get("items.0.serialNumber"); value.Exists() {
		data.SerialNumber = types.StringValue(value.String())
	} else {
		data.SerialNumber = types.StringNull()
	}
	if value := res.Get("items.0.serverVersion"); value.Exists() {
		data.ServerVersion = types.StringValue(value.String())
	} else {
		data.ServerVersion = types.StringNull()
	}
	if value := res.Get("items.0.sruVersion"); value.Exists() {
		data.SruVersion = types.StringValue(value.String())
	} else {
		data.SruVersion = types.StringNull()
	}
	if value := res.Get("items.0.type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("items.0.uptime"); value.Exists() {
		data.Uptime = types.StringValue(value.String())
	} else {
		data.Uptime = types.StringNull()
	}
	if value := res.Get("items.0.vdbVersion"); value.Exists() {
		data.VdbVersion = types.StringValue(value.String())
	} else {
		data.VdbVersion = types.StringNull()
	}
}

// End of section. //template:end fromBody
