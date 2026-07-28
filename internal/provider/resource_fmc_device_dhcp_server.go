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
	_ resource.Resource                = &DeviceDHCPServerResource{}
	_ resource.ResourceWithImportState = &DeviceDHCPServerResource{}
)

func NewDeviceDHCPServerResource() resource.Resource {
	return &DeviceDHCPServerResource{}
}

type DeviceDHCPServerResource struct {
	client *fmc.Client
}

func (r *DeviceDHCPServerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_dhcp_server"
}

func (r *DeviceDHCPServerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages a Device DHCP Server.").AddMinimumVersionHeaderDescription().AddMinimumVersionDescription("7.4").String,

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
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this value is always 'DHCPServerSettings'.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"ping_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Ping timeout in milliseconds.").AddIntegerRangeDescription(10, 10000).AddDefaultValueDescription("50").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(10, 10000),
				},
				Default: int64default.StaticInt64(50),
			},
			"lease_length": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Duration (in seconds) that a client can use an assigned IP address before it must renew the lease.").AddIntegerRangeDescription(300, 1048575).AddDefaultValueDescription("3600").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(300, 1048575),
				},
				Default: int64default.StaticInt64(3600),
			},
			"auto_config_interface_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the interface to be used for autoconfiguration. Id, name, and type must all be specified together.").String,
				Optional:            true,
			},
			"auto_config_interface_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the interface to be used for autoconfiguration. Id, name, and type must all be specified together.").String,
				Optional:            true,
			},
			"auto_config_interface_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the interface to be used for autoconfiguration. Id, name, and type must all be specified together.").String,
				Optional:            true,
			},
			"domain_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Domain name given to DHCP clients.").String,
				Optional:            true,
			},
			"primary_dns_server_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the Host object used as the primary DNS server offered to DHCP clients.").String,
				Optional:            true,
			},
			"secondary_dns_server_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the Host object used as the secondary DNS server offered to DHCP clients.").String,
				Optional:            true,
			},
			"primary_wins_server_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the Host object used as the primary WINS server offered to DHCP clients.").String,
				Optional:            true,
			},
			"secondary_wins_server_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the Host object used as the secondary WINS server offered to DHCP clients.").String,
				Optional:            true,
			},
			"servers": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of per-interface DHCP server pools.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the interface on which the DHCP server is enabled.").String,
							Required:            true,
						},
						"interface_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the interface on which the DHCP server is enabled.").String,
							Required:            true,
						},
						"interface_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of the interface on which the DHCP server is enabled.").String,
							Required:            true,
						},
						"address_pool": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Range of IP addresses (formatted as `start_ip-end_ip`) the DHCP server can hand out. The pool must be on the same subnet as the associated interface.").String,
							Required:            true,
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable the DHCP server on the associated interface.").String,
							Optional:            true,
						},
					},
				},
			},
			"dhcp_options": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of advanced DHCP option codes (parameters) sent to DHCP clients.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"code": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("DHCP option code.").AddIntegerRangeDescription(2, 255).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(2, 255),
							},
						},
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of the value carried by the option. `IP` uses `first_ip_address_id` (and optionally `second_ip_address_id`), `ASCII` uses `ascii` and `HEX` uses `hex`.").AddStringEnumDescription("IP", "ASCII", "HEX").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("IP", "ASCII", "HEX"),
							},
						},
						"first_ip_address_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the Host object used as the first IP address, when `type` is `IP`.").String,
							Optional:            true,
						},
						"second_ip_address_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the Host object used as the second IP address, when `type` is `IP`.").String,
							Optional:            true,
						},
						"ascii": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ASCII value, when `type` is `ASCII`.").String,
							Optional:            true,
						},
						"hex": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Hexadecimal value, when `type` is `HEX`.").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *DeviceDHCPServerResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *DeviceDHCPServerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	// Check if FMC client is connected to supports this object
	if r.client.FMCVersionParsed.LessThan(minFMCVersionDeviceDHCPServer) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support Device DHCP Server creation, minumum required version is 7.4", r.client.FMCVersion))
		return
	}
	var plan DeviceDHCPServer

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
	body := plan.toBody(ctx, DeviceDHCPServer{})
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

func (r *DeviceDHCPServerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Check if FMC client is connected to supports this object
	if r.client.FMCVersionParsed.LessThan(minFMCVersionDeviceDHCPServer) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("UnsupportedVersion: FMC version %s does not support Device DHCP Server, minimum required version is 7.4", r.client.FMCVersion))
		return
	}
	var state DeviceDHCPServer

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

func (r *DeviceDHCPServerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state DeviceDHCPServer

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

func (r *DeviceDHCPServerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DeviceDHCPServer

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
func (r *DeviceDHCPServerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
