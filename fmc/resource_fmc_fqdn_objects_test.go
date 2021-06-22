package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcFQDNObjectBasic(t *testing.T) {
	name := "test_fqdn_obj"
	value := "cisco.com"
	description := "test fqdn obj"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcFQDNObjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcFQDNObjectConfigBasic(name, value, description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcFQDNObjectExists("fmc_fqdn_objects.test"),
				),
			},
		},
	})
}

func testAccCheckFmcFQDNObjectDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_fqdn_objects" {
			continue
		}

		id := rs.Primary.ID
		ctx := context.Background()
		err := c.DeleteFmcFQDNObject(ctx, id)

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "404") {
			return err
		}
	}

	return nil
}

func testAccCheckFmcFQDNObjectConfigBasic(name, value, description string) string {
	return fmt.Sprintf(`
    resource "fmc_fqdn_objects" "test" {
        name        = "%s"
        value       = "%s"
        description = "%s"
		dns_resolution = "IPV4_ONLY"
    }
    `, name, value, description)
}

func testAccCheckFmcFQDNObjectExists(n string) resource.TestCheckFunc {
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
