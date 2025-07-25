---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_url Resource - terraform-provider-fmc"
subcategory: "Objects"
description: |-
  This resource manages an URL.
---

# fmc_url (Resource)

This resource manages an URL.

## Example Usage

```terraform
resource "fmc_url" "example" {
  name        = "my_url"
  url         = "https://www.example.com/app"
  description = "My URL"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the object.
- `url` (String) URL value.

### Optional

- `description` (String) Description of the object.
- `domain` (String) Name of the FMC domain
- `overridable` (Boolean) Indicates whether object values can be overridden.

### Read-Only

- `id` (String) Id of the object
- `type` (String) Type of the object; this value is always 'Url'.

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import fmc_url.example "<id>"
```
