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
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
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
	_ resource.Resource                = &ChassisLogicalDeviceResource{}
	_ resource.ResourceWithImportState = &ChassisLogicalDeviceResource{}
)

func NewChassisLogicalDeviceResource() resource.Resource {
	return &ChassisLogicalDeviceResource{}
}

type ChassisLogicalDeviceResource struct {
	client *fmc.Client
}

func (r *ChassisLogicalDeviceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_chassis_logical_device"
}

func (r *ChassisLogicalDeviceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages a Chassis Logical Device.").String,

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
			"chassis_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the parent chassis.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the device; this value is always 'LogicalDevice'.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the logical device.").String,
				Required:            true,
			},
			"ftd_version": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Version of the logical device, that should be deployed. Image should be pre-deployed to the chassis.").String,
				Required:            true,
			},
			"ipv4_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Management IPv4 address of the logical device.").String,
				Optional:            true,
			},
			"ipv4_netmask": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Netmask of Management IPv4 address.").String,
				Optional:            true,
			},
			"ipv4_gateway": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Gateway for Management IPv4 address.").String,
				Optional:            true,
			},
			"ipv6_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Management IPv6 address of the logical device.").String,
				Optional:            true,
			},
			"ipv6_prefix_length": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Prefix length of Management IPv6 address.").String,
				Optional:            true,
			},
			"ipv6_gateway": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Gateway for Management IPv6 address.").String,
				Optional:            true,
			},
			"search_domain": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Search domain for the logical device.").String,
				Optional:            true,
			},
			"fqdn": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Fully qualified domain name (FQDN) of the logical device.").String,
				Optional:            true,
			},
			"firewall_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Firewall mode of the logical device.").AddStringEnumDescription("ROUTED", "TRANSPARENT").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("ROUTED", "TRANSPARENT"),
				},
			},
			"dns_servers": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("DNS servers for the logical device. Up to three, comma-separated DNS servers can be specified.").String,
				Optional:            true,
			},
			"device_password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Admin password for the logical device.").String,
				Required:            true,
			},
			"permit_expert_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Permit expert mode for the logical device.").AddStringEnumDescription("yes", "no").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("yes", "no"),
				},
			},
			"resource_profile_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the resource profile.").String,
				Required:            true,
			},
			"resource_profile_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the resource profile.").String,
				Required:            true,
			},
			"assigned_interfaces": schema.SetNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface assignment for the logical device.").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the interface.").String,
							Optional:            true,
						},
					},
				},
			},
			"device_group_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the device group.").String,
				Optional:            true,
			},
			"access_policy_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the Access Control Policy to be assigned to the logical device.").String,
				Required:            true,
			},
			"platform_settings_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the platform settings.").String,
				Optional:            true,
			},
			"license_capabilities": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Array of strings representing the license capabilities on the managed device.").AddStringEnumDescription("MALWARE", "URLFilter", "CARRIER", "PROTECT").String,
				ElementType:         types.StringType,
				Optional:            true,
				Validators: []validator.Set{
					setvalidator.ValueStringsAre(
						stringvalidator.OneOf("MALWARE", "URLFilter", "CARRIER", "PROTECT"),
					),
				},
			},
		},
	}
}

func (r *ChassisLogicalDeviceResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *ChassisLogicalDeviceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ChassisLogicalDevice

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, ChassisLogicalDevice{})
	res, err := r.client.Post(plan.getPath(), body, reqMods...)
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

func (r *ChassisLogicalDeviceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ChassisLogicalDevice

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

func (r *ChassisLogicalDeviceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ChassisLogicalDevice

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

func (r *ChassisLogicalDeviceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ChassisLogicalDevice

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
	res, err := r.client.Delete(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString()), reqMods...)
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import

func (r *ChassisLogicalDeviceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <chassis_id>,<id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("chassis_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin createSubresources

// End of section. //template:end createSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin deleteSubresources

// End of section. //template:end deleteSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin updateSubresources

// End of section. //template:end updateSubresources
