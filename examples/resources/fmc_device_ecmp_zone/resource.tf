resource "fmc_device_ecmp_zone" "example" {
  device_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  name      = "my_ecmp_zone"
  interfaces = [
    {
      id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      name = "GigabitEthernet0/1"
    }
  ]
}
