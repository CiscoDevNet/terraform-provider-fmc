resource "fmc_device_eigrp" "example" {
  device_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  as_number = "100"
  networks = [
    {
      id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
    }
  ]
  auto_summary = false
  passive_interfaces = [
    {
      id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
    }
  ]
  neighbors = [
    {
      interface_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      host_id      = "76d24097-41c4-4558-a4d0-a8c07ac08470"
    }
  ]
  filter_rules = [
    {
      filter_direction        = "INBOUND"
      standard_access_list_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      interface_id            = "76d24097-41c4-4558-a4d0-a8c07ac08470"
    }
  ]
  redistributions = [
    {
      redistribute_protocol = "STATIC"
      route_map_id          = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      metric_bandwidth      = 1000
      metric_delay          = 100
      metric_reliability    = 255
      metric_loading        = 1
      metric_mtu            = 1500
    }
  ]
  summary_addresses = [
    {
      interface_id            = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      network_id              = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      administrative_distance = 5
    }
  ]
  interfaces = [
    {
      interface_id            = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      interface_type          = "PhysicalInterface"
      hello_interval          = 10
      hold_time               = 30
      split_horizon           = false
      delay_time              = 20
      md5_authentication      = true
      authentication_key_type = "UNENCRYPTED"
      authentication_key_id   = 2
      authentication_key      = "my_secret"
    }
  ]
  router_id                                       = "10.10.10.1"
  default_information_in_standard_access_list_id  = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  default_information_out_standard_access_list_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  administrative_distance_internal                = 90
  administrative_distance_external                = 170
  log_neighbor_changes                            = true
  log_neighbor_warnings                           = true
  log_neighbor_warnings_interval                  = 15
  stub_receive_only                               = false
  stub_advertise_connected                        = true
  stub_advertise_redistributed                    = false
  stub_advertise_static                           = false
  stub_advertise_summary                          = true
  default_metric_bandwidth                        = 1000
  default_metric_delay_time                       = 100
  default_metric_reliability                      = 255
  default_metric_loading                          = 1
  default_metric_mtu                              = 1500
}
