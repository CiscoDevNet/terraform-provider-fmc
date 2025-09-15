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

func TestAccFmcVPNS2SEndpoints(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" || os.Getenv("TF_VAR_interface_name") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id and TF_VAR_interface_name")
	}
	var checks []resource.TestCheckFunc

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccFmcVPNS2SEndpointsPrerequisitesConfig + testAccFmcVPNS2SEndpointsConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
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

const testAccFmcVPNS2SEndpointsPrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id
variable "interface_name" {default = null} // tests will set $TF_VAR_interface_name

data "fmc_device" "test" {
  id = var.device_id
}

data "fmc_network" "test" {
  name = "any-ipv4"
}

resource "fmc_device_physical_interface" "test" {
  device_id = data.fmc_device.test.id
  name      = var.interface_name
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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcVPNS2SEndpointsConfig_all() string {
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
	return config
}

// End of section. //template:end testAccConfigAll

func TestAccFmcVPNS2SEndpoints_Sequential(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id")
	}

	step_01 := `resource "fmc_vpn_s2s_endpoints" "test" {` + "\n" +
		`	vpn_s2s_id = fmc_vpn_s2s.test.id` + "\n" +
		`	items = {` + "\n" +
		`	  "ext-dev" = {` + "\n" +
		`		extranet_device     = true` + "\n" +
		`		peer_type           = "PEER"` + "\n" +
		`		extranet_ip_address = "1.1.1.1"` + "\n" +
		`	  }` + "\n" +
		`	}` + "\n" +
		`	depends_on = [fmc_vpn_s2s_ike_settings.test]` + "\n" +
		`}` + "\n"

	step_02 := `resource "fmc_vpn_s2s_endpoints" "test" {` + "\n" +
		`	vpn_s2s_id = fmc_vpn_s2s.test.id` + "\n" +
		`	items = {` + "\n" +
		`	  "ext-dev" = {` + "\n" +
		`		extranet_device     = true` + "\n" +
		`		peer_type           = "PEER"` + "\n" +
		`		extranet_ip_address = "1.1.1.1"` + "\n" +
		`	  }` + "\n" +
		`	  (data.fmc_device.test.name) = {` + "\n" +
		`	  	peer_type = "PEER"` + "\n" +
		`		extranet_device = false` + "\n" +
		`		device_id = data.fmc_device.test.id` + "\n" +
		`		interface_id = fmc_device_physical_interface.test.id` + "\n" +
		`		connection_type = "BIDIRECTIONAL"` + "\n" +
		`		protected_networks = [{` + "\n" +
		`			id = data.fmc_network.test.id` + "\n" +
		`		}]` + "\n" +
		`	  }` + "\n" +
		`   }` + "\n" +
		`	depends_on = [fmc_vpn_s2s_ike_settings.test]` + "\n" +
		`}` + "\n"

	step_03 := `resource "fmc_vpn_s2s_endpoints" "test" {` + "\n" +
		`	vpn_s2s_id = fmc_vpn_s2s.test.id` + "\n" +
		`	items = {` + "\n" +
		`	  "ext-dev" = {` + "\n" +
		`		extranet_device     = true` + "\n" +
		`		peer_type           = "PEER"` + "\n" +
		`		extranet_ip_address = "1.1.1.1"` + "\n" +
		`	  }` + "\n" +
		`	  (data.fmc_device.test.name) = {` + "\n" +
		`	  	peer_type = "PEER"` + "\n" +
		`		extranet_device = false` + "\n" +
		`		device_id = data.fmc_device.test.id` + "\n" +
		`		interface_id = fmc_device_physical_interface.test.id` + "\n" +
		`		connection_type = "BIDIRECTIONAL"` + "\n" +
		`		protected_networks = [{` + "\n" +
		`			id = data.fmc_network.test.id` + "\n" +
		`		}]` + "\n" +
		`		local_identity_type = "EMAILID"` + "\n" +
		`		local_identity_string = "me@cisco.com"` + "\n" +
		`	  }` + "\n" +
		`   }` + "\n" +
		`	depends_on = [fmc_vpn_s2s_ike_settings.test]` + "\n" +
		`}` + "\n"

	step_04 := `resource "fmc_vpn_s2s_endpoints" "test" {` + "\n" +
		`	vpn_s2s_id = fmc_vpn_s2s.test.id` + "\n" +
		`	items = {` + "\n" +
		`	  "ext-dev" = {` + "\n" +
		`		extranet_device     = true` + "\n" +
		`		peer_type           = "PEER"` + "\n" +
		`		extranet_ip_address = "1.1.1.1"` + "\n" +
		`	  }` + "\n" +
		`	  (data.fmc_device.test.name) = {` + "\n" +
		`	  	peer_type = "PEER"` + "\n" +
		`		extranet_device = false` + "\n" +
		`		device_id = data.fmc_device.test.id` + "\n" +
		`		interface_id = fmc_device_physical_interface.test.id` + "\n" +
		`		connection_type = "BIDIRECTIONAL"` + "\n" +
		`		protected_networks = [{` + "\n" +
		`			id = data.fmc_network.test.id` + "\n" +
		`		}]` + "\n" +
		`	  }` + "\n" +
		`   }` + "\n" +
		`	depends_on = [fmc_vpn_s2s_ike_settings.test]` + "\n" +
		`}` + "\n"

	step_05 := `resource "fmc_vpn_s2s_endpoints" "test" {` + "\n" +
		`	vpn_s2s_id = fmc_vpn_s2s.test.id` + "\n" +
		`	items = {` + "\n" +
		`	  "ext-dev" = {` + "\n" +
		`		extranet_device     = true` + "\n" +
		`		peer_type           = "PEER"` + "\n" +
		`		extranet_ip_address = "1.1.1.1"` + "\n" +
		`	  }` + "\n" +
		`	}` + "\n" +
		`	depends_on = [fmc_vpn_s2s_ike_settings.test]` + "\n" +
		`}` + "\n"

	steps := []resource.TestStep{{
		Config: testAccFmcVPNS2SEndpointsPrerequisitesConfig + step_01,
	}, {
		Config: testAccFmcVPNS2SEndpointsPrerequisitesConfig + step_02,
	}, {
		Config: testAccFmcVPNS2SEndpointsPrerequisitesConfig + step_03,
	}, {
		Config: testAccFmcVPNS2SEndpointsPrerequisitesConfig + step_04,
	}, {
		Config: testAccFmcVPNS2SEndpointsPrerequisitesConfig + step_05,
	}}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
