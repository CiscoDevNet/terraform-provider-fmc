---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_ise_sgt_objects Data Source - terraform-provider-fmc"
subcategory: ""
description: |-
  Data source for ise security group tags object from pxgrid in FMC
  An example is shown below:
  hcl
  data "fmc_ise_sgt_objects" "workstation" {
      name = "workstation"
  }
---

# fmc_ise_sgt_objects (Data Source)

Data source for ise security group tags object from pxgrid in FMC

An example is shown below: 
```hcl
data "fmc_ise_sgt_objects" "workstation" {
	name = "workstation"
}
```



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of this resource

### Read-Only

- `id` (String) The ID of this resource
- `type` (String) The type of this resource


