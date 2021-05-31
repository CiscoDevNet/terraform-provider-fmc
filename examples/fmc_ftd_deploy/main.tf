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

data "fmc_devices" "ftd" {
    name = "ftd.adyah.cisco"
}

resource "fmc_ftd_deploy" "ftd" {
    device = data.fmc_devices.ftd.id
    ignore_warning = false
    force_deploy = false
}
