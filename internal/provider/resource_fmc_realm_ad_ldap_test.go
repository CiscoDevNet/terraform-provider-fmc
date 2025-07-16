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
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_realm_ad_ldap.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "description", "My realm"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "realm_type", "LDAP"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "ad_primary_domain", "example.com"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "directory_username", "user@example.com"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "directory_password", "my_password"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "base_dn", "dc=example,dc=com"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "group_dn", "ou=users,dc=example,dc=com"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "directory_configurations.0.hostname", "ldap.example.com"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "directory_configurations.0.port", "389"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "directory_configurations.0.encryption", "LDAPS"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "directory_configurations.0.encryption_certificate", "1234-5678-90AB-CDEF12345678"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "directory_configurations.0.use_routing_to_select_interface", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_realm_ad_ldap.test", "directory_configurations.0.interface_group_id", "1234-5678-90AB-CDEF12345678"))

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
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcRealmADLDAPConfig_all() string {
	config := `resource "fmc_realm_ad_ldap" "test" {` + "\n"
	config += `	name = "my_ldap_realm"` + "\n"
	config += `	description = "My realm"` + "\n"
	config += `	realm_type = "LDAP"` + "\n"
	config += `	ad_primary_domain = "example.com"` + "\n"
	config += `	directory_username = "user@example.com"` + "\n"
	config += `	directory_password = "my_password"` + "\n"
	config += `	base_dn = "dc=example,dc=com"` + "\n"
	config += `	group_dn = "ou=users,dc=example,dc=com"` + "\n"
	config += `	directory_configurations = [{` + "\n"
	config += `		hostname = "ldap.example.com"` + "\n"
	config += `		port = 389` + "\n"
	config += `		encryption = "LDAPS"` + "\n"
	config += `		encryption_certificate = "1234-5678-90AB-CDEF12345678"` + "\n"
	config += `		use_routing_to_select_interface = true` + "\n"
	config += `		interface_group_id = "1234-5678-90AB-CDEF12345678"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
