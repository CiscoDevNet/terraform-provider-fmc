resource "fmc_internal_certificate" "example" {
  name        = "my_internal_certificate"
  certificate = "-----BEGIN CERTIFICATE-----\nMII(...)\n-----END CERTIFICATE-----"
  private_key = "-----BEGIN RSA PRIVATE KEY-----\nProc-Type: 4,ENCRYPTED\nDEK-Info: AES-128-CBC,D55(...)\n-----END RSA PRIVATE KEY-----"
}
