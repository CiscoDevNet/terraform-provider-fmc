resource "fmc_prefilter_policy" "example" {
  name                                = "my_prefilter_policy"
  description                         = "My Prefilter policy"
  default_action                      = "BLOCK_TUNNELS"
  default_action_log_connection_begin = true
  default_action_log_connection_end   = false
  default_action_send_events_to_fmc   = true
  default_action_syslog_alert_id      = "35e197ca-33a8-11ef-b2d1-d98ae17766e7"
  default_action_snmp_alert_id        = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  rules = [
    {
      name           = "rule1"
      rule_type      = "PREFILTER"
      enabled        = true
      action         = "FASTPATH"
      tunnel_zone_id = "0050568A-7F57-0ed3-0000-004294975576"
      time_range_id  = "0050568A-7F57-0ed3-0000-004294975576"
      source_interfaces = [
        {
          id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
          type = "SecurityZone"
        }
      ]
      destination_interfaces = [
        {
          id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
          type = "SecurityZone"
        }
      ]
      source_network_literals = [
        {
          value = "10.1.1.0/24"
        }
      ]
      source_network_objects = [
        {
          id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
          type = "Network"
        }
      ]
      destination_network_literals = [
        {
          value = "10.2.2.0/24"
        }
      ]
      destination_network_objects = [
        {
          id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
          type = "Network"
        }
      ]
      vlan_tag_literals = [
        {
          start_tag = "11"
          end_tag   = "22"
        }
      ]
      vlan_tag_objects = [
        {
          id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
        }
      ]
      source_port_literals = [
        {
          protocol = "6"
          port     = "80"
        }
      ]
      source_port_objects = [
        {
          id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
        }
      ]
      destination_port_literals = [
        {
          protocol = "6"
          port     = "80"
        }
      ]
      destination_port_objects = [
        {
          id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
        }
      ]
      encapsulation_ports = ["GRE"]
      log_begin           = true
      log_end             = true
      send_events_to_fmc  = true
      send_syslog         = true
      syslog_alert_id     = "35e197ca-33a8-11ef-b2d1-d98ae17766e7"
      syslog_severity     = "DEBUG"
      snmp_alert_id       = "76d24097-41c4-4558-a4d0-a8c07ac08470"
    }
  ]
}
