resource "fmc_ikev2_ipsec_proposals" "example" {
  items = {
    my_ikev2_ipsec_proposals = {
      description    = "IKEv2 IPsec Proposal 1"
      esp_encryption = ["AES-256"]
      esp_hash       = ["SHA-256"]
    }
  }
}
