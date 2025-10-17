resource "fmc_icmpv4" "example" {
  name        = "my_icmpv4_object"
  description = "ICMPv4 network unreachable response, type 3, code 0"
  icmp_type   = 3
  code        = 0
}
