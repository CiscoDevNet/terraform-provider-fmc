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
	"regexp"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &DeviceDDNSResource{}
	_ resource.ResourceWithImportState = &DeviceDDNSResource{}
)

func NewDeviceDDNSResource() resource.Resource {
	return &DeviceDDNSResource{}
}

type DeviceDDNSResource struct {
	client *fmc.Client
}

func (r *DeviceDDNSResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_ddns"
}

func (r *DeviceDDNSResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages a Device DDNS.").AddMinimumVersionHeaderDescription().AddMinimumVersionDescription("7.4").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "Name of the FMC domain",
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"device_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the parent device.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this value is always 'DDNSSetting'.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"dhcp_client_request_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Determine which records you want the DHCP server to update.").AddStringEnumDescription("NOT_SELECTED", "NO_UPDATE", "ONLY_PTR", "BOTH_A_AND_PTR_RECORD").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("NOT_SELECTED", "NO_UPDATE", "ONLY_PTR", "BOTH_A_AND_PTR_RECORD"),
				},
			},
			"dhcp_client_broadcast": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Request that the DHCP server broadcast the DHCP reply (DHCP option 1).").String,
				Optional:            true,
			},
			"dynamic_dns_update": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Which DNS RRs you want the DHCP server to update.").AddStringEnumDescription("NOT_SELECTED", "ONLY_PTR", "BOTH_A_AND_PTR_RECORD").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("NOT_SELECTED", "ONLY_PTR", "BOTH_A_AND_PTR_RECORD"),
				},
			},
			"dhcp_client_request_override": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Override the update actions requested by the DHCP client.").String,
				Optional:            true,
			},
			"dhcp_client_id_interfaces": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of interfaces on which the DHCP client sets the client identifier.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the interface.").String,
							Required:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the interface.").String,
							Required:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of the interface.").String,
							Required:            true,
						},
					},
				},
			},
			"ddns_update_methods": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of DDNS update methods.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the DDNS update method.").String,
							Required:            true,
						},
						"method": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("DDNS update method type.").AddStringEnumDescription("DDNS", "WEB", "FMC_ONLY").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("DDNS", "WEB", "FMC_ONLY"),
							},
						},
						"update_interval_day": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Update interval day portion.").AddIntegerRangeDescription(0, 364).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 364),
							},
						},
						"update_interval_hour": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Update interval hour portion.").AddIntegerRangeDescription(0, 23).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 23),
							},
						},
						"update_interval_minute": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Update interval minute portion.").AddIntegerRangeDescription(0, 59).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 59),
							},
						},
						"update_interval_second": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Update interval second portion.").AddIntegerRangeDescription(0, 59).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 59),
							},
						},
						"web_url": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("URL of the web update server, used when `method` is `WEB`. Syntax https://username:password@provider-domain/xyz?hostname=<h>&myip=<a>.").String,
							Optional:            true,
						},
						"web_update_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Addresses to update when `method` is `WEB`.").AddStringEnumDescription("ALL_ADDRESSES", "IPV4_ADDRESS", "IPV4_AND_ONE_IPV6_ADDRESS", "ONE_IPV6_ADDRESS", "ALL_IPV6_ADDRESSES").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("ALL_ADDRESSES", "IPV4_ADDRESS", "IPV4_AND_ONE_IPV6_ADDRESS", "ONE_IPV6_ADDRESS", "ALL_IPV6_ADDRESSES"),
							},
						},
						"update_records": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("DNS records to update.").AddStringEnumDescription("NOT_DEFINED", "A_RECORDS", "BOTH_A_AND_PTR_RECORDS").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("NOT_DEFINED", "A_RECORDS", "BOTH_A_AND_PTR_RECORDS"),
							},
						},
					},
				},
			},
			"ddns_interface_settings": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of per-interface DDNS settings.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the interface. Id, name, and type must all be specified together.").String,
							Required:            true,
						},
						"interface_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the interface. Id, name, and type must all be specified together.").String,
							Required:            true,
						},
						"interface_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of the interface. Id, name, and type must all be specified together.").String,
							Required:            true,
						},
						"method_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the DDNS update method (from `update_methods`) applied to the interface.").String,
							Optional:            true,
						},
						"hostname": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set hostname for the interface.").String,
							Optional:            true,
						},
						"dhcp_client_request_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("How the DHCP client requests the DHCP server to perform DNS updates.").AddStringEnumDescription("NOT_SELECTED", "NO_UPDATE", "ONLY_PTR", "BOTH_A_AND_PTR_RECORD").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("NOT_SELECTED", "NO_UPDATE", "ONLY_PTR", "BOTH_A_AND_PTR_RECORD"),
							},
						},
						"dynamic_dns_update": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("DNS records that the DDNS method updates.").AddStringEnumDescription("NOT_SELECTED", "ONLY_PTR", "BOTH_A_AND_PTR_RECORD").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("NOT_SELECTED", "ONLY_PTR", "BOTH_A_AND_PTR_RECORD"),
							},
						},
						"dhcp_client_request_override": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Override the DDNS updates that the DHCP client requests the DHCP server to perform.").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *DeviceDDNSResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *DeviceDDNSResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	// Check if FMC client is connected to supports this object
	if r.client.FMCVersionParsed.LessThan(minFMCVersionDeviceDDNS) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support Device DDNS creation, minumum required version is 7.4", r.client.FMCVersion))
		return
	}
	var plan DeviceDDNS

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !plan.Domain.IsNull() && plan.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(plan.Domain.ValueString()))
	}
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

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, DeviceDDNS{})
	res, err := r.client.Put(plan.getPath()+"/"+url.PathEscape(plan.Id.ValueString()), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("id").String())
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *DeviceDDNSResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Check if FMC client is connected to supports this object
	if r.client.FMCVersionParsed.LessThan(minFMCVersionDeviceDDNS) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support Device DDNS, minimum required version is 7.4", r.client.FMCVersion))
		return
	}
	var state DeviceDDNS

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !state.Domain.IsNull() && state.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(state.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	urlPath := state.getPath() + "/" + url.QueryEscape(state.Id.ValueString())
	res, err := r.client.Get(urlPath, reqMods...)

	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
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

func (r *DeviceDDNSResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state DeviceDDNS

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
	reqMods := [](func(*fmc.Req)){}
	if !plan.Domain.IsNull() && plan.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(plan.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	body := plan.toBody(ctx, state)
	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *DeviceDDNSResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DeviceDDNS

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !state.Domain.IsNull() && state.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(state.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	body := state.toBodyPutDelete(ctx)
	res, err := r.client.Put(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString()), body, reqMods...)
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *DeviceDDNSResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Parse import ID
	var inputPattern = regexp.MustCompile(`^(?:(?P<domain>[^\s,]+),)?(?P<device_id>[^\s,]+),(?P<id>[^\s,]+?)$`)
	match := inputPattern.FindStringSubmatch(req.ID)
	if match == nil {
		errMsg := "Failed to parse import parameters.\nPlease provide import string in the following format: <domain>,<device_id>,<id>\n<domain> is optional. If not provided, `Global` is used implicitly and resource's `domain` attribute is not set.\n" + fmt.Sprintf("Got: %q", req.ID)
		resp.Diagnostics.AddError("Import error", errMsg)
		return
	}

	// Set domain, if provided
	if tmpDomain := match[inputPattern.SubexpIndex("domain")]; tmpDomain != "" {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("domain"), tmpDomain)...)
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), match[inputPattern.SubexpIndex("id")])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("device_id"), match[inputPattern.SubexpIndex("device_id")])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin createSubresources

// End of section. //template:end createSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin deleteSubresources

// End of section. //template:end deleteSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin updateSubresources

// End of section. //template:end updateSubresources
