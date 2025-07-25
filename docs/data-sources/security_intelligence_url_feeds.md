---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_security_intelligence_url_feeds Data Source - terraform-provider-fmc"
subcategory: "Objects"
description: |-
  This data source reads the Security Intelligence URL Feeds.
---

# fmc_security_intelligence_url_feeds (Data Source)

This data source reads the Security Intelligence URL Feeds.

## Example Usage

```terraform
data "fmc_security_intelligence_url_feeds" "example" {
  items = {
    "my_si_url_feeds" = {
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `domain` (String) Name of the FMC domain
- `items` (Attributes Map) Map of Security Intelligence URL Feeds. The key of the map is the name of the individual Security Intelligence URL Feed. (see [below for nested schema](#nestedatt--items))

### Read-Only

- `id` (String) Id of the object

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `checksum_url` (String) URL with md5 checksum of the feed.
- `feed_url` (String) Security Intelligence Url Feed location.
- `id` (String) Id of the managed Security Intelligence URL Feed.
- `type` (String) Type of the object; this value is always 'SIURLFeed'.
- `update_frequency` (Number) Feed update frequency (in minutes).
