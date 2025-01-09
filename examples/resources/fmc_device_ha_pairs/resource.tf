resource "fmc_device_ha_pairs" "example" {
  name                             = "FTD_HA"
  primary_device_id                = "<FTD1_ID>"
  secondary_device_id              = "<FTD2_ID>"
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
  lan_failover_interface_id        = "<interface uuid>"
  lan_failover_interface_type      = "PhysicalInterface"
  stateful_failover_standby_ip     = "1.1.1.2"
  stateful_failover_active_ip      = "1.1.1.1"
  stateful_failover_name           = "Stateful-INTERFACE"
  stateful_failover_subnet_mask    = "255.255.255.0"
  stateful_failover_ipv6_addr      = false
  stateful_failover_interface_name = "GigabitEthernet0/0"
  stateful_failover_interface_id   = "<interface uuid>"
  stateful_failover_interface_type = "PhysicalInterface"
  action                           = "SWITCH"
  force_break                      = false
}
