resource "fmc_bfd_template" "example" {
  name                    = "BFD_Template1"
  hop_type                = "SINGLE_HOP"
  echo                    = "ENABLED"
  interval_time           = "MILLISECONDS"
  min_transmit            = 300
  tx_rx_multiplier        = 3
  min_receive             = 300
  authentication_password = "Cisco123!"
  authentication_key_id   = 1
  authentication_type     = "MD5"
}
