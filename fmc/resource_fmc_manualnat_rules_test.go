package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcManualNatRuleBasic(t *testing.T) {
	name := "Test NAT Rule Policy"
	description := "test manualnat rule"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcManualNatRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcManualNatRuleConfigBasic(name, description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcManualNatRuleExists("fmc_ftd_manualnat_rules.test"),
				),
			},
		},
	})
}

func testAccCheckFmcManualNatRuleDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_ftd_manualnat_rules" {
			continue
		}

		id := rs.Primary.ID
		ctx := context.Background()
		err := c.DeleteFmcManualNatRule(ctx, rs.Primary.Attributes["nat_policy"], id)

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "404") {
			return err
		}
	}

	return nil
}

func testAccCheckFmcManualNatRuleConfigBasic(name, description string) string {
	return fmt.Sprintf(`
	resource "fmc_network_objects" "test" {
        name        = "test_manual_nat_network_obj"
        value       = "10.10.10.0/24"
        description = "Testing"
    }

    resource "fmc_ftd_nat_policies" "nat_policy" {
		name = "%s"
		description = "Test NAT policy!"
	}
	
	resource "fmc_ftd_manualnat_rules" "test" {
		nat_policy = fmc_ftd_nat_policies.nat_policy.id
		description = "%s"
		nat_type = "static"
		original_source {
			id = fmc_network_objects.test.id
			type = fmc_network_objects.test.type
		}
        translated_source {
			id = fmc_network_objects.test.id
			type = fmc_network_objects.test.type
		}
	}
    `, name, description)
}

func testAccCheckFmcManualNatRuleExists(n string) resource.TestCheckFunc {
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
