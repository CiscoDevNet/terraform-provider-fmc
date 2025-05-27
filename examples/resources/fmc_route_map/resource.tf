resource "fmc_route_map" "example" {
  name        = "my_route_map"
  overridable = true
  entries = [
    {
      sequence_number = 10
      action          = ""
      security_zones = [
        {
          id = "0050568A-4E02-1ed3-0000-004294969198"
        }
      ]
      ipv4_access_list_addresses = [
        {
          id   = "0050568A-4E02-1ed3-0000-004294969199"
          type = "0050568A-4E02-1ed3-0000-004294969199"
        }
      ]
      ipv4_access_list_next_hops = [
        {
          id   = "0050568A-4E02-1ed3-0000-004294969199"
          type = "0050568A-4E02-1ed3-0000-004294969199"
        }
      ]
      ipv4_access_list_route_sources = [
        {
          id   = "0050568A-4E02-1ed3-0000-004294969199"
          type = "0050568A-4E02-1ed3-0000-004294969199"
        }
      ]
    }
  ]
}
