resource "fmc_mac_address_pools" "example" {
  items = {
    my_mac_address_pools = {
      description       = "My MAC Address Pool object"
      mac_address_range = "aaaa.bbbb.1130-aaaa.bbbb.1140"
      overridable       = true
    }
  }
}
