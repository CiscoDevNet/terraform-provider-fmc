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

func TestAccDataSourceFmcBFDTemplates(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_bfd_templates.test", "items.my_bfd_templates.id"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_bfd_templates.test", "items.my_bfd_templates.type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_bfd_templates.test", "items.my_bfd_templates.hop_type", "SINGLE_HOP"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_bfd_templates.test", "items.my_bfd_templates.echo", "ENABLED"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_bfd_templates.test", "items.my_bfd_templates.interval_type", "MILLISECONDS"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_bfd_templates.test", "items.my_bfd_templates.multiplier", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_bfd_templates.test", "items.my_bfd_templates.minimum_transmit", "300"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_bfd_templates.test", "items.my_bfd_templates.minimum_receive", "300"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_bfd_templates.test", "items.my_bfd_templates.authentication_type", "MD5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_bfd_templates.test", "items.my_bfd_templates.authentication_password_encryption", "UN_ENCRYPTED"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_bfd_templates.test", "items.my_bfd_templates.authentication_key_id", "1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcBFDTemplatesConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcBFDTemplatesConfig() string {
	config := `resource "fmc_bfd_templates" "test" {` + "\n"
	config += `	items = { "my_bfd_templates" = {` + "\n"
	config += `		hop_type = "SINGLE_HOP"` + "\n"
	config += `		echo = "ENABLED"` + "\n"
	config += `		interval_type = "MILLISECONDS"` + "\n"
	config += `		multiplier = 3` + "\n"
	config += `		minimum_transmit = 300` + "\n"
	config += `		minimum_receive = 300` + "\n"
	config += `		authentication_type = "MD5"` + "\n"
	config += `		authentication_password = "ThisIsMySecretPassword"` + "\n"
	config += `		authentication_password_encryption = "UN_ENCRYPTED"` + "\n"
	config += `		authentication_key_id = 1` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_bfd_templates" "test" {
			depends_on = [fmc_bfd_templates.test]
			items = {
				"my_bfd_templates" = {
				}
			}
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
