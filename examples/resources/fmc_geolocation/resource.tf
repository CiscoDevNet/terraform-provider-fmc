resource "fmc_geolocation" "example" {
  name = "my_geolocation"
  continents = [
    {
      id = 6
    }
  ]
  countries = [
    {
      id = 616
    }
  ]
}
