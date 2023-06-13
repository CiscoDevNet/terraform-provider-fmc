package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcSGTObjectBasic(t *testing.T) {
	name := "sgt_objct-test"
	nameUpdated := "sgt_objct-updated"
	tag := "25"
	description := "Test SGT object"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcSGTObjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcSGTObjectConfigBasic(name, tag, description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcSGTObjectExists("fmc_sgt_objects.test"),
				),
			},
			{
				Config: testAccCheckFmcSGTObjectConfigBasic(nameUpdated, tag, description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcSGTObjectExists("fmc_sgt_objects.test"),
				),
			},
		},
	})
}

func testAccCheckFmcSGTObjectDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_sgt_objects" {
			continue
		}

		id := rs.Primary.ID
		ctx := context.Background()
		err := c.DeleteFmcSGTObjects(ctx, id)

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "400") {
			return err
		}
	}

	return nil
}

func testAccCheckFmcSGTObjectConfigBasic(name, tag string, description string) string {
	return fmt.Sprintf(`
    resource "fmc_sgt_objects" "test" {
		name = "%s"
		tag = "%s"
		description = "%s"
	}	  
    `, name, tag, description)
}

func testAccCheckFmcSGTObjectExists(n string) resource.TestCheckFunc {
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
