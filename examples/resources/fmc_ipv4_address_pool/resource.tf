resource "fmc_ipv4_address_pool" "example" {
  name        = "my_ipv4_address_pool"
  description = "My IPv4 Address Pool object"
  range       = "10.0.0.10-10.0.0.20"
  netmask     = "255.255.255.0"
  overridable = true
}
