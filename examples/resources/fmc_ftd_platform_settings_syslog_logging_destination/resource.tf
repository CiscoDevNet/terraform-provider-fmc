resource "fmc_ftd_platform_settings_syslog_logging_destination" "example" {
  ftd_platform_settings_id    = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  logging_destination         = "INTERNAL_BUFFER"
  event_class_filter_criteria = "SEVERITY"
  event_class_filter_value    = "EMERG"
  event_class_filters = [
    {
      class    = "AUTH"
      severity = "ERR"
    }
  ]
}
