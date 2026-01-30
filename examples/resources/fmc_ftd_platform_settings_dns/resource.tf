resource "fmc_ftd_platform_settings_dns" "example" {
  ftd_platform_settings_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  dns_server_groups = [
    {
      id         = "12345678-1234-1234-1234-123456789012"
      is_default = true
    }
  ]
  expire_entry_timer = 1
  poll_timer         = 240
  interface_objects = [
    {
      id   = "12345678-1234-1234-1234-123456789012"
      type = "SecurityZone"
    }
  ]
  use_management_interface = true
}
