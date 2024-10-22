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

func TestAccFmcHosts(t *testing.T) {
	var checks_step01 []resource.TestCheckFunc
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_hosts.test", "items.hosts_1.ip", "1.2.3.1"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_hosts.test", "items.hosts_1.description", "host1"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_hosts.test", "items.hosts_2.ip", "1.2.3.2"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_hosts.test", "items.hosts_3.ip", "1.2.3.3"))

	var checks_step02 []resource.TestCheckFunc
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttr("fmc_hosts.test", "items.hosts_1.ip", "1.2.3.1"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttr("fmc_hosts.test", "items.hosts_1.description", "host1 new description"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttr("fmc_hosts.test", "items.hosts_2.ip", "1.2.3.2"))
	checks_step02 = append(checks_step02, resource.TestCheckNoResourceAttr("fmc_hosts.test", "items.hosts_3"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttr("fmc_hosts.test", "items.hosts_4.ip", "1.2.3.4"))

	var steps []resource.TestStep

	steps = append(steps, resource.TestStep{
		Config: testAccFmcHostsConfig_step01(),
		Check:  resource.ComposeTestCheckFunc(checks_step01...),
	})
	steps = append(steps, resource.TestStep{
		Config: testAccFmcHostsConfig_step02(),
		Check:  resource.ComposeTestCheckFunc(checks_step02...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcHostsConfig_minimum() string {
	config := `resource "fmc_hosts" "test" {` + "\n"
	config += `	items = ` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcHostsConfig_all() string {
	config := `resource "fmc_hosts" "test" {` + "\n"
	config += `	items = { "hosts_1" = {` + "\n"
	config += `		description = "My Host 1"` + "\n"
	config += `		overridable = true` + "\n"
	config += `		ip = "10.1.1.1"` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

func testAccFmcHostsConfig_step01() string {
	config := `resource "fmc_hosts" "test" {` + "\n"
	config += `	items = {` + "\n"
	config += `		"hosts_1" = {` + "\n"
	config += `			ip = "1.2.3.1",` + "\n"
	config += `			description = "host1"` + "\n"
	config += `			overridable = true` + "\n"
	config += `		},` + "\n"
	config += `		"hosts_2" = {` + "\n"
	config += `			ip = "1.2.3.2",` + "\n"
	config += `		},` + "\n"
	config += `		"hosts_3" = {` + "\n"
	config += `			ip = "1.2.3.3",` + "\n"
	config += `		},` + "\n"
	config += `	} ` + "\n"
	config += `}` + "\n"
	return config
}

func testAccFmcHostsConfig_step02() string {
	config := `resource "fmc_hosts" "test" {` + "\n"
	config += `	items = {` + "\n"
	config += `		"hosts_1" = {` + "\n"
	config += `			ip = "1.2.3.1",` + "\n"
	config += `			description = "host1 new description"` + "\n"
	config += `		},` + "\n"
	config += `		"hosts_2" = {` + "\n"
	config += `			ip = "1.2.3.2",` + "\n"
	config += `		},` + "\n"
	config += `		"hosts_4" = {` + "\n"
	config += `			ip = "1.2.3.4",` + "\n"
	config += `		},` + "\n"
	config += `	} ` + "\n"
	config += `}` + "\n"
	return config
}
