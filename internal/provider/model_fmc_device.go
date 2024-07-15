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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/terraform-provider-fmc/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type Device struct {
	Id                     types.String `tfsdk:"id"`
	Domain                 types.String `tfsdk:"domain"`
	Name                   types.String `tfsdk:"name"`
	HostName               types.String `tfsdk:"host_name"`
	NatId                  types.String `tfsdk:"nat_id"`
	LicenseCapabilities    types.Set    `tfsdk:"license_capabilities"`
	RegistrationKey        types.String `tfsdk:"registration_key"`
	Type                   types.String `tfsdk:"type"`
	AccessPolicyId         types.String `tfsdk:"access_policy_id"`
	NatPolicyId            types.String `tfsdk:"nat_policy_id"`
	ProhibitPacketTransfer types.Bool   `tfsdk:"prohibit_packet_transfer"`
	PerformanceTier        types.String `tfsdk:"performance_tier"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data Device) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data Device) toBody(ctx context.Context, state Device) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.HostName.IsNull() {
		body, _ = sjson.Set(body, "hostName", data.HostName.ValueString())
	}
	if !data.NatId.IsNull() {
		body, _ = sjson.Set(body, "natID", data.NatId.ValueString())
	}
	if !data.LicenseCapabilities.IsNull() {
		var values []string
		data.LicenseCapabilities.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "license_caps", values)
	}
	if !data.RegistrationKey.IsNull() {
		body, _ = sjson.Set(body, "regKey", data.RegistrationKey.ValueString())
	}
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, "type", data.Type.ValueString())
	}
	if !data.AccessPolicyId.IsNull() {
		body, _ = sjson.Set(body, "accessPolicy.id", data.AccessPolicyId.ValueString())
	}
	if !data.NatPolicyId.IsNull() {
		body, _ = sjson.Set(body, "dummy_nat_policy_id", data.NatPolicyId.ValueString())
	}
	if !data.ProhibitPacketTransfer.IsNull() {
		body, _ = sjson.Set(body, "prohibitPacketTransfer", data.ProhibitPacketTransfer.ValueBool())
	}
	if !data.PerformanceTier.IsNull() {
		body, _ = sjson.Set(body, "performanceTier", data.PerformanceTier.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *Device) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("hostName"); value.Exists() {
		data.HostName = types.StringValue(value.String())
	} else {
		data.HostName = types.StringNull()
	}
	if value := res.Get("license_caps"); value.Exists() {
		data.LicenseCapabilities = helpers.GetStringSet(value.Array())
	} else {
		data.LicenseCapabilities = types.SetNull(types.StringType)
	}
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringValue("Device")
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody

func (data *Device) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("hostName"); value.Exists() && !data.HostName.IsNull() {
		data.HostName = types.StringValue(value.String())
	} else {
		data.HostName = types.StringNull()
	}
	if value := res.Get("license_caps"); value.Exists() && !data.LicenseCapabilities.IsNull() {
		data.LicenseCapabilities = helpers.GetStringSet(value.Array())
	} else {
		data.LicenseCapabilities = types.SetNull(types.StringType)
	}
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else if data.Type.ValueString() != "Device" {
		data.Type = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

func (data *Device) fromPolicyBody(ctx context.Context, res gjson.Result) {
	query := fmt.Sprintf(`items.#(targets.#(id=="%s"))#.policy`, data.Id.ValueString())
	list := res.Get(query)
	tflog.Debug(ctx, fmt.Sprintf("gjson path %s resulted in %d policies: %s", query, len(list.Array()), list))

	if !list.Exists() {
		data.AccessPolicyId = types.StringNull()
		data.NatPolicyId = types.StringNull()
		return
	}

	value := list.Get(`#(type=="AccessPolicy").id`)
	tflog.Debug(ctx, fmt.Sprintf("gjson search AccessPolicy resulted in: %s", value))
	if value.Exists() {
		data.AccessPolicyId = types.StringValue(value.String())
	} else {
		data.AccessPolicyId = types.StringNull()
	}

	value = list.Get(`#(type=="FTDNatPolicy").id`)
	tflog.Debug(ctx, fmt.Sprintf("gjson search FTDNatPolicy resulted in: %s", value))
	if value.Exists() {
		data.NatPolicyId = types.StringValue(value.String())
	} else {
		data.NatPolicyId = types.StringNull()
	}
}

func (data *Device) updateFromPolicyBody(ctx context.Context, res gjson.Result) {
	query := fmt.Sprintf(`items.#(targets.#(id=="%s"))#.policy`, data.Id.ValueString())
	list := res.Get(query)
	tflog.Debug(ctx, fmt.Sprintf("gjson path %s resulted in %d policies for update: %s", query, len(list.Array()), list))

	if !list.Exists() {
		data.AccessPolicyId = types.StringNull()
		data.NatPolicyId = types.StringNull()
		return
	}

	value := list.Get(`#(type=="AccessPolicy").id`)
	tflog.Debug(ctx, fmt.Sprintf("gjson search AccessPolicy resulted in: %s", value))
	if value.Exists() && !data.AccessPolicyId.IsNull() {
		data.AccessPolicyId = types.StringValue(value.String())
	} else {
		data.AccessPolicyId = types.StringNull()
	}

	value = list.Get(`#(type=="FTDNatPolicy").id`)
	tflog.Debug(ctx, fmt.Sprintf("gjson search FTDNatPolicy resulted in: %s", value))
	if value.Exists() && !data.NatPolicyId.IsNull() {
		data.NatPolicyId = types.StringValue(value.String())
	} else {
		data.NatPolicyId = types.StringNull()
	}
}

// Section below is generated&owned by "gen/generator.go". //template:begin isNull

func (data *Device) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.HostName.IsNull() {
		return false
	}
	if !data.NatId.IsNull() {
		return false
	}
	if !data.LicenseCapabilities.IsNull() {
		return false
	}
	if !data.RegistrationKey.IsNull() {
		return false
	}
	if !data.Type.IsNull() {
		return false
	}
	if !data.AccessPolicyId.IsNull() {
		return false
	}
	if !data.NatPolicyId.IsNull() {
		return false
	}
	if !data.ProhibitPacketTransfer.IsNull() {
		return false
	}
	if !data.PerformanceTier.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
