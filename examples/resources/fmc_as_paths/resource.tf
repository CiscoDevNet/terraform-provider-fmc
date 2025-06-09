resource "fmc_as_paths" "example" {
  items = {
    240 = {
      overridable = false
      entries = [
        {
          action             = "PERMIT"
          regular_expression = "^(100|200)$"
        }
      ]
    }
  }
}
