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
    name = "Balanced Security and Connectivity"
}

output "existing_ips_policy" {
    value = data.fmc_ips_policies.ips_policy
}

resource "fmc_ips_policies" "ips_policy"{
name="Test-ips-policy"
inspection_mode="DETECTION"
basepolicy_id=data.fmc_ips_policies.ips_policy.id
}