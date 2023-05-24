package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcPhysicalInterfacesCreateBasic(t *testing.T) {
	device := "ftd.adyah.cisco"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcPhysicalInterfacesCreateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcPhysicalInterfacesCreateConfigBasic(device),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcPhysicalInterfacesCreateExists("fmc_device_physical_interfaces.my_fmc_device_physical_interfaces"),
				),
			},
		},
	})
}

func testAccCheckFmcPhysicalInterfacesCreateDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_device_physical_interfaces" {
			continue
		}

		id := rs.Primary.ID
		ctx := context.Background()
		//_, err := c.DeleteFmcPhysicalInterfacesDetails(ctx, rs.Primary.Attributes["device_id"], id)
		_, err := c.UpdateFmcPhysicalInterface(ctx, rs.Primary.Attributes["device_id"], id, &PhysicalInterfaceRequest{
			ID:     id,
			Ifname: "",
			Mode:   "NONE",
			Name:   "TenGigabitEthernet0/0",
			MTU:    1700,
			// Description:  "",
			// SecurityZone: PhysicalInterfaceSecurityZone,
		})

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "404") {
			return err
		}
	}

	return nil
}

func testAccCheckFmcPhysicalInterfacesCreateConfigBasic(device string) string {
	return fmt.Sprintf(`
	data "fmc_devices" "device" {
		name = "%s"
	}
	
	data "fmc_device_physical_interfaces" "device_physical_interface" {
		device_id = data.fmc_devices.device.id
		name = "TenGigabitEthernet0/0"
	}

	resource "fmc_security_zone" "my_security_zone" {
	  name = "physical-interface-zone"
	  interface_mode = "ROUTED"
	}
	resource "fmc_device_physical_interfaces" "my_fmc_device_physical_interfaces" {
		device_id = data.fmc_devices.device.id
   		physical_interface_id= data.fmc_device_physical_interfaces.device_physical_interface.id
    	name =   data.fmc_device_physical_interfaces.device_physical_interface.name
    	security_zone_id= fmc_security_zone.my_security_zone.id
    	if_name = "logical-name"
    	description = "Physical Interface"
    	mtu =  1700
    	mode = "NONE"
    	ipv4_static_address = "10.20.220.45"
    	ipv4_static_netmask = 24
    	ipv4_dhcp_enabled = false
    	ipv4_dhcp_route_metric = 1
    	ipv6_address = "2001:1234:5678:1234::"
    	ipv6_prefix = 32
    	ipv6_enforce_eui = false
	}
    `, device)
}

func testAccCheckFmcPhysicalInterfacesCreateExists(n string) resource.TestCheckFunc {
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
