resource "fmc_extended_community_list" "example" {
  name     = "my_extended_community_list"
  sub_type = ""
  entries = [
    {
      action       = "PERMIT"
      route_target = "123 456 789"
      expression   = "^(123|456|789)$"
    }
  ]
}
