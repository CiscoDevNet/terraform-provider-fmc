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

resource "fmc_network_objects" "range1" {
  name        = "Sample_Range1"
  value       = "192.168.1.0/24"
  description = "Sample Range 1"
}

resource "fmc_network_objects" "range2" {
  name        = "Sample_Range2"
  value       = "192.168.2.0/24"
  description = "Sample Range 2"
}

resource "fmc_network_objects" "range3" {
  name        = "Sample_Range3"
  value       = "192.168.3.0/24"
  description = "Sample Range 3"
}

resource "fmc_network_group_objects" "group" {
  name = "Sample_Ntwk_Grp"
  description = "Testing Group objects"
  objects {
    id = fmc_network_objects.range1.id
    type = fmc_network_objects.range1.type
  }
  objects {
    id = fmc_network_objects.range3.id
    type = fmc_network_objects.range3.type
  }
}
