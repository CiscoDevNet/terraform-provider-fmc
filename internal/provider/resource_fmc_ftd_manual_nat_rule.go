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
	_ resource.Resource                = &FTDManualNATRuleResource{}
	_ resource.ResourceWithImportState = &FTDManualNATRuleResource{}
)

func NewFTDManualNATRuleResource() resource.Resource {
	return &FTDManualNATRuleResource{}
}

type FTDManualNATRuleResource struct {
	client *fmc.Client
}

func (r *FTDManualNATRuleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ftd_manual_nat_rule"
}

func (r *FTDManualNATRuleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages a FTD Manual NAT Rule.").String,

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
			"ftd_nat_policy_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the FTD NAT Policy.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this value is always 'FTDManualNatRule'.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description of Manual NAT rule.").String,
				Optional:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates if the rule is enabled.").String,
				Optional:            true,
			},
			"section": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of section to which the rule belongs.").AddStringEnumDescription("BEFORE_AUTO", "AFTER_AUTO").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("BEFORE_AUTO", "AFTER_AUTO"),
				},
			},
			"nat_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the rule").AddStringEnumDescription("STATIC", "DYNAMIC").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("STATIC", "DYNAMIC"),
				},
			},
			"fall_through": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Fallthrough to Interface PAT (Destination Interface)").String,
				Optional:            true,
			},
			"interface_in_original_destination": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use interface address as original destination").String,
				Optional:            true,
			},
			"interface_in_translated_source": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Translate source network to destination interface address").String,
				Optional:            true,
			},
			"ipv6": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use the IPv6 address of the destination interface for interface PAT.").String,
				Optional:            true,
			},
			"net_to_net": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Net to Net Mapping").String,
				Optional:            true,
			},
			"no_proxy_arp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Do not proxy ARP on Destination Interface").String,
				Optional:            true,
			},
			"unidirectional": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether the rule is unidirectional").String,
				Optional:            true,
			},
			"source_interface_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of source security zone or interface group").String,
				Optional:            true,
			},
			"original_source_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of original source network object (host, network or range)").String,
				Optional:            true,
			},
			"original_source_port_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of original source port object").String,
				Optional:            true,
			},
			"original_destination_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of original destination network object (host, network or range)").String,
				Optional:            true,
			},
			"original_destination_port_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of original destination port object").String,
				Optional:            true,
			},
			"route_lookup": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Perform Route Lookup for Destination Interface").String,
				Optional:            true,
			},
			"destination_interface_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of destination security zone or interface group").String,
				Optional:            true,
			},
			"translated_source_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of translated source network object (host, network or range)").String,
				Optional:            true,
			},
			"translated_source_port_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of translated source port object").String,
				Optional:            true,
			},
			"translate_dns": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Translate DNS replies that match this rule").String,
				Optional:            true,
			},
			"translated_destination_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of translated destination network object (host, network or range)").String,
				Optional:            true,
			},
			"translated_destination_port_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of translated destination port object").String,
				Optional:            true,
			},
		},
	}
}

func (r *FTDManualNATRuleResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

func (r *FTDManualNATRuleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan FTDManualNATRule

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
	body := plan.toBody(ctx, FTDManualNATRule{})
	body = plan.adjustBody(ctx, body)
	res, err := r.client.Post(plan.getPath()+"?section="+strings.ToLower(plan.Section.ValueString()), body, reqMods...)
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

func (r *FTDManualNATRuleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state FTDManualNATRule

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

func (r *FTDManualNATRuleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state FTDManualNATRule

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

func (r *FTDManualNATRuleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state FTDManualNATRule

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
func (r *FTDManualNATRuleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var config FTDManualNATRule

	// Parse import ID
	var inputPattern = regexp.MustCompile(`^(?:(?P<domain>[^\s,]+),)?(?P<ftd_nat_policy_id>[^\s,]+),(?P<id>[^\s,]+?)$`)
	match := inputPattern.FindStringSubmatch(req.ID)
	if match == nil {
		errMsg := "Failed to parse import parameters.\nPlease provide import string in the following format: <domain>,<ftd_nat_policy_id>,<id>\n<domain> is optional.\n" + fmt.Sprintf("Got: %q", req.ID)
		resp.Diagnostics.AddError("Import error", errMsg)
		return
	}

	// Set domain, if provided
	if tmpDomain := match[inputPattern.SubexpIndex("domain")]; tmpDomain != "" {
		config.Domain = types.StringValue(tmpDomain)
	}
	config.Id = types.StringValue(match[inputPattern.SubexpIndex("id")])
	config.FtdNatPolicyId = types.StringValue(match[inputPattern.SubexpIndex("ftd_nat_policy_id")])

	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
