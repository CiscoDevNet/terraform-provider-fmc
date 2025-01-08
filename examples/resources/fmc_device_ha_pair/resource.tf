resource "fmc_device_ha_pair" "example" {
  name                             = "FTD_HA"
  primary_device_id                = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  secondary_device_id              = "96d24097-41c4-4332-a4d0-a8c07ac08482"
  is_encryption_enabled            = false
  use_same_link_for_failovers      = false
  shared_key                       = "cisco123"
  enc_key_generation_scheme        = "CUSTOM"
  lan_failover_standby_ip          = "1.1.1.2"
  lan_failover_active_ip           = "1.1.1.1"
  lan_failover_name                = "LAN-INTERFACE"
  lan_failover_subnet_mask         = "255.255.255.0"
  lan_failover_ipv6_addr           = false
  lan_failover_interface_name      = "GigabitEthernet0/0"
  lan_failover_interface_id        = "757kdgh5-41c4-4558-a4d0-a8c07ac08470"
  lan_failover_interface_type      = "PhysicalInterface"
  stateful_failover_standby_ip     = "10.10.10.2"
  stateful_failover_active_ip      = "10.10.10.1"
  stateful_failover_name           = "Stateful-INTERFACE"
  stateful_failover_subnet_mask    = "255.255.255.0"
  stateful_failover_ipv6_addr      = false
  stateful_failover_interface_name = "GigabitEthernet0/0"
  stateful_failover_interface_id   = "76d24097-hj7r-7786-a4d0-a8c07ac08470"
  stateful_failover_interface_type = "PhysicalInterface"
  action                           = "SWITCH"
  force_break                      = false
}
