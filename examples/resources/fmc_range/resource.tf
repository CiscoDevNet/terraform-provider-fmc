resource "fmc_range" "example" {
  name        = "my_range"
  ip_range    = "10.0.0.1-10.0.0.9"
  description = "My range"
}
