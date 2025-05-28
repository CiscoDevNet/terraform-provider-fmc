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
      ipv6_access_list_addresses = [
        {
          id   = "0050568A-4E02-1ed3-0000-004294969199"
          type = "0050568A-4E02-1ed3-0000-004294969199"
        }
      ]
      ipv6_access_list_next_hops = [
        {
          id   = "0050568A-4E02-1ed3-0000-004294969199"
          type = "0050568A-4E02-1ed3-0000-004294969199"
        }
      ]
      ipv6_access_list_route_sources = [
        {
          id   = "0050568A-4E02-1ed3-0000-004294969199"
          type = "0050568A-4E02-1ed3-0000-004294969199"
        }
      ]
      metric_route_values          = []
      tag_values                   = []
      route_type_external1         = true
      route_type_external2         = true
      route_type_internal          = true
      route_type_local             = true
      route_type_n_s_s_a_external1 = true
      route_type_n_s_s_a_external2 = true
    }
  ]
}
