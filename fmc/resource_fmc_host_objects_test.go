package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcHostObjectBasic(t *testing.T) {
	name := "test_host_obj"
	value := "1.1.1.1"
	description := "test host obj"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcHostObjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcHostObjectConfigBasic(name, value, description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcHostObjectExists("fmc_host_objects.test"),
				),
			},
		},
	})
}

func testAccCheckFmcHostObjectDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_host_objects" {
			continue
		}

		id := rs.Primary.ID
		ctx := context.Background()
		err := c.DeleteFmcHostObject(ctx, id)

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "404") {
			return err
		}
	}

	return nil
}

func testAccCheckFmcHostObjectConfigBasic(name, value, description string) string {
	return fmt.Sprintf(`
    resource "fmc_host_objects" "test" {
        name        = "%s"
        value       = "%s"
        description = "%s"
    }
    `, name, value, description)
}

func testAccCheckFmcHostObjectExists(n string) resource.TestCheckFunc {
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
