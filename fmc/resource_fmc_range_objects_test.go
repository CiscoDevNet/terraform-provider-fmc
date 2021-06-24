package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcRangeObjectBasic(t *testing.T) {
	name := "test_range_obj"
	value := "10.10.10.0-10.10.10.10"
	description := "Range test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcRangeObjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcRangeObjectConfigBasic(name, value, description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcRangeObjectExists("fmc_range_objects.test"),
				),
			},
		},
	})
}

func testAccCheckFmcRangeObjectDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_range_objects" {
			continue
		}

		id := rs.Primary.ID
		ctx := context.Background()
		err := c.DeleteFmcRangeObject(ctx, id)

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "404") {
			return err
		}
	}

	return nil
}

func testAccCheckFmcRangeObjectConfigBasic(name, value, description string) string {
	return fmt.Sprintf(`
    resource "fmc_range_objects" "test" {
		name = "%s"
		value = "%s"
		description = "%s"
	}	  
    `, name, value, description)
}

func testAccCheckFmcRangeObjectExists(n string) resource.TestCheckFunc {
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
