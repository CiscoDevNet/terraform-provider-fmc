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

func TestAccFmcDistinguishedName(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_distinguished_name.test", "name", "my_distinguished_name"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_distinguished_name.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_distinguished_name.test", "distinguished_name", "C=US,CN=example.com,O=Example,OU=IT"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcDistinguishedNameConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcDistinguishedNameConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_distinguished_name.test",
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

func testAccFmcDistinguishedNameConfig_minimum() string {
	config := `resource "fmc_distinguished_name" "test" {` + "\n"
	config += `	name = "my_distinguished_name"` + "\n"
	config += `	distinguished_name = "C=US,CN=example.com,O=Example,OU=IT"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcDistinguishedNameConfig_all() string {
	config := `resource "fmc_distinguished_name" "test" {` + "\n"
	config += `	name = "my_distinguished_name"` + "\n"
	config += `	distinguished_name = "C=US,CN=example.com,O=Example,OU=IT"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// The test below is hand-owned (it lives outside of any generator section so it survives
// `go generate`).

// TestAccFmcDistinguishedName_reorder verifies the distinguished-name reconciliation logic.
// FMC stores the distinguished name in a canonicalized, slash-separated attribute order
// (e.g. "C=US/CN=example.com/O=Example/OU=IT") regardless of the order supplied in the
// request. When the user provides the components in a different order, the provider must
// recognize that the two values describe the same set of components and preserve the user's
// ordering, so that no persistent diff is reported. The test applies a non-canonical order
// and relies on the testing framework's post-apply empty-plan check to prove stability.
func TestAccFmcDistinguishedName_reorder(t *testing.T) {
	const reordered = "OU=IT,O=Example,CN=example.com,C=US"

	config := `resource "fmc_distinguished_name" "reorder" {` + "\n"
	config += `	name = "my_distinguished_name_reorder"` + "\n"
	config += `	distinguished_name = "` + reordered + `"` + "\n"
	config += `}` + "\n"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			// A non-empty plan after apply (the framework checks this automatically) would
			// fail the step, proving the reconciliation suppresses the re-ordering diff.
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("fmc_distinguished_name.reorder", "distinguished_name", reordered),
				),
			},
			// Re-apply the same config explicitly as a plan-only step to double-check no diff.
			{
				Config:   config,
				PlanOnly: true,
			},
		},
	})
}
