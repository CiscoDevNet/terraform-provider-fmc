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
	_ resource.Resource                = &GroupPolicyResource{}
	_ resource.ResourceWithImportState = &GroupPolicyResource{}
)

func NewGroupPolicyResource() resource.Resource {
	return &GroupPolicyResource{}
}

type GroupPolicyResource struct {
	client *fmc.Client
}

func (r *GroupPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_group_policy"
}

func (r *GroupPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages a Group Policy.").String,

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
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the Group Policy object.").String,
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this value is always ''.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description of the object.").String,
				Optional:            true,
			},
			"enable_ssl_protocol": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether the SSL protocol is enabled.").String,
				Optional:            true,
			},
			"enable_ipsec_ikev2_protocol": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether the IPsec IKEv2 protocol is enabled.").String,
				Optional:            true,
			},
			"ipv4_address_pools": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of IPv4 address pools for the Group Policy.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Unique identifier for the IPv4 address pool.").String,
							Required:            true,
						},
					},
				},
			},
			"banner": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Banner text to be displayed to users.").String,
				Optional:            true,
			},
			"primary_dns_server": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Primary DNS server for the Group Policy.").String,
				Optional:            true,
			},
			"secondary_dns_server": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Secondary DNS server for the Group Policy.").String,
				Optional:            true,
			},
			"primary_wins_server": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Primary WINS server for the Group Policy.").String,
				Optional:            true,
			},
			"secondary_wins_server": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Secondary WINS server for the Group Policy.").String,
				Optional:            true,
			},
			"dhcp_scope_network_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Network ID for the DHCP scope.").String,
				Optional:            true,
			},
			"default_domain": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Default domain name for the Group Policy.").String,
				Optional:            true,
			},
			"ipv4_split_tunnel_policy": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("TUNNEL_ALL", "TUNNEL_SPECIFIED", "EXCLUDE_SPECIFIED_OVER_TUNNEL").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("TUNNEL_ALL", "TUNNEL_SPECIFIED", "EXCLUDE_SPECIFIED_OVER_TUNNEL"),
				},
			},
			"ipv6_split_tunnel_policy": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("TUNNEL_ALL", "TUNNEL_SPECIFIED", "EXCLUDE_SPECIFIED_OVER_TUNNEL").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("TUNNEL_ALL", "TUNNEL_SPECIFIED", "EXCLUDE_SPECIFIED_OVER_TUNNEL"),
				},
			},
			"split_tunnel_acl_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ACL ID for the split tunnel configuration.").String,
				Optional:            true,
			},
			"split_d_n_s_request_policy": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("USE_SPLIT_TUNNEL_SETTING", "TUNNEL_ALL", "TUNNEL_SPECIFIED_DOMAINS").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("USE_SPLIT_TUNNEL_SETTING", "TUNNEL_ALL", "TUNNEL_SPECIFIED_DOMAINS"),
				},
			},
			"split_d_n_s_domain_list": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of domains for split DNS requests.").String,
				Optional:            true,
			},
			"secure_client_client_profile_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of the Secure Client profile.").String,
				Optional:            true,
			},
			"secure_client_management_profile_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of the Secure Client management profile.").String,
				Optional:            true,
			},
			"ssl_compression": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("DISABLED", "DEFLATE", "LZS").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("DISABLED", "DEFLATE", "LZS"),
				},
			},
			"dtls_compression": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("DISABLED", "DEFLATE", "LZS").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("DISABLED", "DEFLATE", "LZS"),
				},
			},
			"mtu_size": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Maximum Transmission Unit size for SSL connections.").AddIntegerRangeDescription(576, 1462).AddDefaultValueDescription("1406").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(576, 1462),
				},
				Default: int64default.StaticInt64(1406),
			},
			"ignore_df_bit": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether to ignore the Don't Fragment bit in packets.").String,
				Optional:            true,
			},
			"enable_keep_alive_messages": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether to enable keep-alive messages for the connection.").String,
				Optional:            true,
			},
			"keep_alive_message_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interval for keep-alive messages in seconds.").AddIntegerRangeDescription(15, 600).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(15, 600),
				},
			},
			"enable_gateway_dpd": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether to enable Dead Peer Detection (DPD) for the connection.").String,
				Optional:            true,
			},
			"gateway_dpd_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interval for Dead Peer Detection (DPD) messages in seconds.").AddIntegerRangeDescription(5, 3600).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(5, 3600),
				},
			},
			"enable_client_dpd": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether to enable Dead Peer Detection (DPD) for the connection.").String,
				Optional:            true,
			},
			"client_dpd_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interval for Dead Peer Detection (DPD) messages in seconds.").AddIntegerRangeDescription(5, 3600).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(5, 3600),
				},
			},
			"client_bypass_protocol": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
			},
			"enable_ssl_rekey": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether to enable SSL rekeying.").String,
				Optional:            true,
			},
			"rekey_method": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Method to use for SSL rekeying.").AddStringEnumDescription("NEW_TUNNEL", "EXISTING_TUNNEL").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("NEW_TUNNEL", "EXISTING_TUNNEL"),
				},
			},
			"rekey_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interval for SSL rekeying in minutes.").AddIntegerRangeDescription(4, 10080).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(4, 10080),
				},
			},
			"client_firewall_private_network_rules_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of the client firewall private network rules.").String,
				Optional:            true,
			},
			"client_firewall_public_network_rules_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of the client firewall public network rules.").String,
				Optional:            true,
			},
			"custom_attributes": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of custom attributes for the Group Policy.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of the custom attribute.").String,
							Required:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("AnyConnect specific custom attribute.").AddStringEnumDescription("PER_APP_VPN", "ALLOW_DEFER_UPDATE", "DYNAMIC_SPLIT_TUNNELING", "CUSTOM_TYPE").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("PER_APP_VPN", "ALLOW_DEFER_UPDATE", "DYNAMIC_SPLIT_TUNNELING", "CUSTOM_TYPE"),
							},
						},
					},
				},
			},
			"traffic_filter_acl_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ACL ID for the traffic filter.").String,
				Optional:            true,
			},
			"restrict_vpn_to_vlan_id": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("VLAN ID to restrict VPN access.").AddIntegerRangeDescription(1, 4094).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4094),
				},
			},
			"access_hours_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of the access hours settings.").String,
				Optional:            true,
			},
			"simultaneous_login_per_user": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Maximum number of simultaneous logins allowed per user.").AddIntegerRangeDescription(0, 2147483647).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 2147483647),
				},
			},
			"max_connection_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Maximum connection timeout in minutes.").AddIntegerRangeDescription(1, 4473924).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4473924),
				},
			},
			"max_connection_time_alert_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Alert interval for maximum connection time in minutes.").AddIntegerRangeDescription(1, 30).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 30),
				},
			},
			"vpn_idle_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("VPN idle timeout in minutes.").AddIntegerRangeDescription(1, 4473924).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4473924),
				},
			},
			"vpn_idle_timeout_alert_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Alert interval for VPN idle timeout in minutes.").AddIntegerRangeDescription(1, 30).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 30),
				},
			},
		},
	}
}

func (r *GroupPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *GroupPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan GroupPolicy

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
	body := plan.toBody(ctx, GroupPolicy{})
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

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *GroupPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state GroupPolicy

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

func (r *GroupPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state GroupPolicy

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

func (r *GroupPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state GroupPolicy

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

func (r *GroupPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin createSubresources

// End of section. //template:end createSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin deleteSubresources

// End of section. //template:end deleteSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin updateSubresources

// End of section. //template:end updateSubresources
