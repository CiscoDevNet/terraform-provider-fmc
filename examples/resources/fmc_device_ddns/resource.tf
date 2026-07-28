resource "fmc_device_ddns" "example" {
  device_id                    = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  dhcp_client_request_type     = "BOTH_A_AND_PTR_RECORD"
  dhcp_client_broadcast        = true
  dynamic_dns_update           = "BOTH_A_AND_PTR_RECORD"
  dhcp_client_request_override = true
  ddns_update_methods = [
    {
      name                   = "my_ddns_method"
      method                 = "DDNS"
      update_interval_day    = 0
      update_interval_hour   = 0
      update_interval_minute = 5
      update_interval_second = 0
      update_records         = "BOTH_A_AND_PTR_RECORDS"
    }
  ]
  ddns_interface_settings = [
    {
      interface_id                 = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      interface_name               = "GigabitEthernet0/1"
      interface_type               = "PhysicalInterface"
      method_name                  = "my_ddns_method"
      hostname                     = "host.example.com"
      dhcp_client_request_type     = "NO_UPDATE"
      dynamic_dns_update           = "BOTH_A_AND_PTR_RECORD"
      dhcp_client_request_override = true
    }
  ]
}
