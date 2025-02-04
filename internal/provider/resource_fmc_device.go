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
	"time"

	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/netascode/terraform-provider-fmc/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Mutex to protect policy assignment changes
var policyMu sync.Mutex

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &DeviceResource{}
	_ resource.ResourceWithImportState = &DeviceResource{}
)

func NewDeviceResource() resource.Resource {
	return &DeviceResource{}
}

type DeviceResource struct {
	client *fmc.Client
}

func (r *DeviceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device"
}

func (r *DeviceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Device.").String,

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
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("User-specified name, must be unique.").String,
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the device; this value is always 'Device'.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"host_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Hostname or IP address of the device. Either the host_name or nat_id must be present.").String,
				Required:            true,
			},
			"nat_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("(used for device registration behind NAT) If the device to be registered and the Firepower Management Center are separated by network address translation (NAT), set a unique string identifier.").String,
				Optional:            true,
			},
			"license_capabilities": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Array of strings representing the license capabilities on the managed device. ESSENTIALS is mandatory").AddStringEnumDescription("ESSENTIALS", "IPS", "URL", "MALWARE_DEFENSE", "CARRIER", "SECURE_CLIENT_PREMIER", "SECURE_CLIENT_PREMIER_ADVANTAGE", "SECURE_CLIENT_VPNOnly", "BASE", "THREAT", "PROTECT", "CONTROL", "URLFilter", "MALWARE", "VPN", "SSL").String,
				ElementType:         types.StringType,
				Required:            true,
				Validators: []validator.Set{
					setvalidator.ValueStringsAre(
						stringvalidator.OneOf("ESSENTIALS", "IPS", "URL", "MALWARE_DEFENSE", "CARRIER", "SECURE_CLIENT_PREMIER", "SECURE_CLIENT_PREMIER_ADVANTAGE", "SECURE_CLIENT_VPNOnly", "BASE", "THREAT", "PROTECT", "CONTROL", "URLFilter", "MALWARE", "VPN", "SSL"),
					),
				},
			},
			"registration_key": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Registration Key identical to the one previously configured on the device (`configure manager`).").String,
				Required:            true,
			},
			"device_group_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of the device group.").String,
				Optional:            true,
			},
			"prohibit_packet_transfer": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Value true prohibits the device from sending packet data with events to the Firepower Management Center. Value false allows the transfer when a certain event is triggered. Not all traffic data is sent; connection events do not include a payload, only connection metadata.").String,
				Optional:            true,
			},
			"performance_tier": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Performance tier for the managed device, applicable only to vFTD devices >=6.8.0.").AddStringEnumDescription("FTDv5", "FTDv10", "FTDv20", "FTDv30", "FTDv50", "Legacy").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("FTDv5", "FTDv10", "FTDv20", "FTDv30", "FTDv50", "Legacy"),
				},
			},
			"snort_engine": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Performance tier for the managed device, applicable only to vFTD devices >=6.8.0.").AddStringEnumDescription("SNORT2", "SNORT3").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("SNORT2", "SNORT3"),
				},
			},
			"object_group_search": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enables Object Group Search").String,
				Optional:            true,
			},
			"access_policy_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The UUID of the assigned access control policy. For example `fmc_access_control_policy.example.id`.").String,
				Required:            true,
			},
			"nat_policy_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The UUID of the assigned NAT policy.").String,
				Optional:            true,
			},
			"health_policy_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The UUID of the assigned Health policy.").String,
				Optional:            true,
			},
			"info_version": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Version of the registered device - Informational only.").String,
				Computed:            true,
			},
			"info_health_status": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Health Status of the device - Informational only.").String,
				Computed:            true,
			},
			"info_health_message": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Health Message of the device - Informational only.").String,
				Computed:            true,
			},
			"info_is_connected": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Shows if the device is connected - Informational only.").String,
				Computed:            true,
			},
			"info_deployment_status": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Shows deployment status - Informational only.").String,
				Computed:            true,
			},
			"info_ftd_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("FTD Mode - Informational only.").String,
				Computed:            true,
			},
			"info_device_serial_number": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Device Serial Number - Informational only.").String,
				Computed:            true,
			},
			"info_snort_version": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Snort Version - Informational only.").String,
				Computed:            true,
			},
			"info_vdb_version": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("VDB Version - Informational only.").String,
				Computed:            true,
			},
			"info_lsp_version": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("LSP Version - Informational only.").String,
				Computed:            true,
			},
			"info_deployed_access_policy_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Deployed Access Control Policy Name - Informational only.").String,
				Computed:            true,
			},
			"info_deployed_health_policy_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Deployed Health Policy Name - Informational only.").String,
				Computed:            true,
			},
		},
	}
}

func (r *DeviceResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

func (r *DeviceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan Device
	// time to wait for time-based loops
	const atom time.Duration = 5 * time.Second

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !plan.Domain.IsNull() && plan.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(plan.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, Device{})
	body, _ = sjson.Delete(body, "dummy_nat_policy_id")
	body, _ = sjson.Delete(body, "dummy_health_policy_id")
	res, err := r.client.Post(plan.getPath(), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}

	taskID := res.Get("metadata.task.id").String()
	tflog.Debug(ctx, fmt.Sprintf("%s: Async task initiated successfully", taskID))

	// We need device's UUID, but it only shows after the task succeeds. Poll the task.
	for i := time.Duration(0); i < 5*time.Minute; i += atom {
		task, err := r.client.Get("/api/fmc_config/v1/domain/{DOMAIN_UUID}/job/taskstatuses/"+url.QueryEscape(taskID), reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to read object (GET), got error: %s, %s", err, task.String()))
			return
		}
		stat := strings.ToUpper(task.Get("status").String())
		if stat == "FAILED" {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("API task for the new device failed: %s, %s", task.Get("message"), task.Get("description")))
			return
		}
		if stat != "PENDING" && stat != "RUNNING" && stat != "IN_PROGRESS" {
			break
		}
		time.Sleep(atom)
	}

	bulk, err := r.client.Get(plan.getPath() + "?filter=name:" + url.QueryEscape(plan.Name.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to read object (GET), got error: %s, %s", err, bulk))
		return
	}

	plan.Id = types.StringValue(bulk.Get("items.0.id").String())
	if plan.Id.ValueString() == "" {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("No device named %q: %s", plan.Name.ValueString(), bulk))
		return
	}

	if bulk.Get("items.1").Exists() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Multiple devices named %q: %s", plan.Name.ValueString(), bulk))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Configuring the non-access policy assignments", plan.Id.ValueString()))

	if !plan.NatPolicyId.IsNull() {
		diags = r.updatePolicy(ctx, plan.Id, path.Root("nat_policy_id"), req.Plan, resp.State, reqMods...)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	if !plan.HealthPolicyId.IsNull() {
		diags = r.updatePolicy(ctx, plan.Id, path.Root("health_policy_id"), req.Plan, resp.State, reqMods...)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	// Let long-running deployment finish because it enables DELETE verb. Our tests really expect that.
	for i := time.Duration(0); i < 10*time.Minute; i += atom {
		res, err = r.client.Get(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
			return
		}
		if res.Get("accessPolicy.id").Exists() {
			break // access policy fully deployed
		}
		time.Sleep(atom)
	}

	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Created successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *DeviceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state Device

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

	// Get device object
	res, err := r.client.Get(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString()), reqMods...)
	if err != nil && strings.Contains(err.Error(), "StatusCode 400") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	// Get policy assignments for the device
	policies, err := r.client.Get("/api/fmc_config/v1/domain/{DOMAIN_UUID}/assignment/policyassignments?expanded=true", reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, policies.String()))
		return
	}

	imp, diags := helpers.IsFlagImporting(ctx, req)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Update body with policy assignments
	res = state.fromBodyPolicy(ctx, res, policies)

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

func (r *DeviceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state Device
	var diags diag.Diagnostics

	// Read plan
	diags = req.Plan.Get(ctx, &plan)
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
	body, _ = sjson.Delete(body, "accessPolicy") // usable for POST, but not for PUT
	body, _ = sjson.Delete(body, "dummy_nat_policy_id")
	body, _ = sjson.Delete(body, "healthPolicy") // usable for POST, but not for PUT
	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	if plan.AccessPolicyId != state.AccessPolicyId {
		diags = r.updatePolicy(ctx, plan.Id, path.Root("access_policy_id"), req.Plan, req.State, reqMods...)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	if plan.NatPolicyId != state.NatPolicyId {
		diags = r.updatePolicy(ctx, plan.Id, path.Root("nat_policy_id"), req.Plan, req.State, reqMods...)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	if plan.HealthPolicyId != state.HealthPolicyId {
		diags = r.updatePolicy(ctx, plan.Id, path.Root("health_policy_id"), req.Plan, req.State, reqMods...)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	if plan.DeviceGroupId != state.DeviceGroupId {
		diags = r.updateDeviceGroup(ctx, plan.Id, req.Plan, req.State, reqMods...)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// updatePolicy updates policy-to-device assignment of one specific device (UUID) and of one specific policy type
// (policyPath points to a different attribute for Access Policy, NAT Policy, Platform Settings Policy, etc.).
func (r *DeviceResource) updatePolicy(ctx context.Context, device basetypes.StringValue, policyPath path.Path, plan tfsdk.Plan, state tfsdk.State, reqMods ...func(*fmc.Req)) diag.Diagnostics {
	devId := device.ValueString()

	var planPolicy types.String
	if diags := plan.GetAttribute(ctx, policyPath, &planPolicy); diags.HasError() {
		return diags
	}

	var statePolicy types.String
	if diags := state.GetAttribute(ctx, policyPath, &statePolicy); diags.HasError() {
		return diags
	}

	tflog.Debug(ctx, fmt.Sprintf("policy assignment %s id %s policy planned  %s, state %s", policyPath, devId, planPolicy, statePolicy))

	if statePolicy.Equal(planPolicy) {
		return nil
	}

	// Marginal risk of data race follows (GET/PUT).
	// Partial self-protection in case user specifies terraform -parallelism 2 or more:
	policyMu.Lock()
	defer policyMu.Unlock()

	if planPolicy.IsNull() {
		if policyPath.String() == "health_policy_id" {
			return nil
		}
		res, err := r.client.Get("/api/fmc_config/v1/domain/{DOMAIN_UUID}/assignment/policyassignments/"+url.QueryEscape(statePolicy.ValueString()), reqMods...)
		if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
			return nil
		} else if err != nil {
			return diag.Diagnostics{diag.NewErrorDiagnostic(
				"Client Error",
				fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()),
			)}
		}
		targets := res.Get(fmt.Sprintf(`targets.#(id != "%s")#`, devId))
		body, err := sjson.SetRaw(res.String(), "targets", targets.Raw)
		tflog.Debug(ctx, fmt.Sprintf("policy assignment with device %s removed: %s", devId, res))
		if err != nil {
			return diag.Diagnostics{diag.NewAttributeErrorDiagnostic(
				path.Root("id"),
				"Internal Error",
				fmt.Sprintf("Failed to set JSON \"targets\", got error: %s, %s, %s", err, res, targets),
			)}
		}

		res, err = r.client.Put("/api/fmc_config/v1/domain/{DOMAIN_UUID}/assignment/policyassignments"+"/"+url.QueryEscape(statePolicy.ValueString()), body, reqMods...)
		if err != nil {
			return diag.Diagnostics{diag.NewErrorDiagnostic(
				"Client Error",
				fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()),
			)}
		}
		return nil
	}

	tflog.Debug(ctx, fmt.Sprintf("policy path: %s", policyPath.String()))
	if policyPath.String() == "health_policy_id" {
		stub, _ := sjson.Set("{}", "id", planPolicy.ValueString())
		stub, _ = sjson.Set(stub, "policy.id", planPolicy.ValueString())
		stub, _ = sjson.Set(stub, "policy.type", "HealthPolicy")
		stub, _ = sjson.Set(stub, "targets", []any{})
		stubdev, _ := sjson.Set("{}", "id", devId)
		stubdev, _ = sjson.Set(stubdev, "type", "Device")
		stub, _ = sjson.SetRaw(stub, "targets.-1", stubdev)
		res, err := r.client.Post("/api/fmc_config/v1/domain/{DOMAIN_UUID}/assignment/policyassignments", stub, reqMods...)
		if err != nil {
			return diag.Diagnostics{diag.NewErrorDiagnostic(
				"Client Error",
				fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()),
			)}
		}
		return nil
	}

	res, err := r.client.Get("/api/fmc_config/v1/domain/{DOMAIN_UUID}/assignment/policyassignments"+"/"+url.QueryEscape(planPolicy.ValueString()), reqMods...)
	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
		stub, _ := sjson.Set("{}", "id", planPolicy.ValueString())
		stub, _ = sjson.Set(stub, "policy.id", planPolicy.ValueString())
		stub, _ = sjson.Set(stub, "targets", []any{})
		res = gjson.Parse(stub)
	} else if err != nil {
		return diag.Diagnostics{diag.NewErrorDiagnostic(
			"Client Error",
			fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()),
		)}
	}

	if res.Get(fmt.Sprintf(`targets.#(id=="%s")`, devId)).Exists() {
		tflog.Debug(ctx, fmt.Sprintf("target %s already assigned to policy %s", devId, planPolicy))
		return nil
	}

	polBody, err := sjson.Set(res.String(), `targets.-1`, map[string]any{
		"id":   devId,
		"type": "Device",
	})
	if err != nil {
		return diag.Diagnostics{diag.NewAttributeErrorDiagnostic(
			path.Root("id"),
			"Internal Error",
			fmt.Sprintf("Failed to append to JSON list \"targets\", got error: %s, %s", err, res),
		)}
	}

	res, err = r.client.Put("/api/fmc_config/v1/domain/{DOMAIN_UUID}/assignment/policyassignments"+"/"+url.QueryEscape(planPolicy.ValueString()), polBody, reqMods...)
	if err != nil {
		return diag.Diagnostics{diag.NewErrorDiagnostic(
			"Client Error",
			fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()),
		)}
	}

	return nil
}

func (r *DeviceResource) updateDeviceGroup(ctx context.Context, device basetypes.StringValue, plan tfsdk.Plan, state tfsdk.State, reqMods ...func(*fmc.Req)) diag.Diagnostics {
	deviceId := device.ValueString()

	// Extract IDs of current and planned device group
	var planDeviceGroup types.String
	if diags := plan.GetAttribute(ctx, path.Root("device_group_id"), &planDeviceGroup); diags.HasError() {
		return diags
	}

	var stateDeviceGroup types.String
	if diags := state.GetAttribute(ctx, path.Root("device_group_id"), &stateDeviceGroup); diags.HasError() {
		return diags
	}

	if !stateDeviceGroup.IsNull() {
		// Get Device Group to which device currently belongs
		res, err := r.client.Get("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devicegroups/devicegrouprecords"+"/"+url.QueryEscape(stateDeviceGroup.ValueString()), reqMods...)
		if err != nil {
			return diag.Diagnostics{diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Failed to get current device group (GET), got error: %s, %s", err, res.String()))}
		}

		// Remove device from current device group
		// Device Group endpoint returns 'metadata' twice in the response, which breaks sjson.Delete. Hence we copy needed fields to a new JSON.
		var request = ""
		resString := res.String()
		request, _ = sjson.Set(request, "id", gjson.Get(resString, "id").String())
		request, _ = sjson.Set(request, "type", gjson.Get(resString, "type").String())
		request, _ = sjson.Set(request, "name", gjson.Get(resString, "name").String())

		// Get all members
		members := gjson.Get(resString, "members").Array()
		filteredMembers := []string{}

		// Filter out the device to be removed
		for _, member := range members {
			if member.Get("id").String() != deviceId {
				filteredMembers = append(filteredMembers, member.Raw)
			}
		}

		// Update request
		request, _ = sjson.SetRaw(request, "members", fmt.Sprintf("[%s]", strings.Join(filteredMembers, ",")))

		// Make the PUT request
		res, err = r.client.Put("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devicegroups/devicegrouprecords"+"/"+url.QueryEscape(stateDeviceGroup.ValueString()), request, reqMods...)
		if err != nil {
			return diag.Diagnostics{diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Failed to remove (PUT) device from current device group, got error: %s, %s", err, res.String()))}
		}
	}

	if !planDeviceGroup.IsNull() {
		// Get Device Group to which device needs to be assigned
		res, err := r.client.Get("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devicegroups/devicegrouprecords"+"/"+url.QueryEscape(planDeviceGroup.ValueString()), reqMods...)
		if err != nil {
			return diag.Diagnostics{diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Failed to get destination device group (GET), got error: %s, %s", err, res.String()))}
		}

		// Put device to new device group
		// Device Group endpoint returns 'metadata' twice in the response, which breaks sjson.Delete. Hence we copy needed fields to a new JSON.
		var request = ""
		resString := res.String()
		request, _ = sjson.Set(request, "id", gjson.Get(resString, "id").String())
		request, _ = sjson.Set(request, "type", gjson.Get(resString, "type").String())
		request, _ = sjson.Set(request, "name", gjson.Get(resString, "name").String())
		members := gjson.Get(resString, "members")
		if members.Exists() {
			request, _ = sjson.SetRaw(request, "members", members.Raw)
		} else {
			request, _ = sjson.SetRaw(request, "members", "[]")
		}
		request, _ = sjson.SetRaw(request, "members.-1", fmt.Sprintf(`{"id":"%s"}`, deviceId))
		res, err = r.client.Put("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devicegroups/devicegrouprecords"+"/"+url.QueryEscape(planDeviceGroup.ValueString()), request, reqMods...)
		if err != nil {
			return diag.Diagnostics{diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Failed to add (PUT) device to new device group, got error: %s, %s", err, res.String()))}
		}
	}

	return nil
}

func (r *DeviceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state Device

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

	// Deleting a device implicitly mutates also policyassignments.items.*.targets.
	// The updatePolicy is vulnerable during that mutation, protect it:
	policyMu.Lock()
	defer policyMu.Unlock()

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	res, err := r.client.Delete(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString()), reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// Section below is generated&owned by "gen/generator.go". //template:begin import

func (r *DeviceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
