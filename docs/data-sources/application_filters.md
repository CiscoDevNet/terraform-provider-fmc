---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_application_filters Data Source - terraform-provider-fmc"
subcategory: "Objects"
description: |-
  This data source reads the Application Filters.
---

# fmc_application_filters (Data Source)

This data source reads the Application Filters.

## Example Usage

```terraform
data "fmc_application_filters" "example" {
  items = {
    "my_application_filter" = {
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `domain` (String) Name of the FMC domain
- `items` (Attributes Map) Map of Application Filters. The key of the map is the name of the individual Application Filter. (see [below for nested schema](#nestedatt--items))

### Read-Only

- `id` (String) Id of the object

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `applications` (Attributes Set) Set of Applications. Either `applications` or `filters` must be specified. (see [below for nested schema](#nestedatt--items--applications))
- `filters` (Attributes List) List of Application filtering conditions. Either `applications` or `filters` must be specified. (see [below for nested schema](#nestedatt--items--filters))
- `id` (String) Id of the managed File Type.
- `type` (String) Type of the object; this value is always 'ApplicationFilter'.

<a id="nestedatt--items--applications"></a>
### Nested Schema for `items.applications`

Read-Only:

- `id` (String) Id of the Application.
- `name` (String) Name of the Application.


<a id="nestedatt--items--filters"></a>
### Nested Schema for `items.filters`

Read-Only:

- `business_relevances` (Attributes Set) Set of Application Business Relevances. (see [below for nested schema](#nestedatt--items--filters--business_relevances))
- `categories` (Attributes Set) Set of Application Categories. (see [below for nested schema](#nestedatt--items--filters--categories))
- `risks` (Attributes Set) Set of Application Risks. (see [below for nested schema](#nestedatt--items--filters--risks))
- `tags` (Attributes Set) Set of Application Tags. (see [below for nested schema](#nestedatt--items--filters--tags))
- `types` (Attributes Set) Set of Application Types. (see [below for nested schema](#nestedatt--items--filters--types))

<a id="nestedatt--items--filters--business_relevances"></a>
### Nested Schema for `items.filters.business_relevances`

Read-Only:

- `id` (String) Id of the Application Business Relevance.


<a id="nestedatt--items--filters--categories"></a>
### Nested Schema for `items.filters.categories`

Read-Only:

- `id` (String) Id of the Application Category.


<a id="nestedatt--items--filters--risks"></a>
### Nested Schema for `items.filters.risks`

Read-Only:

- `id` (String) Id of the Application Risk.


<a id="nestedatt--items--filters--tags"></a>
### Nested Schema for `items.filters.tags`

Read-Only:

- `id` (String) Id of the Application Tag.


<a id="nestedatt--items--filters--types"></a>
### Nested Schema for `items.filters.types`

Read-Only:

- `id` (String) Id of the Application Type.
