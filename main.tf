terraform {
  required_providers {
    fmc = {
      source = "CiscoDevNet/fmc"
      version = "0.2.1"
    }
  }
}

provider "fmc" {
    fmc_username = "admin"
    fmc_password = "Cisco@123"
    fmc_host = "13.235.51.192"
    fmc_insecure_skip_verify = true
}

data "fmc_dynamic_objects" "dy" {
  name = "doj1"
}

# output "test" {
#   value = data.fmc_dynamic_objects.dy
# }

# resource "fmc_dynamic_object" "test" {
#   name        = "DynamicObject"
#   object_type = "IP"
#   description = "testing terraform"
# }

resource "fmc_dynamic_object_mapping" "test" {
  dynamic_object_id =  data.fmc_dynamic_objects.dy.id
  mappings = ["1.1.1.1", "8.8.8.8", "4.4.4.4"]
}