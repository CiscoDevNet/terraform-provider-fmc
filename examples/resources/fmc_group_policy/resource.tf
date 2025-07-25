resource "fmc_group_policy" "example" {
  name                     = "my_group_policy"
  description              = "My Group Policy object"
  vpn_protocol_ssl         = true
  vpn_protocol_ipsec_ikev2 = true
  ipv4_address_pools = [
    {
      id = ""
    }
  ]
  banner                               = "Welcome to the VPN Connection."
  primary_dns_server_host_id           = ""
  secondary_dns_server_host_id         = ""
  primary_wins_server_host_id          = ""
  secondary_wins_server_host_id        = ""
  dhcp_network_scope_network_object_id = ""
  default_domain                       = "example.com"
  ipv4_split_tunnel_policy             = "TUNNEL_ALL"
  ipv6_split_tunnel_policy             = "TUNNEL_ALL"
  split_tunnel_acl_id                  = "12345678-1234-1234-1234-123456"
  split_tunnel_acl_type                = "ExtendedAccessList"
  dns_request_split_tunnel_policy      = "USE_SPLIT_TUNNEL_SETTING"
  split_dns_domain_list                = "example.com,example.org"
  secure_client_profile_id             = "12345678-1234-1234-1234-123456"
  secure_client_management_profile_id  = "12345678-1234-1234-1234-123456"
  secure_client_modules = [
    {
      type            = ""
      profile_id      = "12345678-1234-1234-1234-123456"
      download_module = true
    }
  ]
  ssl_compression                              = ""
  dtls_compression                             = ""
  mtu_size                                     = 1406
  ignore_df_bit                                = true
  keep_alive_messages                          = true
  keep_alive_messages_interval                 = 20
  gateway_dpd                                  = true
  gateway_dpd_interval                         = 30
  client_dpd                                   = true
  client_dpd_interval                          = 30
  client_bypass_protocol                       = false
  ssl_rekey                                    = true
  ssl_rekey_method                             = "NEW_TUNNEL"
  ssl_rekey_interval                           = 60
  client_firewall_private_network_rules_acl_id = "12345678-1234-1234-1234-123456"
  client_firewall_public_network_rules_acl_id  = "12345678-1234-1234-1234-123456"
  secure_client_custom_attributes = [
    {
      id = "12345678-1234-1234-1234-123456"
    }
  ]
  traffic_filter_acl_id                  = "12345678-1234-1234-1234-123456"
  restrict_vpn_to_vlan                   = 100
  access_hours_time_range_id             = "12345678-1234-1234-1234-123456"
  simultaneous_logins_per_user           = 3
  maximum_connection_time                = 3600
  maximum_connection_time_alert_interval = 1
  vpn_idle_timeout                       = 3600
  vpn_idle_timeout_alert_interval        = 1
}
