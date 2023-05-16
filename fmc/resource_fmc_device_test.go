package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcDevicesBasic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcDevicesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcDevicesConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcDevicesExists("fmc_devices.device"),
				),
			},
		},
	})
}

func testAccCheckFmcDevicesDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_devices" {
			continue
		}

		id := rs.Primary.ID
		ctx := context.Background()
		err := c.DeleteFmcDevice(ctx, id)

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "404") {
			return err
		}
	}

	return nil
}

func testAccCheckFmcDevicesConfigBasic() string {
	return fmt.Sprintf(`
	data "fmc_access_policies" "access_policy" {
		name = "test-acp"
	}

	resource "fmc_devices" "device"{
		name = "FTD"
		hostname = "13.208.234.220"
		regkey = "cisco"
		type = "Device"
		license_caps = [ "MALWARE"]
		access_policy {
			id = data.fmc_access_policies.access_policy.id
			type = data.fmc_access_policies.access_policy.type
		}
	}
    `)
}

func testAccCheckFmcDevicesExists(n string) resource.TestCheckFunc {
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
