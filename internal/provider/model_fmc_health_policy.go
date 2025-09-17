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
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type HealthPolicy struct {
	Id                          types.String                `tfsdk:"id"`
	Domain                      types.String                `tfsdk:"domain"`
	Name                        types.String                `tfsdk:"name"`
	Description                 types.String                `tfsdk:"description"`
	Type                        types.String                `tfsdk:"type"`
	PolicyType                  types.String                `tfsdk:"policy_type"`
	IsDefaultPolicy             types.Bool                  `tfsdk:"is_default_policy"`
	HealthModules               []HealthPolicyHealthModules `tfsdk:"health_modules"`
	HealthModuleRunTimeInterval types.Int64                 `tfsdk:"health_module_run_time_interval"`
	MetricCollectionInterval    types.Int64                 `tfsdk:"metric_collection_interval"`
}

type HealthPolicyHealthModules struct {
	ModuleId          types.String                               `tfsdk:"module_id"`
	Enabled           types.Bool                                 `tfsdk:"enabled"`
	Type              types.String                               `tfsdk:"type"`
	AlertSeverity     types.String                               `tfsdk:"alert_severity"`
	CriticalThreshold types.Int64                                `tfsdk:"critical_threshold"`
	WarningThreshold  types.Int64                                `tfsdk:"warning_threshold"`
	CustomThreshold   []HealthPolicyHealthModulesCustomThreshold `tfsdk:"custom_threshold"`
	AlertConfig       []HealthPolicyHealthModulesAlertConfig     `tfsdk:"alert_config"`
}

type HealthPolicyHealthModulesCustomThreshold struct {
	Type  types.String `tfsdk:"type"`
	Value types.Int64  `tfsdk:"value"`
}
type HealthPolicyHealthModulesAlertConfig struct {
	Name       types.String                                     `tfsdk:"name"`
	Enabled    types.Bool                                       `tfsdk:"enabled"`
	Thresholds []HealthPolicyHealthModulesAlertConfigThresholds `tfsdk:"thresholds"`
}

type HealthPolicyHealthModulesAlertConfigThresholds struct {
	Type  types.String `tfsdk:"type"`
	Value types.Int64  `tfsdk:"value"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data HealthPolicy) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/healthpolicies"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data HealthPolicy) toBody(ctx context.Context, state HealthPolicy) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.PolicyType.IsNull() {
		body, _ = sjson.Set(body, "policyType", data.PolicyType.ValueString())
	}
	if !data.IsDefaultPolicy.IsNull() {
		body, _ = sjson.Set(body, "defaultPolicy", data.IsDefaultPolicy.ValueBool())
	}
	if len(data.HealthModules) > 0 {
		body, _ = sjson.Set(body, "healthModules", []any{})
		for _, item := range data.HealthModules {
			itemBody := ""
			if !item.ModuleId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "moduleId", item.ModuleId.ValueString())
			}
			if !item.Enabled.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enabled", item.Enabled.ValueBool())
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
			}
			if !item.AlertSeverity.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "severity", item.AlertSeverity.ValueString())
			}
			if !item.CriticalThreshold.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "criticalThreshold", item.CriticalThreshold.ValueInt64())
			}
			if !item.WarningThreshold.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "warningThreshold", item.WarningThreshold.ValueInt64())
			}
			if len(item.CustomThreshold) > 0 {
				itemBody, _ = sjson.Set(itemBody, "customThreshold", []any{})
				for _, childItem := range item.CustomThreshold {
					itemChildBody := ""
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					if !childItem.Value.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Value.ValueInt64())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "customThreshold.-1", itemChildBody)
				}
			}
			if len(item.AlertConfig) > 0 {
				itemBody, _ = sjson.Set(itemBody, "alertConfig", []any{})
				for _, childItem := range item.AlertConfig {
					itemChildBody := ""
					if !childItem.Name.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "name", childItem.Name.ValueString())
					}
					if !childItem.Enabled.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "enabled", childItem.Enabled.ValueBool())
					}
					if len(childItem.Thresholds) > 0 {
						itemChildBody, _ = sjson.Set(itemChildBody, "thresholds", []any{})
						for _, childChildItem := range childItem.Thresholds {
							itemChildChildBody := ""
							if !childChildItem.Type.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "type", childChildItem.Type.ValueString())
							}
							if !childChildItem.Value.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value", childChildItem.Value.ValueInt64())
							}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "thresholds.-1", itemChildChildBody)
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "alertConfig.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "healthModules.-1", itemBody)
		}
	}
	if !data.HealthModuleRunTimeInterval.IsNull() {
		body, _ = sjson.Set(body, "setting.alertEvaluationInterval", data.HealthModuleRunTimeInterval.ValueInt64())
	}
	if !data.MetricCollectionInterval.IsNull() {
		body, _ = sjson.Set(body, "setting.dataCollectionInterval", data.MetricCollectionInterval.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *HealthPolicy) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("policyType"); value.Exists() {
		data.PolicyType = types.StringValue(value.String())
	} else {
		data.PolicyType = types.StringNull()
	}
	if value := res.Get("defaultPolicy"); value.Exists() {
		data.IsDefaultPolicy = types.BoolValue(value.Bool())
	} else {
		data.IsDefaultPolicy = types.BoolNull()
	}
	if value := res.Get("healthModules"); value.Exists() {
		data.HealthModules = make([]HealthPolicyHealthModules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := HealthPolicyHealthModules{}
			if value := res.Get("moduleId"); value.Exists() {
				data.ModuleId = types.StringValue(value.String())
			} else {
				data.ModuleId = types.StringNull()
			}
			if value := res.Get("enabled"); value.Exists() {
				data.Enabled = types.BoolValue(value.Bool())
			} else {
				data.Enabled = types.BoolNull()
			}
			if value := res.Get("type"); value.Exists() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			if value := res.Get("severity"); value.Exists() {
				data.AlertSeverity = types.StringValue(value.String())
			} else {
				data.AlertSeverity = types.StringNull()
			}
			if value := res.Get("criticalThreshold"); value.Exists() {
				data.CriticalThreshold = types.Int64Value(value.Int())
			} else {
				data.CriticalThreshold = types.Int64Null()
			}
			if value := res.Get("warningThreshold"); value.Exists() {
				data.WarningThreshold = types.Int64Value(value.Int())
			} else {
				data.WarningThreshold = types.Int64Null()
			}
			if value := res.Get("customThreshold"); value.Exists() {
				data.CustomThreshold = make([]HealthPolicyHealthModulesCustomThreshold, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := HealthPolicyHealthModulesCustomThreshold{}
					if value := res.Get("type"); value.Exists() {
						data.Type = types.StringValue(value.String())
					} else {
						data.Type = types.StringNull()
					}
					if value := res.Get("value"); value.Exists() {
						data.Value = types.Int64Value(value.Int())
					} else {
						data.Value = types.Int64Null()
					}
					(*parent).CustomThreshold = append((*parent).CustomThreshold, data)
					return true
				})
			}
			if value := res.Get("alertConfig"); value.Exists() {
				data.AlertConfig = make([]HealthPolicyHealthModulesAlertConfig, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := HealthPolicyHealthModulesAlertConfig{}
					if value := res.Get("name"); value.Exists() {
						data.Name = types.StringValue(value.String())
					} else {
						data.Name = types.StringNull()
					}
					if value := res.Get("enabled"); value.Exists() {
						data.Enabled = types.BoolValue(value.Bool())
					} else {
						data.Enabled = types.BoolNull()
					}
					if value := res.Get("thresholds"); value.Exists() {
						data.Thresholds = make([]HealthPolicyHealthModulesAlertConfigThresholds, 0)
						value.ForEach(func(k, res gjson.Result) bool {
							parent := &data
							data := HealthPolicyHealthModulesAlertConfigThresholds{}
							if value := res.Get("type"); value.Exists() {
								data.Type = types.StringValue(value.String())
							} else {
								data.Type = types.StringNull()
							}
							if value := res.Get("value"); value.Exists() {
								data.Value = types.Int64Value(value.Int())
							} else {
								data.Value = types.Int64Null()
							}
							(*parent).Thresholds = append((*parent).Thresholds, data)
							return true
						})
					}
					(*parent).AlertConfig = append((*parent).AlertConfig, data)
					return true
				})
			}
			(*parent).HealthModules = append((*parent).HealthModules, data)
			return true
		})
	}
	if value := res.Get("setting.alertEvaluationInterval"); value.Exists() {
		data.HealthModuleRunTimeInterval = types.Int64Value(value.Int())
	} else {
		data.HealthModuleRunTimeInterval = types.Int64Value(5)
	}
	if value := res.Get("setting.dataCollectionInterval"); value.Exists() {
		data.MetricCollectionInterval = types.Int64Value(value.Int())
	} else {
		data.MetricCollectionInterval = types.Int64Value(1)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *HealthPolicy) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("policyType"); value.Exists() && !data.PolicyType.IsNull() {
		data.PolicyType = types.StringValue(value.String())
	} else {
		data.PolicyType = types.StringNull()
	}
	if value := res.Get("defaultPolicy"); value.Exists() && !data.IsDefaultPolicy.IsNull() {
		data.IsDefaultPolicy = types.BoolValue(value.Bool())
	} else {
		data.IsDefaultPolicy = types.BoolNull()
	}
	for i := 0; i < len(data.HealthModules); i++ {
		keys := [...]string{"moduleId"}
		keyValues := [...]string{data.HealthModules[i].ModuleId.ValueString()}

		parent := &data
		data := (*parent).HealthModules[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("healthModules").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing HealthModules[%d] = %+v",
				i,
				(*parent).HealthModules[i],
			))
			(*parent).HealthModules = slices.Delete((*parent).HealthModules, i, i+1)
			i--

			continue
		}
		if value := res.Get("moduleId"); value.Exists() && !data.ModuleId.IsNull() {
			data.ModuleId = types.StringValue(value.String())
		} else {
			data.ModuleId = types.StringNull()
		}
		if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
			data.Enabled = types.BoolValue(value.Bool())
		} else {
			data.Enabled = types.BoolNull()
		}
		if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		if value := res.Get("severity"); value.Exists() && !data.AlertSeverity.IsNull() {
			data.AlertSeverity = types.StringValue(value.String())
		} else {
			data.AlertSeverity = types.StringNull()
		}
		if value := res.Get("criticalThreshold"); value.Exists() && !data.CriticalThreshold.IsNull() {
			data.CriticalThreshold = types.Int64Value(value.Int())
		} else {
			data.CriticalThreshold = types.Int64Null()
		}
		if value := res.Get("warningThreshold"); value.Exists() && !data.WarningThreshold.IsNull() {
			data.WarningThreshold = types.Int64Value(value.Int())
		} else {
			data.WarningThreshold = types.Int64Null()
		}
		for i := 0; i < len(data.CustomThreshold); i++ {
			keys := [...]string{"type", "value"}
			keyValues := [...]string{data.CustomThreshold[i].Type.ValueString(), strconv.FormatInt(data.CustomThreshold[i].Value.ValueInt64(), 10)}

			parent := &data
			data := (*parent).CustomThreshold[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("customThreshold").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing CustomThreshold[%d] = %+v",
					i,
					(*parent).CustomThreshold[i],
				))
				(*parent).CustomThreshold = slices.Delete((*parent).CustomThreshold, i, i+1)
				i--

				continue
			}
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			if value := res.Get("value"); value.Exists() && !data.Value.IsNull() {
				data.Value = types.Int64Value(value.Int())
			} else {
				data.Value = types.Int64Null()
			}
			(*parent).CustomThreshold[i] = data
		}
		for i := 0; i < len(data.AlertConfig); i++ {
			keys := [...]string{"name", "enabled"}
			keyValues := [...]string{data.AlertConfig[i].Name.ValueString(), strconv.FormatBool(data.AlertConfig[i].Enabled.ValueBool())}

			parent := &data
			data := (*parent).AlertConfig[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("alertConfig").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing AlertConfig[%d] = %+v",
					i,
					(*parent).AlertConfig[i],
				))
				(*parent).AlertConfig = slices.Delete((*parent).AlertConfig, i, i+1)
				i--

				continue
			}
			if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
				data.Enabled = types.BoolValue(value.Bool())
			} else {
				data.Enabled = types.BoolNull()
			}
			for i := 0; i < len(data.Thresholds); i++ {
				keys := [...]string{"type", "value"}
				keyValues := [...]string{data.Thresholds[i].Type.ValueString(), strconv.FormatInt(data.Thresholds[i].Value.ValueInt64(), 10)}

				parent := &data
				data := (*parent).Thresholds[i]
				parentRes := &res
				var res gjson.Result

				parentRes.Get("thresholds").ForEach(
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
					tflog.Debug(ctx, fmt.Sprintf("removing Thresholds[%d] = %+v",
						i,
						(*parent).Thresholds[i],
					))
					(*parent).Thresholds = slices.Delete((*parent).Thresholds, i, i+1)
					i--

					continue
				}
				if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
					data.Type = types.StringValue(value.String())
				} else {
					data.Type = types.StringNull()
				}
				if value := res.Get("value"); value.Exists() && !data.Value.IsNull() {
					data.Value = types.Int64Value(value.Int())
				} else {
					data.Value = types.Int64Null()
				}
				(*parent).Thresholds[i] = data
			}
			(*parent).AlertConfig[i] = data
		}
		(*parent).HealthModules[i] = data
	}
	if value := res.Get("setting.alertEvaluationInterval"); value.Exists() && !data.HealthModuleRunTimeInterval.IsNull() {
		data.HealthModuleRunTimeInterval = types.Int64Value(value.Int())
	} else if data.HealthModuleRunTimeInterval.ValueInt64() != 5 {
		data.HealthModuleRunTimeInterval = types.Int64Null()
	}
	if value := res.Get("setting.dataCollectionInterval"); value.Exists() && !data.MetricCollectionInterval.IsNull() {
		data.MetricCollectionInterval = types.Int64Value(value.Int())
	} else if data.MetricCollectionInterval.ValueInt64() != 1 {
		data.MetricCollectionInterval = types.Int64Null()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *HealthPolicy) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyPutDelete

// End of section. //template:end toBodyPutDelete

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBody

// End of section. //template:end adjustBody

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBodyBulk

// End of section. //template:end adjustBodyBulk
