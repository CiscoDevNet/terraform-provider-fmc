# <domain> is optional. If not provided, `Global` is used implicitly and resource's `domain` attribute is not set.
{{if .RestEndpointVrf -}}
# <vrf_id> is optional.
terraform import fmc_{{snakeCase .Name}}.example "<domain>,<device_id>,<vrf_id>,<id>"
{{- else if .IsBulk -}}
terraform import fmc_{{snakeCase .Name}}.example "<domain>,{{range .Attributes}}{{if .Reference}}<{{.TfName}}>,{{end}}{{end}}[<item1_name>,<item2_name>,...]"
{{else -}}
terraform import fmc_{{snakeCase .Name}}.example "<domain>,{{$id := false}}{{range .Attributes}}{{if .Id}}{{$id = true}}<{{.TfName}}>{{end}}{{end}}{{if not $id}}{{range .Attributes}}{{if .Reference}}<{{.TfName}}>,{{end}}{{end}}<id>{{end}}"
{{- end }}