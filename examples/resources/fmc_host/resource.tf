resource "fmc_host" "example" {
  name        = "my_host"
  description = "My host object"
  ip          = "10.1.1.1"
  overridable = true
}
