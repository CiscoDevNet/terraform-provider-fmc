---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_ftd_nat_policies Data Source - terraform-provider-fmc"
subcategory: ""
description: |-
  Data source for Nat Policies in FMC
  An example is shown below:
  hcl
  data "fmc_ftd_nat_policies" "natp" {
      name = "FTD NATP"
  }
---

# fmc_ftd_nat_policies (Data Source)

Data source for Nat Policies in FMC

An example is shown below: 
```hcl
data "fmc_ftd_nat_policies" "natp" {
	name = "FTD NATP"
}
```



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the FTD natPolicy

### Read-Only

- `description` (String) Description of this resource
- `id` (String) The ID of this resource
- `type` (String) Type of this resource


