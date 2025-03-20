resource "fmc_ipv6_address_pools" "example" {
  items = {
    my_ipv6_address_pool = {
      description         = "My IPv4 Address Pool object"
      start_address       = "2001:db8::1/64"
      number_of_addresses = 10
      overridable         = true
    }
  }
}
