resource "fmc_port_groups" "example" {
  items = {
    my_port_groups = {
      description = "My port group description"
      objects = [
        {
          id   = "0050568A-232D-0ed3-0000-004294971602"
          type = "ProtocolPortObject"
        }
      ]
    }
  }
}
