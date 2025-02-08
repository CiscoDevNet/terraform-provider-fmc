resource "fmc_device_bgp_general_settings" "example" {
  device_id                   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  as_number                   = "65535"
  router_id                   = "AUTOMATIC"
  scanning_interval           = 60
  as_number_in_path_attribute = 50
  tcp_path_mtu_discovery      = true
  reset_session_upon_failover = true
  enforce_first_peer_as       = true
  use_dot_notation            = false
}
