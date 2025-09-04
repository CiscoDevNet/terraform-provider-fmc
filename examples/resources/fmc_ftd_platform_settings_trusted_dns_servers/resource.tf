resource "fmc_ftd_platform_settings_trusted_dns_servers" "example" {
  ftd_platform_settings_id     = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  trust_any_dns_server         = false
  trust_dhcp_pool              = true
  trust_dhcp_relay             = true
  trust_dhcp_client            = true
  trust_dns_server_group       = true
  trusted_dns_servers_literals = ["10.20.30.40"]
  trusted_dns_servers_objects = [
    {
      id = "12345678-1234-1234-1234-123456"
    }
  ]
}
