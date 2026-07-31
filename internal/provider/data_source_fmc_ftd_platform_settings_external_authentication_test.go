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

func TestAccDataSourceFmcFTDPlatformSettingsExternalAuthentication(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_ftd_platform_settings_external_authentication.test", "type"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcFTDPlatformSettingsExternalAuthenticationPrerequisitesConfig + testAccDataSourceFmcFTDPlatformSettingsExternalAuthenticationConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcFTDPlatformSettingsExternalAuthenticationPrerequisitesConfig = `
resource "fmc_ftd_platform_settings" "test" {
  name = "ftd_platform_settings_external_authentication"
}

resource "fmc_radius_external_authentication_object" "test" {
  name           = "ftd_platform_settings_external_authentication_radius"
  server_address = "10.1.1.10"
  key            = "my_secret_key"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcFTDPlatformSettingsExternalAuthenticationConfig() string {
	config := `resource "fmc_ftd_platform_settings_external_authentication" "test" {` + "\n"
	config += `	ftd_platform_settings_id = fmc_ftd_platform_settings.test.id` + "\n"
	config += `	external_auth_server_id = fmc_radius_external_authentication_object.test.id` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_ftd_platform_settings_external_authentication" "test" {
			id = fmc_ftd_platform_settings_external_authentication.test.id
			ftd_platform_settings_id = fmc_ftd_platform_settings.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
