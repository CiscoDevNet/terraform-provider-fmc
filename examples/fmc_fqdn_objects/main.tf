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

resource "fmc_fqdn_objects" "new" {
  name        = "terraform_fqdn_object_208"
  value       = "cisco.com"
  description = "just"
  dns_resolution = "IPV4_ONLY"
}

resource "fmc_fqdn_objects" "new_3" {
  name        = "terraform_fqdn_object_209"
  value       = "google.com"
  description = "My First Terraform Network Object"
  dns_resolution = "IPV4_ONLY"
}

output "new_fmc_fqdn_object" {
  value = fmc_fqdn_objects.new.id
}
output "new_fmc_fqdn_object_3" {
  value = fmc_fqdn_objects.new_3.id
}
