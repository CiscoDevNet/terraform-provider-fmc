package fmc

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcDynamicObjectMappingBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcDynamicObjectMappingDestroy([]string{"8.8.8.8"}),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcDynamicObjectMappingsConfigBasic("8.8.8.8"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcDynamicObjectMappingCreate([]string{"8.8.8.8"}),
				),
			},
		},
	})
}

func testAccCheckFmcDynamicObjectMappingCreate(mappings []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		c := testAccProvider.Meta().(*Client)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "fmc_dynamic_object_mapping" {
				continue
			}

			if rs.Primary.ID == "" {
				return fmt.Errorf("dynamic object mapping resource ID was not set")
			}

			// check if mapping was created for an object
			ctx := context.Background()
			dynamicObjectMappings, err := c.GetFmcDynamicObjectMapping(ctx, &DynamicObjectMapping{
				DynamicObject: DynamicObjectMappingObject{
					ID: rs.Primary.Attributes["dynamic_object_id"],
				},
				Mappings: mappings,
			})

			if err != nil {
				return fmt.Errorf("unable to read dynamic object: %s", err.Error())
			}

			if len(mappings) != len(dynamicObjectMappings.Mappings) {
				return fmt.Errorf("mapping was not created properly. Expected: %v, got: %v", mappings, dynamicObjectMappings.Mappings)
			}
		}

		return nil
	}
}

func testAccCheckFmcDynamicObjectMappingDestroy(mappings []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "fmc_dynamic_object_mapping" {
				continue
			}

			if rs.Primary.ID == "" {
				return fmt.Errorf("dynamic object mapping resource ID was not set")
			}
		}
		return nil
	}
}

func testAccCheckFmcDynamicObjectMappingsConfigBasic(mappings string) string {
	name := randomString(5)

	return fmt.Sprintf(`
    resource "fmc_dynamic_object" "test" {
        name        = "%s"
        object_type = "IP"
    }
	
    resource "fmc_dynamic_object_mapping" "test" {
        dynamic_object_id = fmc_dynamic_object.test.id
        mappings = ["%s"]
    }

    `, name, mappings)
}
