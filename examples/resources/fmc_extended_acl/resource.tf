resource "fmc_extended_acl" "example" {
  name        = "extended_acl_1"
  description = "My Extended Access Control List"
  entries = [
    {
      action               = "DENY"
      log_level            = "WARNING"
      logging              = "PER_ACCESS_LIST_ENTRY"
      log_interval_seconds = 120
      source_network_literals = [
        {
          value = "10.1.1.0/24"
          type  = "Network"
        }
      ]
      destination_network_literals = [
        {
          value = "10.2.2.2"
          type  = "Host"
        }
      ]
      source_network_objects = [
        {
          id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
        }
      ]
      destination_network_objects = [
        {
          id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
        }
      ]
      source_port_objects = [
        {
          id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
        }
      ]
      destination_port_objects = [
        {
          id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
        }
      ]
    }
  ]
}
