resource "fmc_icmpv6s" "example" {
  items = {
    my_icmpv6_objects = {
      description = "ICMPv6 address unreachable response, type 1, code 3"
      icmp_type   = 1
      code        = 3
    }
  }
}
