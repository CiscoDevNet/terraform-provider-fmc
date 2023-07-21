package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcIpsPolicyBasic(t *testing.T) {
	name := "test_ips_policy"
	//description := "test ips policy"
	inspection_mode := "DETECTION"
	basepolicy_id := "6c66b83c-bc23-55b6-879d-c4d847443503"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcIpsPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcIpsPolicyConfigBasic(name, inspection_mode, basepolicy_id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcIpsPolicyExists("fmc_ips_policies.test"),
				),
			},
		},
	})
}

func testAccCheckFmcIpsPolicyDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_ips_policies" {
			continue
		}

		id := rs.Primary.ID
		ctx := context.Background()
		err := c.DeleteFmcIPSPolicy(ctx, id)

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "404") {
			return err
		}
	}

	return nil
}

func testAccCheckFmcIpsPolicyConfigBasic(name, inspection_mode, basepolicy_id string) string {
	return fmt.Sprintf(`
    resource "fmc_ips_policies" "test" {
        name        = "%s"
        inspection_mode = "%s"
		basepolicy_id = "%s"
    }
    `, name, inspection_mode, basepolicy_id)
}

func testAccCheckFmcIpsPolicyExists(n string) resource.TestCheckFunc {
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
