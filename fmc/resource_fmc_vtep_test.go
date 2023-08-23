package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcVtepCreateBasic(t *testing.T) {
	device := "ftd.adyah.cisco"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcVtepCreateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcVtepCreateConfigBasic(device),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcVtepCreateExists("fmc_device_vtep.my_fmc_device_vtep"),
				),
			},
		},
	})
}

func testAccCheckFmcVtepCreateDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_device_vtep" {
			continue
		}

		id := rs.Primary.ID
		ctx := context.Background()
		_, err := c.DeleteFmcVTEPDetails(ctx, rs.Primary.Attributes["device_id"], id)

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "404") {
			return err
		}
	}

	return nil
}

func testAccCheckFmcVtepCreateConfigBasic(device string) string {
	return fmt.Sprintf(`
	data "fmc_devices" "device" {
		name = "%s"
	}
	data "fmc_device_physical_interfaces" "zero_physical_interface" {
		device_id = data.fmc_devices.device.id
		name = "TenGigabitEthernet0/0"
	}
	resource "fmc_host_objects" "test1" {
	  name        = "test1"
	  value       = "172.16.1.1"
	}
	resource "fmc_host_objects" "test2" {
	  name        = "test2"
	  value       = "172.16.2.1"
	}
	
	resource "fmc_network_group_objects" "TestPrivateGroup" {
	  name = "TestPrivateGroup"
	  description = "Testing groups"
	  objects {
		  id = fmc_host_objects.test1.id
		  type = fmc_host_objects.test1.type
	  }
	  objects {
		  id = fmc_host_objects.test2.id
		  type = fmc_host_objects.test2.type
	  }
	}
	resource "fmc_device_vtep" "my_fmc_device_vtep" {
		device_id = data.fmc_devices.device.id
		nve_enabled = true
	
		nve_vtep_id = 1
		nve_encapsulation_type = "VXLAN"
		nve_destination_port = 4789
		source_interface_id = data.fmc_device_physical_interfaces.zero_physical_interface.id 
		
		nve_neighbor_discovery_type= "NONE"
	}
    `, device)
}

func testAccCheckFmcVtepCreateExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID set")
		}

		return nil
	}
}
