resource "fmc_device_vtep_policy" "example" {
  device_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  vteps = [
    {
      source_interface_id      = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      nve_number               = 1
      neighbor_discovery       = "STATIC_PEER_IP"
      neighbor_address_literal = "192.168.0.1"
    }
  ]
}
