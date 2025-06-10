resource "fmc_extended_community_lists" "example" {
  items = {
    my_extended_community_lists = {
      sub_type = ""
      entries = [
        {
          action       = "PERMIT"
          route_target = "123 456 789"
          expression   = "^(123|456|789)$"
        }
      ]
    }
  }
}
