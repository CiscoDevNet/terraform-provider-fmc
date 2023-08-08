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

data "fmc_access_policies" "access_policy"{
  name="test-acp"
}

resource "fmc_devices_bulk" "device"{
  devices {
    name = "FTD1"
    hostname = "<FTD IP ADDRESS>"
    regkey = "<REG_KEY>"
    performance_tier = "FTDv30"
    license_caps = [ "MALWARE" ]
    access_policy {
        id = data.fmc_access_policies.access_policy.id
    }
  }

  devices {
    name = "FTD2"
    hostname = "<FTD IP ADDRESS>"
    regkey = "<REG_KEY>"
    performance_tier = "FTDv30"
    license_caps = [ "MALWARE" ]
    access_policy {
        id = data.fmc_access_policies.access_policy.id
    }
  }
}