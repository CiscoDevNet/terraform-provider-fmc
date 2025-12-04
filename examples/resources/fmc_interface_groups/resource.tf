resource "fmc_interface_groups" "example" {
  items = {
    my_interface_groups = {
      interface_type = "ROUTED"
      interfaces = [
        {
          id = "0050568A-4E02-1ed3-0000-004294969198"
        }
      ]
    }
  }
}
