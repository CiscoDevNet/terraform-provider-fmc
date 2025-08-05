resource "fmc_vpn_ra_load_balancing" "example" {
  vpn_ra_id                               = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  enabled                                 = true
  ipv4_group_address                      = "192.168.1.1"
  ipv6_group_address                      = "2001:db8::1"
  interface_id                            = "76d24097-41c4-4558-a4d0"
  udp_port_number                         = 9023
  ipsec                                   = true
  ipsec_encryption_key                    = "my-secret-key"
  send_fqdn_to_peer_devices_instead_of_ip = false
  ikev2_redirect_phase                    = "DURING_SA_AUTHENTICATION"
}
