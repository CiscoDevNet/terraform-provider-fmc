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
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &VPNRAIPSecCryptoMapResource{}
	_ resource.ResourceWithImportState = &VPNRAIPSecCryptoMapResource{}
)

func NewVPNRAIPSecCryptoMapResource() resource.Resource {
	return &VPNRAIPSecCryptoMapResource{}
}

type VPNRAIPSecCryptoMapResource struct {
	client *fmc.Client
}

func (r *VPNRAIPSecCryptoMapResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_ra_ipsec_crypto_map"
}

func (r *VPNRAIPSecCryptoMapResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages FTD Remote Access (RA) Virtual Private Networks (VPNs) IPSec Crypto Maps.").String,

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
				MarkdownDescription: helpers.NewAttributeDescription("Id of the parent VPN RA Topology.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this value is always 'RaVpnIPsecCryptoMap'.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"interface_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the interface object.").String,
				Required:            true,
			},
			"ikev2_ipsec_proposals": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of IKEv2 IPSec proposals").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the IKEv2 IPSec proposal.").String,
							Required:            true,
						},
					},
				},
			},
			"reverse_route_injection": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Reverse Route Injection (RRI).").String,
				Optional:            true,
			},
			"client_services": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Client Services.").String,
				Optional:            true,
			},
			"client_services_port": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Port for Client Services.").AddIntegerRangeDescription(1, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
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
				Required:            true,
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

func (r *VPNRAIPSecCryptoMapResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *VPNRAIPSecCryptoMapResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan VPNRAIPSecCryptoMap

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

	tflog.Debug(ctx, fmt.Sprintf("%s: considering object interface_id %s", plan.Id, plan.InterfaceId))
	if plan.Id.ValueString() == "" && plan.InterfaceId.ValueString() != "" {
		offset := 0
		limit := 1000
		for page := 1; ; page++ {
			queryString := fmt.Sprintf("?limit=%d&offset=%d&expanded=true", limit, offset)
			res, err := r.client.Get(plan.getPath()+queryString, reqMods...)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects, got error: %s", err))
				return
			}
			if value := res.Get("items"); len(value.Array()) > 0 {
				value.ForEach(func(k, v gjson.Result) bool {
					if plan.InterfaceId.ValueString() == v.Get("interfaceObject.id").String() {
						plan.Id = types.StringValue(v.Get("id").String())
						tflog.Debug(ctx, fmt.Sprintf("%s: Found object with interface_id '%v', id: %s", plan.Id.ValueString(), plan.InterfaceId.ValueString(), plan.Id.ValueString()))
						return false
					}
					return true
				})
			}
			if plan.Id.ValueString() != "" || !res.Get("paging.next.0").Exists() {
				break
			}
			offset += limit
		}

		if plan.Id.ValueString() == "" {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with interface_id: %v", plan.InterfaceId.ValueString()))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, VPNRAIPSecCryptoMap{})
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

func (r *VPNRAIPSecCryptoMapResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state VPNRAIPSecCryptoMap

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

func (r *VPNRAIPSecCryptoMapResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state VPNRAIPSecCryptoMap

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

func (r *VPNRAIPSecCryptoMapResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state VPNRAIPSecCryptoMap

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

func (r *VPNRAIPSecCryptoMapResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

// Section below is generated&owned by "gen/generator.go". //template:begin createSubresources

// End of section. //template:end createSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin deleteSubresources

// End of section. //template:end deleteSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin updateSubresources

// End of section. //template:end updateSubresources
