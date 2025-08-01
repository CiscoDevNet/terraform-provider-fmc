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

func TestAccDataSourceFmcVPNRACertificateMap(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_vpn_ra_certificate_map.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_ra_certificate_map.test", "use_group_url", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_ra_certificate_map.test", "use_certificate_to_connection_profile_mappings", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_ra_certificate_map.test", "certificate_to_connection_profile_mappings.0.certificate_map_id", "76d24097-41c4-4558-a4d0-a8c07ac08470"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_ra_certificate_map.test", "certificate_to_connection_profile_mappings.0.connection_profile_id", "76d24097-41c4-4558-a4d0-a8c07ac08470"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcVPNRACertificateMapConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcVPNRACertificateMapConfig() string {
	config := `resource "fmc_vpn_ra_certificate_map" "test" {` + "\n"
	config += `	vpn_ra_id = fmc_vpn_ra.test.id` + "\n"
	config += `	use_group_url = false` + "\n"
	config += `	use_certificate_to_connection_profile_mappings = false` + "\n"
	config += `	certificate_to_connection_profile_mappings = [{` + "\n"
	config += `		certificate_map_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `		connection_profile_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_vpn_ra_certificate_map" "test" {
			id = fmc_vpn_ra_certificate_map.test.id
			vpn_ra_id = fmc_vpn_ra.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
