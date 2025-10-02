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
	"sync"
	"time"

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
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
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages a Device. This resource is not supported in cdFMC - to register the device in cdFMC, please use Security Cloud Control API instead.").String,

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
				MarkdownDescription: helpers.NewAttributeDescription("Id of the device group.").String,
				Optional:            true,
			},
			"prohibit_packet_transfer": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Value true prohibits the device from sending packet data with events to the Firepower Management Center. Value false allows the transfer when a certain event is triggered. Not all traffic data is sent; connection events do not include a payload, only connection metadata.").String,
				Optional:            true,
			},
			"performance_tier": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Performance tier for the managed device.").AddStringEnumDescription("FTDv5", "FTDv10", "FTDv20", "FTDv30", "FTDv50", "FTDv100", "Legacy").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("FTDv5", "FTDv10", "FTDv20", "FTDv30", "FTDv50", "FTDv100", "Legacy"),
				},
			},
			"snort_engine": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SNORT engine version to be enabled.").AddStringEnumDescription("SNORT2", "SNORT3").String,
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
				MarkdownDescription: helpers.NewAttributeDescription("Id of the assigned Access Control Policy. For example `fmc_access_control_policy.example.id`.").String,
				Required:            true,
			},
			"nat_policy_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the assigned FTD NAT policy.").String,
				Optional:            true,
			},
			"health_policy_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the assigned Health policy. Every device requires health policy assignment, hence removal of this attribute does not trigger health policy de-assignment.").String,
				Optional:            true,
			},
			"container_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the parent container. Empty if device is Standalone.").String,
				Computed:            true,
			},
			"container_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the parent container (DeviceHAPair or DeviceCluster). Empty if device is Standalone.").String,
				Computed:            true,
			},
			"container_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the parent container. Empty if device is Standalone.").String,
				Computed:            true,
			},
			"container_role": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Role of the device in the container (PRIMARY, SECONDARY) for DeviceHAPair or (Control, Data) for DeviceCluster. Empty if device is Standalone.").String,
				Computed:            true,
			},
			"container_status": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Status of the device in DeviceHAPair (Active, Standby, but other possible as well).").String,
				Computed:            true,
			},
			"is_part_of_container": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("True if the device is part of a container (DeviceHAPair or DeviceCluster).").String,
				Computed:            true,
			},
			"is_multi_instance": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("True if the device is part of a multi-instance.").String,
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
	if r.client.IsCDFMC {
		resp.Diagnostics.AddError("Client Error", "UnsupportedVersion: Device resource is not supported in cdFMC. To register the device, please use Security Cloud Control API instead.")
		return
	}

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
	tflog.Debug(ctx, fmt.Sprintf("%s: Async task initiated successfully (id: %s)", plan.Id.ValueString(), taskID))

	diags = FMCWaitForJobToFinish(ctx, r.client, taskID, reqMods)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// ID of the device is known only when the registration is successfully completed.
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
		diags = r.updatePolicy(ctx, plan.Id.ValueString(), "Device", path.Root("nat_policy_id"), req.Plan, resp.State, reqMods...)
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

	// On device registration, default health policy is auto assigned. We are waiting till that is finished. (see loop above)
	// Health policy assignment triggers automatic deployment.
	if !plan.HealthPolicyId.IsNull() {
		diags = r.updatePolicy(ctx, plan.Id.ValueString(), "Device", path.Root("health_policy_id"), req.Plan, resp.State, reqMods...)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve Policy Assignments (GET), got error: %s, %s", err, policies.String()))
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

	// Read Unkonwns (that may change, as device may get added/removed from hapair/cluster)
	state.fromBodyUnknowns(ctx, res)

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

	if state.ContainerType.ValueString() == "DeviceHAPair" && state.ContainerStatus.ValueString() != "Active" {
		tflog.Info(ctx, fmt.Sprintf("%s: Device %s is in HA Pair, with current status: %s, hence cannot be updated. Configuration will be replicated from active node.", state.Id.ValueString(), state.Name.ValueString(), state.ContainerStatus.ValueString()))
		plan.copyComputed(ctx, state)
		diags = resp.State.Set(ctx, &plan)
		resp.Diagnostics.Append(diags...)
		return
	}

	if state.ContainerType.ValueString() == "DeviceCluster" && state.ContainerRole.ValueString() == "Data" {
		tflog.Info(ctx, fmt.Sprintf("%s: Device %s is in cluster as data node, hence cannot be updated. Configuration will be replicated from control node.", state.Id.ValueString(), state.Name.ValueString()))
		plan.copyComputed(ctx, state)
		diags = resp.State.Set(ctx, &plan)
		resp.Diagnostics.Append(diags...)
		return
	}

	body := plan.toBody(ctx, state)
	body, _ = sjson.Delete(body, "accessPolicy") // usable for POST, but not for PUT
	body, _ = sjson.Delete(body, "dummy_nat_policy_id")
	body, _ = sjson.Delete(body, "healthPolicy") // usable for POST, but not for PUT

	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	var deviceId, deviceType string
	if state.ContainerId.ValueString() != "" {
		deviceId = state.ContainerId.ValueString()
		deviceType = state.ContainerType.ValueString()
	} else {
		deviceId = plan.Id.ValueString()
		deviceType = "Device"
	}

	// Update policy assignments
	if plan.AccessPolicyId != state.AccessPolicyId {
		diags = r.updatePolicy(ctx, deviceId, deviceType, path.Root("access_policy_id"), req.Plan, req.State, reqMods...)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	if plan.NatPolicyId != state.NatPolicyId {
		diags = r.updatePolicy(ctx, deviceId, deviceType, path.Root("nat_policy_id"), req.Plan, req.State, reqMods...)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	if plan.HealthPolicyId != state.HealthPolicyId {
		diags = r.updatePolicy(ctx, deviceId, deviceType, path.Root("health_policy_id"), req.Plan, req.State, reqMods...)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	if plan.DeviceGroupId != state.DeviceGroupId {
		diags = FMCupdateDeviceGroup(ctx, r.client, plan.Id, req.Plan, req.State, reqMods)
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
func (r *DeviceResource) updatePolicy(ctx context.Context, deviceId string, deviceType string, policyPath path.Path, plan tfsdk.Plan, state tfsdk.State, reqMods ...func(*fmc.Req)) diag.Diagnostics {
	var planPolicy, statePolicy types.String

	if diags := plan.GetAttribute(ctx, policyPath, &planPolicy); diags.HasError() {
		return diags
	}

	if diags := state.GetAttribute(ctx, policyPath, &statePolicy); diags.HasError() {
		return diags
	}

	tflog.Debug(ctx, fmt.Sprintf("policy assignment %s, id %s, policy planned %s, state %s", policyPath, deviceId, planPolicy, statePolicy))

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
		// If assignment does not exist, do noting
		if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
			return nil
		} else if err != nil {
			return diag.Diagnostics{diag.NewErrorDiagnostic(
				"Client Error",
				fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()),
			)}
		}
		targets := res.Get(fmt.Sprintf(`targets.#(id != "%s")#`, deviceId))
		body, err := sjson.SetRaw(res.String(), "targets", targets.Raw)
		tflog.Debug(ctx, fmt.Sprintf("policy assignment with device %s removed: %s", deviceId, res))
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
		stubdev, _ := sjson.Set("{}", "id", deviceId)
		stubdev, _ = sjson.Set(stubdev, "type", deviceType)
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

	if res.Get(fmt.Sprintf(`targets.#(id=="%s")`, deviceId)).Exists() {
		tflog.Debug(ctx, fmt.Sprintf("target %s already assigned to policy %s", deviceId, planPolicy))
		return nil
	}

	polBody, err := sjson.Set(res.String(), `targets.-1`, map[string]any{
		"id":   deviceId,
		"type": deviceType,
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
	for range 3 {
		_, err := r.client.Delete("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords/"+url.QueryEscape(state.Id.ValueString()), reqMods...)
		if err != nil {
			if strings.Contains(err.Error(), "StatusCode 400") {
				// Check if device is under deployment
				tflog.Debug(ctx, fmt.Sprintf("%s: Checking if device is under deployment...", state.Id.ValueString()))
				diags := FMCWaitForDeploymentToFinish(ctx, r.client, []string{state.Id.ValueString()}, reqMods)
				if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
					return
				}
			}
		} else {
			// No error returned, break the loop
			break
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *DeviceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var config Device

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
		config.Domain = types.StringValue(tmpDomain)
	}
	config.Id = types.StringValue(match[inputPattern.SubexpIndex("id")])

	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
