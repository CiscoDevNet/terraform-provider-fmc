terraform {
  required_providers {
    fmc = {
      source = "CiscoDevNet/fmc"
      # version = "0.2.0"
    }
  }
}

provider "fmc" {
  fmc_username = var.fmc_username
  fmc_password = var.fmc_password
  fmc_host = var.fmc_host
  fmc_insecure_skip_verify = var.fmc_insecure_skip_verify
}

data "fmc_ravpn_policy" "ravpn" {
    name = "ravpn1"
}

output "ravpn_policy" {
    value = data.fmc_ravpn_policy.ravpn
}
