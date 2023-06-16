package fmc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFmcAccessRuleBasic(t *testing.T) {
	name := "FTD"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFmcAccessRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFmcAccessRuleConfigBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFmcAccessRuleExists("fmc_access_rules.access_rule_1"),
				),
			},
		},
	})
}

func testAccCheckFmcAccessRuleDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fmc_access_rules" {
			continue
		}

		id := rs.Primary.ID
		ctx := context.Background()
		err := c.DeleteFmcAccessRule(ctx, rs.Primary.Attributes["acp"], id)

		// Object is already deleted
		if err != nil && !strings.Contains(fmt.Sprint(err), "404") {
			return err
		}
	}

	return nil
}

func testAccCheckFmcAccessRuleConfigBasic(name string) string {
	return fmt.Sprintf(`
	data "fmc_access_policies" "access_policy" {
		name = "%s"
	}
	data "fmc_port_objects" "http" {
		name = "HTTPS"
	}
	
	data "fmc_ips_policies" "ips_policy" {
		name = "Connectivity Over Security"
	}
	resource "fmc_security_zone" "oustide" {
		name            = "outside"
		interface_mode  = "ROUTED"
	}
	resource "fmc_security_zone" "inside" {
		name            = "inside"
		interface_mode  = "ROUTED"
	}
	resource "fmc_network_objects" "source" {
		name        = "test-source"
		value       = "172.1.0.0/24"
		description = "testing terraform"
	}
	resource "fmc_network_objects" "dest" {
		name        = "test-dest"
		value       = "192.1.0.0/24"
		description = "testing terraform"
	}
	resource "fmc_access_rules" "access_rule_1" {
		acp = data.fmc_access_policies.access_policy.id
		section = "mandatory"
		name = "Test rule 1"
		action = "allow"
		enabled = true
		enable_syslog = true
		syslog_severity = "alert"
		send_events_to_fmc = true
		log_files = false
		log_end = true
		source_zones {
			source_zone {
				id = fmc_security_zone.inside.id
				type =  fmc_security_zone.inside.type
			}
		}
		destination_zones {
			destination_zone {
				id = fmc_security_zone.outside.id
				type =  fmc_security_zone.outside.type
			}
		}
		source_networks {
			source_network {
				id = fmc_network_objects.source.id
				type =  fmc_network_objects.source.type
			}
		}
		destination_networks {
			destination_network {
				id = fmc_network_objects.dest.id
				type = fmc_network_objects.dest.type
			}
		}
		destination_ports {
			destination_port {
				id = data.fmc_port_objects.http.id
				type =  data.fmc_port_objects.http.type
			}
		}
		ips_policy = data.fmc_ips_policies.ips_policy.id
		new_comments = [ "New", "comment" ]
	}
    `, name)
}

func testAccCheckFmcAccessRuleExists(n string) resource.TestCheckFunc {
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
