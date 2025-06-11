resource "fmc_standard_community_list" "example" {
  name = "my_standard_community_list"
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
