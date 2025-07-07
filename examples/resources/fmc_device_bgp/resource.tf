resource "fmc_device_bgp" "example" {
  device_id                         = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  ipv4_default_information_orginate = false
  ipv4_auto_summary                 = false
  ipv4_suppress_inactive_routes     = false
  ipv4_synchronization              = false
  ipv4_redistribute_ibgp_into_igp   = false
  ipv4_external_distance            = 20
  ipv4_internal_distance            = 200
  ipv4_local_distance               = 200
  ipv4_number_of_ibgp_paths         = 1
  ipv4_number_of_ebgp_paths         = 1
  ipv4_neighbors = [
    {
      address                    = "10.1.1.1"
      remote_as                  = "65534"
      bfd_fallover               = "SINGLE_HOP"
      update_source_interface_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      enable_address             = true
      as_override                = false
      shutdown_administratively  = false
      description                = "My BGP Peer"
    }
  ]
}
