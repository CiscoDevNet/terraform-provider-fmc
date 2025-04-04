resource "fmc_device_ha_pair_physical_interface_mac_address" "example" {
  ha_pair_id          = "76d24097-41c4-4558-a4d0-a8c07ac13928"
  interface_name      = "GigabitEthernet0/0"
  interface_id        = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  active_mac_address  = "c460.15e4.0edd"
  standby_mac_address = "c460.15e4.0ed0"
}
