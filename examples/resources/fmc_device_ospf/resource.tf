resource "fmc_device_ospf" "example" {
  device_id                          = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  process_id                         = 1
  router_id                          = "AUTOMATIC"
  rfc_1583_compatible                = false
  log_adjacency_changes              = "DEFAULT"
  ignore_lsa_mospf                   = false
  administrative_distance_inter_area = 110
  administrative_distance_intra_area = 110
  administrative_distance_external   = 110
  timer_lsa_group                    = 240
  default_route_always_advertise     = false
  default_route_metric               = 1
  default_route_metric_type          = ""
  default_route_route_map_id         = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  areas = [
    {
      id   = "1"
      type = "normal"
      networks = [
        {
          id   = "123e4567-e89b-12d3-a456-426614174000"
          name = "Office Network"
        }
      ]
      authentication = "MESSAGE_DIGEST"
      default_cost   = 1
      ranges = [
        {
          network_object_id = "123e4567-e89b-12d3-a456-426614174000"
          advertise         = true
        }
      ]
      virtual_links = [
        {
          peer_router_host_id     = "123e4567-e89b-12d3-a456-426614174000"
          hello_interval          = 10
          transmit_delay          = 1
          retransmit_interval     = 5
          dead_interval           = 40
          authentication_password = "ospfAuthKey"
        }
      ]
      inter_area_filters = [
        {
          prefix_list_id   = ""
          prefix_list_name = ""
          filter_direction = ""
        }
      ]
    }
  ]
  redistributions = [
    {
      route_type                 = "RedistributeBGP"
      as_number                  = 65001
      ospf_match_external_1      = true
      ospf_match_external_2      = true
      ospf_match_internal        = true
      ospf_match_nssa_external_1 = true
      ospf_match_nssa_external_2 = true
      process_id                 = 2
      subnets                    = true
      metric                     = 1
      metric_type                = "TYPE_2"
      tag                        = 100
      route_map_id               = "76d24097-41c4-4558-a4d0-a8c07ac08470"
    }
  ]
  filter_rules = [
    {
      access_list_id    = "123e4567-e89b-12d3-a456-426614174000"
      traffic_direction = "incomingroutefilter"
      routing_process   = "CONNECTED"
    }
  ]
  summary_addresses = [
    {
      networks = [
        {
          id = "123e4567-e89b-12d3-a456-426614174000"
        }
      ]
      tag       = 100
      advertise = true
    }
  ]
}
