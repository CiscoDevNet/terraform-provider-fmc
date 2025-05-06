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
	"maps"
	"net/url"
	"slices"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type VPNS2SEndpoints struct {
	Id       types.String                    `tfsdk:"id"`
	Domain   types.String                    `tfsdk:"domain"`
	VpnS2sId types.String                    `tfsdk:"vpn_s2s_id"`
	Items    map[string]VPNS2SEndpointsItems `tfsdk:"items"`
}

type VPNS2SEndpointsItems struct {
	Id                             types.String                            `tfsdk:"id"`
	PeerType                       types.String                            `tfsdk:"peer_type"`
	IsExtranet                     types.Bool                              `tfsdk:"is_extranet"`
	ExtranetIpAddress              types.String                            `tfsdk:"extranet_ip_address"`
	ExtranetIsDynamicIp            types.Bool                              `tfsdk:"extranet_is_dynamic_ip"`
	DeviceId                       types.String                            `tfsdk:"device_id"`
	InterfaceId                    types.String                            `tfsdk:"interface_id"`
	InterfaceIpv6Address           types.String                            `tfsdk:"interface_ipv6_address"`
	InterfacePublicIpAddress       types.String                            `tfsdk:"interface_public_ip_address"`
	ConnectionType                 types.String                            `tfsdk:"connection_type"`
	AllowIncomingIkev2Routes       types.Bool                              `tfsdk:"allow_incoming_ikev2_routes"`
	SendTunnelInterfaceIpToPeer    types.Bool                              `tfsdk:"send_tunnel_interface_ip_to_peer"`
	ProtectedNetworks              []VPNS2SEndpointsItemsProtectedNetworks `tfsdk:"protected_networks"`
	AclId                          types.String                            `tfsdk:"acl_id"`
	EnableNatTraversal             types.Bool                              `tfsdk:"enable_nat_traversal"`
	EnableNatExemption             types.Bool                              `tfsdk:"enable_nat_exemption"`
	InsideInterfaceId              types.String                            `tfsdk:"inside_interface_id"`
	EnableReverseRouteInjection    types.Bool                              `tfsdk:"enable_reverse_route_injection"`
	LocalIdentityType              types.String                            `tfsdk:"local_identity_type"`
	LocalIdentityString            types.String                            `tfsdk:"local_identity_string"`
	VpnFilterAclId                 types.String                            `tfsdk:"vpn_filter_acl_id"`
	OverrideRemoteVpnFilter        types.Bool                              `tfsdk:"override_remote_vpn_filter"`
	RemoteVpnFilterAclId           types.String                            `tfsdk:"remote_vpn_filter_acl_id"`
	BackupInterfaceId              types.String                            `tfsdk:"backup_interface_id"`
	BackupInterfacePublicIpAddress types.String                            `tfsdk:"backup_interface_public_ip_address"`
	BackupLocalIdentityType        types.String                            `tfsdk:"backup_local_identity_type"`
	BackupLocalIdentityString      types.String                            `tfsdk:"backup_local_identity_string"`
}

type VPNS2SEndpointsItemsProtectedNetworks struct {
	Id types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionBulkDeleteVPNS2SEndpoints = version.Must(version.NewVersion("999"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data VPNS2SEndpoints) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/ftds2svpns/%v/endpoints", url.QueryEscape(data.VpnS2sId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data VPNS2SEndpoints) toBody(ctx context.Context, state VPNS2SEndpoints) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if len(data.Items) > 0 {
		body, _ = sjson.Set(body, "items", []interface{}{})
		for key, item := range data.Items {
			itemBody, _ := sjson.Set("{}", "name", key)
			if !item.Id.IsNull() && !item.Id.IsUnknown() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.PeerType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "peerType", item.PeerType.ValueString())
			}
			if !item.IsExtranet.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "extranet", item.IsExtranet.ValueBool())
			}
			if !item.ExtranetIpAddress.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "extranetInfo.ipAddress", item.ExtranetIpAddress.ValueString())
			}
			if !item.ExtranetIsDynamicIp.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "extranetInfo.isDynamicIP", item.ExtranetIsDynamicIp.ValueBool())
			}
			if !item.DeviceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "device.id", item.DeviceId.ValueString())
			}
			if !item.InterfaceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "interface.id", item.InterfaceId.ValueString())
			}
			if !item.InterfaceIpv6Address.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipv6InterfaceAddress", item.InterfaceIpv6Address.ValueString())
			}
			if !item.InterfacePublicIpAddress.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "nattedInterfaceAddress", item.InterfacePublicIpAddress.ValueString())
			}
			if !item.ConnectionType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "connectionType", item.ConnectionType.ValueString())
			}
			if !item.AllowIncomingIkev2Routes.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "allowIncomingIKEv2Routes", item.AllowIncomingIkev2Routes.ValueBool())
			}
			if !item.SendTunnelInterfaceIpToPeer.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sendTunnelInterfaceIpToPeer", item.SendTunnelInterfaceIpToPeer.ValueBool())
			}
			if len(item.ProtectedNetworks) > 0 {
				itemBody, _ = sjson.Set(itemBody, "protectedNetworks.networks", []interface{}{})
				for _, childItem := range item.ProtectedNetworks {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "protectedNetworks.networks.-1", itemChildBody)
				}
			}
			if !item.AclId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "protectedNetworks.acl.id", item.AclId.ValueString())
			}
			if !item.EnableNatTraversal.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enableNatTraversal", item.EnableNatTraversal.ValueBool())
			}
			if !item.EnableNatExemption.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enableNATExempt", item.EnableNatExemption.ValueBool())
			}
			if !item.InsideInterfaceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "insideInterface.0.id", item.InsideInterfaceId.ValueString())
			}
			if !item.EnableReverseRouteInjection.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "dynamicRRIEnabled", item.EnableReverseRouteInjection.ValueBool())
			}
			if !item.LocalIdentityType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "localIdentityType", item.LocalIdentityType.ValueString())
			}
			if !item.LocalIdentityString.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "localIdentityString", item.LocalIdentityString.ValueString())
			}
			if !item.VpnFilterAclId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "vpnFilterAcl.id", item.VpnFilterAclId.ValueString())
			}
			if !item.OverrideRemoteVpnFilter.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "overrideRemoteVpnFilter", item.OverrideRemoteVpnFilter.ValueBool())
			}
			if !item.RemoteVpnFilterAclId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "remoteVpnFilterAclObject.id", item.RemoteVpnFilterAclId.ValueString())
			}
			if !item.BackupInterfaceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "backupInterfaces.0.interface.id", item.BackupInterfaceId.ValueString())
			}
			if !item.BackupInterfacePublicIpAddress.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "backupInterfaces.0.publicIpAddress", item.BackupInterfacePublicIpAddress.ValueString())
			}
			if !item.BackupLocalIdentityType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "backupInterfaces.0.localIdentityType", item.BackupLocalIdentityType.ValueString())
			}
			if !item.BackupLocalIdentityString.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "backupInterfaces.0.localIdentityString", item.BackupLocalIdentityString.ValueString())
			}
			body, _ = sjson.SetRaw(body, "items.-1", itemBody)
		}
	}
	return gjson.Get(body, "items").String()
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *VPNS2SEndpoints) fromBody(ctx context.Context, res gjson.Result) {
	for k := range data.Items {
		parent := &data
		data := (*parent).Items[k]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("items").ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("name").String() == k {
					res = v
					return false // break ForEach
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("subresource not found, removing: name=%v", k))
			delete((*parent).Items, k)
			continue
		}
		if value := res.Get("id"); value.Exists() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("peerType"); value.Exists() {
			data.PeerType = types.StringValue(value.String())
		} else {
			data.PeerType = types.StringNull()
		}
		if value := res.Get("extranet"); value.Exists() {
			data.IsExtranet = types.BoolValue(value.Bool())
		} else {
			data.IsExtranet = types.BoolNull()
		}
		if value := res.Get("extranetInfo.ipAddress"); value.Exists() {
			data.ExtranetIpAddress = types.StringValue(value.String())
		} else {
			data.ExtranetIpAddress = types.StringNull()
		}
		if value := res.Get("extranetInfo.isDynamicIP"); value.Exists() {
			data.ExtranetIsDynamicIp = types.BoolValue(value.Bool())
		} else {
			data.ExtranetIsDynamicIp = types.BoolNull()
		}
		if value := res.Get("device.id"); value.Exists() {
			data.DeviceId = types.StringValue(value.String())
		} else {
			data.DeviceId = types.StringNull()
		}
		if value := res.Get("interface.id"); value.Exists() {
			data.InterfaceId = types.StringValue(value.String())
		} else {
			data.InterfaceId = types.StringNull()
		}
		if value := res.Get("ipv6InterfaceAddress"); value.Exists() {
			data.InterfaceIpv6Address = types.StringValue(value.String())
		} else {
			data.InterfaceIpv6Address = types.StringNull()
		}
		if value := res.Get("nattedInterfaceAddress"); value.Exists() {
			data.InterfacePublicIpAddress = types.StringValue(value.String())
		} else {
			data.InterfacePublicIpAddress = types.StringNull()
		}
		if value := res.Get("connectionType"); value.Exists() {
			data.ConnectionType = types.StringValue(value.String())
		} else {
			data.ConnectionType = types.StringNull()
		}
		if value := res.Get("allowIncomingIKEv2Routes"); value.Exists() {
			data.AllowIncomingIkev2Routes = types.BoolValue(value.Bool())
		} else {
			data.AllowIncomingIkev2Routes = types.BoolNull()
		}
		if value := res.Get("sendTunnelInterfaceIpToPeer"); value.Exists() {
			data.SendTunnelInterfaceIpToPeer = types.BoolValue(value.Bool())
		} else {
			data.SendTunnelInterfaceIpToPeer = types.BoolNull()
		}
		if value := res.Get("protectedNetworks.networks"); value.Exists() {
			data.ProtectedNetworks = make([]VPNS2SEndpointsItemsProtectedNetworks, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := VPNS2SEndpointsItemsProtectedNetworks{}
				if value := res.Get("id"); value.Exists() {
					data.Id = types.StringValue(value.String())
				} else {
					data.Id = types.StringNull()
				}
				(*parent).ProtectedNetworks = append((*parent).ProtectedNetworks, data)
				return true
			})
		}
		if value := res.Get("protectedNetworks.acl.id"); value.Exists() {
			data.AclId = types.StringValue(value.String())
		} else {
			data.AclId = types.StringNull()
		}
		if value := res.Get("enableNatTraversal"); value.Exists() {
			data.EnableNatTraversal = types.BoolValue(value.Bool())
		} else {
			data.EnableNatTraversal = types.BoolNull()
		}
		if value := res.Get("insideInterface.0.id"); value.Exists() {
			data.InsideInterfaceId = types.StringValue(value.String())
		} else {
			data.InsideInterfaceId = types.StringNull()
		}
		if value := res.Get("dynamicRRIEnabled"); value.Exists() {
			data.EnableReverseRouteInjection = types.BoolValue(value.Bool())
		} else {
			data.EnableReverseRouteInjection = types.BoolNull()
		}
		if value := res.Get("localIdentityType"); value.Exists() {
			data.LocalIdentityType = types.StringValue(value.String())
		} else {
			data.LocalIdentityType = types.StringNull()
		}
		if value := res.Get("localIdentityString"); value.Exists() {
			data.LocalIdentityString = types.StringValue(value.String())
		} else {
			data.LocalIdentityString = types.StringNull()
		}
		if value := res.Get("vpnFilterAcl.id"); value.Exists() {
			data.VpnFilterAclId = types.StringValue(value.String())
		} else {
			data.VpnFilterAclId = types.StringNull()
		}
		if value := res.Get("overrideRemoteVpnFilter"); value.Exists() {
			data.OverrideRemoteVpnFilter = types.BoolValue(value.Bool())
		} else {
			data.OverrideRemoteVpnFilter = types.BoolNull()
		}
		if value := res.Get("remoteVpnFilterAclObject.id"); value.Exists() {
			data.RemoteVpnFilterAclId = types.StringValue(value.String())
		} else {
			data.RemoteVpnFilterAclId = types.StringNull()
		}
		if value := res.Get("backupInterfaces.0.interface.id"); value.Exists() {
			data.BackupInterfaceId = types.StringValue(value.String())
		} else {
			data.BackupInterfaceId = types.StringNull()
		}
		if value := res.Get("backupInterfaces.0.publicIpAddress"); value.Exists() {
			data.BackupInterfacePublicIpAddress = types.StringValue(value.String())
		} else {
			data.BackupInterfacePublicIpAddress = types.StringNull()
		}
		if value := res.Get("backupInterfaces.0.localIdentityType"); value.Exists() {
			data.BackupLocalIdentityType = types.StringValue(value.String())
		} else {
			data.BackupLocalIdentityType = types.StringNull()
		}
		if value := res.Get("backupInterfaces.0.localIdentityString"); value.Exists() {
			data.BackupLocalIdentityString = types.StringValue(value.String())
		} else {
			data.BackupLocalIdentityString = types.StringNull()
		}
		(*parent).Items[k] = data
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *VPNS2SEndpoints) fromBodyPartial(ctx context.Context, res gjson.Result) {
	for i := range data.Items {
		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("items").ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("id").String() == data.Id.ValueString() && data.Id.ValueString() != "" {
					res = v
					return false // break ForEach
				}
				return true
			},
		)
		if value := res.Get("id"); value.Exists() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("peerType"); value.Exists() && !data.PeerType.IsNull() {
			data.PeerType = types.StringValue(value.String())
		} else {
			data.PeerType = types.StringNull()
		}
		if value := res.Get("extranet"); value.Exists() && !data.IsExtranet.IsNull() {
			data.IsExtranet = types.BoolValue(value.Bool())
		} else {
			data.IsExtranet = types.BoolNull()
		}
		if value := res.Get("extranetInfo.ipAddress"); value.Exists() && !data.ExtranetIpAddress.IsNull() {
			data.ExtranetIpAddress = types.StringValue(value.String())
		} else {
			data.ExtranetIpAddress = types.StringNull()
		}
		if value := res.Get("extranetInfo.isDynamicIP"); value.Exists() && !data.ExtranetIsDynamicIp.IsNull() {
			data.ExtranetIsDynamicIp = types.BoolValue(value.Bool())
		} else {
			data.ExtranetIsDynamicIp = types.BoolNull()
		}
		if value := res.Get("device.id"); value.Exists() && !data.DeviceId.IsNull() {
			data.DeviceId = types.StringValue(value.String())
		} else {
			data.DeviceId = types.StringNull()
		}
		if value := res.Get("interface.id"); value.Exists() && !data.InterfaceId.IsNull() {
			data.InterfaceId = types.StringValue(value.String())
		} else {
			data.InterfaceId = types.StringNull()
		}
		if value := res.Get("ipv6InterfaceAddress"); value.Exists() && !data.InterfaceIpv6Address.IsNull() {
			data.InterfaceIpv6Address = types.StringValue(value.String())
		} else {
			data.InterfaceIpv6Address = types.StringNull()
		}
		if value := res.Get("nattedInterfaceAddress"); value.Exists() && !data.InterfacePublicIpAddress.IsNull() {
			data.InterfacePublicIpAddress = types.StringValue(value.String())
		} else {
			data.InterfacePublicIpAddress = types.StringNull()
		}
		if value := res.Get("connectionType"); value.Exists() && !data.ConnectionType.IsNull() {
			data.ConnectionType = types.StringValue(value.String())
		} else {
			data.ConnectionType = types.StringNull()
		}
		if value := res.Get("allowIncomingIKEv2Routes"); value.Exists() && !data.AllowIncomingIkev2Routes.IsNull() {
			data.AllowIncomingIkev2Routes = types.BoolValue(value.Bool())
		} else {
			data.AllowIncomingIkev2Routes = types.BoolNull()
		}
		if value := res.Get("sendTunnelInterfaceIpToPeer"); value.Exists() && !data.SendTunnelInterfaceIpToPeer.IsNull() {
			data.SendTunnelInterfaceIpToPeer = types.BoolValue(value.Bool())
		} else {
			data.SendTunnelInterfaceIpToPeer = types.BoolNull()
		}
		for i := 0; i < len(data.ProtectedNetworks); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.ProtectedNetworks[i].Id.ValueString()}

			parent := &data
			data := (*parent).ProtectedNetworks[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("protectedNetworks.networks").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing ProtectedNetworks[%d] = %+v",
					i,
					(*parent).ProtectedNetworks[i],
				))
				(*parent).ProtectedNetworks = slices.Delete((*parent).ProtectedNetworks, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).ProtectedNetworks[i] = data
		}
		if value := res.Get("protectedNetworks.acl.id"); value.Exists() && !data.AclId.IsNull() {
			data.AclId = types.StringValue(value.String())
		} else {
			data.AclId = types.StringNull()
		}
		if value := res.Get("enableNatTraversal"); value.Exists() && !data.EnableNatTraversal.IsNull() {
			data.EnableNatTraversal = types.BoolValue(value.Bool())
		} else {
			data.EnableNatTraversal = types.BoolNull()
		}
		if value := res.Get("insideInterface.0.id"); value.Exists() && !data.InsideInterfaceId.IsNull() {
			data.InsideInterfaceId = types.StringValue(value.String())
		} else {
			data.InsideInterfaceId = types.StringNull()
		}
		if value := res.Get("dynamicRRIEnabled"); value.Exists() && !data.EnableReverseRouteInjection.IsNull() {
			data.EnableReverseRouteInjection = types.BoolValue(value.Bool())
		} else {
			data.EnableReverseRouteInjection = types.BoolNull()
		}
		if value := res.Get("localIdentityType"); value.Exists() && !data.LocalIdentityType.IsNull() {
			data.LocalIdentityType = types.StringValue(value.String())
		} else {
			data.LocalIdentityType = types.StringNull()
		}
		if value := res.Get("localIdentityString"); value.Exists() && !data.LocalIdentityString.IsNull() {
			data.LocalIdentityString = types.StringValue(value.String())
		} else {
			data.LocalIdentityString = types.StringNull()
		}
		if value := res.Get("vpnFilterAcl.id"); value.Exists() && !data.VpnFilterAclId.IsNull() {
			data.VpnFilterAclId = types.StringValue(value.String())
		} else {
			data.VpnFilterAclId = types.StringNull()
		}
		if value := res.Get("overrideRemoteVpnFilter"); value.Exists() && !data.OverrideRemoteVpnFilter.IsNull() {
			data.OverrideRemoteVpnFilter = types.BoolValue(value.Bool())
		} else {
			data.OverrideRemoteVpnFilter = types.BoolNull()
		}
		if value := res.Get("remoteVpnFilterAclObject.id"); value.Exists() && !data.RemoteVpnFilterAclId.IsNull() {
			data.RemoteVpnFilterAclId = types.StringValue(value.String())
		} else {
			data.RemoteVpnFilterAclId = types.StringNull()
		}
		if value := res.Get("backupInterfaces.0.interface.id"); value.Exists() && !data.BackupInterfaceId.IsNull() {
			data.BackupInterfaceId = types.StringValue(value.String())
		} else {
			data.BackupInterfaceId = types.StringNull()
		}
		if value := res.Get("backupInterfaces.0.publicIpAddress"); value.Exists() && !data.BackupInterfacePublicIpAddress.IsNull() {
			data.BackupInterfacePublicIpAddress = types.StringValue(value.String())
		} else {
			data.BackupInterfacePublicIpAddress = types.StringNull()
		}
		if value := res.Get("backupInterfaces.0.localIdentityType"); value.Exists() && !data.BackupLocalIdentityType.IsNull() {
			data.BackupLocalIdentityType = types.StringValue(value.String())
		} else {
			data.BackupLocalIdentityType = types.StringNull()
		}
		if value := res.Get("backupInterfaces.0.localIdentityString"); value.Exists() && !data.BackupLocalIdentityString.IsNull() {
			data.BackupLocalIdentityString = types.StringValue(value.String())
		} else {
			data.BackupLocalIdentityString = types.StringNull()
		}
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *VPNS2SEndpoints) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	for i, val := range data.Items {
		var r gjson.Result
		res.Get("items").ForEach(
			func(_, v gjson.Result) bool {
				if val.Id.IsUnknown() {
					if v.Get("name").String() == i {
						r = v
						return false // break ForEach
					}
				} else {
					if v.Get("id").String() == val.Id.ValueString() && val.Id.ValueString() != "" {
						r = v
						return false // break ForEach
					}
				}

				return true
			},
		)
		if v := data.Items[i]; v.Id.IsUnknown() {
			if value := r.Get("id"); value.Exists() {
				v.Id = types.StringValue(value.String())
			} else {
				v.Id = types.StringNull()
			}
			data.Items[i] = v
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin Clone

func (data *VPNS2SEndpoints) Clone() VPNS2SEndpoints {
	ret := *data
	ret.Items = maps.Clone(data.Items)

	return ret
}

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// Updates done one-by-one require different API body
func (data VPNS2SEndpoints) toBodyNonBulk(ctx context.Context, state VPNS2SEndpoints) string {
	// This is one-by-one update, so only one element to update is expected
	if len(data.Items) > 1 {
		tflog.Error(ctx, "Found more than one element to change. Only one will be changed.")
	}

	// Utilize existing toBody function
	body := data.toBody(ctx, state)

	// Get first element only
	return gjson.Get(body, "0").String()
}

// End of section. //template:end toBodyNonBulk

// Section below is generated&owned by "gen/generator.go". //template:begin findObjectsToBeReplaced

// End of section. //template:end findObjectsToBeReplaced

// Section below is generated&owned by "gen/generator.go". //template:begin clearItemIds

// End of section. //template:end clearItemIds

// VPN S2S Endpoints do have some fileds that are set based on other fields.
// This function is used to fix those fields before sending the request to the API.
func (data VPNS2SEndpoints) fixFields(ctx context.Context, req string) string {
	reqJson := gjson.Parse(req)
	itemsJson := reqJson.Array()

	for k, v := range data.Items {
		for idx, val := range itemsJson {
			if val.Get("name").String() == k {
				var header = fmt.Sprintf("%d.", idx)

				if v.IsExtranet.ValueBool() {
					req, _ = sjson.Set(req, header+"extranetInfo.name", k)
				} else {
					req, _ = sjson.Set(req, header+"device.name", k)
				}

				if v.LocalIdentityType.ValueString() != "" {
					req, _ = sjson.Set(req, header+"isLocalTunnelIdEnabled", true)
				}

				if v.BackupLocalIdentityType.ValueString() != "" {
					req, _ = sjson.Set(req, header+"backupInterfaces.0.isLocalTunnelIdEnabled", true)
				}
			}
		}
	}

	return req
}

func (data VPNS2SEndpoints) fixFieldsNonBulk(ctx context.Context, req string) string {
	fixed := data.fixFields(ctx, fmt.Sprintf("[%s]", req))
	return gjson.Get(fixed, "0").String()
}
