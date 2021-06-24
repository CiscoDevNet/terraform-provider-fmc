package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcPortGroupObjectsBasic(t *testing.T) {
	name := "test_port_group_obj"
	description := "Sample Ports"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcPortGroupObjectsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcPortGroupObjectsConfigBasic(name, description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcPortGroupObjectsExists("fmc_port_group_objects.test"),
				),
			},
		},
	})
}

func testAccCheckFmcPortGroupObjectsDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_port_group_objects" {
			continue
		}

		id := rs.Primary.ID
		ctx := context.Background()
		err := c.DeleteFmcPortGroupObject(ctx, id)

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "404") {
			return err
		}
	}

	return nil
}

func testAccCheckFmcPortGroupObjectsConfigBasic(name, description string) string {
	return fmt.Sprintf(`
	resource "fmc_port_objects" "test1" {
		name        = "test_port_1"
		port        = "80"
		protocol 	= "TCP"
	}
	resource "fmc_icmpv4_objects" "test2" {
		name        = "test_port_2"
		icmp_type   = "3"
		code 		= 2
	}
    resource "fmc_port_group_objects" "test" {
		name = "%s"
		# description = "%s"
		objects {
			id = fmc_port_objects.test1.id
			type = fmc_port_objects.test1.type
		}
		objects {
			id = fmc_icmpv4_objects.test2.id
			type = fmc_icmpv4_objects.test2.type
		}
	}
    `, name, description)
}

func testAccCheckFmcPortGroupObjectsExists(n string) resource.TestCheckFunc {
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
