resource "fmc_route_map" "example" {
  name        = "my_route_map"
  overridable = true
  entries = [
    {
      action = ""
      match_security_zones = [
        {
          id = "0050568A-4E02-1ed3-0000-004294969198"
        }
      ]
      match_interface_names = ["GigabitEthernet0/1"]
      match_ipv4_address_access_lists = [
        {
          id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
          type = "StandardAccessList"
        }
      ]
      match_ipv4_next_hop_access_lists = [
        {
          id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
          type = "StandardAccessList"
        }
      ]
      match_ipv4_route_source_access_lists = [
        {
          id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
          type = "StandardAccessList"
        }
      ]
      match_ipv6_address_extended_access_list_id      = ""
      match_ipv6_next_hop_extended_access_list_id     = ""
      match_ipv6_route_source_extended_access_list_id = ""
      match_bgp_as_path_lists = [
        {
          id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
        }
      ]
      match_bgp_community_lists = [
        {
          id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
        }
      ]
      match_bgp_policy_lists = [
        {
          id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
        }
      ]
      match_metric_route_values                              = []
      match_tag_values                                       = []
      match_route_type_external_1                            = true
      match_route_type_external_2                            = true
      match_route_type_internal                              = true
      match_route_type_local                                 = true
      match_route_type_nssa_external_1                       = true
      match_route_type_nssa_external_2                       = true
      set_metric_bandwidth                                   = 1000000
      set_metric_type                                        = "INTERNAL"
      set_bgp_as_path_prepend                                = []
      set_bgp_as_path_prepend_last_as                        = 2
      set_bgp_as_path_convert_route_tag_into_as_path         = true
      set_bgp_community_none                                 = true
      set_bgp_community_specific_community                   = 100
      set_bgp_community_add_to_existing_communities          = true
      set_bgp_community_internet                             = true
      set_bgp_community_no_advertise                         = true
      set_bgp_community_no_export                            = true
      set_bgp_community_route_target                         = ""
      set_bgp_community_add_to_existing_extended_communities = true
      set_bgp_automatic_tag                                  = true
      set_bgp_local_preference                               = 100
      set_bgp_weight                                         = 100
      set_bgp_origin                                         = "LOCAL_IGP"
      set_bgp_ipv4_next_hop                                  = ""
      set_bgp_ipv4_next_hop_specific_ip                      = [""]
      set_bgp_ipv4_prefix_list_id                            = ""
      set_bgp_ipv6_next_hop                                  = ""
      set_bgp_ipv6_next_hop_specific_ip                      = [""]
      set_bgp_ipv6_prefix_list_id                            = ""
    }
  ]
}
