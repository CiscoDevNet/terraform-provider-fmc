resource "fmc_access_control_policy_advanced_logging" "example" {
  access_control_policy_id  = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  enabled                   = true
  log_format                = "CSV"
  send_to_other_destination = true
  send_to_splunk            = true
  syslog_alert_id           = "35e197ca-33a8-11ef-b2d1-d98ae17766e7"
}
