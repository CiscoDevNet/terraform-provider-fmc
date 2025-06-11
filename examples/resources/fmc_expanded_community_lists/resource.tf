resource "fmc_expanded_community_lists" "example" {
  items = {
    my_expanded_community_lists = {
      entries = [
        {
          action             = "PERMIT"
          regular_expression = "^123$"
        }
      ]
    }
  }
}
