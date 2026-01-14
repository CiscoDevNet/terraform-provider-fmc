resource "fmc_device_vti_interface" "example" {
  device_id                         = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  tunnel_type                       = "STATIC"
  logical_name                      = "my_vti_interface"
  enabled                           = true
  description                       = "My VTI"
  security_zone_id                  = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  priority                          = 100
  tunnel_id                         = 100
  tunnel_source_interface_id        = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  tunnel_source_interface_name      = "GigabitEthernet0/1"
  tunnel_mode                       = "ipv4"
  ipv4_static_address               = "10.10.10.10"
  ipv4_static_netmask               = "24"
  ip_based_monitoring               = true
  ip_based_monitoring_type          = "PEER_IPV4"
  ip_based_monitoring_peer_ip       = "10.10.10.100"
  http_based_application_monitoring = true
}
