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

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccFmcFTDPlatformSettingsHTTPAccess(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_ftd_platform_settings_http_access.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_http_access.test", "server_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_http_access.test", "server_port", "443"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccFmcFTDPlatformSettingsHTTPAccessPrerequisitesConfig + testAccFmcFTDPlatformSettingsHTTPAccessConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
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

const testAccFmcFTDPlatformSettingsHTTPAccessPrerequisitesConfig = `
resource "fmc_ftd_platform_settings" "test" {
  name        = "ftd_platform_settings_http_access"
}

resource "fmc_host" "test" {
  name = "ftd_platform_settings_http_access_host1"
  ip   = "10.0.2.1"
}

resource "fmc_security_zone" "test" {
  name           = "ftd_platform_settings_http_access_zone1"
  interface_type = "ROUTED"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcFTDPlatformSettingsHTTPAccessConfig_all() string {
	config := `resource "fmc_ftd_platform_settings_http_access" "test" {` + "\n"
	config += `	ftd_platform_settings_id = fmc_ftd_platform_settings.test.id` + "\n"
	config += `	server_enabled = true` + "\n"
	config += `	server_port = 443` + "\n"
	config += `	configurations = [{` + "\n"
	config += `		source_network_object_id = fmc_host.test.id` + "\n"
	config += `		interface_objects = [{` + "\n"
	config += `			id = fmc_security_zone.test.id` + "\n"
	config += `			type = fmc_security_zone.test.type` + "\n"
	config += `			name = fmc_security_zone.test.name` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
