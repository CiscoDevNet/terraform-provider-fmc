terraform {
  required_providers {
    fmc = {
      source = "hashicorp.com/edu/fmc"
    }
  }
}

provider "fmc" {
  fmc_username = var.fmc_username
  fmc_password = var.fmc_password
  fmc_host = var.fmc_host
  fmc_insecure_skip_verify = var.fmc_insecure_skip_verify
}


resource "fmc_range_objects" "shbharti-range-1" {
  name        = "shbharti-range-1"
  value       = "1.2.3.4-1.2.3.10"
  description = "testing terraform"
}

resource "fmc_range_objects" "shbharti-range-2" {
  name        = "shbharti-range-2"
  value       = "1.2.3.4-1.2.3.9"
  description = "testing terraform"
}

output "fmc_range_object-1" {
  value = fmc_range_objects.shbharti-range-1.name
}

output "new_fmc_range_object_3" {
  value = fmc_range_objects.shbharti-range-2.name
}
