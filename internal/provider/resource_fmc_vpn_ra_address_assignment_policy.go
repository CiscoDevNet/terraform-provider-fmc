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
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &VPNRAAddressAssignmentPolicyResource{}
	_ resource.ResourceWithImportState = &VPNRAAddressAssignmentPolicyResource{}
)

func NewVPNRAAddressAssignmentPolicyResource() resource.Resource {
	return &VPNRAAddressAssignmentPolicyResource{}
}

type VPNRAAddressAssignmentPolicyResource struct {
	client *fmc.Client
}

func (r *VPNRAAddressAssignmentPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_ra_address_assignment_policy"
}

func (r *VPNRAAddressAssignmentPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages FTD Remote Access (RA) Virtual Private Networks (VPNs) Client Address Assignment Policies.").String,

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
			"vpn_ra_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the parent VPN RA Configuration.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this value is always 'RaVpnAddressAssignmentSetting'.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"ipv4_use_authorization_server": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use authorization server (Only for RADIUS or Realm).").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ipv4_use_dhcp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use DHCP for IPv4 address assignment.").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ipv4_internal_address_pool": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use internal address pool for IPv4 address assignment.").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ipv4_internal_address_pool_reuse_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interval in seconds for reusing IPv4 addresses.").AddIntegerRangeDescription(0, 480).AddDefaultValueDescription("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 480),
				},
				Default: int64default.StaticInt64(0),
			},
			"ipv6_use_authorization_server": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use authorization server (Only for RADIUS or Realm).").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ipv6_internal_address_pool": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use internal address pool for IPv6 address assignment.").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
		},
	}
}

func (r *VPNRAAddressAssignmentPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

func (r *VPNRAAddressAssignmentPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan VPNRAAddressAssignmentPolicy

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

	// FMCBUG CSCwq61583 FMC API: RAVPN sub-endpoints are unstable
	var resId, res fmc.Res
	var err error
	var val []gjson.Result
	for range 5 {
		for range 5 {
			resId, err = r.client.Get(plan.getPath(), reqMods...)
			if err == nil {
				break
			}
			if !strings.Contains(err.Error(), "StatusCode 404") && !strings.Contains(err.Error(), "StatusCode 400") {
				break
			}
			time.Sleep(5 * time.Second)
		}

		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}

		// Check if exactly one object is returned
		val = resId.Get("items").Array()
		if len(val) == 0 {
			time.Sleep(5 * time.Second)
			continue
		}
		if len(val) != 1 {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Expected 1 object, got %d", len(val)))
			return
		}
		break
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
	body := plan.toBody(ctx, VPNRAAddressAssignmentPolicy{})
	urlPath := plan.getPath() + "/" + url.PathEscape(plan.Id.ValueString())
	for range 5 {
		res, err = r.client.Put(urlPath, body, reqMods...)
		if err == nil {
			break
		}
		if !strings.Contains(err.Error(), "StatusCode 404") && !strings.Contains(err.Error(), "StatusCode 400") {
			break
		}
		time.Sleep(5 * time.Second)
	}
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

func (r *VPNRAAddressAssignmentPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state VPNRAAddressAssignmentPolicy

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

	// FMCBUG CSCwq61583 FMC API: RAVPN sub-endpoints are unstable
	var res fmc.Res
	var err error
	for range 5 {
		res, err = r.client.Get(urlPath, reqMods...)
		if err == nil {
			break
		}
		if !strings.Contains(err.Error(), "StatusCode 404") && !strings.Contains(err.Error(), "StatusCode 400") {
			break
		}
		time.Sleep(5 * time.Second)
	}

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

func (r *VPNRAAddressAssignmentPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state VPNRAAddressAssignmentPolicy

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
	// FMCBUG CSCwq61583 FMC API: RAVPN sub-endpoints are unstable
	var res fmc.Res
	var err error
	urlPath := plan.getPath() + "/" + url.QueryEscape(plan.Id.ValueString())
	for range 5 {
		res, err = r.client.Put(urlPath, body, reqMods...)
		if err == nil {
			break
		}
		if !strings.Contains(err.Error(), "StatusCode 404") && !strings.Contains(err.Error(), "StatusCode 400") {
			break
		}
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *VPNRAAddressAssignmentPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state VPNRAAddressAssignmentPolicy

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
func (r *VPNRAAddressAssignmentPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var config VPNRAAddressAssignmentPolicy

	// Parse import ID
	var inputPattern = regexp.MustCompile(`^(?:(?P<domain>[^\s,]+),)?(?P<vpn_ra_id>[^\s,]+),(?P<id>[^\s,]+?)$`)
	match := inputPattern.FindStringSubmatch(req.ID)
	if match == nil {
		errMsg := "Failed to parse import parameters.\nPlease provide import string in the following format: <domain>,<vpn_ra_id>,<id>\n<domain> is optional.\n" + fmt.Sprintf("Got: %q", req.ID)
		resp.Diagnostics.AddError("Import error", errMsg)
		return
	}

	// Set domain, if provided
	if tmpDomain := match[inputPattern.SubexpIndex("domain")]; tmpDomain != "" {
		config.Domain = types.StringValue(tmpDomain)
	}
	config.Id = types.StringValue(match[inputPattern.SubexpIndex("id")])
	config.VpnRaId = types.StringValue(match[inputPattern.SubexpIndex("vpn_ra_id")])

	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
