resource "fmc_device_inline_set" "example" {
  device_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  name      = "my_inline_set"
  mtu       = 1500
  fail_safe = true
  interface_pairs = [
    {
      first_interface_id    = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      first_interface_name  = "GigabitEthernet0/1"
      first_interface_type  = "PhysicalInterface"
      second_interface_id   = "22a5b78e-3f6b-4e83-b8f8-c1a0d2b48c31"
      second_interface_name = "GigabitEthernet0/2"
      second_interface_type = "PhysicalInterface"
    }
  ]
  tap_mode               = false
  propagate_link_state   = true
  strict_tcp_enforcement = true
  snort_fail_open_busy   = false
  snort_fail_open_down   = true
}
