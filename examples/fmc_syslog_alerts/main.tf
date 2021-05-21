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

data "fmc_syslog_alerts" "syslog_alert" {
    name = "Testing Syslog"
}

output "existing_syslog_alert" {
    value = data.fmc_syslog_alerts.syslog_alert
}
