resource "fmc_device_vlan_interface" "example" {
  device_id           = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  logical_name        = "myvlaninterface-7"
  description         = "my description"
  mode                = "NONE"
  security_zone_id    = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  mtu                 = 9000
  vlan_id             = 7
  ipv4_static_address = "10.1.1.1"
  ipv4_static_netmask = "24"
}
