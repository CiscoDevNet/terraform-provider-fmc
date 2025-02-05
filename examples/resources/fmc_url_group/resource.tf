resource "fmc_url_group" "example" {
  name        = "my_url_group"
  description = "My URL group"
  urls = [
    {
      id = "0050568A-FAC7-0ed3-0000-004294987896"
    }
  ]
  literals = [
    {
      url = "https://www.example.com/app"
    }
  ]
}
