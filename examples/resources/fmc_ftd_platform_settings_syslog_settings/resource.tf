resource "fmc_ftd_platform_settings_syslog_settings" "example" {
  ftd_platform_settings_id          = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  facility                          = "LOCAL4"
  timestamp_format                  = "RFC_5424"
  device_id_source                  = "USERDEFINEDID"
  device_id_user_defined            = "my_device_id"
  all_syslog_messages               = true
  all_syslog_messages_logging_level = "ERR"
}
