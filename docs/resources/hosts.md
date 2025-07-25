---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_hosts Resource - terraform-provider-fmc"
subcategory: "Objects"
description: |-
  This resource manages Hosts through bulk operations.
  The following restrictions apply:
  Minimum FMC version for bulk object deletion: 7.4If FMC version does not meet the minimum version requirement for bulk operations, this resource will automatically fall back to processing operations one-by-one.Updates are always done one-by-one.
---

# fmc_hosts (Resource)

This resource manages Hosts through bulk operations.

The following restrictions apply:
  - Minimum FMC version for bulk object deletion: `7.4`
  - If FMC version does not meet the minimum version requirement for bulk operations, this resource will automatically fall back to processing operations one-by-one.
  - Updates are always done one-by-one.

## Example Usage

```terraform
resource "fmc_hosts" "example" {
  items = {
    my_hosts = {
      description = "My Host 1"
      ip          = "10.1.1.1"
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `items` (Attributes Map) Map of hosts. The key of the map is the name of the individual Host. (see [below for nested schema](#nestedatt--items))

### Optional

- `domain` (String) Name of the FMC domain

### Read-Only

- `id` (String) Id of the object

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Required:

- `ip` (String) IP of the host.

Optional:

- `description` (String) Description of the object.
- `overridable` (Boolean) Indicates whether object values can be overridden.

Read-Only:

- `id` (String) Id of the Host object.
- `type` (String) Type of the object; this value is always 'Host'.

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import fmc_hosts.example "<domain>,[<hosts_name>]"
```
