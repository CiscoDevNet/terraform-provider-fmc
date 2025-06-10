resource "fmc_expanded_community_list" "example" {
  name = "my_expanded_community_list"
  entries = [
    {
      action     = "PERMIT"
      expression = "^123$"
    }
  ]
}
