terraform {
  required_providers {
    fmc = {
      source = "hashicorp.com/edu/fmc"
    }
  }
}

provider "fmc" {
  fmc_username = "patrick"
  fmc_password = "CXsecurity!@34"
  fmc_host = "10.106.36.90"
  fmc_insecure_skip_verify = true
}

resource "fmc_host_objects" "test_host_1" {
  name        = "terraform_test_host_1"
  value       = "1.1.1.2"
  description = "testing terraform change"
}

output "test_host_1" {
  value = fmc_host_objects.test_host_1.id
}
