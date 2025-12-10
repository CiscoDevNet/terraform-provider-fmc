resource "fmc_ipv6_prefix_lists" "example" {
  items = {
    my_ipv6_prefix_lists = {
      entries = [
        {
          action            = "PERMIT"
          prefix            = "2001:1::/64"
          min_prefix_length = 65
          max_prefix_length = 120
        }
      ]
    }
  }
}
