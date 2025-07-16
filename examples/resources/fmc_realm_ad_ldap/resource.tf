resource "fmc_realm_ad_ldap" "example" {
  name               = "my_ldap_realm"
  description        = "My realm"
  realm_type         = "LDAP"
  ad_primary_domain  = "example.com"
  directory_username = "user@example.com"
  directory_password = "my_password"
  base_dn            = "dc=example,dc=com"
  group_dn           = "ou=users,dc=example,dc=com"
  directory_configurations = [
    {
      hostname                        = "ldap.example.com"
      port                            = 389
      encryption                      = "LDAPS"
      encryption_certificate          = "1234-5678-90AB-CDEF12345678"
      use_routing_to_select_interface = true
      interface_group_id              = "1234-5678-90AB-CDEF12345678"
    }
  ]
}
