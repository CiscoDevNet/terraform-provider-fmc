resource "fmc_vpn_s2s_endpoints" "example" {
  vpn_s2s_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  items = {
    my_ftd_01 = {
      peer_type                        = "PEER"
      is_extranet                      = false
      device_id                        = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      interface_id                     = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      interface_ipv6_address           = "2001:db8::1"
      connection_type                  = "BIDIRECTIONAL"
      allow_incoming_ikev2_routes      = false
      send_tunnel_interface_ip_to_peer = false
      protected_networks = [
        {
          id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
        }
      ]
      enable_nat_traversal           = false
      enable_nat_exemption           = false
      inside_interface_id            = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      enable_reverse_route_injection = false
      local_identity_type            = "EMAIL"
      local_identity_string          = "me@cisco.com"
      vpn_filter_acl_id              = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      override_remote_vpn_filter     = false
      remote_vpn_filter_acl_id       = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      backup_interface_id            = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      backup_local_identity_type     = "EMAIL"
      backup_local_identity_string   = "me@cisco.com"
    }
  }
}
