---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_secure_client_external_browser_package Resource - terraform-provider-fmc"
subcategory: "Objects"
description: |-
  This resource manages a Secure Client External Browser Package.
---

# fmc_secure_client_external_browser_package (Resource)

This resource manages a Secure Client External Browser Package.

## Example Usage

```terraform
resource "fmc_secure_client_external_browser_package" "example" {
  name        = "my_secure_client_external_browser_package"
  description = "My Secure Client External Browser Package"
  path        = "./secure_client_external_browser_package.pkg"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the Secure Client External Browser Package object.
- `path` (String) Path to the Secure Client External Browser Package file. Supported file types are .pkg, .zip.

### Optional

- `description` (String) Description of the Secure Client External Browser Package object.
- `domain` (String) Name of the FMC domain

### Read-Only

- `file_name` (String) Name of the Secure Client External Browser Package file.
- `id` (String) Id of the object
- `type` (String) Type of the object; this value is always 'AnyConnectExternalBrowserPackage'.

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import fmc_secure_client_external_browser_package.example "<id>"
```
