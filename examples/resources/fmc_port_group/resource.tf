resource "fmc_port_group" "example" {
  name        = "portgroup_obj1"
  description = "My port group description"
  objects = [
    {
      id   = "0050568A-232D-0ed3-0000-004294971602"
      type = "ProtocolPortObject"
    }
  ]
}
