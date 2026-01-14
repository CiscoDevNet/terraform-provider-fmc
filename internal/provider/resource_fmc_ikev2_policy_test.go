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

func TestAccFmcIKEv2Policy(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ikev2_policy.test", "name", "my_ikev2_policy"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ikev2_policy.test", "description", "IKEv2 Policy"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_ikev2_policy.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ikev2_policy.test", "priority", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ikev2_policy.test", "lifetime", "86400"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcIKEv2PolicyConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcIKEv2PolicyConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_ikev2_policy.test",
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

func testAccFmcIKEv2PolicyConfig_minimum() string {
	config := `resource "fmc_ikev2_policy" "test" {` + "\n"
	config += `	name = "my_ikev2_policy"` + "\n"
	config += `	priority = 10` + "\n"
	config += `	lifetime = 86400` + "\n"
	config += `	integrity_algorithms = ["SHA-256"]` + "\n"
	config += `	encryption_algorithms = ["AES-256"]` + "\n"
	config += `	prf_algorithms = ["SHA-256"]` + "\n"
	config += `	dh_groups = ["14"]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcIKEv2PolicyConfig_all() string {
	config := `resource "fmc_ikev2_policy" "test" {` + "\n"
	config += `	name = "my_ikev2_policy"` + "\n"
	config += `	description = "IKEv2 Policy"` + "\n"
	config += `	priority = 10` + "\n"
	config += `	lifetime = 86400` + "\n"
	config += `	integrity_algorithms = ["SHA-256"]` + "\n"
	config += `	encryption_algorithms = ["AES-256"]` + "\n"
	config += `	prf_algorithms = ["SHA-256"]` + "\n"
	config += `	dh_groups = ["14"]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
