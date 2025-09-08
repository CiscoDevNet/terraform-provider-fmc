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

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
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
	_ resource.Resource                = &FTDPlatformSettingsICMPAccessResource{}
	_ resource.ResourceWithImportState = &FTDPlatformSettingsICMPAccessResource{}
)

func NewFTDPlatformSettingsICMPAccessResource() resource.Resource {
	return &FTDPlatformSettingsICMPAccessResource{}
}

type FTDPlatformSettingsICMPAccessResource struct {
	client *fmc.Client
}

func (r *FTDPlatformSettingsICMPAccessResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ftd_platform_settings_icmp_access"
}

func (r *FTDPlatformSettingsICMPAccessResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages FTD Platform Settings - ICMP Access.").AddMinimumVersionHeaderDescription().AddMinimumVersionAnyDescription().AddMinimumVersionCreateDescription("7.7").String,

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
			"ftd_platform_settings_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the parent FTD Platform Settings.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this value is always 'ICMPSetting'.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"rate_limit": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rate limit on ICMPv4 Unreachable messages.").AddIntegerRangeDescription(1, 100).AddDefaultValueDescription("1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 100),
				},
				Default: int64default.StaticInt64(1),
			},
			"burst_size": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Burst size on ICMPv4 Unreachable messages.").AddIntegerRangeDescription(1, 10).AddDefaultValueDescription("1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 10),
				},
				Default: int64default.StaticInt64(1),
			},
			"icmp_configs": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ICMP access rules.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"action": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Action to take on matching ICMP packets.").AddStringEnumDescription("Permit", "Deny").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("Permit", "Deny"),
							},
						},
						"icmp_service_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of the ICMP Service.").String,
							Required:            true,
						},
						"source_network_object_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of network object (host, network, network group) defining the source IP addresses for ICMP access.").String,
							Required:            true,
						},
						"interface_literals": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of interfaces to reach SNMP management host.").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"interface_objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of interface objects (Security Zones or Interface Groups) to reach SNMP management host.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the interface object.").String,
										Required:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of the interface object; either 'SecurityZone' or 'InterfaceGroup'.").AddStringEnumDescription("SecurityZone", "InterfaceGroup").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("SecurityZone", "InterfaceGroup"),
										},
									},
									"name": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Name of the interface object.").String,
										Required:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *FTDPlatformSettingsICMPAccessResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *FTDPlatformSettingsICMPAccessResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Get FMC version
	fmcVersion, _ := version.NewVersion(strings.Split(r.client.FMCVersion, " ")[0])

	// Check if FMC client is connected to supports this object
	if fmcVersion.LessThan(minFMCVersionCreateFTDPlatformSettingsICMPAccess) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support FTD Platform Settings ICMP Access creation, minumum required version is 7.7", r.client.FMCVersion))
		return
	}
	var plan FTDPlatformSettingsICMPAccess

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
	body := plan.toBody(ctx, FTDPlatformSettingsICMPAccess{})
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

func (r *FTDPlatformSettingsICMPAccessResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state FTDPlatformSettingsICMPAccess

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

func (r *FTDPlatformSettingsICMPAccessResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state FTDPlatformSettingsICMPAccess

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

func (r *FTDPlatformSettingsICMPAccessResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state FTDPlatformSettingsICMPAccess

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

func (r *FTDPlatformSettingsICMPAccessResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <ftd_platform_settings_id>,<id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("ftd_platform_settings_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
