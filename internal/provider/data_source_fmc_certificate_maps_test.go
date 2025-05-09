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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource

func TestAccDataSourceFmcCertificateMaps(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_certificate_maps.test", "items.my_certificate_maps.id"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_certificate_maps.test", "items.my_certificate_maps.type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_certificate_maps.test", "items.my_certificate_maps.rules.0.field", "SUBJECT"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_certificate_maps.test", "items.my_certificate_maps.rules.0.component", "COMMON_NAME"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_certificate_maps.test", "items.my_certificate_maps.rules.0.operator", "EQUALS"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_certificate_maps.test", "items.my_certificate_maps.rules.0.value", "cisco.com"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcCertificateMapsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcCertificateMapsConfig() string {
	config := `resource "fmc_certificate_maps" "test" {` + "\n"
	config += `	items = { "my_certificate_maps" = {` + "\n"
	config += `		rules = [{` + "\n"
	config += `			field = "SUBJECT"` + "\n"
	config += `			component = "COMMON_NAME"` + "\n"
	config += `			operator = "EQUALS"` + "\n"
	config += `			value = "cisco.com"` + "\n"
	config += `		}]` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_certificate_maps" "test" {
			depends_on = [fmc_certificate_maps.test]
			items = {
				"my_certificate_maps" = {
				}
			}
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
