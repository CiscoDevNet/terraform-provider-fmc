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

func TestAccFmcVPNRAConnectionProfiles(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.id"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.group_policy_id", "12345678-1234-1234-1234-123456"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.ipv4_address_pools.0.id", "12345678-1234-1234-1234-123456"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.ipv6_address_pools.0.id", "12345678-1234-1234-1234-123456"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.dhcp_servers.0.id", "12345678-1234-1234-1234-123456"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.authentication_method", "AAA_ONLY"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.primary_authentication_server_use_local", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.primary_authentication_server_id", "12345678-1234-1234-1234-123456"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.primary_authentication_server_type", "TBD"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.primary_authentication_fallback_to_local", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.multiple_certificate_authentication", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.saml_and_certificate_username_must_match", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.primary_authentication_prefill_username_from_certificate", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.primary_authentication_prefill_username_from_certificate_map_primary_field", "CN_COMMMON_NAME"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.primary_authentication_prefill_username_from_certificate_map_secondary_field", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.primary_authentication_prefill_username_from_certificate_map_entire_dn", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.primary_authentication_hide_username_in_login_window", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.secondary_authentication", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.secondary_authentication_server_use_local", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.secondary_authentication_server_id", "12345678-1234-1234-1234-123456"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.secondary_authentication_server_type", "TBD"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.secondary_authentication_fallback_to_local", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.secondary_authentication_prompt_for_username", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.secondary_authentication_use_primary_authentication_username", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.use_secondary_authentication_username_for_reporting", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.saml_use_external_browser", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.authorization_server_id", "12345678-1234-1234-1234-123456"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.authorization_server_type", "RadiusServerGroup"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.allow_connection_only_if_user_exists_in_authorization_database", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.accounting_server_id", "12345678-1234-1234-1234-123456"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.accounting_server_type", "RadiusServerGroup"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.strip_realm_from_username", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.strip_group_from_username", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.password_management", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.password_management_notify_user_on_password_expiry_day", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.password_management_advance_password_expiration_notification", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.alias_names.0.name", "my_group_alias"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.alias_names.0.enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.alias_urls.0.url_object_id", "12345678-1234-1234-1234-123456789012"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_connection_profiles.test", "items.my_connection_profile.alias_urls.0.enabled", "true"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccFmcVPNRAConnectionProfilesConfig_all(),
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
// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcVPNRAConnectionProfilesConfig_all() string {
	config := `resource "fmc_vpn_ra_connection_profiles" "test" {` + "\n"
	config += `	vpn_ra_id = fmc_vpn_ra.test.id` + "\n"
	config += `	items = { "my_connection_profile" = {` + "\n"
	config += `		group_policy_id = "12345678-1234-1234-1234-123456"` + "\n"
	config += `		ipv4_address_pools = [{` + "\n"
	config += `			id = "12345678-1234-1234-1234-123456"` + "\n"
	config += `		}]` + "\n"
	config += `		ipv6_address_pools = [{` + "\n"
	config += `			id = "12345678-1234-1234-1234-123456"` + "\n"
	config += `		}]` + "\n"
	config += `		dhcp_servers = [{` + "\n"
	config += `			id = "12345678-1234-1234-1234-123456"` + "\n"
	config += `		}]` + "\n"
	config += `		authentication_method = "AAA_ONLY"` + "\n"
	config += `		primary_authentication_server_use_local = false` + "\n"
	config += `		primary_authentication_server_id = "12345678-1234-1234-1234-123456"` + "\n"
	config += `		primary_authentication_server_type = "TBD"` + "\n"
	config += `		primary_authentication_fallback_to_local = true` + "\n"
	config += `		multiple_certificate_authentication = true` + "\n"
	config += `		saml_and_certificate_username_must_match = true` + "\n"
	config += `		primary_authentication_prefill_username_from_certificate = true` + "\n"
	config += `		primary_authentication_prefill_username_from_certificate_map_primary_field = "CN_COMMMON_NAME"` + "\n"
	config += `		primary_authentication_prefill_username_from_certificate_map_secondary_field = ""` + "\n"
	config += `		primary_authentication_prefill_username_from_certificate_map_entire_dn = true` + "\n"
	config += `		primary_authentication_hide_username_in_login_window = true` + "\n"
	config += `		secondary_authentication = true` + "\n"
	config += `		secondary_authentication_server_use_local = false` + "\n"
	config += `		secondary_authentication_server_id = "12345678-1234-1234-1234-123456"` + "\n"
	config += `		secondary_authentication_server_type = "TBD"` + "\n"
	config += `		secondary_authentication_fallback_to_local = true` + "\n"
	config += `		secondary_authentication_prompt_for_username = ` + "\n"
	config += `		secondary_authentication_use_primary_authentication_username = true` + "\n"
	config += `		use_secondary_authentication_username_for_reporting = ` + "\n"
	config += `		saml_use_external_browser = ` + "\n"
	config += `		authorization_server_id = "12345678-1234-1234-1234-123456"` + "\n"
	config += `		authorization_server_type = "RadiusServerGroup"` + "\n"
	config += `		allow_connection_only_if_user_exists_in_authorization_database = false` + "\n"
	config += `		accounting_server_id = "12345678-1234-1234-1234-123456"` + "\n"
	config += `		accounting_server_type = "RadiusServerGroup"` + "\n"
	config += `		strip_realm_from_username = true` + "\n"
	config += `		strip_group_from_username = true` + "\n"
	config += `		password_management = false` + "\n"
	config += `		password_management_notify_user_on_password_expiry_day = true` + "\n"
	config += `		password_management_advance_password_expiration_notification = ` + "\n"
	config += `		alias_names = [{` + "\n"
	config += `			name = "my_group_alias"` + "\n"
	config += `			enabled = true` + "\n"
	config += `		}]` + "\n"
	config += `		alias_urls = [{` + "\n"
	config += `			url_object_id = "12345678-1234-1234-1234-123456789012"` + "\n"
	config += `			enabled = true` + "\n"
	config += `		}]` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
