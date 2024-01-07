data "fmc_{{snakeCase .Name}}" "example" {
  id = "{{$id := false}}{{range .Attributes}}{{if .Id}}{{$id = true}}{{.Example}}{{end}}{{end}}{{if not $id}}76d24097-41c4-4558-a4d0-a8c07ac08470{{end}}"
  {{- range  .Attributes}}
  {{- if .Reference}}
  {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if eq .Type "StringList"}}["{{.Example}}"]{{else}}{{.Example}}{{end}}
  {{- end}}
  {{- end}}
}
