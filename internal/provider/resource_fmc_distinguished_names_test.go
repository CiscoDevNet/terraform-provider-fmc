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

func TestAccFmcDistinguishedNames(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_distinguished_names.test", "items.my_distinguished_names.id"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_distinguished_names.test", "items.my_distinguished_names.type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_distinguished_names.test", "items.my_distinguished_names.distinguished_name", "C=US,CN=example.com,O=Example,OU=IT"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcDistinguishedNamesConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcDistinguishedNamesConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
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

func testAccFmcDistinguishedNamesConfig_minimum() string {
	config := `resource "fmc_distinguished_names" "test" {` + "\n"
	config += `	items = { "my_distinguished_names" = {` + "\n"
	config += `		distinguished_name = "C=US,CN=example.com,O=Example,OU=IT"` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcDistinguishedNamesConfig_all() string {
	config := `resource "fmc_distinguished_names" "test" {` + "\n"
	config += `	items = { "my_distinguished_names" = {` + "\n"
	config += `		distinguished_name = "C=US,CN=example.com,O=Example,OU=IT"` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// The test below is hand-owned (it lives outside of any generator section so it survives
// `go generate`).

// TestAccFmcDistinguishedNames_reorder verifies the distinguished-name reconciliation logic
// for the bulk (plural) resource across multiple items, each carrying a different distinguished
// name supplied in a non-canonical attribute order. FMC stores each distinguished name in a
// canonicalized, slash-separated attribute order regardless of the order supplied in the
// request. The provider must reconcile each item independently and preserve the user's ordering,
// so that no persistent diff is reported. The test relies on the testing framework's post-apply
// empty-plan check (plus an explicit plan-only step) to prove per-item stability.
func TestAccFmcDistinguishedNames_reorder(t *testing.T) {
	// Each item uses a different set of components in a different, non-canonical order.
	items := map[string]string{
		"my_distinguished_names_reorder_1": "OU=IT,O=Example,CN=example.com,C=US",
		"my_distinguished_names_reorder_2": "O=Acme,C=PL,CN=host.acme.com",
		"my_distinguished_names_reorder_3": "CN=single.example.net",
	}

	config := `resource "fmc_distinguished_names" "reorder" {` + "\n"
	config += `	items = {` + "\n"
	for name, dn := range items {
		config += `		"` + name + `" = {` + "\n"
		config += `			distinguished_name = "` + dn + `"` + "\n"
		config += `		}` + "\n"
	}
	config += `	}` + "\n"
	config += `}` + "\n"

	var checks []resource.TestCheckFunc
	for name, dn := range items {
		checks = append(checks, resource.TestCheckResourceAttr("fmc_distinguished_names.reorder", "items."+name+".distinguished_name", dn))
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			// A non-empty plan after apply (the framework checks this automatically) would
			// fail the step, proving the reconciliation suppresses the re-ordering diff for
			// every item.
			{
				Config: config,
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			// Re-apply the same config explicitly as a plan-only step to double-check no diff.
			{
				Config:   config,
				PlanOnly: true,
			},
		},
	})
}
