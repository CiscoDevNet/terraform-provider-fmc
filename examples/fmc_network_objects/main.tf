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

locals {
  object_names = [for i in range(5): "test-${i}"]
}

data "fmc_network_objects" "PrivateVLAN" {
  name = "VLAN825-Private"
}

resource "fmc_network_objects" "PrivateVLANDR" {
  name        = "${data.fmc_network_objects.PrivateVLAN.name}-DRsite"
  value       = data.fmc_network_objects.PrivateVLAN.value
  description = "testing terraform"
}

/* resource "fmc_network_objects_bulk" "tests" {
  dynamic "objects" {
    for_each = local.object_names

    content {
      name        = objects.value
      value       = data.fmc_network_objects.PrivateVLAN.value
      description = "testing terraform"
    }
  }
} */

resource "fmc_network_objects_bulk" "test" {
  objects {
    name        = "test-1"
    value       = data.fmc_network_objects.PrivateVLAN.value
    description = "testing terraform"
  }

  objects {
    name        = "test-2"
    value       = data.fmc_network_objects.PrivateVLAN.value
    description = "testing terraform"
  }

  objects {
    name        = "test-3"
    value       = data.fmc_network_objects.PrivateVLAN.value
    description = "testing terraform"
  }
}

output "existing_fmc_network_object" {
  value = data.fmc_network_objects.PrivateVLAN
}

output "new_fmc_network_object_3" {
  value = fmc_network_objects.PrivateVLANDR
}
