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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccFmcTimeRange(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_time_range.test", "name", "my_time_range"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_time_range.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_time_range.test", "description", "My time range object"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_time_range.test", "start_time", "2025-01-07T20:20"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_time_range.test", "end_time", "2025-01-22T09:20"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_time_range.test", "recurrence_list.0.recurrence_type", "RANGE"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_time_range.test", "recurrence_list.0.range_start_time", "09:00"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_time_range.test", "recurrence_list.0.range_end_time", "17:00"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_time_range.test", "recurrence_list.0.range_start_day", "MON"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_time_range.test", "recurrence_list.0.range_end_day", "FRI"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcTimeRangeConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcTimeRangeConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_time_range.test",
		ImportState:  true,
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcTimeRangeConfig_minimum() string {
	config := `resource "fmc_time_range" "test" {` + "\n"
	config += `	name = "my_time_range"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcTimeRangeConfig_all() string {
	config := `resource "fmc_time_range" "test" {` + "\n"
	config += `	name = "my_time_range"` + "\n"
	config += `	description = "My time range object"` + "\n"
	config += `	start_time = "2025-01-07T20:20"` + "\n"
	config += `	end_time = "2025-01-22T09:20"` + "\n"
	config += `	recurrence_list = [{` + "\n"
	config += `		recurrence_type = "RANGE"` + "\n"
	config += `		range_start_time = "09:00"` + "\n"
	config += `		range_end_time = "17:00"` + "\n"
	config += `		range_start_day = "MON"` + "\n"
	config += `		range_end_day = "FRI"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
