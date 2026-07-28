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

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DeviceDHCPRelay struct {
	Id                  types.String                 `tfsdk:"id"`
	Domain              types.String                 `tfsdk:"domain"`
	DeviceId            types.String                 `tfsdk:"device_id"`
	Type                types.String                 `tfsdk:"type"`
	Ipv4RelayTimeout    types.Int64                  `tfsdk:"ipv4_relay_timeout"`
	Ipv6RelayTimeout    types.Int64                  `tfsdk:"ipv6_relay_timeout"`
	TrustAllInformation types.Bool                   `tfsdk:"trust_all_information"`
	RelayAgents         []DeviceDHCPRelayRelayAgents `tfsdk:"relay_agents"`
	Servers             []DeviceDHCPRelayServers     `tfsdk:"servers"`
}

type DeviceDHCPRelayRelayAgents struct {
	InterfaceId   types.String `tfsdk:"interface_id"`
	InterfaceName types.String `tfsdk:"interface_name"`
	InterfaceType types.String `tfsdk:"interface_type"`
	Ipv4Relay     types.Bool   `tfsdk:"ipv4_relay"`
	Ipv6Relay     types.Bool   `tfsdk:"ipv6_relay"`
	SetRoute      types.Bool   `tfsdk:"set_route"`
}

type DeviceDHCPRelayServers struct {
	ServerId            types.String                             `tfsdk:"server_id"`
	ServerInterfaceId   types.String                             `tfsdk:"server_interface_id"`
	ServerInterfaceName types.String                             `tfsdk:"server_interface_name"`
	ServerInterfaceType types.String                             `tfsdk:"server_interface_type"`
	ClientInterfaces    []DeviceDHCPRelayServersClientInterfaces `tfsdk:"client_interfaces"`
}

type DeviceDHCPRelayServersClientInterfaces struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionDeviceDHCPRelay = version.Must(version.NewVersion("7.4"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DeviceDHCPRelay) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords/%v/dhcp/dhcprelaysettings", url.QueryEscape(data.DeviceId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DeviceDHCPRelay) toBody(ctx context.Context, state DeviceDHCPRelay) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Ipv4RelayTimeout.IsNull() {
		body, _ = sjson.Set(body, "ipv4TimeoutInSec", data.Ipv4RelayTimeout.ValueInt64())
	}
	if !data.Ipv6RelayTimeout.IsNull() {
		body, _ = sjson.Set(body, "ipv6TimeoutInSec", data.Ipv6RelayTimeout.ValueInt64())
	}
	if !data.TrustAllInformation.IsNull() {
		body, _ = sjson.Set(body, "trustAllInformation", data.TrustAllInformation.ValueBool())
	}
	if len(data.RelayAgents) > 0 {
		body, _ = sjson.Set(body, "dhcpRelayAgent", []any{})
		for _, item := range data.RelayAgents {
			itemBody := ""
			if !item.InterfaceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "interface.id", item.InterfaceId.ValueString())
			}
			if !item.InterfaceName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "interface.name", item.InterfaceName.ValueString())
			}
			if !item.InterfaceType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "interface.type", item.InterfaceType.ValueString())
			}
			if !item.Ipv4Relay.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enableDHCPIPv4Relay", item.Ipv4Relay.ValueBool())
			}
			if !item.Ipv6Relay.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enableDHCPIPv6Relay", item.Ipv6Relay.ValueBool())
			}
			if !item.SetRoute.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "setRoute", item.SetRoute.ValueBool())
			}
			body, _ = sjson.SetRaw(body, "dhcpRelayAgent.-1", itemBody)
		}
	}
	if len(data.Servers) > 0 {
		body, _ = sjson.Set(body, "dhcpRelayServers", []any{})
		for _, item := range data.Servers {
			itemBody := ""
			if !item.ServerId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "server.id", item.ServerId.ValueString())
			}
			if !item.ServerInterfaceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "interface.id", item.ServerInterfaceId.ValueString())
			}
			if !item.ServerInterfaceName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "interface.name", item.ServerInterfaceName.ValueString())
			}
			if !item.ServerInterfaceType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "interface.type", item.ServerInterfaceType.ValueString())
			}
			if len(item.ClientInterfaces) > 0 {
				itemBody, _ = sjson.Set(itemBody, "clientsideInterfaces", []any{})
				for _, childItem := range item.ClientInterfaces {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Name.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "name", childItem.Name.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "clientsideInterfaces.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "dhcpRelayServers.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DeviceDHCPRelay) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("ipv4TimeoutInSec"); value.Exists() {
		data.Ipv4RelayTimeout = types.Int64Value(value.Int())
	} else {
		data.Ipv4RelayTimeout = types.Int64Value(60)
	}
	if value := res.Get("ipv6TimeoutInSec"); value.Exists() {
		data.Ipv6RelayTimeout = types.Int64Value(value.Int())
	} else {
		data.Ipv6RelayTimeout = types.Int64Value(60)
	}
	if value := res.Get("trustAllInformation"); value.Exists() {
		data.TrustAllInformation = types.BoolValue(value.Bool())
	} else {
		data.TrustAllInformation = types.BoolNull()
	}
	if value := res.Get("dhcpRelayAgent"); value.Exists() {
		data.RelayAgents = make([]DeviceDHCPRelayRelayAgents, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceDHCPRelayRelayAgents{}
			if value := res.Get("interface.id"); value.Exists() {
				data.InterfaceId = types.StringValue(value.String())
			} else {
				data.InterfaceId = types.StringNull()
			}
			if value := res.Get("interface.name"); value.Exists() {
				data.InterfaceName = types.StringValue(value.String())
			} else {
				data.InterfaceName = types.StringNull()
			}
			if value := res.Get("interface.type"); value.Exists() {
				data.InterfaceType = types.StringValue(value.String())
			} else {
				data.InterfaceType = types.StringNull()
			}
			if value := res.Get("enableDHCPIPv4Relay"); value.Exists() {
				data.Ipv4Relay = types.BoolValue(value.Bool())
			} else {
				data.Ipv4Relay = types.BoolNull()
			}
			if value := res.Get("enableDHCPIPv6Relay"); value.Exists() {
				data.Ipv6Relay = types.BoolValue(value.Bool())
			} else {
				data.Ipv6Relay = types.BoolNull()
			}
			if value := res.Get("setRoute"); value.Exists() {
				data.SetRoute = types.BoolValue(value.Bool())
			} else {
				data.SetRoute = types.BoolNull()
			}
			(*parent).RelayAgents = append((*parent).RelayAgents, data)
			return true
		})
	}
	if value := res.Get("dhcpRelayServers"); value.Exists() {
		data.Servers = make([]DeviceDHCPRelayServers, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceDHCPRelayServers{}
			if value := res.Get("server.id"); value.Exists() {
				data.ServerId = types.StringValue(value.String())
			} else {
				data.ServerId = types.StringNull()
			}
			if value := res.Get("interface.id"); value.Exists() {
				data.ServerInterfaceId = types.StringValue(value.String())
			} else {
				data.ServerInterfaceId = types.StringNull()
			}
			if value := res.Get("interface.name"); value.Exists() {
				data.ServerInterfaceName = types.StringValue(value.String())
			} else {
				data.ServerInterfaceName = types.StringNull()
			}
			if value := res.Get("interface.type"); value.Exists() {
				data.ServerInterfaceType = types.StringValue(value.String())
			} else {
				data.ServerInterfaceType = types.StringNull()
			}
			if value := res.Get("clientsideInterfaces"); value.Exists() {
				data.ClientInterfaces = make([]DeviceDHCPRelayServersClientInterfaces, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := DeviceDHCPRelayServersClientInterfaces{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
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
					(*parent).ClientInterfaces = append((*parent).ClientInterfaces, data)
					return true
				})
			}
			(*parent).Servers = append((*parent).Servers, data)
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
func (data *DeviceDHCPRelay) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("ipv4TimeoutInSec"); value.Exists() && !data.Ipv4RelayTimeout.IsNull() {
		data.Ipv4RelayTimeout = types.Int64Value(value.Int())
	} else if data.Ipv4RelayTimeout.ValueInt64() != 60 {
		data.Ipv4RelayTimeout = types.Int64Null()
	}
	if value := res.Get("ipv6TimeoutInSec"); value.Exists() && !data.Ipv6RelayTimeout.IsNull() {
		data.Ipv6RelayTimeout = types.Int64Value(value.Int())
	} else if data.Ipv6RelayTimeout.ValueInt64() != 60 {
		data.Ipv6RelayTimeout = types.Int64Null()
	}
	if value := res.Get("trustAllInformation"); value.Exists() && !data.TrustAllInformation.IsNull() {
		data.TrustAllInformation = types.BoolValue(value.Bool())
	} else {
		data.TrustAllInformation = types.BoolNull()
	}
	for i := 0; i < len(data.RelayAgents); i++ {
		keys := [...]string{"interface.id"}
		keyValues := [...]string{data.RelayAgents[i].InterfaceId.ValueString()}

		parent := &data
		data := (*parent).RelayAgents[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("dhcpRelayAgent").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing RelayAgents[%d] = %+v",
				i,
				(*parent).RelayAgents[i],
			))
			(*parent).RelayAgents = slices.Delete((*parent).RelayAgents, i, i+1)
			i--

			continue
		}
		if value := res.Get("interface.id"); value.Exists() && !data.InterfaceId.IsNull() {
			data.InterfaceId = types.StringValue(value.String())
		} else {
			data.InterfaceId = types.StringNull()
		}
		if value := res.Get("interface.name"); value.Exists() && !data.InterfaceName.IsNull() {
			data.InterfaceName = types.StringValue(value.String())
		} else {
			data.InterfaceName = types.StringNull()
		}
		if value := res.Get("interface.type"); value.Exists() && !data.InterfaceType.IsNull() {
			data.InterfaceType = types.StringValue(value.String())
		} else {
			data.InterfaceType = types.StringNull()
		}
		if value := res.Get("enableDHCPIPv4Relay"); value.Exists() && !data.Ipv4Relay.IsNull() {
			data.Ipv4Relay = types.BoolValue(value.Bool())
		} else {
			data.Ipv4Relay = types.BoolNull()
		}
		if value := res.Get("enableDHCPIPv6Relay"); value.Exists() && !data.Ipv6Relay.IsNull() {
			data.Ipv6Relay = types.BoolValue(value.Bool())
		} else {
			data.Ipv6Relay = types.BoolNull()
		}
		if value := res.Get("setRoute"); value.Exists() && !data.SetRoute.IsNull() {
			data.SetRoute = types.BoolValue(value.Bool())
		} else {
			data.SetRoute = types.BoolNull()
		}
		(*parent).RelayAgents[i] = data
	}
	for i := 0; i < len(data.Servers); i++ {
		keys := [...]string{"server.id"}
		keyValues := [...]string{data.Servers[i].ServerId.ValueString()}

		parent := &data
		data := (*parent).Servers[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("dhcpRelayServers").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Servers[%d] = %+v",
				i,
				(*parent).Servers[i],
			))
			(*parent).Servers = slices.Delete((*parent).Servers, i, i+1)
			i--

			continue
		}
		if value := res.Get("server.id"); value.Exists() && !data.ServerId.IsNull() {
			data.ServerId = types.StringValue(value.String())
		} else {
			data.ServerId = types.StringNull()
		}
		if value := res.Get("interface.id"); value.Exists() && !data.ServerInterfaceId.IsNull() {
			data.ServerInterfaceId = types.StringValue(value.String())
		} else {
			data.ServerInterfaceId = types.StringNull()
		}
		if value := res.Get("interface.name"); value.Exists() && !data.ServerInterfaceName.IsNull() {
			data.ServerInterfaceName = types.StringValue(value.String())
		} else {
			data.ServerInterfaceName = types.StringNull()
		}
		if value := res.Get("interface.type"); value.Exists() && !data.ServerInterfaceType.IsNull() {
			data.ServerInterfaceType = types.StringValue(value.String())
		} else {
			data.ServerInterfaceType = types.StringNull()
		}
		for i := 0; i < len(data.ClientInterfaces); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.ClientInterfaces[i].Id.ValueString()}

			parent := &data
			data := (*parent).ClientInterfaces[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("clientsideInterfaces").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing ClientInterfaces[%d] = %+v",
					i,
					(*parent).ClientInterfaces[i],
				))
				(*parent).ClientInterfaces = slices.Delete((*parent).ClientInterfaces, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
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
			(*parent).ClientInterfaces[i] = data
		}
		(*parent).Servers[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DeviceDHCPRelay) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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

// toBodyPutDelete is used to create the body for PUT requests to clear the resource state
func (data DeviceDHCPRelay) toBodyPutDelete(ctx context.Context) string {
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

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBody

// End of section. //template:end adjustBody

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBodyBulk

// End of section. //template:end adjustBodyBulk

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyOverrides

// End of section. //template:end toBodyOverrides

// Section below is generated&owned by "gen/generator.go". //template:begin synthesizeOverrides

// End of section. //template:end synthesizeOverrides
