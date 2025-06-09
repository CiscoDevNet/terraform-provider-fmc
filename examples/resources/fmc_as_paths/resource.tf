resource "fmc_as_paths" "example" {
  items = {
    my_as_paths = {
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
