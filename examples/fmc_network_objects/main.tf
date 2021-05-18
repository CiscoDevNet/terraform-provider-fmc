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

resource "fmc_network_objects" "new" {
  name        = "terraform_network_object_204"
  value       = "1.0.0.0/24"
  description = "just"
}

resource "fmc_network_objects" "new_3" {
  name        = "terraform_network_object_205"
  value       = "1.0.0.0/24"
  description = "My First Terraform Network Object"
}

output "new_fmc_network_object" {
  value = fmc_network_objects.new.id
}
output "new_fmc_network_object_3" {
  value = fmc_network_objects.new_3.id
}
