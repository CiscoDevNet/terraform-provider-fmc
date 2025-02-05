resource "fmc_vlan_tag_group" "example" {
  name        = "fmc_vlan_tag_group"
  description = "My VLAN Tag Group"
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
