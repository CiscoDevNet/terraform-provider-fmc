terraform {
  required_providers {
    fmc = {
      source = "hashicorp.com/edu/fmc"
    }
  }
}

resource "fmc_url_objects" "new" {
  name        = "terraform_url_object_100"
  url       = "http://wwwin.ami.cisco.com"
  description = "just"
  type        = "Url"
}

resource "fmc_url_objects" "new_3" {
  name        = "terraform_url_object_111"
  url       = "http://wwwin.ami.cisco.com"
  description = "My First Terraform Network Object"
  type        = "Url"
}

output "new_fmc_url_object" {
  value = fmc_url_objects.new
}
output "new_fmc_url_object_3" {
  value = fmc_url_objects.new_3
}
