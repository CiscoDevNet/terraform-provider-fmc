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
    name = "ftd2"
}

data "fmc_security_zones" "outside" {
    name = "outside"
}

data "fmc_device_sub_interfaces" "sub_interfaces" {
    device_id   = data.fmc_devices.device.id
    name = "TenGigabitEthernet0/0"
}
# resource "fmc_physical_interfaces" "physical_interfaces" {
#     physicalinterface_id = data.fmc_device_physical_interfaces.physical_interfaces.id
#     ifname = "ljkgjfh"
#     device_id   = data.fmc_devices.device.id
#     name = "outside2022"
#     mode = "NONE"
#     mtu = 1500
#     enabled = true
#     security_zone {
#             id   =  data.fmc_security_zones.outside.id
#             type =  data.fmc_security_zones.outside.type
#     }
#     ipv4{
#         dhcp{
#           enable_default_route_dhcp = true
#           dhcp_route_metric = 1
#         }
#     }
# }

# output "new_physical_interfaces" {
#     value = fmc_device_physical_interfaces.physical_interfaces
# }

output "old_physical_interfaces" {
    value = data.fmc_device_sub_interfaces.sub_interfaces
}


