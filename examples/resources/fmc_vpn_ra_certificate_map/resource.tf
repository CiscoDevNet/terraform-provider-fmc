resource "fmc_vpn_ra_certificate_map" "example" {
  vpn_ra_id                                      = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  use_alias_url                                  = false
  use_certificate_to_connection_profile_mappings = false
  certificate_to_connection_profile_mappings = [
    {
      certificate_map_id    = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      connection_profile_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
    }
  ]
}
