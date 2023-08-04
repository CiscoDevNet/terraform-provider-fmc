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
    name = "10.100.0.218"
}

data "fmc_security_zones" "my_security_zone" {
  name = "TEST2"
}

output "security_zone" {
  value = data.fmc_security_zones.my_security_zone
}

data "fmc_device_vni" "device_vni" {
    device_id = data.fmc_devices.device.id
    name = "vni1"
}
 
output "fmc_device_vni" {
    value = data.fmc_device_vni.device_vni
}

resource "fmc_device_vni" "my_fmc_device_vni" {
    device_id = data.fmc_devices.device.id
    if_name = "VNI1"
    description = "Description Updated"
    security_zone_id= data.fmc_security_zones.my_security_zone.id
    priority = 3
    vnid = 11 
    multicast_groupaddress =  "224.0.0.34"
    segment_id = 4011
    enable_proxy = false
    ipv4 {
       static {
        address = "3.3.3.3"
        netmask = 4
      }
       dhcp {
        enable_default_route_dhcp = false
        dhcp_route_metric = 0
      }
    }
}

 


