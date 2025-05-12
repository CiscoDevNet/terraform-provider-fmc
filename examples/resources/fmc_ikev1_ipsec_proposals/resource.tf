resource "fmc_ikev1_ipsec_proposals" "example" {
  items = {
    my_ikev1_ipsec_proposals = {
      description    = "IKEv1 IPsec Proposal 1"
      esp_encryption = "AES-192"
      esp_hash       = "SHA"
    }
  }
}
