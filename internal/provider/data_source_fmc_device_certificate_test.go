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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource

func TestAccDataSourceFmcDeviceCertificate(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_device_certificate.test", "type"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcDeviceCertificatePrerequisitesConfig + testAccDataSourceFmcDeviceCertificateConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcDeviceCertificatePrerequisitesConfig + testAccNamedByDeviceIdDataSourceFmcDeviceCertificateConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcDeviceCertificatePrerequisitesConfig + testAccNamedByCertificateEnrollmentIdDataSourceFmcDeviceCertificateConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcDeviceCertificatePrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id

resource "fmc_certificate_enrollment" "test" {
  name                    = "device_certificate_enrollment"
  enrollment_type         = "SELF_SIGNED_CERTFICATE"
  certificate_common_name = "ftd.example.com"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcDeviceCertificateConfig() string {
	config := `resource "fmc_device_certificate" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	certificate_enrollment_id = fmc_certificate_enrollment.test.id` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_certificate" "test" {
			id = fmc_device_certificate.test.id
		}
	`
	return config
}

func testAccNamedByDeviceIdDataSourceFmcDeviceCertificateConfig() string {
	config := `resource "fmc_device_certificate" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	certificate_enrollment_id = fmc_certificate_enrollment.test.id` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_certificate" "test" {
			device_id = fmc_device_certificate.test.device_id
		}
	`
	return config
}
func testAccNamedByCertificateEnrollmentIdDataSourceFmcDeviceCertificateConfig() string {
	config := `resource "fmc_device_certificate" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	certificate_enrollment_id = fmc_certificate_enrollment.test.id` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_certificate" "test" {
			certificate_enrollment_id = fmc_device_certificate.test.certificate_enrollment_id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
