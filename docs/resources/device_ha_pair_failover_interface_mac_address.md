---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_device_ha_pair_failover_interface_mac_address Resource - terraform-provider-fmc"
subcategory: "Devices"
description: |-
  This resource manages a Device HA Pair Failover Interface MAC Address.
---

# fmc_device_ha_pair_failover_interface_mac_address (Resource)

This resource manages a Device HA Pair Failover Interface MAC Address.

## Example Usage

```terraform
resource "fmc_device_ha_pair_failover_interface_mac_address" "example" {
  ha_pair_id          = "76d24097-41c4-4558-a4d0-a8c07ac13928"
  interface_name      = "GigabitEthernet0/0"
  interface_id        = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  interface_type      = "PhysicalInterface"
  active_mac_address  = "c460.15e4.0edd"
  standby_mac_address = "c460.15e4.0ed0"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `active_mac_address` (String) MAC address of the active interface.
- `ha_pair_id` (String) Id of the parent HA Pair device.
- `interface_id` (String) Id of the interface or sub-interface.
- `interface_name` (String) Name of the physical interface. In case of sub-interfaces, this is the name of the parent interface (fmc_device_subinterface.x.interface_name).
- `interface_type` (String) Type of the interface.
- `standby_mac_address` (String) MAC address of the standby interface.

### Optional

- `domain` (String) Name of the FMC domain

### Read-Only

- `id` (String) Id of the object
- `type` (String) Type of the resource; This is always `FailoverInterfaceMacAddressConfig`.

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import fmc_device_ha_pair_failover_interface_mac_address.example "<ha_pair_id>,<id>"
```
