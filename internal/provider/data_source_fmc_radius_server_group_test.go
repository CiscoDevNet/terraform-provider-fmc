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

func TestAccDataSourceFmcRadiusServerGroup(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_radius_server_group.test", "name", "my_radius_server_group"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_radius_server_group.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_radius_server_group.test", "description", "My RADIUS Server Group object"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_radius_server_group.test", "group_accounting_mode", "SINGLE"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_radius_server_group.test", "retry_interval", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_radius_server_group.test", "authorize_only", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_radius_server_group.test", "interim_account_update_interval", "24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_radius_server_group.test", "dynamic_authorization", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_radius_server_group.test", "dynamic_authorization_port", "1700"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_radius_server_group.test", "radius_servers.0.hostname", "10.10.10.10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_radius_server_group.test", "radius_servers.0.radius_server_enabled_message_authenticator", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_radius_server_group.test", "radius_servers.0.authentication_port", "1812"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_radius_server_group.test", "radius_servers.0.accounting_port", "1813"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_radius_server_group.test", "radius_servers.0.timeout", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_radius_server_group.test", "radius_servers.0.use_routing_to_select_interface", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcRadiusServerGroupConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccNamedDataSourceFmcRadiusServerGroupConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcRadiusServerGroupConfig() string {
	config := `resource "fmc_radius_server_group" "test" {` + "\n"
	config += `	name = "my_radius_server_group"` + "\n"
	config += `	description = "My RADIUS Server Group object"` + "\n"
	config += `	group_accounting_mode = "SINGLE"` + "\n"
	config += `	retry_interval = 10` + "\n"
	config += `	authorize_only = true` + "\n"
	config += `	interim_account_update_interval = 24` + "\n"
	config += `	dynamic_authorization = true` + "\n"
	config += `	dynamic_authorization_port = 1700` + "\n"
	config += `	radius_servers = [{` + "\n"
	config += `		hostname = "10.10.10.10"` + "\n"
	config += `		radius_server_enabled_message_authenticator = true` + "\n"
	config += `		authentication_port = 1812` + "\n"
	config += `		key = "my_secret_key"` + "\n"
	config += `		accounting_port = 1813` + "\n"
	config += `		timeout = 10` + "\n"
	config += `		use_routing_to_select_interface = true` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_radius_server_group" "test" {
			id = fmc_radius_server_group.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcRadiusServerGroupConfig() string {
	config := `resource "fmc_radius_server_group" "test" {` + "\n"
	config += `	name = "my_radius_server_group"` + "\n"
	config += `	description = "My RADIUS Server Group object"` + "\n"
	config += `	group_accounting_mode = "SINGLE"` + "\n"
	config += `	retry_interval = 10` + "\n"
	config += `	authorize_only = true` + "\n"
	config += `	interim_account_update_interval = 24` + "\n"
	config += `	dynamic_authorization = true` + "\n"
	config += `	dynamic_authorization_port = 1700` + "\n"
	config += `	radius_servers = [{` + "\n"
	config += `		hostname = "10.10.10.10"` + "\n"
	config += `		radius_server_enabled_message_authenticator = true` + "\n"
	config += `		authentication_port = 1812` + "\n"
	config += `		key = "my_secret_key"` + "\n"
	config += `		accounting_port = 1813` + "\n"
	config += `		timeout = 10` + "\n"
	config += `		use_routing_to_select_interface = true` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_radius_server_group" "test" {
			name = fmc_radius_server_group.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
