resource "fmc_bfd_template" "example" {
  name                               = "BFD_Template1"
  hop_type                           = "SINGLE_HOP"
  echo                               = "ENABLED"
  interval_time                      = "MILLISECONDS"
  min_transmit                       = 300
  tx_rx_multiplier                   = 3
  min_receive                        = 300
  authentication_password            = "ThisIsMySecretPassword"
  authentication_key_id              = 1
  authentication_type                = "MD5"
  authentication_password_encryption = "UN_ENCRYPTED"
}
