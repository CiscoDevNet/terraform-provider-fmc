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

func TestAccDataSourceFmcASPath(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_as_path.test", "name", "100"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_as_path.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_as_path.test", "overridable", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_as_path.test", "entries.0.action", "PERMIT"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_as_path.test", "entries.0.regular_expression", "^(100|200)$"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcASPathConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccNamedDataSourceFmcASPathConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcASPathConfig() string {
	config := `resource "fmc_as_path" "test" {` + "\n"
	config += `	name = 100` + "\n"
	config += `	overridable = false` + "\n"
	config += `	entries = [{` + "\n"
	config += `		action = "PERMIT"` + "\n"
	config += `		regular_expression = "^(100|200)$"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_as_path" "test" {
			id = fmc_as_path.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcASPathConfig() string {
	config := `resource "fmc_as_path" "test" {` + "\n"
	config += `	name = 100` + "\n"
	config += `	overridable = false` + "\n"
	config += `	entries = [{` + "\n"
	config += `		action = "PERMIT"` + "\n"
	config += `		regular_expression = "^(100|200)$"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_as_path" "test" {
			name = fmc_as_path.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
