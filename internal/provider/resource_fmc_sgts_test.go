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

func TestAccFmcSGTs(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_sgts.test", "items.my_sgts.id"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_sgts.test", "items.my_sgts.type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_sgts.test", "items.my_sgts.description", "My SGT object"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_sgts.test", "items.my_sgts.tag", "11"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcSGTsConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcSGTsConfig_all(),
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

func testAccFmcSGTsConfig_minimum() string {
	config := `resource "fmc_sgts" "test" {` + "\n"
	config += `	items = { "my_sgts" = {` + "\n"
	config += `		tag = "11"` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcSGTsConfig_all() string {
	config := `resource "fmc_sgts" "test" {` + "\n"
	config += `	items = { "my_sgts" = {` + "\n"
	config += `		description = "My SGT object"` + "\n"
	config += `		tag = "11"` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

func TestAccFmcSGTs_Sequential(t *testing.T) {

	step_01 := `resource "fmc_sgts" "test" {` + "\n" +
		`	items = {` + "\n" +
		`		"sgt_1" = {` + "\n" +
		`			description = "my first tag"` + "\n" +
		`			tag = "100"` + "\n" +
		`		},` + "\n" +
		`		"sgt_2" = {` + "\n" +
		`			tag = "200",` + "\n" +
		`		},` + "\n" +
		`		"sgt_3" = {` + "\n" +
		`			description = "my third tag"` + "\n" +
		`			tag = "300",` + "\n" +
		`		},` + "\n" +
		`	} ` + "\n" +
		`}` + "\n"

	var checks_step01 []resource.TestCheckFunc
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttrSet("fmc_sgts.test", "items.sgt_1.id"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttrSet("fmc_sgts.test", "items.sgt_1.type"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_sgts.test", "items.sgt_1.description", "my first tag"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_sgts.test", "items.sgt_1.tag", "100"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttrSet("fmc_sgts.test", "items.sgt_2.id"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttrSet("fmc_sgts.test", "items.sgt_2.type"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_sgts.test", "items.sgt_2.tag", "200"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttrSet("fmc_sgts.test", "items.sgt_3.id"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttrSet("fmc_sgts.test", "items.sgt_3.type"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_sgts.test", "items.sgt_3.description", "my third tag"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_sgts.test", "items.sgt_3.tag", "300"))

	step_02 := `resource "fmc_sgts" "test" {` + "\n" +
		`	items = {` + "\n" +
		`		"sgt_1" = {` + "\n" +
		`			description = "my first tag"` + "\n" +
		`			tag = "100"` + "\n" +
		`		},` + "\n" +
		`		"sgt_2" = {` + "\n" +
		`			tag = "205",` + "\n" +
		`		},` + "\n" +
		`		"sgt_4" = {` + "\n" +
		`			tag = "400",` + "\n" +
		`		},` + "\n" +
		`	} ` + "\n" +
		`}` + "\n"

	var checks_step02 []resource.TestCheckFunc
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttrSet("fmc_sgts.test", "items.sgt_1.id"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttrSet("fmc_sgts.test", "items.sgt_1.type"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttr("fmc_sgts.test", "items.sgt_1.description", "my first tag"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttr("fmc_sgts.test", "items.sgt_1.tag", "100"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttrSet("fmc_sgts.test", "items.sgt_2.id"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttrSet("fmc_sgts.test", "items.sgt_2.type"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttr("fmc_sgts.test", "items.sgt_2.tag", "205"))
	checks_step02 = append(checks_step02, resource.TestCheckNoResourceAttr("fmc_sgts.test", "items.sgt_3"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttrSet("fmc_sgts.test", "items.sgt_4.id"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttrSet("fmc_sgts.test", "items.sgt_4.type"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttr("fmc_sgts.test", "items.sgt_4.tag", "400"))

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
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps:                    steps,
	})
}
