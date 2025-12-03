resource "fmc_device_ospf" "example" {
  device_id                          = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  process_id                         = 1
  router_id                          = "10.10.10.1"
  rfc_1583_compatible                = false
  log_adjacency_changes              = "DEFAULT"
  ignore_lsa_mospf                   = false
  administrative_distance_inter_area = 110
  administrative_distance_intra_area = 110
  administrative_distance_external   = 110
  timer_lsa_group                    = 240
  default_route_always_advertise     = false
  default_route_metric               = 1
  default_route_metric_type          = "TYPE_2"
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
      inter_area_filters = [
        {
          prefix_list_id   = "123e4567-e89b-12d3-a456-426614174000"
          prefix_list_name = "Office Prefix List"
          filter_direction = "IN"
        }
      ]
    }
  ]
}
