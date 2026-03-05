resource "fmc_host_overrides" "example" {
  parent_name = "my_host"
  parent_id   = "12345678-90ab-cdef-1234-567890abcdef"
  overrides = [
    {
      target_id   = "12345678-90ab-cdef-1234-567890abcdef"
      target_type = "Device"
      description = "My Host object"
      ip          = "10.1.1.1"
    }
  ]
}
