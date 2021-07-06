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

data "fmc_ips_policies" "ips_policy" {
    name = "Connectivity Over Security"
}

output "existing_ips_policy" {
    value = data.fmc_ips_policies.ips_policy
}
