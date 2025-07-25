---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_secure_client_profile Resource - terraform-provider-fmc"
subcategory: "Objects"
description: |-
  This resource manages a Secure Client Profile.
---

# fmc_secure_client_profile (Resource)

This resource manages a Secure Client Profile.

## Example Usage

```terraform
resource "fmc_secure_client_profile" "example" {
  name        = "my_secure_client_profile"
  description = "My Secure Client Profile"
  file_type   = "ANYCONNECT_VPN_PROFILE"
  path        = "./package.xml"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `file_type` (String) Type of the Secure Client Profile file.
  - Choices: `ANYCONNECT_MANAGEMENT_VPN_PROFILE`, `AMP_ENABLER`, `FEEDBACK`, `WEB_SECURITY`, `ANYCONNECT_VPN_PROFILE`, `UMBRELLA_ROAMING`, `NETWORK_ACCESS_MANAGER`, `ISE_POSTURE`, `NETWORK_VISIBILITY`
- `name` (String) Name of the Secure Client Profile object.
- `path` (String) Path to the Secure Client Profile file. Supported file types are .xml, .asp, .fsp, .isp, .nsp, .nvmsp, .json, .wsp, .wso.

### Optional

- `description` (String) Description of the Secure Client Profile object.
- `domain` (String) Name of the FMC domain

### Read-Only

- `file_name` (String) Name of the Secure Client Profile file.
- `id` (String) Id of the object
- `type` (String) Type of the object; this value is always 'AnyConnectProfile'.

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import fmc_secure_client_profile.example "<id>"
```
