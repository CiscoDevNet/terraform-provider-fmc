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

func TestAccFmcFTDPlatformSettingsSyslogServers(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_ftd_platform_settings_syslog_servers.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_syslog_servers.test", "allow_user_traffic_when_tcp_syslog_server_is_down", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_syslog_servers.test", "message_queue_size", "512"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_syslog_servers.test", "servers.0.ip_object_id", "c1a0f5b6-3d4e-11b2-9f8f-0242ac112345"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_syslog_servers.test", "servers.0.protocol", "TCP"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_syslog_servers.test", "servers.0.port", "1470"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_syslog_servers.test", "servers.0.secure_syslog", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_syslog_servers.test", "servers.0.use_management_interface", "false"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccFmcFTDPlatformSettingsSyslogServersPrerequisitesConfig + testAccFmcFTDPlatformSettingsSyslogServersConfig_all(),
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

const testAccFmcFTDPlatformSettingsSyslogServersPrerequisitesConfig = `
resource "fmc_ftd_platform_settings" "test" {
  name        = "ftd_platform_settings_banner"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcFTDPlatformSettingsSyslogServersConfig_all() string {
	config := `resource "fmc_ftd_platform_settings_syslog_servers" "test" {` + "\n"
	config += `	ftd_platform_settings_id = fmc_ftd_platform_settings.test.id` + "\n"
	config += `	allow_user_traffic_when_tcp_syslog_server_is_down = true` + "\n"
	config += `	message_queue_size = 512` + "\n"
	config += `	servers = [{` + "\n"
	config += `		ip_object_id = "c1a0f5b6-3d4e-11b2-9f8f-0242ac112345"` + "\n"
	config += `		protocol = "TCP"` + "\n"
	config += `		port = 1470` + "\n"
	config += `		secure_syslog = true` + "\n"
	config += `		use_management_interface = false` + "\n"
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
