resource "fmc_chassis_logical_device" "example" {
  chassis_id            = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  name                  = "my-logical-device"
  ftd_version           = "7.6.0.113"
  ipv4_address          = "10.10.10.10"
  ipv4_netmask          = "255.255.255.0"
  ipv4_gateway          = "10.10.10.1"
  ipv6_address          = "2001:db8::10"
  ipv6_prefix           = 64
  ipv6_gateway          = "2001:db8::1"
  search_domain         = "cisco.com"
  fqdn                  = "my_logical_device.cisco.com"
  firewall_mode         = "ROUTED"
  dns_servers           = "10.123.10.12,10.123.10.14"
  device_password       = "my_password"
  admin_state           = "ENABLED"
  permit_expert_mode    = "yes"
  resource_profile_id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  resource_profile_name = "my_resource_profile"
  assigned_interfaces = [
    {
      id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
    }
  ]
  device_group_id      = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  access_policy_id     = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  platform_settings_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  license_capabilities = ["MALWARE"]
}
