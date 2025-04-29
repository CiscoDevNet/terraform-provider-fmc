resource "fmc_device_loopback_interface" "example" {
  device_id           = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  logical_name        = "my_loopback_1"
  enabled             = true
  loopback_id         = 1
  description         = "my VTI interface"
  ipv4_static_address = "10.1.1.1"
  ipv4_static_netmask = "24"
  ipv6_addresses = [
    {
      address = "2004::10"
      prefix  = "64"
    }
  ]
}
