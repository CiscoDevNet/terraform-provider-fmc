resource "fmc_hosts" "example" {
  items = {
    hosts_1 = {
      description = "My Host 1"
      ip          = "10.1.1.1"
    }
  }
}
