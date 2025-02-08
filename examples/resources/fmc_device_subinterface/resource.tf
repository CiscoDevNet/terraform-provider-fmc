resource "fmc_device_subinterface" "example" {
  device_id            = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  logical_name         = "myinterface-0-1"
  description          = "my description"
  security_zone_id     = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  mtu                  = 9000
  enable_sgt_propagate = false
  interface_name       = "GigabitEthernet0/1"
  sub_interface_id     = 7
  vlan_id              = 4094
  ipv4_static_address  = "10.1.1.1"
  ipv4_static_netmask  = "24"
}
