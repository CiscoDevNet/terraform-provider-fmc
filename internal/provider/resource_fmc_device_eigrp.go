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
	"github.com/hashicorp/terraform-plugin-framework/path"
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
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &DeviceEIGRPResource{}
	_ resource.ResourceWithImportState = &DeviceEIGRPResource{}
)

func NewDeviceEIGRPResource() resource.Resource {
	return &DeviceEIGRPResource{}
}

type DeviceEIGRPResource struct {
	client *fmc.Client
}

func (r *DeviceEIGRPResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_eigrp"
}

func (r *DeviceEIGRPResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages a Device EIGRP.").String,

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
			"device_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Id of the parent device.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this value is always 'EigrpRoute'.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"as_number": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Autonomous System (AS) number of the EIGRP process.").String,
				Required:            true,
			},
			"networks": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Networks or Hosts that participate in the EIGRP routing process.").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Network object ID.").String,
							Required:            true,
						},
					},
				},
			},
			"auto_summary": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Advertise summary routes to other EIGRP routers.").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"passive_interfaces_all": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set all interfaces as passive (do not send or receive routing updates).").String,
				Optional:            true,
			},
			"passive_interfaces": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interfaces as passive (do not send or receive routing updates). Applicable only when `passive_interfaces_all` is not set to true.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Interface ID.").String,
							Required:            true,
						},
					},
				},
			},
			"neighbors": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("EIGRP neighbors.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of interface through which the neighbor is reachable.").String,
							Required:            true,
						},
						"host_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of Host object representing the neighbor.").String,
							Required:            true,
						},
					},
				},
			},
			"filter_rules": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Filter rules (distribute lists) for EIGRP routing updates.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"filter_direction": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Filter direction.").AddStringEnumDescription("INBOUND", "OUTBOUND").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("INBOUND", "OUTBOUND"),
							},
						},
						"standard_access_list_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Standard Access List ID, that defines the networks that are to be received and suppressed in routing updates.").String,
							Required:            true,
						},
						"interface_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of interface to which the filter rule applies.").String,
							Optional:            true,
						},
						"routing_process": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Routing process whose routes are filtered. Applicable only for `OUTBOUND` direction.").AddStringEnumDescription("CONNECTED", "BGP", "OSPF", "RIP", "STATIC").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("CONNECTED", "BGP", "OSPF", "RIP", "STATIC"),
							},
						},
						"routing_process_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("OSPF process ID or BGP AS number.").String,
							Optional:            true,
						},
					},
				},
			},
			"redistributions": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Redistributions to the EIGRP routing process.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"redistribute_protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Protocol to redistribute.").AddStringEnumDescription("CONNECTED", "BGP", "OSPF", "RIP", "STATIC").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("CONNECTED", "BGP", "OSPF", "RIP", "STATIC"),
							},
						},
						"process_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("OSPF process ID or BGP AS number.").String,
							Optional:            true,
						},
						"route_map_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Route map ID for route filtering.").String,
							Optional:            true,
						},
						"match_internal": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether to match OSPF internal routes.").String,
							Optional:            true,
						},
						"match_external_1": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether to match OSPF external type 1 routes.").String,
							Optional:            true,
						},
						"match_external_2": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether to match OSPF external type 2 routes.").String,
							Optional:            true,
						},
						"match_nssa_external_1": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether to match OSPF NSSA external type 1 routes.").String,
							Optional:            true,
						},
						"match_nssa_external_2": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether to match OSPF NSSA external type 2 routes.").String,
							Optional:            true,
						},
						"metric_bandwidth": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Metric bandwidth (in Kbps).").AddIntegerRangeDescription(1, 4294967295).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4294967295),
							},
						},
						"metric_delay": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Metric delay (in tens of microseconds).").AddIntegerRangeDescription(0, 4294967295).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 4294967295),
							},
						},
						"metric_reliability": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Metric reliability. The value 255 indicates 100 percent reliability; 0 means no reliability.").AddIntegerRangeDescription(0, 255).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 255),
							},
						},
						"metric_loading": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Metric effective bandwidth (loading). 255 indicates 100 percent loading.").AddIntegerRangeDescription(1, 255).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 255),
							},
						},
						"metric_mtu": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Metric path MTU (in bytes).").AddIntegerRangeDescription(1, 65535).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65535),
							},
						},
					},
				},
			},
			"summary_addresses": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Summary addresses advertised on interfaces.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of interface on which the summary address is advertised.").String,
							Required:            true,
						},
						"network_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Network object ID of the summary address.").String,
							Required:            true,
						},
						"administrative_distance": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Administrative distance of the summary route.").AddIntegerRangeDescription(1, 255).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 255),
							},
						},
					},
				},
			},
			"interfaces": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("EIGRP interface settings.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Interface ID.").String,
							Required:            true,
						},
						"interface_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Interface type.").String,
							Required:            true,
						},
						"hello_interval": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Interval (in seconds) between hello transmissions.").AddIntegerRangeDescription(1, 65535).AddDefaultValueDescription("5").String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65535),
							},
							Default: int64default.StaticInt64(5),
						},
						"hold_time": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Time (in seconds) before the neighbor is considered down.").AddIntegerRangeDescription(1, 65535).AddDefaultValueDescription("15").String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65535),
							},
							Default: int64default.StaticInt64(15),
						},
						"split_horizon": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable split horizon.").String,
							Optional:            true,
						},
						"delay_time": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Throughput delay (in tens of microseconds).").AddIntegerRangeDescription(1, 16777215).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 16777215),
							},
						},
						"md5_authentication": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable MD5 authentication.").String,
							Optional:            true,
						},
						"authentication_key_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Authentication key encryption.").AddStringEnumDescription("UNENCRYPTED", "ENCRYPTED", "AUTHKEY").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("UNENCRYPTED", "ENCRYPTED", "AUTHKEY"),
							},
						},
						"authentication_key_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Authentication key ID.").AddIntegerRangeDescription(0, 255).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 255),
							},
						},
						"authentication_key": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Authentication key.").String,
							Optional:            true,
							Sensitive:           true,
						},
					},
				},
			},
			"router_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv4 address used as the router ID. Do not configure for AUTOMATIC router ID selection.").String,
				Optional:            true,
			},
			"default_information_in_standard_access_list_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Standard Access List ID used to accept exterior default routing information.").String,
				Optional:            true,
			},
			"default_information_out_standard_access_list_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Standard Access List ID used to advertise exterior default routing information.").String,
				Optional:            true,
			},
			"administrative_distance_internal": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Administrative distance for EIGRP internal routes.").AddIntegerRangeDescription(1, 255).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
			},
			"administrative_distance_external": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Administrative distance for EIGRP external routes.").AddIntegerRangeDescription(1, 255).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
			},
			"log_neighbor_changes": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable logging of EIGRP neighbor adjacency changes.").String,
				Optional:            true,
			},
			"log_neighbor_warnings": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable logging of EIGRP neighbor warning messages.").String,
				Optional:            true,
			},
			"log_neighbor_warnings_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interval (in seconds) between repeated neighbor warning messages.").AddIntegerRangeDescription(1, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"stub_receive_only": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure the device as a receive-only EIGRP stub router.").String,
				Optional:            true,
			},
			"stub_advertise_connected": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("EIGRP stub router advertises connected routes.").String,
				Optional:            true,
			},
			"stub_advertise_redistributed": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("EIGRP stub router advertises redistributed routes.").String,
				Optional:            true,
			},
			"stub_advertise_static": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("EIGRP stub router advertises static routes.").String,
				Optional:            true,
			},
			"stub_advertise_summary": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("EIGRP stub router advertises summary routes.").String,
				Optional:            true,
			},
			"default_metric_bandwidth": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Default metric bandwidth (in Kbps) for redistributed routes.").AddIntegerRangeDescription(1, 4294967295).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4294967295),
				},
			},
			"default_metric_delay_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Default metric delay (in tens of microseconds) for redistributed routes.").AddIntegerRangeDescription(0, 4294967295).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 4294967295),
				},
			},
			"default_metric_reliability": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Default metric reliability for redistributed routes.").AddIntegerRangeDescription(0, 255).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 255),
				},
			},
			"default_metric_loading": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Default metric effective bandwidth (loading) for redistributed routes.").AddIntegerRangeDescription(1, 255).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
			},
			"default_metric_mtu": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Default metric path MTU (in bytes) for redistributed routes.").AddIntegerRangeDescription(1, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"cluster_ipv4_pool_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv4 Address Pool ID used for router ID assignment on clustered devices.").String,
				Optional:            true,
			},
		},
	}
}

func (r *DeviceEIGRPResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *DeviceEIGRPResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DeviceEIGRP

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
	body := plan.toBody(ctx, DeviceEIGRP{})
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

func (r *DeviceEIGRPResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DeviceEIGRP

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

func (r *DeviceEIGRPResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state DeviceEIGRP

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

func (r *DeviceEIGRPResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DeviceEIGRP

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
func (r *DeviceEIGRPResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Parse import ID
	var inputPattern = regexp.MustCompile(`^(?:(?P<domain>[^\s,]+),)?(?P<device_id>[^\s,]+),(?P<id>[^\s,]+?)$`)
	match := inputPattern.FindStringSubmatch(req.ID)
	if match == nil {
		errMsg := "Failed to parse import parameters.\nPlease provide import string in the following format: <domain>,<device_id>,<id>\n<domain> is optional. If not provided, `Global` is used implicitly and resource's `domain` attribute is not set.\n" + fmt.Sprintf("Got: %q", req.ID)
		resp.Diagnostics.AddError("Import error", errMsg)
		return
	}

	// Set domain, if provided
	if tmpDomain := match[inputPattern.SubexpIndex("domain")]; tmpDomain != "" {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("domain"), tmpDomain)...)
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), match[inputPattern.SubexpIndex("id")])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("device_id"), match[inputPattern.SubexpIndex("device_id")])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin createSubresources

// End of section. //template:end createSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin deleteSubresources

// End of section. //template:end deleteSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin updateSubresources

// End of section. //template:end updateSubresources
