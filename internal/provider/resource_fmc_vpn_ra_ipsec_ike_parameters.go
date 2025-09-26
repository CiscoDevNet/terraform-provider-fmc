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
	"time"

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
	_ resource.Resource                = &VPNRAIPSecIKEParametersResource{}
	_ resource.ResourceWithImportState = &VPNRAIPSecIKEParametersResource{}
)

func NewVPNRAIPSecIKEParametersResource() resource.Resource {
	return &VPNRAIPSecIKEParametersResource{}
}

type VPNRAIPSecIKEParametersResource struct {
	client *fmc.Client
}

func (r *VPNRAIPSecIKEParametersResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_ra_ipsec_ike_parameters"
}

func (r *VPNRAIPSecIKEParametersResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages FTD Remote Access (RA) Virtual Private Networks (VPNs) IPSec/IKEv2 Parameters.").String,

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
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this value is always 'RaVpnIPsecAdvancedSetting'.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"ikev2_identity_sent_to_peer": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Identity sent to the peer during IKEv2 session establishment.").AddStringEnumDescription("IP_ADDRESS", "HOST_NAME", "AUTO_OR_DN").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("IP_ADDRESS", "HOST_NAME", "AUTO_OR_DN"),
				},
			},
			"ikev2_notification_on_tunnel_disconnect": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable notification on tunnel disconnect.").String,
				Optional:            true,
			},
			"ikev2_do_not_reboot_until_all_sessions_are_terminated": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Wait for all active sessions to voluntarily terminate before the system reboots.").String,
				Optional:            true,
			},
			"ikev2_cookie_challenge": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether to send cookie challenges to peer devices in response to SA initiated packets.").AddStringEnumDescription("CUSTOM", "ALWAYS", "NEVER").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("CUSTOM", "ALWAYS", "NEVER"),
				},
			},
			"ikev2_threshold_to_challenge_incoming_cookies": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Percentage of the total allowed SAs that are in-negotiation.").AddIntegerRangeDescription(1, 1000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 1000),
				},
			},
			"ikev2_number_of_sas_allowed_in_negotiation": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Limits the maximum number of SAs that can be in negotiation at any time.").AddIntegerRangeDescription(1, 100).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 100),
				},
			},
			"ikev2_maximum_number_of_sas_allowed": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Maximum number of Security Associations (SAs) allowed.").String,
				Optional:            true,
			},
			"ipsec_path_maximum_transmission_unit_aging_reset_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enter the number of minutes at which the Path Maximum Transission Unit (PMTU) value of an SA is reset to its original value.").AddIntegerRangeDescription(10, 30).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(10, 30),
				},
			},
			"nat_keepalive_message_traversal": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable NAT keepalive message traversal.").String,
				Optional:            true,
			},
			"nat_keepalive_message_traversal_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("NAT keepalive message traversal interval in seconds.").AddIntegerRangeDescription(10, 3600).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(10, 3600),
				},
			},
		},
	}
}

func (r *VPNRAIPSecIKEParametersResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

func (r *VPNRAIPSecIKEParametersResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan VPNRAIPSecIKEParameters

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
	body := plan.toBody(ctx, VPNRAIPSecIKEParameters{})
	body = plan.adjustBody(ctx, body)
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

func (r *VPNRAIPSecIKEParametersResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state VPNRAIPSecIKEParameters

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

func (r *VPNRAIPSecIKEParametersResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state VPNRAIPSecIKEParameters

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
	body = plan.adjustBody(ctx, body)
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

func (r *VPNRAIPSecIKEParametersResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state VPNRAIPSecIKEParameters

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

func (r *VPNRAIPSecIKEParametersResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <vpn_ra_id>,<id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("vpn_ra_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
