resource "fmc_icmpv4_objects" "example" {
  items = {
    my_icmpv4_objects = {
      description = "ICMPv4 network unreachable response, type 3, code 0"
      icmp_type   = "3"
      code        = 0
    }
  }
}
