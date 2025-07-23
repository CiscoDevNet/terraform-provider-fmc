resource "fmc_trusted_certificate_authority" "example" {
  name        = "my_trusted_ca"
  certificate = "-----BEGIN CERTIFICATE-----\nMII(...)\n-----END CERTIFICATE-----"
}
