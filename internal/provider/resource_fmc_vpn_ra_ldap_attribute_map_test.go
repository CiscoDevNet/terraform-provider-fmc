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

func TestAccFmcVPNRALDAPAttributeMap(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_vpn_ra_ldap_attribute_map.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_ldap_attribute_map.test", "realms.0.realm_ad_ldap_id", "76d24097-41c4-4558-a4d0-a8c07ac08470"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_ldap_attribute_map.test", "realms.0.attribute_maps.0.ldap_attribute_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_ldap_attribute_map.test", "realms.0.attribute_maps.0.cisco_attribute_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_ldap_attribute_map.test", "realms.0.attribute_maps.0.value_maps.0.ldap_value", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_ldap_attribute_map.test", "realms.0.attribute_maps.0.value_maps.0.cisco_value", ""))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcVPNRALDAPAttributeMapConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcVPNRALDAPAttributeMapConfig_all(),
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
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcVPNRALDAPAttributeMapConfig_minimum() string {
	config := `resource "fmc_vpn_ra_ldap_attribute_map" "test" {` + "\n"
	config += `	vpn_ra_id = fmc_vpn_ra.test.id` + "\n"
	config += `	realms = [{` + "\n"
	config += `		realm_ad_ldap_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcVPNRALDAPAttributeMapConfig_all() string {
	config := `resource "fmc_vpn_ra_ldap_attribute_map" "test" {` + "\n"
	config += `	vpn_ra_id = fmc_vpn_ra.test.id` + "\n"
	config += `	realms = [{` + "\n"
	config += `		realm_ad_ldap_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `		attribute_maps = [{` + "\n"
	config += `			ldap_attribute_name = ""` + "\n"
	config += `			cisco_attribute_name = ""` + "\n"
	config += `			value_maps = [{` + "\n"
	config += `				ldap_value = ""` + "\n"
	config += `				cisco_value = ""` + "\n"
	config += `			}]` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
