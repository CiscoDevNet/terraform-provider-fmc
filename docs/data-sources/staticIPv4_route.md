---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_staticIPv4_route Data Source - terraform-provider-fmc"
subcategory: ""
description: |-
  Data source for Static IPv4 route in FMC
  An example is shown below:
  hcl
  data "fmc_staticIPv4_route" "route" {
     device_id = "<device ID>"
     network_name = "Random-net"
  }
---

# fmc_staticIPv4_route (Data Source)

Data source for Static IPv4 route in FMC

An example is shown below: 
```hcl
data "fmc_staticIPv4_route" "route" {
   device_id = "<device ID>"
   network_name = "Random-net"
}
```



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `device_id` (String) The ID of this resource
- `network_name` (String) The ID of this resource

### Read-Only

- `id` (String) The ID of this resource.
- `type` (String) Type of this resource


