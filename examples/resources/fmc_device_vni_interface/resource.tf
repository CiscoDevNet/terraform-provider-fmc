resource "fmc_device_vni_interface" "example" {
  device_id               = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  vni_id                  = 42
  multicast_group_address = "224.0.0.24"
  segment_id              = 501
  nve_number              = 1
  logical_name            = "vni42"
  description             = "my description"
  security_zone_id        = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  ipv4_static_address     = "10.2.2.2"
  ipv4_static_netmask     = "24"
  ipv6_enable             = true
  ipv6_enable_auto_config = true
}
