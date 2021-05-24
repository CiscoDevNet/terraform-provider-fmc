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

resource "fmc_ftd_nat_policies" "nat_policy" {
    name = "Terraform NAT Policy"
    description = "New NAT policy!"
}

output "new_ftd_nat_policy" {
    value = fmc_ftd_nat_policies.nat_policy
}
