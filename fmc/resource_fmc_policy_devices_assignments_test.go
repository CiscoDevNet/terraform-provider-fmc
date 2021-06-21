package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcPolicyDeviceAssignmentsBasic(t *testing.T) {
	policy := "FTD"
	device := "ftd.adyah.cisco"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcPolicyDeviceAssignmentsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcPolicyDeviceAssignmentsConfigBasic(policy, device),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcPolicyDeviceAssignmentsExists("fmc_policy_devices_assignments.test"),
				),
			},
		},
	})
}

func testAccCheckFmcPolicyDeviceAssignmentsDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_policy_devices_assignments" {
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

func testAccCheckFmcPolicyDeviceAssignmentsConfigBasic(policy, device string) string {
	return fmt.Sprintf(`
	data "fmc_access_policies" "access_policy" {
		name = "%s"
	}
	data "fmc_devices" "device" {
		name = "%s"
	}
	resource "fmc_access_policies" "access_policy" {
		name = "Pre Test Access Policy"
		default_action = "block"
	}
	resource "fmc_policy_devices_assignments" "pre-test" {
		policy {
			id = fmc_access_policies.access_policy.id
			type = fmc_access_policies.access_policy.type
		}
		target_devices {
			id = data.fmc_devices.device.id
			type = data.fmc_devices.device.type
		}
	}
	resource "fmc_policy_devices_assignments" "test" {
		policy {
			id = data.fmc_access_policies.access_policy.id
			type = data.fmc_access_policies.access_policy.type
		}
		target_devices {
			id = data.fmc_devices.device.id
			type = data.fmc_devices.device.type
		}
		depends_on = [ fmc_policy_devices_assignments.pre-test ]
	}
    `, policy, device)
}

func testAccCheckFmcPolicyDeviceAssignmentsExists(n string) resource.TestCheckFunc {
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
