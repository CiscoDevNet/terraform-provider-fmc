package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcPrefilterPolicyBasic(t *testing.T) {
	name := "Another Test Prefilter Policy"
	description := "Terraform Prefilter Policy description"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcPrefilterPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcPrefilterPolicyConfigBasic(name, description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcPrefilterPolicyExists("fmc_prefilter_policy.prefilter_policy2"),
				),
			},
		},
	})
}

func testAccCheckFmcPrefilterPolicyDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_prefilter_policy" {
			continue
		}

		id := rs.Primary.ID
		ctx := context.Background()
		err := c.DeleteFmcPrefilterPolicy(ctx, id)

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "404") {
			return err
		}
	}

	return nil
}

func testAccCheckFmcPrefilterPolicyConfigBasic(name, description string) string {
	return fmt.Sprintf(`
		resource "fmc_prefilter_policy" "prefilter_policy2" {
		  name        		     = "%s"
		  description 		     = "%s"
		  default_action { 
			log_begin = false
			send_events_to_fmc = false
			action = "ANALYZE_TUNNELS" 
		  }
		}
    `, name, description)
}

func testAccCheckFmcPrefilterPolicyExists(n string) resource.TestCheckFunc {
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
