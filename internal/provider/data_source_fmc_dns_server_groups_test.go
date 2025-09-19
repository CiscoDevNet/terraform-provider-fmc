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

func TestAccDataSourceFmcDNSServerGroups(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_dns_server_groups.test", "items.my_dns_server_groups.id"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_dns_server_groups.test", "items.my_dns_server_groups.type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_dns_server_groups.test", "items.my_dns_server_groups.default_domain", "example.com"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_dns_server_groups.test", "items.my_dns_server_groups.timeout", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_dns_server_groups.test", "items.my_dns_server_groups.retries", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_dns_server_groups.test", "items.my_dns_server_groups.dns_servers.0.ip", "10.10.10.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_dns_server_groups.test", "items.my_dns_server_groups.dns_servers.0.ip", "10.10.10.1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcDNSServerGroupsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcDNSServerGroupsConfig() string {
	config := `resource "fmc_dns_server_groups" "test" {` + "\n"
	config += `	items = { "my_dns_server_groups" = {` + "\n"
	config += `		default_domain = "example.com"` + "\n"
	config += `		timeout = 2` + "\n"
	config += `		retries = 2` + "\n"
	config += `		dns_servers = [{` + "\n"
	config += `			ip = "10.10.10.1"` + "\n"
	config += `		}]` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_dns_server_groups" "test" {
			depends_on = [fmc_dns_server_groups.test]
			items = {
				"my_dns_server_groups" = {
				}
			}
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
