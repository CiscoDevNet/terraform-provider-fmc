resource "fmc_sla_monitor" "example" {
  name              = "my_sla_monitor"
  description       = "My SLA Monitor object"
  sla_monitor_id    = 10
  timeout           = 5000
  frequency         = 60
  threshold         = 1000
  data_size         = 28
  tos               = 20
  number_of_packets = 1
  monitor_address   = "10.10.10.1"
  selected_interfaces = [
    {
      id = "12345678-1234-1234-1234-123456789012"
    }
  ]
}
