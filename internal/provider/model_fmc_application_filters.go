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
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type ApplicationFilters struct {
	Id     types.String                       `tfsdk:"id"`
	Domain types.String                       `tfsdk:"domain"`
	Items  map[string]ApplicationFiltersItems `tfsdk:"items"`
}

type ApplicationFiltersItems struct {
	Id           types.String                          `tfsdk:"id"`
	Type         types.String                          `tfsdk:"type"`
	Applications []ApplicationFiltersItemsApplications `tfsdk:"applications"`
	Filters      []ApplicationFiltersItemsFilters      `tfsdk:"filters"`
}

type ApplicationFiltersItemsApplications struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}
type ApplicationFiltersItemsFilters struct {
	Types              []ApplicationFiltersItemsFiltersTypes              `tfsdk:"types"`
	Risks              []ApplicationFiltersItemsFiltersRisks              `tfsdk:"risks"`
	BusinessRelevances []ApplicationFiltersItemsFiltersBusinessRelevances `tfsdk:"business_relevances"`
	Categories         []ApplicationFiltersItemsFiltersCategories         `tfsdk:"categories"`
	Tags               []ApplicationFiltersItemsFiltersTags               `tfsdk:"tags"`
}

type ApplicationFiltersItemsFiltersTypes struct {
	Id types.String `tfsdk:"id"`
}
type ApplicationFiltersItemsFiltersRisks struct {
	Id types.String `tfsdk:"id"`
}
type ApplicationFiltersItemsFiltersBusinessRelevances struct {
	Id types.String `tfsdk:"id"`
}
type ApplicationFiltersItemsFiltersCategories struct {
	Id types.String `tfsdk:"id"`
}
type ApplicationFiltersItemsFiltersTags struct {
	Id types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplicationFilters) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/applicationfilters"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplicationFilters) toBody(ctx context.Context, state ApplicationFilters) string {
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
			if len(item.Applications) > 0 {
				itemBody, _ = sjson.Set(itemBody, "applications", []interface{}{})
				for _, childItem := range item.Applications {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Name.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "name", childItem.Name.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "applications.-1", itemChildBody)
				}
			}
			if len(item.Filters) > 0 {
				itemBody, _ = sjson.Set(itemBody, "appConditions", []interface{}{})
				for _, childItem := range item.Filters {
					itemChildBody := ""
					if len(childItem.Types) > 0 {
						itemChildBody, _ = sjson.Set(itemChildBody, "applicationTypes", []interface{}{})
						for _, childChildItem := range childItem.Types {
							itemChildChildBody := ""
							if !childChildItem.Id.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "id", childChildItem.Id.ValueString())
							}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "applicationTypes.-1", itemChildChildBody)
						}
					}
					if len(childItem.Risks) > 0 {
						itemChildBody, _ = sjson.Set(itemChildBody, "risks", []interface{}{})
						for _, childChildItem := range childItem.Risks {
							itemChildChildBody := ""
							if !childChildItem.Id.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "id", childChildItem.Id.ValueString())
							}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "risks.-1", itemChildChildBody)
						}
					}
					if len(childItem.BusinessRelevances) > 0 {
						itemChildBody, _ = sjson.Set(itemChildBody, "productivities", []interface{}{})
						for _, childChildItem := range childItem.BusinessRelevances {
							itemChildChildBody := ""
							if !childChildItem.Id.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "id", childChildItem.Id.ValueString())
							}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "productivities.-1", itemChildChildBody)
						}
					}
					if len(childItem.Categories) > 0 {
						itemChildBody, _ = sjson.Set(itemChildBody, "categories", []interface{}{})
						for _, childChildItem := range childItem.Categories {
							itemChildChildBody := ""
							if !childChildItem.Id.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "id", childChildItem.Id.ValueString())
							}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "categories.-1", itemChildChildBody)
						}
					}
					if len(childItem.Tags) > 0 {
						itemChildBody, _ = sjson.Set(itemChildBody, "tags", []interface{}{})
						for _, childChildItem := range childItem.Tags {
							itemChildChildBody := ""
							if !childChildItem.Id.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "id", childChildItem.Id.ValueString())
							}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "tags.-1", itemChildChildBody)
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "appConditions.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "items.-1", itemBody)
		}
	}
	return gjson.Get(body, "items").String()
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplicationFilters) fromBody(ctx context.Context, res gjson.Result) {
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
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		if value := res.Get("applications"); value.Exists() {
			data.Applications = make([]ApplicationFiltersItemsApplications, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := ApplicationFiltersItemsApplications{}
				if value := res.Get("name"); value.Exists() {
					data.Name = types.StringValue(value.String())
				} else {
					data.Name = types.StringNull()
				}
				(*parent).Applications = append((*parent).Applications, data)
				return true
			})
		}
		if value := res.Get("appConditions"); value.Exists() {
			data.Filters = make([]ApplicationFiltersItemsFilters, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := ApplicationFiltersItemsFilters{}
				if value := res.Get("applicationTypes"); value.Exists() {
					data.Types = make([]ApplicationFiltersItemsFiltersTypes, 0)
					value.ForEach(func(k, res gjson.Result) bool {
						parent := &data
						data := ApplicationFiltersItemsFiltersTypes{}
						if value := res.Get("id"); value.Exists() {
							data.Id = types.StringValue(value.String())
						} else {
							data.Id = types.StringNull()
						}
						(*parent).Types = append((*parent).Types, data)
						return true
					})
				}
				if value := res.Get("risks"); value.Exists() {
					data.Risks = make([]ApplicationFiltersItemsFiltersRisks, 0)
					value.ForEach(func(k, res gjson.Result) bool {
						parent := &data
						data := ApplicationFiltersItemsFiltersRisks{}
						if value := res.Get("id"); value.Exists() {
							data.Id = types.StringValue(value.String())
						} else {
							data.Id = types.StringNull()
						}
						(*parent).Risks = append((*parent).Risks, data)
						return true
					})
				}
				if value := res.Get("productivities"); value.Exists() {
					data.BusinessRelevances = make([]ApplicationFiltersItemsFiltersBusinessRelevances, 0)
					value.ForEach(func(k, res gjson.Result) bool {
						parent := &data
						data := ApplicationFiltersItemsFiltersBusinessRelevances{}
						if value := res.Get("id"); value.Exists() {
							data.Id = types.StringValue(value.String())
						} else {
							data.Id = types.StringNull()
						}
						(*parent).BusinessRelevances = append((*parent).BusinessRelevances, data)
						return true
					})
				}
				if value := res.Get("categories"); value.Exists() {
					data.Categories = make([]ApplicationFiltersItemsFiltersCategories, 0)
					value.ForEach(func(k, res gjson.Result) bool {
						parent := &data
						data := ApplicationFiltersItemsFiltersCategories{}
						if value := res.Get("id"); value.Exists() {
							data.Id = types.StringValue(value.String())
						} else {
							data.Id = types.StringNull()
						}
						(*parent).Categories = append((*parent).Categories, data)
						return true
					})
				}
				if value := res.Get("tags"); value.Exists() {
					data.Tags = make([]ApplicationFiltersItemsFiltersTags, 0)
					value.ForEach(func(k, res gjson.Result) bool {
						parent := &data
						data := ApplicationFiltersItemsFiltersTags{}
						if value := res.Get("id"); value.Exists() {
							data.Id = types.StringValue(value.String())
						} else {
							data.Id = types.StringNull()
						}
						(*parent).Tags = append((*parent).Tags, data)
						return true
					})
				}
				(*parent).Filters = append((*parent).Filters, data)
				return true
			})
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
func (data *ApplicationFilters) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
		if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		for i := 0; i < len(data.Applications); i++ {
			keys := [...]string{"name"}
			keyValues := [...]string{data.Applications[i].Name.ValueString()}

			parent := &data
			data := (*parent).Applications[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("applications").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing Applications[%d] = %+v",
					i,
					(*parent).Applications[i],
				))
				(*parent).Applications = slices.Delete((*parent).Applications, i, i+1)
				i--

				continue
			}
			if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			(*parent).Applications[i] = data
		}
		{
			l := len(res.Get("appConditions").Array())
			tflog.Debug(ctx, fmt.Sprintf("appConditions array resizing from %d to %d", len(data.Filters), l))
			for i := len(data.Filters); i < l; i++ {
				data.Filters = append(data.Filters, ApplicationFiltersItemsFilters{})
			}
			if len(data.Filters) > l {
				data.Filters = data.Filters[:l]
			}
		}
		for i := range data.Filters {
			parent := &data
			data := (*parent).Filters[i]
			parentRes := &res
			res := parentRes.Get(fmt.Sprintf("appConditions.%d", i))
			for i := 0; i < len(data.Types); i++ {
				keys := [...]string{"id"}
				keyValues := [...]string{data.Types[i].Id.ValueString()}

				parent := &data
				data := (*parent).Types[i]
				parentRes := &res
				var res gjson.Result

				parentRes.Get("applicationTypes").ForEach(
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
					tflog.Debug(ctx, fmt.Sprintf("removing Types[%d] = %+v",
						i,
						(*parent).Types[i],
					))
					(*parent).Types = slices.Delete((*parent).Types, i, i+1)
					i--

					continue
				}
				if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
					data.Id = types.StringValue(value.String())
				} else {
					data.Id = types.StringNull()
				}
				(*parent).Types[i] = data
			}
			for i := 0; i < len(data.Risks); i++ {
				keys := [...]string{"id"}
				keyValues := [...]string{data.Risks[i].Id.ValueString()}

				parent := &data
				data := (*parent).Risks[i]
				parentRes := &res
				var res gjson.Result

				parentRes.Get("risks").ForEach(
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
					tflog.Debug(ctx, fmt.Sprintf("removing Risks[%d] = %+v",
						i,
						(*parent).Risks[i],
					))
					(*parent).Risks = slices.Delete((*parent).Risks, i, i+1)
					i--

					continue
				}
				if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
					data.Id = types.StringValue(value.String())
				} else {
					data.Id = types.StringNull()
				}
				(*parent).Risks[i] = data
			}
			for i := 0; i < len(data.BusinessRelevances); i++ {
				keys := [...]string{"id"}
				keyValues := [...]string{data.BusinessRelevances[i].Id.ValueString()}

				parent := &data
				data := (*parent).BusinessRelevances[i]
				parentRes := &res
				var res gjson.Result

				parentRes.Get("productivities").ForEach(
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
					tflog.Debug(ctx, fmt.Sprintf("removing BusinessRelevances[%d] = %+v",
						i,
						(*parent).BusinessRelevances[i],
					))
					(*parent).BusinessRelevances = slices.Delete((*parent).BusinessRelevances, i, i+1)
					i--

					continue
				}
				if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
					data.Id = types.StringValue(value.String())
				} else {
					data.Id = types.StringNull()
				}
				(*parent).BusinessRelevances[i] = data
			}
			for i := 0; i < len(data.Categories); i++ {
				keys := [...]string{"id"}
				keyValues := [...]string{data.Categories[i].Id.ValueString()}

				parent := &data
				data := (*parent).Categories[i]
				parentRes := &res
				var res gjson.Result

				parentRes.Get("categories").ForEach(
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
					tflog.Debug(ctx, fmt.Sprintf("removing Categories[%d] = %+v",
						i,
						(*parent).Categories[i],
					))
					(*parent).Categories = slices.Delete((*parent).Categories, i, i+1)
					i--

					continue
				}
				if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
					data.Id = types.StringValue(value.String())
				} else {
					data.Id = types.StringNull()
				}
				(*parent).Categories[i] = data
			}
			for i := 0; i < len(data.Tags); i++ {
				keys := [...]string{"id"}
				keyValues := [...]string{data.Tags[i].Id.ValueString()}

				parent := &data
				data := (*parent).Tags[i]
				parentRes := &res
				var res gjson.Result

				parentRes.Get("tags").ForEach(
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
					tflog.Debug(ctx, fmt.Sprintf("removing Tags[%d] = %+v",
						i,
						(*parent).Tags[i],
					))
					(*parent).Tags = slices.Delete((*parent).Tags, i, i+1)
					i--

					continue
				}
				if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
					data.Id = types.StringValue(value.String())
				} else {
					data.Id = types.StringNull()
				}
				(*parent).Tags[i] = data
			}
			(*parent).Filters[i] = data
		}
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ApplicationFilters) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
		if v := data.Items[i]; v.Type.IsUnknown() {
			if value := r.Get("type"); value.Exists() {
				v.Type = types.StringValue(value.String())
			} else {
				v.Type = types.StringNull()
			}
			data.Items[i] = v
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin Clone

func (data *ApplicationFilters) Clone() ApplicationFilters {
	ret := *data
	ret.Items = maps.Clone(data.Items)

	return ret
}

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// Updates done one-by-one require different API body
func (data ApplicationFilters) toBodyNonBulk(ctx context.Context, state ApplicationFilters) string {
	// This is one-by-one update, so only one element to update is expected
	if len(data.Items) > 1 {
		tflog.Error(ctx, "Found more than one element to chage. Only one will be changed.")
	}

	// Utilize existing toBody function
	body := data.toBody(ctx, state)

	// Get first element only
	return gjson.Get(body, "0").String()
}

// End of section. //template:end toBodyNonBulk
