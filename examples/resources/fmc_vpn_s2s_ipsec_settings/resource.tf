resource "fmc_vpn_s2s_ipsec_settings" "example" {
  vpn_s2s_id      = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  crypto_map_type = "STATIC"
  ikev2_mode      = "TUNNEL"
  ikev1_ipsec_proposals = [
    {
      id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      name = "my_ikev1_ipsec_proposal"
    }
  ]
  ikev2_ipsec_proposals = [
    {
      id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      name = "my_ikev2_ipsec_proposal"
    }
  ]
  enable_sa_strength_enforcement        = false
  enable_reverse_route_injection        = false
  enable_perfect_forward_secrecy        = true
  perfect_forward_secrecy_modulus_group = "2"
  lifetime_duration                     = 28800
  lifetime_size                         = 4608000
  validate_incoming_icmp_error_messages = false
  do_not_fragment_policy                = "NONE"
  enable_tfc_packets                    = false
}
