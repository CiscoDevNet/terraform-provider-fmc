resource "fmc_ftd_platform_settings_syslog_logging_setup" "example" {
  ftd_platform_settings_id                    = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  enable_logging                              = true
  enable_logging_on_the_failover_standby_unit = false
  emblem_format                               = false
  send_debug_messages_as_syslog               = false
  internal_buffer_memory_size                 = 4096
  fmc_logging_type                            = "VPN"
  fmc_logging_level                           = "ERR"
  ftp_server_host_id                          = "d3f5e8c0-1d4b-11b2-9f0e-ecf4bbf7a5c6"
  ftp_server_username                         = "ftpuser"
  ftp_server_password                         = "ftppassword"
  ftp_server_path                             = "/logs"
  ftp_server_interface_groups = [
    {
      id = "e7f5e8c0-1d4b-11b2-9f0e-ecf4bbf7a5c6"
    }
  ]
  flash                    = false
  flash_maximum_space      = 3076
  flash_minimum_free_space = 1024
}
