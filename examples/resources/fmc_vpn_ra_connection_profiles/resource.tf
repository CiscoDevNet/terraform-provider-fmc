resource "fmc_vpn_ra_connection_profiles" "example" {
  vpn_ra_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  items = {
    my_connection_profile = {
      group_policy_id = "12345678-1234-1234-1234-123456"
      ipv4_address_pools = [
        {
          id = "12345678-1234-1234-1234-123456"
        }
      ]
      ipv6_address_pools = [
        {
          id = "12345678-1234-1234-1234-123456"
        }
      ]
      dhcp_servers = [
        {
          id = "12345678-1234-1234-1234-123456"
        }
      ]
      authentication_method                                                        = "AAA_AND_CLIENT_CERTIFICATE"
      multiple_certificate_authentication                                          = true
      primary_authentication_server_use_local                                      = false
      primary_authentication_server_id                                             = "12345678-1234-1234-1234-123456"
      primary_authentication_server_type                                           = "RadiusServerGroup"
      primary_authentication_fallback_to_local                                     = true
      saml_and_certificate_username_must_match                                     = true
      primary_authentication_prefill_username_from_certificate                     = true
      primary_authentication_prefill_username_from_certificate_map_primary_field   = "CN_COMMMON_NAME"
      primary_authentication_prefill_username_from_certificate_map_secondary_field = "NONE"
      primary_authentication_prefill_username_from_certificate_map_entire_dn       = false
      primary_authentication_hide_username_in_login_window                         = true
      secondary_authentication                                                     = true
      secondary_authentication_server_use_local                                    = false
      secondary_authentication_server_id                                           = "12345678-1234-1234-1234-123456"
      secondary_authentication_server_type                                         = "Realm"
      secondary_authentication_fallback_to_local                                   = false
      secondary_authentication_prompt_for_username                                 = false
      secondary_authentication_use_primary_authentication_username                 = true
      authorization_server_id                                                      = "12345678-1234-1234-1234-123456"
      authorization_server_type                                                    = "RadiusServerGroup"
      allow_connection_only_if_user_exists_in_authorization_database               = false
      accounting_server_id                                                         = "12345678-1234-1234-1234-123456"
      accounting_server_type                                                       = "RadiusServerGroup"
      strip_realm_from_username                                                    = true
      strip_group_from_username                                                    = true
      password_management                                                          = true
      password_management_notify_user_on_password_expiry_day                       = true
    }
  }
}
