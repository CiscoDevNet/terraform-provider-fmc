resource "fmc_network_group_overrides" "example" {
  parent_name = "my_network_group"
  parent_id   = "12345678-90ab-cdef-1234-567890abcdef"
  overrides = [
    {
      target_id   = "12345678-90ab-cdef-1234-567890abcdef"
      target_type = "Device"
      description = "My Network Group Overridden name"
      objects = [
        {
          id   = "12345678-1234-1234-1234-123456789012"
          name = "fmc_network_object_1"
        }
      ]
      literals = [
        {
          value = "10.1.1.0/24"
        }
      ]
    }
  ]
}
