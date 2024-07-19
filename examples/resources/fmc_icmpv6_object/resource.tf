resource "fmc_icmpv6_object" "example" {
  icmp_type   = 1
  code        = 3
  name        = "icmpv6_addr_unreachable"
  description = "ICMPv6 address unreachable response, type 1, code 3"
}
