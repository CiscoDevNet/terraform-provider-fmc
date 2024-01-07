resource "fmc_{{snakeCase .Name}}" "example" {
{{- range  .Attributes}}
{{- if and (not .ExcludeTest) (not .ExcludeExample) (not .Value) (not .ResourceId)}}
{{- if or (eq .Type "List") (eq .Type "Set")}}
  {{.TfName}} = [
    {
      {{- range  .Attributes}}
      {{- if and (not .ExcludeTest) (not .ExcludeExample) (not .Value)}}
      {{- if or (eq .Type "List") (eq .Type "Set")}}
        {{.TfName}} = [
          {
          {{- range  .Attributes}}
          {{- if and (not .ExcludeTest) (not .ExcludeExample) (not .Value)}}
          {{- if or (eq .Type "List") (eq .Type "Set")}}
            {{.TfName}} = [
              {
                {{- range  .Attributes}}
                {{- if and (not .ExcludeTest) (not .ExcludeExample) (not .Value)}}
                {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if eq .Type "StringList"}}["{{.Example}}"]{{else}}{{.Example}}{{end}}
                {{- end}}
                {{- end}}
              }
            ]
          {{- else}}
          {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if eq .Type "StringList"}}["{{.Example}}"]{{else}}{{.Example}}{{end}}
          {{- end}}
          {{- end}}
          {{- end}}
          }
        ]
      {{- else}}
      {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if eq .Type "StringList"}}["{{.Example}}"]{{else}}{{.Example}}{{end}}
      {{- end}}
      {{- end}}
      {{- end}}
    }
  ]
{{- else}}
  {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if eq .Type "StringList"}}["{{.Example}}"]{{else}}{{.Example}}{{end}}
{{- end}}
{{- end}}
{{- end}}
}
