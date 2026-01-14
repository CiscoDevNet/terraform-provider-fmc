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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type IKEv1Policy struct {
	Id                   types.String `tfsdk:"id"`
	Domain               types.String `tfsdk:"domain"`
	Name                 types.String `tfsdk:"name"`
	Description          types.String `tfsdk:"description"`
	Type                 types.String `tfsdk:"type"`
	Priority             types.Int64  `tfsdk:"priority"`
	EncryptionAlgorithm  types.String `tfsdk:"encryption_algorithm"`
	Hash                 types.String `tfsdk:"hash"`
	DhGroup              types.String `tfsdk:"dh_group"`
	Lifetime             types.Int64  `tfsdk:"lifetime"`
	AuthenticationMethod types.String `tfsdk:"authentication_method"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data IKEv1Policy) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/ikev1policies"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data IKEv1Policy) toBody(ctx context.Context, state IKEv1Policy) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.Priority.IsNull() {
		body, _ = sjson.Set(body, "priority", data.Priority.ValueInt64())
	}
	if !data.EncryptionAlgorithm.IsNull() {
		body, _ = sjson.Set(body, "encryption", data.EncryptionAlgorithm.ValueString())
	}
	if !data.Hash.IsNull() {
		body, _ = sjson.Set(body, "hash", data.Hash.ValueString())
	}
	if !data.DhGroup.IsNull() {
		body, _ = sjson.Set(body, "diffieHellmanGroup", data.DhGroup.ValueString())
	}
	if !data.Lifetime.IsNull() {
		body, _ = sjson.Set(body, "lifetimeInSeconds", data.Lifetime.ValueInt64())
	}
	if !data.AuthenticationMethod.IsNull() {
		body, _ = sjson.Set(body, "authenticationMethod", data.AuthenticationMethod.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *IKEv1Policy) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("priority"); value.Exists() {
		data.Priority = types.Int64Value(value.Int())
	} else {
		data.Priority = types.Int64Null()
	}
	if value := res.Get("encryption"); value.Exists() {
		data.EncryptionAlgorithm = types.StringValue(value.String())
	} else {
		data.EncryptionAlgorithm = types.StringNull()
	}
	if value := res.Get("hash"); value.Exists() {
		data.Hash = types.StringValue(value.String())
	} else {
		data.Hash = types.StringNull()
	}
	if value := res.Get("diffieHellmanGroup"); value.Exists() {
		data.DhGroup = types.StringValue(value.String())
	} else {
		data.DhGroup = types.StringNull()
	}
	if value := res.Get("lifetimeInSeconds"); value.Exists() {
		data.Lifetime = types.Int64Value(value.Int())
	} else {
		data.Lifetime = types.Int64Null()
	}
	if value := res.Get("authenticationMethod"); value.Exists() {
		data.AuthenticationMethod = types.StringValue(value.String())
	} else {
		data.AuthenticationMethod = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *IKEv1Policy) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("priority"); value.Exists() && !data.Priority.IsNull() {
		data.Priority = types.Int64Value(value.Int())
	} else {
		data.Priority = types.Int64Null()
	}
	if value := res.Get("encryption"); value.Exists() && !data.EncryptionAlgorithm.IsNull() {
		data.EncryptionAlgorithm = types.StringValue(value.String())
	} else {
		data.EncryptionAlgorithm = types.StringNull()
	}
	if value := res.Get("hash"); value.Exists() && !data.Hash.IsNull() {
		data.Hash = types.StringValue(value.String())
	} else {
		data.Hash = types.StringNull()
	}
	if value := res.Get("diffieHellmanGroup"); value.Exists() && !data.DhGroup.IsNull() {
		data.DhGroup = types.StringValue(value.String())
	} else {
		data.DhGroup = types.StringNull()
	}
	if value := res.Get("lifetimeInSeconds"); value.Exists() && !data.Lifetime.IsNull() {
		data.Lifetime = types.Int64Value(value.Int())
	} else {
		data.Lifetime = types.Int64Null()
	}
	if value := res.Get("authenticationMethod"); value.Exists() && !data.AuthenticationMethod.IsNull() {
		data.AuthenticationMethod = types.StringValue(value.String())
	} else {
		data.AuthenticationMethod = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *IKEv1Policy) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin Clone

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// End of section. //template:end toBodyNonBulk

// Section below is generated&owned by "gen/generator.go". //template:begin findObjectsToBeReplaced

// End of section. //template:end findObjectsToBeReplaced

// Section below is generated&owned by "gen/generator.go". //template:begin clearItemIds

// End of section. //template:end clearItemIds

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyPutDelete

// End of section. //template:end toBodyPutDelete

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBody

// End of section. //template:end adjustBody

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBodyBulk

// End of section. //template:end adjustBodyBulk
