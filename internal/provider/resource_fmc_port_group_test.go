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

func TestAccFmcPortGroup(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_port_group.test", "name", "my_port_group"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_port_group.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_port_group.test", "description", "My port group description"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_port_group.test", "objects.0.type", "ProtocolPortObject"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcPortGroupPrerequisitesConfig + testAccFmcPortGroupConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcPortGroupPrerequisitesConfig + testAccFmcPortGroupConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_port_group.test",
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

const testAccFmcPortGroupPrerequisitesConfig = `
resource "fmc_port" "test" {
  name        = "fmc_port_group_port"
  description = "My PORT id"
  protocol    = "TCP"
  port        = "443"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcPortGroupConfig_minimum() string {
	config := `resource "fmc_port_group" "test" {` + "\n"
	config += `	name = "my_port_group"` + "\n"
	config += `	objects = [{` + "\n"
	config += `		id = fmc_port.test.id` + "\n"
	config += `		type = "ProtocolPortObject"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcPortGroupConfig_all() string {
	config := `resource "fmc_port_group" "test" {` + "\n"
	config += `	name = "my_port_group"` + "\n"
	config += `	description = "My port group description"` + "\n"
	config += `	overridable = true` + "\n"
	config += `	objects = [{` + "\n"
	config += `		id = fmc_port.test.id` + "\n"
	config += `		type = "ProtocolPortObject"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
