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

func TestAccFmcNetworks(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_networks.test", "items.my_networks.id"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_networks.test", "items.my_networks.description", "My Network 1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_networks.test", "items.my_networks.prefix", "10.1.1.0/24"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_networks.test", "items.my_networks.type"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcNetworksConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcNetworksConfig_all(),
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

func testAccFmcNetworksConfig_minimum() string {
	config := `resource "fmc_networks" "test" {` + "\n"
	config += `	items = { "my_networks" = {` + "\n"
	config += `		prefix = "10.1.1.0/24"` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcNetworksConfig_all() string {
	config := `resource "fmc_networks" "test" {` + "\n"
	config += `	items = { "my_networks" = {` + "\n"
	config += `		description = "My Network 1"` + "\n"
	config += `		overridable = true` + "\n"
	config += `		prefix = "10.1.1.0/24"` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

func TestAccFmcNetworks_Sequential(t *testing.T) {
	step_01 := `resource "fmc_networks" "test" {` + "\n" +
		`	items = {` + "\n" +
		`		"networks_1" = {` + "\n" +
		`			prefix = "10.20.30.0/24"` + "\n" +
		`			description = "network 1"` + "\n" +
		`			overridable = true` + "\n" +
		`		},` + "\n" +
		`		"networks_2" = {` + "\n" +
		`			prefix = "10.20.40.0/24"` + "\n" +
		`		},` + "\n" +
		`		"networks_3" = {` + "\n" +
		`			prefix = "10.20.50.0/24"` + "\n" +
		`		},` + "\n" +
		`	} ` + "\n" +
		`}` + "\n"

	var checks_step01 []resource.TestCheckFunc
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_networks.test", "items.networks_1.prefix", "10.20.30.0/24"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_networks.test", "items.networks_1.description", "network 1"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_networks.test", "items.networks_2.prefix", "10.20.40.0/24"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_networks.test", "items.networks_3.prefix", "10.20.50.0/24"))

	step_02 := `resource "fmc_networks" "test" {` + "\n" +
		`	items = {` + "\n" +
		`		"networks_1" = {` + "\n" +
		`			prefix = "10.20.30.0/24"` + "\n" +
		`			description = "network 1 new description"` + "\n" +
		`		},` + "\n" +
		`		"networks_2" = {` + "\n" +
		`			prefix = "10.20.40.0/24"` + "\n" +
		`		},` + "\n" +
		`		"networks_4" = {` + "\n" +
		`			prefix = "10.20.60.0/24"` + "\n" +
		`		},` + "\n" +
		`	} ` + "\n" +
		`}` + "\n"

	var checks_step02 []resource.TestCheckFunc
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttr("fmc_networks.test", "items.networks_1.prefix", "10.20.30.0/24"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttr("fmc_networks.test", "items.networks_1.description", "network 1 new description"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttr("fmc_networks.test", "items.networks_2.prefix", "10.20.40.0/24"))
	checks_step02 = append(checks_step02, resource.TestCheckNoResourceAttr("fmc_networks.test", "items.networks_3"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttr("fmc_networks.test", "items.networks_4.prefix", "10.20.60.0/24"))

	steps := []resource.TestStep{{
		Config: step_01,
		Check:  resource.ComposeTestCheckFunc(checks_step01...),
	}, {
		Config: step_02,
		Check:  resource.ComposeTestCheckFunc(checks_step02...),
	}}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
