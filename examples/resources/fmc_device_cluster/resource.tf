resource "fmc_device_cluster" "example" {
  name                          = "MyDeviceClusterName1"
  cluster_key                   = "cisco123"
  control_node_device_id        = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  control_node_vni_prefix       = "10.10.3.0/27"
  control_node_ccl_prefix       = "10.10.4.0/27"
  control_node_interface_id     = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  control_node_interface_name   = "GigabitEthernet0/0"
  control_node_interface_type   = "PhysicalInterface"
  control_node_ccl_ipv4_address = "10.10.4.1"
  control_node_priority         = 1
  data_devices = [
    {
      data_node_device_id        = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      data_node_ccl_ipv4_address = "10.10.4.2"
      data_node_priority         = 2
    }
  ]
}
