resource "fmc_device_ospf_interface" "example" {
  device_id            = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  interface_id         = "123e4567-e89b-12d3-a456-426614174000"
  default_cost         = 10
  priority             = 1
  mtu_missmatch_ignore = false
  hello_interval       = 10
  transmit_delay       = 1
  retransmit_interval  = 5
  dead_interval        = 40
  hello_multiplier     = 4
  point_to_point       = false
  bfd                  = false
}
