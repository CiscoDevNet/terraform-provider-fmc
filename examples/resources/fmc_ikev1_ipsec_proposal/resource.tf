resource "fmc_ikev1_ipsec_proposal" "example" {
  name           = "my_ikev1_ipsec_proposal"
  description    = "IKEv1 IPsec Proposal 1"
  esp_encryption = "AES-192"
  esp_hash       = "SHA"
}
