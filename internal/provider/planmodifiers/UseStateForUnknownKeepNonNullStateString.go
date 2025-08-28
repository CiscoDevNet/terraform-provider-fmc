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

package planmodifiers

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

//
// terraform-plugin-framework v1.15.1 changed behaviour of UseStateForUnknown plan modifier.
// Before the change, if new nested map, that contains Computed/Unknown elements is added to already created resource,
// those elements are treated as Unknown.
// After the change, those elements are treated as Null.
// This breaks bulk operations, as the ID of newly created objects can never be set, as Terraform would expects them to be Null.
// Ref: https://github.com/hashicorp/terraform-plugin-framework/issues/1197
//

var _ planmodifier.String = useStateForUnknownKeepNonNullStateString{}

func UseStateForUnknownKeepNonNullStateString() planmodifier.String {
	return useStateForUnknownKeepNonNullStateString{}
}

type useStateForUnknownKeepNonNullStateString struct{}

func (m useStateForUnknownKeepNonNullStateString) Description(_ context.Context) string {
	return "Once set to a non-null value, the value of this attribute in state will not change. Attributes of new nested objects, that did not exist on create will default to Unknown."
}

func (m useStateForUnknownKeepNonNullStateString) MarkdownDescription(_ context.Context) string {
	return "Once set to a non-null value, the value of this attribute in state will not change. Attributes of new nested objects, that did not exist on create will default to Unknown."
}

func (m useStateForUnknownKeepNonNullStateString) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// Do nothing if there the state value is null.
	if req.StateValue.IsNull() {
		return
	}

	// Do nothing if there is a known planned value.
	if !req.PlanValue.IsUnknown() {
		return
	}

	// Do nothing if there is an unknown configuration value, otherwise interpolation gets messed up.
	if req.ConfigValue.IsUnknown() {
		return
	}

	resp.PlanValue = req.StateValue
}
