
data "fmc_devices" "device" {
    name = "10.100.0.218"
}

data "fmc_security_zones" "my_security_zone" {
  name = "TEST2"
}

output "security_zone" {
  value = data.fmc_security_zones.my_security_zone
}


resource "fmc_device_sub_interface" "my_fmc_device_sub_interface" {
    device_id = data.fmc_devices.device.id
    name = "TenGigabitEthernet0/0"
    type = "SubInterface"
    description = "DescriptionUpdated"
    sub_interface_id = 3
    if_name = "IFNameUpdated"
    mtu = 1504
    priority = 4
    management_only = false
    security_zone_id = data.fmc_security_zones.my_security_zone.id
    vlan_id = 4
    enabled = true
    ipv4 {
      static {
        address = "3.3.3.3"
        netmask = 4
      }
       dhcp {
        enable_default_route_dhcp = true
        dhcp_route_metric = 0
      }
    }
}