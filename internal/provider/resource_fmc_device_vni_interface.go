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
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
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
	_ resource.Resource                = &DeviceVNIInterfaceResource{}
	_ resource.ResourceWithImportState = &DeviceVNIInterfaceResource{}
)

func NewDeviceVNIInterfaceResource() resource.Resource {
	return &DeviceVNIInterfaceResource{}
}

type DeviceVNIInterfaceResource struct {
	client *fmc.Client
}

func (r *DeviceVNIInterfaceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_vni_interface"
}

func (r *DeviceVNIInterfaceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages a Device VNI Interface.").String,

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
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"vni_id": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("User-created VNI number for the interface, not exposed over the wire.").AddIntegerRangeDescription(1, 10000).String,
				Required:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 10000),
				},
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"multicast_group_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Can only be set when VNI interface is mapped to VTEP source interface with `neighbor_discovery` equal to DEFAULT_MULTICAST_GROUP. If unset, the default group from the VTEP source interface is used.").String,
				Optional:            true,
			},
			"segment_id": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("VNI tag value used in packets over the wire. If null, the `enable_proxy` must be true.").AddIntegerRangeDescription(1, 16777215).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 16777215),
				},
			},
			"nve_number": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("VTEP policy NVE number. If null, not mapped to a VTEP.").AddIntegerRangeDescription(1, 10000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 10000),
				},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether to enable the interface.").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"logical_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Customizable logical name of the interface, unique on the device. Should not contain whitespace or slash characters. Can only be used when `segment_id` is set.").String,
				Optional:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Optional user-created description.").String,
				Optional:            true,
			},
			"mtu": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Maximum transmission unit. Can only be used when logical_name is set on the parent interface.").String,
				Optional:            true,
			},
			"priority": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Priority.").AddIntegerRangeDescription(0, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
				},
			},
			"security_zone_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the assigned security zone. Can only be used when `logical_name` is set.").String,
				Optional:            true,
			},
			"ipv4_static_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Static IPv4 address.").String,
				Optional:            true,
			},
			"ipv4_static_netmask": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Netmask (width) for ipv4_static_address.").String,
				Optional:            true,
			},
			"ipv4_dhcp_obtain_default_route": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Any non-null value here indicates to enable DHCPv4. Value `false` indicates to enable DHCPv4 without obtaining default IPv4 route but anyway requires also `ipv4_dhcp_route_metric` to be set to exactly 1. Value `true` indicates to enable DHCPv4 and obtain the route and also requires `ipv4_dhcp_route_metric` to be non-null. The `ipv4_dhcp_obtain_default_route` must be null when using `ipv4_static_address`.").String,
				Optional:            true,
			},
			"ipv4_dhcp_default_route_metric": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The metric for `ipv4_dhcp_obtain_default_route`. Any non-null value enables DHCP as a side effect. Must be null when using `ipv4_static_address`.").AddIntegerRangeDescription(1, 255).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
			},
			"ipv6": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether to enable IPv6.").String,
				Optional:            true,
			},
			"ipv6_enforce_eui": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether to enforce IPv6 Extended Unique Identifier (EUI64 from RFC2373).").String,
				Optional:            true,
			},
			"ipv6_auto_config": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether to enable IPv6 autoconfiguration.").String,
				Optional:            true,
			},
			"ipv6_dhcp_address": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether to enable DHCPv6 for address config.").String,
				Optional:            true,
			},
			"ipv6_dhcp_nonaddress": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether to enable DHCPv6 for non-address config.").String,
				Optional:            true,
			},
			"ipv6_ra": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether to enable IPv6 router advertisement (RA).").String,
				Optional:            true,
			},
			"ipv6_addresses": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv6 address without a slash and prefix.").String,
							Required:            true,
						},
						"prefix": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Prefix width for the IPv6 address.").String,
							Required:            true,
						},
						"enforce_eui": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Indicates whether to enforce IPv6 Extended Unique Identifier (EUI64 from RFC2373).").String,
							Optional:            true,
						},
					},
				},
			},
			"proxy": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether to enable proxy.").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
}

func (r *DeviceVNIInterfaceResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

func (r *DeviceVNIInterfaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DeviceVNIInterface

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
	body := plan.toBody(ctx, DeviceVNIInterface{})
	body, _ = sjson.Set(body, "name", fmt.Sprintf("vni%d", plan.VniId.ValueInt64()))

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

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *DeviceVNIInterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DeviceVNIInterface

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

func (r *DeviceVNIInterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state DeviceVNIInterface

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
	body, _ = sjson.Set(body, "name", fmt.Sprintf("vni%d", plan.VniId.ValueInt64()))

	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *DeviceVNIInterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DeviceVNIInterface

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
func (r *DeviceVNIInterfaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
