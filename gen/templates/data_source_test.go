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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource

func TestAccDataSourceFmc{{camelCase .Name}}(t *testing.T) {
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} && {{end}}os.Getenv("{{$e}}") == ""{{end}} {
        t.Skip("skipping test, set environment variable {{range $i, $e := .TestTags}}{{if $i}} or {{end}}{{$e}}{{end}}")
	}
	{{- end}}
	var checks []resource.TestCheckFunc
	{{- $name := .Name }}
	{{- range  .Attributes}}
	{{- if and (not .WriteOnly) (not .ExcludeTest) (not .Value) (not .TestValue) (not .ResourceId)}}
	{{- if isNestedListSet .}}
	{{- $list := .TfName }}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
	{{- end}}
	{{- range  .Attributes}}
	{{- if and (not .WriteOnly) (not .ExcludeTest) (not .Value) (not .TestValue)}}
	{{- if isNestedListSet .}}
	{{- $clist := .TfName }}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
	{{- end}}
	{{- range  .Attributes}}
	{{- if and (not .WriteOnly) (not .ExcludeTest) (not .Value) (not .TestValue)}}
	{{- if isNestedListSet .}}
	{{- $cclist := .TfName }}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
	{{- end}}
	{{- range  .Attributes}}
	{{- if and (not .WriteOnly) (not .ExcludeTest) (not .Value) (not .TestValue) (not (isSet .))}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		checks = append(checks, resource.TestCheckResourceAttr("data.fmc_{{snakeCase $name}}.test", "{{$list}}.0.{{$clist}}.0.{{$cclist}}.0.{{.TfName}}{{if isList .}}.0{{end}}", "{{.Example}}"))
	}
	{{- else}}
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_{{snakeCase $name}}.test", "{{$list}}.0.{{$clist}}.0.{{$cclist}}.0.{{.TfName}}{{if isList .}}.0{{end}}", "{{.Example}}"))
	{{- end}}
	{{- end}}
	{{- end}}
	{{- if len .TestTags}}
	}
	{{- end}}
	{{- else if not (or (isSet .) (isNestedMap .))}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		checks = append(checks, resource.TestCheckResourceAttr("data.fmc_{{snakeCase $name}}.test", "{{$list}}.0.{{$clist}}.0.{{.TfName}}{{if isList .}}.0{{end}}", "{{.Example}}"))
	}
	{{- else}}
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_{{snakeCase $name}}.test", "{{$list}}.0.{{$clist}}.0.{{.TfName}}{{if isList .}}.0{{end}}", "{{.Example}}"))
	{{- end}}
	{{- end}}
	{{- end}}
	{{- end}}
	{{- if len .TestTags}}
	}
	{{- end}}
	{{- else if not (or (isSet .) (isNestedMap .))}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		checks = append(checks, resource.TestCheckResourceAttr("data.fmc_{{snakeCase $name}}.test", "{{$list}}.0.{{.TfName}}{{if isList .}}.0{{end}}", "{{.Example}}"))
	}
	{{- else}}
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_{{snakeCase $name}}.test", "{{$list}}.0.{{.TfName}}{{if isList .}}.0{{end}}", "{{.Example}}"))
	{{- end}}
	{{- end}}
	{{- end}}
	{{- end}}
	{{- if len .TestTags}}
	}
	{{- end}}
	{{- else if not (or (isSet .) (isNestedMap .))}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		checks = append(checks, resource.TestCheckResourceAttr("data.fmc_{{snakeCase $name}}.test", "{{.TfName}}{{if isList .}}.0{{end}}", "{{.Example}}"))
	}
	{{- else}}
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_{{snakeCase $name}}.test", "{{.TfName}}{{if isList .}}.0{{end}}", "{{.Example}}"))
	{{- end}}
	{{- end}}
	{{- end}}
	{{- end}}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: {{if .TestPrerequisites}}testAccDataSourceFmc{{camelCase .Name}}PrerequisitesConfig+{{end}}testAccDataSourceFmc{{camelCase .Name}}Config(),
				Check: resource.ComposeTestCheckFunc(checks...),
			},
			{{- if .DataSourceNameQuery}}
			{
				Config: {{if .TestPrerequisites}}testAccDataSourceFmc{{camelCase .Name}}PrerequisitesConfig+{{end}}testAccNamedDataSourceFmc{{camelCase .Name}}Config(),
				Check: resource.ComposeTestCheckFunc(checks...),
			},
			{{- end}}
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
{{- if .TestPrerequisites}}

const testAccDataSourceFmc{{camelCase .Name}}PrerequisitesConfig = `
{{.TestPrerequisites}}
`

{{- end}}
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmc{{camelCase .Name}}Config() string {
	config := `resource "fmc_{{snakeCase $name}}" "test" {` + "\n"
	{{- range  .Attributes}}
	{{- if and (not .ExcludeTest) (not .Value) (not .ResourceId)}}
	{{- if isNestedListMapSet .}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
	{{- end}}
	{{- if isNestedListSet .}}
	config += `	{{.TfName}} = [{` + "\n"
	{{- else if isNestedMap .}}
	config += `	{{.TfName}} = { "{{.MapKeyExample}}" = {` + "\n"
	{{- end}}
		{{- range  .Attributes}}
		{{- if and (not .ExcludeTest) (not .Value)}}
		{{- if isNestedListSet .}}
		{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		{{- end}}
	config += `		{{.TfName}} = [{` + "\n"
			{{- range  .Attributes}}
			{{- if and (not .ExcludeTest) (not .Value)}}
			{{- if isNestedListSet .}}
			{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
			{{- end}}
	config += `			{{.TfName}} = [{` + "\n"
				{{- range  .Attributes}}
				{{- if and (not .ExcludeTest) (not .Value)}}
				{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		config += `				{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
	}
				{{- else}}
	config += `				{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
				{{- end}}
				{{- end}}
				{{- end}}
	config += `			}]` + "\n"
			{{- if len .TestTags}}
	}
			{{- end}}
			{{- else}}
			{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		config += `			{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
	}
			{{- else}}
	config += `			{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
			{{- end}}
			{{- end}}
			{{- end}}
			{{- end}}
	config += `		}]` + "\n"
		{{- if len .TestTags}}
	}
		{{- end}}
		{{- else}}
		{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		config += `		{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
	}
			{{- else}}
	config += `		{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
		{{- end}}
		{{- end}}
		{{- end}}
		{{- end}}
	{{- if isNestedListSet .}}
	config += `	}]` + "\n"
	{{- else if isNestedMap .}}
	config += `	}}` + "\n"
	{{- end}}
		{{- if len .TestTags}}
	}
		{{- end}}
	{{- else}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		config += `	{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
	}
	{{- else}}
	config += `	{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
	{{- end}}
	{{- end}}
	{{- end}}
	{{- end}}
	config += `}` + "\n"

	config += `
		data "fmc_{{snakeCase .Name}}" "test" {
			id = fmc_{{snakeCase $name}}.test.id
			{{- range  .Attributes}}
			{{- if .Reference}}
			{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}
			{{- else if isNestedMap .}}
			{{- $map := .TfName}}
			{{- $mapkey := .MapKeyExample}}
			{{.TfName}} = {
				"{{.MapKeyExample}}" = {
					{{- range  .Attributes}}
					{{- if .ResourceId}}
					{{.TfName}} = fmc_{{snakeCase $name}}.test.{{$map}}["{{$mapkey}}"].{{.TfName}}
					{{- end}}
					{{- end}}
				}
			}
			{{- end}}
			{{- end}}
		}
	`
	return config
}

{{if .DataSourceNameQuery -}}
func testAccNamedDataSourceFmc{{camelCase .Name}}Config() string {
	config := `resource "fmc_{{snakeCase $name}}" "test" {` + "\n"
	{{- range  .Attributes}}
	{{- if and (not .ExcludeTest) (not .Value) (not .ResourceId)}}
	{{- if isNestedListSet .}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
	{{- end}}
	config += `	{{.TfName}} = [{` + "\n"
		{{- range  .Attributes}}
		{{- if and (not .ExcludeTest) (not .Value)}}
		{{- if isNestedListSet .}}
		{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		{{- end}}
	config += `		{{.TfName}} = [{` + "\n"
			{{- range  .Attributes}}
			{{- if and (not .ExcludeTest) (not .Value)}}
			{{- if isNestedListSet .}}
			{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
			{{- end}}
	config += `			{{.TfName}} = [{` + "\n"
				{{- range  .Attributes}}
				{{- if and (not .ExcludeTest) (not .Value)}}
				{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		config += `				{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
	}
				{{- else}}
	config += `				{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
				{{- end}}
				{{- end}}
				{{- end}}
	config += `			}]` + "\n"
			{{- if len .TestTags}}
	}
			{{- end}}
			{{- else}}
			{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		config += `			{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
	}
			{{- else}}
	config += `			{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
			{{- end}}
			{{- end}}
			{{- end}}
			{{- end}}
	config += `		}]` + "\n"
		{{- if len .TestTags}}
	}
		{{- end}}
		{{- else}}
		{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		config += `		{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
	}
			{{- else}}
	config += `		{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
		{{- end}}
		{{- end}}
		{{- end}}
		{{- end}}
	config += `	}]` + "\n"
		{{- if len .TestTags}}
	}
		{{- end}}
	{{- else}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		config += `	{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
	}
	{{- else}}
	config += `	{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
	{{- end}}
	{{- end}}
	{{- end}}
	{{- end}}
	config += `}` + "\n"

	config += `
		data "fmc_{{snakeCase .Name}}" "test" {
			name = fmc_{{snakeCase $name}}.test.name
			{{- range  .Attributes}}
			{{- if .Reference}}
			{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}
			{{- end}}
			{{- end}}
		}
	`
	return config
}
{{- end}}

// End of section. //template:end testAccDataSourceConfig
