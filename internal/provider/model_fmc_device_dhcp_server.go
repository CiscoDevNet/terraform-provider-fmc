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
	"strconv"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DeviceDHCPServer struct {
	Id                      types.String                  `tfsdk:"id"`
	Domain                  types.String                  `tfsdk:"domain"`
	DeviceId                types.String                  `tfsdk:"device_id"`
	Type                    types.String                  `tfsdk:"type"`
	PingTimeout             types.Int64                   `tfsdk:"ping_timeout"`
	LeaseLength             types.Int64                   `tfsdk:"lease_length"`
	AutoConfigInterfaceId   types.String                  `tfsdk:"auto_config_interface_id"`
	AutoConfigInterfaceType types.String                  `tfsdk:"auto_config_interface_type"`
	AutoConfigInterfaceName types.String                  `tfsdk:"auto_config_interface_name"`
	DomainName              types.String                  `tfsdk:"domain_name"`
	PrimaryDnsServerId      types.String                  `tfsdk:"primary_dns_server_id"`
	SecondaryDnsServerId    types.String                  `tfsdk:"secondary_dns_server_id"`
	PrimaryWinsServerId     types.String                  `tfsdk:"primary_wins_server_id"`
	SecondaryWinsServerId   types.String                  `tfsdk:"secondary_wins_server_id"`
	Servers                 []DeviceDHCPServerServers     `tfsdk:"servers"`
	DhcpOptions             []DeviceDHCPServerDhcpOptions `tfsdk:"dhcp_options"`
}

type DeviceDHCPServerServers struct {
	InterfaceId   types.String `tfsdk:"interface_id"`
	InterfaceName types.String `tfsdk:"interface_name"`
	InterfaceType types.String `tfsdk:"interface_type"`
	AddressPool   types.String `tfsdk:"address_pool"`
	Enabled       types.Bool   `tfsdk:"enabled"`
}

type DeviceDHCPServerDhcpOptions struct {
	Code              types.Int64  `tfsdk:"code"`
	Type              types.String `tfsdk:"type"`
	FirstIpAddressId  types.String `tfsdk:"first_ip_address_id"`
	SecondIpAddressId types.String `tfsdk:"second_ip_address_id"`
	Ascii             types.String `tfsdk:"ascii"`
	Hex               types.String `tfsdk:"hex"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionDeviceDHCPServer = version.Must(version.NewVersion("7.4"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DeviceDHCPServer) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords/%v/dhcp/dhcpserver", url.QueryEscape(data.DeviceId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DeviceDHCPServer) toBody(ctx context.Context, state DeviceDHCPServer) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.PingTimeout.IsNull() {
		body, _ = sjson.Set(body, "pingTimeoutInMs", data.PingTimeout.ValueInt64())
	}
	if !data.LeaseLength.IsNull() {
		body, _ = sjson.Set(body, "leaseLengthInSec", data.LeaseLength.ValueInt64())
	}
	if !data.AutoConfigInterfaceId.IsNull() {
		body, _ = sjson.Set(body, "autoConfigInterface.id", data.AutoConfigInterfaceId.ValueString())
	}
	if !data.AutoConfigInterfaceType.IsNull() {
		body, _ = sjson.Set(body, "autoConfigInterface.type", data.AutoConfigInterfaceType.ValueString())
	}
	if !data.AutoConfigInterfaceName.IsNull() {
		body, _ = sjson.Set(body, "autoConfigInterface.name", data.AutoConfigInterfaceName.ValueString())
	}
	if !data.DomainName.IsNull() {
		body, _ = sjson.Set(body, "searchDomain", data.DomainName.ValueString())
	}
	if !data.PrimaryDnsServerId.IsNull() {
		body, _ = sjson.Set(body, "primaryDNSServer.id", data.PrimaryDnsServerId.ValueString())
	}
	if !data.SecondaryDnsServerId.IsNull() {
		body, _ = sjson.Set(body, "secondaryDNSServer.id", data.SecondaryDnsServerId.ValueString())
	}
	if !data.PrimaryWinsServerId.IsNull() {
		body, _ = sjson.Set(body, "primaryWINSServer.id", data.PrimaryWinsServerId.ValueString())
	}
	if !data.SecondaryWinsServerId.IsNull() {
		body, _ = sjson.Set(body, "secondaryWINSServer.id", data.SecondaryWinsServerId.ValueString())
	}
	if len(data.Servers) > 0 {
		body, _ = sjson.Set(body, "dhcpServers", []any{})
		for _, item := range data.Servers {
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
			if !item.AddressPool.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "addressPool", item.AddressPool.ValueString())
			}
			if !item.Enabled.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enableDHCP", item.Enabled.ValueBool())
			}
			body, _ = sjson.SetRaw(body, "dhcpServers.-1", itemBody)
		}
	}
	if len(data.DhcpOptions) > 0 {
		body, _ = sjson.Set(body, "dhcpOptions", []any{})
		for _, item := range data.DhcpOptions {
			itemBody := ""
			if !item.Code.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "optionCode", item.Code.ValueInt64())
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "optionType", item.Type.ValueString())
			}
			if !item.FirstIpAddressId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "firstIpAddress.id", item.FirstIpAddressId.ValueString())
			}
			if !item.SecondIpAddressId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "secondIpAddress.id", item.SecondIpAddressId.ValueString())
			}
			if !item.Ascii.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ascii", item.Ascii.ValueString())
			}
			if !item.Hex.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "hex", item.Hex.ValueString())
			}
			body, _ = sjson.SetRaw(body, "dhcpOptions.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DeviceDHCPServer) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("pingTimeoutInMs"); value.Exists() {
		data.PingTimeout = types.Int64Value(value.Int())
	} else {
		data.PingTimeout = types.Int64Value(50)
	}
	if value := res.Get("leaseLengthInSec"); value.Exists() {
		data.LeaseLength = types.Int64Value(value.Int())
	} else {
		data.LeaseLength = types.Int64Value(3600)
	}
	if value := res.Get("autoConfigInterface.id"); value.Exists() {
		data.AutoConfigInterfaceId = types.StringValue(value.String())
	} else {
		data.AutoConfigInterfaceId = types.StringNull()
	}
	if value := res.Get("autoConfigInterface.type"); value.Exists() {
		data.AutoConfigInterfaceType = types.StringValue(value.String())
	} else {
		data.AutoConfigInterfaceType = types.StringNull()
	}
	if value := res.Get("autoConfigInterface.name"); value.Exists() {
		data.AutoConfigInterfaceName = types.StringValue(value.String())
	} else {
		data.AutoConfigInterfaceName = types.StringNull()
	}
	if value := res.Get("searchDomain"); value.Exists() {
		data.DomainName = types.StringValue(value.String())
	} else {
		data.DomainName = types.StringNull()
	}
	if value := res.Get("primaryDNSServer.id"); value.Exists() {
		data.PrimaryDnsServerId = types.StringValue(value.String())
	} else {
		data.PrimaryDnsServerId = types.StringNull()
	}
	if value := res.Get("secondaryDNSServer.id"); value.Exists() {
		data.SecondaryDnsServerId = types.StringValue(value.String())
	} else {
		data.SecondaryDnsServerId = types.StringNull()
	}
	if value := res.Get("primaryWINSServer.id"); value.Exists() {
		data.PrimaryWinsServerId = types.StringValue(value.String())
	} else {
		data.PrimaryWinsServerId = types.StringNull()
	}
	if value := res.Get("secondaryWINSServer.id"); value.Exists() {
		data.SecondaryWinsServerId = types.StringValue(value.String())
	} else {
		data.SecondaryWinsServerId = types.StringNull()
	}
	if value := res.Get("dhcpServers"); value.Exists() {
		data.Servers = make([]DeviceDHCPServerServers, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceDHCPServerServers{}
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
			if value := res.Get("addressPool"); value.Exists() {
				data.AddressPool = types.StringValue(value.String())
			} else {
				data.AddressPool = types.StringNull()
			}
			if value := res.Get("enableDHCP"); value.Exists() {
				data.Enabled = types.BoolValue(value.Bool())
			} else {
				data.Enabled = types.BoolNull()
			}
			(*parent).Servers = append((*parent).Servers, data)
			return true
		})
	}
	if value := res.Get("dhcpOptions"); value.Exists() {
		data.DhcpOptions = make([]DeviceDHCPServerDhcpOptions, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceDHCPServerDhcpOptions{}
			if value := res.Get("optionCode"); value.Exists() {
				data.Code = types.Int64Value(value.Int())
			} else {
				data.Code = types.Int64Null()
			}
			if value := res.Get("optionType"); value.Exists() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			if value := res.Get("firstIpAddress.id"); value.Exists() {
				data.FirstIpAddressId = types.StringValue(value.String())
			} else {
				data.FirstIpAddressId = types.StringNull()
			}
			if value := res.Get("secondIpAddress.id"); value.Exists() {
				data.SecondIpAddressId = types.StringValue(value.String())
			} else {
				data.SecondIpAddressId = types.StringNull()
			}
			if value := res.Get("ascii"); value.Exists() {
				data.Ascii = types.StringValue(value.String())
			} else {
				data.Ascii = types.StringNull()
			}
			if value := res.Get("hex"); value.Exists() {
				data.Hex = types.StringValue(value.String())
			} else {
				data.Hex = types.StringNull()
			}
			(*parent).DhcpOptions = append((*parent).DhcpOptions, data)
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
func (data *DeviceDHCPServer) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("pingTimeoutInMs"); value.Exists() && !data.PingTimeout.IsNull() {
		data.PingTimeout = types.Int64Value(value.Int())
	} else if data.PingTimeout.ValueInt64() != 50 {
		data.PingTimeout = types.Int64Null()
	}
	if value := res.Get("leaseLengthInSec"); value.Exists() && !data.LeaseLength.IsNull() {
		data.LeaseLength = types.Int64Value(value.Int())
	} else if data.LeaseLength.ValueInt64() != 3600 {
		data.LeaseLength = types.Int64Null()
	}
	if value := res.Get("autoConfigInterface.id"); value.Exists() && !data.AutoConfigInterfaceId.IsNull() {
		data.AutoConfigInterfaceId = types.StringValue(value.String())
	} else {
		data.AutoConfigInterfaceId = types.StringNull()
	}
	if value := res.Get("autoConfigInterface.type"); value.Exists() && !data.AutoConfigInterfaceType.IsNull() {
		data.AutoConfigInterfaceType = types.StringValue(value.String())
	} else {
		data.AutoConfigInterfaceType = types.StringNull()
	}
	if value := res.Get("autoConfigInterface.name"); value.Exists() && !data.AutoConfigInterfaceName.IsNull() {
		data.AutoConfigInterfaceName = types.StringValue(value.String())
	} else {
		data.AutoConfigInterfaceName = types.StringNull()
	}
	if value := res.Get("searchDomain"); value.Exists() && !data.DomainName.IsNull() {
		data.DomainName = types.StringValue(value.String())
	} else {
		data.DomainName = types.StringNull()
	}
	if value := res.Get("primaryDNSServer.id"); value.Exists() && !data.PrimaryDnsServerId.IsNull() {
		data.PrimaryDnsServerId = types.StringValue(value.String())
	} else {
		data.PrimaryDnsServerId = types.StringNull()
	}
	if value := res.Get("secondaryDNSServer.id"); value.Exists() && !data.SecondaryDnsServerId.IsNull() {
		data.SecondaryDnsServerId = types.StringValue(value.String())
	} else {
		data.SecondaryDnsServerId = types.StringNull()
	}
	if value := res.Get("primaryWINSServer.id"); value.Exists() && !data.PrimaryWinsServerId.IsNull() {
		data.PrimaryWinsServerId = types.StringValue(value.String())
	} else {
		data.PrimaryWinsServerId = types.StringNull()
	}
	if value := res.Get("secondaryWINSServer.id"); value.Exists() && !data.SecondaryWinsServerId.IsNull() {
		data.SecondaryWinsServerId = types.StringValue(value.String())
	} else {
		data.SecondaryWinsServerId = types.StringNull()
	}
	for i := 0; i < len(data.Servers); i++ {
		keys := [...]string{"interface.id"}
		keyValues := [...]string{data.Servers[i].InterfaceId.ValueString()}

		parent := &data
		data := (*parent).Servers[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("dhcpServers").ForEach(
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
		if value := res.Get("addressPool"); value.Exists() && !data.AddressPool.IsNull() {
			data.AddressPool = types.StringValue(value.String())
		} else {
			data.AddressPool = types.StringNull()
		}
		if value := res.Get("enableDHCP"); value.Exists() && !data.Enabled.IsNull() {
			data.Enabled = types.BoolValue(value.Bool())
		} else {
			data.Enabled = types.BoolNull()
		}
		(*parent).Servers[i] = data
	}
	for i := 0; i < len(data.DhcpOptions); i++ {
		keys := [...]string{"optionCode"}
		keyValues := [...]string{strconv.FormatInt(data.DhcpOptions[i].Code.ValueInt64(), 10)}

		parent := &data
		data := (*parent).DhcpOptions[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("dhcpOptions").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing DhcpOptions[%d] = %+v",
				i,
				(*parent).DhcpOptions[i],
			))
			(*parent).DhcpOptions = slices.Delete((*parent).DhcpOptions, i, i+1)
			i--

			continue
		}
		if value := res.Get("optionCode"); value.Exists() && !data.Code.IsNull() {
			data.Code = types.Int64Value(value.Int())
		} else {
			data.Code = types.Int64Null()
		}
		if value := res.Get("optionType"); value.Exists() && !data.Type.IsNull() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		if value := res.Get("firstIpAddress.id"); value.Exists() && !data.FirstIpAddressId.IsNull() {
			data.FirstIpAddressId = types.StringValue(value.String())
		} else {
			data.FirstIpAddressId = types.StringNull()
		}
		if value := res.Get("secondIpAddress.id"); value.Exists() && !data.SecondIpAddressId.IsNull() {
			data.SecondIpAddressId = types.StringValue(value.String())
		} else {
			data.SecondIpAddressId = types.StringNull()
		}
		if value := res.Get("ascii"); value.Exists() && !data.Ascii.IsNull() {
			data.Ascii = types.StringValue(value.String())
		} else {
			data.Ascii = types.StringNull()
		}
		if value := res.Get("hex"); value.Exists() && !data.Hex.IsNull() {
			data.Hex = types.StringValue(value.String())
		} else {
			data.Hex = types.StringNull()
		}
		(*parent).DhcpOptions[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DeviceDHCPServer) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
func (data DeviceDHCPServer) toBodyPutDelete(ctx context.Context) string {
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
