resource "fmc_extended_community_list" "example" {
  name     = "my_extended_community_list"
  sub_type = "Standard"
  entries = [
    {
      action       = "PERMIT"
      route_target = "64512:1010"
    }
  ]
}
