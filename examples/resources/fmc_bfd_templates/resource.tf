resource "fmc_bfd_templates" "example" {
  items = {
    my_bfd_templates = {
      hop_type                           = "SINGLE_HOP"
      echo                               = "ENABLED"
      interval_type                      = "MILLISECONDS"
      multiplier                         = 3
      minimum_transmit                   = 300
      minimum_receive                    = 300
      authentication_type                = "MD5"
      authentication_password            = "ThisIsMySecretPassword"
      authentication_password_encryption = "UN_ENCRYPTED"
      authentication_key_id              = 1
    }
  }
}
