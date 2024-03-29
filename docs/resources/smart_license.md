---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_smart_license Resource - terraform-provider-fmc"
subcategory: ""
description: |-
  Resource for Smart License in FMC
  Example
  Examples are shown below:
  Enable Evaluation Mode
  hcl
  resource "fmc_smart_license" "license" {
    registration_type = "EVALUATION"
  }
  
  Runs successfully as long as the registration status is REGISTERED
  hcl
  resource "fmc_smart_license" "license" {
    registration_type = "REGISTER"
    token             = "X2M3YmJlY..."
  }
  
  Force to de-register and register with the token provided
  hcl
  resource "fmc_smart_license" "license" {
    registration_type = "REGISTER"
    token             = "X2M3YmJlY..."
    force             = true
  }
  
  De-register on terraform destroy
  hcl
  resource "fmc_smart_license" "license" {
    registration_type = "REGISTER"
    token             = "X2M3YmJlY..."
    retain            = false
  }
---

# fmc_smart_license (Resource)

Resource for Smart License in FMC

## Example
Examples are shown below: 

Enable Evaluation Mode
```hcl
resource "fmc_smart_license" "license" {
  registration_type = "EVALUATION"
}
```

Runs successfully as long as the registration status is REGISTERED
```hcl
resource "fmc_smart_license" "license" {
  registration_type = "REGISTER"
  token             = "X2M3YmJlY..."
}
```

Force to de-register and register with the token provided
```hcl
resource "fmc_smart_license" "license" {
  registration_type = "REGISTER"
  token             = "X2M3YmJlY..."
  force             = true
}
```

De-register on terraform destroy
```hcl
resource "fmc_smart_license" "license" {
  registration_type = "REGISTER"
  token             = "X2M3YmJlY..."
  retain            = false
}
```



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `registration_type` (String) The type of registration

### Optional

- `force` (Boolean) Force to unregister and register
- `retain` (Boolean) Set to false if you want to deregister on destroy
- `token` (String, Sensitive) Smart license token of the FMC

### Read-Only

- `id` (String) The ID of this resource.
- `registration_status` (String)


