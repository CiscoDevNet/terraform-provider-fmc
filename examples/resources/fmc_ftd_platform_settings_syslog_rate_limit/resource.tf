resource "fmc_ftd_platform_settings_syslog_rate_limit" "example" {
  ftd_platform_settings_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  rate_limit_type          = "LOG_LEVEL"
  rate_limit_value         = "ERR"
  number_of_messages       = 100
  interval                 = 60
}
