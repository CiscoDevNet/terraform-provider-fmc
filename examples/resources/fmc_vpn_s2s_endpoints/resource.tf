resource "fmc_vpn_s2s_endpoints" "example" {
  vpn_s2s_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  items = {
    my_ftd_01 = {
      peer_type                                = "PEER"
      extranet_device                          = false
      device_id                                = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      interface_id                             = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      interface_ipv6_address                   = "2001:db8::1"
      interface_public_ip_address              = "10.1.1.1"
      connection_type                          = "BIDIRECTIONAL"
      allow_incoming_ikev2_routes              = false
      send_virtual_tunnel_interface_ip_to_peer = false
      protected_networks = [
        {
          id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
        }
      ]
      local_identity_type          = "EMAILID"
      local_identity_string        = "me@cisco.com"
      backup_interface_id          = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      backup_local_identity_type   = "EMAILID"
      backup_local_identity_string = "me@cisco.com"
    }
  }
}
