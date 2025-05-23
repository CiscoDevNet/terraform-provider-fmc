---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_vpn_s2s_ike_settings Data Source - terraform-provider-fmc"
subcategory: "VPN"
description: |-
  This data source reads the VPN S2S IKE Settings.
---

# fmc_vpn_s2s_ike_settings (Data Source)

This data source reads the VPN S2S IKE Settings.

## Example Usage

```terraform
data "fmc_vpn_s2s_ike_settings" "example" {
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

- `ikev1_authentication_type` (String) Authentication method for IKEv1.
- `ikev1_automatic_pre_shared_key_length` (Number) Length of the automatically generated pre-shared key for IKEv1.
- `ikev1_certificate_id` (String) Id of the certificate for certificate-based authentication for IKEv1.
- `ikev1_manual_pre_shared_key` (String, Sensitive) Manually configured pre-shared key for IKEv1.
- `ikev1_policies` (Attributes Set) Set of policies for IKEv1. (see [below for nested schema](#nestedatt--ikev1_policies))
- `ikev2_authentication_type` (String) Authentication method for IKEv2.
- `ikev2_automatic_pre_shared_key_length` (Number) Length of the automatically generated pre-shared key for IKEv2.
- `ikev2_certificate_id` (String) Id of the certificate for certificate-based authentication for IKEv2.
- `ikev2_enforce_hex_based_pre_shared_key` (Boolean) Enforce use of a hex-based pre-shared key for IKEv2.
- `ikev2_manual_pre_shared_key` (String, Sensitive) Manually configured pre-shared key for IKEv2.
- `ikev2_policies` (Attributes Set) Set of policies for IKEv2 settings. (see [below for nested schema](#nestedatt--ikev2_policies))
- `type` (String) Type of the object; this value is always 'IkeSetting'.

<a id="nestedatt--ikev1_policies"></a>
### Nested Schema for `ikev1_policies`

Read-Only:

- `id` (String) Id of the IKEv1 policy
- `name` (String) Name of the IKEv1 policy


<a id="nestedatt--ikev2_policies"></a>
### Nested Schema for `ikev2_policies`

Read-Only:

- `id` (String) Id of the IKEv2 policy
- `name` (String) Name of the IKEv2 policy
