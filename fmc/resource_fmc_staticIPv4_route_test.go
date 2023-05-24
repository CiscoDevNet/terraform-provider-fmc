package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcStaticIPv4RouteBasic(t *testing.T) {
	device := "ftd.adyah.cisco"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcStaticIPv4RouteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcStaticIPv4RouteConfigBasic(device),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcStaticIPv4RouteExists("fmc_staticIPv4_route.route"),
				),
			},
		},
	})
}

func testAccCheckFmcStaticIPv4RouteDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_staticIPv4_route" {
			continue
		}

		id := rs.Primary.ID
		ctx := context.Background()
		_, err := c.GetFmcStaticIPv4Route(ctx, rs.Primary.Attributes["device_id"], id)

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "404") {
			return err
		}
	}

	return nil
}

func testAccCheckFmcStaticIPv4RouteConfigBasic(device string) string {
	return fmt.Sprintf(`
	data "fmc_devices" "device" {
		name = "%s"
	}

	data "fmc_device_physical_interfaces" "zero_physical_interface" {
		device_id = data.fmc_devices.device.id
		name = "TenGigabitEthernet0/0"
	}
	
	resource "fmc_host_objects" "aws_meta" {
	  name        = "aws_metadata_server"
	  value       = "169.254.169.254"
	}
	resource "fmc_host_objects" "inside-gw" {
	  name        = "inside-gateway1"
	  value       = "172.0.0.1"
	}

	resource "fmc_security_zone" "outside" {
	  name            = "outside"
	  interface_mode  = "ROUTED"
	}
	resource "fmc_device_physical_interfaces" "physical_interfaces00" {
		enabled = true
		device_id = data.fmc_devices.device.id
		physical_interface_id= data.fmc_device_physical_interfaces.zero_physical_interface.id
		name =   data.fmc_device_physical_interfaces.zero_physical_interface.name
		security_zone_id= fmc_security_zone.outside.id
		if_name = "outside"
		description = "Applied by terraform"
		mtu =  1900
		mode = "NONE"
		ipv4_dhcp_enabled = true
		ipv4_dhcp_route_metric = 1
	}
	
	resource "fmc_staticIPv4_route" "route" {
	  depends_on = [fmc_device_physical_interfaces.physical_interfaces00]
	  metric_value = 25
	  device_id  = data.fmc_devices.device.id
	  interface_name = fmc_device_physical_interfaces.physical_interfaces00.if_name
	  selected_networks {
		  id = fmc_host_objects.aws_meta.id
		  type = fmc_host_objects.aws_meta.type
		  name = fmc_host_objects.aws_meta.name
	  }
	  gateway {
		object {
		  id   = fmc_host_objects.inside-gw.id
		  type = fmc_host_objects.inside-gw.type
		  name = fmc_host_objects.inside-gw.name
		}
	  }
	}
    `, device)
}

func testAccCheckFmcStaticIPv4RouteExists(n string) resource.TestCheckFunc {
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
