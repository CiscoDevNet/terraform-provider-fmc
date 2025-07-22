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

func TestAccFmcSingleSignOnServer(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_single_sign_on_server.test", "name", "my_sso_server"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_single_sign_on_server.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_single_sign_on_server.test", "identity_provider_entity_id", "https://idp.example.com/saml"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_single_sign_on_server.test", "sso_url", "https://idp.example.com/sso"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_single_sign_on_server.test", "logout_url", "https://idp.example.com/logout"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_single_sign_on_server.test", "base_url", "https://vpn.example.com/callback"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_single_sign_on_server.test", "identity_provider_certificate_id", "123e4567-e89b-12d3-a456-426614174000"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_single_sign_on_server.test", "service_provider_certificate_id", "123e4567-e89b-12d3-a456-426614174001"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_single_sign_on_server.test", "request_signature_type", "RSA_SHA256"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_single_sign_on_server.test", "request_timeout", "300"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_single_sign_on_server.test", "identity_provider_accessible_only_on_internal_network", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_single_sign_on_server.test", "request_identity_provider_reauthentication_at_each_login", "false"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcSingleSignOnServerConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcSingleSignOnServerConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_single_sign_on_server.test",
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

func testAccFmcSingleSignOnServerConfig_minimum() string {
	config := `resource "fmc_single_sign_on_server" "test" {` + "\n"
	config += `	name = "my_sso_server"` + "\n"
	config += `	identity_provider_entity_id = "https://idp.example.com/saml"` + "\n"
	config += `	sso_url = "https://idp.example.com/sso"` + "\n"
	config += `	identity_provider_certificate_id = "123e4567-e89b-12d3-a456-426614174000"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcSingleSignOnServerConfig_all() string {
	config := `resource "fmc_single_sign_on_server" "test" {` + "\n"
	config += `	name = "my_sso_server"` + "\n"
	config += `	identity_provider_entity_id = "https://idp.example.com/saml"` + "\n"
	config += `	sso_url = "https://idp.example.com/sso"` + "\n"
	config += `	logout_url = "https://idp.example.com/logout"` + "\n"
	config += `	base_url = "https://vpn.example.com/callback"` + "\n"
	config += `	identity_provider_certificate_id = "123e4567-e89b-12d3-a456-426614174000"` + "\n"
	config += `	service_provider_certificate_id = "123e4567-e89b-12d3-a456-426614174001"` + "\n"
	config += `	request_signature_type = "RSA_SHA256"` + "\n"
	config += `	request_timeout = 300` + "\n"
	config += `	identity_provider_accessible_only_on_internal_network = false` + "\n"
	config += `	request_identity_provider_reauthentication_at_each_login = false` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
