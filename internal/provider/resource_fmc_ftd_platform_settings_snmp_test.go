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

func TestAccFmcFTDPlatformSettingsSNMP(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_ftd_platform_settings_snmp.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "enable_snmp_servers", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "system_administrator_name", "admin"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "location", "Data Center 1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "listen_port", "161"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "snmp_management_hosts.0.snmp_version", "SNMPv3"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "snmp_management_hosts.0.username", "snmpuser1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "snmp_management_hosts.0.poll", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "snmp_management_hosts.0.trap", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "snmp_management_hosts.0.trap_port", "162"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "snmp_management_hosts.0.use_device_management_interface", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "snmpv3_users.0.security_level", "Priv"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "snmpv3_users.0.username", "snmpuser1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "snmpv3_users.0.encryption_password_type", "Clear"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "snmpv3_users.0.authentication_algorithm_type", "SHA256"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "snmpv3_users.0.encryption_type", "AES256"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "trap_syslog", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "trap_authentication", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "trap_link_up", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "trap_link_down", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "trap_cold_start", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "trap_warm_start", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "trap_field_replacement_unit_insert", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "trap_field_replacement_unit_delete", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "trap_configuration_change", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "trap_connection_limit_reached", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "trap_nat_packet_discard", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "trap_cpu_rising_threshold", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "trap_cpu_rising_threshold_value", "70"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "trap_cpu_rising_threshold_interval", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "trap_memory_rising_threshold", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "trap_memory_rising_threshold_value", "70"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "trap_failover", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "trap_cluster", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_snmp.test", "trap_peer_flap", "true"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccFmcFTDPlatformSettingsSNMPPrerequisitesConfig + testAccFmcFTDPlatformSettingsSNMPConfig_all(),
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

const testAccFmcFTDPlatformSettingsSNMPPrerequisitesConfig = `
resource "fmc_ftd_platform_settings" "test" {
  name        = "ftd_platform_settings_snmp"
}

resource "fmc_host" "test" {
  name = "ftd_platform_settings_snmp_host1"
  ip   = "10.0.2.1"
}

resource "fmc_security_zone" "test" {
  name           = "ftd_platform_settings_snmp_host1"
  interface_type = "ROUTED"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcFTDPlatformSettingsSNMPConfig_all() string {
	config := `resource "fmc_ftd_platform_settings_snmp" "test" {` + "\n"
	config += `	ftd_platform_settings_id = fmc_ftd_platform_settings.test.id` + "\n"
	config += `	enable_snmp_servers = true` + "\n"
	config += `	read_community_string = "public"` + "\n"
	config += `	system_administrator_name = "admin"` + "\n"
	config += `	location = "Data Center 1"` + "\n"
	config += `	listen_port = 161` + "\n"
	config += `	snmp_management_hosts = [{` + "\n"
	config += `		management_host_ip_object_id = fmc_host.test.id` + "\n"
	config += `		snmp_version = "SNMPv3"` + "\n"
	config += `		username = "snmpuser1"` + "\n"
	config += `		poll = true` + "\n"
	config += `		trap = true` + "\n"
	config += `		trap_port = 162` + "\n"
	config += `		use_device_management_interface = false` + "\n"
	config += `		interface_objects = [{` + "\n"
	config += `			id = fmc_security_zone.test.id` + "\n"
	config += `			type = fmc_security_zone.test.type` + "\n"
	config += `			name = fmc_security_zone.test.name` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `	snmpv3_users = [{` + "\n"
	config += `		security_level = "Priv"` + "\n"
	config += `		username = "snmpuser1"` + "\n"
	config += `		encryption_password_type = "Clear"` + "\n"
	config += `		authentication_algorithm_type = "SHA256"` + "\n"
	config += `		authentication_password = "MyAuthPassword123"` + "\n"
	config += `		encryption_type = "AES256"` + "\n"
	config += `		encryption_password = "MyEncryptionPassword123"` + "\n"
	config += `	}]` + "\n"
	config += `	trap_syslog = true` + "\n"
	config += `	trap_authentication = true` + "\n"
	config += `	trap_link_up = true` + "\n"
	config += `	trap_link_down = true` + "\n"
	config += `	trap_cold_start = true` + "\n"
	config += `	trap_warm_start = true` + "\n"
	config += `	trap_field_replacement_unit_insert = true` + "\n"
	config += `	trap_field_replacement_unit_delete = true` + "\n"
	config += `	trap_configuration_change = true` + "\n"
	config += `	trap_connection_limit_reached = true` + "\n"
	config += `	trap_nat_packet_discard = true` + "\n"
	config += `	trap_cpu_rising_threshold = true` + "\n"
	config += `	trap_cpu_rising_threshold_value = 70` + "\n"
	config += `	trap_cpu_rising_threshold_interval = 1` + "\n"
	config += `	trap_memory_rising_threshold = true` + "\n"
	config += `	trap_memory_rising_threshold_value = 70` + "\n"
	config += `	trap_failover = true` + "\n"
	config += `	trap_cluster = true` + "\n"
	config += `	trap_peer_flap = true` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
