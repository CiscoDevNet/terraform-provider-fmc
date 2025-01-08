resource "fmc_device_bgp" "example" {
  device_id                                = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  ipv4_default_information_orginate        = false
  ipv4_auto_aummary                        = false
  ipv4_bgp_supress_inactive                = false
  ipv4_synchronization                     = false
  ipv4_bgp_redistribute_internal           = false
  ipv4_external_distance                   = 20
  ipv4_internal_distance                   = 200
  ipv4_local_distance                      = 200
  ipv4_forward_packets_over_multipath_ibgp = 1
  ipv4_forward_packets_over_multipath_ebgp = 1
  ipv4_neighbors = [
    {
      neighbor_address      = "10.1.1.1"
      neighbor_remote_as    = "65534"
      neighbor_bfd          = "SINGLE_HOP"
      enable_address_family = true
      neighbor_description  = "My BGP Peer"
    }
  ]
}
