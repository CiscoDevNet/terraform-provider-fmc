resource "fmc_urls" "example" {
  items = {
    my_urls = {
      description = "My URL"
      url         = "https://www.example.com/app"
    }
  }
}
