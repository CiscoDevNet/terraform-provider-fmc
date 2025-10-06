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
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
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
				MarkdownDescription: helpers.NewAttributeDescription("Name of the Group Policy object. Use `DfltGrpPolicy` to manage the default group policy.").String,
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this value is always 'GroupPolicy'.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description of the object.").String,
				Optional:            true,
			},
			"protocol_ssl": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable SSL protocol for VPN connections.").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"protocol_ipsec_ikev2": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable IPsec IKEv2 protocol for VPN connections.").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ipv4_address_pools": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of IPv4 Address Pools for address assignment.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Pool Id.").String,
							Required:            true,
						},
					},
				},
				Validators: []validator.List{
					listvalidator.SizeAtMost(6),
				},
			},
			"banner": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Banner text to be displayed to users.").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 491),
				},
			},
			"primary_dns_server_host_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of host object that represents primary DNS server.").String,
				Optional:            true,
			},
			"secondary_dns_server_host_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of host object that represents secondary DNS server.").String,
				Optional:            true,
			},
			"primary_wins_server_host_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of host object that represents primary WINS server.").String,
				Optional:            true,
			},
			"secondary_wins_server_host_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of host object that represents secondary WINS server.").String,
				Optional:            true,
			},
			"dhcp_network_scope_network_object_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the Network Object used to determine the DHCP scope.").String,
				Optional:            true,
			},
			"default_domain": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the default domain.").String,
				Optional:            true,
			},
			"ipv4_split_tunnel_policy": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv4 split tunnel policy.").AddStringEnumDescription("TUNNEL_ALL", "TUNNEL_SPECIFIED", "EXCLUDE_SPECIFIED_OVER_TUNNEL").AddDefaultValueDescription("TUNNEL_ALL").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("TUNNEL_ALL", "TUNNEL_SPECIFIED", "EXCLUDE_SPECIFIED_OVER_TUNNEL"),
				},
				Default: stringdefault.StaticString("TUNNEL_ALL"),
			},
			"ipv6_split_tunnel_policy": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv6 split tunnel policy.").AddStringEnumDescription("TUNNEL_ALL", "TUNNEL_SPECIFIED", "EXCLUDE_SPECIFIED_OVER_TUNNEL").AddDefaultValueDescription("TUNNEL_ALL").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("TUNNEL_ALL", "TUNNEL_SPECIFIED", "EXCLUDE_SPECIFIED_OVER_TUNNEL"),
				},
				Default: stringdefault.StaticString("TUNNEL_ALL"),
			},
			"split_tunnel_acl_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of standard or extended ACL used for split tunnel configuration.").String,
				Optional:            true,
			},
			"split_tunnel_acl_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of ACL used for split tunnel configuration. Mandatory, when `split_tunnel_acl_id` is set.").AddStringEnumDescription("StandardAccessList", "ExtendedAccessList").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("StandardAccessList", "ExtendedAccessList"),
				},
			},
			"dns_request_split_tunnel_policy": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Define if DNS requests should be send over the tunnel or not.").AddStringEnumDescription("USE_SPLIT_TUNNEL_SETTING", "TUNNEL_ALL", "TUNNEL_SPECIFIED_DOMAINS").AddDefaultValueDescription("USE_SPLIT_TUNNEL_SETTING").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("USE_SPLIT_TUNNEL_SETTING", "TUNNEL_ALL", "TUNNEL_SPECIFIED_DOMAINS"),
				},
				Default: stringdefault.StaticString("USE_SPLIT_TUNNEL_SETTING"),
			},
			"split_dns_domain_list": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Up to 10, comma separated domains for split DNS requests.").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 255),
				},
			},
			"secure_client_profile_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of the Secure Client Profile.").String,
				Optional:            true,
			},
			"secure_client_management_profile_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of the Secure Client Management Profile.").String,
				Optional:            true,
			},
			"secure_client_modules": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of Secure Client Modules to be enabled.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("AMP_ENABLER", "FEEDBACK", "ISE_POSTURE", "NETWORK_ACCESS_MANAGER", "NETWORK_VISIBILITY", "UMBRELLA_ROAMING", "WEB_SECURITY", "START_BEFORE_LOGIN", "DART").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("AMP_ENABLER", "FEEDBACK", "ISE_POSTURE", "NETWORK_ACCESS_MANAGER", "NETWORK_VISIBILITY", "UMBRELLA_ROAMING", "WEB_SECURITY", "START_BEFORE_LOGIN", "DART"),
							},
						},
						"profile_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of the module profile.").String,
							Optional:            true,
						},
						"download_module": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable module download.").AddDefaultValueDescription("true").String,
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(true),
						},
					},
				},
			},
			"ssl_compression": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SSL compression method for the connection.").AddStringEnumDescription("DISABLED", "DEFLATE", "LZS").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("DISABLED", "DEFLATE", "LZS"),
				},
			},
			"dtls_compression": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("DTLS compression method for the connection.").AddStringEnumDescription("DISABLED", "LZS").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("DISABLED", "LZS"),
				},
			},
			"mtu_size": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Maximum Transmission Unit (MTU) size for SSL connections.").AddIntegerRangeDescription(576, 1462).AddDefaultValueDescription("1406").String,
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
			"keep_alive_messages": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Keepalive Messages between Secure Client and VPN gateway.").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"keep_alive_messages_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Keepalive message interval in seconds.").AddIntegerRangeDescription(15, 600).AddDefaultValueDescription("20").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(15, 600),
				},
				Default: int64default.StaticInt64(20),
			},
			"gateway_dpd": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable VPN secure gateway Dead Peer Detection (DPD).").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"gateway_dpd_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("VPN secure gateway Dead Peer Detection (DPD) messages interval in seconds.").AddIntegerRangeDescription(5, 3600).AddDefaultValueDescription("30").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(5, 3600),
				},
				Default: int64default.StaticInt64(30),
			},
			"client_dpd": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable VPN client Dead Peer Detection (DPD).").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"client_dpd_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("VPN client Dead Peer Detection (DPD) messages interval in seconds.").AddIntegerRangeDescription(5, 3600).AddDefaultValueDescription("30").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(5, 3600),
				},
				Default: int64default.StaticInt64(30),
			},
			"client_bypass_protocol": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Drop network traffic for which the headend did not assign an IP address. Applicable if headend assigned only IPv4 or only IPv6 address.").String,
				Optional:            true,
			},
			"ssl_rekey": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enables the client to rekey the connection.").String,
				Optional:            true,
			},
			"ssl_rekey_method": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Method to use for SSL rekeying.").AddStringEnumDescription("NEW_TUNNEL", "EXISTING_TUNNEL").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("NEW_TUNNEL", "EXISTING_TUNNEL"),
				},
			},
			"ssl_rekey_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interval for SSL rekeying in minutes.").AddIntegerRangeDescription(4, 10080).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(4, 10080),
				},
			},
			"client_firewall_private_network_rules_acl_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of extended ACL to configure firewall settings for the VPN client's platform.").String,
				Optional:            true,
			},
			"client_firewall_public_network_rules_acl_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of extended ACL to configure firewall settings for the VPN client's platform.").String,
				Optional:            true,
			},
			"secure_client_custom_attributes": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Secure Client Custom Attributes that are used by the Secure Client to configure features.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the Custom Attribute.").String,
							Required:            true,
						},
					},
				},
			},
			"traffic_filter_acl_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of Extended ACL that determine whether to allow or block tunneled data packets coming through the VPN connection.").String,
				Optional:            true,
			},
			"restrict_vpn_to_vlan": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specifies the egress VLAN ID for sessions to which this Group Policy applies.").AddIntegerRangeDescription(1, 4094).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4094),
				},
			},
			"access_hours_time_range_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of Time Range object that specifies when this group policy is available to be applied to a remote access user.").String,
				Optional:            true,
			},
			"simultaneous_logins_per_user": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Maximum number of simultaneous logins allowed for a user.").AddIntegerRangeDescription(0, 2147483647).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 2147483647),
				},
			},
			"maximum_connection_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Maximum user connection time in minutes.").AddIntegerRangeDescription(1, 4473924).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4473924),
				},
			},
			"maximum_connection_time_alert_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specifies the interval of time before maximum connection time is reached to display a message to the user.").AddIntegerRangeDescription(1, 30).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 30),
				},
			},
			"idle_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("VPN idle timeout in minutes.").AddIntegerRangeDescription(1, 4473924).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4473924),
				},
			},
			"idle_timeout_alert_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interval of time before idle time is reached to display a message to the user.").AddIntegerRangeDescription(1, 30).String,
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

	var res fmc.Res
	var err error

	if plan.Name.ValueString() == "DfltGrpPolicy" {
		// This is a default group policy, we will update it instead of creating a new one.
		tflog.Debug(ctx, fmt.Sprintf("%s: Detected DfltGrpPolicy, updating instead of creating", plan.Id.ValueString()))
		// Get id of the default group policy
		resTmp, err := r.client.Get(plan.getPath()+"?filter=nameOrValue:DfltGrpPolicy", reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve DfltGrpPolicy (GET), got error: %s, %s", err, resTmp.String()))
			return
		}
		if len(resTmp.Get("items").Array()) != 1 {
			resp.Diagnostics.AddError("Error", fmt.Sprintf("Got %d items for DfltGrpPolicy, expected 1", len(resTmp.Get("items").Array())))
			return
		}
		plan.Id = types.StringValue(resTmp.Get("items").Array()[0].Get("id").String())
		// Update object
		body := plan.toBody(ctx, GroupPolicy{})
		res, err = r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body, reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure DfltGrpPolicy (PUT), got error: %s, %s", err, res.String()))
			return
		}
	} else {
		// Create object
		body := plan.toBody(ctx, GroupPolicy{})
		res, err = r.client.Post(plan.getPath(), body, reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
			return
		}
		plan.Id = types.StringValue(res.Get("id").String())
	}
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

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
	if state.Name.ValueString() == "DfltGrpPolicy" {
		tflog.Debug(ctx, fmt.Sprintf("%s: Detected DfltGrpPolicy, clearing configuration", state.Id.ValueString()))
		body := state.toBodyPutDelete(ctx)
		res, err := r.client.Put(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString()), body, reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to clear DfltGrpPolicy (PUT), got error: %s, %s", err, res.String()))
			return
		}
	} else {
		res, err := r.client.Delete(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString()), reqMods...)
		if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *GroupPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var config GroupPolicy

	// Parse import ID
	var inputPattern = regexp.MustCompile(`^(?:(?P<domain>[^\s,]+),)?(?P<id>[^\s,]+?)$`)
	match := inputPattern.FindStringSubmatch(req.ID)
	if match == nil {
		errMsg := "Failed to parse import parameters.\nPlease provide import string in the following format: <domain>,<id>\n<domain> is optional. If not provided, `Global` is used implicitly and resource's `domain` attribute is not set.\n" + fmt.Sprintf("Got: %q", req.ID)
		resp.Diagnostics.AddError("Import error", errMsg)
		return
	}

	// Set domain, if provided
	if tmpDomain := match[inputPattern.SubexpIndex("domain")]; tmpDomain != "" {
		config.Domain = types.StringValue(tmpDomain)
	}
	config.Id = types.StringValue(match[inputPattern.SubexpIndex("id")])

	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
