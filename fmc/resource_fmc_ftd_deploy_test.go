package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcFtdDeployBasic(t *testing.T) {
	device := "ftd.adyah.cisco"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcFtdDeployDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcFtdDeployConfigBasic(device),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcFtdDeployExists("fmc_ftd_deploy.test"),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckFmcFtdDeployDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_ftd_deploy" {
			continue
		}

		id := rs.Primary.ID
		ctx := context.Background()
		err := c.DeleteFmcPortGroupObject(ctx, id)

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "404") {
			return err
		}
	}

	return nil
}

func testAccCheckFmcFtdDeployConfigBasic(device string) string {
	return fmt.Sprintf(`
	data "fmc_devices" "ftd" {
		name = "%s"
	}
	
	resource "fmc_ftd_deploy" "ftd" {
		device = data.fmc_devices.ftd.id
		ignore_warning = false
		force_deploy = false
	}
    `, device)
}

func testAccCheckFmcFtdDeployExists(n string) resource.TestCheckFunc {
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
