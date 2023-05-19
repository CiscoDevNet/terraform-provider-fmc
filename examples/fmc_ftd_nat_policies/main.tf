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

data "fmc_nat_policies" "existing-nat-policy"{
  name="test"
}

resource "fmc_ftd_nat_policies" "nat_policy" {
    name = "Terraform NAT Policy"
    description = "New NAT policy!"
}

output "new_ftd_nat_policy" {
    value = fmc_ftd_nat_policies.nat_policy
}

output "existing-nat" {
    value = data.fmc_nat_policies.existing-nat-policy
}