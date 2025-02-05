resource "fmc_sgts" "example" {
  items = {
    my_sgts = {
      description = "My SGT object"
      tag         = "11"
    }
  }
}
