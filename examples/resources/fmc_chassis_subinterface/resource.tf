resource "fmc_chassis_subinterface" "example" {
  chassis_id            = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  parent_interface_name = "Ethernet1/1"
  parent_interface_id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  sub_interface_id      = 7
  vlan_id               = 4094
  port_type             = "DATA"
}
