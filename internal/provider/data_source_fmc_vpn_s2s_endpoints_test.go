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

func TestAccDataSourceFmcVPNS2SEndpoints(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id")
	}
	var checks []resource.TestCheckFunc
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcVPNS2SEndpointsPrerequisitesConfig + testAccDataSourceFmcVPNS2SEndpointsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcVPNS2SEndpointsPrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id

data "fmc_device" "test" {
  id = var.device_id
}

data "fmc_network" "test" {
  name = "any-ipv4"
}

resource "fmc_device_physical_interface" "test" {
  device_id = data.fmc_device.test.id
  name      = "GigabitEthernet0/1"
  mode      = "NONE"
  logical_name = "my_phy_s2s_vpn_endpoints"
  ipv4_static_address = "10.198.21.1"
  ipv4_static_netmask = "24"
}

resource "fmc_vpn_s2s" "test" {
  name             = "my_s2s_vpn_endpoints"
  route_based      = false
  network_topology = "POINT_TO_POINT"
  ikev1            = false
  ikev2            = true
}

resource "fmc_vpn_s2s_ike_settings" "test" {
  vpn_s2s_id                  = fmc_vpn_s2s.test.id
  ikev2_authentication_type   = "MANUAL_PRE_SHARED_KEY"
  ikev2_manual_pre_shared_key = "my_pre_shared_key123"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcVPNS2SEndpointsConfig() string {
	config := `resource "fmc_vpn_s2s_endpoints" "test" {` + "\n"
	config += `	vpn_s2s_id = fmc_vpn_s2s.test.id` + "\n"
	config += `	items = { (data.fmc_device.test.name) = {` + "\n"
	config += `		peer_type = "PEER"` + "\n"
	config += `		extranet_device = false` + "\n"
	config += `		device_id = data.fmc_device.test.id` + "\n"
	config += `		interface_id = fmc_device_physical_interface.test.id` + "\n"
	config += `		interface_public_ip_address = "10.1.1.1"` + "\n"
	config += `		connection_type = "BIDIRECTIONAL"` + "\n"
	config += `		protected_networks = [{` + "\n"
	config += `			id = data.fmc_network.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		local_identity_type = "EMAILID"` + "\n"
	config += `		local_identity_string = "me@cisco.com"` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_vpn_s2s_endpoints" "test" {
			depends_on = [fmc_vpn_s2s_endpoints.test]
			vpn_s2s_id = fmc_vpn_s2s.test.id
			items = {
				(data.fmc_device.test.name) = {
				}
			}
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
