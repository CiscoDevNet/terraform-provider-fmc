terraform {
  required_providers {
    fmc = {
      source = "CiscoDevNet/fmc"
      # version = "0.1.1"
    }
  }
}

provider "fmc" {
  fmc_username = var.fmc_username
  fmc_password = var.fmc_password
  fmc_host = var.fmc_host
  fmc_insecure_skip_verify = var.fmc_insecure_skip_verify
}

# data "fmc_standard_acl" "test" {
#     name = "test-updated"
# }

data "fmc_network_objects" "any" {
    name = "any-ipv4"
}
data "fmc_host_objects" "test" {
    name = "test-object"
}
resource "fmc_standard_acl" "create1"{
    count = 3
    name = "kadahic${count.index+1}"
    action = "DENY"
    object_id = data.fmc_network_objects.any.id
    literal_type = "Host"
    literal_value = "1.1.1.1"
}
# output "existing_acl" {
#     value = data.fmc_standard_acl.test
# }
