---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_urls Resource - terraform-provider-fmc"
subcategory: "Objects"
description: |-
  This resource manages URLs through bulk operations.
  The following restrictions apply:
  Minimum FMC version for bulk object deletion: 7.4If FMC version does not meet the minimum version requirement for bulk operations, this resource will automatically fall back to processing operations one-by-one.Updates are always done one-by-one.
---

# fmc_urls (Resource)

This resource manages URLs through bulk operations.

The following restrictions apply:
  - Minimum FMC version for bulk object deletion: `7.4`
  - If FMC version does not meet the minimum version requirement for bulk operations, this resource will automatically fall back to processing operations one-by-one.
  - Updates are always done one-by-one.

## Example Usage

```terraform
resource "fmc_urls" "example" {
  items = {
    url_1 = {
      url         = "https://www.example.com/app"
      description = "My URL"
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `items` (Attributes Map) Map of security zones. The key of the map is the name of the individual URL object. (see [below for nested schema](#nestedatt--items))

### Optional

- `domain` (String) Name of the FMC domain

### Read-Only

- `id` (String) Id of the object

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Required:

- `url` (String) URL value.

Optional:

- `description` (String) Description of the object.
- `overridable` (Boolean) Indicates whether object values can be overridden.

Read-Only:

- `id` (String) Id of the managed URL object.
- `type` (String) Type of the object; this value is always 'Url'.

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import fmc_urls.example "<domain>,[<urls_name>]"
```
