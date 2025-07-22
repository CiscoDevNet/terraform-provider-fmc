resource "fmc_radius_server_group" "example" {
  name                            = "my_radius_server_group"
  description                     = "My RADIUS Server Group object"
  group_accounting_mode           = "SINGLE"
  retry_interval                  = 10
  realm_id                        = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  authorize_only                  = true
  interim_account_update_interval = 24
  dynamic_authorization           = true
  dynamic_authorization_port      = 1700
  merge_downloadable_acl_order    = "MERGE_DACL_BEFORE_AV_PAIR_ACL"
  radius_servers = [
    {
      hostname                                    = "10.10.10.10"
      radius_server_enabled_message_authenticator = true
      authentication_port                         = 1812
      key                                         = "my_secret_key"
      accounting_port                             = 1813
      timeout                                     = 10
      use_routing_to_select_interface             = true
    }
  ]
}
