resource "fmc_port_group" "example" {
  name        = "my_port_group"
  description = "My port group description"
  objects = [
    {
      id   = "0050568A-232D-0ed3-0000-004294971602"
      type = "ProtocolPortObject"
    }
  ]
}
