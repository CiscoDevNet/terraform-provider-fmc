resource "fmc_icmpv4_object" "example" {
  icmp_type   = 3
  code        = 0
  name        = "icmpv4_net_unreachable"
  description = "ICMPv4 network unreachable response, type 3, code 0"
}
