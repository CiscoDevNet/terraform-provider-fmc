resource "fmc_extended_community_lists" "example" {
  items = {
    my_extended_community_lists = {
      sub_type = "Standard"
      entries = [
        {
          action       = "PERMIT"
          route_target = "64512:1010"
        }
      ]
    }
  }
}
