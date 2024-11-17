resource "fmc_ranges" "example" {
  items = {
    ranges_1 = {
      description = "My Range 1"
      ip_range    = "10.0.0.1-10.0.0.9"
    }
  }
}
