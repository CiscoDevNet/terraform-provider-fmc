resource "fmc_device_etherchannel_interface" "example" {
  device_id        = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  logical_name     = "myinterface-0-1"
  description      = "my description"
  mode             = "NONE"
  security_zone_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  mtu              = 9000
  sgt_propagate    = false
  ether_channel_id = "1"
  selected_interfaces = [
    {
      id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      name = "GigabitEthernet0/1"
    }
  ]
  ipv4_static_address = "10.1.1.1"
  ipv4_static_netmask = "24"
}
