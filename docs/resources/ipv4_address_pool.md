---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_ipv4_address_pool Resource - terraform-provider-fmc"
subcategory: "Objects"
description: |-
  This resource manages an IPv4 Address Pool.
---

# fmc_ipv4_address_pool (Resource)

This resource manages an IPv4 Address Pool.

## Example Usage

```terraform
resource "fmc_ipv4_address_pool" "example" {
  name        = "my_ipv4_address_pool"
  description = "My IPv4 Address Pool object"
  range       = "10.0.0.10-10.0.0.20"
  netmask     = "255.255.255.0"
  overridable = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the IPv4 Address Pool object.
- `netmask` (String) IP netmask for the range
- `range` (String) IP range

### Optional

- `description` (String) Description of the object.
- `domain` (String) Name of the FMC domain
- `overridable` (Boolean) Whether the object values can be overridden.

### Read-Only

- `id` (String) Id of the object
- `type` (String) Type of the object; this value is always 'IPv4AddressPool'.

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import fmc_ipv4_address_pool.example "<id>"
```
