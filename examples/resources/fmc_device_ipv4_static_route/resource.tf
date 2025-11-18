resource "fmc_device_ipv4_static_route" "example" {
  device_id              = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  interface_logical_name = "myinterface-0-1"
  interface_id           = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  destination_networks = [
    {
      id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
    }
  ]
  metric               = 254
  gateway_host_literal = "10.0.0.1"
  sla_monitor_id       = "76d24097-41c4-4558-a4d0-a8c07ac08470"
}
