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

data "fmc_devices" "ftd" {
    name = "ftd.adyah.cisco"
}

resource "fmc_ftd_deploy" "ftd" {
    device = data.fmc_devices.ftd.id
    ignore_warning = false
    force_deploy = false
}
