resource "fmc_ikev1_policy" "example" {
  name                  = "my_ikev1_policy"
  description           = "IKEv1 Policy 1"
  priority              = 10
  encryption_algorithm  = "AES-192"
  hash                  = "SHA"
  dh_group              = "14"
  lifetime              = 86400
  authentication_method = "Preshared Key"
}
