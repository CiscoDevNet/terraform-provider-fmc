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

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccFmcServiceAccess(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_service_access.test", "name", "my_service_access"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_service_access.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_service_access.test", "default_action", "DENY"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_service_access.test", "rules.0.action", "ALLOW"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcServiceAccessPrerequisitesConfig + testAccFmcServiceAccessConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcServiceAccessPrerequisitesConfig + testAccFmcServiceAccessConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_service_access.test",
		ImportState:  true,
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccFmcServiceAccessPrerequisitesConfig = `
data "fmc_countries" "test" {
  items = {
    "Poland"  = {}
  }
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcServiceAccessConfig_minimum() string {
	config := `resource "fmc_service_access" "test" {` + "\n"
	config += `	name = "my_service_access"` + "\n"
	config += `	default_action = "DENY"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcServiceAccessConfig_all() string {
	config := `resource "fmc_service_access" "test" {` + "\n"
	config += `	name = "my_service_access"` + "\n"
	config += `	default_action = "DENY"` + "\n"
	config += `	rules = [{` + "\n"
	config += `		action = "ALLOW"` + "\n"
	config += `		geolocation_sources = [{` + "\n"
	config += `			id = data.fmc_countries.test.items["Poland"].id` + "\n"
	config += `			type = data.fmc_countries.test.items["Poland"].type` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
