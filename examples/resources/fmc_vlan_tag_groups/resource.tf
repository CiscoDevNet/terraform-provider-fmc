resource "fmc_vlan_tag_groups" "example" {
  items = {
    fmc_vlan_tag_groups = {
      description = "My vlan tag group name"
      overridable = true
      vlan_tags = [
        {
          id = "0050568A-4E02-1ed3-0000-004294969198"
        }
      ]
      literals = [
        {
          start_tag = "11"
          end_tag   = "22"
        }
      ]
    }
  }
}
