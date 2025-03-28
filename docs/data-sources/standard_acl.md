---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_standard_acl Data Source - terraform-provider-fmc"
subcategory: "Objects"
description: |-
  This data source reads the Standard ACL.
---

# fmc_standard_acl (Data Source)

This data source reads the Standard ACL.

## Example Usage

```terraform
data "fmc_standard_acl" "example" {
  id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `domain` (String) Name of the FMC domain
- `id` (String) Id of the object
- `name` (String) Name of the object.

### Read-Only

- `description` (String) Description of the object.
- `entries` (Attributes List) Ordered list of ACL's entries. (see [below for nested schema](#nestedatt--entries))
- `type` (String) Type of the object; this value is always 'StandardAccessList'.

<a id="nestedatt--entries"></a>
### Nested Schema for `entries`

Read-Only:

- `action` (String) Indicates the redistribution access: PERMIT or DENY.
- `literals` (Attributes Set) Set of literal values. (see [below for nested schema](#nestedatt--entries--literals))
- `objects` (Attributes Set) Set of objects (Host, Network, Range). (see [below for nested schema](#nestedatt--entries--objects))

<a id="nestedatt--entries--literals"></a>
### Nested Schema for `entries.literals`

Read-Only:

- `value` (String)


<a id="nestedatt--entries--objects"></a>
### Nested Schema for `entries.objects`

Read-Only:

- `id` (String) Id of the object.
- `type` (String) Type of the object (such as fmc_network.this.type, etc.).
