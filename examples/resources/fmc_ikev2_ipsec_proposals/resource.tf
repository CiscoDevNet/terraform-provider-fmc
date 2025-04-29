resource "fmc_ikev2_ipsec_proposals" "example" {
  items = {
    my_ikev2_ipsec_proposals = {
      description     = "IKEv2 IPsec Proposal 1"
      esp_encryptions = ["AES-256"]
      esp_hashes      = ["SHA-256"]
    }
  }
}
