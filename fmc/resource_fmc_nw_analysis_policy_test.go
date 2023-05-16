package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcNWAnalysisPolicyBasic(t *testing.T) {
	name := "test_nw_analysis_policy"
	description := "Terraform Policy description"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcNWAnalysisPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcNWAnalysisPolicyConfigBasic(name, description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcNWAnalysisPolicyExists("fmc_network_analysis_policy.test"),
				),
			},
		},
	})
}

func testAccCheckFmcNWAnalysisPolicyDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_network_analysis_policy" {
			continue
		}

		id := rs.Primary.ID
		ctx := context.Background()
		err := c.DeleteFmcNetworkAnalysisPolicy(ctx, id)

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "404") {
			Log.Debug("Does it error here?")
			return err
		}
	}

	Log.Debug("FMC Network Analysis Policy destroyed.")
	return nil
}

func testAccCheckFmcNWAnalysisPolicyConfigBasic(name, description string) string {
	return fmt.Sprintf(`
		resource "fmc_network_analysis_policy" "test" {
		  name        		     = "%s"
		  description 		     = "%s"
		  base_policy { 
			name = "Connectivity Over Security"
		  }
		  snort_engine = "SNORT3"
		}
    `, name, description)
}

func testAccCheckFmcNWAnalysisPolicyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			Log.Debug("Does it error here123?")
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			Log.Debug("Does it error here456?")
			return fmt.Errorf("No ID set")
		}
		Log.Debug(rs, ok)
		return nil
	}
}
