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
    name = "FTD1"
}
  
# data "fmc_device_vtep_policies" "fmc_device_vtep" {
#      device_id = data.fmc_devices.device.id
# }
 
# output "fmc_device_vtep" {
#     value = data.fmc_device_vtep_policies.fmc_device_vtep
# }

data "fmc_device_physical_interfaces" "zero_physical_interface" {
    device_id = data.fmc_devices.device.id
    name = "TenGigabitEthernet0/0"
}
resource "fmc_host_objects" "test1" {
  name        = "test1"
  value       = "172.16.1.1"
}
resource "fmc_host_objects" "test2" {
  name        = "test2"
  value       = "172.16.2.1"
}

resource "fmc_network_group_objects" "TestPrivateGroup" {
  name = "TestPrivateGroup"
  description = "Testing groups"
  objects {
      id = fmc_host_objects.test1.id
      type = fmc_host_objects.test1.type
  }
  objects {
      id = fmc_host_objects.test2.id
      type = fmc_host_objects.test2.type
  }
}
resource "fmc_device_vtep" "my_fmc_device_vtep" {
    device_id = data.fmc_devices.device.id
    nve_enabled = true

    nve_vtep_id = 1
    nve_encapsulation_type = "VXLAN"
    nve_destination_port = 4789
    source_interface_id = data.fmc_device_physical_interfaces.zero_physical_interface.id 
    
    nve_neighbor_discovery_type= "NONE"
    # neighbor_addr_object_id = fmc_network_group_objects.TestPrivateGroup.id
    # neighbor_addr_object_overridable = true

    # neighbor_addr_literal_type = "Host"
    # neighbor_addr_literal_value = "1.1.0.1"
}