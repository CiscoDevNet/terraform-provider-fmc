resource "fmc_ftd_platform_settings_snmp" "example" {
  ftd_platform_settings_id  = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  enable_snmp_servers       = true
  read_community_string     = "public"
  system_administrator_name = "admin"
  location                  = "Data Center 1"
  listen_port               = 161
  snmp_management_hosts = [
    {
      management_host_ip_object_id    = "123e4567-e89b-12d3-a456-426614174000"
      snmp_version                    = "SNMPv3"
      username                        = "snmpuser1"
      poll                            = true
      trap                            = true
      trap_port                       = 162
      use_device_management_interface = false
      interface_objects = [
        {
          id   = "123e4567-e89b-12d3-a456-426614174000"
          type = "SecurityZone"
          name = "Outside"
        }
      ]
    }
  ]
  snmpv3_users = [
    {
      security_level                = "Priv"
      username                      = "snmpuser1"
      encryption_password_type      = "Clear"
      authentication_algorithm_type = "SHA256"
      authentication_password       = "MyAuthPassword123"
      encryption_type               = "AES256"
      encryption_password           = "MyEncryptionPassword123"
    }
  ]
  trap_syslog                        = true
  trap_authentication                = true
  trap_link_up                       = true
  trap_link_down                     = true
  trap_cold_start                    = true
  trap_warm_start                    = true
  trap_field_replacement_unit_insert = true
  trap_field_replacement_unit_delete = true
  trap_configuration_change          = true
  trap_connection_limit_reached      = true
  trap_nat_packet_discard            = true
  trap_cpu_rising_threshold          = true
  trap_cpu_rising_threshold_value    = 70
  trap_cpu_rising_threshold_interval = 1
  trap_memory_rising_threshold       = true
  trap_memory_rising_threshold_value = 70
  trap_failover                      = true
  trap_cluster                       = true
  trap_peer_flap                     = true
}
