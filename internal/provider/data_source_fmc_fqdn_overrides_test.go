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

func TestAccDataSourceFmcFQDNOverrides(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_fqdn_overrides.test", "overrides.0.target_type", "Device"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_fqdn_overrides.test", "overrides.0.description", "My FQDN object"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_fqdn_overrides.test", "overrides.0.fqdn", "sub.example.com"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcFQDNOverridesPrerequisitesConfig + testAccDataSourceFmcFQDNOverridesConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcFQDNOverridesPrerequisitesConfig + testAccNamedDataSourceFmcFQDNOverridesConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcFQDNOverridesPrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id

resource "fmc_fqdn" "test" {
  name        = "my_fqdn_override"
  fqdn        = "example.com"
  overridable = true
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcFQDNOverridesConfig() string {
	config := `resource "fmc_fqdn_overrides" "test" {` + "\n"
	config += `	parent_name = fmc_fqdn.test.name` + "\n"
	config += `	parent_id = fmc_fqdn.test.id` + "\n"
	config += `	overrides = [{` + "\n"
	config += `		target_id = var.device_id` + "\n"
	config += `		target_type = "Device"` + "\n"
	config += `		description = "My FQDN object"` + "\n"
	config += `		fqdn = "sub.example.com"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_fqdn_overrides" "test" {
			id = fmc_fqdn_overrides.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcFQDNOverridesConfig() string {
	config := `resource "fmc_fqdn_overrides" "test" {` + "\n"
	config += `	parent_name = fmc_fqdn.test.name` + "\n"
	config += `	parent_id = fmc_fqdn.test.id` + "\n"
	config += `	overrides = [{` + "\n"
	config += `		target_id = var.device_id` + "\n"
	config += `		target_type = "Device"` + "\n"
	config += `		description = "My FQDN object"` + "\n"
	config += `		fqdn = "sub.example.com"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_fqdn_overrides" "test" {
			parent_name = fmc_fqdn_overrides.test.parent_name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
