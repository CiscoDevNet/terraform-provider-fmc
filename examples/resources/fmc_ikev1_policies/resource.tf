resource "fmc_ikev1_policies" "example" {
  items = {
    my_ikev1_policies = {
      description           = "IKEv1 Policy 1"
      priority              = 10
      encryption            = "AES-192"
      hash                  = "SHA"
      dh_group              = "14"
      lifetime              = 86400
      authentication_method = "Preshared Key"
    }
  }
}
