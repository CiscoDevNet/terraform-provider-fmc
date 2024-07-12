resource "fmc_device_physical_interface" "example" {
  device_id                   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  mode                        = "NONE"
  name                        = "GigabitEthernet0/1"
  logical_name                = "myinterface-0-1"
  description                 = "my description"
  mtu                         = 9000
  security_zone_id            = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  ipv4_static_address         = "10.1.1.1"
  ipv4_static_netmask         = "24"
  ipv6_enable                 = true
  ipv6_enforce_eui            = true
  ipv6_enable_auto_config     = true
  ipv6_enable_dhcp_address    = true
  ipv6_enable_dhcp_nonaddress = true
  ipv6_enable_ra              = false
  ipv6_addresses = [
    {
      address = "2004::"
      prefix  = "124"
    }
  ]
}
