terraform {
  required_providers {
    fmc = {
      source = "hashicorp.com/edu/fmc"
    }
  }
}

resource "fmc_url_objects" "new" {
  name        = "terraform_url_object_112"
  url       = "http://wwwin.ami.cisco.com"
  description = "My Second Terraform Network Object"
  type        = "Url"
}

output "new_fmc_url_object" {
  value = fmc_url_objects.new
}
