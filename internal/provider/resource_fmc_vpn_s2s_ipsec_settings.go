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
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
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
				MarkdownDescription: helpers.NewAttributeDescription("The type of the crypto map.").AddStringEnumDescription("STATIC", "DYNAMIC").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("STATIC", "DYNAMIC"),
				},
			},
			"ikev2_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The IKEv2 mode for the IPSEC settings.").AddStringEnumDescription("TUNNEL", "TRANSPORT_PREFERRED", "TRANSPORT_REQUIRE").String,
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
			"enable_sa_strength_enforcement": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether Security Association (SA) strength enforcement is enabled.").String,
				Optional:            true,
			},
			"enable_reverse_route_injection": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether Route Redundancy Interface (RRI) is enabled.").String,
				Optional:            true,
			},
			"enable_perfect_forward_secrecy": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether IPSEC Perfect Forward Secrecy (PFS) is enabled.").String,
				Optional:            true,
			},
			"perfect_forward_secrecy_modulus_group": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The modulus group for IPSEC Perfect Forward Secrecy (PFS).").AddStringEnumDescription("1", "2", "5", "14", "15", "16", "19", "20", "21", "24", "31").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1", "2", "5", "14", "15", "16", "19", "20", "21", "24", "31"),
				},
			},
			"lifetime_duration": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The lifetime duration for the IPSEC settings in seconds.").AddIntegerRangeDescription(120, 2147483647).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(120, 2147483647),
				},
			},
			"lifetime_size": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The lifetime size for the IPSEC settings in kilobytes.").AddIntegerRangeDescription(10, 2147483647).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(10, 2147483647),
				},
			},
			"validate_incoming_icmp_error_messages": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether to validate incoming ICMP error messages.").String,
				Optional:            true,
			},
			"do_not_fragment_policy": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The policy for handling Do Not Fragment (DNF) packets.").AddStringEnumDescription("SET", "COPY", "CLEAR", "NONE").AddDefaultValueDescription("NONE").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("SET", "COPY", "CLEAR", "NONE"),
				},
				Default: stringdefault.StaticString("NONE"),
			},
			"enable_tfc_packets": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether Traffic Flow Confidentiality (TFC) packets are enabled.").String,
				Optional:            true,
			},
			"tfc_burst_bytes": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The burst size in bytes for TFC packets. Set 0 for `auto`").AddIntegerRangeDescription(0, 16).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 16),
				},
			},
			"tfc_payload_bytes": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The payload size in bytes for TFC packets. Set 0 for `auto`, or set to 64-1024.").AddIntegerRangeDescription(0, 1024).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 1024),
				},
			},
			"tfc_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The timeout duration in seconds for TFC packets. Set 0 for `auto`, or set to 10-60.").AddIntegerRangeDescription(1, 60).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 60),
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

	// Get ID
	res, err := r.client.Get(plan.getPath(), reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve ike settings, got error: %s", err))
		return
	}
	if val := res.Get("items.0.id"); val.Exists() {
		plan.Id = types.StringValue(val.String())
		tflog.Debug(ctx, fmt.Sprintf("%s: Found object", plan.Id))
	} else {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to extract ike settings from payload: %s", res.String()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, VPNS2SIPSECSettings{})
	res, err = r.client.Put(plan.getPath()+"/"+url.PathEscape(plan.Id.ValueString()), body, reqMods...)
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

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import

func (r *VPNS2SIPSECSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <vpn_s2s_id>,<id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("vpn_s2s_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin createSubresources

// End of section. //template:end createSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin deleteSubresources

// End of section. //template:end deleteSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin updateSubresources

// End of section. //template:end updateSubresources
