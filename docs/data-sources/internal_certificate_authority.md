---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_internal_certificate_authority Data Source - terraform-provider-fmc"
subcategory: "Objects"
description: |-
  This data source reads the Internal Certificate Authority.
---

# fmc_internal_certificate_authority (Data Source)

This data source reads the Internal Certificate Authority.

## Example Usage

```terraform
data "fmc_internal_certificate_authority" "example" {
  id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `domain` (String) Name of the FMC domain
- `id` (String) Id of the object
- `name` (String) Name of the Internal Certificate Authority (CA) object.

### Read-Only

- `certificate` (String) Certificate in PEM, DER, or PKCS#7 format.
- `password` (String, Sensitive) Private key password.
- `private_key` (String, Sensitive) Private key in PEM, DER, or PKCS#7 format.
- `type` (String) Type of the object; this value is always 'InternalCA'.
