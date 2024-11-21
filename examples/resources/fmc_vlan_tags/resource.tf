resource "fmc_vlan_tags" "example" {
  items = {
    vlan_tag_1 = {
      description = "My TAG id"
      overridable = true
      start_tag   = "12"
      end_tag     = "15"
    }
  }
}
