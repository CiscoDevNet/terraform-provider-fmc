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

func TestAccFmcCipherSuiteList(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_cipher_suite_list.test", "name", "my_cipher_suite_list"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_cipher_suite_list.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_cipher_suite_list.test", "cipher_suites.0.name", "TLS_RSA_WITH_AES_128_CBC_SHA"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcCipherSuiteListConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcCipherSuiteListConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_cipher_suite_list.test",
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

func testAccFmcCipherSuiteListConfig_minimum() string {
	config := `resource "fmc_cipher_suite_list" "test" {` + "\n"
	config += `	name = "my_cipher_suite_list"` + "\n"
	config += `	cipher_suites = [{` + "\n"
	config += `		name = "TLS_RSA_WITH_AES_128_CBC_SHA"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcCipherSuiteListConfig_all() string {
	config := `resource "fmc_cipher_suite_list" "test" {` + "\n"
	config += `	name = "my_cipher_suite_list"` + "\n"
	config += `	cipher_suites = [{` + "\n"
	config += `		name = "TLS_RSA_WITH_AES_128_CBC_SHA"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
