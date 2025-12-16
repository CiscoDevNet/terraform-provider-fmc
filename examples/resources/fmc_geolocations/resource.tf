resource "fmc_geolocations" "example" {
  items = {
    my_geolocations = {
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
  }
}
