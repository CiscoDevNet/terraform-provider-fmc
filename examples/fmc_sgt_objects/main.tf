terraform {
  required_providers {
    fmc = {
      source = "CiscoDevNet/fmc"
      # version = "0.2.0"
    }
  }
}

provider "fmc" {
  fmc_username = var.fmc_username
  fmc_password = var.fmc_password
  fmc_host = var.fmc_host
  fmc_insecure_skip_verify = var.fmc_insecure_skip_verify
}

# Create the security group tag
resource "fmc_sgt_objects" "my_sgt1" {
  name = "sgt_objct-1"
  description = "Applied via TF"
  tag = "26"
}

resource "fmc_sgt_objects" "my_sgt2" {
  name = "sgt_objct-2"
  description = "Applied via TF"
  tag = "27"
}

data "fmc_sgt_objects" "sgt1" {
    name = fmc_sgt_objects.my_sgt1.name
}

data "fmc_sgt_objects" "sgt2" {
    name = fmc_sgt_objects.my_sgt2.name
}

output "security_group_tag_1" {
    value = data.fmc_sgt_objects.sgt1
}

output "security_group_tag_2" {
    value = data.fmc_sgt_objects.sgt2
}