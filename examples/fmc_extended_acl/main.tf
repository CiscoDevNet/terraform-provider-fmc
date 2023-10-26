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

# data "fmc_extended_acl" "test" {
#   name = "test-extended-acl"
# }

data "fmc_network_objects" "nw1" {
  name = "any-ipv4"
}

data "fmc_network_objects" "nw2" {
  name = "IPv4-Link-Local"
}

data "fmc_network_objects" "nw3" {
  name = "IPv4-Multicast"
}

data "fmc_port_objects" "port1" {
  name = "HTTP"
}

data "fmc_port_objects" "port2" {
  name = "HTTPS"
}

data "fmc_port_objects" "port3" {
  name = "FTP"
}

resource "fmc_extended_acl" "new_acl" {
  name         = "new_acl_test"
  action       = "BLOCK" //"PERMIT"//
  log_level    = "ERROR" //"INFORMATIONAL"
  logging      = "PER_ACCESS_LIST_ENTRY"
  log_interval = 545

  source_port_object_ids       = [data.fmc_port_objects.port1.id, data.fmc_port_objects.port2.id]
  source_port_literal_port     = "12311"
  source_port_literal_protocol = "6"

  destination_port_object_ids       = [data.fmc_port_objects.port2.id, data.fmc_port_objects.port3.id]
  destination_port_literal_port     = "12311"
  destination_port_literal_protocol = "6"

  source_network_object_ids    = [data.fmc_network_objects.nw2.id]
  source_network_literal_type  = "Host"
  source_network_literal_value = "172.16.1.2"

  destination_network_object_ids    = [data.fmc_network_objects.nw1.id, data.fmc_network_objects.nw2.id]
  destination_network_literal_type  = "Host"
  destination_network_literal_value = "172.16.1.2"
}

output "new_acl" {
  value = fmc_extended_acl.new_acl
}
