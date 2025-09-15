resource "fmc_ftd_platform_settings_syslog_settings" "example" {
  ftd_platform_settings_id          = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  facility                          = "LOCAL4"
  timestamp_format                  = "RFC_5424"
  device_id_type                    = "USERDEFINEDID"
  device_id_user_defined_id         = "my_device_id"
  device_id_interface_id            = "123e4567-e89b-12d3-a456-426614174000"
  all_syslog_messages               = true
  all_syslog_messages_logging_level = "ERR"
}
