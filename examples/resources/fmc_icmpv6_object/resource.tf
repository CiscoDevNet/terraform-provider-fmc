resource "fmc_icmpv6_object" "example" {
  icmp_type   = 1
  code        = 3
  name        = "my_icmpv6_object"
  description = "ICMPv6 address unreachable response, type 1, code 3"
}
