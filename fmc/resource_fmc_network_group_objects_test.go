package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcNetworkGroupObjectBasic(t *testing.T) {
	net1 := "1.1.1.0/24"
	net2 := "1.1.2.0/24"
	name := "test_network_group"
	literal := "1.1.1.1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcNetworkGroupObjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcNetworkGroupObjectConfigBasic(net1, net2, name, literal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcNetworkGroupObjectExists("fmc_network_group_objects.test"),
				),
			},
		},
	})
}

func testAccCheckFmcNetworkGroupObjectDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_network_group_objects" {
			continue
		}

		id := rs.Primary.ID
		ctx := context.Background()
		err := c.DeleteFmcNetworkGroupObject(ctx, id)

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "404") {
			return err
		}
	}

	return nil
}

func testAccCheckFmcNetworkGroupObjectConfigBasic(net1, net2, name, literal string) string {
	return fmt.Sprintf(`
	resource "fmc_network_objects" "test1" {
		name        = "test1"
		value       = "%s"
		description = "testing 1"
	}

	resource "fmc_network_objects" "test2" {
		name        = "test2"
		value       = "%s"
	}	  
    resource "fmc_network_group_objects" "test" {
		name = "%s"
		objects {
			id = fmc_network_objects.test1.id
			type = fmc_network_objects.test1.type
		}
		objects {
			id = fmc_network_objects.test2.id
			type = fmc_network_objects.test2.type
		}
		literals {
			value = "%s"
			type = "Host"
		}
	}
    `, net1, net2, name, literal)
}

func testAccCheckFmcNetworkGroupObjectExists(n string) resource.TestCheckFunc {
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
