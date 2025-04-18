---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_file_policy Data Source - terraform-provider-fmc"
subcategory: "Policies"
description: |-
  This data source reads the File Policy.
---

# fmc_file_policy (Data Source)

This data source reads the File Policy.

## Example Usage

```terraform
data "fmc_file_policy" "example" {
  id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `domain` (String) Name of the FMC domain
- `id` (String) Id of the object
- `name` (String) Name of file policy.

### Read-Only

- `block_encrypted_archives` (Boolean) Block encrypted archives
- `block_uninspectable_archives` (Boolean) Block uninspectable Archives
- `clean_list` (Boolean) Enable clean list
- `custom_detection_list` (Boolean) Enable custom detection list
- `description` (String) File policy description.
- `file_rules` (Attributes List) The ordered list of file rules. (see [below for nested schema](#nestedatt--file_rules))
- `first_time_file_analysis` (Boolean) Analyze first-seen files while AMP cloud disposition is pending
- `inspect_archives` (Boolean) Inspect Archives
- `max_archive_depth` (Number) Max archive depth
- `threat_score` (String) If AMP Cloud disposition is Unknown, override disposition based upon threat score.
- `type` (String) Type of the object

<a id="nestedatt--file_rules"></a>
### Nested Schema for `file_rules`

Read-Only:

- `action` (String) Action to be performed on a file.
- `application_protocol` (String) Defines a protocol for file inspection.
- `direction_of_transfer` (String) Direction of file transfer.
- `file_categories` (Attributes Set) Defines a list of file categories for inspection. (see [below for nested schema](#nestedatt--file_rules--file_categories))
- `file_types` (Attributes Set) Defines a list of file types for inspection. (see [below for nested schema](#nestedatt--file_rules--file_types))
- `id` (String) Id of File Rule
- `store_files` (Set of String) List of file dispositions that should be stored (MALWARE, CUSTOM, CLEAN, UNKNOWN).
- `type` (String) Type of File Rule.

<a id="nestedatt--file_rules--file_categories"></a>
### Nested Schema for `file_rules.file_categories`

Read-Only:

- `id` (String) The id of file category.
- `name` (String) The name of file category.
- `type` (String) The type of file category.


<a id="nestedatt--file_rules--file_types"></a>
### Nested Schema for `file_rules.file_types`

Read-Only:

- `id` (String) The id of file type.
- `name` (String) The name of file type.
- `type` (String) The name of file type.
