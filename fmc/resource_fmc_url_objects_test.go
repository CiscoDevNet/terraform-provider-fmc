package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcURLObjectBasic(t *testing.T) {
	name := "test_url_obj"
	url := "https://www.cisco.com/"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcURLObjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcURLObjectConfigBasic(name, url),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcURLObjectExists("fmc_url_objects.test"),
				),
			},
		},
	})
}

func testAccCheckFmcURLObjectDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_url_objects" {
			continue
		}

		id := rs.Primary.ID
		ctx := context.Background()
		err := c.DeleteFmcURLObject(ctx, id)

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "404") {
			return err
		}
	}

	return nil
}

func testAccCheckFmcURLObjectConfigBasic(name, url string) string {
	return fmt.Sprintf(`
    resource "fmc_url_objects" "test" {
		name = "%s"
		url = "%s"
	}
    `, name, url)
}

func testAccCheckFmcURLObjectExists(n string) resource.TestCheckFunc {
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
