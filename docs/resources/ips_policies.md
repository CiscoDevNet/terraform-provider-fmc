---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_ips_policies Resource - terraform-provider-fmc"
subcategory: ""
description: |-
  Resource for Intrusion Policies in FMC
  Example
  An example is shown below:
  hcl
  resource "fmc_ips_policies" "ips_policy" {
      name = "test"
      inspection_mode = "DETECTION"
      basepolicy_id = "<basepolicy-id>"
  }
---

# fmc_ips_policies (Resource)

Resource for Intrusion Policies in FMC

## Example
An example is shown below: 
```hcl
resource "fmc_ips_policies" "ips_policy" {
    name = "test"
    inspection_mode = "DETECTION"
    basepolicy_id = "<basepolicy-id>"
}
```



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `basepolicy_id` (String) Id of the base policy
- `inspection_mode` (String) The inspection mode of this policy
- `name` (String) The name of this resource

### Optional

- `basepolicy_name` (String) Name of the base policy

### Read-Only

- `id` (String) The ID of this resource.
- `type` (String) The type of this resource


