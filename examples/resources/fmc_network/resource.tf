resource "fmc_network" "example" {
  name        = "fmc_network_1"
  description = "My network object"
  prefix      = "10.1.2.0/24"
  overridable = true
}
