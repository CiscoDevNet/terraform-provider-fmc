---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_ftd_nat_policies Resource - terraform-provider-fmc"
subcategory: ""
description: |-
  Resource for NAT Policies in FMC
  Example
  An example is shown below:
  hcl
  resource "fmc_ftd_nat_policies" "nat_policy" {
      name = "Terraform NAT Policy"
      description = "New NAT policy!"
  }
---

# fmc_ftd_nat_policies (Resource)

Resource for NAT Policies in FMC

## Example
An example is shown below: 
```hcl
resource "fmc_ftd_nat_policies" "nat_policy" {
    name = "Terraform NAT Policy"
    description = "New NAT policy!"
}
```



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of this resource

### Optional

- `description` (String) The description of this resource

### Read-Only

- `id` (String) The ID of this resource.
- `type` (String) The type of this resource


