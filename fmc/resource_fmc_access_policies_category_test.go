package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcAccessPoliciesCategoryBasic(t *testing.T) {
	name := "test_access_policy"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcAccessPoliciesCategoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcAccessPoliciesCategoryConfigBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcAccessPoliciesCategoryExists("fmc_access_policies_category.access_policy_test_category"),
				),
			},
		},
	})
}

func testAccCheckFmcAccessPoliciesCategoryDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_access_policies_category" {
			continue
		}

		id := rs.Primary.ID
		accessPolicyId := rs.Primary.Attributes["access_policy_id"]
		ctx := context.Background()
		err := c.DeleteFmcAccessPoliciesCategory(ctx, id, accessPolicyId)

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "404") {
			return err
		}
	}

	return nil
}

func testAccCheckFmcAccessPoliciesCategoryConfigBasic(name string) string {
	return fmt.Sprintf(`
resource "fmc_access_policies" "access_policy" {
    name = "Terraform Access Policy for Category Testing"
    default_action = "block" 
}

resource "fmc_access_policies_category" "access_policy_test_category" {
    name             = "%s"
    access_policy_id = fmc_access_policies.access_policy.id
}
    `, name)
}

func testAccCheckFmcAccessPoliciesCategoryExists(n string) resource.TestCheckFunc {
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
