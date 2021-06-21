package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcNatPolicyBasic(t *testing.T) {
	name := "test_nat_policy"
	description := "test nat policy"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcNatPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcNatPolicyConfigBasic(name, description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcNatPolicyExists("fmc_ftd_nat_policies.test"),
				),
			},
		},
	})
}

func testAccCheckFmcNatPolicyDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_ftd_nat_policies" {
			continue
		}

		id := rs.Primary.ID
		ctx := context.Background()
		err := c.DeleteFmcNatPolicy(ctx, id)

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "404") {
			return err
		}
	}

	return nil
}

func testAccCheckFmcNatPolicyConfigBasic(name, description string) string {
	return fmt.Sprintf(`
    resource "fmc_ftd_nat_policies" "test" {
        name        = "%s"
        description = "%s"
    }
    `, name, description)
}

func testAccCheckFmcNatPolicyExists(n string) resource.TestCheckFunc {
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
