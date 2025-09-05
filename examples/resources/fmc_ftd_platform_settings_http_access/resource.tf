resource "fmc_ftd_platform_settings_http_access" "example" {
  ftd_platform_settings_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  enable_http_server       = true
  port                     = 443
  http_configurations = [
    {
      source_network_object_id = "5a9f6d9c-3f8d-11e4-9163-6c4008b8c5d7"
      interface_objects = [
        {
          id   = "123e4567-e89b-12d3-a456-426614174000"
          type = "SecurityZone"
          name = "Outside"
        }
      ]
    }
  ]
}
