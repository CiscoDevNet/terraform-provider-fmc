resource "fmc_ftd_platform_settings_syslog_email_setup" "example" {
  ftd_platform_settings_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  source_email_address     = "test@example.com"
  destinations = [
    {
      email_addresses = ["recipient@example.com"]
      log_level       = "ERR"
    }
  ]
}
