resource "fmc_certificate_maps" "example" {
  items = {
    my_certificate_maps = {
      rules = [
        {
          field     = "SUBJECT"
          component = "COMMON_NAME"
          operator  = "EQUALS"
          value     = "cisco.com"
        }
      ]
    }
  }
}
