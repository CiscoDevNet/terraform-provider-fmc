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
	"net/url"

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DeviceClusterHealthMonitor struct {
	Id                                          types.String  `tfsdk:"id"`
	Domain                                      types.String  `tfsdk:"domain"`
	ClusterId                                   types.String  `tfsdk:"cluster_id"`
	Type                                        types.String  `tfsdk:"type"`
	HealthCheck                                 types.Bool    `tfsdk:"health_check"`
	HoldTime                                    types.Float64 `tfsdk:"hold_time"`
	DebounceTime                                types.Int64   `tfsdk:"debounce_time"`
	DataInterfaceAutoRejoinAttempts             types.Int64   `tfsdk:"data_interface_auto_rejoin_attempts"`
	DataInterfaceAutoRejoinInterval             types.Int64   `tfsdk:"data_interface_auto_rejoin_interval"`
	DataInterfaceAutoRejoinIntervalVariation    types.Int64   `tfsdk:"data_interface_auto_rejoin_interval_variation"`
	ClusterInterfaceAutoRejoinAttempts          types.Int64   `tfsdk:"cluster_interface_auto_rejoin_attempts"`
	ClusterInterfaceAutoRejoinInterval          types.Int64   `tfsdk:"cluster_interface_auto_rejoin_interval"`
	ClusterInterfaceAutoRejoinIntervalVariation types.Int64   `tfsdk:"cluster_interface_auto_rejoin_interval_variation"`
	SystemAutoRejoinAttempts                    types.Int64   `tfsdk:"system_auto_rejoin_attempts"`
	SystemAutoRejoinInterval                    types.Int64   `tfsdk:"system_auto_rejoin_interval"`
	SystemAutoRejoinIntervalVariation           types.Int64   `tfsdk:"system_auto_rejoin_interval_variation"`
	UnmonitoredInterfaces                       types.Set     `tfsdk:"unmonitored_interfaces"`
	ServiceApplicationMonitoring                types.Bool    `tfsdk:"service_application_monitoring"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DeviceClusterHealthMonitor) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/deviceclusters/ftddevicecluster/%v/clusterhealthmonitorsettings", url.QueryEscape(data.ClusterId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DeviceClusterHealthMonitor) toBody(ctx context.Context, state DeviceClusterHealthMonitor) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.HealthCheck.IsNull() {
		body, _ = sjson.Set(body, "isHealthCheckEnabled", data.HealthCheck.ValueBool())
	}
	if !data.HoldTime.IsNull() {
		body, _ = sjson.Set(body, "holdTime", data.HoldTime.ValueFloat64())
	}
	if !data.DebounceTime.IsNull() {
		body, _ = sjson.Set(body, "monitorInterface.debounceTime", data.DebounceTime.ValueInt64())
	}
	if !data.DataInterfaceAutoRejoinAttempts.IsNull() {
		body, _ = sjson.Set(body, "interfaceAutoRejoin.dataInterface.attempts", data.DataInterfaceAutoRejoinAttempts.ValueInt64())
	}
	if !data.DataInterfaceAutoRejoinInterval.IsNull() {
		body, _ = sjson.Set(body, "interfaceAutoRejoin.dataInterface.interval", data.DataInterfaceAutoRejoinInterval.ValueInt64())
	}
	if !data.DataInterfaceAutoRejoinIntervalVariation.IsNull() {
		body, _ = sjson.Set(body, "interfaceAutoRejoin.dataInterface.nextInterval", data.DataInterfaceAutoRejoinIntervalVariation.ValueInt64())
	}
	if !data.ClusterInterfaceAutoRejoinAttempts.IsNull() {
		body, _ = sjson.Set(body, "interfaceAutoRejoin.clusterInterface.attempts", data.ClusterInterfaceAutoRejoinAttempts.ValueInt64())
	}
	if !data.ClusterInterfaceAutoRejoinInterval.IsNull() {
		body, _ = sjson.Set(body, "interfaceAutoRejoin.clusterInterface.interval", data.ClusterInterfaceAutoRejoinInterval.ValueInt64())
	}
	if !data.ClusterInterfaceAutoRejoinIntervalVariation.IsNull() {
		body, _ = sjson.Set(body, "interfaceAutoRejoin.clusterInterface.nextInterval", data.ClusterInterfaceAutoRejoinIntervalVariation.ValueInt64())
	}
	if !data.SystemAutoRejoinAttempts.IsNull() {
		body, _ = sjson.Set(body, "systemAutoRejoin.attempts", data.SystemAutoRejoinAttempts.ValueInt64())
	}
	if !data.SystemAutoRejoinInterval.IsNull() {
		body, _ = sjson.Set(body, "systemAutoRejoin.interval", data.SystemAutoRejoinInterval.ValueInt64())
	}
	if !data.SystemAutoRejoinIntervalVariation.IsNull() {
		body, _ = sjson.Set(body, "systemAutoRejoin.nextInterval", data.SystemAutoRejoinIntervalVariation.ValueInt64())
	}
	if !data.UnmonitoredInterfaces.IsNull() {
		var values []string
		data.UnmonitoredInterfaces.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "unMonitorInterface", values)
	}
	if !data.ServiceApplicationMonitoring.IsNull() {
		body, _ = sjson.Set(body, "monitorInterface.isServiceApplicationEnabled", data.ServiceApplicationMonitoring.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DeviceClusterHealthMonitor) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("isHealthCheckEnabled"); value.Exists() {
		data.HealthCheck = types.BoolValue(value.Bool())
	} else {
		data.HealthCheck = types.BoolNull()
	}
	if value := res.Get("holdTime"); value.Exists() {
		data.HoldTime = types.Float64Value(value.Float())
	} else {
		data.HoldTime = types.Float64Null()
	}
	if value := res.Get("monitorInterface.debounceTime"); value.Exists() {
		data.DebounceTime = types.Int64Value(value.Int())
	} else {
		data.DebounceTime = types.Int64Null()
	}
	if value := res.Get("interfaceAutoRejoin.dataInterface.attempts"); value.Exists() {
		data.DataInterfaceAutoRejoinAttempts = types.Int64Value(value.Int())
	} else {
		data.DataInterfaceAutoRejoinAttempts = types.Int64Null()
	}
	if value := res.Get("interfaceAutoRejoin.dataInterface.interval"); value.Exists() {
		data.DataInterfaceAutoRejoinInterval = types.Int64Value(value.Int())
	} else {
		data.DataInterfaceAutoRejoinInterval = types.Int64Null()
	}
	if value := res.Get("interfaceAutoRejoin.dataInterface.nextInterval"); value.Exists() {
		data.DataInterfaceAutoRejoinIntervalVariation = types.Int64Value(value.Int())
	} else {
		data.DataInterfaceAutoRejoinIntervalVariation = types.Int64Null()
	}
	if value := res.Get("interfaceAutoRejoin.clusterInterface.attempts"); value.Exists() {
		data.ClusterInterfaceAutoRejoinAttempts = types.Int64Value(value.Int())
	} else {
		data.ClusterInterfaceAutoRejoinAttempts = types.Int64Null()
	}
	if value := res.Get("interfaceAutoRejoin.clusterInterface.interval"); value.Exists() {
		data.ClusterInterfaceAutoRejoinInterval = types.Int64Value(value.Int())
	} else {
		data.ClusterInterfaceAutoRejoinInterval = types.Int64Null()
	}
	if value := res.Get("interfaceAutoRejoin.clusterInterface.nextInterval"); value.Exists() {
		data.ClusterInterfaceAutoRejoinIntervalVariation = types.Int64Value(value.Int())
	} else {
		data.ClusterInterfaceAutoRejoinIntervalVariation = types.Int64Null()
	}
	if value := res.Get("systemAutoRejoin.attempts"); value.Exists() {
		data.SystemAutoRejoinAttempts = types.Int64Value(value.Int())
	} else {
		data.SystemAutoRejoinAttempts = types.Int64Null()
	}
	if value := res.Get("systemAutoRejoin.interval"); value.Exists() {
		data.SystemAutoRejoinInterval = types.Int64Value(value.Int())
	} else {
		data.SystemAutoRejoinInterval = types.Int64Null()
	}
	if value := res.Get("systemAutoRejoin.nextInterval"); value.Exists() {
		data.SystemAutoRejoinIntervalVariation = types.Int64Value(value.Int())
	} else {
		data.SystemAutoRejoinIntervalVariation = types.Int64Null()
	}
	if value := res.Get("unMonitorInterface"); value.Exists() {
		data.UnmonitoredInterfaces = helpers.GetStringSet(value.Array())
	} else {
		data.UnmonitoredInterfaces = types.SetNull(types.StringType)
	}
	if value := res.Get("monitorInterface.isServiceApplicationEnabled"); value.Exists() {
		data.ServiceApplicationMonitoring = types.BoolValue(value.Bool())
	} else {
		data.ServiceApplicationMonitoring = types.BoolNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *DeviceClusterHealthMonitor) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("isHealthCheckEnabled"); value.Exists() && !data.HealthCheck.IsNull() {
		data.HealthCheck = types.BoolValue(value.Bool())
	} else {
		data.HealthCheck = types.BoolNull()
	}
	if value := res.Get("holdTime"); value.Exists() && !data.HoldTime.IsNull() {
		data.HoldTime = types.Float64Value(value.Float())
	} else {
		data.HoldTime = types.Float64Null()
	}
	if value := res.Get("monitorInterface.debounceTime"); value.Exists() && !data.DebounceTime.IsNull() {
		data.DebounceTime = types.Int64Value(value.Int())
	} else {
		data.DebounceTime = types.Int64Null()
	}
	if value := res.Get("interfaceAutoRejoin.dataInterface.attempts"); value.Exists() && !data.DataInterfaceAutoRejoinAttempts.IsNull() {
		data.DataInterfaceAutoRejoinAttempts = types.Int64Value(value.Int())
	} else {
		data.DataInterfaceAutoRejoinAttempts = types.Int64Null()
	}
	if value := res.Get("interfaceAutoRejoin.dataInterface.interval"); value.Exists() && !data.DataInterfaceAutoRejoinInterval.IsNull() {
		data.DataInterfaceAutoRejoinInterval = types.Int64Value(value.Int())
	} else {
		data.DataInterfaceAutoRejoinInterval = types.Int64Null()
	}
	if value := res.Get("interfaceAutoRejoin.dataInterface.nextInterval"); value.Exists() && !data.DataInterfaceAutoRejoinIntervalVariation.IsNull() {
		data.DataInterfaceAutoRejoinIntervalVariation = types.Int64Value(value.Int())
	} else {
		data.DataInterfaceAutoRejoinIntervalVariation = types.Int64Null()
	}
	if value := res.Get("interfaceAutoRejoin.clusterInterface.attempts"); value.Exists() && !data.ClusterInterfaceAutoRejoinAttempts.IsNull() {
		data.ClusterInterfaceAutoRejoinAttempts = types.Int64Value(value.Int())
	} else {
		data.ClusterInterfaceAutoRejoinAttempts = types.Int64Null()
	}
	if value := res.Get("interfaceAutoRejoin.clusterInterface.interval"); value.Exists() && !data.ClusterInterfaceAutoRejoinInterval.IsNull() {
		data.ClusterInterfaceAutoRejoinInterval = types.Int64Value(value.Int())
	} else {
		data.ClusterInterfaceAutoRejoinInterval = types.Int64Null()
	}
	if value := res.Get("interfaceAutoRejoin.clusterInterface.nextInterval"); value.Exists() && !data.ClusterInterfaceAutoRejoinIntervalVariation.IsNull() {
		data.ClusterInterfaceAutoRejoinIntervalVariation = types.Int64Value(value.Int())
	} else {
		data.ClusterInterfaceAutoRejoinIntervalVariation = types.Int64Null()
	}
	if value := res.Get("systemAutoRejoin.attempts"); value.Exists() && !data.SystemAutoRejoinAttempts.IsNull() {
		data.SystemAutoRejoinAttempts = types.Int64Value(value.Int())
	} else {
		data.SystemAutoRejoinAttempts = types.Int64Null()
	}
	if value := res.Get("systemAutoRejoin.interval"); value.Exists() && !data.SystemAutoRejoinInterval.IsNull() {
		data.SystemAutoRejoinInterval = types.Int64Value(value.Int())
	} else {
		data.SystemAutoRejoinInterval = types.Int64Null()
	}
	if value := res.Get("systemAutoRejoin.nextInterval"); value.Exists() && !data.SystemAutoRejoinIntervalVariation.IsNull() {
		data.SystemAutoRejoinIntervalVariation = types.Int64Value(value.Int())
	} else {
		data.SystemAutoRejoinIntervalVariation = types.Int64Null()
	}
	if value := res.Get("unMonitorInterface"); value.Exists() && !data.UnmonitoredInterfaces.IsNull() {
		data.UnmonitoredInterfaces = helpers.GetStringSet(value.Array())
	} else {
		data.UnmonitoredInterfaces = types.SetNull(types.StringType)
	}
	if value := res.Get("monitorInterface.isServiceApplicationEnabled"); value.Exists() && !data.ServiceApplicationMonitoring.IsNull() {
		data.ServiceApplicationMonitoring = types.BoolValue(value.Bool())
	} else {
		data.ServiceApplicationMonitoring = types.BoolNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DeviceClusterHealthMonitor) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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

// Section below is generated&owned by "gen/generator.go". //template:begin findObjectsToBeReplaced

// End of section. //template:end findObjectsToBeReplaced

// Section below is generated&owned by "gen/generator.go". //template:begin clearItemIds

// End of section. //template:end clearItemIds
