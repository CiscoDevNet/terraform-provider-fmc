## 2.0.0-rc5 (Unreleased)

- (Enhancement) Add support for `fmc_countries` and `fmc_continents` data sources
- (Enhancement) Add support for `fmc_geolocation` resource and data source
- (Enhancement) Add support for `fmc_service_access` resource and data source
- (Enhancement) Add support for `fmc_secure_client_*` resources and data sources

## 2.0.0-rc4

- (Fix) Corrected URL encoding for multiple resources
- (Fix) `fmc_device`: add missing FTDv100 performance tier
- (Fix) `fmc_vpn_s2s_endpoints`: import
- (Enhancement) Add support for Security Cloud Control (SCC) Firewall Management Base URI
- (Enhancement) Add support for `fmc_security_intelligence_*` DNS/URL/Network feeds and lists data sources and available resources
- (Enhancement) Add support for `fmc_sla_monitor` and `fmc_sla_monitors` resources and data sources
- (Enhancement) Add support for SLA Monitors under IPv4 static routes

## 2.0.0-rc3

- (Fix) Fixes to `fmc_network_groups`, including: support for more than 1000 items per resource and bulk delete for FMC 7.4 and newer

## 2.0.0-rc2

- (BREAKING CHANGE) Multiple fields renamed in `fmc_device_bgp` resource
- (Fix) Multiple fixes to `fmc_device_bgp` resource
- (Fix) `fmc_device_bridge_group_interface`: `logical_name` is no longer `required` field
- (Fix) `fmc_network_groups`: `id` attribute is set to random value

## 2.0.0-rc1

- (Change) Resource `fmc_device_ha_pair_physical_interface_mac_address` is deprecated and replaced with `fmc_device_ha_pair_failover_interface_mac_address`
- (Enhancement) Add support for `fmc_route_map` resources and data sources
- (Enhancement) Add support for `fmc_policy_list` resources and data sources
- (Enhancement) Add support for `fmc_extended_community_list` and `fmc_extended_community_lists` resources and data sources
- (Enhancement) Add support for `fmc_expanded_community_list` and `fmc_expanded_community_lists` resources and data sources
- (Enhancement) Add support for `fmc_standard_community_list` and `fmc_standard_community_lists` resources and data sources
- (Enhancement) Add support for `fmc_ipv6_preflix_list` and `fmc_ipv6_preflix_lists` resources and data sources
- (Enhancement) Add support for `fmc_ipv4_preflix_list` and `fmc_ipv4_preflix_lists` resources and data sources
- (Enhancement) Add support for `fmc_as_path` and `fmc_as_paths` resources and data sources

## 2.0.0-rc0

- (Enhancement) Add support for Resource Profile (`fmc_resource_profiles`) resource and data source
- (Enhancement) Add support for multi-instance (`fmc_chassis_*`) resources and data sources. Tune `fmc_device_subinterface` and `fmc_device_etherchannel_interface` to support multi-instance logical devices.
- (Enhancement) Add support for Bridge Group Interface (BVI)  (`fmc_device_bridge_group_interface`)
- (Enhancement) Add cdFMC (Cloud-Delivered FMC) and FMC 7.7 support
- (Fix) Add `ipv4_address_family_id` attribute to `fmc_device_bgp`

## 2.0.0-beta5

- (Fix) Add `type` attribute to `Device VNI Interface` and `Device VTEP Policy`

## 2.0.0-beta4

- (Enhancement) Add support for ipv4 and ipv6 address pools under subinterface and etherchannel
- (Enhancement) Add support for `fmc_icmpv6_objects`
- (Enhancement) Add support for `fmc_device_loopback_interface`
- (Enhancement) Add support for `fmc_device_vti_interface`
- (Enhancement) Add support for IKEv1 & IKEv2 IPSec Proposals & Policies
- (Enhancement) Add support for `fmc_certificate_map` and `fmc_certificate_maps`
- (Enhancement) Honor proxy settings (`HTTP_PROXY`, `HTTPS_PROXY`, `NO_PROXY` environment variables)
- (Enhancement) Add support for Site-to-Site VPNs (`fmc_vpn_s2s`, `fmc_vpn_s2s_ipsec_settings`, `fmc_vpn_s2s_ike_settings`, `fmc_vpn_s2s_advanced_settings`, `fmc_vpn_s2s_endpoints`)

## 2.0.0-beta3

- (Fix) Change value of `interface_type` within `fmc_security_zones` item should replace just this object, not entire bulk resource
- (Fix) Improved HA Pair and Clustering implementations
- (Enhancement) Add `fmc_device_ha_pair_physical_interface_mac_address` resource and data source
- (Enhancement) Add support for `fmc_ipv4_address_pool` and `fmc_ipv4_address_pools` resource and data source
- (Enhancement) Add support for `fmc_ipv6_address_pool` and `fmc_ipv6_address_pools` resource and data source
- (Enhancement) Add support for `fmc_device_cluster_health_monitor` resource and data source
- (Enhancement) Add support for `fmc_domains` data source
- (Enhancement) Add support for `fmc_endpoint_device_types` data source. It can now be used in Access Control Policy
- (Enhancement) Add support for `fmc_ise_sgts` data source
- (Enhancement) Add support for `destination_sgt_objects` in Access Control Policy Rules

## 2.0.0-beta2

- (Fix) Update minimum FMC version for `fmc_file_type` and `fmc_file_category` data sources
- (Fix) Align fields in ipv4/ipv6/vrf_ipv4/vrf_ipv6 static_route resources
- (Enhancement) Add `type` field to multiple resources
- (Enhancement) Add support for multiple `fmc_application_*` data sources and `fmc_application_filter` & `fmc_application_filters` resources
- (Enhancement) Add support for Applications in Access Rules

## 2.0.0-beta1

- Initial release
