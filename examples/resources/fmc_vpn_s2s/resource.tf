resource "fmc_vpn_s2s" "example" {
  name             = "my_ftd_s2s_vpn"
  route_based      = true
  network_topology = "POINT_TO_POINT"
  ikev1            = false
  ikev2            = true
}
