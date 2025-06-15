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
	"fmt"
	"strings"
)

type AttributeDescription struct {
	String string
}

func NewAttributeDescription(s string) *AttributeDescription {
	return &AttributeDescription{s}
}

func (d *AttributeDescription) AddAttributeDescription(s string) *AttributeDescription {
	d.String = fmt.Sprintf("%s\n%s", d.String, s)
	return d
}

func (d *AttributeDescription) AddMinimumVersionHeaderDescription() *AttributeDescription {
	d.String = fmt.Sprintf("%s\n\nThe following restrictions apply:", d.String)
	return d
}

func (d *AttributeDescription) AddMinimumVersionAnyDescription() *AttributeDescription {
	d.String = fmt.Sprintf("%s\n  - Read operations are supported by any tested FMC version", d.String)
	return d
}

func (d *AttributeDescription) AddMinimumVersionDescription(minimumVersion string) *AttributeDescription {
	d.String = fmt.Sprintf("%s\n  - Minimum FMC version: `%s`", d.String, minimumVersion)
	return d
}

func (d *AttributeDescription) AddMinimumVersionCreateDescription(minimumVersion string) *AttributeDescription {
	d.String = fmt.Sprintf("%s\n  - Minimum FMC version for object management (Create/Update/Delete): `%s`", d.String, minimumVersion)
	return d
}

func (d *AttributeDescription) AddMinimumVersionBulkCreateDescription(minimumVersion string) *AttributeDescription {
	if minimumVersion == "999" {
		d.String = fmt.Sprintf("%s\n  - Bulk object creation is not supported by FMC, it will be handled one-by-one", d.String)
	} else {
		d.String = fmt.Sprintf("%s\n  - Minimum FMC version for bulk object creation: `%s`", d.String, minimumVersion)
	}
	return d
}

func (d *AttributeDescription) AddMinimumVersionBulkDeleteDescription(minimumVersion string) *AttributeDescription {
	if minimumVersion == "999" {
		d.String = fmt.Sprintf("%s\n  - Bulk object deletion is not supported by FMC, it will be handled one-by-one", d.String)
	} else {
		d.String = fmt.Sprintf("%s\n  - Minimum FMC version for bulk object deletion: `%s`", d.String, minimumVersion)
	}
	return d
}

func (d *AttributeDescription) AddMinimumVersionBulkUpdateDescription() *AttributeDescription {
	d.String = fmt.Sprintf("%s\n  - Updates are always done one-by-one.", d.String)
	return d
}

func (d *AttributeDescription) AddMinimumVersionBulkDisclaimerDescription() *AttributeDescription {
	d.String = fmt.Sprintf("%s\n  - If FMC version does not meet the minimum version requirement for bulk operations, this resource will automatically fall back to processing operations one-by-one.", d.String)
	return d
}

func (d *AttributeDescription) AddDefaultValueDescription(defaultValue string) *AttributeDescription {
	d.String = fmt.Sprintf("%s\n  - Default value: `%s`", d.String, defaultValue)
	return d
}

func (d *AttributeDescription) AddStringEnumDescription(values ...string) *AttributeDescription {
	v := make([]string, len(values))
	for i, value := range values {
		v[i] = fmt.Sprintf("`%s`", value)
	}
	d.String = fmt.Sprintf("%s\n  - Choices: %s", d.String, strings.Join(v, ", "))
	return d
}

func (d *AttributeDescription) AddIntegerRangeDescription(min, max int64) *AttributeDescription {
	d.String = fmt.Sprintf("%s\n  - Range: `%v`-`%v`", d.String, min, max)
	return d
}

func (d *AttributeDescription) AddFloatRangeDescription(min, max float64) *AttributeDescription {
	d.String = fmt.Sprintf("%s\n  - Range: `%v`-`%v`", d.String, min, max)
	return d
}
