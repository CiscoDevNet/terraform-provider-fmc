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

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DeviceCertificate struct {
	Id                      types.String `tfsdk:"id"`
	Domain                  types.String `tfsdk:"domain"`
	DeviceId                types.String `tfsdk:"device_id"`
	CertificateEnrollmentId types.String `tfsdk:"certificate_enrollment_id"`
	Type                    types.String `tfsdk:"type"`
	Name                    types.String `tfsdk:"name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionDeviceCertificate = version.Must(version.NewVersion("7.7"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DeviceCertificate) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/certificates"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DeviceCertificate) toBody(ctx context.Context, state DeviceCertificate) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.DeviceId.IsNull() {
		body, _ = sjson.Set(body, "deviceId", data.DeviceId.ValueString())
	}
	if !data.CertificateEnrollmentId.IsNull() {
		body, _ = sjson.Set(body, "certificate.id", data.CertificateEnrollmentId.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DeviceCertificate) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("deviceId"); value.Exists() {
		data.DeviceId = types.StringValue(value.String())
	} else {
		data.DeviceId = types.StringNull()
	}
	if value := res.Get("certificate.id"); value.Exists() {
		data.CertificateEnrollmentId = types.StringValue(value.String())
	} else {
		data.CertificateEnrollmentId = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *DeviceCertificate) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("deviceId"); value.Exists() && !data.DeviceId.IsNull() {
		data.DeviceId = types.StringValue(value.String())
	} else {
		data.DeviceId = types.StringNull()
	}
	if value := res.Get("certificate.id"); value.Exists() && !data.CertificateEnrollmentId.IsNull() {
		data.CertificateEnrollmentId = types.StringValue(value.String())
	} else {
		data.CertificateEnrollmentId = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DeviceCertificate) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
	if data.Name.IsUnknown() {
		if value := res.Get("name"); value.Exists() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin Clone

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// End of section. //template:end toBodyNonBulk

// Section below is generated&owned by "gen/generator.go". //template:begin findObjectsToBeReplaced

// End of section. //template:end findObjectsToBeReplaced

// Section below is generated&owned by "gen/generator.go". //template:begin clearItemIds

// End of section. //template:end clearItemIds

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyPutDelete

// End of section. //template:end toBodyPutDelete

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBody

// End of section. //template:end adjustBody

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBodyBulk

// End of section. //template:end adjustBodyBulk

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyOverrides

// End of section. //template:end toBodyOverrides

// Section below is generated&owned by "gen/generator.go". //template:begin synthesizeOverrides

// End of section. //template:end synthesizeOverrides

// Manual section - helpers for the hand-wired Device Certificate CRUD (Create/Read/Update/Delete) and the
// data source Read. Placed here (outside the autogeneration markers) so `go generate` does not remove them.

// deviceCertificateManagePath is the operational endpoint used to enroll, re-enroll and delete device certificates.
const deviceCertificateManagePath = "/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/operational/managecertificates"

// deviceCertificateManageBody builds the ManageDeviceCertificatesRequestModel payload for the given action
// (ENROLL, RE_ENROLL, REFRESH or DELETE). Only one certificate per device is supported by FMC.
func deviceCertificateManageBody(action, deviceId, certificateEnrollmentId string) string {
	body := ""
	body, _ = sjson.Set(body, "action", action)
	body, _ = sjson.Set(body, "deviceCertificates.0.deviceId", deviceId)
	body, _ = sjson.Set(body, "deviceCertificates.0.certificates.0.certificate.id", certificateEnrollmentId)
	return body
}

// deviceCertificateFindByEnrollment scans a GET /devices/certificates?filter=deviceId list response and returns the
// item whose enrolled certificate references the given certificate enrollment id (empty result if not found).
func deviceCertificateFindByEnrollment(res gjson.Result, certificateEnrollmentId string) gjson.Result {
	var match gjson.Result
	res.Get("items").ForEach(func(_, item gjson.Result) bool {
		item.Get("enrolledCertificates").ForEach(func(_, ec gjson.Result) bool {
			if ec.Get("certificate.id").String() == certificateEnrollmentId {
				match = item
				return false
			}
			return true
		})
		return !match.Exists()
	})
	return match
}
