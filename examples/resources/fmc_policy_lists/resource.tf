resource "fmc_policy_lists" "example" {
  items = {
    my_policy_lists = {
      action = "PERMIT"
      interfaces = [
        {
          id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
        }
      ]
      interface_names = ["GigabitEthernet0/1"]
      address_standard_access_lists = [
        {
          id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
        }
      ]
      next_hop_standard_access_lists = [
        {
          id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
        }
      ]
      route_source_standard_access_lists = [
        {
          id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
        }
      ]
      as_paths = [
        {
          id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
        }
      ]
      community_lists = [
        {
          id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
        }
      ]
      match_community_exactly = true
      metric                  = 100
      tag                     = 100
    }
  }
}
