---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_intrusion_policy Resource - terraform-provider-fmc"
subcategory: "Policies"
description: |-
  This resource manages an Intrusion Policy.
---

# fmc_intrusion_policy (Resource)

This resource manages an Intrusion Policy.

## Example Usage

```terraform
resource "fmc_intrusion_policy" "example" {
  name            = "my_intrusion_policy"
  description     = "My IPS Policy"
  base_policy_id  = "0050568A-4E02-1ed3-0000-004294969198"
  inspection_mode = "PREVENTION"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `base_policy_id` (String) Id of the base policy.
- `name` (String) Name of the policy.

### Optional

- `description` (String) Description of the policy.
- `domain` (String) Name of the FMC domain
- `inspection_mode` (String) Inspection mode.
  - Choices: `PREVENTION`, `DETECTION`

### Read-Only

- `id` (String) Id of the object
- `type` (String) Type of the object; this value is always 'intrusionpolicy'.

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import fmc_intrusion_policy.example "<id>"
```
