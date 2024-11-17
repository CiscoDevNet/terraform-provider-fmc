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

package helpers

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func GetStringList(result []gjson.Result) types.List {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.StringValue(result[r].String())
	}
	return types.ListValueMust(types.StringType, v)
}

func GetInt64List(result []gjson.Result) types.List {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.Int64Value(result[r].Int())
	}
	return types.ListValueMust(types.Int64Type, v)
}

func GetStringSet(result []gjson.Result) types.Set {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.StringValue(result[r].String())
	}
	return types.SetValueMust(types.StringType, v)
}

func GetInt64Set(result []gjson.Result) types.Set {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.Int64Value(result[r].Int())
	}
	return types.SetValueMust(types.Int64Type, v)
}

// ToLower is the same as strings.ToLower, except it cares to not to convert null/unknown strings
// into empty strings.
func ToLower(s basetypes.StringValue) basetypes.StringValue {
	if s.IsUnknown() || s.IsNull() {
		return s
	}

	return types.StringValue(strings.ToLower(s.ValueString()))
}

// IsConfigUpdatingAt checks whether the attribute given by the Path is not Equal() between plan and state.
func IsConfigUpdatingAt(ctx context.Context, tfsdkPlan tfsdk.Plan, tfsdkState tfsdk.State, where path.Path) (bool, diag.Diagnostics) {
	var pv, sv attr.Value

	diags := tfsdkPlan.GetAttribute(ctx, where, &pv)
	if diags.HasError() {
		return false, diags
	}

	diags = tfsdkState.GetAttribute(ctx, where, &sv)
	if diags.HasError() {
		return false, nil
	}

	return !sv.Equal(pv), diags
}

// SetGjson conveniently wraps sjson.SetRaw, so that it acts on gjson.Result directly.
func SetGjson(orig gjson.Result, path string, content gjson.Result) gjson.Result {
	s, err := sjson.SetRaw(orig.String(), path, content.String())
	if err != nil {
		panic(err)
	}

	return gjson.Parse(s)
}

// DifferenceStringSet returns the elements that are present in `b`, but not in `a`.
func DifferenceStringSet(ctx context.Context, a basetypes.SetValue, b basetypes.SetValue) basetypes.SetValue {
	// extract elements from both sets
	var aElements []string
	var bElements []string
	a.ElementsAs(ctx, &aElements, false)
	b.ElementsAs(ctx, &bElements, false)

	diff := []attr.Value{}
	m := map[string]bool{}

	// Mark in `m` all elements from `a`
	for _, v := range aElements {
		m[v] = true
	}

	// Iterate over `b` to find elements not marked in `m`
	for _, v := range bElements {
		// If element is not in `m`, add it to the diff
		if _, ok := m[v]; !ok {
			diff = append(diff, types.StringValue(v))
		}
	}

	return types.SetValueMust(types.StringType, diff)
}
