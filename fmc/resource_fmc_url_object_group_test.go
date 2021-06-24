package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcURLObjectGroupBasic(t *testing.T) {
	name := "test_url_group_obj"
	description := "Cisco URLs"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcURLObjectGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcURLObjectGroupConfigBasic(name, description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcURLObjectGroupExists("fmc_url_object_group.test"),
				),
			},
		},
	})
}

func testAccCheckFmcURLObjectGroupDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_url_object_group" {
			continue
		}

		id := rs.Primary.ID
		ctx := context.Background()
		err := c.DeleteFmcURLObjectGroup(ctx, id)

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "404") {
			return err
		}
	}

	return nil
}

func testAccCheckFmcURLObjectGroupConfigBasic(name, description string) string {
	return fmt.Sprintf(`
	resource "fmc_url_objects" "test1" {
		name        = "test_url_1"
		url         = "https://www.cisco.com/"
	}
	resource "fmc_url_objects" "test2" {
		name        = "test_url_2"
		url         = "https://cisco.com/"
		description = "Test URL 2"
	}
    resource "fmc_url_object_group" "test" {
		name = "%s"
		description = "%s"
		objects {
			id = fmc_url_objects.test1.id
			type = fmc_url_objects.test1.type
		}
		objects {
			id = fmc_url_objects.test2.id
			type = fmc_url_objects.test2.type
		}
		literals {
			url = "www.cisco.com"
		}
		literals {
			url = "cisco.com"
		}
	}
    `, name, description)
}

func testAccCheckFmcURLObjectGroupExists(n string) resource.TestCheckFunc {
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
