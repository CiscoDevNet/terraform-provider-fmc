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

func TestAccFmcBFDTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_bfd_template.test", "name", "BFD_Template1"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_bfd_template.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_bfd_template.test", "hop_type", "SINGLE_HOP"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_bfd_template.test", "echo", "ENABLED"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_bfd_template.test", "interval_time", "MILLISECONDS"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_bfd_template.test", "min_transmit", "300"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_bfd_template.test", "tx_rx_multiplier", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_bfd_template.test", "min_receive", "300"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_bfd_template.test", "authentication_password", "Cisco123!"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_bfd_template.test", "authentication_key_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_bfd_template.test", "authentication_type", "MD5"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcBFDTemplateConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcBFDTemplateConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_bfd_template.test",
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

func testAccFmcBFDTemplateConfig_minimum() string {
	config := `resource "fmc_bfd_template" "test" {` + "\n"
	config += `	name = "BFD_Template1"` + "\n"
	config += `	hop_type = "SINGLE_HOP"` + "\n"
	config += `	echo = "ENABLED"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcBFDTemplateConfig_all() string {
	config := `resource "fmc_bfd_template" "test" {` + "\n"
	config += `	name = "BFD_Template1"` + "\n"
	config += `	hop_type = "SINGLE_HOP"` + "\n"
	config += `	echo = "ENABLED"` + "\n"
	config += `	interval_time = "MILLISECONDS"` + "\n"
	config += `	min_transmit = 300` + "\n"
	config += `	tx_rx_multiplier = 3` + "\n"
	config += `	min_receive = 300` + "\n"
	config += `	authentication_password = "Cisco123!"` + "\n"
	config += `	authentication_key_id = 1` + "\n"
	config += `	authentication_type = "MD5"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
