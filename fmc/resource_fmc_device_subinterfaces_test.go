package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcSubInterfaceCreateBasic(t *testing.T) {
	device := "ftd.adyah.cisco"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcSubInterfaceCreateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcSubInterfaceCreateConfigBasic(device),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcSubInterfaceCreateExists("fmc_device_subinterfaces.sub"),
				),
			},
		},
	})
}

func testAccCheckFmcSubInterfaceCreateDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_device_subinterfaces" {
			continue
		}

		id := rs.Primary.ID
		ctx := context.Background()
		err := c.DeleteFmcSubInterface(ctx, rs.Primary.Attributes["device_id"], id)

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "404") {
			return err
		}
	}

	return nil
}

func testAccCheckFmcSubInterfaceCreateConfigBasic(device string) string {
	return fmt.Sprintf(`
	data "fmc_devices" "device" {
		name = "%s"
	}
	resource "fmc_security_zone" "outside" {
		name            = "outside"
		interface_mode  = "ROUTED"
	}
	    
	resource "fmc_device_subinterfaces" "sub" {
		device_id   = data.fmc_devices.device.id
		ifname = "Testing1"
		subinterface_id = 12345
		vlan_id = 80
		name = "TenGigabitEthernet0/1"
		mode = "NONE"
		mtu = 1600
		enabled = true
		priority = 69
		security_zone_id = fmc_security_zone.outside.id
		ipv4_dhcp_enabled = false
		ipv4_dhcp_route_metric = 1
		enable_ipv6 = true
		ipv6_address = "2001:10:240:ac::a"
		ipv6_prefix = "124"
		ipv6_enforce_eui = false
	}
    `, device)
}

func testAccCheckFmcSubInterfaceCreateExists(n string) resource.TestCheckFunc {
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
