resource "fmc_device_vrf" "example" {
  device_id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  name        = "VRF_A"
  description = "My VRF instance"
  interfaces = [
    {
      interface_id           = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      interface_name         = "GigabitEthernet0/0"
      interface_logical_name = "outside"
    }
  ]
}
