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
  name        = "${data.fmc_network_objects.PrivateVLAN.name}-DRsite"
  value       = data.fmc_network_objects.PrivateVLAN.value
  description = "testing terraform"
}

output "existing_fmc_network_object" {
  value = data.fmc_network_objects.PrivateVLAN
}

output "new_fmc_network_object_3" {
  value = fmc_network_objects.PrivateVLANDR
}
