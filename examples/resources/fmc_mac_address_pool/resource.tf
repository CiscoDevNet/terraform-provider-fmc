resource "fmc_mac_address_pool" "example" {
  name              = "my_mac_address_pool"
  description       = "My MAC Address Pool object"
  mac_address_range = "aaaa.bbbb.1130-aaaa.bbbb.1140"
  overridable       = true
}
