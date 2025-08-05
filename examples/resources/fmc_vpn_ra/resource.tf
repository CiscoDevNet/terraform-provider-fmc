resource "fmc_vpn_ra" "example" {
  name                     = "my_ftd_ra_vpn"
  description              = "My Remote Access VPN Configuration"
  protocol_ssl             = true
  protocol_ipsec_ikev2     = true
  local_realm_id           = "12345678-1234-1234-1234-123456789012"
  dynamic_access_policy_id = "12345678-1234-1234-1234-123456"
  access_interfaces = [
    {
      id                                = "12345678-1234-1234-1234-123456789012"
      protocol_ipsec_ikev2              = true
      protocol_ssl                      = true
      protocol_ssl_dtls                 = true
      interface_specific_certificate_id = "12345678-1234-1234-1234-123456"
    }
  ]
  allow_users_to_select_connection_profile           = true
  web_access_port_number                             = 443
  dtls_port_number                                   = 443
  ssl_global_identity_certificate_id                 = "12345678-1234-1234-1234-123456"
  ipsec_ikev2_identity_certificate_id                = "12345678-1234-1234-1234-123456"
  service_access_object_id                           = "12345678-1234-1234-1234-123456"
  bypass_access_control_policy_for_decrypted_traffic = false
  secure_client_images = [
    {
      id               = "12345678-1234-1234-1234-123456789012"
      operating_system = "WINDOWS"
    }
  ]
  external_browser_package_id = "12345678-1234-1234-1234-123456789012"
  group_policies = [
    {
      id = "12345678-1234-1234-1234-1234567890aa"
    }
  ]
  ikev2_policies = [
    {
      id = "12345678-1234-1234-1234-123456789012"
    }
  ]
}
