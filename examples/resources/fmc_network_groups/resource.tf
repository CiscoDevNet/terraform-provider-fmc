resource "fmc_network_groups" "example" {
  items = {
    my_network_groups = {
      description    = "My Network Group 1"
      network_groups = ["child_group_1"]
      objects = [
        {
          id   = "cf271be2-e80d-4033-8e1d-285f133a5f8c"
          name = "fmc_network_object_1"
        }
      ]
      literals = [
        {
          value = "10.1.1.0/24"
        }
      ]
    }
  }
}
