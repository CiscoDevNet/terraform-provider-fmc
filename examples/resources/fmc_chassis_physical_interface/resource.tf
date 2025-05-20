resource "fmc_chassis_physical_interface" "example" {
  chassis_id  = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  name        = "Ethernet1/1"
  port_type   = "DATA"
  admin_state = "ENABLED"
  speed       = "AUTO"
}
