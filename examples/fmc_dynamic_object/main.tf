terraform {
  required_providers {
    fmc = {
      source = "CiscoDevNet/fmc"
      version = "0.2"
    }
  }
}

provider "fmc" {
  fmc_username = var.fmc_username
  fmc_password = var.fmc_password
  fmc_host = var.fmc_host
  fmc_insecure_skip_verify = var.fmc_insecure_skip_verify
}

data "fmc_dynamic_objects" "test"{
  name="DynamicObject"
}

resource "fmc_dynamic_objects" "test" {
  name        = "DynamicObject"
  object_type = "IP"
  description = "testing terraform"
}

output "value" {
  value=data.fmc_dynamic_objects.test
}
