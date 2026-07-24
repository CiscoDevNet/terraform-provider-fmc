resource "fmc_cipher_suite_lists" "example" {
  items = {
    my_cipher_suite_lists = {
      cipher_suites = [
        {
          name = "TLS_RSA_WITH_AES_128_CBC_SHA"
        }
      ]
    }
  }
}
