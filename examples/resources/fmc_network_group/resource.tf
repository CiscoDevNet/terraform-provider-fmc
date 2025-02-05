resource "fmc_network_group" "example" {
  name        = "my_network_group"
  description = "My Network Group 1"
  objects = [
    {
      id = "12345678-1234-1234-1234-123456789012"
    }
  ]
  literals = [
    {
      value = "10.1.1.0/24"
    }
  ]
}
