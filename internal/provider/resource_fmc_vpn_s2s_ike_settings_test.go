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

func TestAccFmcVPNS2SIKESettings(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_vpn_s2s_ike_settings.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_ike_settings.test", "ikev1_authentication_type", "MANUAL_PRE_SHARED_KEY"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_ike_settings.test", "ikev2_authentication_type", "MANUAL_PRE_SHARED_KEY"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_ike_settings.test", "ikev2_enforce_hex_based_pre_shared_key", "false"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcVPNS2SIKESettingsPrerequisitesConfig + testAccFmcVPNS2SIKESettingsConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcVPNS2SIKESettingsPrerequisitesConfig + testAccFmcVPNS2SIKESettingsConfig_all(),
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

const testAccFmcVPNS2SIKESettingsPrerequisitesConfig = `
resource "fmc_vpn_s2s" "test" {
  name             = "my_s2s_vpn_ike_settings"
  route_based      = false
  network_topology = "POINT_TO_POINT"
  ikev1            = true
  ikev2            = true
}

resource "fmc_ikev2_policies" "test" {
  items = {
    my_s2s_vpn_ike_settings = {
      description           = "IKEv2 Policy"
      priority              = 10
      lifetime              = 86400
      integrity_algorithms  = ["SHA-256"]
      encryption_algorithms = ["AES-256"]
      prf_algorithms        = ["SHA-256"]
      dh_groups             = ["14"]
    }
  }
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcVPNS2SIKESettingsConfig_minimum() string {
	config := `resource "fmc_vpn_s2s_ike_settings" "test" {` + "\n"
	config += `	vpn_s2s_id = fmc_vpn_s2s.test.id` + "\n"
	config += `	ikev1_authentication_type = "MANUAL_PRE_SHARED_KEY"` + "\n"
	config += `	ikev1_manual_pre_shared_key = "my_pre_shared_key123"` + "\n"
	config += `	ikev2_authentication_type = "MANUAL_PRE_SHARED_KEY"` + "\n"
	config += `	ikev2_manual_pre_shared_key = "my_pre_shared_key123"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcVPNS2SIKESettingsConfig_all() string {
	config := `resource "fmc_vpn_s2s_ike_settings" "test" {` + "\n"
	config += `	vpn_s2s_id = fmc_vpn_s2s.test.id` + "\n"
	config += `	ikev1_authentication_type = "MANUAL_PRE_SHARED_KEY"` + "\n"
	config += `	ikev1_manual_pre_shared_key = "my_pre_shared_key123"` + "\n"
	config += `	ikev2_authentication_type = "MANUAL_PRE_SHARED_KEY"` + "\n"
	config += `	ikev2_manual_pre_shared_key = "my_pre_shared_key123"` + "\n"
	config += `	ikev2_enforce_hex_based_pre_shared_key = false` + "\n"
	config += `	ikev2_policies = [{` + "\n"
	config += `		id = fmc_ikev2_policies.test.items["my_s2s_vpn_ike_settings"].id` + "\n"
	config += `		name = "my_s2s_vpn_ike_settings"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
