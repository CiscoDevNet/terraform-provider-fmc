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
	_ resource.Resource                = &RouteMapResource{}
	_ resource.ResourceWithImportState = &RouteMapResource{}
)

func NewRouteMapResource() resource.Resource {
	return &RouteMapResource{}
}

type RouteMapResource struct {
	client *fmc.Client
}

func (r *RouteMapResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_route_map"
}

func (r *RouteMapResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages a Route Map.").String,

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
				MarkdownDescription: helpers.NewAttributeDescription("Name of the Route Map object.").String,
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this value is always 'RouteMap'.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"overridable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether the object values can be overridden.").String,
				Optional:            true,
			},
			"entries": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of route map entries.").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"action": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Indicate the redistribution access.").AddStringEnumDescription("PERMIT", "DENY").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("PERMIT", "DENY"),
							},
						},
						"match_security_zones": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match traffic based on the (ingress/egress) Security Zones.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("ID of the object.").String,
										Required:            true,
									},
								},
							},
						},
						"match_interface_names": schema.ListAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match traffic based on the (ingress/egress) interface names.").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"match_ipv4_address_access_lists": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match routes based on the route address using Standard or Extended IPv4 Access Lists.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the Access List.").String,
										Required:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of the Access List.").String,
										Required:            true,
									},
								},
							},
						},
						"match_ipv4_address_prefix_lists": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match routes based on the route address using Prefix Lists.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the Prefix List.").String,
										Required:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of the Prefix List.").String,
										Required:            true,
									},
								},
							},
						},
						"match_ipv4_next_hop_access_lists": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match routes based on the next hop address of a route using Standard or Extended IPv4 Access Lists.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the Access List.").String,
										Optional:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of the Access List.").String,
										Required:            true,
									},
								},
							},
						},
						"match_ipv4_next_hop_prefix_lists": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match routes based on the next hop address of a route using Prefix Lists.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the Prefix List.").String,
										Required:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of the Prefix List.").String,
										Required:            true,
									},
								},
							},
						},
						"match_ipv4_route_source_access_lists": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match routes based on the advertising source address of the route using Standard or Extended IPv4 Access List.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the Access List.").String,
										Optional:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of the Access List.").String,
										Required:            true,
									},
								},
							},
						},
						"match_ipv4_route_source_prefix_lists": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match routes based on the advertising source address of the route using Prefix Lists.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the Prefix List.").String,
										Required:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of the Prefix List.").String,
										Required:            true,
									},
								},
							},
						},
						"match_ipv6_address_extended_access_list_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match routes based on the route address using IPv6 Extended Access Lists.").String,
							Optional:            true,
						},
						"match_ipv6_address_prefix_list_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match routes based on the route address using IPv6 Prefix Lists.").String,
							Optional:            true,
						},
						"match_ipv6_next_hop_extended_access_list_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match routes based on the next hop address of a route using IPv6 Extended Access Lists.").String,
							Optional:            true,
						},
						"match_ipv6_next_hop_prefix_list_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match routes based on the next hop address of a route using IPv6 Prefix Lists.").String,
							Optional:            true,
						},
						"match_ipv6_route_source_extended_access_list_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match routes based on the advertising source address of the route using IPv6 Extended Access Lists.").String,
							Optional:            true,
						},
						"match_ipv6_route_source_prefix_list_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match routes based on the advertising source address of the route using IPv6 Prefix Lists.").String,
							Optional:            true,
						},
						"match_bgp_as_path_lists": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match a BGP Autonomous System (AS) path with the specified path access list.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of object.").String,
										Optional:            true,
									},
								},
							},
						},
						"match_bgp_community_lists": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match a BGP Community with Standard/Expanded Community Lists.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the object.").String,
										Optional:            true,
									},
								},
							},
						},
						"match_bgp_extended_community_lists": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match a BGP Community with Extended Community Lists.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the object.").String,
										Optional:            true,
									},
								},
							},
						},
						"match_bgp_policy_lists": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Evaluate and process a BGP Policy Lists.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the object.").String,
										Optional:            true,
									},
								},
							},
						},
						"match_metric_route_values": schema.ListAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of Metric values to match.").String,
							ElementType:         types.Int64Type,
							Optional:            true,
						},
						"match_tag_values": schema.ListAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of Tag values to match.").String,
							ElementType:         types.Int64Type,
							Optional:            true,
						},
						"match_route_type_external_1": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match external type 1 routes.").String,
							Optional:            true,
						},
						"match_route_type_external_2": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match external type 2 routes.").String,
							Optional:            true,
						},
						"match_route_type_internal": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match internal routes.").String,
							Optional:            true,
						},
						"match_route_type_local": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match local routes.").String,
							Optional:            true,
						},
						"match_route_type_nssa_external_1": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match NSSA external type 1 routes.").String,
							Optional:            true,
						},
						"match_route_type_nssa_external_2": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Match NSSA external type 2 routes.").String,
							Optional:            true,
						},
						"set_metric_bandwidth": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set the metric bandwidth value in Kbits per second.").AddIntegerRangeDescription(0, 4294967295).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 4294967295),
							},
						},
						"set_metric_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Specify the type of metric for the destination routing protocol.").AddStringEnumDescription("INTERNAL", "TYPE_1", "TYPE_2").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("INTERNAL", "TYPE_1", "TYPE_2"),
							},
						},
						"set_bgp_as_path_prepend": schema.ListAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Prepend an arbitrary Autonomous System (AS) path to BGP routes.").String,
							ElementType:         types.Int64Type,
							Optional:            true,
						},
						"set_bgp_as_path_prepend_last_as": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Number of times to prepend the AS path with the last AS number.").AddIntegerRangeDescription(0, 10).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 10),
							},
						},
						"set_bgp_as_path_convert_route_tag_into_as_path": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Convert the tag of a route into an Autonomous System (AS) path.").String,
							Optional:            true,
						},
						"set_bgp_community_none": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set the specific community to none.").String,
							Optional:            true,
						},
						"set_bgp_community_specific_community": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set the specific community.").AddIntegerRangeDescription(1, 4294967295).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4294967295),
							},
						},
						"set_bgp_community_add_to_existing_communities": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Add the community to the already existing communities.").String,
							Optional:            true,
						},
						"set_bgp_community_internet": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Use Internet well-known community.").String,
							Optional:            true,
						},
						"set_bgp_community_no_advertise": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Use No-Advertise well-known community.").String,
							Optional:            true,
						},
						"set_bgp_community_no_export": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Use No-Export well-known community.").String,
							Optional:            true,
						},
						"set_bgp_community_route_target": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set Route Target number in format ASN:nn format (range 1:1 to 65534:65535). Separate multiple values with a comma.").String,
							Optional:            true,
						},
						"set_bgp_community_add_to_existing_extended_communities": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set the extended community additive.").String,
							Optional:            true,
						},
						"set_bgp_automatic_tag": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Automatically compute the tag value.").String,
							Optional:            true,
						},
						"set_bgp_local_preference": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set the local preference value.").AddIntegerRangeDescription(1, 4294967295).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4294967295),
							},
						},
						"set_bgp_weight": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set the weight value.").AddIntegerRangeDescription(0, 65535).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65535),
							},
						},
						"set_bgp_origin": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Specify the BGP origin code.").AddStringEnumDescription("LOCAL_IGP", "INCOMPLETE").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("LOCAL_IGP", "INCOMPLETE"),
							},
						},
						"set_bgp_ipv4_next_hop": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set the next hop IPv4 address.").AddStringEnumDescription("USE_PEER_ADDRESS", "SPECIFIC_IP").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("USE_PEER_ADDRESS", "SPECIFIC_IP"),
							},
						},
						"set_bgp_ipv4_next_hop_specific_ip": schema.ListAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set the next hop IPv4 address.").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"set_bgp_ipv4_prefix_list_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set the prefix list for IPv4.").String,
							Optional:            true,
						},
						"set_bgp_ipv6_next_hop": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set the next hop IPv6 address.").AddStringEnumDescription("USE_PEER_ADDRESS", "SPECIFIC_IP").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("USE_PEER_ADDRESS", "SPECIFIC_IP"),
							},
						},
						"set_bgp_ipv6_next_hop_specific_ip": schema.ListAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set the next hop IPv6 address.").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"set_bgp_ipv6_prefix_list_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set the prefix list for IPv6.").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *RouteMapResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *RouteMapResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RouteMap

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
	body := plan.toBody(ctx, RouteMap{})
	body = plan.adjustBody(ctx, body)
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

func (r *RouteMapResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RouteMap

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

func (r *RouteMapResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RouteMap

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

func (r *RouteMapResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RouteMap

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

func (r *RouteMapResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
