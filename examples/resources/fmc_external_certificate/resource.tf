resource "fmc_external_certificate" "example" {
  name        = "my_certificate_external"
  certificate = "-----BEGIN CERTIFICATE-----\nMII(...)\n-----END CERTIFICATE-----"
}
