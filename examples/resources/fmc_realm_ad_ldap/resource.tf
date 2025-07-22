resource "fmc_realm_ad_ldap" "example" {
  name                                = "my_ldap_realm"
  enabled                             = true
  description                         = "My realm"
  realm_type                          = "LDAP"
  ad_primary_domain                   = "example.com"
  ad_join_username                    = "user@example.com"
  ad_join_password                    = "my_password"
  directory_username                  = "user@example.com"
  directory_password                  = "my_password"
  base_dn                             = "DC=example,DC=com"
  included_users                      = ["user1"]
  included_groups                     = ["group1"]
  excluded_users                      = ["user2"]
  excluded_groups                     = ["group2"]
  group_dn                            = "CN=users,DC=example,DC=com"
  group_attribute                     = "member"
  timeout_ise_users                   = 1440
  timeout_terminal_server_agent_users = 1440
  timeout_captive_portal_users        = 1440
  timeout_failed_captive_portal_users = 1440
  timeout_guest_captive_portal_users  = 1440
  directory_server_configurations = [
    {
      hostname                        = "ldap.example.com"
      port                            = 389
      encryption_protocol             = "LDAPS"
      encryption_certificate          = "1234-5678-90AB-CDEF12345678"
      use_routing_to_select_interface = true
      interface_group_id              = "1234-5678-90AB-CDEF12345678"
    }
  ]
}
