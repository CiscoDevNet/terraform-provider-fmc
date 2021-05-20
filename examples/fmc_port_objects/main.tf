terraform {
  required_providers {
    fmc = {
      source = "hashicorp.com/edu/fmc"
    }
  }
}

resource "fmc_port_objects" "jay_test_port_object_1" {
  name        = "jay_test_port_object_1"
  port        = "3943"
  protocol    = "TCP"
  description = "testing terraform"
  overridable = false
}
