resource "fmc_vpn_s2s" "example" {
  name          = "my_ftd_s2s_vpn"
  route_based   = true
  topology_type = "POINT_TO_POINT"
  ikev1_enable  = false
  ikev2_enable  = true
}
