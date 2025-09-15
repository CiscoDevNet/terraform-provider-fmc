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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type VPNRASecureClientCustomization struct {
	Id                           types.String                                                 `tfsdk:"id"`
	Domain                       types.String                                                 `tfsdk:"domain"`
	VpnRaId                      types.String                                                 `tfsdk:"vpn_ra_id"`
	Type                         types.String                                                 `tfsdk:"type"`
	GuiAndTextMessages           []VPNRASecureClientCustomizationGuiAndTextMessages           `tfsdk:"gui_and_text_messages"`
	IconsAndImages               []VPNRASecureClientCustomizationIconsAndImages               `tfsdk:"icons_and_images"`
	Scripts                      []VPNRASecureClientCustomizationScripts                      `tfsdk:"scripts"`
	Binaries                     []VPNRASecureClientCustomizationBinaries                     `tfsdk:"binaries"`
	CustomInstallerTransforms    []VPNRASecureClientCustomizationCustomInstallerTransforms    `tfsdk:"custom_installer_transforms"`
	LocalizedInstallerTransforms []VPNRASecureClientCustomizationLocalizedInstallerTransforms `tfsdk:"localized_installer_transforms"`
}

type VPNRASecureClientCustomizationGuiAndTextMessages struct {
	Id types.String `tfsdk:"id"`
}

type VPNRASecureClientCustomizationIconsAndImages struct {
	Id types.String `tfsdk:"id"`
}

type VPNRASecureClientCustomizationScripts struct {
	Id types.String `tfsdk:"id"`
}

type VPNRASecureClientCustomizationBinaries struct {
	Id types.String `tfsdk:"id"`
}

type VPNRASecureClientCustomizationCustomInstallerTransforms struct {
	Id types.String `tfsdk:"id"`
}

type VPNRASecureClientCustomizationLocalizedInstallerTransforms struct {
	Id types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data VPNRASecureClientCustomization) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/ravpns/%v/secureclientcustomizationsettings", url.QueryEscape(data.VpnRaId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data VPNRASecureClientCustomization) toBody(ctx context.Context, state VPNRASecureClientCustomization) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	body, _ = sjson.Set(body, "type", "RaVpnSecureClientCustomization")
	if len(data.GuiAndTextMessages) > 0 {
		body, _ = sjson.Set(body, "languageTranslations", []any{})
		for _, item := range data.GuiAndTextMessages {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "languageTranslations.-1", itemBody)
		}
	}
	if len(data.IconsAndImages) > 0 {
		body, _ = sjson.Set(body, "imagesAndIcons", []any{})
		for _, item := range data.IconsAndImages {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "imagesAndIcons.-1", itemBody)
		}
	}
	if len(data.Scripts) > 0 {
		body, _ = sjson.Set(body, "scripts", []any{})
		for _, item := range data.Scripts {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "scripts.-1", itemBody)
		}
	}
	if len(data.Binaries) > 0 {
		body, _ = sjson.Set(body, "binaries", []any{})
		for _, item := range data.Binaries {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "binaries.-1", itemBody)
		}
	}
	if len(data.CustomInstallerTransforms) > 0 {
		body, _ = sjson.Set(body, "customizedInstallerTransforms", []any{})
		for _, item := range data.CustomInstallerTransforms {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "customizedInstallerTransforms.-1", itemBody)
		}
	}
	if len(data.LocalizedInstallerTransforms) > 0 {
		body, _ = sjson.Set(body, "localizedInstallerTransforms", []any{})
		for _, item := range data.LocalizedInstallerTransforms {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "localizedInstallerTransforms.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *VPNRASecureClientCustomization) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("languageTranslations"); value.Exists() {
		data.GuiAndTextMessages = make([]VPNRASecureClientCustomizationGuiAndTextMessages, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := VPNRASecureClientCustomizationGuiAndTextMessages{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).GuiAndTextMessages = append((*parent).GuiAndTextMessages, data)
			return true
		})
	}
	if value := res.Get("imagesAndIcons"); value.Exists() {
		data.IconsAndImages = make([]VPNRASecureClientCustomizationIconsAndImages, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := VPNRASecureClientCustomizationIconsAndImages{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).IconsAndImages = append((*parent).IconsAndImages, data)
			return true
		})
	}
	if value := res.Get("scripts"); value.Exists() {
		data.Scripts = make([]VPNRASecureClientCustomizationScripts, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := VPNRASecureClientCustomizationScripts{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).Scripts = append((*parent).Scripts, data)
			return true
		})
	}
	if value := res.Get("binaries"); value.Exists() {
		data.Binaries = make([]VPNRASecureClientCustomizationBinaries, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := VPNRASecureClientCustomizationBinaries{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).Binaries = append((*parent).Binaries, data)
			return true
		})
	}
	if value := res.Get("customizedInstallerTransforms"); value.Exists() {
		data.CustomInstallerTransforms = make([]VPNRASecureClientCustomizationCustomInstallerTransforms, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := VPNRASecureClientCustomizationCustomInstallerTransforms{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).CustomInstallerTransforms = append((*parent).CustomInstallerTransforms, data)
			return true
		})
	}
	if value := res.Get("localizedInstallerTransforms"); value.Exists() {
		data.LocalizedInstallerTransforms = make([]VPNRASecureClientCustomizationLocalizedInstallerTransforms, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := VPNRASecureClientCustomizationLocalizedInstallerTransforms{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).LocalizedInstallerTransforms = append((*parent).LocalizedInstallerTransforms, data)
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
func (data *VPNRASecureClientCustomization) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	for i := 0; i < len(data.GuiAndTextMessages); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.GuiAndTextMessages[i].Id.ValueString()}

		parent := &data
		data := (*parent).GuiAndTextMessages[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("languageTranslations").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing GuiAndTextMessages[%d] = %+v",
				i,
				(*parent).GuiAndTextMessages[i],
			))
			(*parent).GuiAndTextMessages = slices.Delete((*parent).GuiAndTextMessages, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).GuiAndTextMessages[i] = data
	}
	for i := 0; i < len(data.IconsAndImages); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.IconsAndImages[i].Id.ValueString()}

		parent := &data
		data := (*parent).IconsAndImages[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("imagesAndIcons").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing IconsAndImages[%d] = %+v",
				i,
				(*parent).IconsAndImages[i],
			))
			(*parent).IconsAndImages = slices.Delete((*parent).IconsAndImages, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).IconsAndImages[i] = data
	}
	for i := 0; i < len(data.Scripts); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.Scripts[i].Id.ValueString()}

		parent := &data
		data := (*parent).Scripts[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("scripts").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Scripts[%d] = %+v",
				i,
				(*parent).Scripts[i],
			))
			(*parent).Scripts = slices.Delete((*parent).Scripts, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).Scripts[i] = data
	}
	for i := 0; i < len(data.Binaries); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.Binaries[i].Id.ValueString()}

		parent := &data
		data := (*parent).Binaries[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("binaries").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Binaries[%d] = %+v",
				i,
				(*parent).Binaries[i],
			))
			(*parent).Binaries = slices.Delete((*parent).Binaries, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).Binaries[i] = data
	}
	for i := 0; i < len(data.CustomInstallerTransforms); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.CustomInstallerTransforms[i].Id.ValueString()}

		parent := &data
		data := (*parent).CustomInstallerTransforms[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("customizedInstallerTransforms").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing CustomInstallerTransforms[%d] = %+v",
				i,
				(*parent).CustomInstallerTransforms[i],
			))
			(*parent).CustomInstallerTransforms = slices.Delete((*parent).CustomInstallerTransforms, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).CustomInstallerTransforms[i] = data
	}
	for i := 0; i < len(data.LocalizedInstallerTransforms); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.LocalizedInstallerTransforms[i].Id.ValueString()}

		parent := &data
		data := (*parent).LocalizedInstallerTransforms[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("localizedInstallerTransforms").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing LocalizedInstallerTransforms[%d] = %+v",
				i,
				(*parent).LocalizedInstallerTransforms[i],
			))
			(*parent).LocalizedInstallerTransforms = slices.Delete((*parent).LocalizedInstallerTransforms, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).LocalizedInstallerTransforms[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *VPNRASecureClientCustomization) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
func (data VPNRASecureClientCustomization) toBodyPutDelete(ctx context.Context) string {
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
