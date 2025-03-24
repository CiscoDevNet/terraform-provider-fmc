resource "fmc_ipv4_address_pools" "example" {
  items = {
    my_ipv4_address_pools = {
      description = "My IPv4 Address Pool object"
      range       = "10.0.0.10-10.0.0.20"
      mask        = "255.255.255.0"
      overridable = true
    }
  }
}
