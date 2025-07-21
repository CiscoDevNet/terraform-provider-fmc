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

func TestAccDataSourceFmcRealmADLDAP(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_realm_ad_ldap.test", "name", "my_ldap_realm"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_realm_ad_ldap.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_realm_ad_ldap.test", "version"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_realm_ad_ldap.test", "description", "My realm"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_realm_ad_ldap.test", "realm_type", "LDAP"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_realm_ad_ldap.test", "ad_primary_domain", "example.com"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_realm_ad_ldap.test", "ad_join_username", "user@example.com"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_realm_ad_ldap.test", "directory_username", "user@example.com"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_realm_ad_ldap.test", "base_dn", "DC=example,DC=com"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_realm_ad_ldap.test", "group_dn", "CN=users,DC=example,DC=com"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_realm_ad_ldap.test", "update_hour", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_realm_ad_ldap.test", "update_interval", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_realm_ad_ldap.test", "group_attribute", "member"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_realm_ad_ldap.test", "timeout_ise_users", "1440"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_realm_ad_ldap.test", "timeout_terminal_server_agent_users", "1440"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_realm_ad_ldap.test", "timeout_captive_portal_users", "1440"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_realm_ad_ldap.test", "timeout_failed_captive_portal_users", "1440"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_realm_ad_ldap.test", "timeout_guest_captive_portal_users", "1440"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_realm_ad_ldap.test", "directory_server_configurations.0.hostname", "ldap.example.com"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_realm_ad_ldap.test", "directory_server_configurations.0.port", "389"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_realm_ad_ldap.test", "directory_server_configurations.0.encryption_protocol", "LDAPS"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_realm_ad_ldap.test", "directory_server_configurations.0.encryption_certificate", "1234-5678-90AB-CDEF12345678"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_realm_ad_ldap.test", "directory_server_configurations.0.use_routing_to_select_interface", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_realm_ad_ldap.test", "directory_server_configurations.0.interface_group_id", "1234-5678-90AB-CDEF12345678"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcRealmADLDAPConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccNamedDataSourceFmcRealmADLDAPConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcRealmADLDAPConfig() string {
	config := `resource "fmc_realm_ad_ldap" "test" {` + "\n"
	config += `	name = "my_ldap_realm"` + "\n"
	config += `	description = "My realm"` + "\n"
	config += `	realm_type = "LDAP"` + "\n"
	config += `	ad_primary_domain = "example.com"` + "\n"
	config += `	ad_join_username = "user@example.com"` + "\n"
	config += `	ad_join_password = "my_password"` + "\n"
	config += `	directory_username = "user@example.com"` + "\n"
	config += `	directory_password = "my_password"` + "\n"
	config += `	base_dn = "DC=example,DC=com"` + "\n"
	config += `	group_dn = "CN=users,DC=example,DC=com"` + "\n"
	config += `	update_hour = ` + "\n"
	config += `	update_interval = ""` + "\n"
	config += `	group_attribute = "member"` + "\n"
	config += `	timeout_ise_users = 1440` + "\n"
	config += `	timeout_terminal_server_agent_users = 1440` + "\n"
	config += `	timeout_captive_portal_users = 1440` + "\n"
	config += `	timeout_failed_captive_portal_users = 1440` + "\n"
	config += `	timeout_guest_captive_portal_users = 1440` + "\n"
	config += `	directory_server_configurations = [{` + "\n"
	config += `		hostname = "ldap.example.com"` + "\n"
	config += `		port = 389` + "\n"
	config += `		encryption_protocol = "LDAPS"` + "\n"
	config += `		encryption_certificate = "1234-5678-90AB-CDEF12345678"` + "\n"
	config += `		use_routing_to_select_interface = true` + "\n"
	config += `		interface_group_id = "1234-5678-90AB-CDEF12345678"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_realm_ad_ldap" "test" {
			id = fmc_realm_ad_ldap.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcRealmADLDAPConfig() string {
	config := `resource "fmc_realm_ad_ldap" "test" {` + "\n"
	config += `	name = "my_ldap_realm"` + "\n"
	config += `	description = "My realm"` + "\n"
	config += `	realm_type = "LDAP"` + "\n"
	config += `	ad_primary_domain = "example.com"` + "\n"
	config += `	ad_join_username = "user@example.com"` + "\n"
	config += `	ad_join_password = "my_password"` + "\n"
	config += `	directory_username = "user@example.com"` + "\n"
	config += `	directory_password = "my_password"` + "\n"
	config += `	base_dn = "DC=example,DC=com"` + "\n"
	config += `	group_dn = "CN=users,DC=example,DC=com"` + "\n"
	config += `	update_hour = ` + "\n"
	config += `	update_interval = ""` + "\n"
	config += `	group_attribute = "member"` + "\n"
	config += `	timeout_ise_users = 1440` + "\n"
	config += `	timeout_terminal_server_agent_users = 1440` + "\n"
	config += `	timeout_captive_portal_users = 1440` + "\n"
	config += `	timeout_failed_captive_portal_users = 1440` + "\n"
	config += `	timeout_guest_captive_portal_users = 1440` + "\n"
	config += `	directory_server_configurations = [{` + "\n"
	config += `		hostname = "ldap.example.com"` + "\n"
	config += `		port = 389` + "\n"
	config += `		encryption_protocol = "LDAPS"` + "\n"
	config += `		encryption_certificate = "1234-5678-90AB-CDEF12345678"` + "\n"
	config += `		use_routing_to_select_interface = true` + "\n"
	config += `		interface_group_id = "1234-5678-90AB-CDEF12345678"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_realm_ad_ldap" "test" {
			name = fmc_realm_ad_ldap.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
