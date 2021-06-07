terraform {
  required_providers {
    fmc = {
      source = "hashicorp.com/cisco/fmc"
    }
  }
}

provider "fmc" {
  fmc_username = var.fmc_username
  fmc_password = var.fmc_password
  fmc_host = var.fmc_host
  fmc_insecure_skip_verify = var.fmc_insecure_skip_verify
}

data "fmc_port_objects" "existing" {
  name = "DNS_over_TCP"
  port = "53"
}

# resource "fmc_port_objects" "jay_test_port_object_1" {
#   name        = "jay_test_port_object_1"
#   port        = "3943"
#   protocol    = "TCP"
#   description = "testing terraform"
#   overridable = false
# }

resource "fmc_port_objects" "similar" {
  name = "${data.fmc_port_objects.existing.name}-Test"
  port = data.fmc_port_objects.existing.port
  protocol = data.fmc_port_objects.existing.protocol
}

output "existing_port" {
  value = data.fmc_port_objects.existing
}

output "similar_port" {
  value = fmc_port_objects.similar
}
