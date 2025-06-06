---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_hosts Data Source - terraform-provider-fmc"
subcategory: "Objects"
description: |-
  This data source reads the Hosts.
---

# fmc_hosts (Data Source)

This data source reads the Hosts.

## Example Usage

```terraform
data "fmc_hosts" "example" {
  items = {
    "my_hosts" = {
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `domain` (String) Name of the FMC domain
- `items` (Attributes Map) Map of hosts. The key of the map is the name of the individual Host. (see [below for nested schema](#nestedatt--items))

### Read-Only

- `id` (String) Id of the object

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `description` (String) Description of the object.
- `id` (String) Id of the Host object.
- `ip` (String) IP of the host.
- `overridable` (Boolean) Indicates whether object values can be overridden.
- `type` (String) Type of the object; this value is always 'Host'.
