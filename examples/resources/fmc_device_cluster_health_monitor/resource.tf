resource "fmc_device_cluster_health_monitor" "example" {
  cluster_id                                       = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  health_check                                     = true
  hold_time                                        = 3.0
  debounce_time                                    = 9000
  data_interface_auto_rejoin_attempts              = 3
  data_interface_auto_rejoin_interval              = 5
  data_interface_auto_rejoin_interval_variation    = 2
  cluster_interface_auto_rejoin_attempts           = -1
  cluster_interface_auto_rejoin_interval           = 5
  cluster_interface_auto_rejoin_interval_variation = 1
  system_auto_rejoin_attempts                      = 3
  system_auto_rejoin_interval                      = 5
  system_auto_rejoin_interval_variation            = 2
  unmonitored_interfaces                           = ["GigabitEthernet0/1"]
  service_application_monitoring                   = true
}
