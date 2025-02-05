resource "fmc_security_zones" "example" {
  items = {
    my_security_zones = {
      interface_type = "ROUTED"
    }
  }
}
