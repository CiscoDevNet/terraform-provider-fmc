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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource

func TestAccDataSourceFmcDeviceDDNS(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" || os.Getenv("TF_VAR_interface_name") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id and TF_VAR_interface_name")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_device_ddns.test", "type"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcDeviceDDNSPrerequisitesConfig + testAccDataSourceFmcDeviceDDNSConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcDeviceDDNSPrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id
variable "interface_name" {default = null} // tests will set $TF_VAR_interface_name

resource "fmc_device_physical_interface" "test" {
  device_id           = var.device_id
  name                = var.interface_name
  logical_name        = "myinterface-0-1"
  mode                = "NONE"
  enabled             = true
  ipv4_static_address = "10.1.1.1"
  ipv4_static_netmask = "24"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcDeviceDDNSConfig() string {
	config := `resource "fmc_device_ddns" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	dhcp_client_request_type = "BOTH_A_AND_PTR_RECORD"` + "\n"
	config += `	dhcp_client_broadcast = true` + "\n"
	config += `	dynamic_dns_update = "BOTH_A_AND_PTR_RECORD"` + "\n"
	config += `	dhcp_client_request_override = true` + "\n"
	config += `	ddns_update_methods = [{` + "\n"
	config += `		name = "my_ddns_method"` + "\n"
	config += `		method = "DDNS"` + "\n"
	config += `		update_interval_day = 1` + "\n"
	config += `		update_interval_minute = 5` + "\n"
	config += `		update_records = "BOTH_A_AND_PTR_RECORDS"` + "\n"
	config += `	}]` + "\n"
	config += `	ddns_interface_settings = [{` + "\n"
	config += `		interface_id = fmc_device_physical_interface.test.id` + "\n"
	config += `		interface_name = fmc_device_physical_interface.test.name` + "\n"
	config += `		interface_type = fmc_device_physical_interface.test.type` + "\n"
	config += `		method_name = "my_ddns_method"` + "\n"
	config += `		hostname = "host.example.com"` + "\n"
	config += `		dynamic_dns_update = "BOTH_A_AND_PTR_RECORD"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_ddns" "test" {
			id = fmc_device_ddns.test.id
			device_id = var.device_id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
