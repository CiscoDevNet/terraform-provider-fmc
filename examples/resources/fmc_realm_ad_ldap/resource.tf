resource "fmc_realm_ad_ldap" "example" {
  name                                    = "my_ldap_realm"
  enabled                                 = true
  description                             = "My realm"
  realm_type                              = "LDAP"
  directory_username                      = "user@example.com"
  directory_password                      = "my_password"
  base_dn                                 = "DC=example,DC=com"
  group_dn                                = "CN=users,DC=example,DC=com"
  included_users                          = ["user1"]
  included_groups                         = ["group1"]
  excluded_users                          = ["user2"]
  excluded_groups                         = ["group2"]
  update_hour                             = 2
  update_interval                         = "4"
  group_attribute                         = "member"
  timeout_ise_and_passive_indentity_users = 1440
  timeout_terminal_server_agent_users     = 1440
  timeout_captive_portal_users            = 1440
  timeout_failed_captive_portal_users     = 1440
  timeout_guest_captive_portal_users      = 1440
  directory_servers = [
    {
      hostname                        = "ldap.example.com"
      port                            = 389
      encryption_protocol             = "LDAPS"
      ca_certificate_id               = "12345678-1234-1234-1234-123456789012"
      use_routing_to_select_interface = false
    }
  ]
}
