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
	_ resource.Resource                = &VPNS2SAdvancedSettingsResource{}
	_ resource.ResourceWithImportState = &VPNS2SAdvancedSettingsResource{}
)

func NewVPNS2SAdvancedSettingsResource() resource.Resource {
	return &VPNS2SAdvancedSettingsResource{}
}

type VPNS2SAdvancedSettingsResource struct {
	client *fmc.Client
}

func (r *VPNS2SAdvancedSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_s2s_advanced_settings"
}

func (r *VPNS2SAdvancedSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages FTD Site-to-Site (S2S) Virtual Private Networks (VPNs) Advanced settings.").String,

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
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this value is always 'AdvancedSetting'.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"ike_keepalive": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable IKE keepalives.").AddStringEnumDescription("DISABLED", "ENABLED", "ENABLED_INFINITE").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("DISABLED", "ENABLED", "ENABLED_INFINITE"),
				},
			},
			"ike_keepalive_threshold": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("IKE keepalive threshold in seconds. This interval is the number of seconds allowing a peer to idle before beginning keepalive monitoring.").AddIntegerRangeDescription(10, 3600).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(10, 3600),
				},
			},
			"ike_keepalive_retry_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("IKE keepalive retry interval in seconds.").AddIntegerRangeDescription(2, 10).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(2, 10),
				},
			},
			"ike_identity_sent_to_peers": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("identity that the peers will use to identify themselves during IKE negotiations.").AddStringEnumDescription("IP_ADDRESS", "HOST_NAME", "AUTO_OR_DN").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("IP_ADDRESS", "HOST_NAME", "AUTO_OR_DN"),
				},
			},
			"ike_peer_identity_validation": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Peer identity validation.").AddStringEnumDescription("DO_NOT_CHECK", "REQUIRED", "IF_SUPPORTED_BY_CERT").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("DO_NOT_CHECK", "REQUIRED", "IF_SUPPORTED_BY_CERT"),
				},
			},
			"ike_aggressive_mode": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable IKEv1 aggressive mode.").String,
				Optional:            true,
			},
			"ike_notification_on_tunnel_disconnect": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable notification on tunnel disconnect.").String,
				Optional:            true,
			},
			"ikev2_cookie_challenge": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Send cookie challenges to peer devices in response to SA initiate packets.").AddStringEnumDescription("CUSTOM", "ALWAYS", "NEVER").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("CUSTOM", "ALWAYS", "NEVER"),
				},
			},
			"ikev2_threshold_to_challenge_incoming_cookies": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The percentage of the total allowed SAs that are in-negotiation.").AddIntegerRangeDescription(1, 100).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 100),
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
				MarkdownDescription: helpers.NewAttributeDescription("Limits the number of allowed IKEv2 connections.").String,
				Optional:            true,
			},
			"ipsec_fragmentation_before_encryption": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable fragmentation before encryption.").String,
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
			"vpn_idle_timeout": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable VPN idle timeout monitoring.").String,
				Optional:            true,
			},
			"vpn_idle_timeout_value": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("VPN idle timeout in minutes.").AddIntegerRangeDescription(1, 35791394).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 35791394),
				},
			},
			"bypass_access_control_traffic_for_decrypted_traffic": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable bypass access control traffic for decrypted traffic (sysopt permit-vpn).").String,
				Optional:            true,
			},
			"cert_use_map_configured_in_endpoint_to_determine_tunnel": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use certificate map configured in endpoint to determine tunnel.").String,
				Optional:            true,
			},
			"cert_use_ou_to_determine_tunnel": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use certificate OU to determine tunnel.").String,
				Optional:            true,
			},
			"cert_use_ike_identity_to_determine_tunnel": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use IKE identity OU to determine tunnel.").String,
				Optional:            true,
			},
			"cert_use_peer_ip_address_to_determine_tunnel": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use peer IP address to determine tunnel.").String,
				Optional:            true,
			},
		},
	}
}

func (r *VPNS2SAdvancedSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *VPNS2SAdvancedSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan VPNS2SAdvancedSettings

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
	body := plan.toBody(ctx, VPNS2SAdvancedSettings{})
	body = plan.adjustBody(ctx, body)
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

func (r *VPNS2SAdvancedSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state VPNS2SAdvancedSettings

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

func (r *VPNS2SAdvancedSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state VPNS2SAdvancedSettings

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

func (r *VPNS2SAdvancedSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state VPNS2SAdvancedSettings

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

func (r *VPNS2SAdvancedSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
