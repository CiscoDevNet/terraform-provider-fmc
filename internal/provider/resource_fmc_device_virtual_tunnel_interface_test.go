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

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccFmcDeviceVirtualTunnelInterface(t *testing.T) {
	if v := os.Getenv("FMC_VERSION"); v != "" && slices.Contains([]string{"7.2"}, v) {
		t.Skip("skipping test for FMC version " + v)
	}
	if os.Getenv("TF_VAR_device_id") == "" || os.Getenv("TF_VAR_interface_name") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id and TF_VAR_interface_name")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_device_virtual_tunnel_interface.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_device_virtual_tunnel_interface.test", "name"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_virtual_tunnel_interface.test", "tunnel_type", "STATIC"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_virtual_tunnel_interface.test", "logical_name", "my_vti_interface"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_virtual_tunnel_interface.test", "enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_virtual_tunnel_interface.test", "description", "My VTI"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_virtual_tunnel_interface.test", "priority", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_virtual_tunnel_interface.test", "tunnel_id", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_virtual_tunnel_interface.test", "tunnel_mode", "ipv4"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_virtual_tunnel_interface.test", "ipv4_static_address", "10.10.10.10"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_virtual_tunnel_interface.test", "ipv4_static_netmask", "24"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_virtual_tunnel_interface.test", "ip_based_monitoring", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_virtual_tunnel_interface.test", "ip_based_monitoring_type", "PEER_IPV4"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_virtual_tunnel_interface.test", "ip_based_monitoring_peer_ip", "10.10.10.100"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcDeviceVirtualTunnelInterfacePrerequisitesConfig + testAccFmcDeviceVirtualTunnelInterfaceConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcDeviceVirtualTunnelInterfacePrerequisitesConfig + testAccFmcDeviceVirtualTunnelInterfaceConfig_all(),
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

const testAccFmcDeviceVirtualTunnelInterfacePrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id
variable "interface_name" {default = null} // tests will set $TF_VAR_interface_name

resource "fmc_device_physical_interface" "test" {
  device_id           = var.device_id
  name                = var.interface_name
  mode                = "NONE"
  logical_name        = "myinterface-0-1"
  ipv4_static_address = "10.12.1.1"
  ipv4_static_netmask = "255.255.255.240"
}

resource "fmc_security_zone" "test" {
  name           = "my_sec_zone_vti"
  interface_type = "ROUTED"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcDeviceVirtualTunnelInterfaceConfig_minimum() string {
	config := `resource "fmc_device_virtual_tunnel_interface" "test" {` + "\n"
	config += `	device_id = fmc_device_physical_interface.test.device_id` + "\n"
	config += `	tunnel_type = "STATIC"` + "\n"
	config += `	logical_name = "my_vti_interface"` + "\n"
	config += `	tunnel_id = 100` + "\n"
	config += `	tunnel_source_interface_id = fmc_device_physical_interface.test.id` + "\n"
	config += `	tunnel_source_interface_name = fmc_device_physical_interface.test.name` + "\n"
	config += `	tunnel_mode = "ipv4"` + "\n"
	config += `	ipv4_static_address = "10.10.10.10"` + "\n"
	config += `	ipv4_static_netmask = "24"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcDeviceVirtualTunnelInterfaceConfig_all() string {
	config := `resource "fmc_device_virtual_tunnel_interface" "test" {` + "\n"
	config += `	device_id = fmc_device_physical_interface.test.device_id` + "\n"
	config += `	tunnel_type = "STATIC"` + "\n"
	config += `	logical_name = "my_vti_interface"` + "\n"
	config += `	enabled = true` + "\n"
	config += `	description = "My VTI"` + "\n"
	config += `	security_zone_id = fmc_security_zone.test.id` + "\n"
	config += `	priority = 100` + "\n"
	config += `	tunnel_id = 100` + "\n"
	config += `	tunnel_source_interface_id = fmc_device_physical_interface.test.id` + "\n"
	config += `	tunnel_source_interface_name = fmc_device_physical_interface.test.name` + "\n"
	config += `	tunnel_mode = "ipv4"` + "\n"
	config += `	ipv4_static_address = "10.10.10.10"` + "\n"
	config += `	ipv4_static_netmask = "24"` + "\n"
	config += `	ip_based_monitoring = true` + "\n"
	config += `	ip_based_monitoring_type = "PEER_IPV4"` + "\n"
	config += `	ip_based_monitoring_peer_ip = "10.10.10.100"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

func TestAccFmcDeviceVirtualTunnelInterface_DynamicInterface(t *testing.T) {
	if v := os.Getenv("FMC_VERSION"); v != "" && slices.Contains([]string{"7.2"}, v) {
		t.Skip("skipping test for FMC version " + v)
	}
	if os.Getenv("TF_VAR_device_id") == "" || os.Getenv("TF_VAR_interface_name") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id and TF_VAR_interface_name")
	}

	prerequisites := `resource "fmc_device_loopback_interface" "test" {
					device_id           = fmc_device_physical_interface.test.device_id
					loopback_id         = 10
					logical_name        = "my_loopback_virtual_tunnel_interface"
					ipv4_static_address = "10.19.10.1"
					ipv4_static_netmask = "255.255.255.0"
				}
	`

	step_01 := `resource "fmc_device_virtual_tunnel_interface" "test-dynamic" {
					device_id                    = fmc_device_physical_interface.test.device_id
					tunnel_type                  = "DYNAMIC"
					logical_name                 = "my_vti_interface_static_dynamic"
					tunnel_id                    = 150
					tunnel_source_interface_id   = fmc_device_physical_interface.test.id
					tunnel_source_interface_name = fmc_device_physical_interface.test.name
					tunnel_mode                  = "ipv4"
					borrow_ip_interface_id       = fmc_device_loopback_interface.test.id
					borrow_ip_interface_name     = fmc_device_loopback_interface.test.name
				}
	`

	var checks_step01 []resource.TestCheckFunc
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttrSet("fmc_device_virtual_tunnel_interface.test-dynamic", "type"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttrSet("fmc_device_virtual_tunnel_interface.test-dynamic", "name"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_device_virtual_tunnel_interface.test-dynamic", "tunnel_type", "DYNAMIC"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_device_virtual_tunnel_interface.test-dynamic", "logical_name", "my_vti_interface_static_dynamic"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_device_virtual_tunnel_interface.test-dynamic", "tunnel_id", "150"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccFmcDeviceVirtualTunnelInterfacePrerequisitesConfig + prerequisites + step_01,
		Check:  resource.ComposeTestCheckFunc(checks_step01...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps:                    steps,
	})
}

// FMCBUG: CSCwp02259 - POST response body may contain incorrect object id (Creation of `Dynamic` interface may return `Static` interface id)
// Hence both interface types are created to confirm this is handled correctly.
func TestAccFmcDeviceVirtualTunnelInterface_StaticAndDynamicInterface(t *testing.T) {
	if v := os.Getenv("FMC_VERSION"); v != "" && slices.Contains([]string{"7.2"}, v) {
		t.Skip("skipping test for FMC version " + v)
	}
	if os.Getenv("TF_VAR_device_id") == "" || os.Getenv("TF_VAR_interface_name") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id and TF_VAR_interface_name")
	}

	prerequisites := `resource "fmc_device_loopback_interface" "test" {
					device_id           = fmc_device_physical_interface.test.device_id
					loopback_id         = 10
					logical_name        = "my_loopback_virtual_tunnel_interface"
					ipv4_static_address = "10.19.10.1"
					ipv4_static_netmask = "255.255.255.0"
				}
	`

	step_01 := `resource "fmc_device_virtual_tunnel_interface" "test-static" {
					device_id                    = fmc_device_physical_interface.test.device_id
					tunnel_type                  = "STATIC"
					logical_name                 = "my_vti_interface_static"
					tunnel_id                    = 200
					tunnel_source_interface_id   = fmc_device_physical_interface.test.id
					tunnel_source_interface_name = fmc_device_physical_interface.test.name
					tunnel_mode                  = "ipv4"
					ipv4_address                 = "10.1.1.3"
					ipv4_netmask                 = "255.255.255.0"
				}

				resource "fmc_device_virtual_tunnel_interface" "test-dynamic" {
					device_id                    = fmc_device_physical_interface.test.device_id
					tunnel_type                  = "DYNAMIC"
					logical_name                 = "my_vti_interface_static_dynamic"
					tunnel_id                    = 200
					tunnel_source_interface_id   = fmc_device_physical_interface.test.id
					tunnel_source_interface_name = fmc_device_physical_interface.test.name
					tunnel_mode                  = "ipv4"
					borrow_ip_interface_id       = fmc_device_loopback_interface.test.id
					borrow_ip_interface_name     = fmc_device_loopback_interface.test.name
				}
	`

	var checks_step01 []resource.TestCheckFunc
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttrSet("fmc_device_virtual_tunnel_interface.test-static", "type"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttrSet("fmc_device_virtual_tunnel_interface.test-static", "name"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_device_virtual_tunnel_interface.test-static", "tunnel_type", "STATIC"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_device_virtual_tunnel_interface.test-static", "logical_name", "my_vti_interface_static"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_device_virtual_tunnel_interface.test-static", "tunnel_id", "200"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttrSet("fmc_device_virtual_tunnel_interface.test-dynamic", "type"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttrSet("fmc_device_virtual_tunnel_interface.test-dynamic", "name"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_device_virtual_tunnel_interface.test-dynamic", "tunnel_type", "DYNAMIC"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_device_virtual_tunnel_interface.test-dynamic", "logical_name", "my_vti_interface_static_dynamic"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_device_virtual_tunnel_interface.test-dynamic", "tunnel_id", "200"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccFmcDeviceVirtualTunnelInterfacePrerequisitesConfig + prerequisites + step_01,
		Check:  resource.ComposeTestCheckFunc(checks_step01...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps:                    steps,
	})
}
