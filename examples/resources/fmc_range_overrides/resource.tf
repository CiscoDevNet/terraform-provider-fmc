resource "fmc_range_overrides" "example" {
  parent_name = "my_range"
  parent_id   = "12345678-90ab-cdef-1234-567890abcdef"
  overrides = [
    {
      target_id   = "12345678-90ab-cdef-1234-567890abcdef"
      target_type = "Device"
      description = "My Range object"
      ip_range    = "10.1.0.1-10.1.0.9"
    }
  ]
}
