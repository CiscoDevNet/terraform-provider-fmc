//go:build ignore
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
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type {{camelCase .Name}} struct {
	Id types.String `tfsdk:"id"`
{{- if isDomainDependent .}}
	Domain types.String `tfsdk:"domain"`
{{- end}}
{{- range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSet .}}
	{{toGoName .TfName}} []{{.GoTypeName}} `tfsdk:"{{.TfName}}"`
{{- else if isNestedMap .}}
	{{toGoName .TfName}} map[string]{{.GoTypeName}} `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
{{- end}}
}

{{range .Attributes}}
{{- if not .Value}}
{{- if isNestedListMapSet .}}
type {{.GoTypeName}} struct {
{{- range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSet .}}
	{{toGoName .TfName}} []{{.GoTypeName}} `tfsdk:"{{.TfName}}"`
{{- else if isNestedMap .}}
	{{toGoName .TfName}} map[string]{{.GoTypeName}} `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
{{- end}}
}
{{- end}}
{{- end}}
{{end}}

{{range .Attributes}}
{{- if not .Value}}
{{- if isNestedListMapSet .}}
{{range .Attributes}}
{{- if not .Value}}
{{- if isNestedListMapSet .}}
type {{.GoTypeName}} struct {
{{- range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSet .}}
	{{toGoName .TfName}} []{{.GoTypeName}} `tfsdk:"{{.TfName}}"`
{{- else if isNestedMap .}}
	{{toGoName .TfName}} map[string]{{.GoTypeName}} `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
{{- end}}
}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{end}}

{{range .Attributes}}
{{- if not .Value}}
{{- if isNestedListMapSet .}}
{{range .Attributes}}
{{- if not .Value}}
{{- if isNestedListMapSet .}}
{{range .Attributes}}
{{- if not .Value}}
{{- if isNestedListMapSet .}}
type {{.GoTypeName}} struct {
{{- range .Attributes}}
{{- if not .Value}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{end}}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

{{- if .MinimumVersion}}
var minFMCVersion{{camelCase .Name}} = version.Must(version.NewVersion("{{.MinimumVersion}}"))
{{- end}}
{{- if .MinimumVersionCreate}}
var minFMCVersionCreate{{camelCase .Name}} = version.Must(version.NewVersion("{{.MinimumVersionCreate}}"))
{{- end}}
{{- if .MinimumVersionBulkCreate}}
var minFMCVersionBulkCreate{{camelCase .Name}} = version.Must(version.NewVersion("{{.MinimumVersionBulkCreate}}"))
{{- end}}
{{- if .MinimumVersionBulkDelete}}
var minFMCVersionBulkDelete{{camelCase .Name}} = version.Must(version.NewVersion("{{.MinimumVersionBulkDelete}}"))
{{- end}}
{{- if .BulkSizeCreate}}
const bulkSizeCreate{{camelCase .Name}} int = {{.BulkSizeCreate}}
{{- end}}

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data {{camelCase .Name}}) getPath() string {
	{{- if hasReference .Attributes}}
		return fmt.Sprintf("{{.RestEndpoint}}"{{range .Attributes}}{{if .Reference}}, url.QueryEscape(data.{{toGoName .TfName}}.Value{{.Type}}()){{end}}{{end}})
	{{- else}}
		return "{{.RestEndpoint}}"
	{{- end}}
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
{{if not .NoResource}}
func (data {{camelCase .Name}}) toBody(ctx context.Context, state {{camelCase .Name}}) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	{{- range .Attributes}}
	{{- if and .Computed (not .ComputedBodyParam)}}{{- continue}}{{- end}}
	{{- if .Value}}
	body, _ = sjson.Set(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
	{{- else if .ResourceId}}
	if state.{{toGoName .TfName}}.ValueString() != "" {
		body, _ = sjson.Set(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", state.{{toGoName .TfName}}.ValueString())
	}
	{{- else if not .Reference}}
	{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
	if !data.{{toGoName .TfName}}.IsNull() {{if .WriteChangesOnly}}&& data.{{toGoName .TfName}} != state.{{toGoName .TfName}}{{end}} {{if .ComputedBodyParam}} && !data.{{toGoName .TfName}}.IsUnknown(){{end}} {
		body, _ = sjson.Set(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", data.{{toGoName .TfName}}.Value{{.Type}}())
	}
	{{- else if isListSet .}}
	if !data.{{toGoName .TfName}}.IsNull() {
		var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
		data.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", values)
	}
	{{- else if isNestedListMapSet .}}
	if len(data.{{toGoName .TfName}}) > 0 {
		body, _ = sjson.Set(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", []interface{}{})
		{{- if isNestedMap .}}
		for key, item := range data.{{toGoName .TfName}} {
			itemBody, _ := sjson.Set("{}", "name", key)
		{{- else}}
		for _, item := range data.{{toGoName .TfName}} {
			itemBody := ""
		{{- end}}
			{{- range .Attributes}}
			{{- if and .Computed (not .ComputedBodyParam)}}{{- continue}}{{- end}}
			{{- if .Value}}
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
			{{- else if not .Reference}}
			{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
			if !item.{{toGoName .TfName}}.IsNull() {{ if .ResourceId -}} && !item.{{toGoName .TfName}}.IsUnknown() {{- end}} {{if .ComputedBodyParam}} && !data.{{toGoName .TfName}}.IsUnknown(){{end}} {
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", item.{{toGoName .TfName}}.Value{{.Type}}())
			}
			{{- else if isListSet .}}
			if !item.{{toGoName .TfName}}.IsNull() {
				var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
				item.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", values)
			}
			{{- else if isNestedListSet .}}
			if len(item.{{toGoName .TfName}}) > 0 {
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", []interface{}{})
				for _, childItem := range item.{{toGoName .TfName}} {
					itemChildBody := ""
					{{- range .Attributes}}
					{{- if and .Computed (not .ComputedBodyParam)}}{{- continue}}{{- end}}
					{{- if .Value}}
					itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
					{{- else if not .Reference}}
					{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
					if !childItem.{{toGoName .TfName}}.IsNull() {{if .ComputedBodyParam}} && !data.{{toGoName .TfName}}.IsUnknown(){{end}} {
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", childItem.{{toGoName .TfName}}.Value{{.Type}}())
					}
					{{- else if isListSet .}}
					if !childItem.{{toGoName .TfName}}.IsNull() {
						var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
						childItem.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", values)
					}
					{{- else if isNestedListSet .}}
					if len(childItem.{{toGoName .TfName}}) > 0 {
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", []interface{}{})
						for _, childChildItem := range childItem.{{toGoName .TfName}} {
							itemChildChildBody := ""
							{{- range .Attributes}}
							{{- if and .Computed (not .ComputedBodyParam)}}{{- continue}}{{- end}}
							{{- if .Value}}
							itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
							{{- else if not .Reference}}
							{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
							if !childChildItem.{{toGoName .TfName}}.IsNull() {{if .ComputedBodyParam}} && !data.{{toGoName .TfName}}.IsUnknown(){{end}} {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", childChildItem.{{toGoName .TfName}}.Value{{.Type}}())
							}
							{{- else if isListSet .}}
							if !childChildItem.{{toGoName .TfName}}.IsNull() {
								var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
								childChildItem.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", values)
							}
							{{- end}}
							{{- end}}
							{{- end}}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.-1", itemChildChildBody)
						}
					}
					{{- end}}
					{{- end}}
					{{- end}}
					itemBody, _ = sjson.SetRaw(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.-1", itemChildBody)
				}
			}
			{{- end}}
			{{- end}}
			{{- end}}
			body, _ = sjson.SetRaw(body, "{{range .DataPath}}{{.}}.{{end}}{{if .ModelName}}{{.ModelName}}.{{end}}-1", itemBody)
		}
	}
	{{- end}}
	{{- end}}
	{{- end}}
	{{- if .IsBulk}}
	return gjson.Get(body, "items").String()
	{{- else}}
	return body
	{{- end}}
}
{{end}}
// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *{{camelCase .Name}}) fromBody(ctx context.Context, res gjson.Result) {
{{- define "fromBodyTemplate"}}
	{{- range .Attributes}}
	{{- if and (not .Value) (not .WriteOnly) (not .Reference)}}
	{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
	if value := res.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() {
		data.{{toGoName .TfName}} = types.{{.Type}}Value(value.{{if eq .Type "Int64"}}Int{{else if eq .Type "Float64"}}Float{{else}}{{.Type}}{{end}}())
	} else {
		{{- if .DefaultValue}}
		data.{{toGoName .TfName}} = types.{{.Type}}Value({{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}})
		{{- else}}
		data.{{toGoName .TfName}} = types.{{.Type}}Null()
		{{- end}}
	}
	{{- else if isListSet .}}
	if value := res.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() {
		data.{{toGoName .TfName}} = helpers.Get{{.ElementType}}{{.Type}}(value.Array())
	} else {
		data.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
	}
	{{- else if isNestedListSet .}}
	if value := res{{if .ModelName}}.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"){{end}}; value.Exists() {
		data.{{toGoName .TfName}} = make([]{{.GoTypeName}}, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := {{.GoTypeName}}{}
			{{- template "fromBodyTemplate" .}}
			(*parent).{{toGoName .TfName}} = append((*parent).{{toGoName .TfName}}, data)
			return true
		})
	}
	{{- else if isNestedMap .}}
	for k := range data.{{toGoName .TfName}} {
		parent := &data
		data := (*parent).{{toGoName .TfName}}[k]
		parentRes := &res
		var res gjson.Result

		parentRes.{{if .ModelName}}Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}").{{end}}ForEach(
			func(_, v gjson.Result) bool {
				{{- if or $.ImportNameQuery $.IsBulk -}}
				if v.Get("name").String() == k {
				{{- else -}}
				if v.Get("id").String() == data.Id.ValueString() && data.Id.ValueString() != "" {
				{{- end}}
					res = v
					return false // break ForEach
				}
				return true
			},
		)
		if !res.Exists() {
			{{- if or $.ImportNameQuery $.IsBulk -}}
			tflog.Debug(ctx, fmt.Sprintf("subresource not found, removing: name=%v", k))
			{{- else -}}
			tflog.Debug(ctx, fmt.Sprintf("subresource not found, removing: uuid=%s, key=%v", data.Id, k))
			{{- end}}
			delete((*parent).{{toGoName .TfName}}, k)
			continue
		}
		{{- template "fromBodyTemplate" .}}
		(*parent).{{toGoName .TfName}}[k] = data
	}
	{{- end}}
	{{- end}}
	{{- end}}
{{- end}}
{{- template "fromBodyTemplate" .}}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial
{{- define "fromBodyPartialTemplate"}}
	{{- range .Attributes}}
	{{- if and (not .Value) (not .WriteOnly) (not .Reference)}}
	{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
	if value := res.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists(){{if not .ResourceId}}{{if not .ComputedRefreshValue}} && !data.{{toGoName .TfName}}.IsNull(){{end}}{{end}} {
		data.{{toGoName .TfName}} = types.{{.Type}}Value(value.{{if eq .Type "Int64"}}Int{{else if eq .Type "Float64"}}Float{{else}}{{.Type}}{{end}}())
	} else {{if .DefaultValue}}if data.{{toGoName .TfName}}.Value{{.Type}}() != {{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}} {{end}}{
		data.{{toGoName .TfName}} = types.{{.Type}}Null()
	}
	{{- else if isListSet .}}
	if value := res.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() && !data.{{toGoName .TfName}}.IsNull() {
		data.{{toGoName .TfName}} = helpers.Get{{.ElementType}}{{.Type}}(value.Array())
	} else {
		data.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
	}
	{{- else if isNestedListMapSet .}}
	{{- $list := (toGoName .TfName)}}
	{{- if isNestedMap .}}
	for i := range data.{{toGoName .TfName}} {
		parent := &data
		data := (*parent).{{toGoName .TfName}}[i]
		parentRes := &res
		var res gjson.Result

		parentRes.{{if .ModelName}}Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}").{{end}}ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("id").String() == data.Id.ValueString() && data.Id.ValueString() != "" {
					res = v
					return false // break ForEach
				}
				return true
			},
		)
	{{- else if .OrderedList }}
	{
		l := len(res.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}").Array())
		tflog.Debug(ctx, fmt.Sprintf("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}} array resizing from %d to %d", len(data.{{toGoName .TfName}}), l))
		for i := len(data.{{toGoName .TfName}}); i < l; i++ {
			data.{{toGoName .TfName}} = append(data.{{toGoName .TfName}}, {{.GoTypeName}}{})
		}
		if len(data.{{toGoName .TfName}}) > l {
			data.{{toGoName .TfName}} = data.{{toGoName .TfName}}[:l]
		}
	}
	for i := range data.{{toGoName .TfName}} {
		parent := &data
		data := (*parent).{{toGoName .TfName}}[i]
		parentRes := &res
		res := parentRes.Get(fmt.Sprintf("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.%d", i))
	{{- else }}
	for i := 0; i < len(data.{{toGoName .TfName}}); i++ {
		keys := [...]string{ {{$noId := not (hasId .Attributes)}}{{range .Attributes}}{{if or .Id (and $noId (not .Value))}}{{if or (eq .Type "Int64") (eq .Type "Bool") (eq .Type "String")}}"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{end}}{{end}}{{end}} }
		keyValues := [...]string{ {{$noId := not (hasId .Attributes)}}{{range .Attributes}}{{if or .Id (and $noId (not .Value))}}{{if eq .Type "Int64"}}strconv.FormatInt(data.{{$list}}[i].{{toGoName .TfName}}.ValueInt64(), 10), {{else if eq .Type "Bool"}}strconv.FormatBool(data.{{$list}}[i].{{toGoName .TfName}}.ValueBool()), {{else if eq .Type "String"}}data.{{$list}}[i].{{toGoName .TfName}}.Value{{.Type}}(), {{end}}{{end}}{{end}} }

		parent := &data
		data := (*parent).{{toGoName .TfName}}[i]
		parentRes := &res
		var res gjson.Result

		parentRes.{{if .ModelName}}Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}").{{end}}ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing {{toGoName .TfName}}[%d] = %+v",
				i,
				(*parent).{{toGoName .TfName}}[i],
			))
			(*parent).{{toGoName .TfName}} = slices.Delete((*parent).{{toGoName .TfName}}, i, i+1)
			i--

			continue
		}
	{{- end}}

		{{- template "fromBodyPartialTemplate" .}}
		(*parent).{{toGoName .TfName}}[i] = data
	}
	{{- end}}
	{{- end}}
	{{- end}}
{{- end}}

{{if not .NoResource}}
// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *{{camelCase .Name}}) fromBodyPartial(ctx context.Context, res gjson.Result) {
{{- template "fromBodyPartialTemplate" .}}
}
{{end}}
// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns
{{if not .NoResource}}
// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *{{camelCase .Name}}) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	{{- range .Attributes}}
	{{- if or (and .ResourceId (not .Reference)) .Computed}}
	{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
	{{- if not .ComputedRefreshValue }}
	if data.{{toGoName .TfName}}.IsUnknown() {
	{{- end}}
		if value := res.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() {
			data.{{toGoName .TfName}} = types.{{.Type}}Value(value.{{if eq .Type "Int64"}}Int{{else if eq .Type "Float64"}}Float{{else}}{{.Type}}{{end}}())
		} else {
			data.{{toGoName .TfName}} = types.{{.Type}}Null()
		}
	{{- if not .ComputedRefreshValue }}
	}
	{{- end}}
	{{- else}}
	{{- errorf "resource_id is not yet implemented for type %v" .Type}}
	{{- end}}
	{{- else if isNestedListMapSet .}}
	{{- if hasResourceId .Attributes}}
	{{- $list := (toGoName .TfName)}}
	{{- if .OrderedList }}
	for i := range data.{{toGoName .TfName}} {
		r := res.Get(fmt.Sprintf("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.%d", i))
	{{- else if isNestedListSet .}}
	for i := range data.{{toGoName .TfName}} {
		keys := [...]string{ {{$noId := not (hasId .Attributes)}}{{range .Attributes}}{{if or .Id (and $noId (not .Value))}}{{if or (eq .Type "Int64") (eq .Type "Bool") (eq .Type "String")}}"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{end}}{{end}}{{end}} }
		keyValues := [...]string{ {{$noId := not (hasId .Attributes)}}{{range .Attributes}}{{if or .Id (and $noId (not .Value))}}{{if eq .Type "Int64"}}strconv.FormatInt(data.{{$list}}[i].{{toGoName .TfName}}.ValueInt64(), 10), {{else if eq .Type "Bool"}}strconv.FormatBool(data.{{$list}}[i].{{toGoName .TfName}}.ValueBool()), {{else if eq .Type "String"}}data.{{$list}}[i].{{toGoName .TfName}}.Value{{.Type}}(), {{end}}{{end}}{{end}} }

		var r gjson.Result
		res.{{if .ModelName}}Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}").{{end}}ForEach(
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
					r = v
					return false
				}
				return true
			},
		)
	{{- else if isNestedMap .}}
	for i, val := range data.{{toGoName .TfName}} {
		var r gjson.Result
		res.{{if .ModelName}}Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}").{{end}}ForEach(
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
	{{- end}}

		{{- range .Attributes}}
		{{- if or (and .ResourceId (not .Reference)) .Computed}}
		{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
		if v := data.{{$list}}[i]; v.{{toGoName .TfName}}.IsUnknown() {
			if value := r.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() {
				v.{{toGoName .TfName}} = types.{{.Type}}Value(value.{{if eq .Type "Int64"}}Int{{else if eq .Type "Float64"}}Float{{else}}{{.Type}}{{end}}())
			} else {
				v.{{toGoName .TfName}} = types.{{.Type}}Null()
			}
			data.{{$list}}[i] = v
		}
		{{- else}}
		{{- errorf "resource_id is not yet implemented for type %v" .Type}}
		{{- end}}
		{{- end}}
		{{- end}}
	}
	{{- end}}
	{{- end}}
	{{- end}}
}
{{end}}
// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin Clone

{{if and .IsBulk (not .NoResource)}}
func (data *{{camelCase .Name}}) Clone() {{camelCase .Name}} {
	ret := *data
	ret.Items = maps.Clone(data.Items)

	return ret
}
{{- end}}

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

{{if and .IsBulk (not .NoResource) }}
// Updates done one-by-one require different API body
func (data {{camelCase .Name}}) toBodyNonBulk(ctx context.Context, state {{camelCase .Name}}) string {
	// This is one-by-one update, so only one element to update is expected
	if len(data.Items) > 1 {
		tflog.Error(ctx, "Found more than one element to change. Only one will be changed.")
	}

	// Utilize existing toBody function
	body := data.toBody(ctx, state)

	// Get first element only
	return gjson.Get(body, "0").String()
}
{{- end}}

// End of section. //template:end toBodyNonBulk

// Section below is generated&owned by "gen/generator.go". //template:begin findObjectsToBeReplaced

{{if and .IsBulk (hasRequiresReplace .Attributes) }}
// Check if single object within bulk requires replace due to `requires_replace`
// Since here we assume object has changed, it must be present in both state and plan (data)
func (data {{camelCase .Name}}) findObjectsToBeReplaced(ctx context.Context, state {{camelCase .Name}}) {{camelCase .Name}} {
	// Prepare empty object to be filled in with objects that require replace
	var toBeReplaced {{camelCase .Name}}
	toBeReplaced.Items = make(map[string]{{camelCase .Name}}Items)

	// Iterate over all objects in plan
	for key, item := range data.Items {
		// Check if object is present in state
		if _, ok := state.Items[key]; !ok {
			// Object is not present in state, hence it's not a candidate for replace
			continue
		}

		// Check if any field marked as `requires_replace` has changed
		{{- range .Attributes}}
			{{- if eq .TfName "items"}}
			{{- range .Attributes}}
			{{- if .RequiresReplace }}
				{{- if (eq .Type "String")}}
					if item.{{toGoName .TfName}} != state.Items[key].{{toGoName .TfName}} {
						toBeReplaced.Items[key] = item
						continue
					}
				{{- else}}
					{{- errorf "requires_replace is not supported for %v" .Type }}
				{{- end}}
			{{- end}}
			{{- end}}
			{{- end}}
		{{- end}}
	}

	return toBeReplaced
}
{{- end}}

// End of section. //template:end findObjectsToBeReplaced

// Section below is generated&owned by "gen/generator.go". //template:begin clearItemIds

{{if and .IsBulk (hasRequiresReplace .Attributes) (not .NoResource) }}
func (data *{{camelCase .Name}}) clearItemsIds(ctx context.Context) {
	for key, value := range data.Items {
		tmp := value
		tmp.Id = types.StringNull()
		data.Items[key] = tmp
	}
}
{{- end}}

// End of section. //template:end clearItemIds

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyPutDelete

{{if .PutDelete}}
// toBodyPutDelete is used to create the body for PUT requests to clear the resource state
func (data {{camelCase .Name}}) toBodyPutDelete(ctx context.Context) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	return body
}
{{- end}}

// End of section. //template:end toBodyPutDelete

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBody

{{if .AdjustBody}}
func (data {{camelCase .Name}}) adjustBody(ctx context.Context, req string) string {
	return req
}
{{- end}}

// End of section. //template:end adjustBody

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBodyBulk

{{if and .AdjustBody .IsBulk}}
func (data {{camelCase .Name}}) adjustBodyBulk(ctx context.Context, req string) string {
	return req
}
{{- end}}

// End of section. //template:end adjustBodyBulk

{{- range .Attributes}}
	{{- if isNestedMap .}}
		{{- $found := false }}
		{{- range .Attributes }}
			{{- if and (eq .ModelName "id") (eq .TfName "id") .ResourceId}}
				{{- $found = true }}
			{{- end}}
		{{- end}}
		{{- if and (not $found) (not $.NoResource) }}
			{{- errorf "type Map with attributes has a limitation for now: it must always contain attribute with `model_name: id` and `tf_name: id` and `resource_id: true`, because it must always be used to track subresources. This limitation is applicable for Resources only."}}
		{{- end}}
	{{- end}}
	{{- range .Attributes}}
		{{- if isNestedMap .}}
			{{- errorf "Map not yet implemented at this depth"}}
		{{- end}}
		{{- if hasResourceId .Attributes}}
			{{- errorf "resource_id not yet implemented at this depth"}}
		{{- end}}

		{{- range .Attributes}}
			{{- if isNestedMap .}}
				{{- errorf "Map not yet implemented at this depth"}}
			{{- end}}
			{{- if .OrderedList }}
				{{- errorf "ordered_list not yet implemented at this depth"}}
			{{- end}}
			{{- if hasResourceId .Attributes}}
				{{- errorf "resource_id not yet implemented at this depth"}}
			{{- end}}
			{{- if and $.IsBulk .RequiresReplace}}
				{{- errorf "requires_replace is not supported for nested objects in bulk operations" }}
			{{- end}}

			{{- range .Attributes}}
				{{- if isNestedMap .}}
					{{- errorf "Map not yet implemented at this depth"}}
				{{- end}}
				{{- if .OrderedList }}
					{{- errorf "ordered_list not yet implemented at this depth"}}
				{{- end}}
				{{- if hasResourceId .Attributes}}
					{{- errorf "resource_id not yet implemented at this depth"}}
				{{- end}}
				{{- range .Attributes}}
					{{- errorf "attributes not yet implemented at this depth"}}
				{{- end}}
				{{- if and $.IsBulk .RequiresReplace}}
					{{- errorf "requires_replace is not supported for nested objects in bulk operations" }}
				{{- end}}
			{{- end}}
		{{- end}}
	{{- end}}
{{- end}}
