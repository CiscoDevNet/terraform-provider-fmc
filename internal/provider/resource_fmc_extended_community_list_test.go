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

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccFmcExtendedCommunityList(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_extended_community_list.test", "name", "my_extended_community_list"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_extended_community_list.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_extended_community_list.test", "sub_type", "Standard"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_extended_community_list.test", "entries.0.action", "PERMIT"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_extended_community_list.test", "entries.0.route_target", "64512:1010"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccFmcExtendedCommunityListConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_extended_community_list.test",
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
// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcExtendedCommunityListConfig_all() string {
	config := `resource "fmc_extended_community_list" "test" {` + "\n"
	config += `	name = "my_extended_community_list"` + "\n"
	config += `	sub_type = "Standard"` + "\n"
	config += `	entries = [{` + "\n"
	config += `		action = "PERMIT"` + "\n"
	config += `		route_target = "64512:1010"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
