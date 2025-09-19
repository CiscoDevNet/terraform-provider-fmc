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

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccFmcFTDPlatformSettingsSyslogRateLimit(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_ftd_platform_settings_syslog_rate_limit.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_syslog_rate_limit.test", "rate_limit_type", "LOG_LEVEL"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_syslog_rate_limit.test", "rate_limit_value", "ERR"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_syslog_rate_limit.test", "number_of_messages", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ftd_platform_settings_syslog_rate_limit.test", "interval", "60"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccFmcFTDPlatformSettingsSyslogRateLimitPrerequisitesConfig + testAccFmcFTDPlatformSettingsSyslogRateLimitConfig_all(),
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

const testAccFmcFTDPlatformSettingsSyslogRateLimitPrerequisitesConfig = `
resource "fmc_ftd_platform_settings" "test" {
  name        = "ftd_platform_settings_rate_limit"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcFTDPlatformSettingsSyslogRateLimitConfig_all() string {
	config := `resource "fmc_ftd_platform_settings_syslog_rate_limit" "test" {` + "\n"
	config += `	ftd_platform_settings_id = fmc_ftd_platform_settings.test.id` + "\n"
	config += `	rate_limit_type = "LOG_LEVEL"` + "\n"
	config += `	rate_limit_value = "ERR"` + "\n"
	config += `	number_of_messages = 100` + "\n"
	config += `	interval = 60` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
