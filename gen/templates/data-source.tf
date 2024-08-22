{{- $name := .Name -}}
data "fmc_{{snakeCase .Name}}" "example" {
  id = "{{$id := false}}{{range .Attributes}}{{if .Id}}{{$id = true}}{{.Example}}{{end}}{{end}}{{if not $id}}76d24097-41c4-4558-a4d0-a8c07ac08470{{end}}"
  {{- range  .Attributes}}
  {{- if .Reference}}
  {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}
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
