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

func TestAccFmcRadiusExternalAuthenticationObject(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_radius_external_authentication_object.test", "name", "my_radius_auth_object"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_radius_external_authentication_object.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_radius_external_authentication_object.test", "description", "My RADIUS external authentication object"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_radius_external_authentication_object.test", "server_address", "10.1.1.10"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_radius_external_authentication_object.test", "server_port", "1812"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_radius_external_authentication_object.test", "timeout", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_radius_external_authentication_object.test", "retries", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_radius_external_authentication_object.test", "message_authenticator_enabled", "true"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcRadiusExternalAuthenticationObjectConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcRadiusExternalAuthenticationObjectConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_radius_external_authentication_object.test",
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

func testAccFmcRadiusExternalAuthenticationObjectConfig_minimum() string {
	config := `resource "fmc_radius_external_authentication_object" "test" {` + "\n"
	config += `	name = "my_radius_auth_object"` + "\n"
	config += `	server_address = "10.1.1.10"` + "\n"
	config += `	key = "my_secret_key"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcRadiusExternalAuthenticationObjectConfig_all() string {
	config := `resource "fmc_radius_external_authentication_object" "test" {` + "\n"
	config += `	name = "my_radius_auth_object"` + "\n"
	config += `	description = "My RADIUS external authentication object"` + "\n"
	config += `	server_address = "10.1.1.10"` + "\n"
	config += `	server_port = "1812"` + "\n"
	config += `	key = "my_secret_key"` + "\n"
	config += `	timeout = 30` + "\n"
	config += `	retries = 3` + "\n"
	config += `	message_authenticator_enabled = true` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
