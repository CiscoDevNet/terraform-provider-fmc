---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_bfd_template Data Source - terraform-provider-fmc"
subcategory: "Objects"
description: |-
  This data source reads the BFD Template.
  The following restrictions apply:
  Minimum FMC version: 7.4
---

# fmc_bfd_template (Data Source)

This data source reads the BFD Template.

The following restrictions apply:
  - Minimum FMC version: `7.4`

## Example Usage

```terraform
data "fmc_bfd_template" "example" {
  id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `domain` (String) Name of the FMC domain
- `id` (String) Id of the object
- `name` (String) Name of the BFD Template object.

### Read-Only

- `authentication_key_id` (Number) Authentication Key ID
- `authentication_password` (String) Password for BFD Authentication (1-24 characters)
- `authentication_password_encryption` (String) Determines if `authentication_password` is encrypted
- `authentication_type` (String) Authentication type.
- `echo` (String) BFD echo status.
- `hop_type` (String) Hop type.
- `interval_time` (String) Interval unit of measurement of time.
- `min_receive` (Number) BFD Minimum Receive unit value in ranges: 50-999 miliseconds, 50000-999000 microseconds
- `min_transmit` (Number) BFD Minimum Transmit unit value.
- `tx_rx_multiplier` (Number) BFD Multipler value.
- `type` (String) Type of the object; this value is always 'BFDTemplate'.
