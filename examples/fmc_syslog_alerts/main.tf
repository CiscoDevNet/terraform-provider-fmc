terraform {
  required_providers {
    fmc = {
      source = "hashicorp.com/edu/fmc"
    }
  }
}

provider "fmc" {
  fmc_username = var.fmc_username
  fmc_password = var.fmc_password
  fmc_host = var.fmc_host
  fmc_insecure_skip_verify = var.fmc_insecure_skip_verify
}

data "fmc_syslog_alerts" "syslog_alert" {
    name = "Testing Syslog"
}

output "existing_syslog_alert" {
    value = data.fmc_syslog_alerts.syslog_alert
}
