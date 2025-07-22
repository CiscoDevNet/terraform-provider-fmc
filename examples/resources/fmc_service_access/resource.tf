resource "fmc_service_access" "example" {
  name           = "my_service_access"
  default_action = "DENY"
  rules = [
    {
      action = "ALLOW"
      geolocation_sources = [
        {
          id   = "616"
          type = "Country"
        }
      ]
    }
  ]
}
