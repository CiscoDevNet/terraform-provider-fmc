resource "fmc_device_ha_pair_monitoring" "example" {
  ha_pair_id           = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  logical_name         = "outside"
  monitor_interface    = true
  ipv4_standby_address = "10.1.1.2"
  ipv6_addresses = [
    {
      active_address  = "2006::1/30"
      standby_address = "2006::2"
    }
  ]
}
