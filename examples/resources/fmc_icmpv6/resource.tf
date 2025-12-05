resource "fmc_icmpv6" "example" {
  name        = "my_icmpv6_object"
  description = "ICMPv6 address unreachable response, type 1, code 3"
  icmp_type   = "1"
  code        = 3
}
