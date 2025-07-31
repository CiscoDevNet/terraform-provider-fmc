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

func TestAccFmcVPNRA(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra.test", "name", "my_ftd_ra_vpn"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra.test", "description", "description of my_fmc_ra_vpn"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_vpn_ra.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra.test", "protocol_ssl", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra.test", "protocol_ipsec_ikev2", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra.test", "local_realm_id", "12345678-1234-1234-1234-123456789012"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra.test", "dap_policy_id", "12345678-1234-1234-1234-123456"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra.test", "access_interfaces.0.id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra.test", "access_interfaces.0.protocol_ipsec_ikev2", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra.test", "access_interfaces.0.protocol_ssl", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra.test", "access_interfaces.0.protocol_ssl_dtls", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra.test", "access_interfaces.0.interface_specific_certificate_id", "12345678-1234-1234-1234-123456"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra.test", "allow_users_to_select_connection_profile", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra.test", "web_port", "443"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra.test", "dtls_port", "443"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra.test", "ssl_global_identity_certificate_id", "12345678-1234-1234-1234-123456"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra.test", "ipsec_global_identity_certificate_id", "12345678-1234-1234-1234-123456"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra.test", "service_access_id", "12345678-1234-1234-1234-123456"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra.test", "bypass_access_control_policy_for_decrypted_traffic", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra.test", "secure_client_images.0.id", "12345678-1234-1234-1234-123456789012"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra.test", "secure_client_images.0.operating_system", "WINDOWS"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra.test", "external_browser_package_id", "12345678-1234-1234-1234-123456789012"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_vpn_ra.test", "secure_client_customization_id"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_vpn_ra.test", "address_assignment_policy_id"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_vpn_ra.test", "certificate_map_id"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra.test", "group_policies.0.id", "12345678-1234-1234-1234-123456"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_vpn_ra.test", "ldap_attribute_map_id"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_vpn_ra.test", "load_balance_id"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra.test", "ikev2_policies.0.id", "12345678-1234-1234-1234-123456789012"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_vpn_ra.test", "ipsec_advanced_settings_id"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcVPNRAConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcVPNRAConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_vpn_ra.test",
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

func testAccFmcVPNRAConfig_minimum() string {
	config := `resource "fmc_vpn_ra" "test" {` + "\n"
	config += `	name = "my_ftd_ra_vpn"` + "\n"
	config += `	group_policies = [{` + "\n"
	config += `		id = "12345678-1234-1234-1234-123456"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcVPNRAConfig_all() string {
	config := `resource "fmc_vpn_ra" "test" {` + "\n"
	config += `	name = "my_ftd_ra_vpn"` + "\n"
	config += `	description = "description of my_fmc_ra_vpn"` + "\n"
	config += `	protocol_ssl = true` + "\n"
	config += `	protocol_ipsec_ikev2 = true` + "\n"
	config += `	local_realm_id = "12345678-1234-1234-1234-123456789012"` + "\n"
	config += `	dap_policy_id = "12345678-1234-1234-1234-123456"` + "\n"
	config += `	access_interfaces = [{` + "\n"
	config += `		id = ""` + "\n"
	config += `		protocol_ipsec_ikev2 = true` + "\n"
	config += `		protocol_ssl = true` + "\n"
	config += `		protocol_ssl_dtls = true` + "\n"
	config += `		interface_specific_certificate_id = "12345678-1234-1234-1234-123456"` + "\n"
	config += `	}]` + "\n"
	config += `	allow_users_to_select_connection_profile = true` + "\n"
	config += `	web_port = 443` + "\n"
	config += `	dtls_port = 443` + "\n"
	config += `	ssl_global_identity_certificate_id = "12345678-1234-1234-1234-123456"` + "\n"
	config += `	ipsec_global_identity_certificate_id = "12345678-1234-1234-1234-123456"` + "\n"
	config += `	service_access_id = "12345678-1234-1234-1234-123456"` + "\n"
	config += `	bypass_access_control_policy_for_decrypted_traffic = true` + "\n"
	config += `	secure_client_images = [{` + "\n"
	config += `		id = "12345678-1234-1234-1234-123456789012"` + "\n"
	config += `		operating_system = "WINDOWS"` + "\n"
	config += `	}]` + "\n"
	config += `	external_browser_package_id = "12345678-1234-1234-1234-123456789012"` + "\n"
	config += `	group_policies = [{` + "\n"
	config += `		id = "12345678-1234-1234-1234-123456"` + "\n"
	config += `	}]` + "\n"
	config += `	ikev2_policies = [{` + "\n"
	config += `		id = "12345678-1234-1234-1234-123456789012"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
