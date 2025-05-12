resource "fmc_vpn_s2s_ipsec_settings" "example" {
  vpn_s2s_id      = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  crypto_map_type = "STATIC"
  ikev2_mode      = "TUNNEL"
  ikev2_ipsec_proposals = [
    {
      id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      name = "my_ikev2_ipsec_proposal"
    }
  ]
  security_association_strength_enforcement = false
  reverse_route_injection                   = true
  perfect_forward_secrecy                   = true
  perfect_forward_secrecy_modulus_group     = "14"
  lifetime_duration                         = 28800
  lifetime_size                             = 4608000
  validate_incoming_icmp_error_messages     = false
  do_not_fragment_policy                    = "NONE"
  tfc                                       = true
  tfc_burst_bytes                           = 0
  tfc_payload_bytes                         = 0
  tfc_timeout                               = 0
}
