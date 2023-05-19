
data "fmc_devices" "device" {
    name = "10.100.0.218"
}
  
# data "fmc_device_vtep_policies" "fmc_device_vtep" {
#      device_id = data.fmc_devices.device.id
# }
 
# output "fmc_device_vtep" {
#     value = data.fmc_device_vtep_policies.fmc_device_vtep
# }


resource "fmc_device_vtep" "my_fmc_device_vtep" {
    device_id = data.fmc_devices.device.id
    nve_enabled = true
    source_interface_id = "02E8B76C-CB94-0ed3-0000-008589935642" // GET This using PHYSICAL INTERFACE Datasource
    nve_vtp_id = 1
    nve_encapsulation_type = "GENEVE"
    nve_destination_port = 6064
}