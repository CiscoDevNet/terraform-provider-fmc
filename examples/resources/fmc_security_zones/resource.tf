resource "fmc_security_zones" "example" {
  items = {
    security_zone_1 = {
      interface_type = "ROUTED"
    }
  }
}
