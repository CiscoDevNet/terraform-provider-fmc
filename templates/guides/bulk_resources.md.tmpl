---
subcategory: "Guides"
page_title: "Bulk resources"
description: |-
    Bulk resources
---

# Bulk resources

## Introduction

The Secure Firewall Management Center (FMC) API supports bulk operations for certain resources. These operations allow you to:

- **Create** / **Delete** multiple resources in a single API call
- **Update** resources individually (updates are always performed one by one)

This brings several benefits:

- **Improved Performance**: Significantly faster when managing large numbers of resources
- **Reduced API Calls**: Fewer requests to the FMC API

## FMC API Limitations

Currently, the following FMC API restrictions apply:
- Maximum of 10 concurrent read operations
- Maximum of 1 concurrent create/update/delete operation
- Maximum of 120 req/min or 300 req/min (FMC 7.4.1 and later)
- Bulk update operations are not supported

## Usage
Bulk resources can be distinguished by the plural form in their names, eg. `fmc_hosts` is bulk version of `fmc_host`.

**Example:**
```hcl
# Individual resources
resource "fmc_host" "web_server" {
  name = "web_server"
  ip   = "10.0.1.10"
}

resource "fmc_host" "app_server" {
  name = "app_server"
  ip   = "10.0.1.20"
}

resource "fmc_host" "db_server" {
  name = "db_server"
  ip   = "10.0.1.30"
}

# Bulk resource
resource "fmc_hosts" "example" {
  items = {
    "web_server"   = { ip = "10.0.1.10" }
    "app_server"   = { ip = "10.0.1.20" }
    "db_server"    = { ip = "10.0.1.30" }
  }
}
```

## Performance comparision
The following table illustrates the performance difference between using individual resources and bulk resources for managing 1200 hosts.

| Operation | Individual Resources | Bulk Resource |
|-----------|----------------------|---------------|
| Create    | ~30 minutes          | ~6 minutes    |
| Refresh   | ~4 minutes           | ~17 seconds   |
| Destroy   | ~46 minutes          | ~20 minutes   |

Tests were performed in the following conditions:
- FMC version: 7.7.10
- Terraform version: 1.13.4
- Terraform provider version: 2.0.0
- FMC configured with default settings and no pre-existing objects
- Table values are approximate and may vary based on environment and conditions, like FMC load and existing objects
- Updates are done one by one in both cases, therefore update time of a single object is mostly equal to Refresh time

## Limitations

### Object Replacement

When replacing objects referenced by other resources, you may encounter dependency conflicts that require staged deployments.

#### Example Scenario

**Initial Configuration:**
```hcl
resource "fmc_hosts" "hosts" {
  items = {
    "host_1" = { ip = "10.10.1.1" }
    "host_2" = { ip = "10.10.1.2" }
  }
}

resource "fmc_network_group" "network_group" {
  name = "network_group"
  objects = [
    { id = fmc_hosts.hosts.items["host_1"].id },
    { id = fmc_hosts.hosts.items["host_2"].id },
  ]
}
```

**Desired Change:** Replace `host_2` with `host_3`

```hcl
resource "fmc_hosts" "hosts" {
  items = {
    "host_1" = { ip = "10.10.1.1" }
    "host_3" = { ip = "10.10.1.3" }
  }
}

resource "fmc_network_group" "network_group" {
  name = "network_group"
  objects = [
    { id = fmc_hosts.hosts.items["host_1"].id },
    { id = fmc_hosts.hosts.items["host_3"].id },
  ]
}
```

**Problem:** Terraform cannot apply this change in a single run because:
- It cannot modify `fmc_hosts` first, as this would break `fmc_network_group` reference to `host_2`
- It cannot modify `fmc_network_group` first, as `host_3` doesn't exist yet

#### Solution: Staged Deployment

**Stage 1:** Add the new host and update references
```hcl
resource "fmc_hosts" "hosts" {
  items = {
    "host_1" = { ip = "10.10.1.1" }
    "host_2" = { ip = "10.10.1.2" }  # Keep temporarily
    "host_3" = { ip = "10.10.1.3" }  # Add new host
  }
}

resource "fmc_network_group" "network_group" {
  name = "network_group"
  objects = [
    { id = fmc_hosts.hosts.items["host_1"].id },
    { id = fmc_hosts.hosts.items["host_3"].id },  # Switch to new host
  ]
}
```

**Stage 2:** Remove the old host
```hcl
resource "fmc_hosts" "hosts" {
  items = {
    "host_1" = { ip = "10.10.1.1" }
    "host_3" = { ip = "10.10.1.3" }  # host_2 removed
  }
}

# network_group remains unchanged
resource "fmc_network_group" "network_group" {
  name = "network_group"
  objects = [
    { id = fmc_hosts.hosts.items["host_1"].id },
    { id = fmc_hosts.hosts.items["host_3"].id },
  ]
}
```

This two-stage approach ensures all dependencies are properly maintained throughout the replacement process.

### Object Removal

When removing objects that are referenced by other resources, you cannot do it in a single step.

#### Example Scenario

**Initial Configuration:**
```hcl
resource "fmc_hosts" "hosts" {
  items = {
    "host_1" = { ip = "10.10.1.1" }
    "host_2" = { ip = "10.10.1.2" }
  }
}

resource "fmc_network_group" "network_group" {
  name = "network_group"
  objects = [
    { id = fmc_hosts.hosts.items["host_1"].id },
    { id = fmc_hosts.hosts.items["host_2"].id },
  ]
}
```

**Desired Change:** Remove `host_2`

```hcl
resource "fmc_hosts" "hosts" {
  items = {
    "host_1" = { ip = "10.10.1.1" }
  }
}

resource "fmc_network_group" "network_group" {
  name = "network_group"
  objects = [
    { id = fmc_hosts.hosts.items["host_1"].id },
  ]
}
```

**Problem:** Terraform cannot determine the correct order of operations.

#### Solution: Staged Deployment

**Stage 1:** Remove references to the object

```hcl
# hosts unchanged
resource "fmc_hosts" "hosts" {
  items = {
    "host_1" = { ip = "10.10.1.1" }
    "host_2" = { ip = "10.10.1.2" }
  }
}

resource "fmc_network_group" "network_group" {
  name = "network_group"
  objects = [
    { id = fmc_hosts.hosts.items["host_1"].id },  # host_2 removed
  ]
}
```

**Stage 2:** Remove the object itself
```hcl
resource "fmc_hosts" "hosts" {
  items = {
    "host_1" = { ip = "10.10.1.1" }   # host_2 removed
  }
}

# network group remains unchanged
resource "fmc_network_group" "network_group" {
  name = "network_group"
  objects = [
    { id = fmc_hosts.hosts.items["host_1"].id },
  ]
}
```