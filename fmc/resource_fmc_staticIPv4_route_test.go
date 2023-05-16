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

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcStaticIPv4RouteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcStaticIPv4RouteConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcStaticIPv4RouteExists("fmc_staticIPv4_route.route1"),
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
		_ ,err := c.GetFmcStaticIPv4Route(ctx, rs.Primary.Attributes["device_id"], id)

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "404") {
			return err
		}
	}

	return nil
}

func testAccCheckFmcStaticIPv4RouteConfigBasic() string {
	return fmt.Sprintf(`
	data "fmc_devices" "device1" {
		name = "FTD1"
	}

	resource "fmc_network_objects" "test" {
        name        = "tes_object"
        value       = "10.10.10.0/24"
        description = "Testing"
    }

	resource "fmc_host_objects" "igw" {
        name        = "igw-kds"
        value       = "10.10.10.1"
    }

	resource "fmc_staticIPv4_route" "route1" {
		device_id  = data.fmc_devices.device1.id
		interface_name = "inside"
		metric_value = 23
		is_tunneled = false
		selected_networks {
			id = fmc_network_objects.test.id
			type = fmc_network_objects.test.type
			name =fmc_network_objects.test.name
		}
		gateway {
		  object {
			id = fmc_host_objects.igw.id
			type = fmc_host_objects.igw.type
			name = fmc_host_objects.igw.name
		  }
		}
	}
    `)
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
