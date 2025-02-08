resource "fmc_hosts" "example" {
  items = {
    my_hosts = {
      description = "My Host 1"
      ip          = "10.1.1.1"
    }
  }
}
