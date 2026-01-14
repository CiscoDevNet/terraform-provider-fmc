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
	"slices"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource

func TestAccDataSourceFmcGeolocations(t *testing.T) {
	if v := os.Getenv("FMC_VERSION"); v != "" && slices.Contains([]string{"7.2"}, v) {
		t.Skip("skipping test for FMC version " + v)
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_geolocations.test", "items.my_geolocations.id"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_geolocations.test", "items.my_geolocations.type"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcGeolocationsPrerequisitesConfig + testAccDataSourceFmcGeolocationsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcGeolocationsPrerequisitesConfig = `
data "fmc_countries" "test" {
  items = {
    "Poland"  = {}
  }
}

data "fmc_continents" "test" {
  items = {
    "North America" = {}
  }
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcGeolocationsConfig() string {
	config := `resource "fmc_geolocations" "test" {` + "\n"
	config += `	items = { "my_geolocations" = {` + "\n"
	config += `		continents = [{` + "\n"
	config += `			id = data.fmc_continents.test.items["North America"].id` + "\n"
	config += `		}]` + "\n"
	config += `		countries = [{` + "\n"
	config += `			id = data.fmc_countries.test.items["Poland"].id` + "\n"
	config += `		}]` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_geolocations" "test" {
			depends_on = [fmc_geolocations.test]
			items = {
				"my_geolocations" = {
				}
			}
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
