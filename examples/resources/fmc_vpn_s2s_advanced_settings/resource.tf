resource "fmc_vpn_s2s_advanced_settings" "example" {
  vpn_s2s_id                                              = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  ike_keepalive                                           = "ENABLED"
  ike_keepalive_threshold                                 = 10
  ike_keepalive_retry_interval                            = 2
  ike_identity_sent_to_peers                              = "AUTO_OR_DN"
  ike_peer_identity_validation                            = "REQUIRED"
  ike_enable_aggressive_mode                              = false
  ike_enable_notification_on_tunnel_disconnect            = false
  ikev2_cookie_challenge                                  = "CUSTOM"
  ikev2_threshold_to_challenge_incoming_cookies           = 50
  ikev2_percentage_of_sas_allowed_in_negotiation          = 100
  ipsec_enable_fragmentation_before_encryption            = true
  nat_keepalive_message_traversal_interval                = 20
  vpn_idle_timeout                                        = 10
  bypass_access_control_traffic_for_decrypted_traffic     = false
  use_cert_map_configured_in_endpoint_to_determine_tunnel = false
  use_certificate_ou_to_determine_tunnel                  = false
  use_ike_identity_ou_to_determine_tunnel                 = false
  use_peer_ip_address_to_determine_tunnel                 = false
}
