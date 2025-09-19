resource "fmc_ftd_platform_settings_syslog_event_list" "example" {
  ftd_platform_settings_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  name                     = "my_event_list"
  event_classes = [
    {
      class    = "AUTH"
      severity = "ERR"
    }
  ]
  message_ids = ["100000-200000"]
}
