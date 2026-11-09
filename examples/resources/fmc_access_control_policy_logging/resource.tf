resource "fmc_access_control_policy_logging" "example" {
  access_control_policy_id                = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  syslog_alert_id                         = "35e197ca-33a8-11ef-b2d1-d98ae17766e7"
  syslog_use_platform_settings            = true
  syslog_platform_settings_severity       = "CRIT"
  ips_events_send_syslog                  = true
  ips_events_syslog_severity              = "ERR"
  ips_events_syslog_alert_id              = "35e197ca-33a8-11ef-b2d1-d98ae17766e7"
  file_and_malware_events_send_syslog     = true
  file_and_malware_events_syslog_severity = "ALERT"
  file_and_malware_events_syslog_alert_id = "35e197ca-33a8-11ef-b2d1-d98ae17766e7"
}
