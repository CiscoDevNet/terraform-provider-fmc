// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DeviceCluster struct {
	Id                        types.String             `tfsdk:"id"`
	Domain                    types.String             `tfsdk:"domain"`
	Name                      types.String             `tfsdk:"name"`
	Type                      types.String             `tfsdk:"type"`
	ClusterKey                types.String             `tfsdk:"cluster_key"`
	ControlNodeVniPrefix      types.String             `tfsdk:"control_node_vni_prefix"`
	ControlNodeCclPrefix      types.String             `tfsdk:"control_node_ccl_prefix"`
	ControlNodeInterfaceId    types.String             `tfsdk:"control_node_interface_id"`
	ControlNodeInterfaceName  types.String             `tfsdk:"control_node_interface_name"`
	ControlNodeInterfaceType  types.String             `tfsdk:"control_node_interface_type"`
	ControlNodeDeviceId       types.String             `tfsdk:"control_node_device_id"`
	ControlNodeCclIpv4Address types.String             `tfsdk:"control_node_ccl_ipv4_address"`
	ControlNodePriority       types.Int64              `tfsdk:"control_node_priority"`
	DataNodes                 []DeviceClusterDataNodes `tfsdk:"data_nodes"`
}

type DeviceClusterDataNodes struct {
	DeviceId       types.String `tfsdk:"device_id"`
	CclIpv4Address types.String `tfsdk:"ccl_ipv4_address"`
	Priority       types.Int64  `tfsdk:"priority"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DeviceCluster) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/deviceclusters/ftddevicecluster"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DeviceCluster) toBody(ctx context.Context, state DeviceCluster) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.ClusterKey.IsNull() {
		body, _ = sjson.Set(body, "commonBootstrap.clusterKey", data.ClusterKey.ValueString())
	}
	if !data.ControlNodeVniPrefix.IsNull() {
		body, _ = sjson.Set(body, "commonBootstrap.vniNetwork", data.ControlNodeVniPrefix.ValueString())
	}
	if !data.ControlNodeCclPrefix.IsNull() {
		body, _ = sjson.Set(body, "commonBootstrap.cclNetwork", data.ControlNodeCclPrefix.ValueString())
	}
	if !data.ControlNodeInterfaceId.IsNull() {
		body, _ = sjson.Set(body, "commonBootstrap.cclInterface.id", data.ControlNodeInterfaceId.ValueString())
	}
	if !data.ControlNodeInterfaceName.IsNull() {
		body, _ = sjson.Set(body, "commonBootstrap.cclInterface.name", data.ControlNodeInterfaceName.ValueString())
	}
	if !data.ControlNodeInterfaceType.IsNull() {
		body, _ = sjson.Set(body, "commonBootstrap.cclInterface.type", data.ControlNodeInterfaceType.ValueString())
	}
	if !data.ControlNodeDeviceId.IsNull() {
		body, _ = sjson.Set(body, "controlDevice.deviceDetails.id", data.ControlNodeDeviceId.ValueString())
	}
	if !data.ControlNodeCclIpv4Address.IsNull() {
		body, _ = sjson.Set(body, "controlDevice.clusterNodeBootstrap.cclIp", data.ControlNodeCclIpv4Address.ValueString())
	}
	if !data.ControlNodePriority.IsNull() {
		body, _ = sjson.Set(body, "controlDevice.clusterNodeBootstrap.priority", data.ControlNodePriority.ValueInt64())
	}
	if len(data.DataNodes) > 0 {
		body, _ = sjson.Set(body, "dataDevices", []any{})
		for _, item := range data.DataNodes {
			itemBody := ""
			if !item.DeviceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "deviceDetails.id", item.DeviceId.ValueString())
			}
			if !item.CclIpv4Address.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "clusterNodeBootstrap.cclIp", item.CclIpv4Address.ValueString())
			}
			if !item.Priority.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "clusterNodeBootstrap.priority", item.Priority.ValueInt64())
			}
			body, _ = sjson.SetRaw(body, "dataDevices.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DeviceCluster) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("commonBootstrap.vniNetwork"); value.Exists() {
		data.ControlNodeVniPrefix = types.StringValue(value.String())
	} else {
		data.ControlNodeVniPrefix = types.StringNull()
	}
	if value := res.Get("commonBootstrap.cclNetwork"); value.Exists() {
		data.ControlNodeCclPrefix = types.StringValue(value.String())
	} else {
		data.ControlNodeCclPrefix = types.StringNull()
	}
	if value := res.Get("commonBootstrap.cclInterface.id"); value.Exists() {
		data.ControlNodeInterfaceId = types.StringValue(value.String())
	} else {
		data.ControlNodeInterfaceId = types.StringNull()
	}
	if value := res.Get("commonBootstrap.cclInterface.name"); value.Exists() {
		data.ControlNodeInterfaceName = types.StringValue(value.String())
	} else {
		data.ControlNodeInterfaceName = types.StringNull()
	}
	if value := res.Get("commonBootstrap.cclInterface.type"); value.Exists() {
		data.ControlNodeInterfaceType = types.StringValue(value.String())
	} else {
		data.ControlNodeInterfaceType = types.StringNull()
	}
	if value := res.Get("controlDevice.deviceDetails.id"); value.Exists() {
		data.ControlNodeDeviceId = types.StringValue(value.String())
	} else {
		data.ControlNodeDeviceId = types.StringNull()
	}
	if value := res.Get("controlDevice.clusterNodeBootstrap.cclIp"); value.Exists() {
		data.ControlNodeCclIpv4Address = types.StringValue(value.String())
	} else {
		data.ControlNodeCclIpv4Address = types.StringNull()
	}
	if value := res.Get("controlDevice.clusterNodeBootstrap.priority"); value.Exists() {
		data.ControlNodePriority = types.Int64Value(value.Int())
	} else {
		data.ControlNodePriority = types.Int64Null()
	}
	if value := res.Get("dataDevices"); value.Exists() {
		data.DataNodes = make([]DeviceClusterDataNodes, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceClusterDataNodes{}
			if value := res.Get("deviceDetails.id"); value.Exists() {
				data.DeviceId = types.StringValue(value.String())
			} else {
				data.DeviceId = types.StringNull()
			}
			if value := res.Get("clusterNodeBootstrap.cclIp"); value.Exists() {
				data.CclIpv4Address = types.StringValue(value.String())
			} else {
				data.CclIpv4Address = types.StringNull()
			}
			if value := res.Get("clusterNodeBootstrap.priority"); value.Exists() {
				data.Priority = types.Int64Value(value.Int())
			} else {
				data.Priority = types.Int64Null()
			}
			(*parent).DataNodes = append((*parent).DataNodes, data)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *DeviceCluster) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("commonBootstrap.vniNetwork"); value.Exists() && !data.ControlNodeVniPrefix.IsNull() {
		data.ControlNodeVniPrefix = types.StringValue(value.String())
	} else {
		data.ControlNodeVniPrefix = types.StringNull()
	}
	if value := res.Get("commonBootstrap.cclNetwork"); value.Exists() && !data.ControlNodeCclPrefix.IsNull() {
		data.ControlNodeCclPrefix = types.StringValue(value.String())
	} else {
		data.ControlNodeCclPrefix = types.StringNull()
	}
	if value := res.Get("commonBootstrap.cclInterface.id"); value.Exists() && !data.ControlNodeInterfaceId.IsNull() {
		data.ControlNodeInterfaceId = types.StringValue(value.String())
	} else {
		data.ControlNodeInterfaceId = types.StringNull()
	}
	if value := res.Get("commonBootstrap.cclInterface.name"); value.Exists() && !data.ControlNodeInterfaceName.IsNull() {
		data.ControlNodeInterfaceName = types.StringValue(value.String())
	} else {
		data.ControlNodeInterfaceName = types.StringNull()
	}
	if value := res.Get("commonBootstrap.cclInterface.type"); value.Exists() && !data.ControlNodeInterfaceType.IsNull() {
		data.ControlNodeInterfaceType = types.StringValue(value.String())
	} else {
		data.ControlNodeInterfaceType = types.StringNull()
	}
	if value := res.Get("controlDevice.deviceDetails.id"); value.Exists() && !data.ControlNodeDeviceId.IsNull() {
		data.ControlNodeDeviceId = types.StringValue(value.String())
	} else {
		data.ControlNodeDeviceId = types.StringNull()
	}
	if value := res.Get("controlDevice.clusterNodeBootstrap.cclIp"); value.Exists() && !data.ControlNodeCclIpv4Address.IsNull() {
		data.ControlNodeCclIpv4Address = types.StringValue(value.String())
	} else {
		data.ControlNodeCclIpv4Address = types.StringNull()
	}
	if value := res.Get("controlDevice.clusterNodeBootstrap.priority"); value.Exists() && !data.ControlNodePriority.IsNull() {
		data.ControlNodePriority = types.Int64Value(value.Int())
	} else {
		data.ControlNodePriority = types.Int64Null()
	}
	for i := 0; i < len(data.DataNodes); i++ {
		keys := [...]string{"deviceDetails.id"}
		keyValues := [...]string{data.DataNodes[i].DeviceId.ValueString()}

		parent := &data
		data := (*parent).DataNodes[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("dataDevices").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing DataNodes[%d] = %+v",
				i,
				(*parent).DataNodes[i],
			))
			(*parent).DataNodes = slices.Delete((*parent).DataNodes, i, i+1)
			i--

			continue
		}
		if value := res.Get("deviceDetails.id"); value.Exists() && !data.DeviceId.IsNull() {
			data.DeviceId = types.StringValue(value.String())
		} else {
			data.DeviceId = types.StringNull()
		}
		if value := res.Get("clusterNodeBootstrap.cclIp"); value.Exists() && !data.CclIpv4Address.IsNull() {
			data.CclIpv4Address = types.StringValue(value.String())
		} else {
			data.CclIpv4Address = types.StringNull()
		}
		if value := res.Get("clusterNodeBootstrap.priority"); value.Exists() && !data.Priority.IsNull() {
			data.Priority = types.Int64Value(value.Int())
		} else {
			data.Priority = types.Int64Null()
		}
		(*parent).DataNodes[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DeviceCluster) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin Clone

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// End of section. //template:end toBodyNonBulk

// Instead of deleting the cluster (which will de-register the appliances) we break it,
// making devices go back into standalone mode
func (data DeviceCluster) toBodyPutDelete(ctx context.Context, state DeviceCluster) string {
	body := ""
	body, _ = sjson.Set(body, "id", data.Id.ValueString())
	body, _ = sjson.Set(body, "action", "BREAK")
	return body
}

// Check if the bootstrap configuration has changed
func (data DeviceCluster) hasBootstrapChanged(ctx context.Context, plan DeviceCluster) bool {
	// if data.ClusterKey != plan.ClusterKey {
	// 	return true
	// }
	if data.ControlNodeVniPrefix != plan.ControlNodeVniPrefix {
		return true
	}
	if data.ControlNodeCclPrefix != plan.ControlNodeCclPrefix {
		return true
	}
	if data.ControlNodeInterfaceId != plan.ControlNodeInterfaceId {
		return true
	}
	if data.ControlNodeInterfaceName != plan.ControlNodeInterfaceName {
		return true
	}
	if data.ControlNodeInterfaceType != plan.ControlNodeInterfaceType {
		return true
	}
	if data.ControlNodePriority != plan.ControlNodePriority {
		return true
	}
	for _, dataNode := range data.DataNodes {
		for _, planNode := range plan.DataNodes {
			if dataNode.DeviceId == planNode.DeviceId {
				if dataNode.Priority != planNode.Priority {
					return true
				}
				if dataNode.CclIpv4Address != planNode.CclIpv4Address {
					return true
				}
			}
		}
	}

	return false
}
