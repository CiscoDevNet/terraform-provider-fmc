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
    name = "FTD1"
}

resource "fmc_security_zone" "outside" {
  name            = "outside"
  interface_mode  = "ROUTED"
}


# data "fmc_device_subinterfaces" "sub_interfaces" {
#     device_id   = data.fmc_devices.device.id 
#     subinterface_id = 12  # Example: Represented by the integer value after ' . ' in GigabitEthernet0/1.12
# }

resource "fmc_device_subinterfaces" "sub" {
    device_id   = data.fmc_devices.device.id
    ifname = "Testing12"
    subinterface_id = 12345
    vlan_id = 80
    name = "GigabitEthernet0/1"
    mode = "NONE"
    mtu = 1600
    enabled = true
    priority = 69
    security_zone_id = fmc_security_zone.outside.id
    ipv4_dhcp_enabled = false
    ipv4_dhcp_route_metric = 1

    enable_ipv6 = true
    ipv6_address = "2001:10:240:ac::a"
    ipv6_prefix = "124"
    ipv6_enforce_eui = false

}

output "new_subinterface_sub" {
    value = fmc_device_subinterfaces.sub
}

# output "old_physical_interfaces" {
#     value = data.fmc_device_subinterfaces.sub_interfaces
# }


