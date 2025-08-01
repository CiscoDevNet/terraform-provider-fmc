resource "fmc_vpn_ra_ipsec_crypto_map" "example" {
  vpn_ra_id    = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  interface_id = "76d24097-41c4-4558-a4d0"
  ikev2_ipsec_proposals = [
    {
      id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
    }
  ]
  client_services                       = true
  client_services_port                  = 443
  perfect_forward_secrecy               = true
  perfect_forward_secrecy_modulus_group = "14"
  lifetime_duration                     = 28800
  lifetime_size                         = 4608000
  validate_incoming_icmp_error_messages = false
  do_not_fragment_policy                = "NONE"
  tfc                                   = true
  tfc_burst_bytes                       = 0
  tfc_payload_bytes                     = 0
  tfc_timeout                           = 0
}
