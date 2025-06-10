resource "fmc_standard_community_lists" "example" {
  items = {
    my_standard_community_lists = {
      entries = [
        {
          action       = "PERMIT"
          communities  = "123 456 789"
          internet     = true
          no_advertise = true
          no_export    = true
        }
      ]
    }
  }
}
