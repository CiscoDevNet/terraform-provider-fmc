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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource

func TestAccDataSourceFmcIKEv1Policies(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_ikev1_policies.test", "items.my_ikev1_policies.id"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ikev1_policies.test", "items.my_ikev1_policies.description", "IKEv1 Policy 1"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_ikev1_policies.test", "items.my_ikev1_policies.type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ikev1_policies.test", "items.my_ikev1_policies.priority", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ikev1_policies.test", "items.my_ikev1_policies.encryption", "AES-192"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ikev1_policies.test", "items.my_ikev1_policies.hash", "SHA"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ikev1_policies.test", "items.my_ikev1_policies.dh_group", "14"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ikev1_policies.test", "items.my_ikev1_policies.lifetime", "86400"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ikev1_policies.test", "items.my_ikev1_policies.authentication_method", "Preshared Key"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcIKEv1PoliciesConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcIKEv1PoliciesConfig() string {
	config := `resource "fmc_ikev1_policies" "test" {` + "\n"
	config += `	items = { "my_ikev1_policies" = {` + "\n"
	config += `		description = "IKEv1 Policy 1"` + "\n"
	config += `		priority = 10` + "\n"
	config += `		encryption = "AES-192"` + "\n"
	config += `		hash = "SHA"` + "\n"
	config += `		dh_group = "14"` + "\n"
	config += `		lifetime = 86400` + "\n"
	config += `		authentication_method = "Preshared Key"` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_ikev1_policies" "test" {
			depends_on = [fmc_ikev1_policies.test]
			items = {
				"my_ikev1_policies" = {
				}
			}
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
