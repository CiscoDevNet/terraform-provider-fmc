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
	"github.com/netascode/terraform-provider-fmc/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
{{- $name := camelCase .Name}}

type {{camelCase .Name}} struct {
	Id types.String `tfsdk:"id"`
	Domain types.String `tfsdk:"domain"`
{{- range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSet .}}
	{{toGoName .TfName}} []{{$name}}{{toGoName .TfName}} `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
{{- end}}
}

{{range .Attributes}}
{{- if not .Value}}
{{- $childName := toGoName .TfName}}
{{- if isNestedListSet .}}
type {{$name}}{{toGoName .TfName}} struct {
{{- range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSet .}}
	{{toGoName .TfName}} []{{$name}}{{$childName}}{{toGoName .TfName}} `tfsdk:"{{.TfName}}"`
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
{{- $childName := toGoName .TfName}}
{{- if isNestedListSet .}}
{{range .Attributes}}
{{- if not .Value}}
{{- $childChildName := toGoName .TfName}}
{{- if isNestedListSet .}}
type {{$name}}{{$childName}}{{toGoName .TfName}} struct {
{{- range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSet .}}
	{{toGoName .TfName}} []{{$name}}{{$childName}}{{$childChildName}}{{toGoName .TfName}} `tfsdk:"{{.TfName}}"`
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
{{- $childName := toGoName .TfName}}
{{- if isNestedListSet .}}
{{range .Attributes}}
{{- if not .Value}}
{{- $childChildName := toGoName .TfName}}
{{- if isNestedListSet .}}
{{range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSet .}}
type {{$name}}{{$childName}}{{$childChildName}}{{toGoName .TfName}} struct {
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

func (data {{camelCase .Name}}) toBody(ctx context.Context, state {{camelCase .Name}}) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	{{- range .Attributes}}
	{{- if .Value}}
	body, _ = sjson.Set(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
	{{- else if .ResourceId}}
	if state.{{toGoName .TfName}}.ValueString() != "" {
		body, _ = sjson.Set(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", state.{{toGoName .TfName}}.ValueString())
	}
	{{- else if not .Reference}}
	{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
	if !data.{{toGoName .TfName}}.IsNull() {{if .WriteChangesOnly}}&& data.{{toGoName .TfName}} != state.{{toGoName .TfName}}{{end}} {
		body, _ = sjson.Set(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", data.{{toGoName .TfName}}.Value{{.Type}}())
	}
	{{- else if isListSet .}}
	if !data.{{toGoName .TfName}}.IsNull() {
		var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
		data.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", values)
	}
	{{- else if isNestedListSet .}}
	if len(data.{{toGoName .TfName}}) > 0 {
		body, _ = sjson.Set(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", []interface{}{})
		for _, item := range data.{{toGoName .TfName}} {
			itemBody := ""
			{{- range .Attributes}}
			{{- if .Value}}
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
			{{- else if not .Reference}}
			{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
			if !item.{{toGoName .TfName}}.IsNull() {
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
					{{- if .Value}}
					itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
					{{- else if not .Reference}}
					{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
					if !childItem.{{toGoName .TfName}}.IsNull() {
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
							{{- if .Value}}
							itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
							{{- else if not .Reference}}
							{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
							if !childChildItem.{{toGoName .TfName}}.IsNull() {
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
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *{{camelCase .Name}}) fromBody(ctx context.Context, res gjson.Result) {
	{{- range .Attributes}}
	{{- if and (not .Value) (not .WriteOnly) (not .Reference)}}
	{{- $cname := toGoName .TfName}}
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
		data.{{toGoName .TfName}} = make([]{{$name}}{{toGoName .TfName}}, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := {{$name}}{{toGoName .TfName}}{}
			{{- range .Attributes}}
			{{- $ccname := toGoName .TfName}}
			{{- if and (not .Value) (not .WriteOnly) (not .Reference)}}
			{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); cValue.Exists() {
				item.{{toGoName .TfName}} = types.{{.Type}}Value(cValue.{{if eq .Type "Int64"}}Int{{else if eq .Type "Float64"}}Float{{else}}{{.Type}}{{end}}())
			} else {
				{{- if .DefaultValue}}
				item.{{toGoName .TfName}} = types.{{.Type}}Value({{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}})
				{{- else}}
				item.{{toGoName .TfName}} = types.{{.Type}}Null()
				{{- end}}
			}
			{{- else if isListSet .}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); cValue.Exists() {
				item.{{toGoName .TfName}} = helpers.Get{{.ElementType}}{{.Type}}(cValue.Array())
			} else {
				item.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
			}
			{{- else if isNestedListSet .}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); cValue.Exists() {
				item.{{toGoName .TfName}} = make([]{{$name}}{{$cname}}{{toGoName .TfName}}, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := {{$name}}{{$cname}}{{toGoName .TfName}}{}
					{{- range .Attributes}}
					{{- if and (not .Value) (not .WriteOnly) (not .Reference)}}
					{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); ccValue.Exists() {
						cItem.{{toGoName .TfName}} = types.{{.Type}}Value(ccValue.{{if eq .Type "Int64"}}Int{{else if eq .Type "Float64"}}Float{{else}}{{.Type}}{{end}}())
					} else {
						{{- if .DefaultValue}}
						cItem.{{toGoName .TfName}} = types.{{.Type}}Value({{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}})
						{{- else}}
						cItem.{{toGoName .TfName}} = types.{{.Type}}Null()
						{{- end}}
					}
					{{- else if isListSet .}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); ccValue.Exists() {
						cItem.{{toGoName .TfName}} = helpers.Get{{.ElementType}}{{.Type}}(ccValue.Array())
					} else {
						cItem.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
					}
					{{- else if isNestedListSet .}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); ccValue.Exists() {
						cItem.{{toGoName .TfName}} = make([]{{$name}}{{$cname}}{{$ccname}}{{toGoName .TfName}}, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := {{$name}}{{$cname}}{{$ccname}}{{toGoName .TfName}}{}
							{{- range .Attributes}}
							{{- if and (not .Value) (not .WriteOnly) (not .Reference)}}
							{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
							if cccValue := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); cccValue.Exists() {
								ccItem.{{toGoName .TfName}} = types.{{.Type}}Value(cccValue.{{if eq .Type "Int64"}}Int{{else if eq .Type "Float64"}}Float{{else}}{{.Type}}{{end}}())
							} else {
								{{- if .DefaultValue}}
								ccItem.{{toGoName .TfName}} = types.{{.Type}}Value({{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}})
								{{- else}}
								ccItem.{{toGoName .TfName}} = types.{{.Type}}Null()
								{{- end}}
							}
							{{- else if isListSet .}}
							if cccValue := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); cccValue.Exists() {
								ccItem.{{toGoName .TfName}} = helpers.Get{{.ElementType}}{{.Type}}(cccValue.Array())
							} else {
								ccItem.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
							}
							{{- end}}
							{{- end}}
							{{- end}}
							cItem.{{toGoName .TfName}} = append(cItem.{{toGoName .TfName}}, ccItem)
							return true
						})
					}
					{{- end}}
					{{- end}}
					{{- end}}
					item.{{toGoName .TfName}} = append(item.{{toGoName .TfName}}, cItem)
					return true
				})
			}
			{{- end}}
			{{- end}}
			{{- end}}
			data.{{toGoName .TfName}} = append(data.{{toGoName .TfName}}, item)
			return true
		})
	}
	{{- end}}
	{{- end}}
	{{- end}}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *{{camelCase .Name}}) fromBodyPartial(ctx context.Context, res gjson.Result) {
	{{- range .Attributes}}
	{{- if and (not .Value) (not .WriteOnly) (not .Reference)}}
	{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
	if value := res.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists(){{if not .ResourceId}} && !data.{{toGoName .TfName}}.IsNull(){{end}} {
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
	{{- else if isNestedListSet .}}
	{{- $list := (toGoName .TfName)}}
	{{- if .OrderedList }}
	{{- if (hasId .Attributes)}}
	{{- errorf "ordered_list %s must not contain elements with `id: true`, as it treats list index ([i]) as the only unique id" .TfName}}
	{{- end}}
	{
		l := len(res.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}").Array())
		tflog.Debug(ctx, fmt.Sprintf("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}} array resizing from %d to %d", len(data.{{toGoName .TfName}}), l))
		for i := len(data.{{toGoName .TfName}}); i < l; i++ {
			data.{{toGoName .TfName}} = append(data.{{toGoName .TfName}}, {{$name}}{{toGoName .TfName}}{})
		}
		if len(data.{{toGoName .TfName}}) > l {
			data.{{toGoName .TfName}} = data.{{toGoName .TfName}}[:l]
		}
	}
	for i := range data.{{toGoName .TfName}} {
		r := res.Get(fmt.Sprintf("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.%d", i))
	{{- else }}
	for i := 0; i < len(data.{{toGoName .TfName}}); i++ {
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
		if !r.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing data.{{toGoName .TfName}}[%d] = %+v",
				i,
				data.{{toGoName .TfName}}[i],
			))
			data.{{toGoName .TfName}} = slices.Delete(data.{{toGoName .TfName}}, i, i+1)
			i--

			continue
		}
	{{- end}}

		{{- range .Attributes}}
		{{- if and (not .Value) (not .WriteOnly) (not .Reference)}}
		{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
		if value := r.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() && !data.{{$list}}[i].{{toGoName .TfName}}.IsNull() {
			data.{{$list}}[i].{{toGoName .TfName}} = types.{{.Type}}Value(value.{{if eq .Type "Int64"}}Int{{else if eq .Type "Float64"}}Float{{else}}{{.Type}}{{end}}())
		} else {{if .DefaultValue}}if data.{{$list}}[i].{{toGoName .TfName}}.Value{{.Type}}() != {{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}} {{end}}{
			data.{{$list}}[i].{{toGoName .TfName}} = types.{{.Type}}Null()
		}
		{{- else if isListSet .}}
		if value := r.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() && !data.{{$list}}[i].{{toGoName .TfName}}.IsNull() {
			data.{{$list}}[i].{{toGoName .TfName}} = helpers.Get{{.ElementType}}{{.Type}}(value.Array())
		} else {
			data.{{$list}}[i].{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
		}
		{{- else if isNestedListSet .}}
		{{- $clist := (toGoName .TfName)}}
		{{- if .OrderedList }}
		{{- errorf "ordered_list cannot be nested that deep"}}
		{{- end}}
		for ci := 0; ci < len(data.{{$list}}[i].{{toGoName .TfName}}); ci++ {
			keys := [...]string{ {{$noId := not (hasId .Attributes)}}{{range .Attributes}}{{if or .Id (and $noId (not .Value))}}{{if or (eq .Type "Int64") (eq .Type "Bool") (eq .Type "String")}}"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{end}}{{end}}{{end}} }
			keyValues := [...]string{ {{$noId := not (hasId .Attributes)}}{{range .Attributes}}{{if or .Id (and $noId (not .Value))}}{{if eq .Type "Int64"}}strconv.FormatInt(data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.ValueInt64(), 10), {{else if eq .Type "Bool"}}strconv.FormatBool(data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.ValueBool()), {{else if eq .Type "String"}}data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.Value{{.Type}}(), {{end}}{{end}}{{end}} }

			var cr gjson.Result
			r.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}").ForEach(
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
						cr = v
						return false
					}
					return true
				},
			)
			if !cr.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing data.{{$list}}[i].{{toGoName .TfName}}[%d] = %+v",
					ci,
					data.{{$list}}[i].{{toGoName .TfName}}[ci],
				))
				data.{{$list}}[i].{{toGoName .TfName}} = slices.Delete(data.{{$list}}[i].{{toGoName .TfName}}, ci, ci+1)
				ci--

				continue
			}

			{{- range .Attributes}}
			{{- if and (not .Value) (not .WriteOnly) (not .Reference)}}
			{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
			if value := cr.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() && !data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.IsNull() {
				data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}} = types.{{.Type}}Value(value.{{if eq .Type "Int64"}}Int{{else if eq .Type "Float64"}}Float{{else}}{{.Type}}{{end}}())
			} else {{if .DefaultValue}}if data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.Value{{.Type}}() != {{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}} {{end}}{
				data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}} = types.{{.Type}}Null()
			}
			{{- else if isListSet .}}
			if value := cr.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() && !data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.IsNull() {
				data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}} = helpers.Get{{.ElementType}}{{.Type}}(value.Array())
			} else {
				data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
			}
			{{- else if isNestedListSet .}}
			{{- $cclist := (toGoName .TfName)}}
			{{- if .OrderedList }}
			{{- errorf "ordered_list cannot be nested that deep"}}
			{{- end}}
			for cci := 0; cci < len(data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}); cci++ {
				keys := [...]string{ {{$noId := not (hasId .Attributes)}}{{range .Attributes}}{{if or .Id (and $noId (not .Value))}}{{if or (eq .Type "Int64") (eq .Type "Bool") (eq .Type "String")}}"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{end}}{{end}}{{end}} }
				keyValues := [...]string{ {{$noId := not (hasId .Attributes)}}{{range .Attributes}}{{if or .Id (and $noId (not .Value))}}{{if eq .Type "Int64"}}strconv.FormatInt(data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}}.ValueInt64(), 10), {{else if eq .Type "Bool"}}strconv.FormatBool(data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}}.ValueBool()), {{else if eq .Type "String"}}data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}}.Value{{.Type}}(), {{end}}{{end}}{{end}} }

				var ccr gjson.Result
				cr.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}").ForEach(
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
							ccr = v
							return false
						}
						return true
					},
				)
				if !cr.Exists() {
					tflog.Debug(ctx, fmt.Sprintf("removing data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}[%d] = %+v",
						cci,
						data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}[cci],
					))
					data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}} = slices.Delete(data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}, cci, cci+1)
					cci--

					continue
				}

				{{- range .Attributes}}
				{{- if and (not .Value) (not .WriteOnly) (not .Reference)}}
				{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
				if value := ccr.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() && !data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}}.IsNull() {
					data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}} = types.{{.Type}}Value(value.{{if eq .Type "Int64"}}Int{{else if eq .Type "Float64"}}Float{{else}}{{.Type}}{{end}}())
				} else {{if .DefaultValue}}if data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}}.Value{{.Type}}() != {{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}} {{end}}{
					data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}} = types.{{.Type}}Null()
				}
				{{- else if isListSet .}}
				if value := ccr.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() && !data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}}.IsNull() {
					data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}} = helpers.Get{{.ElementType}}{{.Type}}(value.Array())
				} else {
					data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
				}
				{{- end}}
				{{- end}}
				{{- end}}
			}

			{{- end}}
			{{- end}}
			{{- end}}
		}
		{{- end}}
		{{- end}}
		{{- end}}
	}
	{{- end}}
	{{- end}}
	{{- end}}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *{{camelCase .Name}}) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	{{- range .Attributes}}
	{{- if and .ResourceId (not .Reference)}}
	{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
	if data.{{toGoName .TfName}}.IsUnknown() {
		if value := res.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() {
			data.{{toGoName .TfName}} = types.{{.Type}}Value(value.{{if eq .Type "Int64"}}Int{{else if eq .Type "Float64"}}Float{{else}}{{.Type}}{{end}}())
		} else {
			data.{{toGoName .TfName}} = types.{{.Type}}Null()
		}
	}
	{{- else}}
	{{- errorf "resource_id is not yet implemented for type %v" .Type}}
	{{- end}}
	{{- else if isNestedListSet .}}
	{{- if hasResourceId .Attributes}}
	{{- $list := (toGoName .TfName)}}
	{{- if .OrderedList }}
	for i := range data.{{toGoName .TfName}} {
		r := res.Get(fmt.Sprintf("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.%d", i))
	{{- else }}
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
	{{- end}}

		{{- range .Attributes}}
		{{- if and .ResourceId (not .Reference)}}
		{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
		if data.{{$list}}[i].{{toGoName .TfName}}.IsUnknown() {
			if value := r.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() {
				data.{{$list}}[i].{{toGoName .TfName}} = types.{{.Type}}Value(value.{{if eq .Type "Int64"}}Int{{else if eq .Type "Float64"}}Float{{else}}{{.Type}}{{end}}())
			} else {
				data.{{$list}}[i].{{toGoName .TfName}} = types.{{.Type}}Null()
			}
		}
		{{- else}}
		{{- errorf "resource_id is not yet implemented for type %v" .Type}}
		{{- end}}
		{{- else if isNestedListSet .}}
		{{- if hasResourceId .Attributes}}
		{{- errorf "nesting resource_id so deep is not yet implemented"}}
		{{- end}}
		{{- end}}
		{{- end}}
	}
	{{- end}}
	{{- end}}
	{{- end}}
}

// End of section. //template:end fromBodyUnknowns
