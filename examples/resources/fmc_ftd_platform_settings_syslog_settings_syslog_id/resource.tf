resource "fmc_ftd_platform_settings_syslog_settings_syslog_id" "example" {
  ftd_platform_settings_id                 = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  ftd_platform_settings_syslog_settings_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  syslog_id                                = 106004
  logging_level                            = "ERR"
  enabled                                  = true
}
