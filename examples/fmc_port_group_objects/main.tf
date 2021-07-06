terraform {
  required_providers {
    fmc = {
      source = "CiscoDevNet/fmc"
      # version = "0.1.1"
    }
  }
}

provider "fmc" {
  fmc_username = var.fmc_username
  fmc_password = var.fmc_password
  fmc_host = var.fmc_host
  fmc_insecure_skip_verify = var.fmc_insecure_skip_verify
}

resource "fmc_port_objects" "shbharti_port_1" {
  name        = "shbharti_test_port_object_1"
  port        = "3943"
  protocol    = "TCP"
  description = "testing terraform"
  overridable = false
}

resource "fmc_icmpv4_objects" "shbharti-icmpv4-1" {
  name        = "shbharti-icmpv4-1"
  icmp_type = "3"
  code  = 2
}

resource "fmc_port_group_objects" "TestPortGroup" {
  name = "TestPortGroup"
  description = "Testing groups"
  objects {
      id = fmc_port_objects.shbharti_port_1.id
      type = fmc_port_objects.shbharti_port_1.type
  }
  objects {
      id = fmc_icmpv4_objects.shbharti-icmpv4-1.id
      type = fmc_icmpv4_objects.shbharti-icmpv4-1.type
  }
}

output "new_fmc_port_group_object" {
  value = fmc_port_group_objects.TestPortGroup
}