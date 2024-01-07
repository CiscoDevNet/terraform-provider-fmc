resource "fmc_network" "example" {
  name        = "NET1"
  description = "My network object"
  prefix      = "10.1.2.0/24"
  overridable = true
}
