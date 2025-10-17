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

func TestAccFmcPorts(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_ports.test", "items.my_ports.id"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_ports.test", "items.my_ports.type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ports.test", "items.my_ports.protocol", "TCP"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ports.test", "items.my_ports.port", "443"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ports.test", "items.my_ports.description", "Port TCP/443 (HTTPS)"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcPortsConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcPortsConfig_all(),
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

func testAccFmcPortsConfig_minimum() string {
	config := `resource "fmc_ports" "test" {` + "\n"
	config += `	items = { "my_ports" = {` + "\n"
	config += `		protocol = "TCP"` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcPortsConfig_all() string {
	config := `resource "fmc_ports" "test" {` + "\n"
	config += `	items = { "my_ports" = {` + "\n"
	config += `		protocol = "TCP"` + "\n"
	config += `		port = "443"` + "\n"
	config += `		description = "Port TCP/443 (HTTPS)"` + "\n"
	config += `		overridable = true` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
func TestAccFmcPorts_Sequential(t *testing.T) {
	step_01 := `resource "fmc_ports" "test" {` + "\n" +
		`	items = {` + "\n" +
		`		"ports_1" = {` + "\n" +
		`			port = "1234",` + "\n" +
		`			protocol = "TCP",` + "\n" +
		`			description = "port 1"` + "\n" +
		`			overridable = true` + "\n" +
		`		},` + "\n" +
		`		"ports_2" = {` + "\n" +
		`			port = "1235",` + "\n" +
		`			protocol = "TCP",` + "\n" +
		`		},` + "\n" +
		`		"ports_3" = {` + "\n" +
		`			port = "1236",` + "\n" +
		`			protocol = "TCP",` + "\n" +
		`		},` + "\n" +
		`	} ` + "\n" +
		`}` + "\n"

	var checks_step01 []resource.TestCheckFunc
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttrSet("fmc_ports.test", "items.ports_1.id"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_ports.test", "items.ports_1.port", "1234"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_ports.test", "items.ports_1.description", "port 1"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_ports.test", "items.ports_2.port", "1235"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_ports.test", "items.ports_3.port", "1236"))

	step_02 := `resource "fmc_ports" "test" {` + "\n" +
		`	items = {` + "\n" +
		`		"ports_1" = {` + "\n" +
		`			port = "1233",` + "\n" +
		`			protocol = "TCP",` + "\n" +
		`			description = "port 1 new description"` + "\n" +
		`		},` + "\n" +
		`		"ports_2" = {` + "\n" +
		`			port = "1235",` + "\n" +
		`			protocol = "TCP",` + "\n" +
		`		},` + "\n" +
		`		"ports_4" = {` + "\n" +
		`			port = "1239",` + "\n" +
		`			protocol = "TCP",` + "\n" +
		`		},` + "\n" +
		`	} ` + "\n" +
		`}` + "\n"

	var checks_step02 []resource.TestCheckFunc
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttrSet("fmc_ports.test", "items.ports_1.id"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttr("fmc_ports.test", "items.ports_1.port", "1233"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttr("fmc_ports.test", "items.ports_1.description", "port 1 new description"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttr("fmc_ports.test", "items.ports_2.port", "1235"))
	checks_step02 = append(checks_step02, resource.TestCheckNoResourceAttr("fmc_ports.test", "items.ports_3"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttr("fmc_ports.test", "items.ports_4.port", "1239"))

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
