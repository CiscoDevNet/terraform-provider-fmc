terraform {
  required_providers {
    fmc = {
      source = "CiscoDevNet/fmc"
      # version = "0.2.0"
    }
  }
}

provider "fmc" {
  fmc_username = var.fmc_username
  fmc_password = var.fmc_password
  fmc_host = var.fmc_host
  fmc_insecure_skip_verify = var.fmc_insecure_skip_verify
}

data "fmc_ips_policies" "ips_policy" {
    name = "Connectivity Over Security"
}

data "fmc_syslog_alerts" "syslog_alert" {
    name = "Testing Syslog"
}

data "fmc_access_policies" "access_policy" {
    name = "FTD"
}

resource "fmc_access_policies" "access_policy" {
    name = "Terraform Access Policy"
    default_action = "block" # Cannot have block with base IPS policy
    # default_action = "permit"
    # default_action_base_intrusion_policy_id = data.fmc_ips_policies.ips_policy.id
    default_action_send_events_to_fmc = true
    default_action_log_begin = true
    default_action_syslog_config_id = data.fmc_syslog_alerts.syslog_alert.id
}

output "existing_fmc_access_policy" {
    value = data.fmc_access_policies.access_policy
}

output "new_fmc_access_policy" {
    value = fmc_access_policies.access_policy
}
