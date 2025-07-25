---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_security_intelligence_dns_feed Data Source - terraform-provider-fmc"
subcategory: "Objects"
description: |-
  This data source reads the Security Intelligence DNS Feed.
---

# fmc_security_intelligence_dns_feed (Data Source)

This data source reads the Security Intelligence DNS Feed.

## Example Usage

```terraform
data "fmc_security_intelligence_dns_feed" "example" {
  id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `domain` (String) Name of the FMC domain
- `id` (String) Id of the object
- `name` (String) Name of the Security Intelligence DNS Feed.

### Read-Only

- `checksum_url` (String) URL with md5 checksum of the feed.
- `feed_url` (String) Security Intelligence DNS Feed location.
- `type` (String) Type of the object; this value is always 'SIDNSFeed'.
- `update_frequency` (Number) Feed update frequency (in minutes).
