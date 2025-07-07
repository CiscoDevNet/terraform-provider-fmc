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
	"slices"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/google/uuid"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource = &NetworkGroupsResource{}
)

func NewNetworkGroupsResource() resource.Resource {
	return &NetworkGroupsResource{}
}

type NetworkGroupsResource struct {
	client *fmc.Client
}

func (r *NetworkGroupsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_groups"
}

func (r *NetworkGroupsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages Network Groups through bulk operations.").AddMinimumVersionHeaderDescription().AddMinimumVersionBulkDeleteDescription("7.4").AddMinimumVersionBulkDisclaimerDescription().String,

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
			"items": schema.MapNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Map of network groups. The key of the map is the name of the individual Network Group.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of the managed Network Group.").String,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"description": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Description of the object.").String,
							Optional:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this value is always 'NetworkGroup'.").String,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"overridable": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Indicates whether object values can be overridden.").String,
							Optional:            true,
						},
						"network_groups": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of names (not Ids) of child Network Groups. The names must be defined in the same instance of fmc_network_groups resource. This is an auxiliary way to add a child Network Group: the suggested way is to instead add it inside `objects` by its Ids.").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of network objects (Hosts, Networs, Ranges or FQDNs).").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Id of the network object.").String,
										Optional:            true,
									},
								},
							},
						},
						"literals": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of literal values.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"value": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("IP address or network in CIDR format. Please do not use /32 mask for host.").String,
										Optional:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *NetworkGroupsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// networkGroup is an internal representation of a single fmc_network_group.
type networkGroup struct {
	name     string   // name of the network group
	children []string // slice of children network group names
	json     string   // raw JSON of the network group
	bulk     int      // ID of the bulk in which the group will be created
}

// networkGroupsBulk is a bulk of network groups to be created in one POST request.
type networkGroupsBulk struct {
	groups []networkGroup
}

// networkGroupsBulkDelete is a bulk of network groups to be deleted in one DELETE request.
type networkGroupsBulkDelete struct {
	ids   string
	names []string
}

func (r *NetworkGroupsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan NetworkGroups

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

	body := plan.toBody(ctx, NetworkGroups{})
	plan.Id = types.StringValue(uuid.New().String())

	state := plan
	if len(plan.Items) > 0 {
		state.Items = map[string]NetworkGroupsItems{}
	}

	state, diags = r.updateSubresources(ctx, req.Plan, plan, body, tfsdk.State{}, state, reqMods...)
	resp.Diagnostics.Append(diags...)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished", state.Id.ValueString()))

	// On error we do Set anyway. Terraform taints our resource, and the next run is responsible to call Delete() for us.
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *NetworkGroupsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state NetworkGroups

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

	// Get all Network Groups.
	res, err := r.client.Get(state.getPath()+"?expanded=true", reqMods...)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects (GET), got error: %s", err))
		return
	}

	// Modify API response, by moving child network groups from "objects" to "network_groups" attribute.
	res = synthesizeNetworkGroups(ctx, res, &state)

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

// synthesizeNetworkGroups takes a real API Result (json) and converts some of the entries of the original attribute "objects"
// into synthetic attribute "network_groups". It returns a modified json.
func synthesizeNetworkGroups(ctx context.Context, res gjson.Result, state *NetworkGroups) gjson.Result {
	items := `[]`
	if !res.Get("items").IsArray() {
		return res
	}

	ownedIds := map[string]string{}
	for name, item := range state.Items {
		if item.Id.IsUnknown() || item.Id.IsNull() {
			continue
		}
		ownedIds[item.Id.ValueString()] = name
	}

	for _, item := range res.Get("items").Array() {
		item := synthesizeNetworkGroupsItem(ctx, item, ownedIds)
		items, _ = sjson.SetRaw(items, "-1", item)
	}

	return helpers.SetGjson(res, "items", gjson.Parse(items))
}

// synthesizeNetworkGroupsItem takes a single item from the API response and converts some of the entries of the original attribute "objects"
// into synthetic attribute "network_groups".
func synthesizeNetworkGroupsItem(ctx context.Context, item gjson.Result, ownedIds map[string]string) string {
	ret := item.String()
	if _, owned := ownedIds[item.Get("id").String()]; !owned {
		return ret
	}

	ret, _ = sjson.Delete(ret, "objects")
	for _, obj := range item.Get("objects").Array() {
		name, owned := ownedIds[obj.Get("id").String()]

		if owned && strings.ToLower(obj.Get("type").String()) == "networkgroup" {
			tflog.Debug(ctx, fmt.Sprintf("%s: child %q: adding it to network_groups and removing it from objects: %s",
				item.Get("id").String(), name, obj.String()))

			ret, _ = sjson.Set(ret, "network_groups.-1", name)
		} else {
			ret, _ = sjson.SetRaw(ret, "objects.-1", obj.String())
		}
	}

	return ret
}

func (r *NetworkGroupsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state NetworkGroups

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

	body := plan.toBody(ctx, state)

	//orig := state
	//state = plan
	//state.Items = orig.Items

	state, diags = r.updateSubresources(ctx, req.Plan, plan, body, req.State, state, reqMods...)
	resp.Diagnostics.Append(diags...)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworkGroupsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworkGroups

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

	state, diags = r.updateSubresources(ctx, tfsdk.Plan{}, NetworkGroups{}, "{}", req.State, state, reqMods...)
	resp.Diagnostics.Append(diags...)

	diags = resp.State.Set(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
	tflog.Debug(ctx, fmt.Sprintf("%s: Delete successful", state.Id.ValueString()))
}

// Section below is generated&owned by "gen/generator.go". //template:begin import
// End of section. //template:end import

// updateSubresource creates, updates and deletes subresources of the Network Groups resource.
// updateSubresources returns a coherent state whether it fails or succeeds. Caller should always persist that state
// into the Response (UpdateResponse, CreateResponse, ...), otherwise the API's UUIDs may go out-of-sync with
// terraform.tfstate, an immediate user-visible bug.
func (r *NetworkGroupsResource) updateSubresources(ctx context.Context, tfsdkPlan tfsdk.Plan, plan NetworkGroups, planBody string, tfsdkState tfsdk.State, state NetworkGroups, reqMods ...func(*fmc.Req)) (NetworkGroups, diag.Diagnostics) {
	// Sort objects in a way that any child comes before its parent. 'seq' is sorted list.
	seq, diags := graphTopologicalSeq(ctx, planBody)
	if diags.HasError() {
		return state, diags
	}

	// Group objects into bulks to be created (bulks) and objects to be updated (seq)
	bulks, seq := divideToBulks(ctx, seq, plan)
	if diags.HasError() {
		return state, diags
	}

	// Log bulks for debugging purposes.
	for _, bulk := range bulks {
		readable := slices.Clone(bulk.groups)
		for i := range readable {
			readable[i].json = ""
		}
		tflog.Debug(ctx, fmt.Sprintf("%s: bulk ordered for Create: %+v", plan.Id.ValueString(), readable))
	}

	// Create subresources in bulks.
	for _, bulk := range bulks {
		state, diags = bulk.Create(ctx, plan, state, r.client, reqMods...)
		if diags.HasError() {
			return state, diags
		}

		// Put IDs of created subresources (individual network groups) into the plan. This is needed for successful resolution of
		// the children network groups IDs in the following group.Body() call.
		for _, group := range bulk.groups {
			tmp := plan.Items[group.name]
			tmp.Id = state.Items[group.name].Id
			plan.Items[group.name] = tmp
		}
	}

	// Subresources to Update.
	tflog.Debug(ctx, fmt.Sprintf("%s: considering remaining subresources for Update: %+v", plan.Id.ValueString(), seq))
	for _, group := range seq {
		// Check if the subresource has changed.
		ok, diags := helpers.IsConfigUpdatingAt(ctx, tfsdkPlan, tfsdkState, path.Root("items").AtMapKey(group.name))
		if diags.HasError() {
			return state, diags
		}

		// If the subresource is not changing, skip it.
		if !ok {
			continue
		}

		updating := plan.Items[group.name].Id.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("%s: Subresource %s: Beginning Update", updating, group.name))

		// Generate the body for the subresource. This includes the children network groups ID resolution.
		body, diags := group.Body(ctx, plan)
		if diags.HasError() {
			return state, diags
		}

		res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(updating), body, reqMods...)
		if err != nil {
			return state, diag.Diagnostics{
				diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String())),
			}
		}

		// Update the state with plan for updated subresource.
		state.Items[group.name] = plan.Items[group.name]

		tflog.Debug(ctx, fmt.Sprintf("%s: Subresource %s: Update finished successfully", updating, group.name))
	}

	// Subresources to Delete.
	stateBody := state.toBody(ctx, NetworkGroups{})
	// Sort objects in a way that any child comes before its parent. 'seq' is sorted list.
	delSeq, diags := graphTopologicalSeq(ctx, stateBody)
	if diags.HasError() {
		return state, diags
	}

	// Get FMC version from the client
	fmcVersion, _ := version.NewVersion(strings.Split(r.client.FMCVersion, " ")[0])
	// Check if FMC version supports bulk deletes
	if fmcVersion.LessThan(minFMCVersionBulkDeleteNetworkGroups) {
		// Traverse delSeq in reverse order, so that any parent comes before its children.
		for i := len(delSeq) - 1; i >= 0; i-- {
			gn := delSeq[i].name
			if _, found := plan.Items[gn]; found {
				// item present both in state and in plan, do not delete
				continue
			}

			deleting := state.Items[gn].Id.ValueString()
			tflog.Debug(ctx, fmt.Sprintf("%s: Subresource %s: Beginning Delete", deleting, gn))

			res, err := r.client.Delete(state.getPath()+"/"+url.QueryEscape(deleting), reqMods...)
			if err != nil {
				return state, diag.Diagnostics{
					diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String())),
				}
			}

			delete(state.Items, gn)
			tflog.Debug(ctx, fmt.Sprintf("%s: Subresource %s: Delete finished successfully", deleting, gn))
		}
	} else {
		var idsToRemove strings.Builder
		var namesToRemove []string
		var deleteGroups []networkGroupsBulkDelete
		tflog.Debug(ctx, fmt.Sprintf("%s: Bulk Delete of subresources: Beginning", plan.Id.ValueString()))
		// Traverse delSeq in reverse order, so that any parent comes before its children.
		for i := len(delSeq) - 1; i >= 0; i-- {
			gn := delSeq[i].name
			if _, found := plan.Items[gn]; !found {
				// item not present in plan, so it is to be deleted

				// Prepare URL filter for bulk delete.
				idsToRemove.WriteString(url.QueryEscape(state.Items[gn].Id.ValueString()) + ",")

				// Save names of the groups to be deleted, so that we can remove them from the state.
				namesToRemove = append(namesToRemove, gn)
			}

			// If this is last element, or the following element is in a different bulk, or the URL parameter length exceeds the limit, AND there are elements to delete
			if (i == 0 || delSeq[i].bulk != delSeq[i-1].bulk || len(idsToRemove.String()) >= maxUrlParamLength) && len(idsToRemove.String()) > 0 {
				// Create new bulk group for deletion
				deleteGroups = append(deleteGroups, networkGroupsBulkDelete{
					ids:   idsToRemove.String(),
					names: slices.Clone(namesToRemove),
				})

				// Reset filter URL
				idsToRemove.Reset()

				// Reset the slice
				namesToRemove = namesToRemove[:0]
			}
		}

		// Process each bulk group for deletion
		for _, group := range deleteGroups {
			urlPath := state.getPath() + "?bulk=true&filter=ids:" + url.QueryEscape(group.ids)
			res, err := r.client.Delete(urlPath, reqMods...)
			if err != nil {
				return state, diag.Diagnostics{
					diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String())),
				}
			}

			// Remove groups from state
			for _, name := range group.names {
				delete(state.Items, name)
			}
		}

		tflog.Debug(ctx, fmt.Sprintf("%s: Bulk Delete of subresources: finished successfully", plan.Id.ValueString()))
	}

	return state, nil
}

// graphTopologicalSeq takes "items" of the body and parses their parent-child dependencies (attribute "network_groups").
// The goal is to ensure that any child is created before its parent.
// Having the "items" as a graph the func runs the topological sort algorithm to convert it to a sequence:
// https://en.wikipedia.org/wiki/Topological_sorting
//
// Returns ordered sequence of networkGroup objects, with indication (bulk attribute) of which bulk group the object belongs to.
//
// And if you iterate the result sequence in reverse, any parent is guaranteed to be placed before its children, which
// is useful for delete operations.
func graphTopologicalSeq(ctx context.Context, body string) ([]networkGroup, diag.Diagnostics) {
	b := gjson.Parse(body).Get("items")
	m := map[string]networkGroup{}
	parentCount := map[string]int{}
	for _, item := range b.Array() {
		g := networkGroup{
			name: item.Get("name").String(),
			json: item.String(),
		}
		for _, child := range item.Get("network_groups").Array() {
			g.children = append(g.children, child.String())
		}
		m[g.name] = g
		parentCount[g.name] = 0
	}

	for k := range m {
		for _, child := range m[k].children {
			parentCount[child]++
		}
	}

	var curr []networkGroup
	for k := range parentCount {
		if parentCount[k] == 0 {
			delete(parentCount, k)
			curr = append(curr, m[k])
		}
	}

	var diags diag.Diagnostics
	var ret []networkGroup
	for bulk := 1; len(curr) > 0; bulk++ {
		next := []networkGroup{}
		for _, group := range curr {
			for _, child := range group.children {
				parentCount[child]--
				if parentCount[child] == 0 {
					delete(parentCount, child)
					next = append(next, m[child])
				}
			}
			group.bulk = bulk
			ret = append(ret, group)
		}
		curr = next
	}

	slices.Reverse(ret)

	var cycle []string
	for k := range parentCount {
		cycle = append(cycle, k)
	}

	if len(cycle) > 0 {
		diags.AddAttributeError(
			path.Root("items"),
			"Cycle in network_groups",
			fmt.Sprintf("Children contained in network_groups must not be their own ancestors: %+v", cycle),
		)

		return ret, diags
	}

	return ret, diags
}

// Body generates the body of the network group, which includes refresh of the children network group IDs.
func (group *networkGroup) Body(ctx context.Context, state NetworkGroups) (string, diag.Diagnostics) {
	ret := group.json
	ret, _ = sjson.Delete(ret, "network_groups")

	for _, child := range group.children {
		existing := state.Items[child].Id
		if existing.IsUnknown() {
			return "", diag.Diagnostics{
				diag.NewErrorDiagnostic("Internal Error", "bug in topological sort"),
			}
		}
		tflog.Debug(ctx, fmt.Sprintf("%s: appending child group %s (%s) to objects", group.name, child, existing))

		obj := "{}"
		obj, _ = sjson.Set(obj, "id", existing.ValueString())
		obj, _ = sjson.Set(obj, "type", "AnyNonEmptyString")

		ret, _ = sjson.SetRaw(ret, "objects.-1", obj)
	}

	return ret, nil
}

// divideToBulks takes seq (ordered list of network groups to be created) and divides it into bulks (ret) to be created (bulk-POST), and leftovers (for one-by-one PUTs),
// as there is no bulk-PUT in this API).
func divideToBulks(ctx context.Context, seq []networkGroup, plan NetworkGroups) (ret []networkGroupsBulk, leftovers []networkGroup) {
	var g []networkGroup
	for _, group := range seq {
		// If the group has ID set, it means it's subject to update, not create. Putting that in leftovers.
		if !plan.Items[group.name].Id.IsUnknown() {
			leftovers = append(leftovers, group)
			continue
		}
		// Otherwise, we need to create it.
		g = append(g, group)
	}

	b := networkGroupsBulk{}
	for i := range g {
		// Add group to the current bulk.
		b.groups = append(b.groups, g[i])
		// If this is last element, or the following element is in a different bulk, or number of elements in the bulk group exceeds the limit,
		// then we need to create a new bulk group.
		if i == len(g)-1 || g[i].bulk != g[i+1].bulk || len(b.groups) >= bulkSizeCreate {
			ret = append(ret, b)
			b = networkGroupsBulk{}
		}
	}

	return ret, leftovers
}

// Create creates a bulk of network groups in one POST request.
func (bulk *networkGroupsBulk) Create(ctx context.Context, plan, state NetworkGroups, client *fmc.Client, reqMods ...func(*fmc.Req)) (NetworkGroups, diag.Diagnostics) {
	ret := state.Clone()
	bodies := "[]"
	for i := range bulk.groups {
		body, diags := bulk.groups[i].Body(ctx, state)
		if diags.HasError() {
			return ret, diags
		}

		bodies, _ = sjson.SetRaw(bodies, "-1", body)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Bulk of subresources: Beginning Create", plan.Id.ValueString()))

	res, err := client.Post(plan.getPath()+"?bulk=true", bodies, reqMods...)
	if err != nil {
		return ret, diag.Diagnostics{
			diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Failed to create a bulk (POST), got error: %s, %s", err, res.String())),
		}
	}

	if ret.Items == nil && len(bulk.groups) != 0 {
		ret.Items = map[string]NetworkGroupsItems{}
	}

	// Bulk Create is all-or-nothing, so now persist all in the tfstate.
	for _, g := range bulk.groups {
		ret.Items[g.name] = plan.Items[g.name]
	}

	ret.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Bulk of subresources: Create finished successfully", plan.Id.ValueString()))

	return ret, nil
}
