terraform {
  required_providers {
    fmc = {
      source = "CiscoDevNet/fmc"
      # version = "0.2.0"
    }
  }
}



provider "fmc" {
    fmc_username = "admin"
    fmc_password = "P@ssw123"
    fmc_host = "192.168.25.174"
    fmc_insecure_skip_verify = true
}

# Create the security group tag
resource "fmc_sgt_objects" "my_sgt1" {
  name = "another_one"
  description = "bites the dust"
  tag = "26"
}

resource "fmc_sgt_objects" "my_sgt2" {
  name = "and_another_one"
  description = "gone and another one bites the dust"
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