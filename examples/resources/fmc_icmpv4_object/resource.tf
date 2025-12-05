resource "fmc_icmpv4_object" "example" {
  icmp_type   = "3"
  code        = 0
  name        = "my_icmpv4_object"
  description = "ICMPv4 network unreachable response, type 3, code 0"
}
