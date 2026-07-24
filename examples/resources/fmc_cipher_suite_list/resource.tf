resource "fmc_cipher_suite_list" "example" {
  name = "my_cipher_suite_list"
  cipher_suites = [
    {
      name = "TLS_RSA_WITH_AES_128_CBC_SHA"
    }
  ]
}
