terraform {
  required_providers {
    fmc = {
      source = "hashicorp.com/edu/fmc"
    }
  }
}

resource "fmc_url_objects" "new" {
  name        = "terraform_bulk_url_object_1"
  url       = "http://wwwin.ami.cisco.com"
  description = "My Scale Test Terraform URL Object"
  type        = "Url"
}

resource "fmc_url_objects" "new2" {
  name        = "terraform_bulk_url_object_2"
  url       = "http://wwwin.ami.cisco.com"
  description = "My Scale Test Terraform URL Object"
  type        = "Url"
}

resource "fmc_url_objects" "new3" {
  name        = "terraform_bulk_url_object_3"
  url       = "http://wwwin.ami.cisco.com"
  description = "My Scale Test Terraform URL Object"
  type        = "Url"
}

resource "fmc_url_objects" "new4" {
  name        = "terraform_bulk_url_object_4"
  url       = "http://wwwin.ami.cisco.com"
  description = "My Scale Test Terraform URL Object"
  type        = "Url"
}

resource "fmc_url_objects" "new5" {
  name        = "terraform_bulk_url_object_5"
  url       = "http://wwwin.ami.cisco.com"
  description = "My Scale Test Terraform URL Object"
  type        = "Url"
}

resource "fmc_url_objects" "new6" {
  name        = "terraform_bulk_url_object_6"
  url       = "http://wwwin.ami.cisco.com"
  description = "My Scale Test Terraform URL Object"
  type        = "Url"
}

resource "fmc_url_objects" "new7" {
  name        = "terraform_bulk_url_object_7"
  url       = "http://wwwin.ami.cisco.com"
  description = "My Scale Test Terraform URL Object"
  type        = "Url"
}

resource "fmc_url_objects" "new8" {
  name        = "terraform_bulk_url_object_8"
  url       = "http://wwwin.ami.cisco.com"
  description = "My Scale Test Terraform URL Object"
  type        = "Url"
}

resource "fmc_url_objects" "new9" {
  name        = "terraform_bulk_url_object_9"
  url       = "http://wwwin.ami.cisco.com"
  description = "My Scale Test Terraform URL Object"
  type        = "Url"
}

resource "fmc_url_objects" "new10" {
  name        = "terraform_bulk_url_object_10"
  url       = "http://wwwin.ami.cisco.com"
  description = "My Scale Test Terraform URL Object"
  type        = "Url"
}

output "new_fmc_url_object" {
  value = fmc_url_objects.new
}
