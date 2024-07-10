resource "fmc_device_subinterface" "example" {
  device_id                   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  interface_id                = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  index                       = 7
  vlan_id                     = 4094
  logical_name                = "mysubinterface-0-1.7"
  description                 = "my description"
  security_zone_id            = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  ipv4_static_address         = "10.2.2.2"
  ipv4_static_netmask         = "24"
  ipv6_enable                 = true
  ipv6_enforce_eui            = true
  ipv6_enable_auto_config     = true
  ipv6_enable_dhcp_address    = true
  ipv6_enable_dhcp_nonaddress = true
  ipv6_enable_ra              = false
  ipv6_addresses = [
    {
      address = "2005::"
      prefix  = "124"
    }
  ]
}
