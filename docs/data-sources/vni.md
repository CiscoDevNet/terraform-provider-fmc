---
page_title: "fmc_device_vni Data Source - terraform-provider-fmc"
subcategory: ""
description: |-
  Data source for VNI in FMC
  An example is shown below:
  hcl
  data "fmc_device_vni" "test_fmc_device_vni" {
     device_id = "<<>>"
     name = "<<>>"
  }
---

# fmc_device_vni (Data Source)

Data source for VNI in FMC

An example is shown below: 
```hcl
data "fmc_device_vni" "test_fmc_device_vni" {
	device_id = "<<>>"
    name = "<<>>"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `device_id` (String) The Device ID of the VNI
- `name` (String) Name of the VNI

### Read-Only
- `description` (String) VNI description
- `segment_id` (Number) The Segment ID of this resource.
- `vnid` (Number) VNID of VNI
- `enable_proxy` (Bool) EnableProxy of the VNI
- `multicast_groupaddress` (String) MulticastGroupAddress of the VNI
- `priority` (Number) Priority of the VNI
- `if_name` (String) Logical name for VNI
- `security_zone_id` (String) Security Zone ID for VNI
