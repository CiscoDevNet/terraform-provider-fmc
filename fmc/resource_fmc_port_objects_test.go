package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcPortObjectBasic(t *testing.T) {
	name := "test_port_obj"
	port := "1212"
	protocol := "TCP"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcPortObjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcPortObjectConfigBasic(name, port, protocol),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcPortObjectExists("fmc_port_objects.test"),
				),
			},
		},
	})
}

func testAccCheckFmcPortObjectDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_port_objects" {
			continue
		}

		id := rs.Primary.ID
		ctx := context.Background()
		err := c.DeleteFmcPortObject(ctx, id)

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "404") {
			return err
		}
	}

	return nil
}

func testAccCheckFmcPortObjectConfigBasic(name, port, protocol string) string {
	return fmt.Sprintf(`
    resource "fmc_port_objects" "test" {
		name = "%s"
		port = "%s"
		protocol = "%s"
	}	  
    `, name, port, protocol)
}

func testAccCheckFmcPortObjectExists(n string) resource.TestCheckFunc {
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
