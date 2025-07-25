---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_file_policy Resource - terraform-provider-fmc"
subcategory: "Policies"
description: |-
  This resource manages File Policy with corresponding rules.
  The following restrictions apply:
  Read operations are supported by any tested FMC versionMinimum FMC version for object management (Create/Update/Delete): 7.4
---

# fmc_file_policy (Resource)

This resource manages File Policy with corresponding rules.

The following restrictions apply:
  - Read operations are supported by any tested FMC version
  - Minimum FMC version for object management (Create/Update/Delete): `7.4`

## Example Usage

```terraform
resource "fmc_file_policy" "example" {
  name                         = "my_file_policy"
  description                  = "My file policy"
  first_time_file_analysis     = true
  custom_detection_list        = true
  clean_list                   = true
  threat_score                 = "DISABLED"
  inspect_archives             = false
  block_encrypted_archives     = false
  block_uninspectable_archives = false
  max_archive_depth            = 2
  file_rules = [
    {
      application_protocol  = "ANY"
      action                = "DETECT"
      direction_of_transfer = "ANY"
      file_categories = [
        {
          id   = "5"
          name = "PDF files"
        }
      ]
      file_types = [
        {
          id   = "19"
          name = "7Z"
        }
      ]
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of file policy.

### Optional

- `block_encrypted_archives` (Boolean) Block encrypted archives
- `block_uninspectable_archives` (Boolean) Block uninspectable Archives
- `clean_list` (Boolean) Enable clean list
- `custom_detection_list` (Boolean) Enable custom detection list
- `description` (String) File policy description.
- `domain` (String) Name of the FMC domain
- `file_rules` (Attributes List) The ordered list of file rules. (see [below for nested schema](#nestedatt--file_rules))
- `first_time_file_analysis` (Boolean) Analyze first-seen files while AMP cloud disposition is pending
- `inspect_archives` (Boolean) Inspect Archives
- `max_archive_depth` (Number) Max archive depth
  - Range: `1`-`3`
- `threat_score` (String) If AMP Cloud disposition is Unknown, override disposition based upon threat score.
  - Choices: `DISABLED`, `MEDIUM`, `High`, `VERY_HIGH`

### Read-Only

- `id` (String) Id of the object
- `type` (String) Type of the object

<a id="nestedatt--file_rules"></a>
### Nested Schema for `file_rules`

Required:

- `action` (String) Action to be performed on a file.
  - Choices: `DETECT`, `BLOCK_WITH_RESET`, `DETECT_MALWARE`, `BLOCK_MALWARE_WITH_RESET`
- `application_protocol` (String) Defines a protocol for file inspection.
  - Choices: `ANY`, `HTTP`, `SMTP`, `IMAP`, `POP3`, `FTP`, `SMB`
- `direction_of_transfer` (String) Direction of file transfer.
  - Choices: `ANY`, `UPLOAD`, `DOWNLOAD`

Optional:

- `file_categories` (Attributes Set) Defines a list of file categories for inspection. (see [below for nested schema](#nestedatt--file_rules--file_categories))
- `file_types` (Attributes Set) Defines a list of file types for inspection. (see [below for nested schema](#nestedatt--file_rules--file_types))
- `store_files` (Set of String) List of file dispositions that should be stored (MALWARE, CUSTOM, CLEAN, UNKNOWN).

Read-Only:

- `id` (String) Id of File Rule
- `type` (String) Type of File Rule.

<a id="nestedatt--file_rules--file_categories"></a>
### Nested Schema for `file_rules.file_categories`

Required:

- `id` (String) The id of file category.
- `name` (String) The name of file category.

Optional:

- `type` (String) The type of file category.
  - Default value: `FileCategory`


<a id="nestedatt--file_rules--file_types"></a>
### Nested Schema for `file_rules.file_types`

Required:

- `id` (String) The id of file type.
- `name` (String) The name of file type.

Optional:

- `type` (String) The name of file type.
  - Default value: `FileType`

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import fmc_file_policy.example "<id>"
```
