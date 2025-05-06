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

func TestAccDataSourceFmcVPNS2SIKESettings(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_vpn_s2s_ike_settings.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_ike_settings.test", "ikev1_authentication_type", "MANUAL_PRE_SHARED_KEY"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_ike_settings.test", "ikev1_automatic_pre_shared_key_length", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_ike_settings.test", "ikev1_policies.0.name", "my_ikev1_policy"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_ike_settings.test", "ikev2_authentication_type", "MANUAL_PRE_SHARED_KEY"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_ike_settings.test", "ikev2_automatic_pre_shared_key_length", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_ike_settings.test", "ikev2_enforce_hex_based_pre_shared_key_only", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_ike_settings.test", "ikev2_policies.0.name", "my_ikev2_policy"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcVPNS2SIKESettingsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcVPNS2SIKESettingsConfig() string {
	config := `resource "fmc_vpn_s2s_ike_settings" "test" {` + "\n"
	config += `	vpn_s2s_id = TBD` + "\n"
	config += `	ikev1_authentication_type = "MANUAL_PRE_SHARED_KEY"` + "\n"
	config += `	ikev1_automatic_pre_shared_key_length = ` + "\n"
	config += `	ikev1_manual_pre_shared_key = "my_pre_shared_key123"` + "\n"
	config += `	ikev1_policies = [{` + "\n"
	config += `		id = TBD` + "\n"
	config += `		name = "my_ikev1_policy"` + "\n"
	config += `	}]` + "\n"
	config += `	ikev2_authentication_type = "MANUAL_PRE_SHARED_KEY"` + "\n"
	config += `	ikev2_automatic_pre_shared_key_length = ` + "\n"
	config += `	ikev2_manual_pre_shared_key = "my_pre_shared_key123"` + "\n"
	config += `	ikev2_enforce_hex_based_pre_shared_key_only = false` + "\n"
	config += `	ikev2_policies = [{` + "\n"
	config += `		id = TBD` + "\n"
	config += `		name = "my_ikev2_policy"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_vpn_s2s_ike_settings" "test" {
			id = fmc_vpn_s2s_ike_settings.test.id
			vpn_s2s_id = TBD
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
