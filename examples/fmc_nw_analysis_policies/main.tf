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

resource "fmc_network_analysis_policy" "test" {
  name = "test-NWA-policy"
  description = "NWA-policy created by Terraform"
  base_policy {
    name = "Connectivity Over Security"
  }
  snort_engine = "SNORT3"
}

data "fmc_network_analysis_policy" "data-test" {
  name = fmc_network_analysis_policy.test.name
}

output "fmc_network_analysis_policy" {
  value = data.fmc_network_analysis_policy.data-test
}
