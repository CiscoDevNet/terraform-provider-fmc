resource "fmc_device_physical_interface" "example" {
  device_id           = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  logical_name        = "myinterface-0-1"
  description         = "my description"
  mode                = "NONE"
  security_zone_id    = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  name                = "GigabitEthernet0/1"
  mtu                 = 1400
  sgt_propagate       = false
  ipv4_static_address = "10.1.1.1"
  ipv4_static_netmask = "24"
}
