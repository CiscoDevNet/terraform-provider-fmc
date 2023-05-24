package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcExtendedAclBasic(t *testing.T) {
	hostObjectName := "acl-host-object"
	hostObjectValue := "10.10.0.2"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcExtendedAclDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcExtendedAclConfigBasic(hostObjectName, hostObjectValue),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcExtendedAclExists("fmc_extended_acl.new_acl"),
				),
			},
		},
	})
}

func testAccCheckFmcExtendedAclDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_extended_acl" {
			continue
		}

		id := rs.Primary.ID
		ctx := context.Background()
		err := c.DeleteFmcExtendedAcl(ctx, id)

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "404") {
			return err
		}
	}

	return nil
}

func testAccCheckFmcExtendedAclConfigBasic(hostObjectName, hostObjectValue string) string {
	return fmt.Sprintf(`
	resource "fmc_host_objects" "test"{
		name="%s"
		value="%s"
	  }
	  
	  data "fmc_port_objects" "test"{
		name="HTTP"
	  }
	  
	  resource "fmc_extended_acl" "new_acl" {
		name = "new_acl_test-2"
		action = "DENY"//"PERMIT"//
		log_level = "ERROR"//"INFORMATIONAL"
		logging="PER_ACCESS_LIST_ENTRY"
		log_interval=545
		source_port_object_id = data.fmc_port_objects.test.id
		source_port_literal_port="12311"
		source_port_literal_protocol="6"
	  
		source_network_object_id = fmc_host_objects.test.id
		source_network_literal_type="Host"
		source_network_literal_value="172.16.1.2"
	  
		destination_network_object_id = fmc_host_objects.test.id
		destination_network_literal_type="Host"
		destination_network_literal_value="172.16.1.2"
	  }`, hostObjectName, hostObjectValue)
}

func testAccCheckFmcExtendedAclExists(n string) resource.TestCheckFunc {
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
