resource "fmc_icmpv6_objects" "example" {
  items = {
    my_icmpv6_objects = {
      icmp_type   = 1
      code        = 3
      description = "ICMPv6 address unreachable response, type 1, code 3"
    }
  }
}
