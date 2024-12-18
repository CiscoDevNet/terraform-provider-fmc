resource "fmc_sgts" "example" {
  items = {
    sgt_1 = {
      description = "My SGT object"
      tag         = "11"
    }
  }
}
