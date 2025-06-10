resource "fmc_ipv4_prefix_list" "example" {
  name = "my_ipv4_prefix_list"
  entries = [
    {
      action            = "PERMIT"
      ip_address        = "10.10.10.0/24"
      min_prefix_length = 25
      max_prefix_length = 30
    }
  ]
}
