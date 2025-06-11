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

func TestAccDataSourceFmcChassis(t *testing.T) {
	if os.Getenv("TF_VAR_chassis_registration_key") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_chassis_registration_key")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_chassis.test", "name", "my_chassis"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_chassis.test", "type"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcChassisConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccNamedDataSourceFmcChassisConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcChassisConfig() string {
	config := `resource "fmc_chassis" "test" {` + "\n"
	config += `	name = "my_chassis"` + "\n"
	config += `	host_name = var.chassis_addr` + "\n"
	config += `	registration_key = var.chassis_registration_key` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_chassis" "test" {
			id = fmc_chassis.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcChassisConfig() string {
	config := `resource "fmc_chassis" "test" {` + "\n"
	config += `	name = "my_chassis"` + "\n"
	config += `	host_name = var.chassis_addr` + "\n"
	config += `	registration_key = var.chassis_registration_key` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_chassis" "test" {
			name = fmc_chassis.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
