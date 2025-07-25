resource "fmc_access_rule" "example" {
  access_control_policy_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  category_name            = "category_1"
  action                   = "ALLOW"
  name                     = "rule_1"
  source_network_literals = [
    {
      value = "10.1.1.0/24"
    }
  ]
  destination_network_literals = [
    {
      value = "10.2.2.0/24"
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
  source_network_objects = [
    {
      id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      type = "Network"
    }
  ]
  destination_network_objects = [
    {
      id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      type = "Network"
    }
  ]
  source_dynamic_objects = [
    {
      id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
    }
  ]
  destination_dynamic_objects = [
    {
      id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
    }
  ]
  source_port_literals = [
    {
      type      = "PortLiteral"
      port      = "80"
      protocol  = "6"
      icmp_type = "0"
      icmp_code = "0"
    }
  ]
  source_port_objects = [
    {
      id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
    }
  ]
  destination_port_literals = [
    {
      type      = "PortLiteral"
      port      = "80"
      protocol  = "6"
      icmp_type = "0"
      icmp_code = "0"
    }
  ]
  destination_port_objects = [
    {
      id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
    }
  ]
  source_sgt_objects = [
    {
      name = "my_source_ise_sgt"
      id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      type = "ISESecurityGroupTag"
    }
  ]
  endpoint_device_types = [
    {
      name = "my_endpoint_device_types"
      id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      type = "EndPointDeviceType"
    }
  ]
  destination_sgt_objects = [
    {
      name = "my_destination_ise_sgt"
      id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      type = "ISESecurityGroupTag"
    }
  ]
  source_zones = [
    {
      id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
    }
  ]
  destination_zones = [
    {
      id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
    }
  ]
  url_literals = [
    {
      url = "https://www.example.com/app"
    }
  ]
  url_objects = [
    {
      id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
    }
  ]
  url_categories = [
    {
      id         = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      reputation = "QUESTIONABLE_AND_UNKNOWN"
    }
  ]
  log_begin           = true
  log_end             = true
  send_events_to_fmc  = true
  send_syslog         = true
  syslog_config_id    = "35e197ca-33a8-11ef-b2d1-d98ae17766e7"
  syslog_severity     = "DEBUG"
  snmp_config_id      = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  file_policy_id      = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  intrusion_policy_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  time_range_id       = "0050568A-7F57-0ed3-0000-004294975576"
  variable_set_id     = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  applications = [
    {
      id = "7967"
    }
  ]
  application_filter_objects = [
    {
      id = "bb18bf88-eddc-11ef-83d2-b4300fadd560"
    }
  ]
  application_filters = [
    {
      types = [
        {
          id = "WEBAPP"
        }
      ]
      risks = [
        {
          id = "VERY_LOW"
        }
      ]
      business_relevances = [
        {
          id = "LOW"
        }
      ]
      categories = [
        {
          id = "118"
        }
      ]
      tags = [
        {
          id = "24"
        }
      ]
    }
  ]
}
