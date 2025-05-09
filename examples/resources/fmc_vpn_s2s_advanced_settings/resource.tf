resource "fmc_vpn_s2s_advanced_settings" "example" {
  vpn_s2s_id                                                = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  ike_keepalive                                             = "ENABLED"
  ike_keepalive_threshold                                   = 15
  ike_keepalive_retry_interval                              = 5
  ike_identity_sent_to_peers                                = "AUTO_OR_DN"
  ike_peer_identity_validation                              = "DO_NOT_CHECK"
  ike_aggressive_mode                                       = false
  ike_notification_on_tunnel_disconnect                     = true
  ikev2_cookie_challenge                                    = "CUSTOM"
  ikev2_threshold_to_challenge_incoming_cookies             = 55
  ikev2_number_of_sas_allowed_in_negotiation                = 90
  ipsec_fragmentation_before_encryption                     = true
  ipsec_path_maximum_transmission_unit_aging_reset_interval = 30
  nat_keepalive_message_traversal                           = true
  nat_keepalive_message_traversal_interval                  = 20
  vpn_idle_timeout                                          = true
  vpn_idle_timeout_value                                    = 25
  bypass_access_control_traffic_for_decrypted_traffic       = false
  cert_use_map_configured_in_endpoint_to_determine_tunnel   = false
  cert_use_ou_to_determine_tunnel                           = false
  cert_use_ike_identity_to_determine_tunnel                 = false
  cert_use_peer_ip_address_to_determine_tunnel              = false
}
