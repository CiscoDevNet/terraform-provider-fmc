terraform {
  required_providers {
    fmc = {
      source = "hashicorp.com/edu/fmc"
    }
  }
}

provider "fmc" {
  fmc_username = "api"
  fmc_password = "CXsecurity!@34"
  fmc_host = "10.106.107.228"
  fmc_insecure_skip_verify = true
}

data "fmc_security_zones" "inside" {
    name = "inside"
}

data "fmc_security_zones" "outside" {
    name = "outside"
}

data "fmc_network_objects" "source" {
    name = "VLAN825-Public"
}

data "fmc_network_objects" "dest" {
    name = "VLAN825-Private"
}

# data "fmc_port_objects" "http" {
#     name = "HTTP"
# }

data "fmc_ips_policies" "ips_policy" {
    name = "Connectivity Over Security"
}

data "fmc_syslog_alerts" "syslog_alert" {
    name = "Testing Syslog"
}

resource "fmc_url_objects" "dest_url" {
    name = "Guacamole"
    url = "http://guacamole.adyah.cisco"
    description = "Testing ACR"
}

resource "fmc_access_policies" "access_policy" {
    name = "Terraform Access Policy"
    # default_action = "block" # Cannot have block with base IPS policy
    default_action = "permit"
    default_action_base_intrusion_policy_id = data.fmc_ips_policies.ips_policy.id
    default_action_send_events_to_fmc = "true"
    default_action_log_end = "true"
    default_action_syslog_config_id = data.fmc_syslog_alerts.syslog_alert.id
}

resource "fmc_access_rules" "access_rule" {
    acp = fmc_access_policies.access_policy.id
    name = "Test rule"
    action = "allow"
    enabled = true
    enable_syslog = true
    syslog_severity = "alert"
    send_events_to_fmc = true
    log_files = true
    log_end = true
    source_zones = [ data.fmc_security_zones.inside.id ]
    destination_zones = [ data.fmc_security_zones.outside.id ]
    source_networks = [ data.fmc_network_objects.source.id ]
    destination_networks = [ data.fmc_network_objects.dest.id ]
    # destination_ports = [ data.fmc_port_objects.http.id ]
    urls = [ fmc_url_objects.dest_url.id ]
    ips_policy = data.fmc_ips_policies.ips_policy.id
    syslog_config = data.fmc_syslog_alerts.syslog_alert.id
    comments = [ "Hello", "world" ]
}

output "new_fmc_access_policy" {
    value = fmc_access_policies.access_policy
}

output "new_fmc_access_rule" {
    value = fmc_access_rules.access_rule
}
