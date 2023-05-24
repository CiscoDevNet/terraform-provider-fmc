terraform {
  required_providers {
    fmc = {
      source = "CiscoDevNet/fmc"
      # version = "0.1.1"
    }
  }
}

provider "fmc" {
  fmc_username = var.fmc_username
  fmc_password = var.fmc_password
  fmc_host = var.fmc_host
  fmc_insecure_skip_verify = var.fmc_insecure_skip_verify
}

data "fmc_devices" "device" {
    name = "ftd.adyah.cisco"
}

output "existing_device" {
    value = data.fmc_devices.device
}

data "fmc_access_policies" "access_policy"{
  name="test-acp"
}

resource "fmc_devices" "device"{
  name = "FTD"
  hostname = "<Device IP>"
  regkey = "cisco"
  performance_tier = "FTDv30"
  license_caps = [ "MALWARE"]
  access_policy {
      id = data.fmc_access_policies.access_policy.id
  }
}

output "fmc_devicess" {
    value = fmc_devices.device
}
