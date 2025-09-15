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

func TestAccDataSourceFmcFTDPlatformSettingsSyslogSettings(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_ftd_platform_settings_syslog_settings.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_platform_settings_syslog_settings.test", "facility", "LOCAL4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_platform_settings_syslog_settings.test", "timestamp_format", "RFC_5424"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_platform_settings_syslog_settings.test", "device_id_type", "USERDEFINEDID"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_platform_settings_syslog_settings.test", "device_id_user_defined_id", "my_device_id"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_platform_settings_syslog_settings.test", "device_id_interface_id", "123e4567-e89b-12d3-a456-426614174000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_platform_settings_syslog_settings.test", "all_syslog_messages", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_platform_settings_syslog_settings.test", "all_syslog_messages_logging_level", "ERR"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcFTDPlatformSettingsSyslogSettingsPrerequisitesConfig + testAccDataSourceFmcFTDPlatformSettingsSyslogSettingsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcFTDPlatformSettingsSyslogSettingsPrerequisitesConfig = `
resource "fmc_ftd_platform_settings" "test" {
  name        = "ftd_platform_settings_banner"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcFTDPlatformSettingsSyslogSettingsConfig() string {
	config := `resource "fmc_ftd_platform_settings_syslog_settings" "test" {` + "\n"
	config += `	ftd_platform_settings_id = fmc_ftd_platform_settings.test.id` + "\n"
	config += `	facility = "LOCAL4"` + "\n"
	config += `	timestamp_format = "RFC_5424"` + "\n"
	config += `	device_id_type = "USERDEFINEDID"` + "\n"
	config += `	device_id_user_defined_id = "my_device_id"` + "\n"
	config += `	device_id_interface_id = "123e4567-e89b-12d3-a456-426614174000"` + "\n"
	config += `	all_syslog_messages = true` + "\n"
	config += `	all_syslog_messages_logging_level = "ERR"` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_ftd_platform_settings_syslog_settings" "test" {
			id = fmc_ftd_platform_settings_syslog_settings.test.id
			ftd_platform_settings_id = fmc_ftd_platform_settings.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
