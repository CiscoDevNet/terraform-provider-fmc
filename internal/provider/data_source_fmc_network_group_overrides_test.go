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

func TestAccDataSourceFmcNetworkGroupOverrides(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_network_group_overrides.test", "overrides.0.target_type", "Device"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_network_group_overrides.test", "overrides.0.description", "My Network Group Overridden name"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_network_group_overrides.test", "overrides.0.literals.0.value", "10.1.1.0/24"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcNetworkGroupOverridesPrerequisitesConfig + testAccDataSourceFmcNetworkGroupOverridesConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcNetworkGroupOverridesPrerequisitesConfig + testAccNamedDataSourceFmcNetworkGroupOverridesConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcNetworkGroupOverridesPrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id

resource "fmc_network_group" "test" {
  name        = "my_network_group_override"
  literals    = [{ value = "10.1.0.1" }]
  overridable = true
}

resource "fmc_range" "test" {
  name      = "fmc_network_group_override"
  ip_range  = "10.0.1.1-10.0.1.10"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcNetworkGroupOverridesConfig() string {
	config := `resource "fmc_network_group_overrides" "test" {` + "\n"
	config += `	parent_name = fmc_network_group.test.name` + "\n"
	config += `	parent_id = fmc_network_group.test.id` + "\n"
	config += `	overrides = [{` + "\n"
	config += `		target_id = var.device_id` + "\n"
	config += `		target_type = "Device"` + "\n"
	config += `		description = "My Network Group Overridden name"` + "\n"
	config += `		objects = [{` + "\n"
	config += `			id = fmc_range.test.id` + "\n"
	config += `			name = fmc_range.test.name` + "\n"
	config += `		}]` + "\n"
	config += `		literals = [{` + "\n"
	config += `			value = "10.1.1.0/24"` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_network_group_overrides" "test" {
			id = fmc_network_group_overrides.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcNetworkGroupOverridesConfig() string {
	config := `resource "fmc_network_group_overrides" "test" {` + "\n"
	config += `	parent_name = fmc_network_group.test.name` + "\n"
	config += `	parent_id = fmc_network_group.test.id` + "\n"
	config += `	overrides = [{` + "\n"
	config += `		target_id = var.device_id` + "\n"
	config += `		target_type = "Device"` + "\n"
	config += `		description = "My Network Group Overridden name"` + "\n"
	config += `		objects = [{` + "\n"
	config += `			id = fmc_range.test.id` + "\n"
	config += `			name = fmc_range.test.name` + "\n"
	config += `		}]` + "\n"
	config += `		literals = [{` + "\n"
	config += `			value = "10.1.1.0/24"` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_network_group_overrides" "test" {
			parent_name = fmc_network_group_overrides.test.parent_name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
