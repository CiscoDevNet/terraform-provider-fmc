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

func TestAccDataSourceFmcBGDTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_bgd_template.test", "name", "BFD_Template1"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_bgd_template.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_bgd_template.test", "hop_type", "SINGLE_HOP"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_bgd_template.test", "echo", "ENABLED"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_bgd_template.test", "interval_time", "MILLISECONDS"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_bgd_template.test", "min_transmit", "100000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_bgd_template.test", "tx_rx_multiplier", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_bgd_template.test", "min_receive", "100000"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcBGDTemplateConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccNamedDataSourceFmcBGDTemplateConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcBGDTemplateConfig() string {
	config := `resource "fmc_bgd_template" "test" {` + "\n"
	config += `	name = "BFD_Template1"` + "\n"
	config += `	hop_type = "SINGLE_HOP"` + "\n"
	config += `	echo = "ENABLED"` + "\n"
	config += `	interval_time = "MILLISECONDS"` + "\n"
	config += `	min_transmit = 100000` + "\n"
	config += `	tx_rx_multiplier = 3` + "\n"
	config += `	min_receive = 100000` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_bgd_template" "test" {
			id = fmc_bgd_template.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcBGDTemplateConfig() string {
	config := `resource "fmc_bgd_template" "test" {` + "\n"
	config += `	name = "BFD_Template1"` + "\n"
	config += `	hop_type = "SINGLE_HOP"` + "\n"
	config += `	echo = "ENABLED"` + "\n"
	config += `	interval_time = "MILLISECONDS"` + "\n"
	config += `	min_transmit = 100000` + "\n"
	config += `	tx_rx_multiplier = 3` + "\n"
	config += `	min_receive = 100000` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_bgd_template" "test" {
			name = fmc_bgd_template.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
