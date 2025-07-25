---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_internal_certificate Resource - terraform-provider-fmc"
subcategory: "Objects"
description: |-
  This resource manages an Internal Certificate.
---

# fmc_internal_certificate (Resource)

This resource manages an Internal Certificate.

## Example Usage

```terraform
resource "fmc_internal_certificate" "example" {
  name        = "my_internal_certificate"
  certificate = "-----BEGIN CERTIFICATE-----\nMII(...)\n-----END CERTIFICATE-----"
  private_key = "-----BEGIN RSA PRIVATE KEY-----\nProc-Type: 4,ENCRYPTED\nDEK-Info: AES-128-CBC,D55(...)\n-----END RSA PRIVATE KEY-----"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the Internal Certificate object.

### Optional

- `certificate` (String) Certificate in PEM, DER, or PKCS#7 format.
- `domain` (String) Name of the FMC domain
- `password` (String, Sensitive) Private key password.
- `private_key` (String, Sensitive) Private key in PEM, DER, or PKCS#7 format.

### Read-Only

- `id` (String) Id of the object
- `type` (String) Type of the object; this value is always 'InternalCertificate'.

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import fmc_internal_certificate.example "<id>"
```
