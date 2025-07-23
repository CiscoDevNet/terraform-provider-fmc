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

func TestAccFmcRealmADLDAP(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "name", "my_ldap_realm"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_realm_ad_ldap.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_realm_ad_ldap.test", "version"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "description", "My realm"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "realm_type", "LDAP"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "directory_username", "user@example.com"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "base_dn", "DC=example,DC=com"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "group_dn", "CN=users,DC=example,DC=com"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "included_users.0", "user1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "included_groups.0", "group1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "excluded_users.0", "user2"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "excluded_groups.0", "group2"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "update_hour", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "update_interval", "4"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "group_attribute", "member"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "timeout_ise_and_passive_indentity_users", "1440"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "timeout_terminal_server_agent_users", "1440"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "timeout_captive_portal_users", "1440"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "timeout_failed_captive_portal_users", "1440"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "timeout_guest_captive_portal_users", "1440"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "directory_server_configurations.0.hostname", "ldap.example.com"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "directory_server_configurations.0.port", "389"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "directory_server_configurations.0.use_routing_to_select_interface", "false"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcRealmADLDAPConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcRealmADLDAPConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_realm_ad_ldap.test",
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
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcRealmADLDAPConfig_minimum() string {
	config := `resource "fmc_realm_ad_ldap" "test" {` + "\n"
	config += `	name = "my_ldap_realm"` + "\n"
	config += `	realm_type = "LDAP"` + "\n"
	config += `	directory_username = "user@example.com"` + "\n"
	config += `	directory_password = "my_password"` + "\n"
	config += `	base_dn = "DC=example,DC=com"` + "\n"
	config += `	group_dn = "CN=users,DC=example,DC=com"` + "\n"
	config += `	directory_server_configurations = [{` + "\n"
	config += `		hostname = "ldap.example.com"` + "\n"
	config += `		port = 389` + "\n"
	config += `		encryption_protocol = "NONE"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcRealmADLDAPConfig_all() string {
	config := `resource "fmc_realm_ad_ldap" "test" {` + "\n"
	config += `	name = "my_ldap_realm"` + "\n"
	config += `	enabled = true` + "\n"
	config += `	description = "My realm"` + "\n"
	config += `	realm_type = "LDAP"` + "\n"
	config += `	directory_username = "user@example.com"` + "\n"
	config += `	directory_password = "my_password"` + "\n"
	config += `	base_dn = "DC=example,DC=com"` + "\n"
	config += `	group_dn = "CN=users,DC=example,DC=com"` + "\n"
	config += `	included_users = ["user1"]` + "\n"
	config += `	included_groups = ["group1"]` + "\n"
	config += `	excluded_users = ["user2"]` + "\n"
	config += `	excluded_groups = ["group2"]` + "\n"
	config += `	update_hour = 2` + "\n"
	config += `	update_interval = "4"` + "\n"
	config += `	group_attribute = "member"` + "\n"
	config += `	timeout_ise_and_passive_indentity_users = 1440` + "\n"
	config += `	timeout_terminal_server_agent_users = 1440` + "\n"
	config += `	timeout_captive_portal_users = 1440` + "\n"
	config += `	timeout_failed_captive_portal_users = 1440` + "\n"
	config += `	timeout_guest_captive_portal_users = 1440` + "\n"
	config += `	directory_server_configurations = [{` + "\n"
	config += `		hostname = "ldap.example.com"` + "\n"
	config += `		port = 389` + "\n"
	config += `		encryption_protocol = "NONE"` + "\n"
	config += `		use_routing_to_select_interface = false` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
