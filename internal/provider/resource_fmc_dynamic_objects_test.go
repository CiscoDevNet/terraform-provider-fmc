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

func TestAccFmcDynamicObjects(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_dynamic_objects.test", "items.dynamic_object_1.id"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_dynamic_objects.test", "items.dynamic_object_1.description", "My Dynamic Object 1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_dynamic_objects.test", "items.dynamic_object_1.object_type", "IP"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcDynamicObjectsConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcDynamicObjectsConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcDynamicObjectsConfig_minimum() string {
	config := `resource "fmc_dynamic_objects" "test" {` + "\n"
	config += `	items = { "dynamic_object_1" = {` + "\n"
	config += `		object_type = "IP"` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcDynamicObjectsConfig_all() string {
	config := `resource "fmc_dynamic_objects" "test" {` + "\n"
	config += `	items = { "dynamic_object_1" = {` + "\n"
	config += `		description = "My Dynamic Object 1"` + "\n"
	config += `		object_type = "IP"` + "\n"
	config += `		mappings = ["10.0.0.1"]` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

func TestAccFmcDynamicObjects_Sequential(t *testing.T) {

	step_01 := `resource "fmc_dynamic_objects" "test" {` + "\n" +
		`	items = {` + "\n" +
		`		"dynamic_object_1" = {` + "\n" +
		`			object_type = "IP",` + "\n" +
		`			description = "Dynamic Object 1"` + "\n" +
		`			mappings    = ["10.0.0.1","10.0.0.2","10.0.0.3"]` + "\n" +
		`		},` + "\n" +
		`		"dynamic_object_2" = {` + "\n" +
		`			object_type = "IP",` + "\n" +
		`			description = "Dynamic Object 2"` + "\n" +
		`			mappings    = ["10.1.0.1","10.1.0.2"]` + "\n" +
		`		},` + "\n" +
		`		"dynamic_object_3" = {` + "\n" +
		`			object_type = "IP",` + "\n" +
		`			mappings    = []` + "\n" +
		`		},` + "\n" +
		`		"dynamic_object_4" = {` + "\n" +
		`			object_type = "IP",` + "\n" +
		`			description = "Dynamic Object 4"` + "\n" +
		`		},` + "\n" +
		`		"dynamic_object_5" = {` + "\n" +
		`			object_type = "IP",` + "\n" +
		`			description = "Dynamic Object 5"` + "\n" +
		`			mappings    = ["10.5.0.1","10.5.0.2","10.5.0.3"]` + "\n" +
		`		},` + "\n" +
		`		"dynamic_object_6" = {` + "\n" +
		`			object_type = "IP",` + "\n" +
		`			description = "Dynamic Object 6"` + "\n" +
		`			mappings    = ["10.6.0.1","10.6.0.2","10.6.0.3"]` + "\n" +
		`		},` + "\n" +
		`	} ` + "\n" +
		`}` + "\n"

	step_02 := `resource "fmc_dynamic_objects" "test" {` + "\n" +
		`	items = {` + "\n" +
		`		"dynamic_object_1" = {` + "\n" +
		`			object_type = "IP",` + "\n" +
		`			description = "Dynamic Object 1"` + "\n" +
		`			mappings    = ["10.0.0.1","10.0.0.2","10.0.0.3"]` + "\n" +
		`		},` + "\n" +
		`		"dynamic_object_2" = {` + "\n" +
		`			object_type = "IP",` + "\n" +
		`			description = "Dynamic Object 2 new description"` + "\n" +
		`			mappings    = ["10.1.0.1","10.1.0.2","10.1.0.3"]` + "\n" +
		`		},` + "\n" +
		`		"dynamic_object_3" = {` + "\n" +
		`			object_type = "IP",` + "\n" +
		`			mappings    = ["10.1.0.1","10.1.0.2","10.1.0.3"]` + "\n" +
		`		},` + "\n" +
		`		"dynamic_object_4" = {` + "\n" +
		`			object_type = "IP",` + "\n" +
		`			description = "Dynamic Object 4"` + "\n" +
		`			mappings    = ["10.4.0.1","10.4.0.2","10.4.0.3"]` + "\n" +
		`		},` + "\n" +
		`		"dynamic_object_6" = {` + "\n" +
		`			object_type = "IP",` + "\n" +
		`			description = "Dynamic Object 6"` + "\n" +
		`			mappings    = ["10.6.0.1","10.6.0.2","10.6.0.4"]` + "\n" +
		`		},` + "\n" +
		`		"dynamic_object_7" = {` + "\n" +
		`			object_type = "IP",` + "\n" +
		`			description = "Dynamic Object 7"` + "\n" +
		`			mappings    = ["10.6.0.1","10.6.0.2","10.6.0.3"]` + "\n" +
		`		},` + "\n" +
		`	} ` + "\n" +
		`}` + "\n"

	steps := []resource.TestStep{{
		Config: step_01,
	}, {
		Config: step_02,
	}}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
