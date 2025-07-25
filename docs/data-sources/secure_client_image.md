---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_secure_client_image Data Source - terraform-provider-fmc"
subcategory: "Objects"
description: |-
  This data source reads the Secure Client Image.
---

# fmc_secure_client_image (Data Source)

This data source reads the Secure Client Image.

## Example Usage

```terraform
data "fmc_secure_client_image" "example" {
  id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `domain` (String) Name of the FMC domain
- `id` (String) Id of the object
- `name` (String) Name of the Secure Client Image object.

### Read-Only

- `description` (String) Description of the Secure Client Image object.
- `file_name` (String) Name of the Secure Client Image file.
- `path` (String) Path to the Secure Client Image file. Supported file type is .pkg.
- `type` (String) Type of the object; this value is always 'AnyConnectPackage'.
