resource "fmc_identity_policy" "example" {
  name                                          = "my_identity_policy"
  description                                   = "My Identity Policy"
  identity_mapping_filter_network_object_id     = "12345678-90ab-cdef-1234-567890abcdef"
  identity_mapping_filter_network_object_type   = "Network"
  active_authentication_server_certificate_id   = "12345678-90ab-cdef-1234-567890abcdef"
  active_authentication_server_certificate_type = "PKI_InternalCert"
  active_authentication_redirect_fqdn_id        = "12345678-90ab-cdef-1234-567890abcdef"
  active_authentication_redirect_fqdn_type      = "FQDN"
  active_authentication_redirect_port           = 885
  active_authentication_maximum_login_attempts  = 3
  active_authentication_session_sharing         = true
  active_authentication_page                    = "DEFAULT"
  categories = [
    {
      name = "category_1"
    }
  ]
  rules = [
    {
      name                           = "rule_1"
      enabled                        = true
      category                       = "category_1"
      authentication_type            = "PASSIVE"
      authentication_realm_id        = "12345678-90ab-cdef-1234-567890abcdef"
      authentication_realm_type      = "IdentityRealm"
      guest_access_fallback          = false
      active_authentication_fallback = false
      authentication_protocol        = ""
      excluded_user_agents = [
        {
          id   = "12345678-90ab-cdef-1234-567890abcdef"
          type = "Application"
        }
      ]
      source_zones = [
        {
          id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
          type = "SecurityZone"
        }
      ]
      destination_zones = [
        {
          id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
          type = "SecurityZone"
        }
      ]
      source_network_literals = [
        {
          type  = "Network"
          value = "10.1.1.0/24"
        }
      ]
      source_network_objects = [
        {
          id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
          type = "Network"
        }
      ]
      destination_network_literals = [
        {
          type  = "Network"
          value = "10.2.2.0/24"
        }
      ]
      destination_network_objects = [
        {
          id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
          type = "Network"
        }
      ]
      vlan_tag_literals = [
        {
          type      = "VlanTagLiteral"
          start_tag = "11"
          end_tag   = "22"
        }
      ]
      vlan_tag_objects = [
        {
          id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
          type = "VlanTag"
        }
      ]
      source_port_literals = [
        {
          type     = "PortLiteral"
          protocol = "6"
          port     = "80"
        }
      ]
      source_port_objects = [
        {
          id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
          type = "ProtocolPortObject"
        }
      ]
      destination_port_literals = [
        {
          type     = "PortLiteral"
          port     = "80"
          protocol = "6"
        }
      ]
      destination_port_objects = [
        {
          id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
          type = "ProtocolPortObject"
        }
      ]
    }
  ]
}
