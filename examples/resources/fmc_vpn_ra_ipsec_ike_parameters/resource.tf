resource "fmc_vpn_ra_ipsec_ike_parameters" "example" {
  vpn_ra_id                                                 = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  ikev2_identity_sent_to_peer                               = "AUTO_OR_DN"
  ikev2_notification_on_tunnel_disconnect                   = true
  ikev2_do_not_reboot_until_all_sessions_are_terminated     = false
  ikev2_cookie_challenge                                    = "true"
  ikev2_threshold_to_challenge_incoming_cookies             = 50
  ikev2_number_of_sas_allowed_in_negotiation                = 90
  ikev2_maximum_number_of_sas_allowed                       = 50
  ipsec_path_maximum_transmission_unit_aging                = true
  ipsec_path_maximum_transmission_unit_aging_reset_interval = 30
  nat_keepalive_message_traversal                           = true
  nat_keepalive_message_traversal_interval                  = 20
}
