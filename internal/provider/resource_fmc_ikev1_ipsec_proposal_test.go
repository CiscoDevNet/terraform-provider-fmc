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

func TestAccFmcIKEv1IPsecProposal(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ikev1_ipsec_proposal.test", "name", "my_ikev1_ipsec_proposal"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ikev1_ipsec_proposal.test", "description", "IKEv1 IPsec Proposal 1"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_ikev1_ipsec_proposal.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ikev1_ipsec_proposal.test", "esp_encryption", "AES-192"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ikev1_ipsec_proposal.test", "esp_hash", "SHA"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcIKEv1IPsecProposalConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcIKEv1IPsecProposalConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_ikev1_ipsec_proposal.test",
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

func testAccFmcIKEv1IPsecProposalConfig_minimum() string {
	config := `resource "fmc_ikev1_ipsec_proposal" "test" {` + "\n"
	config += `	name = "my_ikev1_ipsec_proposal"` + "\n"
	config += `	esp_encryption = "AES-192"` + "\n"
	config += `	esp_hash = "SHA"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcIKEv1IPsecProposalConfig_all() string {
	config := `resource "fmc_ikev1_ipsec_proposal" "test" {` + "\n"
	config += `	name = "my_ikev1_ipsec_proposal"` + "\n"
	config += `	description = "IKEv1 IPsec Proposal 1"` + "\n"
	config += `	esp_encryption = "AES-192"` + "\n"
	config += `	esp_hash = "SHA"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
