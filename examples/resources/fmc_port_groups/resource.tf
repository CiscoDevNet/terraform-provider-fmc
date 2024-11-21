resource "fmc_port_groups" "example" {
  items = {
    port_group_1 = {
      description = "My port group description"
      ports = [
        {
          id   = "0050568A-232D-0ed3-0000-004294971602"
          type = "ProtocolPortObject"
        }
      ]
    }
  }
}
