resource "fmc_host" "example" {
  name        = "HOST1"
  description = "My host object"
  ip          = "10.1.1.1"
  overridable = true
}
