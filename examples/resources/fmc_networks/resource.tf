resource "fmc_networks" "example" {
  items = {
    my_networks_object = {
      description = "My Network 1"
      prefix      = "10.1.1.0/24"
    }
  }
}
