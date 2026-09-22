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

import (
	"testing"

	"github.com/tidwall/gjson"
)

// Shape of a real policy assignment, so that the path the parking logic relies on is
// caught if it ever changes. FMC is not consistent about the case of ids, hence the
// mixture of cases here.
const policyAssignment = `{"type":"PolicyAssignment",
 "policy":{"type":"AccessPolicy","name":"MyAccessPolicyName1","id":"005056B0-7B7F-0ed3-0000-012885486397"},
 "targets":[{"id":"95ade646-ebcf-11f0-bbd0-dd11de3ab65f","type":"Device","name":"ftd-02","keepLocalEvents":false},
            {"id":"40FA6CC4-B298-11F1-9882-026CD0381200","type":"DeviceHAPair","name":"ftd-ha"}]}`

func TestAssignedTargetIds(t *testing.T) {
	ids := assignedTargetIds(gjson.Parse(policyAssignment))

	if got, want := len(ids), 2; got != want {
		t.Fatalf("assigned targets = %d, want %d", got, want)
	}

	// Every target has to be found regardless of its type, as an HA Pair is assigned as a
	// whole and is not a device record
	for _, id := range []string{"95ade646-ebcf-11f0-bbd0-dd11de3ab65f", "40fa6cc4-b298-11f1-9882-026cd0381200"} {
		if _, found := ids[id]; !found {
			t.Errorf("target %s not found in %v", id, ids)
		}
	}

	// A target that FMC no longer reports on the policy must not be found, as that is what
	// stops it from being assigned to the after destroy policy
	if _, found := ids["972b12dc-ebcf-11f0-9a23-82f6023139ee"]; found {
		t.Error("ftd-01 found, although it is not assigned to this policy")
	}

	// An assignment without targets yields nothing, rather than everything
	if got := assignedTargetIds(gjson.Parse(`{"type":"PolicyAssignment"}`)); len(got) != 0 {
		t.Errorf("targets of an empty assignment = %d, want 0", len(got))
	}
}
