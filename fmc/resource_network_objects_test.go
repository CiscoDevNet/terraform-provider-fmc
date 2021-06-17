package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcNetworkObjectBasic(t *testing.T) {
	name := "test_network_obj"
	value := "1.1.1.0/24"
	description := "test network obj"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcNetworkObjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcNetworkObjectConfigBasic(name, value, description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcNetworkObjectExists("fmc_network_objects.test"),
				),
			},
		},
	})
}

func testAccCheckFmcNetworkObjectDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_network_objects" {
			continue
		}

		id := rs.Primary.ID
		ctx := context.Background()
		err := c.DeleteNetworkObject(ctx, id)

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "404") {
			return err
		}
	}

	return nil
}

func testAccCheckFmcNetworkObjectConfigBasic(name, value, description string) string {
	return fmt.Sprintf(`
    resource "fmc_network_objects" "test" {
        name        = "%s"
        value       = "%s"
        description = "%s"
    }
    `, name, value, description)
}

func testAccCheckFmcNetworkObjectExists(n string) resource.TestCheckFunc {
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
