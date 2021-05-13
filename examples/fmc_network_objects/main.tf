terraform {
  required_providers {
    fmc = {
      source = "hashicorp.com/edu/fmc"
    }
  }
}

resource "fmc_network_objects" "new" {
  name        = "terraform_network_object_100"
  value       = "1.0.0.0/24"
  description = "just"
  type        = "Network"
}

resource "fmc_network_objects" "new_3" {
  name        = "terraform_network_object_111"
  value       = "1.0.0.0/24"
  description = "My First Terraform Network Object"
  type        = "Network"
}

output "new_fmc_network_object" {
  value = fmc_network_objects.new
}
output "new_fmc_network_object_3" {
  value = fmc_network_objects.new_3
}
