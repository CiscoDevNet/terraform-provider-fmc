resource "fmc_vlan_tags" "example" {
  items = {
    my_vlan_tags = {
      description = "My TAG id"
      overridable = true
      start_tag   = "12"
      end_tag     = "15"
    }
  }
}
