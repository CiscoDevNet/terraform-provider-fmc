package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcAutoNatRuleBasic(t *testing.T) {
	name := "Test NAT Rule Policy"
	description := "test autonat rule"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcAutoNatRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcAutoNatRuleConfigBasic(name, description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcAutoNatRuleExists("fmc_ftd_autonat_rules.test"),
				),
			},
		},
	})
}

func testAccCheckFmcAutoNatRuleDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_ftd_autonat_rules" {
			continue
		}

		id := rs.Primary.ID
		ctx := context.Background()
		err := c.DeleteFmcAutoNatRule(ctx, rs.Primary.Attributes["nat_policy"], id)

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "404") {
			return err
		}
	}

	return nil
}

func testAccCheckFmcAutoNatRuleConfigBasic(name, description string) string {
	return fmt.Sprintf(`
	resource "fmc_network_objects" "test" {
        name        = "test_auto_nat_network_obj"
        value       = "10.10.10.0/24"
        description = "Testing"
    }

    resource "fmc_ftd_nat_policies" "nat_policy" {
		name = "%s"
		description = "Test NAT policy!"
	}
	
	resource "fmc_ftd_autonat_rules" "test" {
		nat_policy = fmc_ftd_nat_policies.nat_policy.id
		description = "%s"
		nat_type = "static"
		original_network {
			id = fmc_network_objects.test.id
			type = fmc_network_objects.test.type
		}
		translated_network {
			id = fmc_network_objects.test.id
			type = fmc_network_objects.test.type
		}
	}
    `, name, description)
}

func testAccCheckFmcAutoNatRuleExists(n string) resource.TestCheckFunc {
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
