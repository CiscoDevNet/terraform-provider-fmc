resource "fmc_range" "example" {
  name        = "range1"
  ip_range    = "10.0.0.1-10.0.0.9"
  description = "My range"
}
