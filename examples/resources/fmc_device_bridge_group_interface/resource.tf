resource "fmc_device_bridge_group_interface" "example" {
  device_id       = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  logical_name    = "my_bridge_group_interface"
  description     = "My Bridge Group Interface"
  bridge_group_id = 100
  selected_interfaces = [
    {
      id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      name = "GigabitEthernet0/0"
    }
  ]
  ipv4_static_address = "10.1.1.1"
  ipv4_static_netmask = "24"
  ipv6_addresses = [
    {
      address = "2004::1"
      prefix  = "64"
    }
  ]
  ipv6_dad            = true
  ipv6_dad_attempts   = 1
  ipv6_ns_interval    = 1000
  ipv6_reachable_time = 0
  arp_table_entries = [
    {
      mac_address  = "0123.4567.89ab"
      ip_address   = "10.1.1.10"
      enable_alias = true
    }
  ]
}
