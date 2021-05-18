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

resource "fmc_url_objects" "new" {
  name        = "terraform_bulk_url_object_21"
  url       = "http://wwwin.ami.cisco.com"
  description = "My Scale Test Terraform URL Object"
}

output "new_fmc_url_object" {
  value = fmc_url_objects.new
}
