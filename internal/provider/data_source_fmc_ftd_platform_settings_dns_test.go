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

func TestAccDataSourceFmcFTDPlatformSettingsDNS(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_ftd_platform_settings_dns.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_platform_settings_dns.test", "server_groups.0.is_default", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_platform_settings_dns.test", "expire_entry_timer", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_platform_settings_dns.test", "poll_timer", "240"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_platform_settings_dns.test", "lookup_via_management_diagnostic_interface", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcFTDPlatformSettingsDNSPrerequisitesConfig + testAccDataSourceFmcFTDPlatformSettingsDNSConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcFTDPlatformSettingsDNSPrerequisitesConfig = `
resource "fmc_ftd_platform_settings" "test" {
  name        = "ftd_platform_settings_banner"
}

resource "fmc_interface_group" "test" {
  name           = "ftd_platform_settings_dns_ig"
  interface_mode = "ROUTED"
}

resource "fmc_dns_server_group" "test" {
  name           = "ftd_platform_settings_dns_dns_sg"
  default_domain = "example.com"
  dns_servers = [
    { ip = "10.10.10.10" }
  ]
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcFTDPlatformSettingsDNSConfig() string {
	config := `resource "fmc_ftd_platform_settings_dns" "test" {` + "\n"
	config += `	ftd_platform_settings_id = fmc_ftd_platform_settings.test.id` + "\n"
	config += `	server_groups = [{` + "\n"
	config += `		server_group_id = fmc_dns_server_group.test.id` + "\n"
	config += `		is_default = true` + "\n"
	config += `	}]` + "\n"
	config += `	expire_entry_timer = 1` + "\n"
	config += `	poll_timer = 240` + "\n"
	config += `	interface_objects = [{` + "\n"
	config += `		id = fmc_interface_group.test.id` + "\n"
	config += `		type = fmc_interface_group.test.type` + "\n"
	config += `	}]` + "\n"
	config += `	lookup_via_management_diagnostic_interface = true` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_ftd_platform_settings_dns" "test" {
			id = fmc_ftd_platform_settings_dns.test.id
			ftd_platform_settings_id = fmc_ftd_platform_settings.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
