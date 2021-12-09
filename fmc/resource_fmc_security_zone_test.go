package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcSecurityZoneBasic(t *testing.T) {
	name := "test"
	nameUpdated := "test1"
	interfaceMode := "INLINE"
	interfaceModeUpdated := "ASA"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcSecurityZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcSecurityZoneBasic(name, interfaceMode),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcSecurityZone("fmc_security_zone.test", map[string]string{
						"name":           name,
						"interface_mode": interfaceMode,
					}),
				),
			},
			{
				Config: testAccCheckFmcSecurityZoneBasic(nameUpdated, interfaceModeUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcSecurityZone("fmc_security_zone.test", map[string]string{
						"name":           nameUpdated,
						"interface_mode": interfaceModeUpdated,
					}),
				),
			},
		},
	})
}

func testAccCheckFmcSecurityZoneDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_security_zone" {
			continue
		}

		id := rs.Primary.ID
		ctx := context.Background()
		_, err := c.GetFmcSecurityZone(ctx, id)

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "404") {
			return err
		}
	}

	return nil
}

func testAccCheckFmcSecurityZoneBasic(name, interfaceMode string) string {
	return fmt.Sprintf(`
    resource "fmc_security_zone" "test" {
        name           = "%s"
        interface_mode = "%s"
    }
    `, name, interfaceMode)
}

func testAccCheckFmcSecurityZone(n string, properties map[string]string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		c := testAccProvider.Meta().(*Client)
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

		ctx := context.Background()
		_, err := c.GetFmcSecurityZone(ctx, rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("unable to read security zone from device: %s", err.Error())
		}
		return nil
	}
}
