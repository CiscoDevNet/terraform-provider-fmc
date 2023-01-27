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
data "fmc_device_physical_interfaces" "physical_interfaces" {
    id   = data.fmc_devices.device.id
    name = "TenGigabitEthernet0/0"
}
# data "fmc_security_zones" "outside" {
#     name = "outside"
# }

# data "fmc_device_vtep_policies" "vtep_policy" {
#     device_id   = data.fmc_devices.device.id
# }

resource "fmc_device_vtep_policies" "vtep_policy" {
    device_id = data.fmc_devices.device.id
    nve_enable = true
    vtep_entries{
      source_interface{
        id = data.fmc_device_physical_interfaces.physical_interfaces.id
        type = "PhysicalInterface"
        name = "TenGigabitEthernet0/0"
      }
      nve_vtep_id   = 1
      nve_destination_port = 6088
      nve_encapsulation_type = "GENEVE"
    }
    description = "wkjnk"
}

output "vtep" {
    value = fmc_device_vtep_policies.vtep_policy
}

# output "created" {
#     value = fmc_device_vtep_policies.vtep_policy
# }

# output "devicesklms" {
#     value = data.fmc_devices.device
# }

# output "sexurityzopnesd" {
#     value = data.fmc_security_zones.outside
# }

