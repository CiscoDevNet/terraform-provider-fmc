package fmc

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcDynamicObjectMappingBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcDynamicObjectMappingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcDynamicObjectMappingsConfigBasic("8.8.8.8"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResource("fmc_dynamic_object_mapping.test", map[string]string{
						"mappings": "8.8.8.8",
					}),
				),
			},
		},
	})
}

func testAccCheckFmcDynamicObjectMappingDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_dynamic_object_mapping" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Write Memory resource ID was not set")
		}
	}

	return nil
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
        mappings = "%s"
    }

    `, name, mappings)
}


