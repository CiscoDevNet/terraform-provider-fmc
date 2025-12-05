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
	"github.com/google/uuid"
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
	_ resource.Resource                = &DeviceOSPFInterfaceResource{}
	_ resource.ResourceWithImportState = &DeviceOSPFInterfaceResource{}
)

func NewDeviceOSPFInterfaceResource() resource.Resource {
	return &DeviceOSPFInterfaceResource{}
}

type DeviceOSPFInterfaceResource struct {
	client *fmc.Client
}

func (r *DeviceOSPFInterfaceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_ospf_interface"
}

func (r *DeviceOSPFInterfaceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages a Device OSPF Interface.").AddMinimumVersionHeaderDescription().AddMinimumVersionDescription("7.6").String,

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
			"vrf_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the parent VRF.").String,
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
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this is always 'OspfInterface'").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"interface_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of the device interface.").String,
				Required:            true,
			},
			"default_cost": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Cost of sending a packet through the interface.").AddIntegerRangeDescription(1, 65535).AddDefaultValueDescription("10").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(10),
			},
			"priority": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Designated router for a network.").AddIntegerRangeDescription(0, 255).AddDefaultValueDescription("1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 255),
				},
				Default: int64default.StaticInt64(1),
			},
			"mtu_missmatch_ignore": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Ignore MTU mismatch on the interface.").String,
				Optional:            true,
			},
			"hello_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interval, in seconds, between hello packets sent on an interface.").AddIntegerRangeDescription(1, 8192).AddDefaultValueDescription("10").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 8192),
				},
				Default: int64default.StaticInt64(10),
			},
			"transmit_delay": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Estimated time in seconds to send an LSA packet on the interface.").AddIntegerRangeDescription(1, 65535).AddDefaultValueDescription("1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(1),
			},
			"retransmit_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Time in seconds between LSA retransmissions for adjacencies that belong to the interface.").AddIntegerRangeDescription(1, 65535).AddDefaultValueDescription("5").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(5),
			},
			"dead_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Time period in seconds for which hello packets must not be seen before neighbors indicate that the router is down.").AddIntegerRangeDescription(1, 65535).AddDefaultValueDescription("40").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(40),
			},
			"hello_multiplier": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Number of Hello packets to be sent per second.").AddIntegerRangeDescription(3, 20).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(3, 20),
				},
			},
			"point_to_point": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure the interface as point-to-point non-broadcast.").String,
				Optional:            true,
			},
			"bfd": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Bidirectional Forwarding Detection (BFD) on the interface.").String,
				Optional:            true,
			},
			"authentication_password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Password for authentication.").String,
				Optional:            true,
				Sensitive:           true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 8),
				},
			},
			"authentication_area_password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Password for area authentication.").String,
				Optional:            true,
				Sensitive:           true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 8),
				},
			},
			"authentication_area_md5s": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Area MD5 authentication keys.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Key ID.").String,
							Required:            true,
						},
						"key": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("MD5 authentication key.").String,
							Required:            true,
							Sensitive:           true,
						},
					},
				},
			},
			"authentication_md5s": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("MD5 authentication keys.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Key ID.").String,
							Required:            true,
						},
						"key": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("MD5 authentication key.").String,
							Required:            true,
							Sensitive:           true,
						},
					},
				},
			},
			"authentication_key_chain_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Key chain object ID for authentication.").String,
				Optional:            true,
			},
		},
	}
}

func (r *DeviceOSPFInterfaceResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *DeviceOSPFInterfaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	// Check if FMC client is connected to supports this object
	if r.client.FMCVersionParsed.LessThan(minFMCVersionDeviceOSPFInterface) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support Device OSPF Interface creation, minumum required version is 7.6", r.client.FMCVersion))
		return
	}
	var plan DeviceOSPFInterface

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
	body := plan.toBody(ctx, DeviceOSPFInterface{})
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

func (r *DeviceOSPFInterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Check if FMC client is connected to supports this object
	if r.client.FMCVersionParsed.LessThan(minFMCVersionDeviceOSPFInterface) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support Device OSPF Interface, minimum required version is 7.6", r.client.FMCVersion))
		return
	}
	var state DeviceOSPFInterface

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

func (r *DeviceOSPFInterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state DeviceOSPFInterface

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

func (r *DeviceOSPFInterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DeviceOSPFInterface

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
func (r *DeviceOSPFInterfaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	errMsg := "Failed to parse import parameters.\nPlease provide import string in the following format: <domain>,<device_id>,<vrf_id>,<id>\n<domain> is optional. If not provided, `Global` is used implicitly and resource's `domain` attribute is not set.\n<vrf_id> is optional.\n" + fmt.Sprintf("Got: %q", req.ID)
	parts := strings.Split(req.ID, ",")
	if len(parts) < 2 || len(parts) > 4 {
		resp.Diagnostics.AddError("Import error", errMsg)
		return
	}

	for i := range parts {
		if parts[i] == "" {
			resp.Diagnostics.AddError("Import error", errMsg)
			return
		}
	}

	if len(parts) == 2 {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("device_id"), parts[0])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), parts[1])...)
	} else if len(parts) == 3 {
		if err := uuid.Validate(parts[0]); err == nil {
			// First part is UUID, so it's device_id
			resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("device_id"), parts[0])...)
			resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("vrf_id"), parts[1])...)
		} else {
			// First part is domain
			resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("domain"), parts[0])...)
			resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("device_id"), parts[1])...)
		}
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), parts[2])...)

	} else if len(parts) == 4 {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("domain"), parts[0])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("device_id"), parts[1])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("vrf_id"), parts[2])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), parts[3])...)
	}

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin createSubresources

// End of section. //template:end createSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin deleteSubresources

// End of section. //template:end deleteSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin updateSubresources

// End of section. //template:end updateSubresources
