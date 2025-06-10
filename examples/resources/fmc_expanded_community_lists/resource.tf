resource "fmc_expanded_community_lists" "example" {
  items = {
    my_expanded_community_lists = {
      entries = [
        {
          action     = "PERMIT"
          expression = "^123$"
        }
      ]
    }
  }
}
