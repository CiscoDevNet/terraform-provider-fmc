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
	"strings"
	"sync"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/go-version"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &{{camelCase .Name}}Resource{}
	{{- if not .NoImport}}
	_ resource.ResourceWithImportState = &{{camelCase .Name}}Resource{}
	{{- end}}
)

func New{{camelCase .Name}}Resource() resource.Resource {
	return &{{camelCase .Name}}Resource{}
}

type {{camelCase .Name}}Resource struct {
	client *fmc.Client
}

func (r *{{camelCase .Name}}Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_{{snakeCase .Name}}"
}

func (r *{{camelCase .Name}}Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("{{.ResDescription}}")
		{{- if .DeprecationMessage -}}
		.AddAttributeDescription("{{.DeprecationMessage}}")
		{{- end -}}
		{{- if or .MinimumVersion .MinimumVersionCreate .MinimumVersionBulkCreate .MinimumVersionBulkDelete .IsBulk -}}
		.AddMinimumVersionHeaderDescription()
		{{- end -}}
		{{- if .MinimumVersion -}}
		.AddMinimumVersionDescription("{{.MinimumVersion}}")
		{{- end -}}
		{{- if and .MinimumVersionCreate (not .MinimumVersion) -}}
		.AddMinimumVersionAnyDescription()
		{{- end -}}
		{{- if .MinimumVersionCreate -}}
		.AddMinimumVersionCreateDescription("{{.MinimumVersionCreate}}")
		{{- end -}}
		{{- if .MinimumVersionBulkCreate -}}
		.AddMinimumVersionBulkCreateDescription("{{.MinimumVersionBulkCreate}}")
		{{- end -}}
		{{- if .MinimumVersionBulkDelete -}}
		.AddMinimumVersionBulkDeleteDescription("{{.MinimumVersionBulkDelete}}")
		{{- end -}}
		{{- if and (or .MinimumVersionBulkCreate .MinimumVersionBulkDelete) (ne .MinimumVersionBulkCreate "999") (ne .MinimumVersionBulkDelete "999")  -}}
		.AddMinimumVersionBulkDisclaimerDescription()
		{{- end -}}
		{{- if .IsBulk -}}
		.AddMinimumVersionBulkUpdateDescription()
		{{- end -}}
		.String,
		{{- if .DeprecationMessage }}
		DeprecationMessage:  helpers.NewAttributeDescription("{{.DeprecationMessage}}").String,
		{{- end}}

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			{{- if isDomainDependent .}}
			"domain": schema.StringAttribute{
				MarkdownDescription: "Name of the FMC domain",
				Optional:			true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			{{- end}}
			{{- range  .Attributes}}
			{{- if not .Value}}
			"{{.TfName}}": schema.{{if isNestedListMapSet .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if eq .Type "Versions"}}List{{else if eq .Type "Version"}}Int64{{else}}{{.Type}}{{end}}Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("{{.Description}}")
					{{- if len .EnumValues -}}
					.AddStringEnumDescription({{range .EnumValues}}"{{.}}", {{end}})
					{{- end -}}
					{{- if or (ne .MinInt 0) (ne .MaxInt 0) -}}
					.AddIntegerRangeDescription({{.MinInt}}, {{.MaxInt}})
					{{- end -}}
					{{- if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0) -}}
					.AddFloatRangeDescription({{.MinFloat}}, {{.MaxFloat}})
					{{- end -}}
					{{- if .DefaultValue -}}
					.AddDefaultValueDescription("{{.DefaultValue}}")
					{{- end -}}
					.String,
				{{- if isListSet .}}
				ElementType:         types.{{.ElementType}}Type,
				{{- end}}
				{{- if or .Reference .Mandatory}}
				Required:            true,
				{{- else if and (not .ResourceId) (not .Computed)}}
				Optional:            true,
				{{- end}}
				{{- if or (len .DefaultValue) .ResourceId .Computed}}
				Computed:            true,
				{{- end}}
				{{- if .Sensitive}}
				Sensitive:           true,
				{{- end}}
				{{- if len .EnumValues}}
				{{- if isSet .}}
				Validators: []validator.Set{
					{{- if eq .ElementType "String"}}
					setvalidator.ValueStringsAre(
						stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
					),
					{{- end}}
				},
				{{- else}}
				Validators: []validator.String{
					stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
				},
				{{- end}}
				{{- else if or (len .StringPatterns) (ne .StringMinLength 0) (ne .StringMaxLength 0) }}
				Validators: []validator.String{
					{{- if or (ne .StringMinLength 0) (ne .StringMaxLength 0)}}
					stringvalidator.LengthBetween({{.StringMinLength}}, {{.StringMaxLength}}),
					{{- end}}
					{{- range .StringPatterns}}
					stringvalidator.RegexMatches(regexp.MustCompile(`{{.}}`), ""),
					{{- end}}
				},
				{{- else if or (ne .MinInt 0) (ne .MaxInt 0)}}
				Validators: []validator.Int64{
					int64validator.Between({{.MinInt}}, {{.MaxInt}}),
				},
				{{- else if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0)}}
				Validators: []validator.Float64{
					float64validator.Between({{.MinFloat}}, {{.MaxFloat}}),
				},
				{{- end}}
				{{- if and (len .DefaultValue) (eq .Type "Int64")}}
				Default:             int64default.StaticInt64({{.DefaultValue}}),
				{{- else if and (len .DefaultValue) (eq .Type "Bool")}}
				Default:             booldefault.StaticBool({{.DefaultValue}}),
				{{- else if and (len .DefaultValue) (eq .Type "String")}}
				Default:             stringdefault.StaticString("{{.DefaultValue}}"),
				{{- end}}
				{{- if or .Id .Reference .RequiresReplace (and .Computed (not .ComputedRefreshValue))}}
				PlanModifiers: []planmodifier.{{.Type}}{
					{{- if or .Id .Reference (and .RequiresReplace (not $.IsBulk))}}
					{{snakeCase .Type}}planmodifier.RequiresReplace(),
					{{end}}
					{{- if and .Computed (not .ComputedRefreshValue)}}
					{{snakeCase .Type}}planmodifier.UseStateForUnknown(),
					{{end}}
				},
				{{- end}}
				{{- if isNestedListMapSet .}}
				{{- $useStateForUnknown := isNestedMap .}}
				{{- $itemsList := .}}
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						{{- range  .Attributes}}
						{{- if not .Value}}
						"{{.TfName}}": schema.{{if isNestedListMapSet .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if eq .Type "Versions"}}List{{else if eq .Type "Version"}}Int64{{else}}{{.Type}}{{end}}Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("{{.Description}}")
								{{- if len .EnumValues -}}
								.AddStringEnumDescription({{range .EnumValues}}"{{.}}", {{end}})
								{{- end -}}
								{{- if or (ne .MinInt 0) (ne .MaxInt 0) -}}
								.AddIntegerRangeDescription({{.MinInt}}, {{.MaxInt}})
								{{- end -}}
								{{- if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0) -}}
								.AddFloatRangeDescription({{.MinFloat}}, {{.MaxFloat}})
								{{- end -}}
								{{- if .DefaultValue -}}
								.AddDefaultValueDescription("{{.DefaultValue}}")
								{{- end -}}
								.String,
							{{- if isListSet .}}
							ElementType:         types.{{.ElementType}}Type,
							{{- end}}
							{{- if or .Reference .Mandatory}}
							Required:            true,
							{{- else if and (not .ResourceId) (not .Computed)}}
							Optional:            true,
							{{- end}}
							{{- if or (len .DefaultValue) .ResourceId .Computed}}
							Computed:            true,
							{{- end}}
							{{- if .Sensitive}}
							Sensitive:           true,
							{{- end}}
							{{- if len .EnumValues}}
							{{- if isSet .}}
							Validators: []validator.Set{
								{{- if eq .ElementType "String"}}
								setvalidator.ValueStringsAre(
									stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
								),
								{{- end}}
							},
							{{- else}}
							Validators: []validator.String{
								stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
							},
							{{- end}}
							{{- else if or (len .StringPatterns) (ne .StringMinLength 0) (ne .StringMaxLength 0) }}
							Validators: []validator.String{
								{{- if or (ne .StringMinLength 0) (ne .StringMaxLength 0)}}
								stringvalidator.LengthBetween({{.StringMinLength}}, {{.StringMaxLength}}),
								{{- end}}
								{{- range .StringPatterns}}
								stringvalidator.RegexMatches(regexp.MustCompile(`{{.}}`), ""),
								{{- end}}
							},
							{{- else if or (ne .MinInt 0) (ne .MaxInt 0)}}
							Validators: []validator.Int64{
								int64validator.Between({{.MinInt}}, {{.MaxInt}}),
							},
							{{- else if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0)}}
							Validators: []validator.Float64{
								float64validator.Between({{.MinFloat}}, {{.MaxFloat}}),
							},
							{{- end}}
							{{- if and (len .DefaultValue) (eq .Type "Int64")}}
							Default:             int64default.StaticInt64({{.DefaultValue}}),
							{{- else if and (len .DefaultValue) (eq .Type "Bool")}}
							Default:             booldefault.StaticBool({{.DefaultValue}}),
							{{- else if and (len .DefaultValue) (eq .Type "String")}}
							Default:             stringdefault.StaticString("{{.DefaultValue}}"),
							{{- end}}
							{{- if or (and .ResourceId $useStateForUnknown) (and .Computed (not .ComputedRefreshValue))}}
							PlanModifiers: []planmodifier.{{.Type}}{
								{{- if and $.IsBulk (eq .TfName "id") (hasRequiresReplace $.Attributes) }}
								{{- if eq .Type "String"}}
								{{- range $itemsList.Attributes }}
								{{- if .RequiresReplace}}
								planmodifiers.ConditionalUseStateForUnknownString("{{.TfName}}"),
								{{- end}}
								{{- end }}
								{{- end}}
								{{- else}}
								planmodifiers.UseStateForUnknownKeepNonNullState{{toGoName .Type}}(),
								{{- end}}
							},
							{{- else if and .RequiresReplace (not $.IsBulk)}}
							PlanModifiers: []planmodifier.{{.Type}}{
								{{snakeCase .Type}}planmodifier.RequiresReplace(),
							},
							{{- end}}
							{{- if isNestedListMapSet .}}
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									{{- range  .Attributes}}
									{{- if not .Value}}
									"{{.TfName}}": schema.{{if isNestedListMapSet .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if eq .Type "Versions"}}List{{else if eq .Type "Version"}}Int64{{else}}{{.Type}}{{end}}Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("{{.Description}}")
											{{- if len .EnumValues -}}
											.AddStringEnumDescription({{range .EnumValues}}"{{.}}", {{end}})
											{{- end -}}
											{{- if or (ne .MinInt 0) (ne .MaxInt 0) -}}
											.AddIntegerRangeDescription({{.MinInt}}, {{.MaxInt}})
											{{- end -}}
											{{- if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0) -}}
											.AddFloatRangeDescription({{.MinFloat}}, {{.MaxFloat}})
											{{- end -}}
											{{- if .DefaultValue -}}
											.AddDefaultValueDescription("{{.DefaultValue}}")
											{{- end -}}
											.String,
										{{- if isListSet .}}
										ElementType:         types.{{.ElementType}}Type,
										{{- end}}
										{{- if or .Reference .Mandatory}}
										Required:            true,
										{{- else if and (not .ResourceId) (not .Computed)}}
										Optional:            true,
										{{- end}}
										{{- if or (len .DefaultValue) .ResourceId .Computed}}
										Computed:            true,
										{{- end}}
										{{- if len .EnumValues}}
										{{- if isSet .}}
										Validators: []validator.Set{
											{{- if eq .ElementType "String"}}
											setvalidator.ValueStringsAre(
												stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
											),
											{{- end}}
										},
										{{- else}}
										Validators: []validator.String{
											stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
										},
										{{- end}}
										{{- else if or (len .StringPatterns) (ne .StringMinLength 0) (ne .StringMaxLength 0) }}
										Validators: []validator.String{
											{{- if or (ne .StringMinLength 0) (ne .StringMaxLength 0)}}
											stringvalidator.LengthBetween({{.StringMinLength}}, {{.StringMaxLength}}),
											{{- end}}
											{{- range .StringPatterns}}
											stringvalidator.RegexMatches(regexp.MustCompile(`{{.}}`), ""),
											{{- end}}
										},
										{{- else if or (ne .MinInt 0) (ne .MaxInt 0)}}
										Validators: []validator.Int64{
											int64validator.Between({{.MinInt}}, {{.MaxInt}}),
										},
										{{- else if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0)}}
										Validators: []validator.Float64{
											float64validator.Between({{.MinFloat}}, {{.MaxFloat}}),
										},
										{{- end}}
										{{- if and (len .DefaultValue) (eq .Type "Int64")}}
										Default:             int64default.StaticInt64({{.DefaultValue}}),
										{{- else if and (len .DefaultValue) (eq .Type "Bool")}}
										Default:             booldefault.StaticBool({{.DefaultValue}}),
										{{- else if and (len .DefaultValue) (eq .Type "String")}}
										Default:             stringdefault.StaticString("{{.DefaultValue}}"),
										{{- end}}
										{{- if or .RequiresReplace .Computed}}
										PlanModifiers: []planmodifier.{{.Type}}{
											{{- if and .RequiresReplace (not $.IsBulk)}}
											{{snakeCase .Type}}planmodifier.RequiresReplace(),
											{{end}}
											{{- if and .Computed (not .ComputedRefreshValue)}}
											planmodifiers.UseStateForUnknownKeepNonNullState{{toGoName .Type}}(),
											{{end}}
										},
										{{- end}}
										{{- if .Sensitive}}
										Sensitive:           true,
										{{- end}}
										{{- if isNestedListMapSet .}}
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												{{- range  .Attributes}}
												{{- if not .Value}}
												"{{.TfName}}": schema.{{if isNestedListMapSet .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if eq .Type "Versions"}}List{{else if eq .Type "Version"}}Int64{{else}}{{.Type}}{{end}}Attribute{
													MarkdownDescription: helpers.NewAttributeDescription("{{.Description}}")
														{{- if len .EnumValues -}}
														.AddStringEnumDescription({{range .EnumValues}}"{{.}}", {{end}})
														{{- end -}}
														{{- if or (ne .MinInt 0) (ne .MaxInt 0) -}}
														.AddIntegerRangeDescription({{.MinInt}}, {{.MaxInt}})
														{{- end -}}
														{{- if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0) -}}
														.AddFloatRangeDescription({{.MinFloat}}, {{.MaxFloat}})
														{{- end -}}
														{{- if .DefaultValue -}}
														.AddDefaultValueDescription("{{.DefaultValue}}")
														{{- end -}}
														.String,
													{{- if isListSet .}}
													ElementType:         types.{{.ElementType}}Type,
													{{- end}}
													{{- if or .Reference .Mandatory}}
													Required:            true,
													{{- else if and (not .ResourceId) (not .Computed)}}
													Optional:            true,
													{{- end}}
													{{- if or (len .DefaultValue) .ResourceId .Computed}}
													Computed:            true,
													{{- end}}
													{{- if .Sensitive}}
													Sensitive:           true,
													{{- end}}
													{{- if len .EnumValues}}
													{{- if isSet .}}
													Validators: []validator.Set{
														{{- if eq .ElementType "String"}}
														setvalidator.ValueStringsAre(
															stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
														),
														{{- end}}
													},
													{{- else}}
													Validators: []validator.String{
														stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
													},
													{{- end}}
													{{- else if or (len .StringPatterns) (ne .StringMinLength 0) (ne .StringMaxLength 0) }}
													Validators: []validator.String{
														{{- if or (ne .StringMinLength 0) (ne .StringMaxLength 0)}}
														stringvalidator.LengthBetween({{.StringMinLength}}, {{.StringMaxLength}}),
														{{- end}}
														{{- range .StringPatterns}}
														stringvalidator.RegexMatches(regexp.MustCompile(`{{.}}`), ""),
														{{- end}}
													},
													{{- else if or (ne .MinInt 0) (ne .MaxInt 0)}}
													Validators: []validator.Int64{
														int64validator.Between({{.MinInt}}, {{.MaxInt}}),
													},
													{{- else if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0)}}
													Validators: []validator.Float64{
														float64validator.Between({{.MinFloat}}, {{.MaxFloat}}),
													},
													{{- end}}
													{{- if and (len .DefaultValue) (eq .Type "Int64")}}
													Default:             int64default.StaticInt64({{.DefaultValue}}),
													{{- else if and (len .DefaultValue) (eq .Type "Bool")}}
													Default:             booldefault.StaticBool({{.DefaultValue}}),
													{{- else if and (len .DefaultValue) (eq .Type "String")}}
													Default:             stringdefault.StaticString("{{.DefaultValue}}"),
													{{- end}}
													{{- if or .RequiresReplace .Computed}}
													PlanModifiers: []planmodifier.{{.Type}}{
														{{- if and .RequiresReplace (not $.IsBulk)}}
														{{snakeCase .Type}}planmodifier.RequiresReplace(),
														{{end}}
														{{- if and .Computed (not .ComputedRefreshValue)}}
														planmodifiers.UseStateForUnknownKeepNonNullState{{toGoName .Type}}(),
														{{end}}
													},
													{{- end}}
												},
												{{- end}}
												{{- end}}
											},
										},
										{{- if or (ne .MinList 0) (ne .MaxList 0)}}
										Validators: []validator.List{
											{{- if ne .MinList 0}}
											listvalidator.SizeAtLeast({{.MinList}}),
											{{- end}}
											{{- if ne .MaxList 0}}
											listvalidator.SizeAtMost({{.MaxList}}),
											{{- end}}
										},
										{{- end}}
										{{- end}}
									},
									{{- end}}
									{{- end}}
								},
							},
							{{- if or (ne .MinList 0) (ne .MaxList 0)}}
							Validators: []validator.List{
								{{- if ne .MinList 0}}
								listvalidator.SizeAtLeast({{.MinList}}),
								{{- end}}
								{{- if ne .MaxList 0}}
								listvalidator.SizeAtMost({{.MaxList}}),
								{{- end}}
							},
							{{- end}}
							{{- end}}
						},
						{{- end}}
						{{- end}}
					},
				},
				{{- if or (ne .MinList 0) (ne .MaxList 0)}}
				Validators: []validator.List{
					{{- if ne .MinList 0}}
					listvalidator.SizeAtLeast({{.MinList}}),
					{{- end}}
					{{- if ne .MaxList 0}}
					listvalidator.SizeAtMost({{.MaxList}}),
					{{- end}}
				},
				{{- end}}
				{{- end}}
			},
			{{- end}}
			{{- end}}
		},
	}
}

func (r *{{camelCase .Name}}Resource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *{{camelCase .Name}}Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	{{- if or .MinimumVersion .MinimumVersionCreate}}
	// Get FMC version
	fmcVersion, _ := version.NewVersion(strings.Split(r.client.FMCVersion, " ")[0])

	// Check if FMC client is connected to supports this object
	if fmcVersion.LessThan(minFMCVersion{{if .MinimumVersionCreate}}Create{{end}}{{camelCase .Name}}) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support {{.Name}} creation, minumum required version is {{if .MinimumVersionCreate}}{{.MinimumVersionCreate}}{{else}}{{.MinimumVersion}}{{end}}", r.client.FMCVersion))
		return
	}
	{{- end}}
	var plan {{camelCase .Name}}

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	{{- if isDomainDependent .}}
	if !plan.Domain.IsNull() && plan.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(plan.Domain.ValueString()))
	}
	{{- end}}

	{{- if and .PutCreate (not .IsBulk)}}

	{{- if .RetrieveId}}
	//// ID needs to be retrieved from FMC, however we are expecting exactly one object
	// Get objects from FMC
	resId, err := r.client.Get(plan.getPath(), reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	// Check if exactly one object is returned
	val := resId.Get("items").Array()
	if len(val) != 1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Expected 1 object, got %d", len(val)))
		return
	}

	// Extract ID from the object
	if retrievedId := val[0].Get("id"); retrievedId.Exists() {
		plan.Id = types.StringValue(retrievedId.String())
		tflog.Debug(ctx, fmt.Sprintf("%s: Found object", plan.Id))
	} else {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object id from payload: %s", resId.String()))
	}

	{{- else}}
	{{- $dataSourceAttribute := getAttributeByTfName .Attributes "name"}}
	{{- if hasDataSourceQuery .Attributes}}
	{{- $dataSourceAttribute = getDataSourceQueryAttribute . }}
	{{- end}}
	
	tflog.Debug(ctx, fmt.Sprintf("%s: considering object {{$dataSourceAttribute.TfName}} %s", plan.Id, plan.{{toGoName $dataSourceAttribute.TfName}}))
	if plan.Id.ValueString() == "" && plan.{{toGoName $dataSourceAttribute.TfName}}.ValueString() != "" {
		offset := 0
		limit := 1000
		for page := 1; ; page++ {
			queryString := fmt.Sprintf("?limit=%d&offset=%d&expanded=true", limit, offset)
			res, err := r.client.Get(plan.getPath()+queryString, reqMods...)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects, got error: %s", err))
				return
			}
			if value := res.Get("items"); len(value.Array()) > 0 {
				value.ForEach(func(k, v gjson.Result) bool {
					if plan.{{toGoName $dataSourceAttribute.TfName}}.
						{{- if eq $dataSourceAttribute.Type "Int64" -}}ValueInt64()
						{{- else -}}ValueString(){{- end -}} == v.Get("{{range $dataSourceAttribute.DataPath}}{{.}}.{{end}}{{$dataSourceAttribute.ModelName}}").
						{{- if eq $dataSourceAttribute.Type "Int64" -}}Int()
						{{- else -}}String(){{- end -}} {
						plan.Id = types.StringValue(v.Get("id").String())
						tflog.Debug(ctx, fmt.Sprintf("%s: Found object with {{$dataSourceAttribute.TfName}} '%v', id: %s", plan.Id.ValueString(), plan.{{toGoName $dataSourceAttribute.TfName}}.{{if eq $dataSourceAttribute.Type "Int64"}}ValueInt64(){{else}}ValueString(){{end}}, plan.Id.ValueString()))
						return false
					}
					return true
				})
			}
			if plan.Id.ValueString() != "" || !res.Get("paging.next.0").Exists() {
				break
			}
			offset += limit
		}

		if plan.Id.ValueString() == "" {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with {{$dataSourceAttribute.TfName}}: %v", plan.{{toGoName $dataSourceAttribute.TfName}}.{{if eq $dataSourceAttribute.Type "Int64"}}ValueInt64(){{else}}ValueString(){{end}}))
			return
		}
	}
	{{- end}}
	{{- end}}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	{{- if .IsBulk}}
	
	//// Prepare state to track creation process. Create request is split to multiple requests, where just subset of them may be successful
    // Copy fields, as those may contain domain information or other references
    state := plan
    // Create random ID to track bulk resource. This does not relate to FMC in any way
    state.Id = types.StringValue(uuid.New().String())
	// Erase all Items, those will be filled in after creation
    state.Items = make(map[string]{{camelCase .Name}}Items, len(plan.Items))
    // Creation process is put in a separate function, as that same proces will be needed with `Update`
    plan, diags = r.createSubresources(ctx, state, plan, reqMods...)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        // Save state for whatever was already created
        diags = resp.State.Set(ctx, &plan)
		tflog.Debug(ctx, fmt.Sprintf("%s: Create failed, some items might have been created", plan.Id.ValueString()))
        resp.Diagnostics.Append(diags...)
        return
    }
	{{- end}}

	{{- if not .IsBulk}}

	// Create object
	body := plan.toBody(ctx, {{camelCase .Name}}{})
	{{- if .AdjustBody}}
	body = plan.adjustBody(ctx, body)
	{{- end}}

	{{- if .PutCreate}}
	res, err := r.client.Put(plan.getPath()+"/"+url.PathEscape(plan.Id.ValueString()), body, reqMods...)
	{{- else}}
	res, err := r.client.Post(plan.getPath(), body, reqMods...)
	{{- end}}
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("id").String())
	plan.fromBodyUnknowns(ctx, res)

	{{- if hasResourceId .Attributes}}
	res, err = r.client.Get(plan.getPath() + "/" + url.QueryEscape(plan.Id.ValueString()), reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	plan.fromBodyUnknowns(ctx, res)
	{{- end}}
	{{- end}}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *{{camelCase .Name}}Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	{{- if .MinimumVersion}}
	// Get FMC version
	fmcVersion, _ := version.NewVersion(strings.Split(r.client.FMCVersion, " ")[0])

	// Check if FMC client is connected to supports this object
	if fmcVersion.LessThan(minFMCVersion{{camelCase .Name}}) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support {{.Name}}, minimum required version is {{.MinimumVersion}}", r.client.FMCVersion))
		return
	}
	{{- end}}
	var state {{camelCase .Name}}

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	{{- if isDomainDependent .}}
	if !state.Domain.IsNull() && state.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(state.Domain.ValueString()))
	}
	{{- end}}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	{{if .IsBulk}}
	// Get all objects from FMC
	urlPath := state.getPath() + "?expanded=true"
	{{- else}}
	urlPath := state.getPath() + "/" + url.QueryEscape(state.Id.ValueString())
	{{- end}}
	res, err := r.client.Get(urlPath, reqMods...)
	{{if not .IsBulk}}
	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
		resp.State.RemoveResource(ctx)
		return
	} else {{end}} if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	imp, diags := helpers.IsFlagImporting(ctx, req)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// After `terraform import` we switch to a full read.
	if imp {
		state.fromBody(ctx, res)
	} else {
		state.fromBodyPartial(ctx, res)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update

func (r *{{camelCase .Name}}Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state {{camelCase .Name}}

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Read state
	diags = req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	{{- if not .NoUpdate}}
	reqMods := [](func(*fmc.Req)){}
	{{- if isDomainDependent .}}
	if !plan.Domain.IsNull() && plan.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(plan.Domain.ValueString()))
	}
	{{- end}}
	{{- end}}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))
	{{- if not .NoUpdate}}
	{{- if not .IsBulk}}

	body := plan.toBody(ctx, state)
	{{- if .AdjustBody}}
	body = plan.adjustBody(ctx, body)
	{{- end}}
	res, err := r.client.Put(plan.getPath() + "/" + url.QueryEscape(plan.Id.ValueString()), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	{{/* 
	This part of the code was not used to generate any code. Also, this seems to be using harcoded to id, regardless of the actual resource_id is.
	
	{{- if hasResourceId .Attributes}}
	res, err = r.client.Get(plan.getPath() + "/" + url.QueryEscape(plan.Id.ValueString()), reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	plan.fromBodyUnknowns(ctx, res)
	{{- end}}
	*/}}
	{{- if hasComputedRefreshValue .Attributes}}
	plan.fromBodyUnknowns(ctx, res)
	{{- end}}

	{{- end}}

	{{- if .IsBulk}}

	{{- if hasRequiresReplace $.Attributes}}
	// Get objects that need to be replaced due to `requires_replace` flag
	toBeReplaced := plan.findObjectsToBeReplaced(ctx, state)
	{{- end}}

	// DELETE
	// Delete objects (that are present in state, but missing in plan)
	{{- if hasRequiresReplace $.Attributes}}
	toDelete := toBeReplaced.Clone()
	{{- else}}
	var toDelete {{camelCase .Name}}
	toDelete.Items = make(map[string]{{camelCase .Name}}Items, len(state.Items))
	{{- end}}
	planOwnedIDs := make(map[string]string, len(plan.Items))

	// Prepare list of ID that are in plan
	for k, v := range plan.Items {
		if !v.Id.IsUnknown() && v.Id.ValueString() != "" {
			planOwnedIDs[v.Id.ValueString()] = k
		}
	}

	// Check if ID from state list is in plan as well. If not, mark it for delete
	for k, v := range state.Items {
		if _, ok := planOwnedIDs[v.Id.ValueString()]; !ok {
			toDelete.Items[k] = v
		}
	}

	// If there are objects marked to be deleted
	if len(toDelete.Items) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to delete: %d", state.Id.ValueString(), len(toDelete.Items)))
		state, diags = r.deleteSubresources(ctx, state, toDelete, reqMods...)
		if diags != nil {
			resp.Diagnostics.Append(diags...)
			diags = resp.State.Set(ctx, &state)
			resp.Diagnostics.Append(diags...)
			return
		}
	}

	// CREATE
	// Create new objects (objects that have missing IDs in plan)
	{{- if hasRequiresReplace $.Attributes}}
	toCreate := toBeReplaced.Clone()
	toCreate.clearItemsIds(ctx)
	{{- else}}
	var toCreate {{camelCase .Name}}
	toCreate.Items = make(map[string]{{camelCase .Name}}Items, len(plan.Items))
	{{- end}}
	// Scan plan for items with no ID
	for k, v := range plan.Items {
		if v.Id.IsUnknown() || v.Id.IsNull() {
			toCreate.Items[k] = v
		}
	}

	// If there are objects marked for create
	if len(toCreate.Items) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to create: %d", state.Id.ValueString(), len(toCreate.Items)))
		state, diags = r.createSubresources(ctx, state, toCreate, reqMods...)
		if diags != nil {
			resp.Diagnostics.Append(diags...)
			diags = resp.State.Set(ctx, &state)
			resp.Diagnostics.Append(diags...)
			return
		}
	}

	// UPDATE
	// Update objects (objects that have different definition in plan and state)
	var notEqual bool
	var toUpdate {{camelCase .Name}}
	toUpdate.Items = make(map[string]{{camelCase .Name}}Items, len(plan.Items))

	{{- if hasRequiresReplace $.Attributes}}

	for tmp, valueState := range state.Items {
		// Check if the ID from state is on toBeReplaced list
		if _, ok := toBeReplaced.Items[tmp]; ok {
			// If it is, skip it as it was handled by delete/create processes
			continue
		}
	{{- else }}
	
	for _, valueState := range state.Items {
	{{- end}}

		// Check if the ID from plan exists on list of ID owned by state
		if keyState, ok := planOwnedIDs[valueState.Id.ValueString()]; ok {

			// Check if items in state and plan are qual
			notEqual, diags = helpers.IsConfigUpdatingAt(ctx, req.Plan, req.State, path.Root("items").AtMapKey(keyState))
			if diags != nil {
				resp.Diagnostics.Append(diags...)
				diags = resp.State.Set(ctx, &state)
				resp.Diagnostics.Append(diags...)
				return
			}

			// If definitions differ, add object to update list
			if notEqual {
				toUpdate.Items[keyState] = plan.Items[keyState]
			}
		}
	}

	// If there are objects marked for update
	if len(toUpdate.Items) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to update: %d", state.Id.ValueString(), len(toUpdate.Items)))
		state, diags = r.updateSubresources(ctx, state, toUpdate, reqMods...)
		if diags != nil {
			resp.Diagnostics.Append(diags...)
			diags = resp.State.Set(ctx, &state)
			resp.Diagnostics.Append(diags...)
			return
		}
	}
	plan = state
	{{- end}}

	{{- end}}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *{{camelCase .Name}}Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state {{camelCase .Name}}

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	{{- if isDomainDependent .}}
	if !state.Domain.IsNull() && state.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(state.Domain.ValueString()))
	}
	{{- end}}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	{{- if not .NoDelete}}
	{{- if not .IsBulk}}
	{{- if .PutDelete}}
	body := state.toBodyPutDelete(ctx)
	res, err := r.client.Put(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString()), body, reqMods...)
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (PUT), got error: %s, %s", err, res.String()))
		return
	}
	{{- else}}
	res, err := r.client.Delete(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()), reqMods...)
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}
	{{- end}}
	{{- end}}
	{{- if .IsBulk}}

	// Execute delete
	state, diags = r.deleteSubresources(ctx, state, state, reqMods...)
	resp.Diagnostics.Append(diags...)

	// Check if every element was removed
	if len(state.Items) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Not all elements have been removed", state.Id.ValueString()))
		diags = resp.State.Set(ctx, &state)
		resp.Diagnostics.Append(diags...)
		return
	}

	{{- end}}
	{{- end}}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
{{- if not .NoImport}}

{{- if not .IsBulk}}

func (r *{{camelCase .Name}}Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	{{- if hasReference .Attributes}}
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != {{importParts .Attributes}}{{range $index, $attr := .Attributes}}{{if $attr.Reference}} || idParts[{{$index}}] == ""{{end}}{{end}}  || idParts[{{subtract (importParts .Attributes) 1}}] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: {{range $index, $attr := .Attributes}}{{if $attr.Reference}}{{if $index}},{{end}}<{{$attr.TfName}}>{{end}}{{end}},<id>. Got: %q", req.ID),
		)
		return
	}

	{{- range $index, $attr := .Attributes}}
	{{- if or $attr.Reference $attr.Id}}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("{{$attr.TfName}}"), idParts[{{$index}}])...)
	{{- end}}
	{{- end}}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[{{subtract (importParts .Attributes) 1}}])...)
	{{- else}}
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	{{- end}}

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}
{{- end}}

{{- if .IsBulk}}
func (r *{{camelCase .Name}}Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import looks for string in the following format: <domain_name>,<ref_id>,[<object1_name>,<object2_name>,...]
	// <domain_name> is optional
	// <ref_id> for objects that have `reference` attributes
	// <object1_name>,<object2_name>,... is coma-separated list of object names
	var config {{camelCase .Name}}

	// Compile pattern for import command parsing
	var inputPattern = regexp.MustCompile(`^(?:(?P<domain>[^\s,]+),)?
	{{- if hasReference .Attributes -}}{{- range $index, $attr := .Attributes -}}{{- if $attr.Reference -}}
	(?P<{{$attr.TfName}}>[^\s,]+),
	{{- end -}}{{- end -}}{{- end -}}
	\[(?P<names>.*?)\]$`)

	// Parse parameter
	match := inputPattern.FindStringSubmatch(req.ID)

	// Check if regex matched
	if match == nil {
		resp.Diagnostics.AddError("Import error", "Failed to parse import parameters. Please provide import string in the following format: <domain_name>,[<object1_name>,<object2_name>,...]")
		return
	}

	// Extract values
	if tmpDomain := match[inputPattern.SubexpIndex("domain")]; tmpDomain != "" {
		config.Domain = types.StringValue(tmpDomain)
	}
	names := strings.Split(match[inputPattern.SubexpIndex("names")], ",")

	// Fill state with names of objects to import
	config.Items = make(map[string]{{camelCase .Name}}Items, len(names))
	for _, v := range names {
		config.Items[v] = {{camelCase .Name}}Items{}
	}

	{{if hasReference .Attributes -}}
	// Set reference attributes
	{{range $index, $attr := .Attributes -}}
	{{- if $attr.Reference -}}
	config.{{toGoName $attr.TfName}} = types.StringValue(match[inputPattern.SubexpIndex("{{$attr.TfName}}")])
	{{- end -}}
	{{- end -}}
	{{- end}}

	// Generate new ID
	config.Id = types.StringValue(uuid.New().String())

	// Set filled in structure
	diags := resp.State.Set(ctx, config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set import flag
	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}
{{- end}}

{{- end}}
// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin createSubresources

{{- if .IsBulk}}
// createSubresources takes list of objects, splits them into bulks and creates them
// We want to save the state after each create event, to be able track already created resources
func (r *{{camelCase .Name}}Resource) createSubresources(ctx context.Context, state, plan {{camelCase .Name}}, reqMods ...func(*fmc.Req)) ({{camelCase .Name}}, diag.Diagnostics) {	
	{{- if .MinimumVersionBulkCreate}}
	// Get FMC version from the clinet
	fmcVersion, _ := version.NewVersion(strings.Split(r.client.FMCVersion, " ")[0])

	// Check if FMC version supports bulk creates
	if fmcVersion.LessThan(minFMCVersionBulkCreate{{camelCase .Name}}) {
		tflog.Debug(ctx, fmt.Sprintf("%s: One-by-one creation mode ({{.Name}})", state.Id.ValueString()))
		var tmpObject {{camelCase .Name}}
		tmpObject.Items = make(map[string]{{camelCase .Name}}Items, 1)
		for k, v := range plan.Items {
			tmpObject.Items[k] = v

			body := tmpObject.toBodyNonBulk(ctx, state)
			{{- if .AdjustBody}}
			body = tmpObject.adjustBody(ctx, body)
			{{- end}}
			res, err := r.client.Post(state.getPath(), body, reqMods...)
			if err != nil {
				return state, diag.Diagnostics{
					diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("%s: Failed to create object (POST) id %s, got error: %s, %s", state.Id.ValueString(), v.Id.ValueString(), err, res.String())),
				}
			}

			// fromBodyUnknowns expect result to be listed under "items" key
			body, _ = sjson.SetRaw("{}", "items.-1", res.String())
			res = gjson.Parse(body)

			// Read computed values
			tmpObject.fromBodyUnknowns(ctx, res)

			// Save object to plan
			state.Items[k] = tmpObject.Items[k]

			// Clear tmpObject.Items
			delete(tmpObject.Items, k)

		}
	} else { {{- end}}
		var idx = 0
		var bulk {{camelCase .Name}}
		bulk.Items = make(map[string]{{camelCase .Name}}Items, bulkSizeCreate{{if .BulkSizeCreate}}{{camelCase .Name}}{{end}})

		tflog.Debug(ctx, fmt.Sprintf("%s: Bulk creation mode ({{.Name}})", state.Id.ValueString()))

		// iterate over all items
		for k, v := range plan.Items {
			// count loops
			idx++

			// add object to current bulk
			bulk.Items[k] = v

			// If bulk size was reached or all entries have been processed
			if idx%bulkSizeCreate{{if .BulkSizeCreate}}{{camelCase .Name}}{{end}} == 0 || idx == len(plan.Items) {

				// Parse body of the request to string
				body := bulk.toBody(ctx, {{camelCase .Name}}{})

				{{- if .AdjustBody}}
				body = bulk.adjustBodyBulk(ctx, body)
				{{- end}}

				// Execute request
				urlPath := state.getPath() + "?bulk=true"
				res, err := r.client.Post(urlPath, body, reqMods...)
				if err != nil {
					return state, diag.Diagnostics{
						diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Failed to create a bulk (POST) id: %s, got error: %s, %s", state.Id.ValueString(), err, res.String())),
					}
				}

				// Read result and save it to the state
				bulk.fromBodyUnknowns(ctx, res)
				maps.Copy(state.Items, bulk.Items)

				// Clear bulk item for next run
				bulk.Items = make(map[string]{{camelCase .Name}}Items, bulkSizeCreate{{if .BulkSizeCreate}}{{camelCase .Name}}{{end}})
			}
		}
	{{- if .MinimumVersionBulkCreate}}
	}
	{{- end}}

	return state, nil
}
{{- end}}

// End of section. //template:end createSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin deleteSubresources

{{- if .IsBulk}}
// deleteSubresources takes list of objects and deletes them either in bulk, or one-by-one, depending on FMC version
func (r *{{camelCase .Name}}Resource) deleteSubresources(ctx context.Context, state, plan {{camelCase .Name}}, reqMods ...func(*fmc.Req)) ({{camelCase .Name}}, diag.Diagnostics) {
	objectsToRemove := plan.Clone()
	
	{{- if .MinimumVersionBulkDelete}}
	
	// Get FMC version from the clinet
	fmcVersion, _ := version.NewVersion(strings.Split(r.client.FMCVersion, " ")[0])

	// Check if FMC version supports bulk deletes
	if fmcVersion.LessThan(minFMCVersionBulkDelete{{camelCase .Name}}) {
		tflog.Debug(ctx, fmt.Sprintf("%s: One-by-one deletion mode ({{.Name}})", state.Id.ValueString()))
		for k, v := range objectsToRemove.Items {
			// Check if the object was not already deleted
			if v.Id.IsNull() {
				delete(state.Items, k)
				continue
			}

			urlPath := state.getPath() + "/" + url.QueryEscape(v.Id.ValueString())
			res, err := r.client.Delete(urlPath, reqMods...)
			if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
				return state, diag.Diagnostics{
					diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("%s: Failed to delete object (DELETE) id %s, got error: %s, %s", state.Id.ValueString(), v.Id.ValueString(), err, res.String())),
				}
			}

			// Remove deleted item from state
			delete(state.Items, k)
		}
	} else { {{- end}}
		tflog.Debug(ctx, fmt.Sprintf("%s: Bulk deletion mode ({{.Name}})", state.Id.ValueString()))

		var idx = 0

		estimatedIDLength := 37 // UUID length + comma
		estimatedCapacity := min(len(objectsToRemove.Items)*estimatedIDLength, maxUrlParamLength)
		var idsToRemove strings.Builder
		idsToRemove.Grow(estimatedCapacity)

		for k, v := range objectsToRemove.Items {
			// Counter
			idx++

			// Check if the object was not already deleted
			if v.Id.IsNull() {
				delete(state.Items, k)
				continue
			}

			// Create list of IDs of items to delete
			idsToRemove.WriteString(v.Id.ValueString())
			idsToRemove.WriteString(",")

			// If bulk size was reached or all entries have been processed
			if idsToRemove.Len() >= maxUrlParamLength || idx == len(objectsToRemove.Items) {
				urlPath := state.getPath() + "?bulk=true&filter=ids:" + url.QueryEscape(idsToRemove.String())
				res, err := r.client.Delete(urlPath, reqMods...)
				if err != nil {
					return state, diag.Diagnostics{
						diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("%s: Failed to delete subobject(s) (DELETE), got error: %s, %s", state.Id.ValueString(), err, res.String())),
					}
				}

				// Read result and remove deleted items from state
				deletedItems := res.Get("items.#.name").Array()
				for _, name := range deletedItems {
					delete(state.Items, name.String())
				}

				// Reset ID string
				idsToRemove.Reset()
			}
		}

	{{- if .MinimumVersionBulkDelete}}
	}
	{{- end}}

	return state, nil
}
{{- end}}

// End of section. //template:end deleteSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin updateSubresources

{{- if .IsBulk}}

// updateSubresources take elements one-by-one and updates them, as bulks are not supported
func (r *{{camelCase .Name}}Resource) updateSubresources(ctx context.Context, state, plan {{camelCase .Name}}, reqMods ...func(*fmc.Req)) ({{camelCase .Name}}, diag.Diagnostics) {
	var tmpObject {{camelCase .Name}}
	tmpObject.Items = make(map[string]{{camelCase .Name}}Items, 1)

	tflog.Debug(ctx, fmt.Sprintf("%s: One-by-one update mode ({{.Name}})", state.Id.ValueString()))

	for k, v := range plan.Items {
		tmpObject.Items[k] = v

		body := tmpObject.toBodyNonBulk(ctx, state)
		{{- if .AdjustBody}}
		body = tmpObject.adjustBody(ctx, body)
		{{- end}}
		urlPath := state.getPath() + "/" + url.QueryEscape(v.Id.ValueString())
		res, err := r.client.Put(urlPath, body, reqMods...)
		if err != nil {
			return state, diag.Diagnostics{
				diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Failed to update object (PUT) id %s, got error: %s, %s", state.Id.ValueString(), err, res.String())),
			}
		}

		// Update state
		state.Items[k] = v

		// Clear tmpObject.Items
		delete(tmpObject.Items, k)
	}

	return state, nil
}
{{- end}}

// End of section. //template:end updateSubresources