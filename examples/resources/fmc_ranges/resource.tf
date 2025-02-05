resource "fmc_ranges" "example" {
  items = {
    my_ranges = {
      description = "My Range 1"
      ip_range    = "10.0.0.1-10.0.0.9"
    }
  }
}
