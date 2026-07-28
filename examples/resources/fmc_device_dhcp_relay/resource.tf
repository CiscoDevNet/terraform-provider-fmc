resource "fmc_device_dhcp_relay" "example" {
  device_id             = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  ipv4_relay_timeout    = 60
  ipv6_relay_timeout    = 60
  trust_all_information = true
  relay_agents = [
    {
      interface_id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      interface_name = "GigabitEthernet0/1"
      interface_type = "PhysicalInterface"
      ipv4_relay     = true
      ipv6_relay     = false
      set_route      = true
    }
  ]
  servers = [
    {
      server_id             = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      server_interface_id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      server_interface_name = "GigabitEthernet0/1"
      server_interface_type = "PhysicalInterface"
      client_interfaces = [
        {
          id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
          name = "GigabitEthernet0/1"
          type = "PhysicalInterface"
        }
      ]
    }
  ]
}
