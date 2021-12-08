package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcDynamicObjectBasic(t *testing.T) {
	name := "test_dynamic_obj"
	objectType := "IP"
	description := "test network obj"
	descriptionUpdated := "test network object updated description"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcDynamicObjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcDynamicObjectConfigBasic(name, objectType, description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcDynamicObjectExists("fmc_dynamic_object.test", map[string]string{
						"name":        name,
						"description": description,
					}),
				),
			},
			{
				Config: testAccCheckFmcDynamicObjectConfigBasic(name, objectType, descriptionUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcDynamicObjectExists("fmc_dynamic_object.test", map[string]string{
						"name":        name,
						"description": descriptionUpdated,
					}),
				),
			},
		},
	})
}

func testAccCheckFmcDynamicObjectDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_dynamic_object" {
			continue
		}

		id := rs.Primary.ID
		ctx := context.Background()
		err := c.DeleteFmcDynamicObject(ctx, id)

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "404") {
			return err
		}
	}

	return nil
}

func testAccCheckFmcDynamicObjectConfigBasic(name, objectType, description string) string {
	return fmt.Sprintf(`
    resource "fmc_dynamic_object" "test" {
        name        = "%s"
        object_type = "%s"
        description = "%s"
    }
    `, name, objectType, description)
}

func testAccCheckFmcDynamicObjectExists(n string, properties map[string]string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID set")
		}

		if properties != nil {
			for key, value := range properties {
				if rs.Primary.Attributes[key] != value {
					return fmt.Errorf("attribute mismatch for key: %s. Expected: %s, got: %s", key, value, rs.Primary.Attributes[key])
				}
			}
		}

		//
		return nil
	}
}
