terraform {
  required_providers {
    fmc = {
      source = "hashicorp.com/edu/fmc"
    }
  }
}

provider "fmc" {
  fmc_username = "api"
  fmc_password = "CXsecurity!@34"
  fmc_host = "10.106.107.228"
  fmc_insecure_skip_verify = true
}

data "fmc_network_objects" "PrivateVLAN" {
  name = "VLAN825-Private"
}

resource "fmc_network_objects" "PrivateVLANDR" {
  name        = "DRsite-VLAN"
  value       = data.fmc_network_objects.PrivateVLAN.value
  description = "testing terraform"
}

resource "fmc_network_group_objects" "TestPrivateGroup" {
  name = "TestPrivateGroup"
  description = "Testing groups"
  objects {
      id = data.fmc_network_objects.PrivateVLAN.id
      type = data.fmc_network_objects.PrivateVLAN.type
  }
  objects {
      id = fmc_network_objects.PrivateVLANDR.id
      type = fmc_network_objects.PrivateVLANDR.type
  }
  literals {
      value = "10.10.10.10"
      type = "Host"
  }
}

output "existing_fmc_network_object" {
  value = data.fmc_network_objects.PrivateVLAN
}

output "new_fmc_network_object" {
  value = fmc_network_objects.PrivateVLANDR
}

output "new_fmc_network_group_object" {
  value = fmc_network_group_objects.TestPrivateGroup
}
