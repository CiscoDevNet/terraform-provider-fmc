resource "fmc_ikev2_policies" "example" {
  items = {
    my_ikev2_policies = {
      description           = "IKEv2 Policy"
      priority              = 10
      lifetime              = 86400
      integrity_algorithms  = ["SHA-256"]
      encryption_algorithms = ["AES-256"]
      prf_algorithms        = ["SHA-256"]
      dh_groups             = ["14"]
    }
  }
}
