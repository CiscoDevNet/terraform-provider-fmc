package fmc

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcSmartLicenseBasic(t *testing.T) {
	registrationType := "EVALUATION"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcSmartLicenseDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcSmartLicenseBasic(registrationType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcSmartLicense("fmc_smart_license.test", map[string]string{
						"registration_type": registrationType,
					}),
				),
			},
		},
	})
}

func testAccCheckFmcSmartLicenseDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_smart_license" {
			continue
		}

		ctx := context.Background()
		_, err := c.GetFmcSmartLicense(ctx)

		if err != nil {
			return err
		}
	}

	return nil
}

func testAccCheckFmcSmartLicenseBasic(registrationType string) string {
	return fmt.Sprintf(`
    resource "fmc_smart_license" "test" {
        registration_type	= "%s"
    }
    `, registrationType)
}

func testAccCheckFmcSmartLicense(n string, properties map[string]string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		c := testAccProvider.Meta().(*Client)
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID set")
		}

		for key, value := range properties {
			if rs.Primary.Attributes[key] != value {
				return fmt.Errorf("attribute mismatch for key: %s. Expected: %s, got: %s", key, value, rs.Primary.Attributes[key])
			}
		}

		ctx := context.Background()
		_, err := c.GetFmcSmartLicense(ctx)

		if err != nil {
			return fmt.Errorf("unable to read smart license: %s", err.Error())
		}
		return nil
	}
}
