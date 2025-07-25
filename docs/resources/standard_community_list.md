---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_standard_community_list Resource - terraform-provider-fmc"
subcategory: "Objects"
description: |-
  This resource manages a Standard Community List.
---

# fmc_standard_community_list (Resource)

This resource manages a Standard Community List.

## Example Usage

```terraform
resource "fmc_standard_community_list" "example" {
  name = "my_standard_community_list"
  entries = [
    {
      action       = "PERMIT"
      communities  = "123 456 789"
      internet     = true
      no_advertise = true
      no_export    = true
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `entries` (Attributes List) List of entries (see [below for nested schema](#nestedatt--entries))
- `name` (String) Name of the Standard Community List object.

### Optional

- `domain` (String) Name of the FMC domain

### Read-Only

- `id` (String) Id of the object
- `type` (String) Type of the object; this value is always 'CommunityListObject'.

<a id="nestedatt--entries"></a>
### Nested Schema for `entries`

Required:

- `action` (String) Indicate redistribution access.
  - Choices: `PERMIT`, `DENY`
- `communities` (String) List of communities. Separate multiple values by space. Valid values can be from 1 to 4294967295 or from 0:1 to 65534:65535

Optional:

- `internet` (Boolean) Specify Internet well-known community.
- `no_advertise` (Boolean) Specify No-advertise well-known community.
- `no_export` (Boolean) Specify No-export well-known community.

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import fmc_standard_community_list.example "<id>"
```
