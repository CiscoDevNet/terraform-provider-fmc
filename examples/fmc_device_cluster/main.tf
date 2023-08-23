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

# data "fmc_device_cluster" "inside" {
#     name = "ftd_1"
# }
data "fmc_devices" "device1" {
    name = "FTD1"
}
data "fmc_devices" "device2" {
    name = "FTD2"
}
data "fmc_devices" "device3" {
    name = "FTD3"
}
data "fmc_device_physical_interfaces" "ccl_physical_interface" {
    device_id = data.fmc_devices.device1.id
    name = "GigabitEthernet0/0"
}

# Note this will only work on VMware, not on public cloud.
resource "fmc_device_cluster" "cluster" {
  name = "test-cluster"
  control_device {
    cluster_node_bootstrap {
      priority = 1
      cclip = "10.10.11.1"
    }
    device_details  {
      id = data.fmc_devices.device1.id
      name = data.fmc_devices.device1.name
    }
  }
  common_bootstrap {
    ccl_interface {
      id = data.fmc_device_physical_interfaces.ccl_physical_interface.id
      name = data.fmc_device_physical_interfaces.ccl_physical_interface.name
    }
    ccl_network = "10.10.11.0/27"
    vni_network = "10.10.10.0/27"
  }

  data_devices{
      cluster_node_bootstrap{
        priority = 2
        cclip = "10.10.11.2"
      }
      device_details{
        id = data.fmc_devices.device2.id
        name = data.fmc_devices.device2.name
      }
  }
  data_devices{
      cluster_node_bootstrap{
        priority = 3
        cclip = "10.10.11.3"
      }
      device_details{
        id = data.fmc_devices.device3.id
        name = data.fmc_devices.device3.name
      }
  }
}
output "data_cluster"{
    value = fmc_device_cluster.cluster
}
# output "data_cluster"{
#     value = data.fmc_device_cluster.inside
# }
