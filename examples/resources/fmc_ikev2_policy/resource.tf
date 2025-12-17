resource "fmc_ikev2_policy" "example" {
  name                  = "my_ikev2_policy"
  description           = "IKEv2 Policy"
  priority              = 10
  lifetime              = 86400
  integrity_algorithms  = ["SHA-256"]
  encryption_algorithms = ["AES-256"]
  prf_algorithms        = ["SHA-256"]
  dh_groups             = ["14"]
}
