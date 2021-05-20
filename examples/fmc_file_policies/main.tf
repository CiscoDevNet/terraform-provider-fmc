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

data "fmc_file_policies" "file_policy" {
    name = "AMP Policy"
}

output "existing_file_policy" {
    value = data.fmc_file_policies.file_policy
}
