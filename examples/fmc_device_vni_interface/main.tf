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

data "fmc_devices" "device" {
    name = "ftd1"
}

data "fmc_security_zones" "outside" {
    name = "outside"
}

data "fmc_device_vni_interface" "vni_interfaces" {
    id   = data.fmc_devices.device.id
    name = "vni1"
}
resource "fmc_device_vni_interface" "vni_interfaces" {
    device_id = data.fmc_devices.device.id
    name = "VNI9"
    ifname = "vni"
    vni_id =  5
    vtep_id = 1
    mode = "NONE"
    mtu= 1500
    description = "Not applicable"
    enabled = true
    enable_proxy = true
    security_zone{
      id = data.fmc_security_zones.outside.id
      type = data.fmc_security_zones.outside.type
    }
}

output "vni_interfaces" {
    value = data.fmc_device_vni_interface.vni_interfaces
}

output "vni_interfaces_create" {
    value = fmc_device_vni_interface.vni_interfaces
}


