resource "fmc_certificate_map" "example" {
  name = "my_certificate_map"
  rules = [
    {
      field     = "SUBJECT"
      component = "COMMON_NAME"
      operator  = "EQUALS"
      value     = "cisco.com"
    }
  ]
}
