---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_device_cluster Data Source - terraform-provider-fmc"
subcategory: ""
description: |-
  Data source for FTD Device Cluster in FMC
  An example is shown below:
  hcl
  data "fmc_device_cluster" "cluster" {
      name = "ftd-cluster1"
  }
---

# fmc_device_cluster (Data Source)

Data source for FTD Device Cluster in FMC

An example is shown below: 
```hcl
data "fmc_device_cluster" "cluster" {
	name = "ftd-cluster1"
}
```



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the cluster

### Read-Only

- `id` (String) The ID of this resource.
- `type` (String) Type of this resource


