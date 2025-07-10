{{- if .IsBulk -}}
terraform import fmc_{{snakeCase .Name}}.example "<domain>,{{range .Attributes}}{{if .Reference}}<{{.TfName}}>,{{end}}{{end}}[<{{snakeCase .Name}}_name>]"
{{- else -}}
terraform import fmc_{{snakeCase .Name}}.example "{{$id := false}}{{range .Attributes}}{{if .Id}}{{$id = true}}<{{.TfName}}>{{end}}{{end}}{{if not $id}}{{range .Attributes}}{{if .Reference}}<{{.TfName}}>,{{end}}{{end}}<id>{{end}}"
{{- end}}