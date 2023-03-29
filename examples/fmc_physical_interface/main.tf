
data "fmc_devices" "device" {
    name = "10.100.0.218"
}

data "fmc_device_physical_interfaces" "device_physical_interface" {
    device_id = data.fmc_devices.device.id
    name = "TenGigabitEthernet0/0"
}

data "fmc_security_zones" "my_security_zone" {
  name = "TEST2"
}

output "existing_device" {
  value = data.fmc_devices.device
}

output "physical_interface" {
    value = data.fmc_device_physical_interfaces.device_physical_interface
}

output "existing_zone" {
  value = data.fmc_security_zones.my_security_zone
}

resource "fmc_device_physical_interfaces" "my_fmc_device_physical_interfaces" {
    device_id = data.fmc_devices.device.id
    physical_interface_id= data.fmc_device_physical_interfaces.device_physical_interface.id
    name =   data.fmc_device_physical_interfaces.device_physical_interface.name
    security_zone_id= data.fmc_security_zones.my_security_zone.id
    if_name = "IFNameChangedDuringDemo"
    description = "DescAddedByDuringDemo"
    mtu =  1700
    mode = "NONE"
    ipv4_static_address = "10.20.220.45"
    ipv4_static_netmask = 24
    ipv4_dhcp_enabled = false
    ipv4_dhcp_route_metric = 1
 }


