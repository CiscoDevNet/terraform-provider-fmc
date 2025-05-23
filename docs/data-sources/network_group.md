---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_network_group Data Source - terraform-provider-fmc"
subcategory: "Objects"
description: |-
  This data source reads the Network Group.
---

# fmc_network_group (Data Source)

This data source reads the Network Group.

## Example Usage

```terraform
data "fmc_network_group" "example" {
  id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `domain` (String) Name of the FMC domain
- `id` (String) Id of the object
- `name` (String) Name of the Network Group object.

### Read-Only

- `description` (String) Description of the ojbect.
- `literals` (Attributes Set) Set of literal values (Host or Network). (see [below for nested schema](#nestedatt--literals))
- `objects` (Attributes Set) Set of network objects (Host, Network, Range, FQDN or Network Group). (see [below for nested schema](#nestedatt--objects))
- `overridable` (Boolean) Indicates whether object values can be overridden.
- `type` (String) Type of the object; this value is always 'NetworkGroup'.

<a id="nestedatt--literals"></a>
### Nested Schema for `literals`

Read-Only:

- `value` (String)


<a id="nestedatt--objects"></a>
### Nested Schema for `objects`

Read-Only:

- `id` (String) ID of the network object
