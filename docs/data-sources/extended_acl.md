---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_extended_acl Data Source - terraform-provider-fmc"
subcategory: ""
description: |-
  Data source for Extended accesslist in FMC
  An example is shown below:
  hcl
  data "fmc_extended_acl" "extended-acl" {
      name = "TEST-ACL"
  }
---

# fmc_extended_acl (Data Source)

Data source for Extended accesslist in FMC

An example is shown below: 
```hcl
data "fmc_extended_acl" "extended-acl" {
	name = "TEST-ACL"
}
```



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the Extended Acl resource

### Read-Only

- `id` (String) The ID of this resource
- `type` (String) Type of this resource


