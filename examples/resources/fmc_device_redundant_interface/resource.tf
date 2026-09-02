resource "fmc_device_redundant_interface" "example" {
  device_id                = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  logical_name             = "my_redundant_interface"
  description              = "my description"
  security_zone_id         = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  mtu                      = 9000
  sgt_propagate            = false
  redundant_interface_id   = 1
  primary_interface_id     = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  primary_interface_name   = "GigabitEthernet0/1"
  primary_interface_type   = "PhysicalInterface"
  secondary_interface_id   = "76d24097-41c4-4558-a4d0-a8c07ac08471"
  secondary_interface_name = "GigabitEthernet0/2"
  secondary_interface_type = "PhysicalInterface"
  ipv4_static_address      = "10.1.1.1"
  ipv4_static_netmask      = "24"
}
