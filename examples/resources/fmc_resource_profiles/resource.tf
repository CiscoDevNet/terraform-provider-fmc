resource "fmc_resource_profiles" "example" {
  items = {
    my_resource_profiles = {
      description    = "My Resource Profile object"
      number_of_cpus = 10
    }
  }
}
