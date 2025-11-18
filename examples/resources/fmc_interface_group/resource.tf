resource "fmc_interface_group" "example" {
  name           = "my_interface_group"
  interface_type = "ROUTED"
  interfaces = [
    {
      id = "0050568A-4E02-1ed3-0000-004294969198"
    }
  ]
}
