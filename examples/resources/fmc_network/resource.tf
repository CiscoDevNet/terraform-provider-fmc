resource "fmc_network" "example" {
  name        = "my_network"
  description = "My network object"
  prefix      = "10.1.2.0/24"
  overridable = true
}
