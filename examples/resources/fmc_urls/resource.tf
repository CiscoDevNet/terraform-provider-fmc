resource "fmc_urls" "example" {
  items = {
    url_1 = {
      url         = "https://www.example.com/app"
      description = "My URL"
    }
  }
}
