resource "fmc_group_policy" "example" {
  name                        = "my_group_policy"
  description                 = "My Group Policy object"
  enable_ssl_protocol         = true
  enable_ipsec_ikev2_protocol = true
  ipv4_address_pools = [
    {
      id = ""
    }
  ]
  banner                                   = "Welcome to the VPN Connection."
  primary_dns_server                       = ""
  secondary_dns_server                     = ""
  primary_wins_server                      = ""
  secondary_wins_server                    = ""
  dhcp_scope_network_id                    = ""
  default_domain                           = ""
  ipv4_split_tunnel_policy                 = "TUNNEL_ALL"
  ipv6_split_tunnel_policy                 = "TUNNEL_ALL"
  split_tunnel_acl_id                      = "12345678-1234-1234-1234-123456"
  split_d_n_s_request_policy               = "TUNNEL_ALL"
  split_d_n_s_domain_list                  = ""
  secure_client_client_profile_id          = "12345678-1234-1234-1234-123456"
  secure_client_management_profile_id      = "12345678-1234-1234-1234-123456"
  ssl_compression                          = ""
  dtls_compression                         = ""
  mtu_size                                 = 1406
  ignore_df_bit                            = true
  enable_keep_alive_messages               = true
  keep_alive_message_interval              = 20
  enable_gateway_dpd                       = true
  gateway_dpd_interval                     = 30
  enable_client_dpd                        = true
  client_dpd_interval                      = 30
  client_bypass_protocol                   = false
  enable_ssl_rekey                         = true
  rekey_method                             = "NEW_TUNNEL"
  rekey_interval                           = 60
  client_firewall_private_network_rules_id = "12345678-1234-1234-1234-123456"
  client_firewall_public_network_rules_id  = "12345678-1234-1234-1234-123456"
  custom_attributes = [
    {
      id   = "12345678-1234-1234-1234-123456"
      type = "myAnyConnectAttribute"
    }
  ]
  traffic_filter_acl_id              = "12345678-1234-1234-1234-123456"
  restrict_vpn_to_vlan_id            = 100
  access_hours_id                    = "12345678-1234-1234-1234-123456"
  simultaneous_login_per_user        = 3
  max_connection_time                = 3600
  max_connection_time_alert_interval = 1
  vpn_idle_timeout                   = 3600
  vpn_idle_timeout_alert_interval    = 1
}
