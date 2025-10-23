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

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &{{camelCase .Name}}DataSource{}
	_ datasource.DataSourceWithConfigure = &{{camelCase .Name}}DataSource{}
)

func New{{camelCase .Name}}DataSource() datasource.DataSource {
	return &{{camelCase .Name}}DataSource{}
}

type {{camelCase .Name}}DataSource struct {
	client *fmc.Client
}

func (d *{{camelCase .Name}}DataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_{{snakeCase .Name}}"
}

func (d *{{camelCase .Name}}DataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("{{.DsDescription}}")
		{{- if .DeprecationMessage -}}
		.AddAttributeDescription("{{.DeprecationMessage}}")
		{{- end -}}
		{{- if .MinimumVersion -}}
		.AddMinimumVersionHeaderDescription().AddMinimumVersionDescription("{{.MinimumVersion}}")
		{{- end -}}
		.String,
		{{- if .DeprecationMessage }}
		DeprecationMessage:  helpers.NewAttributeDescription("{{.DeprecationMessage}}").String,
		{{- end}}

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the object",
				{{- if and (not (hasDataSourceQuery .Attributes)) (not .IsBulk) }}
				Required:            true,
				{{- else}}
				{{- if not .IsBulk}}
				Optional:            true,
				{{- end}}
				Computed:            true,
				{{- end}}
			},
			{{- if isDomainDependent .}}
			"domain": schema.StringAttribute{
				MarkdownDescription: "Name of the FMC domain",
				Optional:			true,
			},
			{{- end}}
			{{- if .RestEndpointVrf }}
			"vrf_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the parent VRF.").String,
				Optional:            true,
			},
			{{- end}}
			{{- range .Attributes}}
			{{- if not .Value}}
			"{{.TfName}}": schema.{{if isNestedListMapSet .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if eq .Type "Versions"}}List{{else if eq .Type "Version"}}Int64{{else}}{{.Type}}{{end}}Attribute{
				MarkdownDescription: "{{.Description}}",
				{{- if isListSet .}}
				ElementType:         types.{{.ElementType}}Type,
				{{- end}}
				{{- if .Reference}}
				Required:            true,
				{{- else}}
				{{- if or .DataSourceQuery .DataSourceOptionalParameter }}
				Optional:            true,
				{{- end}}
				{{- if isNestedMap .}}
				Optional:            true,
				{{- end}}
				Computed:            true,
				{{- end}}
				{{- if .Sensitive}}
				Sensitive:           true,
				{{- end}}
				{{- if isNestedListMapSet .}}
				{{- $parentNestedMap := isNestedMap .}}
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						{{- range .Attributes}}
						{{- if not .Value}}
						"{{.TfName}}": schema.{{if isNestedListMapSet .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if eq .Type "Versions"}}List{{else if eq .Type "Version"}}Int64{{else}}{{.Type}}{{end}}Attribute{
							MarkdownDescription: "{{.Description}}",
							{{- if isListSet .}}
							ElementType:         types.{{.ElementType}}Type,
							{{- end}}
							{{- if and .ResourceId $parentNestedMap (not $.IsBulk)}}
							Required:            true,
							{{- else}}
							Computed:            true,
							{{- end}}
							{{- if .Sensitive}}
							Sensitive:           true,
							{{- end}}
							{{- if isNestedListMapSet .}}
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									{{- range .Attributes}}
									{{- if not .Value}}
									"{{.TfName}}": schema.{{if isNestedListMapSet .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if eq .Type "Versions"}}List{{else if eq .Type "Version"}}Int64{{else}}{{.Type}}{{end}}Attribute{
										MarkdownDescription: "{{.Description}}",
										{{- if isListSet .}}
										ElementType:         types.{{.ElementType}}Type,
										{{- end}}
										Computed:            true,
										{{- if .Sensitive}}
										Sensitive:           true,
										{{- end}}
										{{- if isNestedListMapSet .}}
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												{{- range .Attributes}}
												{{- if not .Value}}
												"{{.TfName}}": schema.{{if isNestedListMapSet .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if eq .Type "Versions"}}List{{else if eq .Type "Version"}}Int64{{else}}{{.Type}}{{end}}Attribute{
													MarkdownDescription: "{{.Description}}",
													{{- if isListSet .}}
													ElementType:         types.{{.ElementType}}Type,
													{{- end}}
													Computed:            true,
													{{- if .Sensitive}}
													Sensitive:           true,
													{{- end}}
												},
												{{- end}}
												{{- end}}
											},
										},
										{{- end}}
									},
									{{- end}}
									{{- end}}
								},
							},
							{{- end}}
						},
						{{- end}}
						{{- end}}
					},
				},
				{{- end}}
			},
			{{- end}}
			{{- end}}
		},
	}
}
{{- $dataSourceAttribute := getDataSourceQueryAttribute .}}
{{- if and (hasDataSourceQuery .Attributes) (not .IsBulk)}}
func (d *{{camelCase .Name}}DataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
    return []datasource.ConfigValidator{
        datasourcevalidator.ExactlyOneOf(
            path.MatchRoot("id"),
			path.MatchRoot("{{$dataSourceAttribute.TfName}}"),
        ),
    }
}
{{- end}}

func (d *{{camelCase .Name}}DataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *{{camelCase .Name}}DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	{{- if .MinimumVersion}}

	// Check if FMC client is connected to supports this object
	if d.client.FMCVersionParsed.LessThan(minFMCVersion{{camelCase .Name}}) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support {{.Name}}, minimum required version is {{.MinimumVersion}}", d.client.FMCVersion))
		return
	}
	{{- end}}
	var config {{camelCase .Name}}

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	{{- if isDomainDependent .}}
	if !config.Domain.IsNull() && config.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(config.Domain.ValueString()))
	}
	{{- end}}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	{{- if and (hasDataSourceQuery .Attributes) (not .IsBulk)}}
	if config.Id.IsNull() && !config.{{toGoName $dataSourceAttribute.TfName}}.IsNull() {
		offset := 0
		limit := 1000
		for page := 1; ; page++ {
			queryString := fmt.Sprintf("?limit=%d&offset=%d&expanded=true", limit, offset)
			res, err := d.client.Get(config.getPath() + queryString, reqMods...)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects, got error: %s", err))
				return
			}
			if value := res.Get("items"); len(value.Array()) > 0 {
				value.ForEach(func(k, v gjson.Result) bool {
					if config.{{toGoName $dataSourceAttribute.TfName}}.
						{{- if eq $dataSourceAttribute.Type "Int64" -}}ValueInt64()
						{{- else -}}ValueString(){{- end -}} == v.Get("{{range $dataSourceAttribute.DataPath}}{{.}}.{{end}}{{$dataSourceAttribute.ModelName}}").
						{{- if eq $dataSourceAttribute.Type "Int64" -}}Int()
						{{- else -}}String(){{- end -}} {
						config.Id = types.StringValue(v.Get("id").String())
						tflog.Debug(ctx, fmt.Sprintf("%s: Found object with {{$dataSourceAttribute.TfName}} '%v', id: %v", config.Id.ValueString(), config.{{toGoName $dataSourceAttribute.TfName}}.{{if eq $dataSourceAttribute.Type "Int64"}}ValueInt64(){{else}}ValueString(){{end}}, config.Id.ValueString()))
						return false
					}
					return true
				})
			}
			if !config.Id.IsNull() || !res.Get("paging.next.0").Exists() {
				break
			}
			offset += limit
		}

		if config.Id.IsNull() {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with {{$dataSourceAttribute.TfName}}: %v", config.{{toGoName $dataSourceAttribute.TfName}}.{{if eq $dataSourceAttribute.Type "Int64"}}ValueInt64(){{else}}ValueString(){{end}}))
			return
		}
	}
	{{- end}}

	{{- if .IsBulk}}
	
	// Get all objects from FMC
	urlPath := config.getPath() + "?expanded=true"
	{{- else}}
	urlPath := config.getPath()+"/"+url.QueryEscape(config.Id.ValueString())
	{{- end}}
	res, err := d.client.Get(urlPath, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
