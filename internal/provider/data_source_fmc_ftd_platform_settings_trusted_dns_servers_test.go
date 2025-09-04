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

func TestAccDataSourceFmcFTDPlatformSettingsTrustedDNSServers(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_ftd_platform_settings_trusted_dns_servers.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_platform_settings_trusted_dns_servers.test", "trust_any_dns_server", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_platform_settings_trusted_dns_servers.test", "trust_dhcp_pool", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_platform_settings_trusted_dns_servers.test", "trust_dhcp_relay", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_platform_settings_trusted_dns_servers.test", "trust_dhcp_client", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_platform_settings_trusted_dns_servers.test", "trust_dns_server_group", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcFTDPlatformSettingsTrustedDNSServersPrerequisitesConfig + testAccDataSourceFmcFTDPlatformSettingsTrustedDNSServersConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcFTDPlatformSettingsTrustedDNSServersPrerequisitesConfig = `
resource "fmc_ftd_platform_settings" "test" {
  name        = "ftd_platform_settings_trusted_dns_server"
}

resource "fmc_host" "test" {
  name = "ftd_platform_settings_trusted_dns_server"
  ip   = "10.0.2.1"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcFTDPlatformSettingsTrustedDNSServersConfig() string {
	config := `resource "fmc_ftd_platform_settings_trusted_dns_servers" "test" {` + "\n"
	config += `	ftd_platform_settings_id = fmc_ftd_platform_settings.test.id` + "\n"
	config += `	trust_any_dns_server = false` + "\n"
	config += `	trust_dhcp_pool = true` + "\n"
	config += `	trust_dhcp_relay = true` + "\n"
	config += `	trust_dhcp_client = true` + "\n"
	config += `	trust_dns_server_group = true` + "\n"
	config += `	trusted_dns_servers_literals = ["10.20.30.40"]` + "\n"
	config += `	trusted_dns_servers_objects = [{` + "\n"
	config += `		id = fmc_host.test.id` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_ftd_platform_settings_trusted_dns_servers" "test" {
			id = fmc_ftd_platform_settings_trusted_dns_servers.test.id
			ftd_platform_settings_id = fmc_ftd_platform_settings.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
