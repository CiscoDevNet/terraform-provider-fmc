---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_vpn_s2s_advanced_settings Data Source - terraform-provider-fmc"
subcategory: "VPN"
description: |-
  This data source reads the VPN S2S Advanced Settings.
---

# fmc_vpn_s2s_advanced_settings (Data Source)

This data source reads the VPN S2S Advanced Settings.

## Example Usage

```terraform
data "fmc_vpn_s2s_advanced_settings" "example" {
  id         = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  vpn_s2s_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Id of the object
- `vpn_s2s_id` (String) Id of the parent VPN S2S Topology.

### Optional

- `domain` (String) Name of the FMC domain

### Read-Only

- `bypass_access_control_policy_for_decrypted_traffic` (Boolean) Enable bypass access control traffic for decrypted traffic (sysopt permit-vpn).
- `cert_use_certificate_map_configured_in_endpoint_to_determine_tunnel` (Boolean) Use certificate map configured in endpoint to determine tunnel.
- `cert_use_ike_identity_to_determine_tunnel` (Boolean) Use IKE identity OU to determine tunnel.
- `cert_use_ou_to_determine_tunnel` (Boolean) Use certificate OU to determine tunnel.
- `cert_use_peer_ip_address_to_determine_tunnel` (Boolean) Use peer IP address to determine tunnel.
- `ike_aggressive_mode` (Boolean) Enable IKEv1 aggressive mode.
- `ike_identity_sent_to_peers` (String) identity that the peers will use to identify themselves during IKE negotiations.
- `ike_keepalive` (String) Enable IKE keepalives.
- `ike_keepalive_retry_interval` (Number) IKE keepalive retry interval in seconds.
- `ike_keepalive_threshold` (Number) IKE keepalive threshold in seconds. This interval is the number of seconds allowing a peer to idle before beginning keepalive monitoring.
- `ike_notification_on_tunnel_disconnect` (Boolean) Enable notification on tunnel disconnect.
- `ike_peer_identity_validation` (String) Peer identity validation.
- `ikev2_cookie_challenge` (String) Send cookie challenges to peer devices in response to SA initiate packets.
- `ikev2_maximum_number_of_sas_allowed` (Number) Limits the number of allowed IKEv2 connections.
- `ikev2_number_of_sas_allowed_in_negotiation` (Number) Limits the maximum number of SAs that can be in negotiation at any time.
- `ikev2_threshold_to_challenge_incoming_cookies` (Number) The percentage of the total allowed SAs that are in-negotiation.
- `ipsec_fragmentation_before_encryption` (Boolean) Enable fragmentation before encryption.
- `ipsec_path_maximum_transmission_unit_aging_reset_interval` (Number) Enter the number of minutes at which the Path Maximum Transission Unit (PMTU) value of an SA is reset to its original value.
- `nat_keepalive_message_traversal` (Boolean) Enable NAT keepalive message traversal.
- `nat_keepalive_message_traversal_interval` (Number) NAT keepalive message traversal interval in seconds.
- `type` (String) Type of the object; this value is always 'AdvancedSetting'.
- `vpn_idle_timeout` (Boolean) Enable VPN idle timeout monitoring.
- `vpn_idle_timeout_value` (Number) VPN idle timeout in minutes.
