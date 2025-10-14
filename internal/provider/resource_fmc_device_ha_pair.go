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
	"github.com/hashicorp/terraform-plugin-framework-validators/resourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &DeviceHAPairResource{}
	_ resource.ResourceWithImportState = &DeviceHAPairResource{}
)

func NewDeviceHAPairResource() resource.Resource {
	return &DeviceHAPairResource{}
}

type DeviceHAPairResource struct {
	client *fmc.Client
}

func (r *DeviceHAPairResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_ha_pair"
}

func (r *DeviceHAPairResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This device manages FTD HA Pair configuration.\n Configuration (like interfaces) of the HA Pair is replicated from the primary device. Nevertheless, please make sure that the configuration of both nodes is consistent.\n On desroy, the HA Pair will be broken and both devices will remain managed by FMC as standalone devices.").String,

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
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the High Availability (HA) Pair.").String,
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object; This is always `DeviceHAPair`.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"primary_device_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of primary FTD in the HA Pair.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"secondary_device_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of secondary FTD in the HA Pair.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"ha_link_interface_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of High Availability Link interface taken from the primary FTD device.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"ha_link_interface_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of High Availability Link interface.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"ha_link_interface_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of High Availability Link interface.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"ha_link_logical_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Logical name of High Availability Link interface.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"ha_link_use_ipv6": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use IPv6 addressing for High Availability communication.").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
			},
			"ha_link_primary_ip": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP of primary node on High Availability interface.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"ha_link_secondary_ip": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP of secondary node on High Availability interface.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"ha_link_netmask": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Subnet mask for High Availability link.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"state_link_use_same_as_ha": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use the same link for state and High Availability.").String,
				Required:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
			},
			"state_link_interface_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of state link physical interface taken from the primary FTD device.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"state_link_interface_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of state link interface.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"state_link_interface_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of state link interface.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"state_link_logical_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Logical name of state link interface.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"state_link_use_ipv6": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use IPv6 addressing for state link communication.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
			},
			"state_link_primary_ip": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP of primary node on state link interface.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"state_link_secondary_ip": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP of secondary node on state link interface.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"state_link_netmask": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Subnet mask for state link.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"encryption_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use encryption for communication.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
			},
			"encryption_key_generation_scheme": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Select the encyption key generation scheme.").AddStringEnumDescription("AUTO", "CUSTOM").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("AUTO", "CUSTOM"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"encryption_key": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Preshared key for encryption if CUSTOM key generation scheme is selected.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"failed_interfaces_percent": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Percentage of Failed Interfaces that triggers failover.").AddIntegerRangeDescription(1, 100).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 100),
				},
			},
			"failed_interfaces_limit": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Number of Failed Interfaces that triggers failover.").AddIntegerRangeDescription(1, 211).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 211),
				},
			},
			"peer_poll_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Peer Pool Time (1-15 SEC or 200-999 MSEC)").AddIntegerRangeDescription(1, 999).AddDefaultValueDescription("1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 999),
				},
				Default: int64default.StaticInt64(1),
			},
			"peer_poll_time_unit": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Peer Pool Time Unit").AddStringEnumDescription("SEC", "MSEC").AddDefaultValueDescription("SEC").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("SEC", "MSEC"),
				},
				Default: stringdefault.StaticString("SEC"),
			},
			"peer_hold_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Peer Hold Time (3-45 SEC or 800-999 MSEC)").AddIntegerRangeDescription(3, 999).AddDefaultValueDescription("15").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(3, 999),
				},
				Default: int64default.StaticInt64(15),
			},
			"peer_hold_time_unit": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Peer Hold Time Unit").AddStringEnumDescription("SEC", "MSEC").AddDefaultValueDescription("SEC").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("SEC", "MSEC"),
				},
				Default: stringdefault.StaticString("SEC"),
			},
			"interface_poll_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Peer Pool Time (1-15 SEC or 500-999 MSEC)").AddIntegerRangeDescription(1, 999).AddDefaultValueDescription("5").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 999),
				},
				Default: int64default.StaticInt64(5),
			},
			"interface_poll_time_unit": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Peer Pool Time Unit").AddStringEnumDescription("SEC", "MSEC").AddDefaultValueDescription("SEC").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("SEC", "MSEC"),
				},
				Default: stringdefault.StaticString("SEC"),
			},
			"interface_hold_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface Hold Time in seconds").AddIntegerRangeDescription(25, 75).AddDefaultValueDescription("25").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(25, 75),
				},
				Default: int64default.StaticInt64(25),
			},
			"action": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("FTD HA PUT operation action. Specifically used for manual switch.").AddStringEnumDescription("SWITCH", "HABREAK").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("SWITCH", "HABREAK"),
				},
			},
		},
	}
}

func (r *DeviceHAPairResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

func (r DeviceHAPairResource) ConfigValidators(ctx context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{
		resourcevalidator.ExactlyOneOf(
			path.MatchRoot("failed_interfaces_percent"),
			path.MatchRoot("failed_interfaces_limit"),
		),
	}
}

func (r *DeviceHAPairResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DeviceHAPair

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

	// Deploy devices in HA Pair (pre-requisite for cluster creation)
	deploy := DeviceDeploy{
		Id:            types.StringValue(plan.Id.ValueString()),
		IgnoreWarning: types.BoolValue(true),
		DeviceIdList:  helpers.GetStringListFromStringSlice([]string{plan.PrimaryDeviceId.ValueString(), plan.SecondaryDeviceId.ValueString()}),
	}
	diags = FMCDeviceDeploy(ctx, r.client, deploy, reqMods)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Create object
	body := plan.toBody(ctx, DeviceHAPair{})
	res, err := r.client.Post(plan.getPath(), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}

	// Wait for the ha pair to be created
	taskID := res.Get("metadata.task.id").String()
	tflog.Debug(ctx, fmt.Sprintf("%s: Async task initiated successfully (id: %s)", plan.Id.ValueString(), taskID))

	diags = FMCWaitForJobToFinish(ctx, r.client, taskID, reqMods)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	check, err := r.client.Get(plan.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to read object (GET), got error: %s, %s", err, check))
		return
	}
	name := "items.#(name==" + url.QueryEscape(plan.Name.ValueString()) + ").id"
	id := check.Get(name).String()
	plan.Id = types.StringValue(id)
	if plan.Id.ValueString() == "" {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("No HA Pair named %q: %s", plan.Name.ValueString(), check))
		return
	}

	// Send second request to configure missing pieces
	body = plan.toBodyUpdateTimers(ctx, DeviceHAPair{})
	if body != "" {
		res, err = r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body, reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT) after create, got error: %s, %s", err, res.String()))
			// Save state, as at this point HA Pair is created, though not fully configured
			diags = resp.State.Set(ctx, &plan)
			resp.Diagnostics.Append(diags...)
			return
		}
	}
	// Ending code to poll object
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *DeviceHAPairResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DeviceHAPair

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

func (r *DeviceHAPairResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state DeviceHAPair

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

	// Update object 'name'. Other attributer are either immutable or updated later
	if state.Name != plan.Name {
		body := ""
		body, _ = sjson.Set(body, "id", plan.Id.ValueString())
		body, _ = sjson.Set(body, "name", plan.Name.ValueString())

		res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body, reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
			return
		}
	}

	// Update object 'timers'
	body := plan.toBodyUpdateTimers(ctx, state)
	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *DeviceHAPairResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DeviceHAPair

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

	// Start of HA Break code
	body := state.toBodyPutDelete(ctx, DeviceHAPair{})
	res, err := r.client.Put(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString()), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to remove object configuration (PUT), got error: %s, %s", err, res.String()))
		return
	}
	// End of HA Break code

	// Adding code to poll object
	taskID := res.Get("metadata.task.id").String()
	tflog.Debug(ctx, fmt.Sprintf("%s: Async task initiated successfully (id: %s)", state.Id.ValueString(), taskID))

	diags = FMCWaitForJobToFinish(ctx, r.client, taskID, reqMods)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *DeviceHAPairResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Parse import ID
	var inputPattern = regexp.MustCompile(`^(?:(?P<domain>[^\s,]+),)?(?P<id>[^\s,]+?)$`)
	match := inputPattern.FindStringSubmatch(req.ID)
	if match == nil {
		errMsg := "Failed to parse import parameters.\nPlease provide import string in the following format: <domain>,<id>\n<domain> is optional. If not provided, `Global` is used implicitly and resource's `domain` attribute is not set.\n" + fmt.Sprintf("Got: %q", req.ID)
		resp.Diagnostics.AddError("Import error", errMsg)
		return
	}

	// Set domain, if provided
	if tmpDomain := match[inputPattern.SubexpIndex("domain")]; tmpDomain != "" {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("domain"), tmpDomain)...)
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), match[inputPattern.SubexpIndex("id")])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
