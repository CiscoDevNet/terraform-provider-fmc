---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_device_vtep_policies Data Source - terraform-provider-fmc"
subcategory: ""
description: |-
  Data source for VTEP policy on FMC
  An example is shown below:
  hcl
  data "fmc_device_vtep_policies" "vtep" {
     device_id = "<device ID>"
  }
---

# fmc_device_vtep_policies (Data Source)

Data source for VTEP policy on FMC

An example is shown below: 
```hcl
data "fmc_device_vtep_policies" "vtep" {
   device_id = "<device ID>"
}
```



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `device_id` (String) The ID of this VTEP Policies

### Read-Only

- `id` (String) The ID of this resource.
- `nveenable` (Boolean) NveEnable of this VTEP Policies
- `type` (String) Type of this VTEP Policies


