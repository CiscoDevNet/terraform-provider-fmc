
data "fmc_devices" "device" {
    name = "10.100.0.218"
}

data "fmc_device_physical_interfaces" "device_physical_interfaces" {
    name = "TenGigabitEthernet0/0"
    id = data.fmc_devices.device.id
}


resource "fmc_device_physical_interfaces" "physical_interfaces" {
    physicalinterface_id = data.fmc_device_physical_interfaces.device_physical_interfaces.id
    ifname = data.fmc_device_physical_interfaces.device_physical_interfaces.name
    device_id   = data.fmc_devices.device.id
    name = "testing_ashish"
    mode = "NONE"
    mtu = 1500
    enabled = true
    # security_zone {
    #         id   =  data.fmc_security_zones.outside.id
    #         type =  data.fmc_security_zones.outside.type
    # }
    ipv4{
        dhcp{
          enable_default_route_dhcp = true
          dhcp_route_metric = 1
        }
    }
}

data "fmc_device_vni_interface" "vni" {
  name = "vni1234"
  id = data.fmc_devices.device.id 
}

resource "fmc_device_vni_interface" "vni_interfaces" {
    #id = data.fmc_devices.device.id
    name = "VNI9"
    #ifname = "vni"
    vni_id =  5
    #vtep_id = 1
    #mode = "NONE"
    #mtu= 1500
    description = "Not applicable"
    enabled = true
    #enable_proxy = true
    # security_zone{
    #   id = data.fmc_security_zones.outside.id
    #   type = data.fmc_security_zones.outside.type
    # }
}


output "existing_device" {
  value = data.fmc_devices.device
}

output "physical_interface" {
    value = data.fmc_device_physical_interfaces.device_physical_interfaces
}

output "name" {
  value = fmc_device_physical_interfaces.physical_interfaces
}

output "vni" {
  value = data.fmc_device_vni_interface.vni 
}

output "vni_interfaces" {
  value = fmc_device_vni_interface.vni_interfaces
  
}