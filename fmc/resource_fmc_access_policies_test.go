package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcAccessPolicyBasic(t *testing.T) {
	name := "test_access_policy"
	default_action := "block"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcAccessPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcAccessPolicyConfigBasic(name, default_action),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcAccessPolicyExists("fmc_access_policies.test"),
				),
			},
		},
	})
}

func testAccCheckFmcAccessPolicyDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_access_policies" {
			continue
		}

		id := rs.Primary.ID
		ctx := context.Background()
		err := c.DeleteFmcAccessPolicy(ctx, id)

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "404") {
			return err
		}
	}

	return nil
}

func testAccCheckFmcAccessPolicyConfigBasic(name, default_action string) string {
	return fmt.Sprintf(`
    resource "fmc_access_policies" "test" {
        name        = "%s"
        default_action = "%s"
    }
    `, name, default_action)
}

func testAccCheckFmcAccessPolicyExists(n string) resource.TestCheckFunc {
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
