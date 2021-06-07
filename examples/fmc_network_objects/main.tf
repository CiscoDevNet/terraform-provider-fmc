terraform {
  required_providers {
    fmc = {
      source = "hashicorp.com/cisco/fmc"
    }
  }
}

provider "fmc" {
  fmc_username = var.fmc_username
  fmc_password = var.fmc_password
  fmc_host = var.fmc_host
  fmc_insecure_skip_verify = var.fmc_insecure_skip_verify
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
