---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_application_tag Data Source - terraform-provider-fmc"
subcategory: "Objects"
description: |-
  This data source reads the Application Tag.
---

# fmc_application_tag (Data Source)

This data source reads the Application Tag.

## Example Usage

```terraform
data "fmc_application_tag" "example" {
  id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `domain` (String) Name of the FMC domain
- `id` (String) Id of the object
- `name` (String) Name of the Application Tag.

### Read-Only

- `type` (String) Type of the object; this value is always 'ApplicationTag'.
