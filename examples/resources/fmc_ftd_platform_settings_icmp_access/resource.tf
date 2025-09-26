resource "fmc_ftd_platform_settings_icmp_access" "example" {
  ftd_platform_settings_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  rate_limit               = 1
  burst_size               = 1
  configurations = [
    {
      action                   = "Permit"
      icmp_service_object_id   = "5a9f2d3c-08d1-11e6-b939-00155d0a1eb1"
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
