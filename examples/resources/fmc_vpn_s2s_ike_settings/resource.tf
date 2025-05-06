resource "fmc_vpn_s2s_ike_settings" "example" {
  vpn_s2s_id                  = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  ikev1_authentication_type   = "MANUAL_PRE_SHARED_KEY"
  ikev1_manual_pre_shared_key = "my_pre_shared_key123"
  ikev1_policies = [
    {
      id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      name = "my_ikev1_policy"
    }
  ]
  ikev2_authentication_type                   = "MANUAL_PRE_SHARED_KEY"
  ikev2_manual_pre_shared_key                 = "my_pre_shared_key123"
  ikev2_enforce_hex_based_pre_shared_key_only = false
  ikev2_policies = [
    {
      id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      name = "my_ikev2_policy"
    }
  ]
}
