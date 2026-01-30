resource "fmc_vpn_ra_address_assignment_policy" "example" {
  vpn_ra_id                                 = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  ipv4_use_authorization_server             = true
  ipv4_use_dhcp                             = true
  ipv4_use_internal_address_pool            = true
  ipv4_internal_address_pool_reuse_interval = 0
  ipv6_use_authorization_server             = true
  ipv6_use_internal_address_pool            = true
}
