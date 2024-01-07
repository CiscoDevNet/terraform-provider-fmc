resource "fmc_access_control_policy" "example" {
  name                              = "POLICY1"
  description                       = "My access control policy"
  default_action                    = "BLOCK"
  default_action_log_begin          = true
  default_action_log_end            = true
  default_action_send_events_to_fmc = true
  default_action_send_syslog        = true
}
