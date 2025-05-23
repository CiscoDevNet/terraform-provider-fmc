---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_device_cluster Data Source - terraform-provider-fmc"
subcategory: "Devices"
description: |-
  This data source reads the Device Cluster.
---

# fmc_device_cluster (Data Source)

This data source reads the Device Cluster.

## Example Usage

```terraform
data "fmc_device_cluster" "example" {
  id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `domain` (String) Name of the FMC domain
- `id` (String) Id of the object
- `name` (String) Name of the FTD Cluster.

### Read-Only

- `cluster_key` (String) Secret key for the cluster, between 1 nd 63 characters.
- `control_node_ccl_ipv4_address` (String) Cluster control link IPv4 address / VTEP IPv4 address.
- `control_node_ccl_prefix` (String) Cluster Control Link Network / Virtual Tunnel Endpoint (VTEP) Network
- `control_node_device_id` (String) Cluster Control Node device ID.
- `control_node_interface_id` (String) Cluster control link interface ID.
- `control_node_interface_name` (String) Cluster control link interface Name.
- `control_node_interface_type` (String) Cluster control link interface Type.
- `control_node_priority` (Number) Priority of cluster controle node.
- `control_node_vni_prefix` (String) Cluster Control VXLAN Network Identifier (VNI) Network
- `data_devices` (Attributes List) List of cluster data nodes. (see [below for nested schema](#nestedatt--data_devices))
- `type` (String) Type of the resource; This is always `DeviceCluster`.

<a id="nestedatt--data_devices"></a>
### Nested Schema for `data_devices`

Read-Only:

- `data_node_ccl_ipv4_address` (String) Cluster Data Node link IPv4 address / VTEP IPv4 address.
- `data_node_device_id` (String) Cluster Data Node device ID.
- `data_node_priority` (Number) Priority of cluster data node.
