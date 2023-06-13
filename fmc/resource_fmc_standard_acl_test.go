package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcStandardAclBasic(t *testing.T) {
	nwObjName := "any-ipv4"
	aclName := "Test1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcStandardAclDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcStandardAclConfigBasic(nwObjName, aclName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcStandardAclExists("fmc_standard_acl.test"),
				),
			},
		},
	})
}

func testAccCheckFmcStandardAclDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_standard_acl" {
			continue
		}

		id := rs.Primary.ID
		ctx := context.Background()
		err := c.DeleteFmcStandardAcl(ctx, id)

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "404") {
			return err
		}
	}

	return nil
}

func testAccCheckFmcStandardAclConfigBasic(nwObjName string, aclName string) string {
	return fmt.Sprintf(`
	data "fmc_network_objects" "any" {
		name = "%s"
	}

	resource "fmc_standard_acl" "test" {
		name = "%s"
		action = "DENY"
		object_id = data.fmc_network_objects.any.id
	}
    `, nwObjName, aclName)
}

func testAccCheckFmcStandardAclExists(n string) resource.TestCheckFunc {
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
