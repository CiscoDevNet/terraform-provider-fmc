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

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/netascode/terraform-provider-fmc/internal/provider/helpers"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &DeviceVTEPPolicyResource{}
	_ resource.ResourceWithImportState = &DeviceVTEPPolicyResource{}
)

func NewDeviceVTEPPolicyResource() resource.Resource {
	return &DeviceVTEPPolicyResource{}
}

type DeviceVTEPPolicyResource struct {
	client *fmc.Client
}

func (r *DeviceVTEPPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_vtep_policy"
}

func (r *DeviceVTEPPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the Device VTEP Policy. Practicioners should ensure only one resource `fmc_device_vtep_policy` exists for a single `fmc_device`, because the FMC API responds with the same single UUID for every request to create a new VTEP Policy on the same Device. Thus multiple resources per Device would unexpectedly overwrite the same set of settings.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "The name of the FMC domain",
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"device_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("UUID of the parent device (fmc_device.example.id).").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"nve_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether to enable NVE on the `device_id`. Can only be false if `vteps` are empty.").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"vteps": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List that can either be empty or contain one VTEP object.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"source_interface_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("UUID of the source interface (e.g. fmc_physical_interface.example.id). It cannot refer to a subinterface.").String,
							Optional:            true,
						},
						"nve_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("VTEP NVE number, currently must always be 1.").AddIntegerRangeDescription(1, 1).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 1),
							},
						},
						"encapsulation_port": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Encapsulation port number. For VXLAN suggested 4789 (default), for GENEVE suggested 6081.").AddIntegerRangeDescription(1024, 65535).AddDefaultValueDescription("4789").String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.Int64{
								int64validator.Between(1024, 65535),
							},
							Default: int64default.StaticInt64(4789),
						},
						"encapsulation_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Encapsulation type.").AddStringEnumDescription("VXLAN", "GENEVE").AddDefaultValueDescription("VXLAN").String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("VXLAN", "GENEVE"),
							},
							Default: stringdefault.StaticString("VXLAN"),
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplace(),
							},
						},
						"neighbor_discovery": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("How to discover addresses of the neighbor VTEPs for the VTEP-to-VTEP communication. For STATIC_PEER_IP and DEFAULT_MULTICAST_GROUP you must set `neighbor_address_literal` to a single IP address. For STATIC_PEER_GROUP you must however set `neighbor_address_id` to a UUID of a network group and such network group can contain only IPv4 Hosts and IPv4 Ranges (but not Networks, etc.).").AddStringEnumDescription("NONE", "STATIC_PEER_IP", "STATIC_PEER_GROUP", "DEFAULT_MULTICAST_GROUP").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("NONE", "STATIC_PEER_IP", "STATIC_PEER_GROUP", "DEFAULT_MULTICAST_GROUP"),
							},
						},
						"neighbor_address_literal": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Used for neighbor_discovery STATIC_PEER_IP, where it holds any unicast IP address. Used for neighbor_discovery DEFAULT_MULTICAST_GROUP, where it holds IP address in range 224.0.0.0 to 239.255.255.255.").String,
							Optional:            true,
						},
						"neighbor_address_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Used for neighbor_discovery STATIC_PEER_GROUP, where it holds UUID of the network group and such network group can contain only IPv4 Hosts and IPv4 Ranges (but not Networks, etc.).").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *DeviceVTEPPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

func (r *DeviceVTEPPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DeviceVTEPPolicy

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

	// Problem workaround:
	// Check if API resource is absent. If present, the POST will return an existing UUID instead of failing.
	// This non-RESTful behavior is a risk: the UUID can be owned/modified by two parties causing clashes.
	existing, err := r.client.Get(plan.getPath(), reqMods...)
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve VTEP policy (GET), got error: %s, %s", err, existing.String()))
		return
	}
	if items := existing.Get("items"); items.IsArray() && len(items.Array()) > 0 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf(
			"VTEP policy already exists for device %s, please consider:\n\n    terraform import fmc_device_vtep_policy.name %s,%s\n\nin order to get it managed within this Terraform config.",
			plan.DeviceId,
			plan.DeviceId.ValueString(),
			items.Array()[0].Get("id"),
		))
		return
	}

	// Create object
	body := plan.toBody(ctx, DeviceVTEPPolicy{})
	res, err := r.client.Post(plan.getPath(), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("id").String())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *DeviceVTEPPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DeviceVTEPPolicy

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

func (r *DeviceVTEPPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state DeviceVTEPPolicy

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

func (r *DeviceVTEPPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DeviceVTEPPolicy

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
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import

func (r *DeviceVTEPPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <device_id>,<id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("device_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
