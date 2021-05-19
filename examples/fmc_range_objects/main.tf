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
