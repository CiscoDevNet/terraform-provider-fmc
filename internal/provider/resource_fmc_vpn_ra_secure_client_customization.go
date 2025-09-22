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
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &VPNRASecureClientCustomizationResource{}
	_ resource.ResourceWithImportState = &VPNRASecureClientCustomizationResource{}
)

func NewVPNRASecureClientCustomizationResource() resource.Resource {
	return &VPNRASecureClientCustomizationResource{}
}

type VPNRASecureClientCustomizationResource struct {
	client *fmc.Client
}

func (r *VPNRASecureClientCustomizationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_ra_secure_client_customization"
}

func (r *VPNRASecureClientCustomizationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages FTD Remote Access (RA) Virtual Private Networks (VPNs) Secure Client Customizations.").String,

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
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this value is always 'RaVpnSecureClientCustomization'.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"gui_and_text_messages": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure language translation tables.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of Secure Client Customization - Language Translation.").String,
							Required:            true,
						},
					},
				},
			},
			"icons_and_images": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure icons and images for the secure client customization.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of Secure Client Customization - Icon/Image.").String,
							Required:            true,
						},
					},
				},
			},
			"scripts": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure Secure Client customization scripts.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of Secure Client Customization - Script.").String,
							Required:            true,
						},
					},
				},
			},
			"binaries": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure Secure Client Binary executable modules on device.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of Secure Client Customization - Binary.").String,
							Required:            true,
						},
					},
				},
			},
			"custom_installer_transforms": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure Secure Client customized installer transforms on device.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of Secure Client Customization - Custom Installer Transform.").String,
							Required:            true,
						},
					},
				},
			},
			"localized_installer_transforms": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure Secure Client localized installer transforms on device.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Id of Secure Client Customization - Localized Installer Transform.").String,
							Required:            true,
						},
					},
				},
			},
		},
	}
}

func (r *VPNRASecureClientCustomizationResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

func (r *VPNRASecureClientCustomizationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan VPNRASecureClientCustomization

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
	body := plan.toBody(ctx, VPNRASecureClientCustomization{})
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

func (r *VPNRASecureClientCustomizationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state VPNRASecureClientCustomization

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

func (r *VPNRASecureClientCustomizationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state VPNRASecureClientCustomization

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

func (r *VPNRASecureClientCustomizationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state VPNRASecureClientCustomization

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

func (r *VPNRASecureClientCustomizationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
