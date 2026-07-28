resource "fmc_device_dhcp_server" "example" {
  device_id             = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  ping_timeout          = 50
  lease_length          = 3600
  domain_name           = "example.com"
  primary_dns_server_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  servers = [
    {
      interface_id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      interface_name = "GigabitEthernet0/1"
      interface_type = "PhysicalInterface"
      address_pool   = "10.1.1.10-10.1.1.20"
      enabled        = true
    }
  ]
  dhcp_options = [
    {
      code                = 3
      type                = "IP"
      first_ip_address_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
    }
  ]
}
