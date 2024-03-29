---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_time_range_object Resource - terraform-provider-fmc"
subcategory: ""
description: |-
  Resource for Time Range Object in FMC
  Example
  An example is shown below:
  hcl
  resource "fmc_time_range_object" "test" {
      name                  = "test-time-range"
      description           = "Test time range"
      effective_start_date = "2019-09-19T15:53:00"
      effective_end_date   = "2019-09-21T17:53:00"
  }
---

# fmc_time_range_object (Resource)

Resource for Time Range Object in FMC

## Example
An example is shown below: 
```hcl
resource "fmc_time_range_object" "test" {
    name        		  = "test-time-range"
    description 		  = "Test time range"
    effective_start_date = "2019-09-19T15:53:00"
    effective_end_date   = "2019-09-21T17:53:00"
}
```



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `effective_end_date` (String) Effective end date for this time range object (time in RFC3339 format)
- `effective_start_date` (String) Effective start date for this time range object (time in RFC3339 format)
- `name` (String) The name of this resource

### Optional

- `description` (String) The description of this resource
- `recurrence` (Block List) List of URL objects to add (see [below for nested schema](#nestedblock--recurrence))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--recurrence"></a>
### Nested Schema for `recurrence`

Required:

- `recurrence_type` (String) Type of recurrence. Allowed values: "DAILY_INTERVAL", "RANGE"

Optional:

- `daily_end_time` (String) Daily end time for this recurrence (time in RFC3339 format)
- `daily_start_time` (String) Daily start time for this recurrence (time in RFC3339 format)
- `days` (List of String)
- `end_day` (String) End day for this recurrence (time in RFC3339 format)
- `end_time` (String) End date for this recurrence (time in RFC3339 format)
- `start_day` (String) Start day for this recurrence (time in RFC3339 format)
- `start_time` (String) Start date for this recurrence (time in RFC3339 format)


