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


data "fmc_devices" "device1" {
    name = "FTD"
}
data "fmc_network_objects" "anyipv4"{
  name = "any-ipv4"
}
data "fmc_network_objects" "rand"{
  name = "rand-net"
}
data "fmc_host_objects" "igw" {
  name = "igw"
}

resource "fmc_staticIPv4_route" "route1" {
  device_id  = data.fmc_devices.device1.id
  interface_name = "outside" # Such interface should be configured on the device first.
  metric_value = 23
  is_tunneled = false
  selected_networks {
      id = data.fmc_network_objects.rand.id
      type = data.fmc_network_objects.rand.type
      name = data.fmc_network_objects.rand.name
  }
  gateway {
    object {
      id = data.fmc_host_objects.igw.id
      type = data.fmc_host_objects.igw.type
      name = data.fmc_host_objects.igw.name
    }
  }
  route_tracking {
    id = "<ID OF THE RESOURCE>"
    type = "SLAMonitor"
    name = "<NAME OF THE RESOURCE>"
  }
}

resource "fmc_staticIPv4_route" "route2" {
  device_id  = data.fmc_devices.device1.id
  interface_name = "inside"
  metric_value = 23
  is_tunneled = false
  selected_networks {
      id = data.fmc_network_objects.anyipv4.id
      type = data.fmc_network_objects.anyipv4.type
      name = data.fmc_network_objects.anyipv4.name
  }
  gateway {
    object {
      id = data.fmc_host_objects.igw.id
      type = data.fmc_host_objects.igw.type
      name = data.fmc_host_objects.igw.name
    }
  }
}

output "route1"{
  value = fmc_staticIPv4_route.route1
}
output "route2"{
  value = fmc_staticIPv4_route.route2
}

# EXAMPLE OF DATA BLOCK
# data "fmc_staticIPv4_route" "route"{
#   device_id  = data.fmc_devices.device1.id
#   network_name = "rand-net"
# }

# output "route" {
#     value = data.fmc_staticIPv4_route.route
# }