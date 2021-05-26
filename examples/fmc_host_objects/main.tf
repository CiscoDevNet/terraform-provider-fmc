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

data "fmc_host_objects" "existing_host_1" {
  name = "CUCM-Pub"
}
resource "fmc_host_objects" "test_host_2" {
  name        = "terraform_test_host_2"
  value       = "1.1.1.2"
  description = "testing terraform change"
}

output "test_host_1" {
  value = fmc_host_objects.test_host_2
}

output "test_exiting_host" {
  value = data.fmc_host_objects.existing_host_1
}