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

func TestAccDataSourceFmcFTDPlatformSettingsSyslogLoggingSetup(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_ftd_platform_settings_syslog_logging_setup.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_platform_settings_syslog_logging_setup.test", "enable_logging", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_platform_settings_syslog_logging_setup.test", "enable_logging_on_the_failover_standby_unit", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_platform_settings_syslog_logging_setup.test", "emblem_format", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_platform_settings_syslog_logging_setup.test", "send_debug_messages_as_syslog", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_platform_settings_syslog_logging_setup.test", "internal_buffer_memory_size", "4096"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_platform_settings_syslog_logging_setup.test", "fmc_logging_type", "VPN"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_platform_settings_syslog_logging_setup.test", "fmc_logging_level", "ERR"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_platform_settings_syslog_logging_setup.test", "ftp_server_host_id", "d3f5e8c0-1d4b-11b2-9f0e-ecf4bbf7a5c6"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_platform_settings_syslog_logging_setup.test", "ftp_server_username", "ftpuser"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_platform_settings_syslog_logging_setup.test", "ftp_server_path", "/logs"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_platform_settings_syslog_logging_setup.test", "ftp_server_interface_groups.0.id", "e7f5e8c0-1d4b-11b2-9f0e-ecf4bbf7a5c6"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_platform_settings_syslog_logging_setup.test", "flash", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_platform_settings_syslog_logging_setup.test", "flash_maximum_space", "3076"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_platform_settings_syslog_logging_setup.test", "flash_minimum_free_space", "1024"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcFTDPlatformSettingsSyslogLoggingSetupPrerequisitesConfig + testAccDataSourceFmcFTDPlatformSettingsSyslogLoggingSetupConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcFTDPlatformSettingsSyslogLoggingSetupPrerequisitesConfig = `
resource "fmc_ftd_platform_settings" "test" {
  name        = "ftd_platform_settings_banner"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcFTDPlatformSettingsSyslogLoggingSetupConfig() string {
	config := `resource "fmc_ftd_platform_settings_syslog_logging_setup" "test" {` + "\n"
	config += `	ftd_platform_settings_id = fmc_ftd_platform_settings.test.id` + "\n"
	config += `	enable_logging = true` + "\n"
	config += `	enable_logging_on_the_failover_standby_unit = false` + "\n"
	config += `	emblem_format = false` + "\n"
	config += `	send_debug_messages_as_syslog = false` + "\n"
	config += `	internal_buffer_memory_size = 4096` + "\n"
	config += `	fmc_logging_type = "VPN"` + "\n"
	config += `	fmc_logging_level = "ERR"` + "\n"
	config += `	ftp_server_host_id = "d3f5e8c0-1d4b-11b2-9f0e-ecf4bbf7a5c6"` + "\n"
	config += `	ftp_server_username = "ftpuser"` + "\n"
	config += `	ftp_server_password = "ftppassword"` + "\n"
	config += `	ftp_server_path = "/logs"` + "\n"
	config += `	ftp_server_interface_groups = [{` + "\n"
	config += `		id = "e7f5e8c0-1d4b-11b2-9f0e-ecf4bbf7a5c6"` + "\n"
	config += `	}]` + "\n"
	config += `	flash = false` + "\n"
	config += `	flash_maximum_space = 3076` + "\n"
	config += `	flash_minimum_free_space = 1024` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_ftd_platform_settings_syslog_logging_setup" "test" {
			id = fmc_ftd_platform_settings_syslog_logging_setup.test.id
			ftd_platform_settings_id = fmc_ftd_platform_settings.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
