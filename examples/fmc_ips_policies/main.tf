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

data "fmc_ips_policies" "ips_policy" {
    name = "Connectivity Over Security"
}

output "existing_ips_policy" {
    value = data.fmc_ips_policies.ips_policy
}
