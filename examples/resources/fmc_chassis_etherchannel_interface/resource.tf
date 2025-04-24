resource "fmc_chassis_etherchannel_interface" "example" {
  chassis_id       = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  ether_channel_id = 10
  port_type        = "DATA"
  admin_state      = "ENABLED"
  selected_interfaces = [
    {
      id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      name = "Ethernet1/1"
    }
  ]
  speed     = "AUTO"
  lacp_mode = "ACTIVE"
  lacp_rate = "DEFAULT"
}
