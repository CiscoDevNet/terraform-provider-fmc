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

func TestAccFmcDeviceEtherChannelInterface(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" || os.Getenv("TF_VAR_interface_name") == "" || os.Getenv("FMC_DEVICE_ETHERCHANNEL_INTERFACE") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id and TF_VAR_interface_name and FMC_DEVICE_ETHERCHANNEL_INTERFACE")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_device_etherchannel_interface.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_device_etherchannel_interface.test", "is_multi_instance"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_etherchannel_interface.test", "logical_name", "myinterface-0-1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_etherchannel_interface.test", "description", "my description"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_device_etherchannel_interface.test", "mode"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_device_etherchannel_interface.test", "name"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_etherchannel_interface.test", "mtu", "9000"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_etherchannel_interface.test", "ether_channel_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_etherchannel_interface.test", "ipv4_static_address", "10.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_etherchannel_interface.test", "ipv4_static_netmask", "24"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcDeviceEtherChannelInterfacePrerequisitesConfig + testAccFmcDeviceEtherChannelInterfaceConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcDeviceEtherChannelInterfacePrerequisitesConfig + testAccFmcDeviceEtherChannelInterfaceConfig_all(),
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

const testAccFmcDeviceEtherChannelInterfacePrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id
variable "interface_name" { default = null } // tests will set $TF_VAR_interface_name

data "fmc_device_physical_interface" "test" {
  device_id = var.device_id
  name      = var.interface_name
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcDeviceEtherChannelInterfaceConfig_minimum() string {
	config := `resource "fmc_device_etherchannel_interface" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	logical_name = "iface_minimum"` + "\n"
	config += `	ether_channel_id = "1"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcDeviceEtherChannelInterfaceConfig_all() string {
	config := `resource "fmc_device_etherchannel_interface" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	logical_name = "myinterface-0-1"` + "\n"
	config += `	enabled = true` + "\n"
	config += `	description = "my description"` + "\n"
	config += `	mtu = 9000` + "\n"
	config += `	ether_channel_id = "1"` + "\n"
	config += `	selected_interfaces = [{` + "\n"
	config += `		id = data.fmc_device_physical_interface.test.id` + "\n"
	config += `		name = data.fmc_device_physical_interface.test.name` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_static_address = "10.1.1.1"` + "\n"
	config += `	ipv4_static_netmask = "24"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// TestAccFmcDeviceEtherChannelInterfaceMode covers `mode`, which is optional and computed, so the
// generated test above cannot exercise it - attributes marked as computed are left out of the
// generated configurations. The first step leaves `mode` unset (implicit: adjustBody sends NONE
// and FMC owns the value), the second one sets it explicitly, which also covers the transition
// from an unset to an explicitly configured `mode` and, in the last step, a change from one
// explicitly configured value to another.
// Only NONE and PASSIVE are used here: INLINE and TAP are assigned by FMC itself once the
// interface becomes a member of an inline set and are rejected otherwise, ERSPAN additionally
// requires `erspan_source_ip` and `erspan_flow_id`, and SWITCHPORT is limited to models with
// switch-capable ports.
func TestAccFmcDeviceEtherChannelInterfaceMode(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" || os.Getenv("TF_VAR_interface_name") == "" || os.Getenv("FMC_DEVICE_ETHERCHANNEL_INTERFACE") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id and TF_VAR_interface_name and FMC_DEVICE_ETHERCHANNEL_INTERFACE")
	}

	steps := []resource.TestStep{{
		// `mode` not configured
		Config: testAccFmcDeviceEtherChannelInterfacePrerequisitesConfig + testAccFmcDeviceEtherChannelInterfaceConfig_mode(""),
		Check:  resource.TestCheckResourceAttr("fmc_device_etherchannel_interface.test", "mode", "NONE"),
	}, {
		// `mode` configured explicitly
		Config: testAccFmcDeviceEtherChannelInterfacePrerequisitesConfig + testAccFmcDeviceEtherChannelInterfaceConfig_mode("NONE"),
		Check:  resource.TestCheckResourceAttr("fmc_device_etherchannel_interface.test", "mode", "NONE"),
	}, {
		// `mode` configured explicitly to a value other than the implicit NONE
		Config: testAccFmcDeviceEtherChannelInterfacePrerequisitesConfig + testAccFmcDeviceEtherChannelInterfaceConfig_mode("PASSIVE"),
		Check:  resource.TestCheckResourceAttr("fmc_device_etherchannel_interface.test", "mode", "PASSIVE"),
	}}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps:                    steps,
	})
}

// testAccFmcDeviceEtherChannelInterfaceConfig_mode renders the configuration with `mode` set to
// the provided value, or with `mode` omitted altogether if the value is empty.
func testAccFmcDeviceEtherChannelInterfaceConfig_mode(mode string) string {
	config := `resource "fmc_device_etherchannel_interface" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	ether_channel_id = "1"` + "\n"
	config += `	logical_name = "iface_mode"` + "\n"
	if mode != "" {
		config += `	mode = "` + mode + `"` + "\n"
	}
	config += `}` + "\n"
	return config
}
