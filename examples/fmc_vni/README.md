# Steps to use VNI Interface Data Source and Resource Provider

## Datasource
VNI details can be retervied based on VNI Name & Device ID, below is example

#### Terraform File
```
data "fmc_devices" "device" {
    name = "10.100.0.218"
}

data "fmc_security_zones" "my_security_zone" {
  name = "TEST2"
}

data "fmc_device_vni" "device_vni" {
    device_id = data.fmc_devices.device.id
    name = "vni11"
}
 
 
output "fmc_device_vni" {
    value = data.fmc_device_vni.device_vni
}
```
#### After applying you will see below output
![image](https://user-images.githubusercontent.com/23295769/233247843-e1b13c50-ca0b-411a-94af-bed78b3278a9.png)

## Resource Provider
Using resource provider you can create, update & destroy the VNI

#### Create VNI 

```

data "fmc_devices" "device" {
    name = "10.100.0.218"
}

data "fmc_security_zones" "my_security_zone" {
  name = "TEST2"
}

resource "fmc_device_vni" "my_fmc_device_vni" {
    device_id = data.fmc_devices.device.id
    if_name = "VNI1"
    description = "Description"
    security_zone_id= data.fmc_security_zones.my_security_zone.id
    priority = 3
    vnid = 2
    multicast_groupaddress =  "224.0.0.34"
    segment_id = 4011
    enable_proxy = false
    ipv4 {
       static {
        address = "3.3.3.3"
        netmask = 4
      }
    }
}
```
#### After applying you will see below output
![image](https://user-images.githubusercontent.com/23295769/233249010-1a33960e-052c-4eea-a768-0aabad2908f2.png)

#### Update VNI
Make some change to resource provider and apply, in below example we have update the Description & changed the IPv4 from static to dhcp

![image](https://user-images.githubusercontent.com/23295769/233249466-10efac0b-e695-4811-9833-176828294aea.png)

#### Destroy VNI
Using destroy command VNI interface will destroy the VNI interface which was created

![image](https://user-images.githubusercontent.com/23295769/233249600-82f0e28e-414d-4e17-814f-edc9da269d95.png)






