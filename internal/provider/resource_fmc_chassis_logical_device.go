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
	"time"

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
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
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages a Chassis Logical Device.\n Creating this resource will initiate a chassis-level deployment, triggering the device creation process based on the logical device configuration defined within this resource.\n Newly created device will be auto-registered with FMC.\n Destruction of the resource will de-register deployed device if it is registered to FMC and remove it from chassis.\n Adding or removing interfaces from logical device will trigger deployment to the chassis and logical device sync.\n Changing resource profile will not trigger automatic deployment to apply the settings.\n Currently, policies assignment is not supported at logical device level. Please use policy assignemnt resource.").String,

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
			"device_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the device that is deployed as result of this configuration.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"device_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the device that is deployed as result of this configuration; this value is always 'Device'.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the logical device. This is also a name of the device that will be deployed on the chassis.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"ftd_version": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Version of the device, that should be deployed. Image should be pre-deployed to the chassis.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"ipv4_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Management IPv4 address of the device.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"ipv4_netmask": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Netmask of Management IPv4 address.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"ipv4_gateway": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Gateway for Management IPv4 address.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"ipv6_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Management IPv6 address of the device.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"ipv6_prefix": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Prefix length of Management IPv6 address.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"ipv6_gateway": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Gateway for Management IPv6 address.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"search_domain": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Search domain for the device.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"fqdn": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Fully qualified domain name (FQDN) of the device.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"firewall_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Firewall mode of the device.").AddStringEnumDescription("ROUTED", "TRANSPARENT").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("ROUTED", "TRANSPARENT"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"dns_servers": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("DNS servers for the device. Up to three, comma-separated DNS servers can be specified.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"device_password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Admin password for the device.").String,
				Required:            true,
				Sensitive:           true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"admin_state": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Admin state of the device.").AddStringEnumDescription("ENABLED", "DISABLED").AddDefaultValueDescription("ENABLED").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("ENABLED", "DISABLED"),
				},
				Default: stringdefault.StaticString("ENABLED"),
			},
			"permit_expert_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Permit expert mode for the device.").AddStringEnumDescription("yes", "no").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("yes", "no"),
				},
			},
			"resource_profile_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the resource profile. Changing resource profile will trigger instance restart on deployment, however changing this value will not trigger automatic deployment.").String,
				Required:            true,
			},
			"resource_profile_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the resource profile. Changing resource profile will trigger instance restart on deployment, however changing this value will not trigger automatic deployment.").String,
				Required:            true,
			},
			"assigned_interfaces": schema.SetNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface assignment for the device.").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the interface.").String,
							Required:            true,
						},
					},
				},
			},
			"device_group_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the device group.").String,
				Optional:            true,
			},
			"access_control_policy_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the Access Control Policy (ACP) to be assigned to the device. This is used only as bootstrap configuration.").String,
				Required:            true,
			},
			"platform_settings_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the platform settings.").String,
				Optional:            true,
			},
			"licenses": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("License capabilities to be assigned to the device. This is used only as bootstrap configuration.").AddStringEnumDescription("MALWARE", "URLFilter", "CARRIER", "PROTECT", "THREAT").String,
				ElementType:         types.StringType,
				Optional:            true,
				Validators: []validator.Set{
					setvalidator.ValueStringsAre(
						stringvalidator.OneOf("MALWARE", "URLFilter", "CARRIER", "PROTECT", "THREAT"),
					),
				},
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

func (r *ChassisLogicalDeviceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ChassisLogicalDevice
	var res2 fmc.Res
	const atom time.Duration = 60 * time.Second

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

	// Create logical device configuration
	body := plan.toBody(ctx, ChassisLogicalDevice{})
	res, err := r.client.Post(plan.getPath(), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("id").String())

	// Trigger chassis deployment, to actually create the device
	deploy := DeviceDeploy{
		Id:            plan.Id,
		IgnoreWarning: types.BoolValue(true),
		DeviceIdList:  helpers.GetStringListFromStringSlice([]string{plan.ChassisId.ValueString()}),
	}
	diags = FMCDeviceDeploy(ctx, r.client, deploy, reqMods)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		// Something went wrong with deployment, we need to delete the logical device
		res, err := r.client.Delete(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), reqMods...)
		if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete logical device after failed deployment (DELETE), got error: %s, %s", err, res.String()))
		}
		return
	}

	// Deployment will trigger the creation of the logical device, but it will take some time. Wait for this to finish to get the device ID.
	for i := time.Duration(0); i < 45*time.Minute; i += atom {
		// FMCBUG: plan.getPath() endpoint also returns the device ID once deployed, but for whatever reason it does not show up if Get request is looped
		//         for that reason the devicerecords endpoint is used
		// filter needs to be exact device name, so search for `ftd` won't find `ftd-1`
		res2, err = r.client.Get("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords?filter=name:"+url.QueryEscape(plan.Name.ValueString()), reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to get device status, got error: %s, %s", err, res2.String()))
			return
		}

		// Once device is found, break the loop
		if res2.Get("items.0.id").Exists() {
			break
		}
		time.Sleep(atom)
	}

	// Check if the device was created
	if !res2.Get("items.0.id").Exists() {
		// Something went wrong with deployment, we need to delete the logical device
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("%s: Failed to get deployed device Id", plan.Id.ValueString()))
		res, err := r.client.Delete(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), reqMods...)
		if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete logical device after failed deployment (DELETE), got error: %s, %s", err, res.String()))
		}
		return
	}

	plan.fromBodyUnknowns(ctx, res)

	// Get deployed device details
	plan.DeviceId = types.StringValue(res2.Get("items.0.id").String())
	plan.DeviceType = types.StringValue(res2.Get("items.0.type").String())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

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

	// If interfaces are updated, trigger chassis deployment
	changed, diags := helpers.IsConfigUpdatingAt(ctx, req.Plan, req.State, path.Root("assigned_interfaces"))
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}
	if changed {
		// Trigger chassis deployment
		deploy := DeviceDeploy{
			Id:            plan.Id,
			IgnoreWarning: types.BoolValue(true),
			DeviceIdList:  helpers.GetStringListFromStringSlice([]string{plan.ChassisId.ValueString()}),
		}
		diags = FMCDeviceDeploy(ctx, r.client, deploy, reqMods)
		if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
			return
		}

		// Sync device to get the interfaces available
		interfaceEventsPath := "/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords/" + url.QueryEscape(plan.DeviceId.ValueString()) + "/interfaceevents"
		tflog.Debug(ctx, fmt.Sprintf("%s: Syncing device", plan.Id.ValueString()))
		body, _ := sjson.Set("", "action", "SYNC_WITH_DEVICE")
		res, err := r.client.Post(interfaceEventsPath, body, reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to sync device (POST), got error: %s, %s", err, res.String()))
			return
		}

		// Save changes to the device
		if res.Get("hasPendingChanges").Exists() && res.Get("hasPendingChanges").Bool() {
			tflog.Debug(ctx, fmt.Sprintf("%s: Saving changes to device", plan.Id.ValueString()))
			body, _ = sjson.Set("", "action", "ACCEPT_CHANGES")
			res, err := r.client.Post(interfaceEventsPath, body, reqMods...)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save the device configuration (POST), got error: %s, %s", err, res.String()))
				return
			}
		}
	}

	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

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

	// Try to unregister the device first
	if state.DeviceId.ValueString() != "" {
		for range 5 {
			res, err := r.client.Delete("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords/"+url.QueryEscape(state.DeviceId.ValueString()), reqMods...)
			if err != nil {
				if strings.Contains(strings.ToLower(res.Raw), "error retrieving device from database") {
					// Device is not registered, break the loop
					break
				}

				if strings.Contains(strings.ToLower(res.Raw), "deployment is in progress") {
					// Deployment is in progress, wait and try again
					tflog.Debug(ctx, fmt.Sprintf("%s: Device is still being deployed, waiting...", state.Id.ValueString()))
					diags := FMCWaitForDeploymentToFinish(ctx, r.client, []string{state.DeviceId.ValueString()}, reqMods)
					if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
						return
					}
					continue
				}
			}
			// No error returned, break the loop
			break
		}
	}

	// Delete logical device configuration
	res, err := r.client.Delete(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString()), reqMods...)
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	// Trigger chassis deployment, to actually delete the device
	deploy := DeviceDeploy{
		Id:            state.Id,
		IgnoreWarning: types.BoolValue(true),
		DeviceIdList:  helpers.GetStringListFromStringSlice([]string{state.ChassisId.ValueString()}),
	}
	diags = FMCDeviceDeploy(ctx, r.client, deploy, reqMods)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *ChassisLogicalDeviceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Parse import ID
	var inputPattern = regexp.MustCompile(`^(?:(?P<domain>[^\s,]+),)?(?P<chassis_id>[^\s,]+),(?P<id>[^\s,]+?)$`)
	match := inputPattern.FindStringSubmatch(req.ID)
	if match == nil {
		errMsg := "Failed to parse import parameters.\nPlease provide import string in the following format: <domain>,<chassis_id>,<id>\n<domain> is optional. If not provided, `Global` is used implicitly and resource's `domain` attribute is not set.\n" + fmt.Sprintf("Got: %q", req.ID)
		resp.Diagnostics.AddError("Import error", errMsg)
		return
	}

	// Set domain, if provided
	if tmpDomain := match[inputPattern.SubexpIndex("domain")]; tmpDomain != "" {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("domain"), tmpDomain)...)
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), match[inputPattern.SubexpIndex("id")])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("chassis_id"), match[inputPattern.SubexpIndex("chassis_id")])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
