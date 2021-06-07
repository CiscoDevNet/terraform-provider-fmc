terraform {
  required_providers {
    fmc = {
      source = "hashicorp.com/cisco/fmc"
    }
  }
}

provider "fmc" {
  fmc_username = var.fmc_username
  fmc_password = var.fmc_password
  fmc_host = var.fmc_host
  fmc_insecure_skip_verify = var.fmc_insecure_skip_verify
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

data "fmc_port_objects" "http" {
    name = "HTTPS"
}

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

resource "fmc_access_rules" "access_rule_1" {
    acp = fmc_access_policies.access_policy.id
    section = "mandatory"
    name = "Test rule 1"
    action = "allow"
    enabled = true
    enable_syslog = true
    syslog_severity = "alert"
    send_events_to_fmc = true
    log_files = true
    log_end = true
    source_zones {
        source_zone {
            id = data.fmc_security_zones.inside.id
            type =  data.fmc_security_zones.inside.type
        }
        source_zone {
            id = data.fmc_security_zones.outside.id
            type =  data.fmc_security_zones.outside.type
        }
    }
    destination_zones {
        destination_zone {
            id = data.fmc_security_zones.outside.id
            type =  data.fmc_security_zones.outside.type
        }
    }
    source_networks {
        source_network {
            id = data.fmc_network_objects.source.id
            type =  data.fmc_network_objects.source.type
        }
    }
    destination_networks {
        destination_network {
            id = data.fmc_network_objects.dest.id
            type =  data.fmc_network_objects.dest.type
        }
    }
    destination_ports {
        destination_port {
            id = data.fmc_port_objects.http.id
            type =  data.fmc_port_objects.http.type
        }
    }
    urls {
        url {
            id = fmc_url_objects.dest_url.id
            type = "Url"
        }
    }
    ips_policy = data.fmc_ips_policies.ips_policy.id
    syslog_config = data.fmc_syslog_alerts.syslog_alert.id
    new_comments = [ "New", "comment" ]
}

resource "fmc_access_rules" "access_rule_2" {
    acp = fmc_access_policies.access_policy.id
    section = "mandatory"
    insert_before = 2 # Wont work as assumed since terraform does not 
    name = "Test rule 2"
    action = "allow"
    enabled = true
    enable_syslog = true
    syslog_severity = "alert"
    send_events_to_fmc = true
    log_files = true
    log_end = true
    source_zones {
        source_zone {
            id = data.fmc_security_zones.inside.id
            type =  data.fmc_security_zones.inside.type
        }
        # source_zone {
        #     id = data.fmc_security_zones.outside.id
        #     type =  data.fmc_security_zones.outside.type
        # }
    }
    destination_zones {
        destination_zone {
            id = data.fmc_security_zones.outside.id
            type =  data.fmc_security_zones.outside.type
        }
    }
    source_networks {
        source_network {
            id = data.fmc_network_objects.source.id
            type =  data.fmc_network_objects.source.type
        }
    }
    destination_networks {
        destination_network {
            id = data.fmc_network_objects.dest.id
            type =  data.fmc_network_objects.dest.type
        }
    }
    destination_ports {
        destination_port {
            id = data.fmc_port_objects.http.id
            type =  data.fmc_port_objects.http.type
        }
    }
    urls {
        url {
            id = fmc_url_objects.dest_url.id
            type = "Url"
        }
    }
    ips_policy = data.fmc_ips_policies.ips_policy.id
    syslog_config = data.fmc_syslog_alerts.syslog_alert.id
    new_comments = [ "New comment" ]
    depends_on = [
        fmc_access_rules.access_rule_1
    ]
}

resource "fmc_access_rules" "access_rule_3" {
    acp = fmc_access_policies.access_policy.id
    section = "mandatory"
    # insert_before = 1 # Wont work as assumed since terraform does not 
    name = "Test rule 3"
    action = "allow"
    enabled = true
    enable_syslog = true
    syslog_severity = "alert"
    send_events_to_fmc = true
    log_files = true
    log_end = true
    source_zones {
        source_zone {
            id = data.fmc_security_zones.inside.id
            type =  data.fmc_security_zones.inside.type
        }
        # source_zone {
        #     id = data.fmc_security_zones.outside.id
        #     type =  data.fmc_security_zones.outside.type
        # }
    }
    destination_zones {
        destination_zone {
            id = data.fmc_security_zones.outside.id
            type =  data.fmc_security_zones.outside.type
        }
    }
    source_networks {
        source_network {
            id = data.fmc_network_objects.source.id
            type =  data.fmc_network_objects.source.type
        }
    }
    destination_networks {
        destination_network {
            id = data.fmc_network_objects.dest.id
            type =  data.fmc_network_objects.dest.type
        }
    }
    destination_ports {
        destination_port {
            id = data.fmc_port_objects.http.id
            type =  data.fmc_port_objects.http.type
        }
    }
    urls {
        url {
            id = fmc_url_objects.dest_url.id
            type = "Url"
        }
    }
    ips_policy = data.fmc_ips_policies.ips_policy.id
    syslog_config = data.fmc_syslog_alerts.syslog_alert.id
    new_comments = [ "New comment" ]
    depends_on = [
        fmc_access_rules.access_rule_2
    ]
}

output "new_fmc_access_policy" {
    value = fmc_access_policies.access_policy
}

output "new_fmc_access_rule_1" {
    value = fmc_access_rules.access_rule_1
}

output "new_fmc_access_rule_3" {
    value = fmc_access_rules.access_rule_3
}
