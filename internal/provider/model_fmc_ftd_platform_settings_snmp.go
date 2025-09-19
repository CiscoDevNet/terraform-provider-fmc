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
	"slices"

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type FTDPlatformSettingsSNMP struct {
	Id                             types.String                                 `tfsdk:"id"`
	Domain                         types.String                                 `tfsdk:"domain"`
	FtdPlatformSettingsId          types.String                                 `tfsdk:"ftd_platform_settings_id"`
	Type                           types.String                                 `tfsdk:"type"`
	EnableSnmpServer               types.Bool                                   `tfsdk:"enable_snmp_server"`
	ReadCommunityString            types.String                                 `tfsdk:"read_community_string"`
	SystemAdministratorName        types.String                                 `tfsdk:"system_administrator_name"`
	Location                       types.String                                 `tfsdk:"location"`
	SnmpServerPort                 types.Int64                                  `tfsdk:"snmp_server_port"`
	SnmpManagementHosts            []FTDPlatformSettingsSNMPSnmpManagementHosts `tfsdk:"snmp_management_hosts"`
	Snmpv3Users                    []FTDPlatformSettingsSNMPSnmpv3Users         `tfsdk:"snmpv3_users"`
	TrapSyslog                     types.Bool                                   `tfsdk:"trap_syslog"`
	TrapAuthentication             types.Bool                                   `tfsdk:"trap_authentication"`
	TrapLinkUp                     types.Bool                                   `tfsdk:"trap_link_up"`
	TrapLinkDown                   types.Bool                                   `tfsdk:"trap_link_down"`
	TrapColdStart                  types.Bool                                   `tfsdk:"trap_cold_start"`
	TrapWarmStart                  types.Bool                                   `tfsdk:"trap_warm_start"`
	TrapFieldReplacementUnitInsert types.Bool                                   `tfsdk:"trap_field_replacement_unit_insert"`
	TrapFieldReplacementUnitDelete types.Bool                                   `tfsdk:"trap_field_replacement_unit_delete"`
	TrapConfigurationChange        types.Bool                                   `tfsdk:"trap_configuration_change"`
	TrapConnectionLimitReached     types.Bool                                   `tfsdk:"trap_connection_limit_reached"`
	TrapNatPacketDiscard           types.Bool                                   `tfsdk:"trap_nat_packet_discard"`
	TrapCpuRisingThreshold         types.Bool                                   `tfsdk:"trap_cpu_rising_threshold"`
	TrapCpuRisingThresholdValue    types.Int64                                  `tfsdk:"trap_cpu_rising_threshold_value"`
	TrapCpuRisingThresholdInterval types.Int64                                  `tfsdk:"trap_cpu_rising_threshold_interval"`
	TrapMemoryRisingThreshold      types.Bool                                   `tfsdk:"trap_memory_rising_threshold"`
	TrapMemoryRisingThresholdValue types.Int64                                  `tfsdk:"trap_memory_rising_threshold_value"`
	TrapFailover                   types.Bool                                   `tfsdk:"trap_failover"`
	TrapCluster                    types.Bool                                   `tfsdk:"trap_cluster"`
	TrapPeerFlap                   types.Bool                                   `tfsdk:"trap_peer_flap"`
}

type FTDPlatformSettingsSNMPSnmpManagementHosts struct {
	IpObjectId             types.String                                                 `tfsdk:"ip_object_id"`
	SnmpVersion            types.String                                                 `tfsdk:"snmp_version"`
	Username               types.String                                                 `tfsdk:"username"`
	ReadCommunityString    types.String                                                 `tfsdk:"read_community_string"`
	Poll                   types.Bool                                                   `tfsdk:"poll"`
	Trap                   types.Bool                                                   `tfsdk:"trap"`
	TrapPort               types.Int64                                                  `tfsdk:"trap_port"`
	UseManagementInterface types.Bool                                                   `tfsdk:"use_management_interface"`
	InterfaceLiterals      types.Set                                                    `tfsdk:"interface_literals"`
	InterfaceObjects       []FTDPlatformSettingsSNMPSnmpManagementHostsInterfaceObjects `tfsdk:"interface_objects"`
}

type FTDPlatformSettingsSNMPSnmpv3Users struct {
	SecurityLevel               types.String `tfsdk:"security_level"`
	Username                    types.String `tfsdk:"username"`
	PasswordType                types.String `tfsdk:"password_type"`
	AuthenticationAlgorithmType types.String `tfsdk:"authentication_algorithm_type"`
	AuthenticationPassword      types.String `tfsdk:"authentication_password"`
	EncryptionType              types.String `tfsdk:"encryption_type"`
	EncryptionPassword          types.String `tfsdk:"encryption_password"`
}

type FTDPlatformSettingsSNMPSnmpManagementHostsInterfaceObjects struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
	Name types.String `tfsdk:"name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionFTDPlatformSettingsSNMP = version.Must(version.NewVersion("7.7"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data FTDPlatformSettingsSNMP) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/ftdplatformsettingspolicies/%v/snmpsettings", url.QueryEscape(data.FtdPlatformSettingsId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data FTDPlatformSettingsSNMP) toBody(ctx context.Context, state FTDPlatformSettingsSNMP) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.EnableSnmpServer.IsNull() {
		body, _ = sjson.Set(body, "enableSNMPServers", data.EnableSnmpServer.ValueBool())
	}
	if !data.ReadCommunityString.IsNull() {
		body, _ = sjson.Set(body, "communityString", data.ReadCommunityString.ValueString())
	}
	if !data.SystemAdministratorName.IsNull() {
		body, _ = sjson.Set(body, "sysAdminName", data.SystemAdministratorName.ValueString())
	}
	if !data.Location.IsNull() {
		body, _ = sjson.Set(body, "location", data.Location.ValueString())
	}
	if !data.SnmpServerPort.IsNull() {
		body, _ = sjson.Set(body, "port", data.SnmpServerPort.ValueInt64())
	}
	if len(data.SnmpManagementHosts) > 0 {
		body, _ = sjson.Set(body, "snmpMgmtHosts", []any{})
		for _, item := range data.SnmpManagementHosts {
			itemBody := ""
			if !item.IpObjectId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipAddress.id", item.IpObjectId.ValueString())
			}
			if !item.SnmpVersion.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "snmpVersion", item.SnmpVersion.ValueString())
			}
			if !item.Username.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "userName", item.Username.ValueString())
			}
			if !item.ReadCommunityString.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "communityString", item.ReadCommunityString.ValueString())
			}
			if !item.Poll.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "poll", item.Poll.ValueBool())
			}
			if !item.Trap.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "trap", item.Trap.ValueBool())
			}
			if !item.TrapPort.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "port", item.TrapPort.ValueInt64())
			}
			if !item.UseManagementInterface.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ifSnmpHostMgmt", item.UseManagementInterface.ValueBool())
			}
			if !item.InterfaceLiterals.IsNull() {
				var values []string
				item.InterfaceLiterals.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "interfaces.literals", values)
			}
			if len(item.InterfaceObjects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "interfaces.objects", []any{})
				for _, childItem := range item.InterfaceObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					if !childItem.Name.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "name", childItem.Name.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "interfaces.objects.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "snmpMgmtHosts.-1", itemBody)
		}
	}
	if len(data.Snmpv3Users) > 0 {
		body, _ = sjson.Set(body, "snmpv3Users", []any{})
		for _, item := range data.Snmpv3Users {
			itemBody := ""
			if !item.SecurityLevel.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "securityLevel", item.SecurityLevel.ValueString())
			}
			if !item.Username.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "userName", item.Username.ValueString())
			}
			if !item.PasswordType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "encryptionPasswordType", item.PasswordType.ValueString())
			}
			if !item.AuthenticationAlgorithmType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "authAlgoType", item.AuthenticationAlgorithmType.ValueString())
			}
			if !item.AuthenticationPassword.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "authenticationPassword", item.AuthenticationPassword.ValueString())
			}
			if !item.EncryptionType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "encryptionType", item.EncryptionType.ValueString())
			}
			if !item.EncryptionPassword.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "encryptionPassword", item.EncryptionPassword.ValueString())
			}
			body, _ = sjson.SetRaw(body, "snmpv3Users.-1", itemBody)
		}
	}
	if !data.TrapSyslog.IsNull() {
		body, _ = sjson.Set(body, "snmpTrap.syslog", data.TrapSyslog.ValueBool())
	}
	if !data.TrapAuthentication.IsNull() {
		body, _ = sjson.Set(body, "snmpTrap.authentication", data.TrapAuthentication.ValueBool())
	}
	if !data.TrapLinkUp.IsNull() {
		body, _ = sjson.Set(body, "snmpTrap.linkUp", data.TrapLinkUp.ValueBool())
	}
	if !data.TrapLinkDown.IsNull() {
		body, _ = sjson.Set(body, "snmpTrap.linkDown", data.TrapLinkDown.ValueBool())
	}
	if !data.TrapColdStart.IsNull() {
		body, _ = sjson.Set(body, "snmpTrap.coldStart", data.TrapColdStart.ValueBool())
	}
	if !data.TrapWarmStart.IsNull() {
		body, _ = sjson.Set(body, "snmpTrap.warmStart", data.TrapWarmStart.ValueBool())
	}
	if !data.TrapFieldReplacementUnitInsert.IsNull() {
		body, _ = sjson.Set(body, "snmpTrap.fruinsert", data.TrapFieldReplacementUnitInsert.ValueBool())
	}
	if !data.TrapFieldReplacementUnitDelete.IsNull() {
		body, _ = sjson.Set(body, "snmpTrap.fruremove", data.TrapFieldReplacementUnitDelete.ValueBool())
	}
	if !data.TrapConfigurationChange.IsNull() {
		body, _ = sjson.Set(body, "snmpTrap.entityConfigChange", data.TrapConfigurationChange.ValueBool())
	}
	if !data.TrapConnectionLimitReached.IsNull() {
		body, _ = sjson.Set(body, "snmpTrap.connLimitReached", data.TrapConnectionLimitReached.ValueBool())
	}
	if !data.TrapNatPacketDiscard.IsNull() {
		body, _ = sjson.Set(body, "snmpTrap.natPacketDiscard", data.TrapNatPacketDiscard.ValueBool())
	}
	if !data.TrapCpuRisingThreshold.IsNull() {
		body, _ = sjson.Set(body, "snmpTrap.cputhreshold", data.TrapCpuRisingThreshold.ValueBool())
	}
	if !data.TrapCpuRisingThresholdValue.IsNull() {
		body, _ = sjson.Set(body, "snmpTrap.cputhresholdValueInPercent", data.TrapCpuRisingThresholdValue.ValueInt64())
	}
	if !data.TrapCpuRisingThresholdInterval.IsNull() {
		body, _ = sjson.Set(body, "snmpTrap.cputhresholdIntervalInMins", data.TrapCpuRisingThresholdInterval.ValueInt64())
	}
	if !data.TrapMemoryRisingThreshold.IsNull() {
		body, _ = sjson.Set(body, "snmpTrap.memoryThreshold", data.TrapMemoryRisingThreshold.ValueBool())
	}
	if !data.TrapMemoryRisingThresholdValue.IsNull() {
		body, _ = sjson.Set(body, "snmpTrap.memThresholdValueInPercent", data.TrapMemoryRisingThresholdValue.ValueInt64())
	}
	if !data.TrapFailover.IsNull() {
		body, _ = sjson.Set(body, "snmpTrap.failoverState", data.TrapFailover.ValueBool())
	}
	if !data.TrapCluster.IsNull() {
		body, _ = sjson.Set(body, "snmpTrap.clusterState", data.TrapCluster.ValueBool())
	}
	if !data.TrapPeerFlap.IsNull() {
		body, _ = sjson.Set(body, "snmpTrap.peerFlap", data.TrapPeerFlap.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *FTDPlatformSettingsSNMP) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("enableSNMPServers"); value.Exists() {
		data.EnableSnmpServer = types.BoolValue(value.Bool())
	} else {
		data.EnableSnmpServer = types.BoolNull()
	}
	if value := res.Get("sysAdminName"); value.Exists() {
		data.SystemAdministratorName = types.StringValue(value.String())
	} else {
		data.SystemAdministratorName = types.StringNull()
	}
	if value := res.Get("location"); value.Exists() {
		data.Location = types.StringValue(value.String())
	} else {
		data.Location = types.StringNull()
	}
	if value := res.Get("port"); value.Exists() {
		data.SnmpServerPort = types.Int64Value(value.Int())
	} else {
		data.SnmpServerPort = types.Int64Value(161)
	}
	if value := res.Get("snmpMgmtHosts"); value.Exists() {
		data.SnmpManagementHosts = make([]FTDPlatformSettingsSNMPSnmpManagementHosts, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := FTDPlatformSettingsSNMPSnmpManagementHosts{}
			if value := res.Get("ipAddress.id"); value.Exists() {
				data.IpObjectId = types.StringValue(value.String())
			} else {
				data.IpObjectId = types.StringNull()
			}
			if value := res.Get("snmpVersion"); value.Exists() {
				data.SnmpVersion = types.StringValue(value.String())
			} else {
				data.SnmpVersion = types.StringNull()
			}
			if value := res.Get("userName"); value.Exists() {
				data.Username = types.StringValue(value.String())
			} else {
				data.Username = types.StringNull()
			}
			if value := res.Get("poll"); value.Exists() {
				data.Poll = types.BoolValue(value.Bool())
			} else {
				data.Poll = types.BoolNull()
			}
			if value := res.Get("trap"); value.Exists() {
				data.Trap = types.BoolValue(value.Bool())
			} else {
				data.Trap = types.BoolNull()
			}
			if value := res.Get("port"); value.Exists() {
				data.TrapPort = types.Int64Value(value.Int())
			} else {
				data.TrapPort = types.Int64Value(162)
			}
			if value := res.Get("ifSnmpHostMgmt"); value.Exists() {
				data.UseManagementInterface = types.BoolValue(value.Bool())
			} else {
				data.UseManagementInterface = types.BoolNull()
			}
			if value := res.Get("interfaces.literals"); value.Exists() {
				data.InterfaceLiterals = helpers.GetStringSet(value.Array())
			} else {
				data.InterfaceLiterals = types.SetNull(types.StringType)
			}
			if value := res.Get("interfaces.objects"); value.Exists() {
				data.InterfaceObjects = make([]FTDPlatformSettingsSNMPSnmpManagementHostsInterfaceObjects, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := FTDPlatformSettingsSNMPSnmpManagementHostsInterfaceObjects{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					if value := res.Get("type"); value.Exists() {
						data.Type = types.StringValue(value.String())
					} else {
						data.Type = types.StringNull()
					}
					if value := res.Get("name"); value.Exists() {
						data.Name = types.StringValue(value.String())
					} else {
						data.Name = types.StringNull()
					}
					(*parent).InterfaceObjects = append((*parent).InterfaceObjects, data)
					return true
				})
			}
			(*parent).SnmpManagementHosts = append((*parent).SnmpManagementHosts, data)
			return true
		})
	}
	if value := res.Get("snmpv3Users"); value.Exists() {
		data.Snmpv3Users = make([]FTDPlatformSettingsSNMPSnmpv3Users, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := FTDPlatformSettingsSNMPSnmpv3Users{}
			if value := res.Get("securityLevel"); value.Exists() {
				data.SecurityLevel = types.StringValue(value.String())
			} else {
				data.SecurityLevel = types.StringNull()
			}
			if value := res.Get("userName"); value.Exists() {
				data.Username = types.StringValue(value.String())
			} else {
				data.Username = types.StringNull()
			}
			if value := res.Get("encryptionPasswordType"); value.Exists() {
				data.PasswordType = types.StringValue(value.String())
			} else {
				data.PasswordType = types.StringNull()
			}
			if value := res.Get("authAlgoType"); value.Exists() {
				data.AuthenticationAlgorithmType = types.StringValue(value.String())
			} else {
				data.AuthenticationAlgorithmType = types.StringNull()
			}
			if value := res.Get("encryptionType"); value.Exists() {
				data.EncryptionType = types.StringValue(value.String())
			} else {
				data.EncryptionType = types.StringNull()
			}
			(*parent).Snmpv3Users = append((*parent).Snmpv3Users, data)
			return true
		})
	}
	if value := res.Get("snmpTrap.syslog"); value.Exists() {
		data.TrapSyslog = types.BoolValue(value.Bool())
	} else {
		data.TrapSyslog = types.BoolNull()
	}
	if value := res.Get("snmpTrap.authentication"); value.Exists() {
		data.TrapAuthentication = types.BoolValue(value.Bool())
	} else {
		data.TrapAuthentication = types.BoolNull()
	}
	if value := res.Get("snmpTrap.linkUp"); value.Exists() {
		data.TrapLinkUp = types.BoolValue(value.Bool())
	} else {
		data.TrapLinkUp = types.BoolNull()
	}
	if value := res.Get("snmpTrap.linkDown"); value.Exists() {
		data.TrapLinkDown = types.BoolValue(value.Bool())
	} else {
		data.TrapLinkDown = types.BoolNull()
	}
	if value := res.Get("snmpTrap.coldStart"); value.Exists() {
		data.TrapColdStart = types.BoolValue(value.Bool())
	} else {
		data.TrapColdStart = types.BoolNull()
	}
	if value := res.Get("snmpTrap.warmStart"); value.Exists() {
		data.TrapWarmStart = types.BoolValue(value.Bool())
	} else {
		data.TrapWarmStart = types.BoolNull()
	}
	if value := res.Get("snmpTrap.fruinsert"); value.Exists() {
		data.TrapFieldReplacementUnitInsert = types.BoolValue(value.Bool())
	} else {
		data.TrapFieldReplacementUnitInsert = types.BoolNull()
	}
	if value := res.Get("snmpTrap.fruremove"); value.Exists() {
		data.TrapFieldReplacementUnitDelete = types.BoolValue(value.Bool())
	} else {
		data.TrapFieldReplacementUnitDelete = types.BoolNull()
	}
	if value := res.Get("snmpTrap.entityConfigChange"); value.Exists() {
		data.TrapConfigurationChange = types.BoolValue(value.Bool())
	} else {
		data.TrapConfigurationChange = types.BoolNull()
	}
	if value := res.Get("snmpTrap.connLimitReached"); value.Exists() {
		data.TrapConnectionLimitReached = types.BoolValue(value.Bool())
	} else {
		data.TrapConnectionLimitReached = types.BoolNull()
	}
	if value := res.Get("snmpTrap.natPacketDiscard"); value.Exists() {
		data.TrapNatPacketDiscard = types.BoolValue(value.Bool())
	} else {
		data.TrapNatPacketDiscard = types.BoolNull()
	}
	if value := res.Get("snmpTrap.cputhreshold"); value.Exists() {
		data.TrapCpuRisingThreshold = types.BoolValue(value.Bool())
	} else {
		data.TrapCpuRisingThreshold = types.BoolNull()
	}
	if value := res.Get("snmpTrap.cputhresholdValueInPercent"); value.Exists() {
		data.TrapCpuRisingThresholdValue = types.Int64Value(value.Int())
	} else {
		data.TrapCpuRisingThresholdValue = types.Int64Null()
	}
	if value := res.Get("snmpTrap.cputhresholdIntervalInMins"); value.Exists() {
		data.TrapCpuRisingThresholdInterval = types.Int64Value(value.Int())
	} else {
		data.TrapCpuRisingThresholdInterval = types.Int64Null()
	}
	if value := res.Get("snmpTrap.memoryThreshold"); value.Exists() {
		data.TrapMemoryRisingThreshold = types.BoolValue(value.Bool())
	} else {
		data.TrapMemoryRisingThreshold = types.BoolNull()
	}
	if value := res.Get("snmpTrap.memThresholdValueInPercent"); value.Exists() {
		data.TrapMemoryRisingThresholdValue = types.Int64Value(value.Int())
	} else {
		data.TrapMemoryRisingThresholdValue = types.Int64Null()
	}
	if value := res.Get("snmpTrap.failoverState"); value.Exists() {
		data.TrapFailover = types.BoolValue(value.Bool())
	} else {
		data.TrapFailover = types.BoolNull()
	}
	if value := res.Get("snmpTrap.clusterState"); value.Exists() {
		data.TrapCluster = types.BoolValue(value.Bool())
	} else {
		data.TrapCluster = types.BoolNull()
	}
	if value := res.Get("snmpTrap.peerFlap"); value.Exists() {
		data.TrapPeerFlap = types.BoolValue(value.Bool())
	} else {
		data.TrapPeerFlap = types.BoolNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *FTDPlatformSettingsSNMP) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("enableSNMPServers"); value.Exists() && !data.EnableSnmpServer.IsNull() {
		data.EnableSnmpServer = types.BoolValue(value.Bool())
	} else {
		data.EnableSnmpServer = types.BoolNull()
	}
	if value := res.Get("sysAdminName"); value.Exists() && !data.SystemAdministratorName.IsNull() {
		data.SystemAdministratorName = types.StringValue(value.String())
	} else {
		data.SystemAdministratorName = types.StringNull()
	}
	if value := res.Get("location"); value.Exists() && !data.Location.IsNull() {
		data.Location = types.StringValue(value.String())
	} else {
		data.Location = types.StringNull()
	}
	if value := res.Get("port"); value.Exists() && !data.SnmpServerPort.IsNull() {
		data.SnmpServerPort = types.Int64Value(value.Int())
	} else if data.SnmpServerPort.ValueInt64() != 161 {
		data.SnmpServerPort = types.Int64Null()
	}
	for i := 0; i < len(data.SnmpManagementHosts); i++ {
		keys := [...]string{"ipAddress.id"}
		keyValues := [...]string{data.SnmpManagementHosts[i].IpObjectId.ValueString()}

		parent := &data
		data := (*parent).SnmpManagementHosts[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("snmpMgmtHosts").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing SnmpManagementHosts[%d] = %+v",
				i,
				(*parent).SnmpManagementHosts[i],
			))
			(*parent).SnmpManagementHosts = slices.Delete((*parent).SnmpManagementHosts, i, i+1)
			i--

			continue
		}
		if value := res.Get("ipAddress.id"); value.Exists() && !data.IpObjectId.IsNull() {
			data.IpObjectId = types.StringValue(value.String())
		} else {
			data.IpObjectId = types.StringNull()
		}
		if value := res.Get("snmpVersion"); value.Exists() && !data.SnmpVersion.IsNull() {
			data.SnmpVersion = types.StringValue(value.String())
		} else {
			data.SnmpVersion = types.StringNull()
		}
		if value := res.Get("userName"); value.Exists() && !data.Username.IsNull() {
			data.Username = types.StringValue(value.String())
		} else {
			data.Username = types.StringNull()
		}
		if value := res.Get("poll"); value.Exists() && !data.Poll.IsNull() {
			data.Poll = types.BoolValue(value.Bool())
		} else {
			data.Poll = types.BoolNull()
		}
		if value := res.Get("trap"); value.Exists() && !data.Trap.IsNull() {
			data.Trap = types.BoolValue(value.Bool())
		} else {
			data.Trap = types.BoolNull()
		}
		if value := res.Get("port"); value.Exists() && !data.TrapPort.IsNull() {
			data.TrapPort = types.Int64Value(value.Int())
		} else if data.TrapPort.ValueInt64() != 162 {
			data.TrapPort = types.Int64Null()
		}
		if value := res.Get("ifSnmpHostMgmt"); value.Exists() && !data.UseManagementInterface.IsNull() {
			data.UseManagementInterface = types.BoolValue(value.Bool())
		} else {
			data.UseManagementInterface = types.BoolNull()
		}
		if value := res.Get("interfaces.literals"); value.Exists() && !data.InterfaceLiterals.IsNull() {
			data.InterfaceLiterals = helpers.GetStringSet(value.Array())
		} else {
			data.InterfaceLiterals = types.SetNull(types.StringType)
		}
		for i := 0; i < len(data.InterfaceObjects); i++ {
			keys := [...]string{"id", "type", "name"}
			keyValues := [...]string{data.InterfaceObjects[i].Id.ValueString(), data.InterfaceObjects[i].Type.ValueString(), data.InterfaceObjects[i].Name.ValueString()}

			parent := &data
			data := (*parent).InterfaceObjects[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("interfaces.objects").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing InterfaceObjects[%d] = %+v",
					i,
					(*parent).InterfaceObjects[i],
				))
				(*parent).InterfaceObjects = slices.Delete((*parent).InterfaceObjects, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			(*parent).InterfaceObjects[i] = data
		}
		(*parent).SnmpManagementHosts[i] = data
	}
	for i := 0; i < len(data.Snmpv3Users); i++ {
		keys := [...]string{"userName"}
		keyValues := [...]string{data.Snmpv3Users[i].Username.ValueString()}

		parent := &data
		data := (*parent).Snmpv3Users[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("snmpv3Users").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Snmpv3Users[%d] = %+v",
				i,
				(*parent).Snmpv3Users[i],
			))
			(*parent).Snmpv3Users = slices.Delete((*parent).Snmpv3Users, i, i+1)
			i--

			continue
		}
		if value := res.Get("securityLevel"); value.Exists() && !data.SecurityLevel.IsNull() {
			data.SecurityLevel = types.StringValue(value.String())
		} else {
			data.SecurityLevel = types.StringNull()
		}
		if value := res.Get("userName"); value.Exists() && !data.Username.IsNull() {
			data.Username = types.StringValue(value.String())
		} else {
			data.Username = types.StringNull()
		}
		if value := res.Get("encryptionPasswordType"); value.Exists() && !data.PasswordType.IsNull() {
			data.PasswordType = types.StringValue(value.String())
		} else {
			data.PasswordType = types.StringNull()
		}
		if value := res.Get("authAlgoType"); value.Exists() && !data.AuthenticationAlgorithmType.IsNull() {
			data.AuthenticationAlgorithmType = types.StringValue(value.String())
		} else {
			data.AuthenticationAlgorithmType = types.StringNull()
		}
		if value := res.Get("encryptionType"); value.Exists() && !data.EncryptionType.IsNull() {
			data.EncryptionType = types.StringValue(value.String())
		} else {
			data.EncryptionType = types.StringNull()
		}
		(*parent).Snmpv3Users[i] = data
	}
	if value := res.Get("snmpTrap.syslog"); value.Exists() && !data.TrapSyslog.IsNull() {
		data.TrapSyslog = types.BoolValue(value.Bool())
	} else {
		data.TrapSyslog = types.BoolNull()
	}
	if value := res.Get("snmpTrap.authentication"); value.Exists() && !data.TrapAuthentication.IsNull() {
		data.TrapAuthentication = types.BoolValue(value.Bool())
	} else {
		data.TrapAuthentication = types.BoolNull()
	}
	if value := res.Get("snmpTrap.linkUp"); value.Exists() && !data.TrapLinkUp.IsNull() {
		data.TrapLinkUp = types.BoolValue(value.Bool())
	} else {
		data.TrapLinkUp = types.BoolNull()
	}
	if value := res.Get("snmpTrap.linkDown"); value.Exists() && !data.TrapLinkDown.IsNull() {
		data.TrapLinkDown = types.BoolValue(value.Bool())
	} else {
		data.TrapLinkDown = types.BoolNull()
	}
	if value := res.Get("snmpTrap.coldStart"); value.Exists() && !data.TrapColdStart.IsNull() {
		data.TrapColdStart = types.BoolValue(value.Bool())
	} else {
		data.TrapColdStart = types.BoolNull()
	}
	if value := res.Get("snmpTrap.warmStart"); value.Exists() && !data.TrapWarmStart.IsNull() {
		data.TrapWarmStart = types.BoolValue(value.Bool())
	} else {
		data.TrapWarmStart = types.BoolNull()
	}
	if value := res.Get("snmpTrap.fruinsert"); value.Exists() && !data.TrapFieldReplacementUnitInsert.IsNull() {
		data.TrapFieldReplacementUnitInsert = types.BoolValue(value.Bool())
	} else {
		data.TrapFieldReplacementUnitInsert = types.BoolNull()
	}
	if value := res.Get("snmpTrap.fruremove"); value.Exists() && !data.TrapFieldReplacementUnitDelete.IsNull() {
		data.TrapFieldReplacementUnitDelete = types.BoolValue(value.Bool())
	} else {
		data.TrapFieldReplacementUnitDelete = types.BoolNull()
	}
	if value := res.Get("snmpTrap.entityConfigChange"); value.Exists() && !data.TrapConfigurationChange.IsNull() {
		data.TrapConfigurationChange = types.BoolValue(value.Bool())
	} else {
		data.TrapConfigurationChange = types.BoolNull()
	}
	if value := res.Get("snmpTrap.connLimitReached"); value.Exists() && !data.TrapConnectionLimitReached.IsNull() {
		data.TrapConnectionLimitReached = types.BoolValue(value.Bool())
	} else {
		data.TrapConnectionLimitReached = types.BoolNull()
	}
	if value := res.Get("snmpTrap.natPacketDiscard"); value.Exists() && !data.TrapNatPacketDiscard.IsNull() {
		data.TrapNatPacketDiscard = types.BoolValue(value.Bool())
	} else {
		data.TrapNatPacketDiscard = types.BoolNull()
	}
	if value := res.Get("snmpTrap.cputhreshold"); value.Exists() && !data.TrapCpuRisingThreshold.IsNull() {
		data.TrapCpuRisingThreshold = types.BoolValue(value.Bool())
	} else {
		data.TrapCpuRisingThreshold = types.BoolNull()
	}
	if value := res.Get("snmpTrap.cputhresholdValueInPercent"); value.Exists() && !data.TrapCpuRisingThresholdValue.IsNull() {
		data.TrapCpuRisingThresholdValue = types.Int64Value(value.Int())
	} else {
		data.TrapCpuRisingThresholdValue = types.Int64Null()
	}
	if value := res.Get("snmpTrap.cputhresholdIntervalInMins"); value.Exists() && !data.TrapCpuRisingThresholdInterval.IsNull() {
		data.TrapCpuRisingThresholdInterval = types.Int64Value(value.Int())
	} else {
		data.TrapCpuRisingThresholdInterval = types.Int64Null()
	}
	if value := res.Get("snmpTrap.memoryThreshold"); value.Exists() && !data.TrapMemoryRisingThreshold.IsNull() {
		data.TrapMemoryRisingThreshold = types.BoolValue(value.Bool())
	} else {
		data.TrapMemoryRisingThreshold = types.BoolNull()
	}
	if value := res.Get("snmpTrap.memThresholdValueInPercent"); value.Exists() && !data.TrapMemoryRisingThresholdValue.IsNull() {
		data.TrapMemoryRisingThresholdValue = types.Int64Value(value.Int())
	} else {
		data.TrapMemoryRisingThresholdValue = types.Int64Null()
	}
	if value := res.Get("snmpTrap.failoverState"); value.Exists() && !data.TrapFailover.IsNull() {
		data.TrapFailover = types.BoolValue(value.Bool())
	} else {
		data.TrapFailover = types.BoolNull()
	}
	if value := res.Get("snmpTrap.clusterState"); value.Exists() && !data.TrapCluster.IsNull() {
		data.TrapCluster = types.BoolValue(value.Bool())
	} else {
		data.TrapCluster = types.BoolNull()
	}
	if value := res.Get("snmpTrap.peerFlap"); value.Exists() && !data.TrapPeerFlap.IsNull() {
		data.TrapPeerFlap = types.BoolValue(value.Bool())
	} else {
		data.TrapPeerFlap = types.BoolNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *FTDPlatformSettingsSNMP) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyPutDelete

// toBodyPutDelete is used to create the body for PUT requests to clear the resource state
func (data FTDPlatformSettingsSNMP) toBodyPutDelete(ctx context.Context) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if data.Type.ValueString() != "" {
		body, _ = sjson.Set(body, "type", data.Type.ValueString())
	}
	return body
}

// End of section. //template:end toBodyPutDelete
