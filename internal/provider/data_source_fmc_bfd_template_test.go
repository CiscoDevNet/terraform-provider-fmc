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

func TestAccDataSourceFmcBFDTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_bfd_template.test", "name", "fmc_bfd_template"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_bfd_template.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_bfd_template.test", "hop_type", "SINGLE_HOP"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_bfd_template.test", "echo", "ENABLED"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_bfd_template.test", "interval_time", "MILLISECONDS"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_bfd_template.test", "min_transmit", "300"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_bfd_template.test", "tx_rx_multiplier", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_bfd_template.test", "min_receive", "300"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_bfd_template.test", "authentication_key_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_bfd_template.test", "authentication_type", "MD5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_bfd_template.test", "authentication_password_encryption", "UN_ENCRYPTED"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcBFDTemplateConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccNamedDataSourceFmcBFDTemplateConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcBFDTemplateConfig() string {
	config := `resource "fmc_bfd_template" "test" {` + "\n"
	config += `	name = "fmc_bfd_template"` + "\n"
	config += `	hop_type = "SINGLE_HOP"` + "\n"
	config += `	echo = "ENABLED"` + "\n"
	config += `	interval_time = "MILLISECONDS"` + "\n"
	config += `	min_transmit = 300` + "\n"
	config += `	tx_rx_multiplier = 3` + "\n"
	config += `	min_receive = 300` + "\n"
	config += `	authentication_password = "ThisIsMySecretPassword"` + "\n"
	config += `	authentication_key_id = 1` + "\n"
	config += `	authentication_type = "MD5"` + "\n"
	config += `	authentication_password_encryption = "UN_ENCRYPTED"` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_bfd_template" "test" {
			id = fmc_bfd_template.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcBFDTemplateConfig() string {
	config := `resource "fmc_bfd_template" "test" {` + "\n"
	config += `	name = "fmc_bfd_template"` + "\n"
	config += `	hop_type = "SINGLE_HOP"` + "\n"
	config += `	echo = "ENABLED"` + "\n"
	config += `	interval_time = "MILLISECONDS"` + "\n"
	config += `	min_transmit = 300` + "\n"
	config += `	tx_rx_multiplier = 3` + "\n"
	config += `	min_receive = 300` + "\n"
	config += `	authentication_password = "ThisIsMySecretPassword"` + "\n"
	config += `	authentication_key_id = 1` + "\n"
	config += `	authentication_type = "MD5"` + "\n"
	config += `	authentication_password_encryption = "UN_ENCRYPTED"` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_bfd_template" "test" {
			name = fmc_bfd_template.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
