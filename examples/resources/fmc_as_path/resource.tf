resource "fmc_as_path" "example" {
  name        = 100
  overridable = false
  entries = [
    {
      action             = "PERMIT"
      regular_expression = "^(100|200)$"
    }
  ]
}
