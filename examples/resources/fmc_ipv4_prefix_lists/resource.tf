resource "fmc_ipv4_prefix_lists" "example" {
  items = {
    my_ipv4_prefix_lists = {
      entries = [
        {
          action            = "PERMIT"
          prefix            = "10.10.10.0/24"
          min_prefix_length = 25
          max_prefix_length = 30
        }
      ]
    }
  }
}
