terraform {
  required_providers {
    fmc = {
      source = "CiscoDevNet/fmc"
      # version = "0.1.1"
    }
  }
}

provider "fmc" {
  fmc_username             = var.fmc_username
  fmc_password             = var.fmc_password
  fmc_host                 = var.fmc_host
  fmc_insecure_skip_verify = var.fmc_insecure_skip_verify
}

data "fmc_extended_acl" "test" {
  name = "test-extended-acl"
}

data "fmc_host_objects" "test" {
  name = "test-object"
}

data "fmc_port_objects" "test" {
  name = "HTTP"
}

resource "fmc_extended_acl" "new_acl" {
  name         = "new_acl_test"
  action       = "DENY"  //"PERMIT"//
  log_level    = "ERROR" //"INFORMATIONAL"
  logging      = "PER_ACCESS_LIST_ENTRY"
  log_interval = 545

  source_port_object_id        = data.fmc_port_objects.test.id
  source_port_literal_port     = "12311"
  source_port_literal_protocol = "6"

  destination_port_object_id        = data.fmc_port_objects.test.id
  destination_port_literal_port     = "12311"
  destination_port_literal_protocol = "6"

  source_network_object_id     = data.fmc_host_objects.test.id
  source_network_literal_type  = "Host"
  source_network_literal_value = "172.16.1.2"

  destination_network_object_id     = data.fmc_host_objects.test.id
  destination_network_literal_type  = "Host"
  destination_network_literal_value = "172.16.1.2"
}

output "existing_acl" {
  value = data.fmc_extended_acl.test
}
