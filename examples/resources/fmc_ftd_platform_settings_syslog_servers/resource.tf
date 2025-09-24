resource "fmc_ftd_platform_settings_syslog_servers" "example" {
  ftd_platform_settings_id                          = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  allow_user_traffic_when_tcp_syslog_server_is_down = true
  message_queue_size                                = 512
  syslog_servers = [
    {
      network_object_id        = "c1a0f5b6-3d4e-11b2-9f8f-0242ac112345"
      protocol                 = "TCP"
      port                     = 1470
      secure_syslog            = true
      use_management_interface = false
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
