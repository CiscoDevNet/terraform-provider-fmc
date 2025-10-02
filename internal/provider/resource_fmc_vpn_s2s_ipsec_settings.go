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
	_ resource.Resource                = &VPNS2SIPSECSettingsResource{}
	_ resource.ResourceWithImportState = &VPNS2SIPSECSettingsResource{}
)

func NewVPNS2SIPSECSettingsResource() resource.Resource {
	return &VPNS2SIPSECSettingsResource{}
}

type VPNS2SIPSECSettingsResource struct {
	client *fmc.Client
}

func (r *VPNS2SIPSECSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_s2s_ipsec_settings"
}

func (r *VPNS2SIPSECSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages FTD Site-to-Site (S2S) Virtual Private Networks (VPNs) IPSEC settings.").String,

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
			"vpn_s2s_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the parent VPN S2S Topology.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this value is always 'IPSecSetting'.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"crypto_map_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the crypto map.").AddStringEnumDescription("STATIC", "DYNAMIC").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("STATIC", "DYNAMIC"),
				},
			},
			"ikev2_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IKEv2 mode.").AddStringEnumDescription("TUNNEL", "TRANSPORT_PREFERRED", "TRANSPORT_REQUIRE").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("TUNNEL", "TRANSPORT_PREFERRED", "TRANSPORT_REQUIRE"),
				},
			},
			"ikev1_ipsec_proposals": schema.SetNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set of IKEv1 IPSEC proposals.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the IKEv1 IPSEC proposal").String,
							Required:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the IKEv1 IPSEC proposal").String,
							Required:            true,
						},
					},
				},
			},
			"ikev2_ipsec_proposals": schema.SetNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set of IKEv2 IPSEC proposals.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the IKEv2 IPSEC proposal").String,
							Required:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the IKEv2 IPSEC proposal").String,
							Required:            true,
						},
					},
				},
			},
			"security_association_strength_enforcement": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Security Association (SA) strength enforcement.").String,
				Optional:            true,
			},
			"reverse_route_injection": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Reverse Route Injection (RRI).").String,
				Optional:            true,
			},
			"perfect_forward_secrecy": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable IPSEC Perfect Forward Secrecy (PFS).").String,
				Optional:            true,
			},
			"perfect_forward_secrecy_modulus_group": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Modulus group for IPSEC Perfect Forward Secrecy (PFS).").AddStringEnumDescription("1", "2", "5", "14", "15", "16", "19", "20", "21", "24", "31").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1", "2", "5", "14", "15", "16", "19", "20", "21", "24", "31"),
				},
			},
			"lifetime_duration": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Number of seconds a security association exists before expiring.").AddIntegerRangeDescription(120, 2147483647).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(120, 2147483647),
				},
			},
			"lifetime_size": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Volume of traffic (in kilobytes) that can pass between IPsec peers using a given security association before it expires.").AddIntegerRangeDescription(10, 2147483647).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(10, 2147483647),
				},
			},
			"validate_incoming_icmp_error_messages": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable incoming ICMP error messages validation.").String,
				Optional:            true,
			},
			"do_not_fragment_policy": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Policy for handling Do Not Fragment (DNF) packets.").AddStringEnumDescription("SET", "COPY", "CLEAR", "NONE").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("SET", "COPY", "CLEAR", "NONE"),
				},
			},
			"tfc": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Traffic Flow Confidentiality (TFC) packets.").String,
				Optional:            true,
			},
			"tfc_burst_bytes": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Burst size in bytes for TFC packets. Set 0 for `auto` or value in range 1-16.").AddIntegerRangeDescription(0, 16).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 16),
				},
			},
			"tfc_payload_bytes": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Payload size in bytes for TFC packets. Set 0 for `auto` or value in range 64-1024.").AddIntegerRangeDescription(0, 1024).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 1024),
				},
			},
			"tfc_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Timeout duration in seconds for TFC packets. Set 0 for `auto` or value in range 10-60.").AddIntegerRangeDescription(0, 60).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 60),
				},
			},
		},
	}
}

func (r *VPNS2SIPSECSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *VPNS2SIPSECSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan VPNS2SIPSECSettings

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
	body := plan.toBody(ctx, VPNS2SIPSECSettings{})
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

func (r *VPNS2SIPSECSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state VPNS2SIPSECSettings

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

func (r *VPNS2SIPSECSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state VPNS2SIPSECSettings

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

func (r *VPNS2SIPSECSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state VPNS2SIPSECSettings

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
func (r *VPNS2SIPSECSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var config VPNS2SIPSECSettings

	// Parse import ID
	var inputPattern = regexp.MustCompile(`^(?:(?P<domain>[^\s,]+),)?(?P<vpn_s2s_id>[^\s,]+),(?P<id>[^\s,]+?)$`)
	match := inputPattern.FindStringSubmatch(req.ID)
	if match == nil {
		errMsg := "Failed to parse import parameters.\nPlease provide import string in the following format: <domain>,<vpn_s2s_id>,<id>\n<domain> is optional. If not provided, `Global` is used implicitly and resource's `domain` attribute is not set.\n" + fmt.Sprintf("Got: %q", req.ID)
		resp.Diagnostics.AddError("Import error", errMsg)
		return
	}

	// Set domain, if provided
	if tmpDomain := match[inputPattern.SubexpIndex("domain")]; tmpDomain != "" {
		config.Domain = types.StringValue(tmpDomain)
	}
	config.Id = types.StringValue(match[inputPattern.SubexpIndex("id")])
	config.VpnS2sId = types.StringValue(match[inputPattern.SubexpIndex("vpn_s2s_id")])

	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
