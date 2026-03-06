resource "fmc_network_overrides" "example" {
  parent_name = "my_network"
  parent_id   = "12345678-90ab-cdef-1234-567890abcdef"
  overrides = [
    {
      target_id   = "12345678-90ab-cdef-1234-567890abcdef"
      target_type = "Device"
      description = "My Network object"
      prefix      = "10.1.2.0/24"
    }
  ]
}
