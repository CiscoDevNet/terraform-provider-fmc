package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcICMPV4ObjectBasic(t *testing.T) {
	name := "test_network_obj"
	icmp_type := "3"
	code := 2

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcICMPV4ObjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcICMPV4ObjectConfigBasic(name, icmp_type, code),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcICMPV4ObjectExists("fmc_icmpv4_objects.test"),
				),
			},
		},
	})
}

func testAccCheckFmcICMPV4ObjectDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_icmpv4_objects" {
			continue
		}

		id := rs.Primary.ID
		ctx := context.Background()
		err := c.DeleteFmcICMPV4Object(ctx, id)

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "404") {
			return err
		}
	}

	return nil
}

func testAccCheckFmcICMPV4ObjectConfigBasic(name, icmp_type string, code int) string {
	return fmt.Sprintf(`
    resource "fmc_icmpv4_objects" "test" {
        name        = "%s"
        icmp_type  	= "%s"
  		code  		= %d
    }
    `, name, icmp_type, code)
}

func testAccCheckFmcICMPV4ObjectExists(n string) resource.TestCheckFunc {
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
