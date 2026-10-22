resource "fmc_ftd_platform_settings_external_authentication" "example" {
  ftd_platform_settings_id          = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  external_authentication_server_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  primary_server_interfaces = [
    {
      id   = "12345678-1234-1234-1234-123456789abc"
      type = "PhysicalInterface"
      name = "outside"
    }
  ]
  backup_server_interfaces = [
    {
      id   = "12345678-1234-1234-1234-426614174000"
      type = "PhysicalInterface"
      name = "inside"
    }
  ]
}
