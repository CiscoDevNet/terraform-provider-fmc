resource "fmc_sinkhole" "example" {
  name                = "my_sinkhole"
  ipv4_address        = "10.1.1.1"
  ipv6_address        = "2001:db8::1"
  action              = "MONITOR"
  log_connection_type = "SINKHOLE_CNC"
}
