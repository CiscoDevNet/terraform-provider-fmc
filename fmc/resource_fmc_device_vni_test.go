package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcVNICreateBasic(t *testing.T) {
	device := "ftd.adyah.cisco"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcVNICreateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcVNICreateConfigBasic(device),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcVNICreateExists("fmc_device_vni.test"),
				),
			},
		},
	})
}

func testAccCheckFmcVNICreateDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_device_vni" {
			continue
		}

		id := rs.Primary.ID
		ctx := context.Background()
		_, err := c.DeleteFmcVNIDetails(ctx, rs.Primary.Attributes["device_id"], id)

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "404") {
			return err
		}
	}

	return nil
}

func testAccCheckFmcVNICreateConfigBasic(device string) string {
	return fmt.Sprintf(`
	data "fmc_devices" "device" {
		name = "%s"
	}
	
	resource "fmc_security_zone" "my_security_zone" {
	  name = "vni-zone"
	  interface_mode = "ROUTED"
	}
	
	resource "fmc_device_vni" "test" {
		device_id = data.fmc_devices.device.id
		if_name = "VNI1"
		description = "Description Updated"
		security_zone_id= fmc_security_zone.my_security_zone.id
		priority = 3
		vnid = 11 
		multicast_groupaddress =  "224.0.0.34"
		segment_id = 4011
		enable_proxy = false
		ipv4 {
		   static {
			address = "3.3.3.3"
			netmask = 4
		  }
		   dhcp {
			enable_default_route_dhcp = false
			dhcp_route_metric = 0
		  }
		}
	}
    `, device)
}

func testAccCheckFmcVNICreateExists(n string) resource.TestCheckFunc {
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
