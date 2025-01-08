resource "fmc_bgd_template" "example" {
  name             = "BFD_Template1"
  hop_type         = "SINGLE_HOP"
  echo             = "ENABLED"
  interval_time    = "MILLISECONDS"
  min_transmit     = 100000
  tx_rx_multiplier = 3
  min_receive      = 100000
}
